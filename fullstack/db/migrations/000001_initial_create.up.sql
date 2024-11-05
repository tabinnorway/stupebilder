create table users (
    id                  serial primary key,
    created_at          timestamp not null,
    updated_at          timestamp,
    email               varchar(128) unique not null,
    passwd              varchar(128),
    username            varchar(128),
    first_name          varchar(128) not null,
    last_name           varchar(128) not null,
    primary_phone       varchar(64),
    confirmed           boolean not null default false,
    primary_club_id     int
);

create table clubs (
    id                  serial primary key,
    created_at          timestamp not null,
    updated_at          timestamp null,
    email               varchar(128) not null,
    club_name           varchar(64) not null,
    short_name          varchar(16) not null,
    phone_number        varchar(64),
    street_address      varchar(256),
    postal_code         varchar(16),
    city                varchar(255),
    country_id          int,
    primary_contact_id  int
);

create table countries
(
    id                   serial primary key,
    created_at           timestamp not null,
    updated_at           timestamp,
    country_name         varchar(128) not null unique,
    country_code_a2      varchar(2),
    country_code_a3      varchar(3),
    country_code_numeric varchar(3),
    country_phone_prefix varchar(8),
    tld                  varchar(8)
);

create table albums
(
    id                   serial primary key,
    created_at           timestamp not null,
    updated_at           timestamp,
    album_folder         varchar(4096) not null unique,
    title                varchar(128),
    datestring           varchar(32),
    owner_id             int
);

alter table clubs
    add constraint club_primary_contact foreign key (primary_contact_id) references users(id)
;
alter table clubs
    add constraint club_country foreign key (country_id) references countries(id)
;
alter table users
    add constraint user_primary_club foreign key (primary_club_id) references clubs(id)
;
alter table albums
    add constraint album_owner foreign key (owner_id) references users(id)
;
