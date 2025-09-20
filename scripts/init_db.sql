-- 创建表
CREATE TABLE IF NOT EXISTS users
(
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(100) NOT NULL,
    email      VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS products
(
    id    SERIAL PRIMARY KEY,
    name  VARCHAR(200)   NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    stock INTEGER DEFAULT 0
);

-- 插入示例数据
INSERT INTO users (name, email)
VALUES ('John Doe', 'john@example.com'),
       ('Jane Smith', 'jane@example.com')
ON CONFLICT
    (email)
    DO NOTHING;

INSERT INTO products (name, price, stock)
VALUES ('Product A', 19.99, 100),
       ('Product B', 29.99, 50),
       ('Product C', 9.99, 200);