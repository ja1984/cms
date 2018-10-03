
-- +migrate Up
CREATE TABLE public.site
(
    id serial PRIMARY KEY NOT NULL,
    name varchar(255) NOT NULL,S
    creator_user_id int NOT NULL,
    url_production varchar(512) NULL,
    url_staging varchar(512) NULL,
    url_test varchar(512) NULL
);

-- +migrate Down
DROP TABLE public.site;