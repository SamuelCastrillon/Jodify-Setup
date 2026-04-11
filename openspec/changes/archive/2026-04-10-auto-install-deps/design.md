# Design: Auto-install de Dependencias del Sistema

## Technical Approach

Agregar detección e instalación automática de herramientas del sistema (gcc, ripgrep, fd, fzf, zoxide, lazygit) como parte del flujo de `jodify-setup install`.

## Architecture Decisions

### Decision: Dónde implementar la lógica de dependencias

**Choice**: Crear nuevo módulo `internal/dependencies/` independiente
**Alternatives considered**: 
- Extender `platform.Platform` con nuevo método
- Agregar lógica en `installer.Install()`
**Rationale**: Separación de responsabilidades. El módulo dependencies maneja solo la detección e instalación de herramientas del sistema, no la config de Neovim.

### Decision: Cómo detectar herramientas disponibles

**Choice**: Verificar cada herramienta ejecutando el comando con `--version` o `which`
**Alternatives considered**:
- Verificar solo si el gestor de paquetes existe
- Mantener lista de herramientas instaladas en cache
**Rationale**: Más confiable. Verificar cada herramienta individualmente asegura que realmente está disponible en PATH.

### Decision: Cuándo instalar las dependencias

**Choice**: Antes de descargar la config de Neovim
**Alternatives considered**:
- Después de descargar la config
- De forma diferida (lazy load)
**Rationale**: Treesitter necesita gcc para compilar parsers al primer inicio de Neovim. Es mejor tenerlas instaladas antes.

## Data Flow

```
jodify-setup install
    │
    ▼
installer.Install()
    │
    ├── CheckPrerequisites() [Neovim + Git]
    │
    ├── dependencies.CheckAndInstall() [NEW]
    │       │
    │       ├── DetectPackageManager() [Scoop/Brew]
    │       │
    │       ├── For each tool:
    │       │   └── IsInstalled() → Install() if missing
    │       │
    │       └── VerifyAll() → return errors if any
    │
    ├── Download() [config zip]
    │
    └── Extract() [to config dir]
```

## File Changes

| File | Action | Description |
|------|--------|-------------|
| `internal/dependencies/manager.go` | Create | Nueva interfaz y struct para gestión de dependencias |
| `internal/dependencies/windows.go` | Create | Implementación Windows (Scoop) |
| `internal/dependencies/darwin.go` | Create | Implementación macOS (Homebrew) |
| `internal/dependencies/tool.go` | Create | Definición de herramientas y comandos |
| `internal/platform/platform.go` | Modify | Agregar método `GetPackageManager()` a interfaz |
| `internal/installer/installer.go` | Modify | Llamar dependencies en Install() |
| `internal/cli/install.go` | Modify | Agregar flag `--skip-deps` |

## Interfaces / Contracts

```go
// DependenciesManager define interface for dependency management
type DependenciesManager interface {
    // CheckAndInstall verifica e instala las dependencias faltantes
    CheckAndInstall(ctx context.Context, opts InstallOptions) (*InstallResult, error)
    
    // IsInstalled verifica si una herramienta está instalada
    IsInstalled(tool string) (bool, error)
}

// InstallResult contiene el resultado de la instalación
type InstallResult struct {
    Installed []string
    Failed    map[string]error
    Skipped   []string
}

// Tool representa una herramienta del sistema
type Tool struct {
    Name        string
    Description string
    WindowsCmd  string  // comando Scoop
    MacCmd      string  // comando Brew
}
```

## Testing Strategy

| Layer | What to Test | Approach |
|-------|-------------|----------|
| Unit | IsInstalled(), parse version output | Mock exec.Command |
| Unit | Tool struct definitions | Assert fields |
| Integration | Install flow end-to-end | Skip if no Scoop/Brew |
| Integration | Fallback cuando no hay gestor | Mock platform detection |

## Migration / Rollout

No migration required. Este es un feature nuevo que no afecta datos existentes.

## Open Questions

- [ ] ¿Cómo manejar el caso donde el usuario tiene algunas herramientas pero no todas?
- [ ] ¿Debe el CLI fallar si alguna dependencia no se puede instalar, o solo warn?