-- 查询字符串函数
select id, name, length(name) as name_length from users;


-- 使用 coalesce 和 jsonb 查询
select
email,
coalesce(metadata->>'nickname', '无昵称') as nickname,
metadata->>'role' as role
from users;


-- 查询课程时长函数
select
title,
start_time,
end_time,
public.course_duration_min(courses.*) as duration_min
from courses;


-- limit + offset
select * from courses order by start_time desc limit 1 offset 1;


-- 子查询示例：找出每位用户的课程数
select
u.name,
(select count(*) from courses c where c.user_id = u.id) as course_count
from users u;