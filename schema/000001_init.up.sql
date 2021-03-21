CREATE TABLE queue
(
    id serial not null unique,
    value varchar(255) not null,
    datetime date not null
)