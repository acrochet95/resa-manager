CREATE DATABASE resa;

CREATE TABLE apartment (
    a_id serial PRIMARY KEY,
    a_name VARCHAR(50) NOT NULL,
    a_created_at TIMESTAMP NOT NULL
);

CREATE TABLE tenant (
    t_id serial PRIMARY KEY,
    t_name VARCHAR(50) NOT NULL,
    t_created_at TIMESTAMP NOT NULL
);