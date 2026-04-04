-- Jodify Neovim Configuration
-- Main module initialization

-- ============================================
-- Core Options (must be first)
-- ============================================
require("jodify.core.options")

-- ============================================
-- Theme (must be before plugins)
-- ============================================
require("jodify.theme")

-- ============================================
-- Plugin Manager & Plugins
-- ============================================
require("jodify.lazy")

-- Load plugin configs
require("jodify.plugins.navigation")
require("jodify.plugins.lsp")
require("jodify.plugins.git_terminal")
require("jodify.plugins.alpha")

-- ============================================
-- Keymaps
-- ============================================
require("jodify.keymaps")

-- ============================================
-- Welcome message
-- ============================================
vim.api.nvim_create_autocmd("VimEnter", {
  callback = function()
    vim.defer_fn(function()
      vim.print("Welcome to Jodify Neovim Configuration!")
    end, 100)
  end,
})

-- ============================================
-- Final setup
-- ============================================

-- Report plugin load time
vim.api.nvim_create_autocmd("VeryLazy", {
  callback = function()
    local lazy_ok, lazy = pcall(require, "lazy")
    if lazy_ok then
      local stats = lazy.stats()
      vim.schedule(function()
        vim.print(string.format(
          "Jodify loaded %d plugins in %.2fms",
          stats.count,
          stats.startuptime
        ))
      end)
    end
  end,
})
