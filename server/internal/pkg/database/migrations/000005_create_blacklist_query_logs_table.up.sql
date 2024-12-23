CREATE TABLE blacklist_query_logs (
    id SERIAL PRIMARY KEY,
    merchant_id INTEGER NOT NULL,
    phone VARCHAR(20) NOT NULL,
    query_time TIMESTAMP NOT NULL,
    ip VARCHAR(50) NOT NULL,
    user_agent VARCHAR(255),
    result BOOLEAN NOT NULL,
    CONSTRAINT fk_blacklist_query_logs_merchant FOREIGN KEY (merchant_id) REFERENCES merchants(id)
);

CREATE INDEX idx_blacklist_query_logs_merchant_id ON blacklist_query_logs(merchant_id);
CREATE INDEX idx_blacklist_query_logs_query_time ON blacklist_query_logs(query_time);
CREATE INDEX idx_blacklist_query_logs_phone ON blacklist_query_logs(phone);