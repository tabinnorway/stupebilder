create table if not exists users (
    id                  varchar(36) unique not null,
    created_at          timestamp not null,
    created_by_email    varchar(256) not null,
    updated_at          timestamp,
    updated_by          varchar(36),
    deleted_at          timestamp,
    deleted_by          varchar(36),
    email               varchar(128) unique not null,
    passwd              varchar(128),
    username            varchar(128),
    first_name          varchar(128) not null,
    last_name           varchar(128) not null,
    primary_phone       varchar(64),
    confirmed           boolean not null default false,
    primary_club_id     varchar(36) 
);

create table if not exists clubs (
    id                  varchar(36) unique not null,
    created_at          timestamp not null,
    created_by          varchar(36) not null,
    updated_at          timestamp,
    updated_by          varchar(36),
    deleted_at          timestamp,
    deleted_by          varchar(36),
    email               varchar(128) not null,
    club_name           varchar(64) not null,
    short_name          varchar(16) not null,
    phone_number        varchar(64),
    street_address      varchar(256),
    postal_code         varchar(16),
    city                varchar(255),
    country_id          int,
    primary_contact_id  varchar(36) 
);

create table if not exists countries
(
    id                   serial primary key,
    created_at           timestamp not null,
    created_by           varchar(36) not null,
    updated_at           timestamp,
    updated_by           varchar(36),
    deleted_at           timestamp,
    deleted_by           varchar(36),
    country_name         varchar(128) not null unique,
    country_code_a2      varchar(2),
    country_code_a3      varchar(3),
    country_code_numeric varchar(3),
    country_phone_prefix varchar(8),
    tld                  varchar(8)
);

create table if not exists albums
(
    id                  varchar(36) unique not null,
    created_at          timestamp not null,
    created_by          varchar(36) not null,
    updated_at          timestamp,
    updated_by          varchar(36),
    deleted_at          timestamp,
    deleted_by          varchar(36),
    album_path          varchar(4096) not null unique,
    title               varchar(128) not null,
    datestring          varchar(32),
    owner_id            varchar(36)
);

create table if not exists folders
(
    id                  varchar(36) unique not null,
    created_at          timestamp not null,
    created_by          varchar(36) not null,
    updated_at          timestamp,
    updated_by          varchar(36),
    deleted_at          timestamp,
    deleted_by          varchar(36),
    folder_path         varchar(4096) not null unique,
    title               varchar(128) not null,
    num_images          int,
    owner_id            varchar(36)
);

alter table if exists clubs
    add constraint club_primary_contact foreign key (primary_contact_id) references users(id)
;
alter table if exists clubs
    add constraint club_country foreign key (country_id) references countries(id)
;
alter table if exists clubs
    add constraint club_created_by_user foreign key (created_by) references users(id)
;
alter table if exists clubs
    add constraint club_updated_by_user foreign key (updated_by) references users(id)
;
alter table if exists clubs
    add constraint club_deleted_by_user foreign key (deleted_by) references users(id)
;

alter table if exists users
    add constraint user_primary_club foreign key (primary_club_id) references clubs(id)
;
alter table if exists users
    add constraint users_updated_by_user foreign key (updated_by) references users(id)
;
alter table if exists users
    add constraint users_deleted_by_user foreign key (deleted_by) references users(id)
;

alter table if exists albums
    add constraint album_owner foreign key (owner_id) references users(id)
;
alter table if exists albums
    add constraint album_created_by_user foreign key (created_by) references users(id)
;
alter table if exists albums
    add constraint album_updated_by_user foreign key (updated_by) references users(id)
;
alter table if exists albums
    add constraint album_deleted_by_user foreign key (deleted_by) references users(id)
;

alter table if exists folders
    add constraint folders_owner foreign key (owner_id) references users(id)
;
alter table if exists folders
    add constraint folders_created_by_user foreign key (created_by) references users(id)
;
alter table if exists folders
    add constraint folders_updated_by_user foreign key (updated_by) references users(id)
;
alter table if exists folders
    add constraint folders_deleted_by_user foreign key (deleted_by) references users(id)
;

