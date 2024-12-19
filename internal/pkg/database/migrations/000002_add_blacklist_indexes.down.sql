DROP INDEX IF EXISTS idx_blacklist_users_name;
DROP INDEX IF EXISTS idx_blacklist_users_phone;
DROP INDEX IF EXISTS idx_blacklist_users_id_card;
DROP INDEX IF EXISTS idx_blacklist_users_email;
DROP INDEX IF EXISTS idx_blacklist_users_address;
DROP INDEX IF EXISTS idx_blacklist_users_remark;

DROP EXTENSION IF EXISTS pg_trgm; 