CREATE TABLE IF NOT EXISTS users (
	id int generated always as identity ( cache 10 ) primary key
    , uid varchar(36) not null unique
    , username text not null unique
    , password text not null
 )