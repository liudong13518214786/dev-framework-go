CREATE TABLE blogs
(
  id SERIAL primary key ,
  uuid VARCHAR(64) UNIQUE,
  useruuid VARCHAR(64),
  title  VARCHAR(64),
  img_url VARCHAR(64),
	info text,
	tag jsonb,
	build_time TIMESTAMP,
	read_num int
)