--user
--STORING PLAIN TEXT PASSWORD VERY BAD DO NOT DO EVER
--WILL CHANGE LATER
PRAGMA foreign_keys=OFF;
BEGIN TRANSACTION;
CREATE TABLE User (
  id integer primary key autoincrement,
  username varchar(100),
  password varchar(1000),
  email varchar(100)
);

--category
CREATE TABLE Category(
  id integer primary key autoincrement,
  name varchar(1000) not null,
);
COMMIT;

--task

PRAGMA foreign_keys=OFF;
BEGIN TRANSACTION;
CREATE TABLE Article (
  id integer primary key autoincrement,
  title varchar(100),
  content text,
  created_date timestamp,
  last_modified_at timestamp,
  cat_id references category(id),
  user_id references user(id),
);

COMMIT;
