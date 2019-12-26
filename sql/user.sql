CREATE TABLE users (
	ID SERIAL PRIMARY KEY,
	uuid VARCHAR ( 64 ) UNIQUE,
	email VARCHAR ( 32 ),
	real_name VARCHAR ( 64 ),
	tel VARCHAR ( 16 ),
status VARCHAR ( 16 )
)