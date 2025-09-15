## 1. 什么是 PostgreSQL 的 superuser

- 在 PostgreSQL 里，用户（role）有权限等级。
- **superuser = 超级用户**，相当于数据库里的 **root** 或 **admin**。
- 默认安装时，会创建一个名为 **`postgres`** 的 superuser。

---

## 2. superuser 拥有哪些权限

- 可以访问、修改所有数据库和所有对象（表、视图、函数、扩展等）。
- 可以绕过所有权限检查（grant/revoke 都对它无效）。
- 可以创建、删除数据库。
- 可以创建其他用户（role）、赋予权限。
- 可以加载扩展、修改系统配置。
- 可以直接访问操作系统层面的一些功能（比如 `COPY` 从文件导入/导出）。

👉 简单说：**superuser 在数据库里没有任何限制**。

---

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

---

## 5. 类比理解

- Linux 里的 `root` 用户。
- Windows 里的 **Administrator**。
- MySQL 里的 `root` 用户。

PostgreSQL superuser 的地位就是 **数据库里的系统管理员**。

---

✅ 总结：
**PostgreSQL 的 superuser 就是数据库的“全权管理员”，默认是 `postgres` 用户，你安装时设的密码就是它的密码。**
