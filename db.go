package main

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Contact struct {
	Name  string `bson:"name"`
	Email string `bson:"email"`
}

func openDB(uri string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		client.Disconnect(ctx)
		return nil, err
	}

	return client, nil
}

func ensureContactIndex(db *mongo.Database) error {
	ctx := context.Background()
	collection := db.Collection("contacts")

	indexModel := mongo.IndexModel{
		Keys:    bson.M{"email": 1},
		Options: options.Index().SetUnique(true),
	}
	_, err := collection.Indexes().CreateOne(ctx, indexModel)
	return err
}

func loadRecipientsFromDB(db *mongo.Database, ch chan Recipient) error {
	defer close(ch)

	ctx := context.Background()
	collection := db.Collection("contacts")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var contact Contact
		if err := cursor.Decode(&contact); err != nil {
			return err
		}
		ch <- Recipient{
			Name:  contact.Name,
			Email: contact.Email,
		}

	}
	return cursor.Err()
}
