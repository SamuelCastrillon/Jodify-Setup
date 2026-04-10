# Project Rules - Jodify-Setup

Este archivo documenta las convenciones y reglas específicas del proyecto para SDD.

## Stack del Proyecto

- **Lenguaje**: Go 1.23
- **CLI Framework**: spf13/cobra v1.10.2
- **Testing**: Go testing (`*_test.go`)
- **Distribución**: GoReleaser
- **Package Manager**: Scoop (Windows), Brew (macOS)

## Estructura de Código

```
cmd/jodify-setup/          # Entry point
internal/                  # Código privado (no exportado)
├── cli/                   # Comandos Cobra
├── installer/             # Lógica de instalación
├── platform/              # Abstracción Windows/macOS
└── config/                # Gestor de configuración
pkg/                       # Código exportable (librerías)
```

## Convenciones de Go

| Aspecto | Convención |
|---------|------------|
| **Tests** | Archivos `*_test.go` junto al código que testea |
| **Nombres** | `camelCase` para variables y funciones, `PascalCase` para exportados |
| **Errores** | Usar `errors.Wrap()` o `fmt.Errorf()` con contexto |
| **Contexto** | Pasar `context.Context` como primer argumento |
| **Paquetes** | Paquetes pequeños y cohesivos |

## Convenciones de Git

| Aspecto | Convención |
|---------|------------|
| **Commits** | Conventional commits (`feat:`, `fix:`, `docs:`, `chore:`, `refactor:`) |
| **Ramas** | `feature/<nombre>`, `fix/<nombre>`, `docs/<nombre>` |
| **Tags** | versioning semántico (`v1.0.0`) |

## Documentación

| Archivo | Propósito |
|---------|-----------|
| `PRD.md` | Requisitos completos del producto (fuente de verdad) |
| `README.md` | Guía de uso para usuarios finales |
| `AGENTS.md` | Convenciones para agentes de IA |
| `docs/*.md` | Documentos técnicos de diseño |

## SDD Workflow

### Modo de Persistencia

- **Modo**: hybrid (Engram + openspec)
- **Config**: `openspec/config.yaml`

### Fases SDD

| Fase | Artefacto | Descripción |
|-----|----------|-------------|
| Explore | `sdd/{change}/explore` | Investigación de ideas |
| Propose | `sdd/{change}/proposal` | Propuesta de cambio |
| Spec | `sdd/{change}/spec` | Especificaciones (Given/When/Then) |
| Design | `sdd/{change}/design` | Diseño técnico |
| Tasks | `sdd/{change}/tasks` | Lista de tareas |
| Apply | `sdd/{change}/apply` | Implementación |
| Verify | `sdd/{change}/verify` | Verificación contra specs |
| Archive | `sdd/{change}/archive` | Cambio completado |

### Reglas de Tasks

- Usar numeración jerárquica (1.1, 1.2, 2.1, etc.)
- Agrupar por fase (infraestructura, implementación, testing)
- Mantener tareas pequeñas para completar en una sesión

## Skills Disponibles

### Locales (en .agents/skills/)

| Skill | Descripción |
|-------|-------------|
| `golang-testing` | Patrones de testing en Go |

### Registry

- Ubicación: `.atl/skill-registry.md`
- Se actualiza automáticamente con `sdd-init`

## Commands Útiles

```bash
# Compilar
go build -o jodify-setup.exe ./cmd/jodify-setup

# Ejecutar tests
go test ./...

# Formatear código
go fmt ./...

# Lint
go vet ./...

# Release con GoReleaser
goreleaser release --clean
```

## Contacto

- **Owner**: SamuelCastrillon
- **Repo**: Jodify-Setup