create extension if not exists "uuid-ossp";

create table schools (
    id uuid not null primary key default uuid_generate_v4() ,
    name varchar not null,
    city varchar not null
);

create table user_roles (
    id serial not null primary key,
    name varchar not null
);

create table users (
    id serial not null primary key,
    username varchar not null unique,
    encrypted_password varchar not null,
    school_id uuid references schools(id) on delete set null,
    role_id integer references user_roles(id) on delete set null
);
