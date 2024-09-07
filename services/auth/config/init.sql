CREATE TABLE user_profiles (
	user_id serial PRIMARY KEY,
	first_name text CHECK (char_length(last_name) <= 50),
	last_name text CHECK (char_length(last_name) <= 50)
);

CREATE TABLE user_credentials (
	user_id int NOT NULL,
	email_address text CHECK (char_length(email_address) <= 50) UNIQUE,
	password_hash text CHECK (char_length(password_hash) <= 250),
	password_salt text CHECK (char_length(password_salt) <= 100),
	FOREIGN KEY (user_id) REFERENCES user_profiles (user_id)
);