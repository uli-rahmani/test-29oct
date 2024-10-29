CREATE TABLE "user" (
	id bigserial NOT NULL,
	"name" varchar(300) NOT NULL,
	email varchar(300) NOT NULL,
	password varchar(300) NOT NULL,
	"role" varchar(50) NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NULL,
    deleted_at timestamp NULL,
	CONSTRAINT banks_name_key UNIQUE (email)
);