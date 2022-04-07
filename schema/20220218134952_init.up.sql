CREATE TABLE list_users
(
    id serial not null unique,
    email varchar(125) not null unique,
    password_hash varchar (255) not null
);

CREATE TABLE list_titles
(
    id serial not null  unique,
    title varchar(255),
    script text
);

CREATE TABLE user_titles
(
    id serial not null unique,
    user_id int references list_users (id) on delete cascade not null,
    title_id int references list_titles(id) on delete cascade not null
);


