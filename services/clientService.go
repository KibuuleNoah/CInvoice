package services

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"

	"CInvoice/db"
	"CInvoice/forms"
	"CInvoice/models"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type ClientServiceConfig struct {
	DB *sqlx.DB
}

type ClientService struct {
	cfg ClientServiceConfig
}

func NewClientService() *ClientService {
	return &ClientService{cfg: ClientServiceConfig{DB: db.AppDB()}}
}

var (
	ErrClientNotFound   = errors.New("client record not found")
	ErrClientNameExists = errors.New("client name already exists")
	ErrClientCreate     = errors.New("failed creating client record")
	ErrClientUpdate     = errors.New("failed updating client record")
	ErrClientList       = errors.New("failed fetching clients")
	ErrClientDelete     = errors.New("failed deleting client record")
)

func (s *ClientService) ListClients(limit, offset string) (*[]models.Client, error) {
	query := `
		SELECT
			id, client_name, address, armed_unit_price, armed_count, unarmed_unit_price, unarmed_count, total_guards, monthly_pay, start_date, contact, created_at, updated_at
		FROM clients LIMIT $1 OFFSET $2
		`

	rows, err := s.cfg.DB.Queryx(query, 500, 0)
	if err != nil {
		log.Println("ClientService.ListClients ERROR: ", err)
		return nil, fmt.Errorf("failed fetching clients: %w", err)
	}
	defer rows.Close()

	var clients []models.Client
	for rows.Next() {
		var row models.Client
		if err := rows.StructScan(&row); err != nil {
			log.Println("ClientService.ListClients ERROR: ", err)
			return nil, ErrClientList
		}

		clients = append(clients, row)
	}
	return &clients, nil
}

func (s *ClientService) GetClientByID(id string) (*models.Client, error) {
	var client models.Client
	err := s.cfg.DB.Get(&client, "SELECT id, client_name, address, armed_unit_price, armed_count, unarmed_unit_price, unarmed_count, total_guards, monthly_pay, start_date, contact, created_at, updated_at FROM clients WHERE id = $1", id)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) && pqErr.Code == "23505" {
			log.Println("ClientService.CreateClient ERROR: ", err)
			return nil, ErrClientNameExists
		}

		if errors.Is(err, sql.ErrNoRows) {
			log.Println("ClientService.GetClientByID ERROR: ", err)
			return nil, ErrClientNotFound
		}

		log.Println("ClientService.GetClientByID ERROR: ", err)
		return nil, ErrClientNotFound
	}
	return &client, nil
}

func (s *ClientService) CreateClient(form forms.CreateClientForm) (*models.Client, error) {
	var client models.Client
	err := s.cfg.DB.QueryRowx(`
		INSERT INTO clients (client_name, address, armed_unit_price, armed_count, unarmed_unit_price, unarmed_count, total_guards, start_date, contact)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id, client_name, address, armed_unit_price, armed_count, unarmed_unit_price, unarmed_count, total_guards, monthly_pay, start_date, contact, created_at, updated_at`,
		form.ClientName, form.Address, form.ArmedUnitPrice, form.ArmedCount, form.UnarmedUnitPrice, form.UnarmedCount, form.TotalGuards, form.StartDate, form.Contact,
	).StructScan(&client)
	if err != nil {
		log.Println("ClientService.CreateClient ERROR: ", err)
		return nil, ErrClientCreate
	}
	return &client, nil
}

func (s *ClientService) UpdateClient(form forms.UpdateClientForm) (*models.Client, error) {
	sets := []string{}
	args := []interface{}{}
	i := 1

	if form.ClientName != nil {
		sets = append(sets, fmt.Sprintf("client_name = $%d", i))
		args = append(args, *form.ClientName)
		i++
	}
	if form.Address != nil {
		sets = append(sets, fmt.Sprintf("address = $%d", i))
		args = append(args, *form.Address)
		i++
	}
	if form.ArmedUnitPrice != nil {
		sets = append(sets, fmt.Sprintf("armed_unit_price = $%d", i))
		args = append(args, *form.ArmedUnitPrice)
		i++
	}
	if form.ArmedCount != nil {
		sets = append(sets, fmt.Sprintf("armed_count = $%d", i))
		args = append(args, *form.ArmedCount)
		i++
	}
	if form.UnarmedUnitPrice != nil {
		sets = append(sets, fmt.Sprintf("unarmed_unit_price = $%d", i))
		args = append(args, *form.UnarmedUnitPrice)
		i++
	}
	if form.UnarmedCount != nil {
		sets = append(sets, fmt.Sprintf("unarmed_count = $%d", i))
		args = append(args, *form.UnarmedCount)
		i++
	}
	if form.TotalGuards != nil {
		sets = append(sets, fmt.Sprintf("total_guards = $%d", i))
		args = append(args, *form.TotalGuards)
		i++
	}
	if form.StartDate != nil {
		sets = append(sets, fmt.Sprintf("start_date = $%d", i))
		args = append(args, *form.StartDate)
		i++
	}
	if form.Contact != nil {
		sets = append(sets, fmt.Sprintf("contact = $%d", i))
		args = append(args, *form.Contact)
		i++
	}

	if len(sets) == 0 {
		return s.GetClientByID(form.ID)
	}

	args = append(args, form.ID)

	query := fmt.Sprintf(`
		UPDATE clients SET %s
		WHERE id = $%d
		RETURNING id, client_name, address, armed_unit_price, armed_count, unarmed_unit_price, unarmed_count, total_guards, monthly_pay, start_date, contact, created_at, updated_at`,
		strings.Join(sets, ", "), i)

	var client models.Client
	err := s.cfg.DB.QueryRowx(query, args...).StructScan(&client)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Println("ClientService.UpdateClient ERROR: ", err)
			return nil, ErrClientNotFound
		}
		log.Println("ClientService.UpdateClient ERROR: ", err)
		return nil, ErrClientUpdate
	}
	return &client, nil
}

func (s *ClientService) DeleteClient(id string) error {
	res, err := s.cfg.DB.Exec("DELETE FROM clients WHERE id = $1", id)
	if err != nil {
		log.Println("ClientService.DeleteClient ERROR: ", err)
		return ErrClientDelete
	}

	rows, err := res.RowsAffected()
	if err != nil {
		log.Println("ClientService.DeleteClient ERROR: ", err)
		return ErrClientDelete
	}

	if rows == 0 {
		return ErrClientNotFound
	}
	return nil
}
