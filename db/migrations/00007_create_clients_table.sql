-- +goose Up
-- +goose StatementBegin

CREATE SEQUENCE IF NOT EXISTS clients_id_seq START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1;
CREATE TABLE clients (
    id                  TEXT PRIMARY KEY DEFAULT generate_short_id('clients_id_seq'),
    client_name         TEXT NOT NULL,
    address             TEXT NOT NULL,
    armed_unit_price    INTEGER NOT NULL,
    armed_count         SMALLINT NOT NULL,
    unarmed_unit_price  INTEGER NOT NULL,
    unarmed_count       SMALLINT NOT NULL,
    total_guards        SMALLINT NOT NULL,
    monthly_pay         INTEGER GENERATED ALWAYS AS (
                            (armed_unit_price * armed_count) + (unarmed_unit_price * unarmed_count)
                        ) STORED,
    start_date          DATE NOT NULL,
    contact             VARCHAR(15) NOT NULL,
    created_at          BIGINT NOT NULL,
    updated_at          BIGINT NOT NULL
);

CREATE INDEX idx_clients_client_name ON clients (client_name);
CREATE INDEX idx_clients_contact ON clients (contact);

CREATE TRIGGER create_clients_created_at
    BEFORE INSERT ON clients
    FOR EACH ROW EXECUTE PROCEDURE created_at_column();

CREATE TRIGGER update_clients_updated_at
    BEFORE UPDATE ON clients
    FOR EACH ROW EXECUTE PROCEDURE update_at_column();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP SEQUENCE IF EXISTS clients_id_seq;
DROP TABLE IF EXISTS clients;
-- +goose StatementEnd
