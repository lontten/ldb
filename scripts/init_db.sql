-- 创建表
CREATE TABLE IF NOT EXISTS users
(
    id         SERIAL PRIMARY KEY,
    age        INTEGER,
    name       VARCHAR(100),
    money      DECIMAL(10, 2),
    day1       TIMESTAMP,
    day2       TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);