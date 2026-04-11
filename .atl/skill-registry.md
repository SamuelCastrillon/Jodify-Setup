# Skill Registry

**Orchestrator use only.** Read this registry once per session to resolve skill paths, then pass pre-resolved paths directly to each sub-agent's launch prompt. Sub-agents receive the path and load the skill directly — they do NOT read this registry.

## Local Skills (in .agents/skills/)

These skills are installed locally in the project and versioned in git.

| Trigger | Skill | Path |
|---------|-------|------|
| Writing Go tests, using teatest, or adding test coverage | go-testing | .agents/skills/golang-testing/SKILL.md |
| Neovim config validation, plugins, or lazy.nvim setup | neovim | .agents/skills/neovim/SKILL.md |

## Neovim Validation Strategy

This project uses a **dual-layer validation** for Neovim configs:

### Layer 1: General Validation (neovim skill)
- Validates syntax, best practices, plugin patterns
- References: `neovim/references/migration-0.11.md`

### Layer 2: Jodify-specific (future skill)
- Bootstrap compatibility (vim.uv vs vim.loop)
- Jodify plugins from PRD
- NVIM_APPNAME isolation

## Global Skills (Fallback - may not exist on all machines)

| Trigger | Skill | Path |
|---------|-------|------|
| Creating new AI skills, adding agent instructions, or documenting patterns for AI | skill-creator | global → skill-creator |
| When user says "update skills", "skill registry", "actualizar skills", "update registry", or after installing/removing skills | skill-registry | global → skill-registry |
| Creating GitHub issues, reporting bugs, or requesting features | issue-creation | global → issue-creation |
| Creating pull requests, opening PRs, or preparing changes for review | branch-pr | global → branch-pr |
| User says "judgment day", "review adversarial", "dual review", "doble review", "juzgar", "que lo juzguen" | judgment-day | global → judgment-day |

## SDD Workflow Skills (Global)

These skills are used internally by the SDD orchestrator and should not be called directly by users:

| Phase | Skill | Path |
|-------|-------|------|
| Init | sdd-init | global |
| Explore | sdd-explore | global |
| Propose | sdd-propose | global |
| Spec | sdd-spec | global |
| Design | sdd-design | global |
| Tasks | sdd-tasks | global |
| Apply | sdd-apply | global |
| Verify | sdd-verify | global |
| Archive | sdd-archive | global |
| Onboard | sdd-onboard | global |

## Project Conventions

| File | Path | Notes |
|------|------|-------|
| PRD | PRD.md | Product Requirements Document — main project specification |
| README | README.md | User-facing documentation |
| Agents | AGENTS.md | AI agent conventions |
| Openspec config | openspec/config.yaml | SDD configuration |

---

*Jodify-Setup Skill Registry*

**Note:** Local skills (.agents/skills/) are versioned in git. Global skills are user-level installs.