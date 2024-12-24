-- 创建数据库和用户
-- CREATE DATABASE blackapp;
-- CREATE USER blackuser WITH PASSWORD 'ZaX0VFi3BNfApBUuBQucTk7rI08plgIt';
-- GRANT ALL PRIVILEGES ON DATABASE blackapp TO blackuser;

\c blackapp;

-- 管理员表
CREATE TABLE admins (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    password VARCHAR(100) NOT NULL,
    name VARCHAR(50) NOT NULL,
    status INT DEFAULT 1, -- 1:启用 2:禁用
    last_login TIMESTAMP,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);
CREATE UNIQUE INDEX idx_admins_username ON admins(username);
CREATE INDEX idx_admins_deleted_at ON admins(deleted_at);

-- 商户表
CREATE TABLE merchants (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    password VARCHAR(100) NOT NULL,
    name VARCHAR(50) NOT NULL,
    status INT DEFAULT 1, -- 1:启用 2:禁用
    api_key VARCHAR(32) NOT NULL,
    rate_limit INT DEFAULT 100, -- 每分钟请求限制
    last_login TIMESTAMP,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);
CREATE UNIQUE INDEX idx_merchants_username ON merchants(username);
CREATE UNIQUE INDEX idx_merchants_api_key ON merchants(api_key);
CREATE INDEX idx_merchants_deleted_at ON merchants(deleted_at);

-- 黑名单表
CREATE TABLE blacklists (
    id SERIAL PRIMARY KEY,
    phone VARCHAR(20),
    id_card VARCHAR(18),
    name VARCHAR(50),
    status INT DEFAULT 1, -- 1:启用 2:禁用
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);
CREATE INDEX idx_blacklists_phone ON blacklists(phone);
CREATE INDEX idx_blacklists_id_card ON blacklists(id_card);
CREATE INDEX idx_blacklists_deleted_at ON blacklists(deleted_at);

-- 登录日志表
CREATE TABLE login_logs (
    id SERIAL PRIMARY KEY,
    type INT NOT NULL, -- 1:商户 2:管理员
    user_id INT NOT NULL,
    ip VARCHAR(50) NOT NULL,
    user_agent VARCHAR(255) NOT NULL,
    status INT NOT NULL, -- 1:成功 2:失败
    created_at TIMESTAMP NOT NULL
);
CREATE INDEX idx_login_logs_type ON login_logs(type);
CREATE INDEX idx_login_logs_user_id ON login_logs(user_id);

-- 查询日志表
CREATE TABLE query_logs (
    id SERIAL PRIMARY KEY,
    merchant_id INT NOT NULL,
    phone VARCHAR(20),
    id_card VARCHAR(18),
    name VARCHAR(50),
    ip VARCHAR(50) NOT NULL,
    user_agent VARCHAR(255) NOT NULL,
    exists BOOLEAN NOT NULL,
    created_at TIMESTAMP NOT NULL
);
CREATE INDEX idx_query_logs_merchant_id ON query_logs(merchant_id);

-- 插入默认管理员账号
INSERT INTO admins (username, password, name, status, created_at, updated_at)
VALUES (
    'admin',
    '21232f297a57a5a743894a0e4a801fc3', -- admin的MD5值
    '系统管理员',
    1,
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
);

-- 授权
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO blackuser;
GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public TO blackuser; 