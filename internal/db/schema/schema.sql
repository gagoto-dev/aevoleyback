CREATE TABLE IF NOT EXISTS user (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    password TEXT NOT NULL,
    email TEXT NOT NULL,
    createdAt timestamp DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS blog_entry (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    authorId UUID NOT NULL,
    title TEXT NOT NULL,

    FOREIGN KEY(authorId) REFERENCES users(id)
);

