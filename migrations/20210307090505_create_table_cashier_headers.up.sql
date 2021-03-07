CREATE TABLE IF NOT EXISTS cashier_cash (
    id SERIAL,
    initial_amount DECIMAL(16,2) NOT NULL,
    ending_amount DECIMAL(16,2) NOT NULL,
    created_at DATETIME NULL,
    updated_at DATETIME NULL,
    PRIMARY KEY (id)
);