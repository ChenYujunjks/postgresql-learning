-- rls.sql
alter table courses enable row level security;


create policy "Self can read" on courses for select using (auth.uid() = user_id);
create policy "Self can insert" on courses for insert with check (auth.uid() = user_id);
create policy "Self can update" on courses for update using (auth.uid() = user_id);
create policy "Self can delete" on courses for delete using (auth.uid() = user_id);