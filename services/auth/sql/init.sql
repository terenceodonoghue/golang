CREATE TABLE ext_providers (
	ext_provider_id serial PRIMARY KEY,
	provider_name text CHECK (char_length(provider_name) <= 50) UNIQUE
);

CREATE TABLE user_accounts (
	user_account_id serial PRIMARY KEY,
	first_name text CHECK (char_length(last_name) <= 50),
	last_name text CHECK (char_length(last_name) <= 50)
);

CREATE TABLE user_credentials (
	user_account_id int PRIMARY KEY,
	email_address text CHECK (char_length(email_address) <= 50) UNIQUE,
	password_hash text CHECK (char_length(password_hash) <= 250) NOT NULL,
	password_salt text CHECK (char_length(password_salt) <= 100) NOT NULL,
	FOREIGN KEY (user_account_id) REFERENCES user_accounts (user_account_id)
);

CREATE TABLE user_credentials_ext (
	user_account_id int,
	ext_provider_id int,
	ext_provider_token text CHECK (char_length(ext_provider_token) <= 100) NOT NULL,
	FOREIGN KEY (user_account_id) REFERENCES user_accounts (user_account_id),
	FOREIGN KEY (ext_provider_id) REFERENCES ext_providers (ext_provider_id),
	PRIMARY KEY (user_account_id, ext_provider_id)
);
