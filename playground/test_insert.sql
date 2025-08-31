
insert into users (email, name, metadata) values
('cat@demo.com', 'Cat', '{"role": "admin"}'::jsonb);


insert into courses (user_id, title, location, start_time, end_time, extra)
select id, 'Chemistry', 'Lab 1', now(), now() + interval '90 min', '{"type": "lab"}'::jsonb
from users where email = 'cat@demo.com';


-- playground/test_rls.sql
-- 注意：必须用实际身份执行（模拟 Supabase 用户）
select * from courses;