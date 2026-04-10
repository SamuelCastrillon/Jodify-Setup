# Jodify-Setup - Project Context

##SDD Initialized

**Fecha**: 2026-04-10  
**Proyecto**: Jodify-Setup  
**Stack**: Go 1.23 + Cobra + Lua  
**Modo**: hybrid (engram + openspec)

---

## Stack del Proyecto

- **Lenguaje**: Go 1.23
- **CLI Framework**: spf13/cobra v1.10.2
- **Testing**: Go testing (`*_test.go`)
- **Distribución**: GoReleaser
- **Package Manager**: Scoop (Windows), Brew (macOS)
- **Configuración Neovim**: Lua

## Estructura del Proyecto

```
Jodify-Setup/
├── .atl/                    # Agent tools
│   └── skill-registry.md     # Skill registry
├── .agents/                 # Agent skills (locales)
│   └── skills/
│       └── golang-testing/
├── openspec/                # SDD workflow
│   ├── config.yaml          # SDD config
│   └── specs/
├── cmd/jodify-setup/        # Entry point
├── internal/               # Código privado
│   ├── cli/              # Comandos Cobra
│   ├── installer/        # Instalación
│   ├── platform/         # Windows/macOS
│   └── config/           # Configuración
├── pkg/                   # Código exportable
├── config/                 # Config Lua Neovim
├── docs/                   # Documentación técnica
├── PRD.md                 # Requisitos completos
├── README.md              # Guía de usuario
└── AGENTS.md              # Convenciones agentes
```

## Convenciones

### Go
| Aspecto | Convención |
|---------|------------|
| Tests | Archivos `*_test.go` junto al código |
| Nombres | `camelCase` (variables), `PascalCase` (exportados) |
| Errores | `errors.Wrap()` o `fmt.Errorf()` con contexto |
| Contexto | `context.Context` como primer argumento |

### Git
| Aspecto | Convención |
|---------|------------|
| Commits | Conventional (`feat:`, `fix:`, `docs:`, `chore:`) |
| Ramas | `feature/<nombre>`, `fix/<nombre>`, `docs/<nombre>` |
| Tags | semantic versioning (`v1.0.0`) |

### Documentación
| Archivo | Propósito |
|---------|-----------|
| PRD.md | Requisitos (fuente de verdad) |
| README.md | Guía usuarios finales |
| AGENTS.md | Convenciones agentes IA |
| docs/*.md | Documentos técnicos |

## SDD Workflow

### Modo
- **Persistencia**: hybrid (engram + openspec)
- **Config**: `openspec/config.yaml`

### Fases
| Fase | Descripción |
|------|-----------|
| Explore | Investigación de ideas |
| Propose | Propuesta de cambio |
| Spec | Especificaciones |
| Design | Diseño técnico |
| Tasks | Lista de tareas |
| Apply | Implementación |
| Verify | Verificación |
| Archive | Cambio completado |

### Reglas de Tasks
- Numeración jerárquica (1.1, 1.2, 2.1)
- Agrupar por fase
- Tareas pequeñaspara una sesión

## Skills

### Locales (proyecto)
| Skill | Descripción |
|-------|-------------|
| golang-testing | Patrones testing Go |

### Registry
- Ubicación: `.atl/skill-registry.md`

## Comandos Útiles

```bash
# Compilar
go build -o jodify-setup.exe ./cmd/jodify-setup

# Tests
go test ./...

# Formatear
go fmt ./...

# Lint
go vet ./...

# Release
goreleaser release --clean
```

## Contacto

- **Owner**: SamuelCastrillon
- **Repo**: Jodify-Setup