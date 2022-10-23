CREATE TABLE companies
(
    id                  varchar       not null unique,
    name                varchar(15)   not null unique,
    description         varchar(3000),
    amount_of_employees integer       not null,
    registered          boolean       not null,
    type                varchar(15)   not null
);

CREATE TABLE users
(
    id            serial       not null unique,
    name          varchar(255) not null,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null
);
