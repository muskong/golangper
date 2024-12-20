-- -------------------------------------------------------------
-- TablePlus 6.0.0(550)
--
-- https://tableplus.com/
--
-- Database: blacklist
-- Generation Time: 2024-12-20 13:41:30.5980
-- -------------------------------------------------------------


DROP TABLE IF EXISTS "public"."blacklist_users";
-- This script only contains the table creation statements and does not fully represent the table in the database. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS blacklist_users_id_seq;

-- Table Definition
CREATE TABLE "public"."blacklist_users" (
    "id" int4 NOT NULL DEFAULT nextval('blacklist_users_id_seq'::regclass),
    "name" varchar(100) NOT NULL,
    "phone" varchar(20),
    "id_card" varchar(18),
    "email" varchar(100),
    "address" varchar(200),
    "remark" varchar(500),
    "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" timestamp,
    PRIMARY KEY ("id")
);

DROP TABLE IF EXISTS "public"."merchants";
-- This script only contains the table creation statements and does not fully represent the table in the database. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS merchants_id_seq;

-- Table Definition
CREATE TABLE "public"."merchants" (
    "id" int4 NOT NULL DEFAULT nextval('merchants_id_seq'::regclass),
    "name" varchar(100) NOT NULL,
    "address" varchar(255),
    "contact" varchar(50),
    "phone" varchar(20),
    "remark" text,
    "status" int2 DEFAULT 1,
    "api_key" varchar(32),
    "api_secret" varchar(64),
    "token" varchar(255),
    "token_exp" timestamptz,
    "created_at" timestamptz DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamptz DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" timestamptz,
    PRIMARY KEY ("id")
);

INSERT INTO "public"."blacklist_users" ("id", "name", "phone", "id_card", "email", "address", "remark", "created_at", "updated_at", "deleted_at") VALUES
(1, '高文静', '15054683318', '371702200709164320', NULL, NULL, NULL, '2024-12-19 13:39:39.435317', '2024-12-19 13:39:39.435317', NULL),
(2, '申建', '13386505556', '332624199506220911', NULL, NULL, NULL, '2024-12-19 13:39:39.435317', '2024-12-19 13:39:39.435317', NULL),
(3, '王桥', '13908207911', '500233199708091411', NULL, NULL, NULL, '2024-12-19 13:39:39.435317', '2024-12-19 13:39:39.435317', NULL),
(4, '黄周武', '18851121057', '350623196408161030', NULL, NULL, NULL, '2024-12-19 13:39:39.435317', '2024-12-19 13:39:39.435317', NULL),
(5, '李瑞', '13860065561', '410482199304269301', NULL, NULL, NULL, '2024-12-19 13:39:39.435317', '2024-12-19 13:39:39.435317', NULL),
(6, '朱志', '18473044267', '430621198810108736', NULL, NULL, NULL, '2024-12-19 13:39:39.435317', '2024-12-19 13:39:39.435317', NULL),
(7, '谢兵', '18841982200', '211011199605142010', NULL, NULL, NULL, '2024-12-19 13:39:39.435317', '2024-12-19 13:39:39.435317', NULL),
(8, '李佳欣', '15841951131', '211003199512280123', NULL, NULL, NULL, '2024-12-19 13:39:39.435317', '2024-12-19 13:39:39.435317', NULL),
(9, '谢茂', NULL, '430923198402132314', NULL, NULL, NULL, '2024-12-19 13:39:39.435317', '2024-12-19 13:39:39.435317', NULL),
(10, '陆登龙', '18046992238', '532627200604020518', NULL, NULL, NULL, '2024-12-19 13:39:39.435317', '2024-12-19 13:39:39.435317', NULL),
(11, '郭水金', '18655674459', '43022119821121501X', NULL, NULL, NULL, '2024-12-19 13:39:39.435317', '2024-12-19 13:39:39.435317', NULL),
(12, '庞小平', '18655880747', '411327198103051530', NULL, NULL, NULL, '2024-12-19 13:39:39.435317', '2024-12-19 13:39:39.435317', NULL),
(13, '查丽军', '13979305599', '362324197608010323', NULL, NULL, NULL, '2024-12-19 13:39:39.435317', '2024-12-19 13:39:39.435317', NULL),
(14, '林宥廷', '15259385125', '350921200211190092', NULL, NULL, NULL, '2024-12-19 13:39:39.435317', '2024-12-19 13:39:39.435317', NULL),
(15, '姬江林', '18811323650', '410823200401230335', NULL, NULL, NULL, '2024-12-19 13:39:39.435317', '2024-12-19 13:39:39.435317', NULL),
(16, '刘路遥', '15136851339', '522127200210206513', NULL, NULL, NULL, '2024-12-19 13:39:39.435317', '2024-12-19 13:39:39.435317', NULL),
(17, '刘佳钰', '18664798393', '350802200103101545', NULL, NULL, NULL, '2024-12-19 13:39:39.435317', '2024-12-19 13:39:39.435317', NULL),
(18, '黄文杰', '13247380248', '360402198611133850', NULL, NULL, NULL, '2024-12-19 13:39:39.435317', '2024-12-19 13:39:39.435317', NULL);

INSERT INTO "public"."merchants" ("id", "name", "address", "contact", "phone", "remark", "status", "api_key", "api_secret", "token", "token_exp", "created_at", "updated_at", "deleted_at") VALUES
(2, 't1', NULL, NULL, NULL, NULL, 1, '123', '456', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJNZXJjaGFudElEIjoyLCJleHAiOjE3MzQ3NTI4MzUsImlhdCI6MTczNDY2NjQzNX0.mlc9OrdZxCIe0F9f68xX9sWT3xdS7tJ-IDhjaAIVAr4', '2024-12-21 11:47:15.764487+08', '2024-12-20 09:38:06.245546+08', '2024-12-20 11:47:15.764771+08', NULL);



-- Indices
CREATE UNIQUE INDEX blacklist_users_phone_key ON public.blacklist_users USING btree (phone);
CREATE UNIQUE INDEX blacklist_users_id_card_key ON public.blacklist_users USING btree (id_card);
CREATE INDEX idx_blacklist_users_deleted_at ON public.blacklist_users USING btree (deleted_at);
CREATE INDEX idx_blacklist_users_name ON public.blacklist_users USING gin (name gin_trgm_ops);
CREATE INDEX idx_blacklist_users_phone ON public.blacklist_users USING gin (phone gin_trgm_ops);
CREATE INDEX idx_blacklist_users_id_card ON public.blacklist_users USING gin (id_card gin_trgm_ops);
CREATE INDEX idx_blacklist_users_email ON public.blacklist_users USING gin (email gin_trgm_ops);
CREATE INDEX idx_blacklist_users_address ON public.blacklist_users USING gin (address gin_trgm_ops);
CREATE INDEX idx_blacklist_users_remark ON public.blacklist_users USING gin (remark gin_trgm_ops);


-- Indices
CREATE UNIQUE INDEX merchants_api_key_key ON public.merchants USING btree (api_key);
CREATE INDEX idx_merchants_api_key ON public.merchants USING btree (api_key);
CREATE INDEX idx_merchants_deleted_at ON public.merchants USING btree (deleted_at);
