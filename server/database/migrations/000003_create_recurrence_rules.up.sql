CREATE TABLE recurrence_rules (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    def_id INTEGER NOT NULL,
    frequency TEXT NOT NULL,
    interval TEXT NOT NULL,
    end_after_occurrences TEXT DEFAULT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME,
    FOREIGN KEY(def_id) REFERENCES task_definitions(id)
);
