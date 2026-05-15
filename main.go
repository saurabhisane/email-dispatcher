package main

import (
	"bytes"
	"context"
	"flag"
	"html/template"
	"log"
	"sync"
)

type Recipient struct {
	Name  string
	Email string
}

func main() {

	mongoURI := flag.String("mongoURI", "mongodb://localhost:27017", "MongoDB URI")
	dbName := flag.String("dbName", "email_dispatcher", "Email-dispatcher database name")
	workerCount := flag.Int("workerCount", 5, "Number of email worker goroutines")
	flag.Parse()

	client, err := openDB(*mongoURI)
	if err != nil {
		log.Fatalf("Failed to connect to mongoDB: %v", err)
	}

	defer func() {
		if err = client.Disconnect(context.Background()); err != nil {
			log.Fatalf("Failed to disconnect from mongoDB: %v", err)
		}
	}()

	db := client.Database(*dbName)

	if err := ensureContactIndex(db); err != nil {
		log.Fatalf("Failed to create index on contacts collection: %v", err)
	}

	recipientChannel := make(chan Recipient)

	go func() {
		if err := loadRecipientsFromDB(db, recipientChannel); err != nil{
			log.Fatalf("Failed to load recipients from database: %v", err)
		}
	}()

	var wg sync.WaitGroup

	for i := 1; i <= *workerCount; i++ {
		wg.Add(1)
		go emailWorker(i, recipientChannel, &wg)
	}

	wg.Wait()
}

func executeTemplate(r Recipient) (string, error) {
	t, err := template.ParseFiles("email.tmpl")
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer
	err = t.Execute(&tpl, r)
	if err != nil {
		return "", err
	}
	return tpl.String(), nil
}
