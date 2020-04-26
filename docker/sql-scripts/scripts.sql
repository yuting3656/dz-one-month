CREATE DATABASE  tim_pokemon;

\c tim_pokemon;

CREATE EXTENSION "uuid-ossp";

CREATE TABLE team (
    uuid uuid NOT NULL DEFAULT uuid_generate_v4(),
    name varchar(30) NULL,
    memo varchar(50) NULL,
    create_time timestamptz NULL,
	update_time timestamptz NULL,
    CONSTRAINT team_pkey PRIMARY KEY (uuid)
);

CREATE TABLE pokemon (
    uuid uuid NOT NULL DEFAULT uuid_generate_v4(),
    name varchar(30) NULL,
    photo varchar NULL,
    create_time timestamptz NULL,
	update_time timestamptz NULL,
    CONSTRAINT pokemon_pkey PRIMARY KEY (uuid)
);

CREATE TABLE team_pokemon_list (
     uuid uuid NOT NULL DEFAULT uuid_generate_v4(),
     team_uuid uuid NOT NULL,
     pokemon_uuid uuid NOT NULL,
     pokemon_order int4 NOT NULL,
     note varchar NULL,
     CONSTRAINT team_pokemon_list_pkey PRIMARY KEY (team_uuid, pokemon_uuid),
     CONSTRAINT team_pokemon_list_team_uuid_fkey FOREIGN KEY (team_uuid) REFERENCES team(uuid) ON UPDATE CASCADE ON DELETE SET NULL NOT VALID,
     CONSTRAINT team_pokemon_list_pokemon_uuid_fkey FOREIGN KEY (pokemon_uuid) REFERENCES pokemon(uuid) ON UPDATE CASCADE ON DELETE SET NULL NOT VALID
);

\c tim_pokemon

insert into 
    team (name, memo) 
values 
    ('帥氣小隊', '火系為主');