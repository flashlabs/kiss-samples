# Git Commit Skill

Use this skill when the user asks you to create a Git commit.

## Scope

This skill owns only the commit workflow.

It does not decide code architecture, test strategy, release process, or pull request content.

## Before Committing

- Run `git status --short`.
- Inspect changed files enough to understand the commit scope.
- Stage only files related to the requested change.
- Do not stage unrelated user changes.
- If unrelated changes are present, leave them unstaged and mention them.

## Commit Rules

- Create one commit for the requested work.
- Use the repository's local commit message convention if one exists.
- Use a concise imperative commit subject.
- Do not amend, rebase, reset, or rewrite history unless the user explicitly asks.
- Do not run destructive Git commands unless the user explicitly asks and approval rules allow it.

## Verification

- Run relevant tests when feasible.
- If tests are skipped, explain why.
- Do not hide failed verification.

## Final Response

Include:

- the commit hash
- a short summary of what was committed
- any tests or checks that were run
- any remaining uncommitted changes
