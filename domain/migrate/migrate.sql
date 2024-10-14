CREATE TABLE IF NOT EXISTS "users" (
    "id" varchar primary key not null,
    "first_name" varchar,
    "last_name" varchar,
    "username" varchar not null,
    "email" varchar not null,
    "password" varchar not null
);