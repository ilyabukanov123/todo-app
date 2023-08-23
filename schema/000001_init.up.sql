create TABLE users
(
    id serial primary key,
    name varchar(255) not null,
    username varchar(255) not null unique,
    password_hash varchar(255) not null
);

create table todo_lists
(
    id serial primary key,
    title varchar(255) not null,
    description varchar(255)
);


create table users_lists
(
    id serial primary key,
    user_id integer REFERENCES users(id) ON DELETE CASCADE not null,
    list_id integer REFERENCES todo_lists(id) ON DELETE CASCADE not null
);


create table todo_items
(
    id serial primary key,
    title varchar(255) not null,
    description varchar(255),
    done boolean not null default false
);

create table lists_items
(
    id serial primary key,
    item_id integer REFERENCES todo_items(id) ON DELETE CASCADE not null,
    list_id integer REFERENCES todo_lists(id) ON DELETE CASCADE not null
);