# diseño de Fonction: Auto-Install Scoop (Windows)

## 1. Phase de Détection

Cuando el usuario ejecuta `Jodify-Setup` en Windows, el primer paso es verificar si Scoop está instalado:

```go
// Pseudo-código Go
func detectScoop() (bool, string) {
    // 1. Verificar si 'scoop' está en PATH
    path := os.Getenv("PATH")
    paths := strings.Split(path, ";")
    
    for _, p := range paths {
        if strings.Contains(strings.ToLower(p), "scoop") {
            // Scoop encontrado en PATH
            return true, p
        }
    }
    
    // 2. Verificar si scoop.exe existe en ubicaciones comunes
    commonPaths := []string{
        os.Getenv("USERPROFILE") + "\\scoop",
        os.Getenv("LOCALAPPDATA") + "\\Programs\\scoop",
    }
    
    for _, p := range commonPaths {
        if _, err := os.Stat(p + "\\shims\\scoop.exe"); err == nil {
            return true, p
        }
    }
    
    return false, ""
}
```

**Resultado de detección:**

| Estado | Acción |
|--------|--------|
| Scoop encontrado | ✅ Continuar con instalación normal |
| No encontrado | ℹ️ Mostrar TUI de oferta de instalación automática |

---

## 2. UX de la TUI (con Huh)

Cuando Scoop NO está instalado, la TUI muestra:

```
╭──────────────────────────────────────────────────────╮
│  🔧 Jodify-Setup                                     │
├──────────────────────────────────────────────────────┤
│                                                      │
│  ⚠️  Scoop no está instalado                        │
│                                                      │
│  Scoop es requerido para instalar las dependencias  │
│  de Neovim de forma administrado.                  │
│                                                      │
│  [ Instalar Scoop automaticamente ]  (recomendado)    │
│  [ Usar instalación existente ]                    │
│  [ Cancelar ]                                        │
│                                                      │
│  Powered by Scoop • get.scoop.sh                    │
╰──────────────────────────────────────────────────────╯
```

### Opción A: Instalar Scoop Automáticamente (Default)

Esta es la **experiencia de baja fricción** que el PRD menciona. El flujo:

1. Usuario presiona `Enter` en [ Instalar Scoop automáticamente ]
2. Huh muestra un componente de `Progress` con el log en vivo
3. El comando de instalación se ejecuta en un subprocess:

```go
// Go - Ejecución del installer de Scoop
func installScoop() error {
    cmd := exec.Command("powershell", "-ExecutionPolicy", "Bypass", "-Command", 
        `Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope CurrentUser -Force; `
        + `Invoke-RestMethod -Uri https://get.scoop.sh | Invoke-Expression`)
    
    // Capturar output en tiempo real para la TUI
    stdout, _ := cmd.StdoutPipe()
    stderr, _ := cmd.StderrPipe()
    
    go func() {
        scanner := bufio.NewScanner(stdout)
        for scanner.Scan() {
            // Enviar cada línea al canal de log de Huh
            logChannel <- scanner.Text()
        }
    }()
    
    return cmd.Run()
}
```

**Log esperado visible en TUI:**

```
╭──────────────────────────────────────────────────────╮
│  🔧 Instalando Scoop...                              │
├──────────────────────────────────────────────────────┤
│  ████████████████████░░░░░░░░  75%                  │
│                                                      │
│  Downloading scoop...                                │
│  Extracting...                                       │
│  Creating shim...                                   │
│  Installing main bucket...                           │
│                                                      │
╰──────────────────────────────────────────────────────╯
```

### Opción B: Usar Instalación Existente

Si el usuario eligió esta opción, mostrar un input para que ingrese la ruta manual:

```
Ingrese la ruta donde tiene instalado Scoop:
[ ~/scoop............................... ]
     [ Aceptar ]   [ Cancelar ]
```

---

## 3. Update Proposed al PRD

### Nuevo Requisito Funcional (agregar en secciones 4)

```
RF1.5: Como usuario de Windows, si Scoop no está instalado,
debo ver una opción para instalarlo automáticamente con
un solo clic, sin necesidad de abortar o hacerlo manualmente.
```

### Nueva Historia de Usuario

> **RF1.5 (Auto-Install Scoop):** Como usuario de Windows, al ejecutar `Jodify-Setup`, automáticamente se detecta si Scoop está instalado. Si no lo está, se me ofrece la opción de instalarlo con un solo clic. El instalador muestra progreso en tiempo real en la TUI. Una vez instalado, Scoop se configura automáticamente y la instalación de Jodify continúa sin más intervención.

---

## 4. Détails Técnicos

### 4.1 Verificación Post-Instalación

Después de installer Scoop, verificar que quedó correctamente instalado:

```go
func verifyScoopInstall() bool {
    // Ejecutar 'scoop --version'
    cmd := exec.Command("scoop", "--version")
    output, err := cmd.CombinedOutput()
    
    if err != nil {
        return false
    }
    
    // Verificar que el output contenga version
    return strings.Contains(string(output), "Version")
}
```

### 4.2 Fallback: Si Auto-Install Falla

Si el auto-install falla, mostrar un mensaje claro con instrucciones manuales:

```
╭──────────────────────────────────────────────────────╮
│  ❌ Error al instalar Scoop automáticamente          │
├──────────────────────────────────────────────────────┤
│                                                      │
│  El instalador automático no pudo completar.            │
│  Por favor, instale Scoop manualmente:                 │
│                                                      │
│  1. Abra PowerShell como Administrador              │
│  2. Ejecute:                                         │
│     Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope CurrentUser
│     Invoke-RestMethod -Uri https://get.scoop.sh | Invoke-Expression
│                                                      │
│  [ Reintentar ]    [ Salir ]                         │
╰──────────────────────────────────────────────────────╯
```

### 4.3 Compatibilidad con PowerShell 5.1 vs 7

El comando de instalación funciona en ambas versiones, pero:

| Versión | ¿Funciona? | Notas |
|--------|-----------|-------|
| PowerShell 5.1 | ✅ Sí | Requiere ejecutar como-admin para Set-ExecutionPolicy |
| PowerShell 7 | ✅ Sí | No requiere elevation para CurrentUser |
| Windows Terminal | ✅ Sí | Recomendado |
| cmd.exe | ❌ No | Debe usar `powershell -Command "..."` |

---

## 5. Resumen del Flujo Completo

```
Jodify-Setup (Go CLI)
        │
        ├─[Windows?]──► Detectar Scoop en PATH
        │                      │
        │              ┌──────┴──────┐
        │              │             │
        │         [Encontrado]    [NO Encontrado]
        │              │             │
        │              ▼             ▼
        │      Continuar con    Mostrar TUI Huh:
        │      instalación      ┌─────────────────┐
        │      normal        │  "¿Instalar   │
        │                 │   Scoop auto?" │
        │                 └─────────────┘
        │                        │
        │              ┌─────────┴─────────┐
        │              │                    │
        │      [Usuario acepta]      [Usuario cancela]
        │              │                    │
        │              ▼                    ▼
        │      Ejecutar installer   Mostrar msg de
        │      + Progress TUI     error + exit(1)
        │              │
        │              ▼
        │      Verificar install
        │              │
        │       ┌──────┴──────┐
        │       │   [Éxito]    │
        │       │       │      │
        │       ▼       ▼      ▼
        │   Continuar  Abortar
        │   con flujo normal
```

---

## 6. Referencias

- Comando oficial de instalación: https://get.scoop.sh
- Repositorio: https://github.com/ScoopInstaller/Scoop
- Documentación: https://scoop.sh/docs