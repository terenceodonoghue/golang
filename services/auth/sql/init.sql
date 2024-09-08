CREATE TABLE user_accounts (
	user_id serial PRIMARY KEY,
	first_name text CHECK (char_length(last_name) <= 50),
	last_name text CHECK (char_length(last_name) <= 50)
);

CREATE TABLE user_credentials (
	user_id int PRIMARY KEY,
	email_address text CHECK (char_length(email_address) <= 50) UNIQUE,
	password_hash text CHECK (char_length(password_hash) <= 250),
	password_salt text CHECK (char_length(password_salt) <= 100),
	FOREIGN KEY (user_id) REFERENCES user_accounts (user_id)
);

CREATE TABLE external_providers (
	external_provider_id serial PRIMARY KEY,
	provider_name text CHECK (char_length(provider_name) <= 50) UNIQUE
);

CREATE TABLE user_credentials_ext (
	user_id int,
	external_provider_id int,
	external_provider_token text CHECK (char_length(external_provider_token) <= 100) NOT NULL,
	FOREIGN KEY (user_id) REFERENCES user_accounts (user_id),
	FOREIGN KEY (external_provider_id) REFERENCES external_providers (external_provider_id),
	PRIMARY KEY (user_id, external_provider_id)
);
