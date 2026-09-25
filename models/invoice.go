package models

type Invoice struct {
	ID            string `db:"id" json:"id"`
	Client        Client `db:"client" json:"client"`
	InvoiceNumber string `db:"invoice_number" json:"invoiceNumber"`
	Amount        int    `db:"amount" json:"amount"`
	Status        string `db:"status" json:"status"`
	DueDate       string `db:"due_date" json:"dueDate"`
	CreatedAt     int64  `db:"created_at" json:"createdAt"`
	UpdatedAt     int64  `db:"updated_at" json:"updatedAt"`
}
