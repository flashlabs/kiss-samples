# Minimal Skill Router

The router is intentionally small. It does not contain domain knowledge. It only decides which skill should be active.

## Rules

Use `good-git-commit-skill.md` when the user asks to:

- create a commit
- commit current changes
- commit a specific file or set of files
- prepare a local commit after implementation

Do not use `good-git-commit-skill.md` when the user asks to:

- review code
- explain code
- implement a feature without committing
- create a pull request
- rewrite Git history

## Routing Checklist

Before loading the skill, answer:

- Is the requested output a local Git commit?
- Is the working tree relevant to the task?
- Is there a risk of unrelated user changes?

If the answer to the first question is "no", use a different skill.
