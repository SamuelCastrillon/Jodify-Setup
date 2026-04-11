# Jodify-Setup

Herramienta de configuración automatizada para entornos de desarrollo basados en Neovim. Diseñado para Windows y macOS, con soporte para Scoop (Windows) y gestión inteligente de dependencias.

## ¿Qué es Jodify?

Jodify es un entorno de desarrollo en terminal basado en Neovim, preconfigurado y optimizado para velocidad. Incluye:

- **Configuración completa de Neovim** en Lua
- **Plugins selecionados** para LSP, formateo, navegación y Git
- **Tema personalizado** con estética violeta oscura (Catppuccin Mocha modificado)
- **Integración con herramientas** como Scoop, Brew, lazygit, fzf, y más

## Requisitos

### Windows
- Windows 10/11 con PowerShell 7
- [Scoop](https://scoop.sh) (se instala automáticamente si no está presente)

### macOS
- macOS 10.15+
- [Homebrew](https://brew.sh) (se instala automáticamente si no está presente)

## Instalación

### Opción 1: Descargar binario

```powershell
# Windows
irm get.scoop.sh | iex
scoop install https://github.com/SamuelCastrillon/Jodify-Setup/releases/latest/download/jodify-setup.ps1
```

### Opción 2: Compilar desde código

```bash
git clone https://github.com/SamuelCastrillon/Jodify-Setup.git
cd Jodify-Setup
go build -o jodify-setup.exe ./cmd/jodify-setup
```

## Uso

```bash
# Instalación interactiva
./jodify-setup.exe

# Ver versión
./jodify-setup.exe version

# Desinstalar
./jodify-setup.exe uninstall
```

## Comandos

| Comando | Descripción |
|---------|-------------|
| `jodify-setup` | Iniciar instalación interactiva |
| `jodify-setup install` | Instalar Jodify |
| `jodify-setup install --skip-deps` | Instalar sin dependencias del sistema |
| `jodify-setup dev sync` | Sincronizar config local para desarrollo |
| `jodify-setup dev open` | Abrir Neovim con config de desarrollo |
| `jodify-setup dev clean` | Limpiar config de desarrollo |
| `jodify-setup uninstall` | Desinstalar completamente |
| `jodify-setup version` | Mostrar versión |
| `jodify-setup update` | Actualizar Jodify |

## Jodify-Dev

Una vez instalado, ejecutá:

```powershell
# Windows
Jodify-Dev

# Mostrar versión de Neovim
Jodify-Dev --check
```

Esto abre Neovim con configuración aislada (usando `NVIM_APPNAME=jodify`), sin afectar tu instalación existente de Neovim.

## Estructura del Proyecto

```
Jodify-Setup/
├── cmd/jodify-setup/     # Entry point
├── internal/
│   ├── cli/              # Comandos Cobra
│   ├── installer/        # Instalación y configuración
│   ├── platform/        # Abstracción Windows/macOS
│   ├── config/          # Gestor de configuración
│   └── dependencies/    # Auto-install de dependencias
├── pkg/                  # Paquetes exportables
├── config/               # Configuración Lua de Neovim
└── openspec/             # SDD workflow
```

## Documentación

- [PRD.md](./PRD.md) - Requisitos completos del producto
- [openspec/](./openspec/) - SDD workflow y artefactos
- [AGENTS.md](./AGENTS.md) - Convenciones para agentes de IA

## Licencia

MIT