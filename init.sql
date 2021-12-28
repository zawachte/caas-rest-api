CREATE TABLE accounts (
	account_id serial PRIMARY KEY,
	username VARCHAR ( 50 ) UNIQUE NOT NULL,
	password VARCHAR ( 50 ) NOT NULL,
	email VARCHAR ( 255 ) UNIQUE NOT NULL,
	created_on TIMESTAMP NOT NULL,
    last_login TIMESTAMP 
);

CREATE TABLE clusters (
    cluster_id SERIAL PRIMARY KEY,
    account_id INT NOT NULL,
    kubeconfig varchar(255) NOT NULL,
	created_on TIMESTAMP NOT NULL,
    FOREIGN KEY (account_id)
      REFERENCES accounts (account_id)
);