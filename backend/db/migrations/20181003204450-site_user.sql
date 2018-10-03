
-- +migrate Up
CREATE TABLE public.site_user
(
    id serial PRIMARY KEY NOT NULL,
    site_id int NOT NULL,
    user_id int NOT NULL,
    is_admin boolean NOT NULL
);
CREATE UNIQUE INDEX site_user_ids_uindex ON public.site_user (site_id, user_id);
-- +migrate Down
DROP TABLE public.site_user;