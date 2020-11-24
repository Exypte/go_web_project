-- Adminer 4.7.7 PostgreSQL dump

DROP TABLE IF EXISTS "company";
DROP SEQUENCE IF EXISTS company_id_seq;
CREATE SEQUENCE company_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 START 1 CACHE 1;

CREATE TABLE "public"."company" (
    "id" integer DEFAULT nextval('company_id_seq') NOT NULL,
    "name" text NOT NULL,
    "age" integer NOT NULL,
    "address" text NOT NULL,
    "salary" real NOT NULL,
    CONSTRAINT "company_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

INSERT INTO "company" ("name", "age", "address", "salary") VALUES
('Paul',	32,	'California',	20000),
('Allen',	25,	'Texas',	15000),
('Teddy',	23,	'Norway',	20000),
('Mark',	25,	'Rich-Mond',	65000),
('David',	27,	'Texas',	85000),
('Kim',	22,	'South-Hall',	45000),
('James',	24,	'Houston',	10000);

-- 2020-11-23 11:31:17.505897+00