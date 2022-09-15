-- DROP SCHEMA public;

CREATE SCHEMA public AUTHORIZATION postgres;
-- public.testuser definition

-- Drop table

-- DROP TABLE public.testuser;

CREATE TABLE public.testuser (
	id int4 NOT NULL GENERATED ALWAYS AS IDENTITY,
	"name" varchar NULL
);
