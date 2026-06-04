# email-dispatcher

A Go backend application that reads recipient data from MongoDB and dispatches personalized emails concurrently using a producer-consumer pattern with goroutines.

---

## How It Works

1. **Producer** — reads recipients from MongoDB and pushes them into a channel
2. **Consumer (workers)** — multiple goroutines pick recipients from the channel, render an HTML email template, and send the email via SMTP
3. **WaitGroup** — the main goroutine waits for all workers to finish before exiting

---

## Project Structure

email-dispatcher/
├── main.go                        # Entry point, flag parsing, wires producer & consumers
├── producer.go                    # Reads recipients from MongoDB → channel
├── consumer.go                    # Worker goroutines that send emails via SMTP
├── db.go                          # MongoDB connection, index creation, recipient loader
├── email.tmpl                     # Go HTML template for email body
├── name_emails_200_records.csv    # Sample CSV data (used for seeding)
├── go.mod
└── go.sum  

---

## Tech Stack

- **Language:** Go 1.26.3
- **Database:** MongoDB (via `go.mongodb.org/mongo-driver v1.17.9`)
- **Email:** `net/smtp` (standard library)
- **Templating:** `html/template` (standard library)
- **Concurrency:** goroutines, channels, `sync.WaitGroup`

---

## Prerequisites

- Go 1.26.3+
- MongoDB running at `localhost:27017` (default)
- An SMTP server running at `localhost:1025` (e.g. [MailHog](https://github.com/mailhog/MailHog) for local testing)

---

## Setup

**1. Clone the repository**

```bash
git clone https://github.com/saurabhisane/email-dispatcher.git
cd email-dispatcher
```

**2. Install dependencies**

```bash
go mod tidy
```

**3. Seed MongoDB**

Import the provided CSV into the `contacts` collection in the `email_dispatcher` database. Each document should have `name` and `email` fields.

**4. Run the application**

```bash
go run .
```

With custom flags:

```bash
go run . -mongoURI="mongodb://localhost:27017" -dbName="email_dispatcher" -workerCount=10
```

---

## CLI Flags

| Flag | Default | Description |
|------|---------|-------------|
| `-mongoURI` | `mongodb://localhost:27017` | MongoDB connection URI |
| `-dbName` | `email_dispatcher` | MongoDB database name |
| `-workerCount` | `5` | Number of concurrent email worker goroutines |

---

## MongoDB Schema

Collection: `contacts`

```json
{
  "name":  "John Doe",
  "email": "johndoe@example.com"
}
```

A unique index is automatically created on the `email` field at startup.

---

## Email Template

The email body is rendered from `email.tmpl` using Go's `html/template` engine. The `Recipient` struct (with `Name` and `Email` fields) is passed as template data.
