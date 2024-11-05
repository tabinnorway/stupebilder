insert into users(created_at, email, username, first_name, last_name, primary_phone, primary_club_id)
values (current_timestamp, 'terje@bergesen.info', 'terje', 'Terje Anthon', 'Bergesen', '+47 900 12 465', (select id from clubs where club_name = 'Bergen Stupeklubb'))

insert into users(created_at, email, username, first_name, last_name, primary_club_id)
values (current_timestamp, 'dagligleder@bergen-stupeklubb.no', 'pj', 'Paul Joachim', 'Thorsen', (select id from clubs where club_name = 'Bergen Stupeklubb'))

