---
name: git-commit-generator
description: Generates conventional git commit messages from staged file diffs. Use when writing commit messages or summarizing code changes before pushing.
---

# Git Commit Generator

Generate commit messages strictly following the Conventional Commits specification:

1. Format: `<type>(<scope>): <short summary>`
2. Types: `feat`, `fix`, `docs`, `style`, `refactor`, `test`, `chore`.
3. Imperative mood: Use "add feature" instead of "added feature".
4. Body: Explain *why* the change was made, not just *what* changed.
