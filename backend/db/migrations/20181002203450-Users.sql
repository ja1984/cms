
-- +migrate Up
CREATE TABLE public.users
(
    id serial PRIMARY KEY NOT NULL,
    external_id varchar(255) NOT NULL,
    email varchar(512) NOT NULL
);
CREATE UNIQUE INDEX users_external_id_uindex ON public.users (external_id);
-- +migrate Down
DROP TABLE public.users;