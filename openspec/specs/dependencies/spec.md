# Specification: Auto-install de Dependencias del Sistema

## Purpose

Este documento especifica los requisitos para la detección e instalación automática de herramientas del sistema requeridas por los plugins de Neovim en Jodify-Setup.

## Requirements

### Requirement: Detección de Herramientas Faltantes

El sistema DEBE detectar las herramientas del sistema que faltan antes de la instalación de la configuración de Neovim.

#### Scenario: Herramientas Present en Windows

- GIVEN el usuario ejecuta `jodify-setup install` en Windows
- WHEN Scoop está disponible y las herramientas (gcc, ripgrep, fd, fzf, zoxide, lazygit) están instaladas
- THEN el sistema DEBE continuar sin instalar nada
- AND mostrar un mensaje indicando que las herramientas ya están presentes

#### Scenario: Herramientas Faltantes en Windows

- GIVEN el usuario ejecuta `jodify-setup install` en Windows
- WHEN Scoop está disponible pero algunas herramientas faltan
- THEN el sistema DEBE instalar automáticamente las herramientas faltantes

### Requirement: Instalación Automática de Dependencias

El sistema DEBE instalar automáticamente las herramientas faltantes usando el gestor de paquetes del sistema.

#### Scenario: Instalación Exitosa en Windows

- GIVEN Scoop está disponible y las herramientas faltan
- WHEN el sistema ejecuta la instalación automática
- THEN todas las herramientas DEBEN ser instaladas exitosamente
- AND el sistema DEBE mostrar las herramientas instaladas en stdout

#### Scenario: Fallo en Instalación

- GIVEN Scoop está disponible pero la instalación falla
- WHEN la herramienta no se puede instalar (red, permisos, etc.)
- THEN el sistema DEBE mostrar un mensaje claro con los comandos manuales para cada herramienta que falló

### Requirement: Integración con Comando Install

El sistema DEBE integrar la detección e instalación de dependencias con el comando `jodify-setup install`.

#### Scenario: Install con Flags

- GIVEN el usuario ejecuta `jodify-setup install`
- WHEN el usuario proporciona `--skip-deps`
- THEN el sistema DEBE omitir la detección e instalación de dependencias
- AND continuar con la instalación de la configuración

- GIVEN el usuario ejecuta `jodify-setup install`
- WHEN el usuario NO proporciona `--skip-deps`
- THEN el sistema DEBE ejecutar la detección e instalación de dependencias

### Requirement: Soporte Multi-Plataforma

El sistema DEBE soportar la instalación de dependencias en Windows (Scoop) y macOS (Homebrew).

#### Scenario: Windows con Scoop

- GIVEN el usuario ejecuta `jodify-setup install` en Windows
- WHEN Scoop está disponible
- THEN el sistema DEBE usar Scoop para instalar las herramientas

#### Scenario: macOS con Homebrew

- GIVEN el usuario ejecuta `jodify-setup install` en macOS
- WHEN Homebrew está disponible
- THEN el sistema DEBE usar Homebrew para instalar las herramientas equivalentes

#### Scenario: Ningún Gestor de Paquetes Disponible

- GIVEN el usuario ejecuta `jodify-setup install`
- WHEN ni Scoop (Windows) ni Homebrew (macOS) están disponibles
- THEN el sistema DEBE mostrar un mensaje indicando que no se pueden instalar dependencias automáticamente
- AND proporcionar instrucciones manuales

## Additional Requirements

### Requisito: Lista de Herramientas a Instalar

El sistema DEBE intentar instalar las siguientes herramientas:

| Herramienta | Propósito | Windows (Scoop) | macOS (Brew) |
|-------------|-----------|-----------------|--------------|
| gcc | Compilador C para Treesitter | `scoop install gcc` | `brew install gcc` |
| ripgrep | Búsqueda para Telescope | `scoop install ripgrep` | `brew install ripgrep` |
| fd | Alternativa a find | `scoop install fd` | `brew install fd` |
| fzf | Buscador difuso | `scoop install fzf` | `brew install fzf` |
| zoxide | Smart cd | `scoop install zoxide` | `brew install zoxide` |
| lazygit | TUI para Git | `scoop install lazygit` | `brew install lazygit` |

### Requisito: Verificación Post-Instalación

El sistema DEBE verificar que cada herramienta se instaló correctamente después de la instalación.

- GIVEN una herramienta fue instalada
- WHEN el sistema verifica la instalación
- THEN el sistema DEBE confirmar que la herramienta está disponible en el PATH