-- 创建 pgcrypto 扩展（用于生成 uuid）
create extension if not exists "pgcrypto";

-- 创建用户表（用于演示 uuid 主键）
drop table if exists users cascade;
create table users (
  id uuid primary key default gen_random_uuid(), -- PostgreSQL 原生支持 uuid（MySQL 需用 char(36) 代替）
  email text unique not null,                    -- text 字段无长度限制，MySQL 通常会用 varchar
  name text,
  is_active boolean default true,                -- PostgreSQL 原生 boolean 类型（MySQL 通常用 tinyint）
  metadata jsonb,                                -- jsonb 支持结构化数据 + 查询优化（MySQL 虽支持 JSON 但功能弱）
  created_at timestamptz default now()           -- 含时区时间戳，适合 SaaS/多时区应用（MySQL 无时区支持）
);

-- 创建课程表，绑定 user_id 外键
drop table if exists courses cascade;
create table courses (
  id bigserial primary key,                      -- 自增 bigint 主键（bigserial 是 PostgreSQL 专属）
  user_id uuid references users(id),             -- 外键关联
  title text not null,
  location text,
  start_time timestamptz not null,
  end_time timestamptz not null,
  extra jsonb,                                   -- 附加字段（如教室/类型等结构数据）
  created_at timestamptz default now()
);
