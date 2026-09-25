package services

import (
	"CInvoice/db"
	"CInvoice/models"
	"database/sql"
	"errors"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/robfig/cron/v3"
)

type InvoiceServiceConfig struct {
	DB *sqlx.DB
}

type InvoiceService struct {
	cfg InvoiceServiceConfig
}

func NewInvoiceService() *InvoiceService {
	return &InvoiceService{cfg: InvoiceServiceConfig{DB: db.AppDB()}}
}

var (
	ErrInvoiceNotFound = errors.New("invoice record not found")
	ErrInvoiceCreate   = errors.New("failed creating invoice record")
	ErrInvoiceUpdate   = errors.New("failed updating invoice record")
	ErrInvoiceList     = errors.New("failed fetching invoices")
	ErrInvoiceDelete   = errors.New("failed deleting invoice record")
)

func (s *InvoiceService) ListInvoices(limit, offset string) (*[]models.Invoice, error) {
	query := `
		SELECT
			i.id, i.invoice_number, i.amount, i.status, i.due_date, i.created_at, i.updated_at,
			c.id                 "client.id",
			c.client_name        "client.client_name",
			c.address            "client.address",
			c.armed_unit_price   "client.armed_unit_price",
			c.armed_count        "client.armed_count",
			c.unarmed_unit_price "client.unarmed_unit_price",
			c.unarmed_count      "client.unarmed_count",
			c.total_guards       "client.total_guards",
			c.monthly_pay        "client.monthly_pay",
			c.start_date         "client.start_date",
			c.contact            "client.contact",
			c.created_at         "client.created_at",
			c.updated_at         "client.updated_at"
		FROM invoices i
		JOIN clients c ON c.id = i.client_id
		ORDER BY i.created_at DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := s.cfg.DB.Queryx(query, 500, 0)
	if err != nil {
		log.Println("InvoiceService.ListInvoices ERROR: ", err)
		return nil, ErrInvoiceList
	}
	defer rows.Close()

	var invoices []models.Invoice
	for rows.Next() {
		var row models.Invoice
		if err := rows.StructScan(&row); err != nil {
			log.Println("InvoiceService.ListInvoices ERROR: ", err)
			return nil, ErrInvoiceList
		}
		invoices = append(invoices, row)
	}
	return &invoices, nil
}
func (s *InvoiceService) GetInvoiceByID(id string) (*models.Invoice, error) {
	var invoice models.Invoice
	query := `
		SELECT
			i.id, i.invoice_number, i.amount, i.status, i.due_date, i.created_at, i.updated_at,
			c.id                 "client.id",
			c.client_name        "client.client_name",
			c.address            "client.address",
			c.armed_unit_price   "client.armed_unit_price",
			c.armed_count        "client.armed_count",
			c.unarmed_unit_price "client.unarmed_unit_price",
			c.unarmed_count      "client.unarmed_count",
			c.total_guards       "client.total_guards",
			c.monthly_pay        "client.monthly_pay",
			c.start_date         "client.start_date",
			c.contact            "client.contact",
			c.created_at         "client.created_at",
			c.updated_at         "client.updated_at"
		FROM invoices i
		JOIN clients c ON c.id = i.client_id
		WHERE i.id = $1
	`
	err := s.cfg.DB.Get(&invoice, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Println("InvoiceService.GetInvoiceByID ERROR: ", err)
			return nil, ErrInvoiceNotFound
		}
		log.Println("InvoiceService.GetInvoiceByID ERROR: ", err)
		return nil, ErrInvoiceNotFound
	}
	return &invoice, nil
}
func StartInvoiceCronJobs() {
	c := cron.New()

	// Run every hour (at minute 0)
	_, err := c.AddFunc("0 * * * *", func() {
		query := `
    INSERT INTO invoices (client_id, invoice_number, amount, status, due_date, created_at, updated_at)
    SELECT 
        c.id,
        'INV-' || TO_CHAR(NOW(), 'YYYYMM') || '-' || SUBSTRING(c.id FROM 1 FOR 6),
        c.monthly_pay,
        'pending',
        DATE_TRUNC('month', NOW() + INTERVAL '1 month') + INTERVAL '4 day', 
        EXTRACT(EPOCH FROM NOW()) * 1000,
        EXTRACT(EPOCH FROM NOW()) * 1000
    FROM clients c
    WHERE 
        -- Triggers if 5 or fewer days remain before their start day this month, or if the day has already passed
        NOW() >= (
            DATE_TRUNC('month', NOW()) + 
            (LEAST(EXTRACT(DAY FROM c.start_date)::int, 28) - 1) * INTERVAL '1 day'
            - INTERVAL '7 day'
        )
        AND NOT EXISTS (
            SELECT 1 FROM invoices i 
            WHERE i.invoice_number = 'INV-' || TO_CHAR(NOW(), 'YYYYMM') || '-' || SUBSTRING(c.id FROM 1 FOR 6)
        );
`
		res, err := db.AppDB().Exec(query)
		if err != nil {
			log.Printf("Failed to generate recurring invoices: %v", err)
		}

		if rowCount, _ := res.RowsAffected(); rowCount > 0 {
			log.Printf("Generated %d Invoices\n", rowCount)
		} else {
			log.Println("Invoices Uptodate, Nothing to generate")
		}
	})

	if err != nil {
		log.Fatalf("Failed to schedule cron job: %v", err)
	}

	c.Start()
}

func (s *InvoiceService) MarkInvoiceAsPaid(id string) error {
	res, err := s.cfg.DB.Exec("UPDATE invoices SET status = 'paid' WHERE id = $1", id)
	if err != nil {
		log.Println("InvoiceService.MarkInvoiceAsPaid ERROR: ", err)
		return ErrInvoiceUpdate
	}

	rows, err := res.RowsAffected()
	log.Printf("InvoiceService.MarkInvoiceAsPaid: %d\n", rows)
	if err != nil {
		log.Println("InvoiceService.MarkInvoiceAsPaid ERROR: ", err)
		return ErrInvoiceUpdate
	}

	if rows == 0 {
		return ErrInvoiceNotFound
	}
	return nil
}
