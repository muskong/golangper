-- 创建用户
CREATE USER blackuser WITH PASSWORD 'ZaX0VFi3BNfApBUuBQucTk7rI08plgIt';

-- 创建数据库
CREATE DATABASE blackapp;

-- 为用户分配数据库所有权
ALTER DATABASE blackapp OWNER TO blackuser;