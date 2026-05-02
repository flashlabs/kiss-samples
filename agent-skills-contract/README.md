# Agent Skills Contract

Companion examples for the article "A Good Agent Skill Is a Contract, Not a Prompt".

This sample compares vague prompt-like instructions with a skill that has a clear trigger, boundaries, verification rules, and output contract.

For details see: https://blog.skopow.ski/a-good-agent-skill-is-a-contract-not-a-prompt

## Files

- `bad-skill.md` - a weak skill that sounds useful but leaves too much room for interpretation.
- `good-git-commit-skill.md` - a more explicit Git commit skill with a focused contract.
- `router.md` - a minimal router that decides when the Git commit skill should be active.
