-- +goose Up
ALTER TABLE clients ADD CONSTRAINT clients_key UNIQUE (client_name, address);

-- +goose Down
ALTER TABLE clients DROP CONSTRAINT IF EXISTS clients_key;
