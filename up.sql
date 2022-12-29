CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	email TEXT UNIQUE NOT NULL,
	name TEXT UNIQUE NOT NULL,
	password TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS workspaces (
	id SERIAL PRIMARY KEY,
	title TEXT NOT NULL,
	logo TEXT
);

CREATE TABLE IF NOT EXISTS users_workspaces(
	user_id INT REFERENCES users(id),
	workspace_id INT,	
	UNIQUE (user_id, workspace_id)
);

CREATE TABLE IF NOT EXISTS boards (
	id SERIAL PRIMARY KEY,
	title TEXT NOT NULL,
	background TEXT,
	workspace_id INT REFERENCES workspaces(id)
);

CREATE TABLE IF NOT EXISTS lists (
	id SERIAL PRIMARY KEY,
	title TEXT NOT NULL,
	board_id INT REFERENCES boards(id)
);

CREATE TABLE IF NOT EXISTS cards (
	id SERIAL PRIMARY KEY,
	title TEXT NOT NULL,
	list_id INT REFERENCES lists(id)
);