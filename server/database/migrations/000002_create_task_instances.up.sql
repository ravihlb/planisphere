CREATE TABLE task_instances (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    def_id INTEGER NOT NULL,
    due_date DATETIME,
    status TEXT CHECK ( status in ('pending', 'done') ) NOT NULL DEFAULT 'pending',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME,
    FOREIGN KEY(def_id) REFERENCES task_definitions(id)
);
