# Tasks: Auto-install de Dependencias del Sistema

## Phase 1: Foundation ( Interfaces y Tipos )

- [x] 1.1 Crear `internal/dependencies/tool.go` con struct Tool y lista de herramientas
- [x] 1.2 Crear `internal/dependencies/manager.go` con interfaz DependenciesManager
- [x] 1.3 Crear `internal/dependencies/result.go` con struct InstallResult

## Phase 2: Implementación Platform (Windows + macOS)

- [x] 2.1 Crear `internal/dependencies/windows.go` con detección Scoop
- [x] 2.2 Implementar `IsInstalled()` en windows.go para cada herramienta
- [x] 2.3 Implementar `Install()` en windows.go usando Scoop
- [x] 2.4 Crear `internal/dependencies/darwin.go` con detección Homebrew
- [x] 2.5 Implementar `IsInstalled()` en darwin.go para cada herramienta
- [x] 2.6 Implementar `Install()` en darwin.go usando Homebrew
- [x] 2.7 Crear función `DetectPackageManager()` para detectar Scoop/Brew

## Phase 3: Integración con Installer

- [x] 3.1 Modificar `internal/installer/installer.go` para llamar dependencies
- [x] 3.2 Crear método `CheckAndInstall()` que ejecute la lógica de deps
- [x] 3.3 Agregar campo `SkipDeps` a InstallOptions
- [x] 3.4 Integrar verificación de deps en flujo de `Install()`

## Phase 4: CLI (Flags)

- [x] 4.1 Modificar `internal/cli/install.go` para agregar flag `--skip-deps`
- [x] 4.2 Pasar flag al Installer
- [x] 4.3 Agregar output informativo (stdout) para tools instaladas/fallidas

## Phase 5: Testing

- [ ] 5.1 Crear `internal/dependencies/manager_test.go` con tests unitarios
- [ ] 5.2 Test: `IsInstalled()` retorna true cuando tool existe
- [ ] 5.3 Test: `IsInstalled()` retorna false cuando tool no existe
- [ ] 5.4 Test: `CheckAndInstall()` instala tools faltantes
- [ ] 5.5 Test: `CheckAndInstall()` omite tools ya instaladas
- [ ] 5.6 Test: Fallback cuando package manager no disponible
- [ ] 5.7 Test: Verificar output con InstallResult

## Phase 6: Documentación (Opcional)

- [ ] 6.1 Actualizar README.md con nuevas dependencias
- [ ] 6.2 Agregar ejemplos de uso del flag --skip-deps