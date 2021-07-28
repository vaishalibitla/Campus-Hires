-- Table: public.members

-- DROP TABLE public.members;

CREATE TABLE public.members
(
    id integer NOT NULL DEFAULT nextval('members_id_seq'::regclass),
    name character varying COLLATE pg_catalog."default" NOT NULL,
    email character varying COLLATE pg_catalog."default",
    phone bigint,
    joining_date date,
    CONSTRAINT pk_members PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE public.members
    OWNER to postgres;