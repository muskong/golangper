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

-- 创建索引
CREATE INDEX idx_merchants_api_key ON merchants(api_key);
CREATE INDEX idx_merchants_deleted_at ON merchants(deleted_at);

-- 创建更新时间触发器
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_merchants_updated_at
    BEFORE UPDATE ON merchants
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

-- 添加字段注释
COMMENT ON COLUMN merchants.name IS '商户名称';
COMMENT ON COLUMN merchants.address IS '商户地址';
COMMENT ON COLUMN merchants.contact IS '商户联系人';
COMMENT ON COLUMN merchants.phone IS '商户联系电话';
COMMENT ON COLUMN merchants.remark IS '商户备注';
COMMENT ON COLUMN merchants.status IS '商户状态 1:正常 2:禁用';
COMMENT ON COLUMN merchants.api_key IS '商户API Key';
COMMENT ON COLUMN merchants.api_secret IS '商户API Secret';
COMMENT ON COLUMN merchants.token IS '商户API Token';
COMMENT ON COLUMN merchants.token_exp IS 'Token过期时间';