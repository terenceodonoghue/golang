CREATE TABLE ext_providers (
	ext_provider_id bigint PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
	provider_name text CHECK (char_length(provider_name) <= 50) UNIQUE
);

CREATE TABLE user_accounts (
	user_account_id bigint PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
	first_name text CHECK (char_length(last_name) <= 50),
	last_name text CHECK (char_length(last_name) <= 50)
);

CREATE TABLE user_sessions (
	user_session_id bigint PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
	user_account_id bigint NOT NULL,
	FOREIGN KEY (user_account_id) REFERENCES user_accounts (user_account_id) ON DELETE CASCADE
);

CREATE TABLE user_credentials (
	user_account_id bigint PRIMARY KEY,
	email_address text CHECK (char_length(email_address) <= 50) UNIQUE NOT NULL,
	password_hash text CHECK (char_length(password_hash) <= 250) NOT NULL,
	password_salt text CHECK (char_length(password_salt) <= 100) NOT NULL,
	FOREIGN KEY (user_account_id) REFERENCES user_accounts (user_account_id) ON DELETE CASCADE
);

CREATE TABLE user_credentials_ext (
	user_account_id bigint,
	ext_provider_id bigint,
	ext_provider_token text CHECK (char_length(ext_provider_token) <= 100) NOT NULL,
	FOREIGN KEY (user_account_id) REFERENCES user_accounts (user_account_id) ON DELETE CASCADE,
	FOREIGN KEY (ext_provider_id) REFERENCES ext_providers (ext_provider_id) ON DELETE CASCADE,
	PRIMARY KEY (user_account_id, ext_provider_id)
);
