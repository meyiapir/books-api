CREATE TABLE books
(
    id serial not null unique,
    title varchar(255) not null,
    author varchar(255) not null,
    year int not null
)