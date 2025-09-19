# 🐘 PostgreSQL 实用语法说明

## 1. 什么是 PostgreSQL 的 superuser

- 在 PostgreSQL 里，用户（role）有权限等级。
- **superuser = 超级用户**，相当于数据库里的 **root** 或 **admin**。
- 默认安装时，会创建一个名为 **`postgres`** 的 superuser。

## 2. superuser 拥有哪些权限

- 可以访问、修改所有数据库和所有对象（表、视图、函数、扩展等）。
- 可以绕过所有权限检查（grant/revoke 都对它无效）。
- 可以创建、删除数据库。
- 可以创建其他用户（role）、赋予权限。
- 可以加载扩展、修改系统配置。
- 可以直接访问操作系统层面的一些功能（比如 `COPY` 从文件导入/导出）。

👉 简单说：**superuser 在数据库里没有任何限制**。

## 3. 和普通用户的区别

- **普通用户**：只能访问自己有权限的数据库对象，需要管理员授予。
- **superuser**：不受限制，可以干任何事。

例如：

```sql
-- 用 superuser 创建一个新用户
CREATE USER kevin WITH PASSWORD '123456';

-- 给 kevin 一些权限
GRANT CONNECT ON DATABASE mydb TO kevin;
GRANT SELECT, INSERT ON ALL TABLES IN SCHEMA public TO kevin;
```

普通用户 `kevin` 就只能做被授予的事情。

---

## 4. 为什么安装时要设置 superuser 密码

因为：

- 这是你管理数据库的“万能钥匙”。
- 后续你要用 `psql -U postgres` 登录数据库时就要用这个密码。
- 生产环境必须保证这个密码安全，否则任何人都能完全控制数据库。

## 5. 类比理解

- Linux 里的 `root` 用户。
- Windows 里的 **Administrator**。
- MySQL 里的 `root` 用户。

---

PostgreSQL superuser 的地位就是 **数据库里的系统管理员**。

## 1. **数据库操作**

```sql
-- 创建数据库
CREATE DATABASE mydb;

-- 删除数据库
DROP DATABASE mydb;

-- 切换数据库 (在 psql 里)
\c mydb;
```

---

## 2. **用户与权限**

```sql
-- 创建用户
CREATE USER kevin WITH PASSWORD 'mypassword';

-- 创建超级用户
CREATE ROLE admin SUPERUSER LOGIN PASSWORD 'secret';

-- 给用户赋权限
GRANT ALL PRIVILEGES ON DATABASE mydb TO kevin;

-- 删除用户
DROP USER kevin;
```

---

## 3. **表操作**

```sql
-- 创建表
CREATE TABLE users (
    id SERIAL PRIMARY KEY,         -- 自增主键
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE,
    age INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 修改表 (加一列)
ALTER TABLE users ADD COLUMN phone VARCHAR(20);

-- 删除表
DROP TABLE users;
```

---

## 4. **数据操作 (CRUD)**

```sql
-- 插入数据
INSERT INTO users (name, email, age) VALUES ('Alice', 'alice@example.com', 25);

-- 查询数据
SELECT * FROM users;
SELECT name, age FROM users WHERE age > 20 ORDER BY age DESC;

-- 更新数据
UPDATE users SET age = 26 WHERE name = 'Alice';

-- 删除数据
DELETE FROM users WHERE id = 1;
```

---

## 5. **查询技巧**

```sql
-- 限制行数
SELECT * FROM users LIMIT 5;

-- 聚合函数
SELECT COUNT(*), AVG(age), MAX(age), MIN(age) FROM users;

-- 分组统计
SELECT age, COUNT(*) FROM users GROUP BY age;

-- 多表连接
SELECT u.name, o.amount
FROM users u
JOIN orders o ON u.id = o.user_id;

-- 子查询
SELECT name FROM users WHERE id IN (SELECT user_id FROM orders WHERE amount > 100);
```

---

## 6. **特殊数据类型 (Postgres 强项)**

```sql
-- JSON 字段
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    data JSONB
);

INSERT INTO products (data) VALUES ('{"name": "Laptop", "price": 1200}');

-- JSON 查询
SELECT data->>'name' AS product_name FROM products;

-- 数组字段
CREATE TABLE tags (
    id SERIAL PRIMARY KEY,
    labels TEXT[]
);

INSERT INTO tags (labels) VALUES (ARRAY['tech','ai','db']);
SELECT * FROM tags WHERE 'ai' = ANY(labels);
```

---

## 7. **索引与性能**

```sql
-- 创建索引
CREATE INDEX idx_users_email ON users(email);

-- 唯一索引 (避免重复)
CREATE UNIQUE INDEX uniq_email ON users(email);

-- 全文搜索索引
CREATE INDEX idx_users_name ON users USING gin(to_tsvector('english', name));
```

---

## 8. **事务 (保证 ACID)**

```sql
BEGIN;

UPDATE users SET age = age + 1 WHERE id = 2;
DELETE FROM orders WHERE id = 10;

COMMIT;  -- 提交事务
-- ROLLBACK;  -- 回滚事务
```

---

## 9. **psql 专属命令（不是 SQL）**

```bash
\l       -- 列出所有数据库
\c mydb  -- 切换数据库
\dt      -- 查看当前数据库的表
\d users -- 查看表结构
\du      -- 查看用户
\x       -- 切换扩展显示模式
\q       -- 退出
```
