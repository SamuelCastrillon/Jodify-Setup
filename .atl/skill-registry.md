# Skill Registry

**Orchestrator use only.** Read this registry once per session to resolve skill paths, then pass pre-resolved paths directly to each sub-agent's launch prompt. Sub-agents receive the path and load the skill directly — they do NOT read this registry.

## Local Skills (in .agents/skills/)

These skills are installed locally in the project and versioned in git.

| Trigger | Skill | Path |
|---------|-------|------|
| Writing Go tests, using teatest, or adding test coverage | go-testing | .agents/skills/golang-testing/SKILL.md |

## Global Skills (Fallback - may not exist on all machines)

| Trigger | Skill | Path |
|---------|-------|------|
| Creating new AI skills, adding agent instructions, or documenting patterns for AI | skill-creator | global → C:\Users\Admin\.config\opencode\skills\skill-creator\SKILL.md |
| When user says "update skills", "skill registry", "actualizar skills", "update registry", or after installing/removing skills | skill-registry | global → C:\Users\Admin\.config\opencode\skills\skill-registry\SKILL.md |
| Creating GitHub issues, reporting bugs, or requesting features | issue-creation | global → C:\Users\Admin\.config\opencode\skills\issue-creation\SKILL.md |
| Creating pull requests, opening PRs, or preparing changes for review | branch-pr | global → C:\Users\Admin\.config\opencode\skills\branch-pr\SKILL.md |
| User says "judgment day", "review adversarial", "dual review", "doble review", "juzgar", "que lo juzguen" | judgment-day | global → C:\Users\Admin\.config\opencode\skills\judgment-day\SKILL.md |

## SDD Workflow Skills (Global)

These skills are used internally by the SDD orchestrator and should not be called directly by users:

| Phase | Skill | Path |
|-------|-------|------|
| Init | sdd-init | global → C:\Users\Admin\.config\opencode\skills\sdd-init\SKILL.md |
| Explore | sdd-explore | global → C:\Users\Admin\.config\opencode\skills\sdd-explore\SKILL.md |
| Propose | sdd-propose | global → C:\Users\Admin\.config\opencode\skills\sdd-propose\SKILL.md |
| Spec | sdd-spec | global → C:\Users\Admin\.config\opencode\skills\sdd-spec\SKILL.md |
| Design | sdd-design | global → C:\Users\Admin\.config\opencode\skills\sdd-design\SKILL.md |
| Tasks | sdd-tasks | global → C:\Users\Admin\.config\opencode\skills\sdd-tasks\SKILL.md |
| Apply | sdd-apply | global → C:\Users\Admin\.config\opencode\skills\sdd-apply\SKILL.md |
| Verify | sdd-verify | global → C:\Users\Admin\.config\opencode\skills\sdd-verify\SKILL.md |
| Archive | sdd-archive | global → C:\Users\Admin\.config\opencode\skills\sdd-archive\SKILL.md |
| Onboard | sdd-onboard | global → C:\Users\Admin\.config\opencode\skills\sdd-onboard\SKILL.md |

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