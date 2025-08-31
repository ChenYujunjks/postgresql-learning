-- init.sql
create extension if not exists "pgcrypto";


-- users 表
create table if not exists users (
id uuid primary key default gen_random_uuid(),
email text unique not null,
name text,
is_active boolean default true,
metadata jsonb,
created_at timestamptz default now()
);


-- courses 表
create table if not exists courses (
id bigserial primary key,
user_id uuid references users(id),
title text not null,
location text,
start_time timestamptz not null,
end_time timestamptz not null,
extra jsonb,
created_at timestamptz default now()
);


-- 索引
create index if not exists idx_users_email on users(email);
create index if not exists idx_users_created on users(created_at);
create index if not exists idx_courses_user on courses(user_id);
create index if not exists idx_courses_start on courses(start_time);
create index if not exists idx_courses_extra on courses using gin (extra);


-- 初始数据
insert into users (email, name, metadata) values
('alice@example.com', 'Alice', '{"role": "student", "interests": ["math", "art"]}'::jsonb),
('bob@example.com', 'Bob', '{"role": "teacher", "department": "science"}'::jsonb);


insert into courses (user_id, title, location, start_time, end_time, extra)
select id, 'Algebra I', 'Room 101', now(), now() + interval '1 hour', '{"level": "beginner"}'::jsonb
from users where email = 'alice@example.com';


insert into courses (user_id, title, location, start_time, end_time, extra)
select id, 'Physics 101', 'Lab 3', now(), now() + interval '2 hours', '{"lab": true, "equipment": ["scope", "sensor"]}'::jsonb
from users where email = 'bob@example.com';


-- 函数示例：计算课程时长（分钟）
create or replace function public.course_duration_min(course courses)
returns int as $$
begin
return extract(epoch from course.end_time - course.start_time) / 60;
end;
$$ language plpgsql;