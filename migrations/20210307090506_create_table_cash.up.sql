CREATE TABLE IF NOT EXISTS cashiers (
    id SERIAL,
    person_id BIGINT NOT NULL,
    source VARCHAR(255) NOT NULL,
    amount DECIMAL(16,2) DEFAULT 0 NOT NULL,
    type VARCHAR(255) NOT NULL,
    category VARCHAR(255) NOT NULL,
    source VARCHAR(255) NOT NULL,
    notes TEXT,
    created_at DATETIME NULL,
    updated_at DATETIME NULL,
    PRIMARY KEY (id)
);