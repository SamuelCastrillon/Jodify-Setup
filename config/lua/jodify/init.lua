-- Jodify Neovim Configuration
-- Main module initialization

-- ============================================
-- Core Options (must be first)
-- ============================================
require("jodify.core.options")

-- ============================================
-- Plugin Manager (must load first to install plugins)
-- ============================================
require("jodify.lazy")

-- ============================================
-- Automatically sync and load plugins
-- ============================================

-- Function to load all plugins after sync
local function load_plugins()
  -- Theme (after plugins are installed)
  require("jodify.theme")

  -- Load plugin configs
  require("jodify.plugins.navigation")
  require("jodify.plugins.lsp")
  require("jodify.plugins.git_terminal")
  require("jodify.plugins.alpha")

  -- Keymaps
  require("jodify.keymaps")

  -- Welcome message
  vim.print("Welcome to Jodify Neovim Configuration!")

  -- Report plugin load time
  local ok, lazy = pcall(require, "lazy")
  if ok then
    local stats = lazy.stats()
    vim.schedule(function()
      vim.print(string.format(
        "Jodify loaded %d plugins in %.2fms",
        stats.count,
        stats.startuptime
      ))
    end)
  end
end

-- Trigger sync on startup with delay, then load plugins
vim.defer_fn(function()
  require("lazy").sync({ show = false })
  
  -- Load plugins after sync completes
  vim.defer_fn(load_plugins, 500)
end, 2000)
