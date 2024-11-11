alter table if exists users drop constraint if exists user_primary_club;
alter table if exists users drop constraint if exists users_created_by_user;
alter table if exists users drop constraint if exists users_updated_by_user;
alter table if exists users drop constraint if exists users_deleted_by_user;

alter table if exists albums drop constraint if exists album_owner;
alter table if exists albums drop constraint if exists album_created_by_user;
alter table if exists albums drop constraint if exists album_updated_by_user;
alter table if exists albums drop constraint if exists album_deleted_by_user;

alter table if exists clubs  drop constraint if exists club_primary_contact;
alter table if exists clubs  drop constraint if exists club_country;
alter table if exists clubs drop constraint if exists club_created_by_user;
alter table if exists clubs drop constraint if exists club_updated_by_user;
alter table if exists clubs drop constraint if exists club_deleted_by_user;


alter table if exists folders drop constraint if exists folders_owner;
alter table if exists folders drop constraint if exists folders_created_by_user;
alter table if exists folders drop constraint if exists folders_updated_by_user;
alter table if exists folders drop constraint if exists folders_deleted_by_user;

drop table if exists albums;
drop table if exists countries;
drop table if exists clubs;
drop table if exists users;
