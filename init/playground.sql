-- playground.sql
-- 基本函数测试
select id, name, length(name) as name_length from users;


select
email,
coalesce(metadata->>'nickname', '无昵称') as nickname,
metadata->>'role' as role
from users;


-- 函数调用
select
title,
start_time,
end_time,
public.course_duration_min(courses.*) as duration_min
from courses;


-- limit + offset 分页
select * from courses order by start_time desc limit 1 offset 1;


-- 子查询：每位用户的课程数
select
u.name,
(select count(*) from courses c where c.user_id = u.id) as course_count
from users u;