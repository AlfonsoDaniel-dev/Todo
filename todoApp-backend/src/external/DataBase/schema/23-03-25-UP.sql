
create extension if not exists "uuis-ossp";

create table users (
    id         UUID        not null default uuid_generate_v4(),
    user_name  VARCHAR(80) not NULL,
    email      VARCHAR(60) not NULL,
    password   VARCHAR(240 not null,
    created_at BIGINT      not null default extract(EPOCH from now(),
    updated_at BIGINT
    deleted_at BIGINT
    constraint users_id_pk primary key (id)
    constraint users_email_uq unique (email)
    constraint users_user_name_uq unique (user_name)
 );


create table tasks (
    id UUID not null default uuid_generate_v4()
    owner_id UUID not null,
    title VARCHAR(80) not null,
    body VARCHAR (240)not null,
    created_at BIGINT not null default extract(EPOCH from now(),
    updated_at BIGINT,
    deleted_at BIGINT,
    constraint tasks_id_pk primary key (id)
    constraint tasks_owner_id_fk foreign key (owner_id) references users (id) on delete restrict on UPDATED CASCADE
 );
