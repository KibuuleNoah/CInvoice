-- +goose Up
-- +goose StatementBegin
CREATE TYPE invoice_status AS ENUM ('pending', 'paid', 'overdue', 'cancelled');
CREATE SEQUENCE IF NOT EXISTS invoices_id_seq START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1;
CREATE TABLE invoices (
    id                  TEXT PRIMARY KEY DEFAULT generate_short_id('invoices_id_seq'),
    client_id           TEXT NOT NULL REFERENCES clients(id),
    invoice_number      TEXT NOT NULL UNIQUE,
    amount              INTEGER NOT NULL,
    status              invoice_status NOT NULL DEFAULT 'pending',
    due_date            DATE NOT NULL,
    paid_date           DATE,
    created_at          BIGINT NOT NULL,
    updated_at          BIGINT NOT NULL
);

CREATE INDEX idx_invoices_client_id ON invoices (client_id);
CREATE INDEX idx_invoices_status ON invoices (status);
CREATE INDEX idx_invoices_due_date ON invoices (due_date);

CREATE TRIGGER create_invoices_created_at
    BEFORE INSERT ON invoices
    FOR EACH ROW EXECUTE PROCEDURE created_at_column();

CREATE TRIGGER update_invoices_updated_at
    BEFORE UPDATE ON invoices
    FOR EACH ROW EXECUTE PROCEDURE update_at_column();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TYPE IF EXISTS invoice_status;
DROP SEQUENCE IF EXISTS invoices_id_seq;
DROP TABLE IF EXISTS invoices;
-- +goose StatementEnd
