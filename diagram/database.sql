# FOR TUTORIAL SQL QUERY
# https://www.w3schools.com/sql/default.asp

CREATE DATABASE IF NOT EXISTS auth;

USE auth;

CREATE TABLE IF NOT EXISTS user_role (
    id int NOT NULL AUTO_INCREMENT,
    role varchar(255) NOT NULL ,
    created_at timestamp NOT NULL DEFAULT now(),

    UNIQUE (role),
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS user (
    id int NOT NULL AUTO_INCREMENT,
    email varchar(255) NOT NULL ,
    password varchar(255) NOT NULL,
    first_name varchar(255),
    last_name varchar(255),
    user_role_id int NOT NULL,
    company_id int NOT NULL,
    principal_id int,
    distributor_id int,
    buyer_id int,
    token_version int,
    is_verified boolean NOT NULL DEFAULT false,
    is_delete boolean NOT NULL DEFAULT false,
    created_at timestamp NOT NULL DEFAULT now(),
    updated_at timestamp NOT NULL DEFAULT now(),
    created_by varchar(255) NOT NULL,
    updated_by varchar(255) NOT NULL,

    UNIQUE (email),
    PRIMARY KEY (id),
    FOREIGN KEY (user_role_id) REFERENCES user_role(id)
);

CREATE TABLE IF NOT EXISTS data_scope (
    id int NOT NULL AUTO_INCREMENT,
    user_id int NOT NULL,
    principal_id int,
    distributor_id int,
    buyer_id int,
    is_delete boolean NOT NULL DEFAULT false,
    created_at timestamp NOT NULL DEFAULT now(),
    updated_at timestamp NOT NULL DEFAULT now(),
    created_by varchar(255) NOT NULL,
    updated_by varchar(255) NOT NULL,

    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES user(id)
);

CREATE TABLE IF NOT EXISTS client (
    id int NOT NULL AUTO_INCREMENT,
    user_id int NOT NULL,
    application_name varchar(255) NOT NULL,
    client_secret varchar(255) NOT NULL,
    is_delete boolean NOT NULL DEFAULT false,
    created_at timestamp NOT NULL DEFAULT now(),
    updated_at timestamp NOT NULL DEFAULT now(),
    created_by varchar(255) NOT NULL,
    updated_by varchar(255) NOT NULL,

    PRIMARY KEY (id)
    FOREIGN KEY (user_id) REFERENCES user(id)
);

#drop database auth;