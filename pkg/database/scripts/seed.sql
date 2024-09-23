CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    email TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_users_email ON users(email);

CREATE TABLE IF NOT EXISTS clients (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    dpi TEXT NOT NULL,
    name TEXT NOT NULL,
    client_type INTEGER NOT NULL CHECK (client_type IN (0, 1))
);

CREATE TABLE IF NOT EXISTS clients_history (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    attended_by INTEGER NOT NULL,
    client_id INTEGER NOT NULL,
    required_operations INTEGER NOT NULL,
    attended_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (attended_by) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE
    FOREIGN KEY (client_id) REFERENCES clients(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_clients_dpi ON clients(dpi);
CREATE INDEX IF NOT EXISTS idx_clients_client_type ON clients(client_type);

INSERT INTO users (email, password_hash)
SELECT 'admin@example.com', '<HASHED_PASSWORD>'
WHERE NOT EXISTS (
    SELECT 1 FROM users WHERE email = 'admin@example.com'
);
