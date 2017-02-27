--user
--STORING PLAIN TEXT PASSWORD VERY BAD DO NOT DO EVER
--WILL CHANGE LATER
PRAGMA foreign_keys=OFF;
CREATE TABLE User (
  username varchar(100) primary key,
  password varchar(1000),
  email varchar(100)
);

--category
CREATE TABLE Category (
  name varchar(1000) primary key not null
);

--task
PRAGMA foreign_keys=OFF;
CREATE TABLE Article (
  username varchar(100) not null references User(username),
  title varchar(100) not null,
  content text,
  created_date timestamp,
  last_modified_at timestamp,
  cat_name references Category(name),
  primary key (username, title)
);
