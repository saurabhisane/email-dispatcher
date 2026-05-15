package main

import (
	"fmt"
	"log"
	"net/smtp"
	"sync"
	"time"
)

func emailWorker(id int, ch chan Recipient, wg *sync.WaitGroup) {
	defer wg.Done()

	for recipient := range ch {
		smtpHost := "localhost"
		smtpPort := "1025"

		msg, err := executeTemplate(recipient)
		if err != nil {
			log.Printf("Worker %d: Error executing template for %s: %v\n", id, recipient.Email, err)
			continue
		}

		fmt.Printf("Worker %d sending email to %s\n", id, recipient.Email)

		err = smtp.SendMail(smtpHost+":"+smtpPort, nil, "isanesaurabh@gmail.com", []string{recipient.Email}, []byte(msg))
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(50 * time.Millisecond)
		fmt.Printf("Worker %d sent email to %s\n", id, recipient.Email)
	}
}
