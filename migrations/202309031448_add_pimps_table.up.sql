CREATE table pimps (
	id SERIAL,
	firstname varchar(30) NOT NULL,
	lastname varchar(30) NOT NULL,
	age INTEGER,
	ADD CONSTRAINT pimps_pk PRIMARY KEY (id) 
);
