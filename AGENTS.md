# AGENTS.md - Guía para Agentes de IA

Este documento establece las convenciones y contexto que los agentes de IA deben seguir al trabajar en Jodify-Setup.

## Stack del Proyecto

- **Lenguaje**: Go 1.23
- **CLI Framework**: spf13/cobra v1.10.2
- **Testing**: Go testing (`*_test.go`)
- **Distribución**: GoReleaser

## Estructura de Código

### Organización de Paquetes

```
cmd/jodify-setup/          # Entry point
internal/                  # Código privado (no exportado)
├── cli/                   # Comandos Cobra
│   ├── root.go            # Comando raíz
│   ├── install.go         # Instalación
│   ├── uninstall.go       # Desinstalación
│   ├── update.go          # Actualización
│   └── version.go         # Versión
├── installer/             # Lógica de instalación
├── platform/              # Abstracción Windows/macOS
│   ├── detector.go       # Detección de plataforma
│   ├── windows.go        # Lógica específica de Windows
│   ├── darwin.go         # Lógica específica de macOS
│   └── platform.go        # Interfaces comunes
├── config/                # Gestor de configuración
└── config/                # Configuración CLI
pkg/                       # Código exportable (librerías)
```

### Convenciones de Go

- **Tests**: Archivos命名为 `*_test.go` junto al código que testea
- **Nombres**: `camelCase` para variables y funciones, `PascalCase` para exportados
- **Errores**: Usar `errors.Wrap()` o `fmt.Errorf()` con contexto
- **Contexto**: Pasar `context.Context` como primer argumento

## Convenciones de Git

- **Commits**: Conventional commits (`feat:`, `fix:`, `docs:`, etc.)
- **Ramas**: `feature/<nombre>`, `fix/<nombre>`, `docs/<nombre>`
- **Tags**: versioning semántico (`v1.0.0`)

## Documentación

- **PRD.md**: Requisitos del producto (fuente de verdad)
- **docs/**: Documentos técnicos de diseño
- **README.md**: Guía de uso para usuarios finales

## Comandos Útiles

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

## SDD - Spec-Driven Development

Este proyecto usa SDD para gestionar cambios sustanciales. Los comandos disponibles:

- `/sdd-init` - Inicializar SDD
- `/sdd-explore <tema>` - Investigar ideas
- `/sdd-new <nombre>` - Crear propuesta de cambio
- `/sdd-spec` - Escribir especificaciones
- `/sdd-design` - Crear diseño técnico
- `/sdd-tasks` - Desglosar tareas
- `/sdd-apply` - Implementar
- `/sdd-verify` - Verificar contra specs
- `/sdd-archive` - Archivar cambio completado

## Recursos Adicionales

- [PRD.md](./PRD.md) - Requisitos completos
- [.atl/skill-registry.md](./.atl/skill-registry.md) - Registro de skills disponibles