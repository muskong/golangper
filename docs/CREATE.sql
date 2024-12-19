CREATE TABLE blacklist_users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    phone VARCHAR(20) UNIQUE,
    id_card VARCHAR(18) UNIQUE,
    email VARCHAR(100),
    address VARCHAR(200),
    remark VARCHAR(500),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- 创建更新时间触发器
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_blacklist_users_updated_at
    BEFORE UPDATE ON blacklist_users
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

-- 创建删除时间索引
CREATE INDEX idx_blacklist_users_deleted_at ON blacklist_users(deleted_at);