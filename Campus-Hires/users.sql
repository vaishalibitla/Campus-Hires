-- Table: public.users

-- DROP TABLE public.users;

CREATE TABLE public.users
(
    username character varying(50) COLLATE pg_catalog."default" NOT NULL,
    password character varying COLLATE pg_catalog."default",
    CONSTRAINT users_pkey PRIMARY KEY (username)
)

TABLESPACE pg_default;

ALTER TABLE public.users
    OWNER to postgres;