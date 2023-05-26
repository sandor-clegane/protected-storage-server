-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public.users (
    id varchar(45) primary key,
    login varchar(45) unique not null,
    password varchar(45) not null
    );

CREATE TABLE IF NOT EXISTS public.raw_data (
    name varchar(45) unique not null,
    data_type int2 not null,
    data bytea,
    user_id varchar(45) references public.users (id)
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users,raw_data;
-- +goose StatementEnd
