CREATE TABLE IF NOT EXISTS merchants (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    address VARCHAR(255),
    contact VARCHAR(50),
    phone VARCHAR(20),
    remark TEXT,
    status SMALLINT DEFAULT 1,
    api_key VARCHAR(32) UNIQUE,
    api_secret VARCHAR(64),
    token VARCHAR(255),
    token_exp TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE INDEX idx_merchants_api_key ON merchants(api_key);
CREATE INDEX idx_merchants_deleted_at ON merchants(deleted_at); 