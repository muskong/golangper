CREATE TABLE merchant_login_logs (
    id SERIAL PRIMARY KEY,
    merchant_id INTEGER NOT NULL,
    ip VARCHAR(50) NOT NULL,
    user_agent VARCHAR(255),
    login_time TIMESTAMP NOT NULL,
    status SMALLINT NOT NULL,
    remark VARCHAR(255),
    CONSTRAINT fk_merchant_login_logs_merchant FOREIGN KEY (merchant_id) REFERENCES merchants(id)
);

CREATE INDEX idx_merchant_login_logs_merchant_id ON merchant_login_logs(merchant_id);
CREATE INDEX idx_merchant_login_logs_login_time ON merchant_login_logs(login_time);