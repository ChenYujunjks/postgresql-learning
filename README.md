# PostgreSQL Learning Project

This is a minimal PostgreSQL learning project designed to help you master core PostgreSQL features through raw SQL, with a focus on:

- Differences from MySQL
- Practical use of data types like `uuid`, `timestamptz`, and `jsonb`
- SQL functions and expressions
- Indexes and performance basics
- Row Level Security (RLS)

No ORMs. No frameworks. Just SQL.

---

## 📁 Project Structure

```bash
postgresql-learning/
├── init.sql         # Create tables, insert data, create indexes and functions
├── rls.sql          # Enable RLS and apply row-level access control policies
├── playground.sql   # Query test: JSON, functions, pagination, subqueries, etc.
└── README.md        # This file
```

---

## 🚀 Getting Started

### 1. Prepare your PostgreSQL database

Either:

- Use local PostgreSQL (`psql`), or
- Use Supabase's SQL Editor or CLI

### 2. Run the SQL scripts

```bash
# Step 1: Create schema and seed data
psql -U postgres -d your_db -f init.sql

# Step 2: (Optional) Enable RLS
psql -U postgres -d your_db -f rls.sql

# Step 3: Run test queries
psql -U postgres -d your_db -f playground.sql
```

---

## 🎯 Learning Goals

| Concept                                                | Covered In       |
| ------------------------------------------------------ | ---------------- |
| `uuid`, `timestamptz`, `jsonb`                         | `init.sql`       |
| SQL functions: `length`, `coalesce`, `now`, `interval` | `playground.sql` |
| JSONB field access & casting                           | `playground.sql` |
| Subqueries & pagination                                | `playground.sql` |
| User-defined SQL functions                             | `init.sql`       |
| Indexing & GIN index on JSONB                          | `init.sql`       |
| Row Level Security (RLS)                               | `rls.sql`        |

---

## ✅ Tips

- You can safely re-run `init.sql` — it uses `if not exists` / `drop` for idempotency.

---

## 📜 License

MIT — use freely for learning and teaching.
