alter table if exists users  drop constraint if exists user_primary_club;
alter table if exists albums drop constraint if exists album_owner;
alter table if exists clubs  drop constraint if exists club_primary_contact;
alter table if exists clubs  drop constraint if exists club_country;

drop table if exists albums;
drop table if exists countries;
drop table if exists clubs;
drop table if exists users;
