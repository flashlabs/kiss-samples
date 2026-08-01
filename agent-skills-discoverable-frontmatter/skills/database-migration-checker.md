---
name: database-migration-checker
description: Validates zero-downtime PostgreSQL schema migrations against expand-contract safety rules. Use when analyzing SQL migration scripts or PRs touching schema files.
---

# Database Migration Checker

Follow these rules to verify zero-downtime safety for PostgreSQL DDL migrations:

1. **Destructive Operations:** Never drop columns or tables in a single deployment step.
2. **Column Additions:** Ensure new columns are either `NULLABLE` or have a safe default without blocking locks.
3. **Index Creation:** Indexes on existing tables must use `CREATE INDEX CONCURRENTLY`.
4. **Expand-Contract Pattern:** Verify that schema expansions do not break older running versions of the application microservice.
