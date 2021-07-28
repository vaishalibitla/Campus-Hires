-- Table: public.queries

-- DROP TABLE public.queries;

CREATE TABLE public.queries
(
    email character varying COLLATE pg_catalog."default",
    query text COLLATE pg_catalog."default"
)

TABLESPACE pg_default;

ALTER TABLE public.queries
    OWNER to postgres;