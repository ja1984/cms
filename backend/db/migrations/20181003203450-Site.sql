
-- +migrate Up
CREATE TABLE public.site
(
    id serial PRIMARY KEY NOT NULL,
    name varchar(255) NOT NULL
);

CREATE TABLE public.teams
(
    id serial PRIMARY KEY NOT NULL,
    site_id int NOT NULL,
    name varchar(255) NOT NULL,
    permission int NOT NULL
);
create table public.teams_users
(
  user_id integer not null,
  team_id integer not null,
  id      serial  not null
    constraint teams_users_pkey
    primary key
);

CREATE TABLE public.environments
(
    id serial PRIMARY KEY NOT NULL,
    name varchar(255) NOT NULL
);

-- +migrate Down
DROP TABLE public.site;
DROP TABLE public.teams;
DROP TABLE public.teams_users;
DROP TABLE public.environments;
