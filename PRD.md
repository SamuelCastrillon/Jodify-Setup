# Documento de Requisitos del Producto (PRD): Jodify

**Nombre del Producto:** Jodify (Jodify-Setup & Jodify-Dev)
**Plataformas:** Windows (PowerShell 7) y macOS
**Stack Principal:** Go (Backend CLI/TUI), Lua (Configuración de Neovim), Git

---

## 1. Visión del Producto

Proporcionar a los desarrolladores una herramienta automatizada y ultrarrápida para desplegar un entorno de desarrollo en terminal basado en Neovim. Este entorno ("Jodify") estará optimizado para la velocidad, preconfigurado con herramientas de IA (Agentes, MCP) y contará con una estética personalizada y fuertemente definida (modo oscuro con acentos morados).

---

## 2. Análisis Crítico y Corrección Arquitectónica

### El problema del planteamiento inicial

Si `Jodify-Setup` simplemente clona tus archivos en la carpeta por defecto de Neovim (`~/.config/nvim` o `~/AppData/Local/nvim`) y `Jodify-Dev` es solo un alias para la palabra `nvim`, **destruirás cualquier configuración previa de Neovim que el usuario ya tenga**. Si un usuario quiere probar Jodify pero no quiere perder su Neovim estándar, no podrá hacerlo.

### La Solución: NVIM_APPNAME

Neovim tiene una característica moderna llamada `NVIM_APPNAME`. Permite aislar completamente diferentes perfiles de Neovim en la misma máquina.

1. `Jodify-Setup` no clonará los archivos en la carpeta de `nvim`. Los clonará en una carpeta aislada llamada `jodify` (ej. `~/AppData/Local/jodify`).
2. `Jodify-Dev` no será un simple alias, será un **script ejecutable (wrapper)** generado por tu CLI en Go que injectará esta variable de entorno temporalmente:
   - Windows (Jodify-Dev.ps1): `$env:NVIM_APPNAME="jodify"; nvim $args`
   - macOS (Jodify-Dev): `NVIM_APPNAME=jodify nvim "$@"`

**Resultado:** El usuario puede tener su Neovim normal escribiendo `nvim`, y lanzar tu super-entorno escribiendo `Jodify-Dev`. Están 100% aislados.

---

## 3. Arquitectura del Sistema

### A. Un Solo Repo con Todo

El repositorio contiene todo el código fuente (Go CLI + Lua config). No hay repos separados.

### B. Distribución: Binario + Auto-Download

El usuario solo descarga **un archivo** (`jodify-setup.exe`). Al ejecutarse, el binario automáticamente:
1. Descarga el `config.zip` desde GitHub Releases
2. Extrae la config a `~/AppData/Local/jodify`
3. Genera el wrapper `Jodify-Dev.ps1`

### C. GoReleaser: Auto-Release

El release se construye automáticamente con **GoReleaser**:
- Compila el binario para Windows + macOS
- Empaqueta `config/` como `jodify-config.zip`
- Crea GitHub Release automáticamente

---

### 3.1 Flujo de Installación (Usuario Final)

```
1. Descarga: jodify-setup.exe         ← Solo ESTO
2. Ejecuta: .\jodify-setup.exe
3. → Se baixa config.zip (auto)
4. → Extrae a ~/AppData/Local/jodify
5. → Genera Jodify-Dev.ps1
6. Listo! → Jodify-Dev
```

---

### 3.2 Jodify-Setup (El Orquestador en Go)

- Un binario único y ligero construido en Go utilizando la librería `Huh` para la TUI.
- Responsabilidades:
  - Auditar el sistema operativo (detectar Windows/macOS)
  - Detectar e instalar Scoop (Windows) o Brew (macOS) automáticamente si no están presentes
  - Ejecutar gestores de paquetes para instalar dependencias base
  - **Descargar config.zip desde GitHub Releases**
  - Extraer config a la carpeta de perfil de Jodify (`NVIM_APPNAME`)
  - Instalar servidores MCP y generar las estructuras de carpetas para las Skills de los agentes
  - Generar e instalar el script ejecutable global `Jodify-Dev` en el PATH del sistema

### 3.3 Jodify Dotfiles (En config/)

- La configuración en Lua vive en la carpeta `config/` del repo
- Se empaqueta como `jodify-config.zip` en cada release
- Responsabilidades:
  - Definir la paleta de colores de la marca (Catppuccin Mocha modificado a base oscura/morada)
  - Configurar `alpha-nvim` para mostrar el logo ASCII de "JODIFY" y los accesos rápidos
  - Cargar plugins mediante `lazy.nvim`

---

## 4. Requisitos Funcionales

### RF1: Interfaz de Configuración

**Como usuario**, al ejecutar `Jodify-Setup`, debo ver una interfaz de terminal interactiva que me permita seleccionar qué componentes instalar (Core Neovim, Servidores MCP, Agent Skills).

### RF1.5: Auto-Install Scoop (Windows)

**Como usuario de Windows**, si Scoop no está instalado, debo ver una opción para instalarlo automáticamente con un solo clic, sin necesidad de abortar o hacerlo manualmente.

- El instalador muestra progreso en tiempo real en la TUI
- Si la instalación automática falla, muestra instrucciones manuales

### RF2: Gestión de Dependencias

**Como usuario de Windows**, la instalación debe utilizar Scoop silenciosamente en segundo plano para manejar las dependencias en C, motores de búsqueda y herramientas de terminal.

**Dependencias del Sistema Operativo a instalar:**

| Herramienta | Función Principal | Comando (Scoop) |
|------------|---------------|---------------|
| gcc | Compilador de C para Treesitter | `scoop install gcc` |
| ripgrep | Motor de búsqueda ultra rápido | `scoop install ripgrep` |
| fd | Alternativa rápida a find | `scoop install fd` |
| fzf | Buscador difuso | `scoop install fzf` |
| zoxide | "Smart cd" | `scoop install zoxide` |
| lazygit | TUI para Git | `scoop install lazygit` |

### RF3: Instalación No Interactiva

**Como usuario**, el proceso no debe requerir interacción manual una vez confirmadas las opciones; la TUI debe mostrar barras de progreso y un log limpio.

### RF4: Entorno Aislado

**Como usuario**, al ejecutar el comando `Jodify-Dev`, se debe abrir Neovim directamente en el entorno aislado de Jodify, mostrando el Dashboard de bienvenida.

### RF5: Terminal Integrada

**Como usuario**, dentro de `Jodify-Dev`, debo poder invocar una terminal flotante (`toggleterm`) que respete la configuración de mi shell principal (PowerShell 7 o Zsh).

### RF6: Actualización

**Como usuario**, al ejecutar `scoop update *`, se deben actualizar todas las herramientas incluyendo Neovim. El flag `--check` de `Jodify-Dev` debe mostrar la versión actual.

### RF7: Detección de Instalación Previa

**Como usuario**, al ejecutar `Jodify-Setup`, si ya tengo Jodify instalado (Existe `~/AppData/Local/jodify`), debo ver un mensaje indicando que ya está instalado y las opciones de actualizar o mantener la versión actual.

- Mostrar versión actual vs versión disponible
- Ofrecer: [Actualizar] [Mantener] [Salir]

### RF8: Desinstalación

**Como usuario**, debo poder desinstalar Jodify completamente con un flag `--uninstall`.

- Eliminar la carpeta `~/AppData/Local/jodify`
- Eliminar el wrapper `Jodify-Dev.ps1`
- Mostrar mensaje de confirmación

### RF9: Fallback Offline

**Como usuario**, si no tengo conexión a internet, debo poder usar un `config.zip` local que venga junto con el binario.

- El binario busca `config.zip` en el mismo directorio
- Si lo encuentra, lo usa en lugar de descargar
- Si no lo encuentra y no hay internet, muestra error claro

---

## 5. Plugins de Neovim

### Core, Interfaz y Navegación

| Plugin | Para qué sirve |
|--------|-------------|
| folke/lazy.nvim | Gestor de paquetes (carga asíncrona) |
| nvim-telescope/telescope.nvim | Buscador visual interactivo |
| nvim-neo-tree/neo-tree.nvim | Explorador de archivos |
| folke/which-key.nvim | Ayuda de memoria de atajos |
| akinsho/bufferline.nvim | Pestañas superiores |
| folke/tokyonight.nvim | Tema de colores |
| goolord/alpha-nvim | Pantalla de inicio (Dashboard) |

### Inteligencia de Código, LSP y Edición

| Plugin | Para qué sirve |
|--------|-------------|
| nvim-treesitter/nvim-treesitter | Resaltado de sintaxis semántico |
| williamboman/mason.nvim | Instalador de servidores LSP |
| neovim/nvim-lspconfig | Conector LSP oficial |
| hrsh7th/nvim-cmp | Motor de autocompletado |
| stevearc/conform.nvim | Formateador de código |
| windwp/nvim-autopairs | Cierre automático de caracteres |
| lukas-reineke/indent-blankline.nvim | Guías de indentación |

### Git, Ventanas y Terminal

| Plugin | Para qué sirve |
|--------|-------------|
| nvim-lualine/lualine.nvim | Barra de estado inferior |
| lewis6991/gitsigns.nvim | Indicadores de Git en el editor |
| akinsho/toggleterm.nvim | Terminal integrada |
| mrjones2014/smart-splits.nvim | Navegación entre paneles |

---

## 6. Requisitos de Experiencia de Usuario (UI/UX)

### Identidad Visual del Editor

- Fondo principal: `#110f18` (Negro con matiz morado)
- Acentos y resaltados: `#b388ff` (Morado vibrante)

### Paleta de Colores Detallada

```css
/* Jodify Violet Palette */
--color-jodify-violet-100: #f3e6ff;
--color-jodify-violet-200: #e7ccff;
--color-jodify-violet-300: #dbb3ff;
--color-jodify-violet-400: #c18fff;
--color-jodify-violet-500: #c380ff; /* Color de marca */
--color-jodify-violet-600: #b866ff;
--color-jodify-violet-700: #ac4dff;
--color-jodify-violet-800: #a033ff;
--color-jodify-violet-900: #941aff;
--color-jodify-violet-950: #8800ff;

/* Dark variants */
--color-jodify-dark-300: #828282;
--color-jodify-dark-400: #3f3f3f;
--color-jodify-dark-500: #2b2b2b;
--color-jodify-dark-900: #1b1c20;
--color-jodify-dark-1000: #0c0c0c;

/* Status colors */
--color-jodify-error-400: #ff5353;
--color-jodify-green-400: #d9ff00;
--color-jodify-green-500: #bdff00;
```

### TUI de Configuración

Los menús generados por `Huh` en el binario de Go deben usar la misma paleta morada/oscura para mantener consistencia de marca.

### Baja Fricción

El `which-key.nvim` estará configurado por defecto para que la curva de aprendizaje sea nula. Al presionar `<Espacio>`, el usuario verá un mapa de todas las acciones.

---

## 7. Wrapper Script (Jodify-Dev)

### Ubicación

- Windows: `%LOCALAPPDATA%\Programs\Jodify\Jodify-Dev.ps1`
- macOS: `~/.local/bin/Jodify-Dev`

### Flags Soportados

| Flag | Descripción |
|------|-----------|
| `--check` | Mostrar versión de Neovim |
| `--help` | Mostrar ayuda |
| (sin flags) | Abrir Neovim con NVIM_APPNAME=jodify |

### Código del Wrapper (Windows)

```powershell
# Jodify-Dev.ps1
param(
    [switch]$Check,
    [switch]$Help
)

if ($Check) {
    $nvimVersion = nvim --version | Select-Object -First 1
    Write-Host "Jodify Neovim: $nvimVersion"
    exit 0
}

if ($Help) {
    Write-Host "Jodify-Dev: Wrapper for isolated Jodify Neovim"
    Write-Host "Usage: Jodify-Dev [options]"
    Write-Host "  --check   Show Neovim version"
    Write-Host "  --help   Show this help"
    exit 0
}

$env:NVIM_APPNAME = "jodify"
nvim $args
```

---

## 8. Manejo de Errores y Casos Límite

- Si `Jodify-Setup` detecta que Scoop (Windows) o Brew (macOS) no están instalados, debe ofrecer instalarlos automáticamente.
- Si la carpeta objetivo de la configuración (`jodify`) ya existe, la TUI debe preguntar si se desea sobrescribir (haciendo un backup automático) o actualizar descargando la nueva versión del config.zip.
- El binario busca primero config.zip local, luego GitHub Releases.
- Fallas de red durante la descarga no deben corromper el entorno; la herramienta debe revertir los archivos a medias en caso de falla crítica.
- Si `scoop update *` actualiza Neovim a una versión incompatible, la configuración debe estar documentada para hacer rollback.
- Si no hay internet y no hay config.zip local, mostrar error claro con instrucciones.

---

## 9. Estructura del Repositorio

```
Jodify-Setup/
├── .goreleaser.yaml              ← Config de release automático
├── .github/
│   └── workflows/
│       └── release.yml           ← GitHub Action trigger
├── cmd/
│   └── jodify-setup/
│       └── main.go               ← Entry point del CLI
├── internal/                     ← Código interno (no exportado)
│   ├── cli/                      ← Comandos Cobra
│   │   ├── root.go               ← Comando raíz
│   │   ├── version.go            ← Comando: jodify-setup version
│   │   ├── install.go            ← Comando: jodify-setup install
│   │   └── uninstall.go          ← Comando: jodify-setup uninstall
│   ├── installer/                ← Lógica de instalación
│   │   ├── config.go             ← Descarga y extracción de config.zip
│   │   ├── wrapper.go            ← Generación de Jodify-Dev wrapper
│   │   └── dependencies.go       ← Gestión de dependencias del sistema
│   ├── platform/                 ← Abstracción de plataforma
│   │   ├── windows.go            ← Windows-specific logic
│   │   ├── darwin.go             ← macOS-specific logic
│   │   └── common.go             ← Interfaces y utilidades compartidas
│   └── tui/                      ← Componentes de TUI (futuro)
│       └── components.go
├── pkg/                          ← Código exportable (librerías reutilizables)
│   └── utils/
│       ├── fs.go                 ← Utilidades de filesystem
│       ├── net.go                ← Utilidades de red
│       └── path.go               ← Utilidades de paths cross-platform
├── config/                       ← ESTO se empaqueta como jodify-config.zip
│   ├── init.lua
│   ├── lua/
│   │   └── jodify/
│   │       ├── core/
│   │       │   └── options.lua
│   │       ├── plugins/
│   │       │   └── init.lua
│   │       ├── keymaps/
│   │       │   └── init.lua
│   │       └── theme/
│   │           └── init.lua
│   └── README.md
├── docs/                         ← Documentación adicional
├── go.mod                        ← Módulo Go
├── go.sum                        ← Checksums de dependencias
└── README.md                     ← Documentación principal
```

### Notas sobre la Estructura

- **`cmd/`**: Contiene solo el entry point (`main.go`). La lógica de comandos está en `internal/cli/`.
- **`internal/`**: Código privado del proyecto. Subpaquetes organizados por responsabilidad:
  - `cli/`: Comandos Cobra (version, install, uninstall, update)
  - `installer/`: Lógica de descarga, extracción, y generación de wrappers
  - `platform/`: Abstracción para manejar diferencias Windows/macOS
  - `tui/`: Componentes visuales Huh (para fases futuras)
- **`pkg/`**: Código que podría ser reutilizado por otros proyectos (utils genéricas)
- **`config/`**: Configuración Lua de Neovim que se empaqueta como `jodify-config.zip` en releases

---

## 10. GoReleaser: Auto-Release

###¿Qué es GoReleaser?

Herramienta que automatiza el release de proyectos Go. Compila, empaqueta y publica automáticamente.

### Configuración (.goreleaser.yaml)

```yaml
project_name: jodify-setup

release:
  github:
    owner: tu-usuario
    name: jodify-setup

archives:
  - id: jodify-setup
    builds:
      - '{{ .Env.GOOS }}_{{ .Env.GOARCH }}'
    format: zip

checksum:
  name: 'checksums.txt'

snapshot:
  name_template: "snapshot-{{ .ShortCommit }}"

release_notes_template: |
  ## Cambios
  - Version: {{ .Tag }}
  - Compilado: {{ .Date }}
```

### Workflow de Release

```bash
# Opción 1: Manual
goreleaser release --clean

# Opción 2: Automático (GitHub Action)
git tag v1.0.0
git push origin v1.0.0
```

Al pushear un tag, GitHub Action:
1. Compila Go (Windows + macOS)
2. Empaqueta `config/` como `jodify-config.zip`
3. Crea GitHub Release
4. Sube binario + config.zip

---

## 11. Notas de Implementación

- Los plugins se cargan mediante **lazy.nvim** para mínimo tiempo de arranque
- El stack MCP usa el de Gentle.ai (no modificar en esta iteración)
- El wrapper se regenera en cada `Jodify-Setup --update`
- El binario descarga config.zip desde GitHub Releases automáticamente