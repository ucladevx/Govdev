package postgresql

const (
	userCreateTable = `
		CREATE TABLE IF NOT EXISTS users (
			user_id varchar(20) PRIMARY KEY,
			username varchar(128) NOT NULL,
			email varchar(512) NOT NULL,
			first_name varchar(128) NOT NULL,
			last_name varchar(128) NOT NULL,
			created_at timestamptz DEFAULT NOW(),
			updated_at timestamptz DEFAULT NOW(),

			UNIQUE ("email")
		);
	`
)
