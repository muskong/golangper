-- 为常用查询字段添加索引
CREATE INDEX IF NOT EXISTS idx_blacklist_users_name ON blacklist_users USING gin (name gin_trgm_ops);
CREATE INDEX IF NOT EXISTS idx_blacklist_users_phone ON blacklist_users USING gin (phone gin_trgm_ops);
CREATE INDEX IF NOT EXISTS idx_blacklist_users_id_card ON blacklist_users USING gin (id_card gin_trgm_ops);
CREATE INDEX IF NOT EXISTS idx_blacklist_users_email ON blacklist_users USING gin (email gin_trgm_ops);
CREATE INDEX IF NOT EXISTS idx_blacklist_users_address ON blacklist_users USING gin (address gin_trgm_ops);
CREATE INDEX IF NOT EXISTS idx_blacklist_users_remark ON blacklist_users USING gin (remark gin_trgm_ops);

-- 启用 pg_trgm 扩展（用于模糊搜索）
CREATE EXTENSION IF NOT EXISTS pg_trgm; 