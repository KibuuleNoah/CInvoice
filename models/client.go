package models

// Client ...
type Client struct {
	ID               *string `db:"id" json:"id"`
	ClientName       *string `db:"client_name" json:"clientName"`
	Address          *string `db:"address" json:"address"`
	ArmedUnitPrice   *int    `db:"armed_unit_price" json:"armedUnitPrice"`
	ArmedCount       *int    `db:"armed_count" json:"armedCount"`
	UnarmedUnitPrice *int    `db:"unarmed_unit_price" json:"unarmedUnitPrice"`
	UnarmedCount     *int    `db:"unarmed_count" json:"unarmedCount"`
	TotalGuards      *int    `db:"total_guards" json:"totalGuards"`
	MonthlyPay       *int    `db:"monthly_pay" json:"monthlyPay"`
	StartDate        *string `db:"start_date" json:"startDate"`
	Contact          *string `db:"contact" json:"contact"`
	CreatedAt        *int64  `db:"created_at" json:"createdAt"`
	UpdatedAt        *int64  `db:"updated_at" json:"updatedAt"`
}
