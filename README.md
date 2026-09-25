# Civoice

Client and invoice management system for security service contracts, This Project is being used by a security company called Calvary Security.
Backend is a Go/Gin API with PostgreSQL; frontend is server-rendered Go templates (`html/template`) styled with Tailwind and Alpine.js. Invoices are recurring per client, auto-generated on a billing cycle, and exportable as PDF.

## Getting Started

Clone the repo, install dependencies, and run the app locally or via Docker Compose.

### Prerequisites

- Go 1.16+
- Gin
_ PostgreSQL
- Docker
- Docker Compose

### Installation

1. Clone the repo
   ```bash
   git clone https://github.com/KibuuleNoah/CInvoive.git
   cd civoice
   ```
### Development

## Project Structure

```
.
├── controllers/       # HTTP handlers (Gin)
├── services/          # Business logic, PDF generation
├── models/            # DB models and query logic
├── web/templates/      # Server-rendered HTML (layout, clients, invoices)
├── db/migrations/         # SQL schema migrations
├── Dockerfile
└── docker-compose.yml
```

## Environment Variables

| Variable       | Description                          |
|----------------|---------------------------------------|
| `DB_HOST`      | PostgreSQL host                       |
| `DB_PORT`      | PostgreSQL port                       |
| `DB_USER`      | PostgreSQL user                       |
| `DB_PASSWORD`  | PostgreSQL password                   |
| `DB_NAME`      | PostgreSQL database name              |
| `PORT`         | App listen port (default `8080`)      |

Copy `.env.example` to `.env` and fill in values before running `make run`.
