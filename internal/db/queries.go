package db

/* INIT TABLES */
const (
	create_users_table_query = `
	CREATE TABLE IF NOT EXISTS users
	(
		id INTEGER PRIMARY KEY,
		login TEXT UNIQUE,
		password TEXT
	);
`

	create_messages_table_query = `
	CREATE TABLE IF NOT EXISTS messages
	(
	    id INTEGER PRIMARY KEY,
	    message TEXT,
	    sender INTEGER NOT NULL,
	    recipient INTEGER NOT NULL,
	    chat_id INTEGER,
	    created_at DATE DEFAULT (datetime('now', 'utc'))
	);
`

	create_chats_table_query = `
	CREATE TABLE IF NOT EXISTS chats
	(
	    id INTEGER PRIMARY KEY,
	    name TEXT,
	    created_at DATE DEFAULT (datetime('now', 'utc')),
	    updated_at DATE
	);
`

	create_chat_members_table = `
	create TABLE IF NOT EXISTS chat_members
	(
	    id INTEGER PRIMARY KEY,
	    user_id INTEGER NOT NULL,
	    chat_id INTEGER NOT NULL
	);
`
)
