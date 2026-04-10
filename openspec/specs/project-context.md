# Jodify-Setup - Project Context

## SDD Inicializado

**Fecha**: 2026-04-10  
**Proyecto**: Jodify-Setup  
**Stack**: Go 1.23 + Cobra + Lua  
**Modo**: hybrid (engram + openspec)

---

## Stack Detectado

- **Lenguaje**: Go 1.23
- **CLI Framework**: spf13/cobra v1.10.2
- **Configuración**: Lua (Neovim dotfiles)
- **Distribución**: GoReleaser (Windows + macOS)
- **Package Manager**: Scoop (Windows), Brew (macOS)

## Estructura del Proyecto

```
Jodify-Setup/
├── cmd/jodify-setup/main.go       # Entry point
├── internal/
│   ├── cli/                       # Comandos Cobra (root, install, uninstall, update, version)
│   ├── installer/                 # Lógica de instalación
│   ├── platform/                  # Abstracción Windows/macOS
│   │   ├── detector.go           # Detección de plataforma
│   │   ├── windows.go            # Lógica específica de Windows
│   │   ├── darwin.go            # Lógica específica de macOS
│   │   └── platform.go          # Interfaces comunes
│   └── config/                    # Gestor de configuración
├── pkg/version/                   # Paquete de versiones
├── config/                        # Configuración Lua (se empaqueta como jodify-config.zip)
├── docs/                          # Documentación técnica
│   ├── auto-install-scoop-design.md
│   └── judgment-prd.md
└── openspec/                      # SDD (this directory)
```

## Convenciones

- **Testing**: Tests en archivos `*_test.go` junto al código
- **Patrón de arquitectura**: Clean Architecture en Go (cmd → internal → pkg)
- **Documentación**: PRD.md en raíz, diseños técnicos en docs/
- **Commits**: Conventional commits (feat:, fix:, docs:, etc.)
- **Ramas**: feature/<nombre>, fix/<nombre>, docs/<nombre>

## Documentos Existentes

- **PRD.md** - Requisitos completos del producto
- **README.md** - Guía de uso para usuarios finales
- **AGENTS.md** - Convenciones para agentes de IA
- docs/auto-install-scoop-design.md - Diseño técnico de Scoop auto-install
- docs/judgment-prd.md - Documentación adicional de PRD
- .atl/skill-registry.md - Registro de skills

## Skill Registry

Ubicado en `.atl/skill-registry.md`. Contiene:
- User skills (go-testing, skill-creator, etc.)
- SDD workflow skills (sdd-init, sdd-explore, etc.)
- Project conventions