CREATE DATABASE gochat;

CREATE EXTENSION "uuid-ossp";

CREATE TABLE IF NOT EXISTS USERS
(
	uid uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
	userid text NOT NULL UNIQUE,
	password text NOT NULL,
	first_name text NOT NULL,
	last_name text NOT NULL,
	email text NOT NULL,
	created_date date NOT NULL DEFAULT 'now'::text::date,
	last_updated timestamp NOT NULL DEFAULT now(),
	last_signin timestamp,
	active bool DEFAULT true
);

INSERT INTO users
(userid, password, first_name, last_name, email, created_date, last_updated, last_signin, active)
VALUES
('', '', '', '', '', 'now'::text::date, now(), now(), true)
RETURNING uid;

DELETE FROM  users
WHERE uid='ed1bc9da-65e7-4e63-9ce0-e23863e24950';

SELECT uid, userid, password, first_name, last_name, email, created_date, last_updated, last_signin, active
FROM users
WHERE uid='';

SELECT uid, userid, password, first_name, last_name, email, created_date, last_updated, last_signin, active
FROM users
WHERE userid='';

UPDATE public.users
SET userid='', password='', first_name='', last_name='', email='', 
created_date='now'::text::date, last_updated=now(), last_signin='', active=true
WHERE uid=uuid_generate_v4();





