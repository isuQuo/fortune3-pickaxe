-- +migrate Up
CREATE TABLE if not exists goals (
    id TEXT PRIMARY KEY,
    name TEXT,
    total_needed REAL,
    current_savings REAL,
    monthly_contribution REAL,
    start_date TEXT,
    end_date TEXT,
    months_to_goal INTEGER
);

-- +migrate Down
DROP TABLE goals;