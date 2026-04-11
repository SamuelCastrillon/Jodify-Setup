# Proposal: Auto-install de Dependencias del Sistema

## Intent

El CLI actual de Jodify-Setup solo verifica Neovim y Git como prerrequisitos. Según el PRD, también debe instalar automáticamente herramientas del sistema (gcc, ripgrep, fd, fzf, zoxide, lazygit) que son necesarias para que los plugins de Neovim funcionen correctamente (Treesitter, Telescope, etc.).

## Scope

### In Scope
- Detección de herramientas faltantes en el sistema (Windows + macOS)
- Auto-install de dependencias via Scoop (Windows) / Homebrew (macOS)
- Integración con el comando `jodify-setup install`
- Flags CLI: `--skip-deps` para omitir instalación automática

### Out of Scope
- TUI interactiva con Huh (reemplazada por flujo automático con output)
- Actualización de dependencias existentes (solo install, no update)
- Configuración de PATH post-install

## Approach

**Opción seleccionada: Silencioso con fallback**

1. **Detección**: Extender `CheckPrerequisites()` para detectar herramientas faltantes
2. **Install automático**: Si Scoop/Brew está disponible, instalar las herramientas faltantes
3. **Output**: Mostrar qué se instala en stdout
4. **Fallback**: Si falla, mostrar mensaje claro con comando manual

**Arquitectura**:
- Extender interfaz `Platform` con nuevo método `InstallDependencies()`
- Crear nuevo módulo `internal/dependencies/` para lógica de instalación
- No usar TUI con Huh para reducir complejidad

## Affected Areas

| Area | Impact | Description |
|------|--------|-------------|
| `internal/platform/platform.go` | Modified | Agregar método `InstallDependencies()` a interfaz |
| `internal/platform/windows.go` | Modified | Implementar detección e install via Scoop |
| `internal/platform/darwin.go` | Modified | Implementar detección e install via Homebrew |
| `internal/installer/installer.go` | Modified | Llamar a InstallDependencies antes de descargar config |
| `internal/cli/install.go` | Modified | Agregar flag `--skip-deps` |

## Risks

| Risk | Likelihood | Mitigation |
|------|------------|------------|
| Scoop no está en PATH post-install | Medium | Agregar shims al PATH explícitamente |
| Fallo en redes restringidas | Medium | Mostrar comando manual como fallback |
| Usuario ya tiene algunas herramientas | Low | Detectar y solo instalar lo que falta |

## Rollback Plan

- Si la instalación falla, el CLI muestra error claro con comando manual para instalar cada herramienta
- El usuario puede ejecutar `jodify-setup install --skip-deps` para saltar el paso
- No se modifica la config existente, solo se agregan herramientas al sistema

## Dependencies

- Neovim y Git deben estar instalados primero (ya verificado por CheckPrerequisites)
- Scoop (Windows) o Homebrew (macOS) debe estar disponible

## Success Criteria

- [ ] `jodify-setup install` detecta herramientas faltantes
- [ ] Auto-install gcc, ripgrep, fd, fzf, zoxide, lazygit en Windows
- [ ] Auto-install equivalentes en macOS
- [ ] Flag `--skip-deps` omite la instalación
- [ ] Si falla, muestra mensaje claro con comandos manuales
- [ ] Tests cubriendo la funcionalidad