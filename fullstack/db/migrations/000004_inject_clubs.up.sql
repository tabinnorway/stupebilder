insert into clubs ( created_at, email, club_name, short_name, phone_number, street_address, postal_code, city, country_id )
    values ( current_timestamp, 'dagligleder@bergen-stupeklubb.no', 'Bergen Stupeklubb', 'BStK', '+47 901 66 005', 'Lungegårdskaien 40', '5015', 'Bergen', (select id from countries where country_name = 'Norway'));
