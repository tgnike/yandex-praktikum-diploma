CREATE TABLE IF NOT EXISTS orders (
    
    ordernumber varchar(20) not null primary key
    , useruid varchar(36) not null
    , balance float not null
    , status varchar(25) not null
     , date TIMESTAMP WITH TIME ZONE
 )