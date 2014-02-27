CREATE TABLE notes (
  id SERIAL NOT NULL PRIMARY KEY,
  tags varchar(100) not null,
  note text not null,
  created_at timestamp default current_timestamp
);
