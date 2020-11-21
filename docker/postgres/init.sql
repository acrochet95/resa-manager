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

CREATE TABLE reservation (
    r_id serial PRIMARY KEY,
    a_id serial REFERENCES apartment,
    t_id serial REFERENCES tenant,
    r_price int NOT NULL,
    r_description VARCHAR(100),
    r_from TIMESTAMP NOT NULL,
    r_to TIMESTAMP NOT NULL,
    r_paid BOOLEAN
);