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
vim.api.nvim_create_autocmd("User", {
  pattern = "LazySync",
  callback = function()
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
    local stats = require("lazy").stats()
    vim.schedule(function()
      vim.print(string.format(
        "Jodify loaded %d plugins in %.2fms",
        stats.count,
        stats.startuptime
      ))
    end)
  end,
})

-- Trigger sync on startup with delay
vim.defer_fn(function()
  require("lazy").sync({ show = false })
end, 2000)
