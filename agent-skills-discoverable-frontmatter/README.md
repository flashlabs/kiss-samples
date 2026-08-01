# Agent Skills Discoverable Frontmatter

Companion examples for the blog post: **"Frontmatter FTW: How to Make Agent Skills Discoverable"** ([https://blog.skopow.ski/frontmatter-ftw-how-to-make-agent-skills-discoverable](https://blog.skopow.ski/frontmatter-ftw-how-to-make-agent-skills-discoverable))

This sample demonstrates how to structure AI agent skill files using YAML frontmatter to decouple discoverability triggers from execution details.

## Structure

- `skills/database-migration-checker.md`: A skill for validating zero-downtime PostgreSQL schema migrations.
- `skills/git-commit-generator.md`: A skill for generating conventional git commits.

## Frontmatter Pattern

Each skill specifies minimal metadata in its YAML header:

```yaml
---
name: database-migration-checker
description: Validates zero-downtime PostgreSQL schema migrations against expand-contract safety rules. Use when analyzing SQL migration scripts or PRs touching schema files.
---
```

When scanning available capabilities, the agent framework loads **only** the frontmatter into the primary context index (~15–20 tokens per skill). The heavy Markdown body is dynamically fetched via file inspection only when the frontmatter description matches the active user request.
