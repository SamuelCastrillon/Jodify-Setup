-- Jodify Neovim Configuration
-- Lazy.nvim plugin manager setup and plugin specifications

-- ============================================
-- Bootstrap lazy.nvim
-- ============================================

local lazypath = vim.fn.stdpath("data") .. "/lazy/lazy.nvim"

-- Check if lazy.nvim is not installed (compatible with Neovim 0.11+)
local lazy_path_exists = (vim.uv or vim.loop).fs_stat(lazypath)
if not lazy_path_exists then
  -- Clone lazy.nvim
  vim.fn.system({
    "git",
    "clone",
    "--filter=blob:none",
    "https://github.com/folke/lazy.nvim.git",
    "--branch=stable",
    lazypath,
  })
end

-- Add lazy.nvim to the runtime path
vim.opt.rtp:prepend(lazypath)

-- ============================================
-- Plugin Specifications
-- ============================================

require("lazy").setup({
  -- ========================================
  -- Core, Interfaz y Navegación
  -- ========================================

  -- Plugin manager
  {
    "folke/lazy.nvim",
    event = "VimEnter",
  },

  -- File explorer
  {
    "nvim-neo-tree/neo-tree.nvim",
    branch = "v3.x",
    cmd = "Neotree",
    dependencies = {
      "nvim-lua/plenary.nvim",
      "MunifTanjim/nui.nvim",
    },
  },

  -- Fuzzy finder
  {
    "nvim-telescope/telescope.nvim",
    cmd = "Telescope",
    dependencies = {
      "nvim-lua/plenary.nvim",
      "nvim-telescope/telescope-fzf-native.nvim",
    },
  },

  -- Keymaps helper
  {
    "folke/which-key.nvim",
    event = "VeryLazy",
  },

  -- Buffer tabs
  {
    "akinsho/bufferline.nvim",
    event = "VeryLazy",
    dependencies = "nvim-tree/nvim-web-devicons",
  },

  -- Color theme
  {
    "folke/tokyonight.nvim",
    lazy = false,
    priority = 1000,
  },

  -- Dashboard
  {
    "goolord/alpha-nvim",
    event = "VimEnter",
    dependencies = "nvim-tree/nvim-web-devicons",
  },

  -- ========================================
  -- Inteligencia de Código, LSP y Edición
  -- ========================================

  -- Treesitter - syntax highlighting
  {
    "nvim-treesitter/nvim-treesitter",
    build = ":TSUpdate",
    event = { "BufReadPost", "BufNewFile" },
  },

  -- LSP installer
  {
    "williamboman/mason.nvim",
    cmd = "Mason",
    build = ":MasonUpdate",
    dependencies = {
      "williamboman/mason-lspconfig.nvim",
      "neovim/nvim-lspconfig",
    },
  },

  -- LSP config
  {
    "neovim/nvim-lspconfig",
    event = { "BufReadPre", "BufNewFile" },
    dependencies = {
      "williamboman/mason.nvim",
      "williamboman/mason-lspconfig.nvim",
    },
  },

  -- Completion
  {
    "hrsh7th/nvim-cmp",
    event = "InsertEnter",
    dependencies = {
      "hrsh7th/cmp-nvim-lsp",
      "hrsh7th/cmp-buffer",
      "hrsh7th/cmp-path",
      "hrsh7th/cmp-cmdline",
      "L3MON4D3/LuaSnip",
      "saadparwaiz1/cmp_luasnip",
      "rafamadriz/friendly-snippets",
    },
  },

  -- Code formatter
  {
    "stevearc/conform.nvim",
    event = { "BufReadPre", "BufNewFile" },
    config = function()
      require("conform").setup()
    end,
  },

  -- Autopairs
  {
    "windwp/nvim-autopairs",
    event = "InsertEnter",
    config = function()
      require("nvim-autopairs").setup()
    end,
  },

  -- Indent guides
  {
    "lukas-reineke/indent-blankline.nvim",
    event = { "BufReadPost", "BufNewFile" },
    main = "ibl",
  },

  -- ========================================
  -- Git, Ventanas y Terminal
  -- ========================================

  -- Status line
  {
    "nvim-lualine/lualine.nvim",
    event = "VeryLazy",
    dependencies = "nvim-tree/nvim-web-devicons",
  },

  -- Git integration
  {
    "lewis6991/gitsigns.nvim",
    event = { "BufReadPre", "BufNewFile" },
    dependencies = "nvim-lua/plenary.nvim",
  },

  -- Terminal
  {
    "akinsho/toggleterm.nvim",
    cmd = { "ToggleTerm", "TermExec" },
    build = ":ToggleTermToggleAll",
  },

  -- Smart splits
  {
    "mrjones2014/smart-splits.nvim",
    event = { "BufReadPost", "BufNewFile" },
  },

  -- ========================================
  -- Utils
  -- ========================================

  -- Comment plugin
  {
    "numToStr/Comment.nvim",
    event = { "BufReadPost", "BufNewFile" },
  },

  -- Nvim Web Devicons
  {
    "nvim-tree/nvim-web-devicons",
    lazy = true,
  },

  -- Nui components
  {
    "MunifTanjim/nui.nvim",
    lazy = true,
  },

  -- Plenary
  {
    "nvim-lua/plenary.nvim",
    lazy = true,
  },
}, {
  defaults = {
    -- By default, lazy.nvim will lazy-load plugins that are not loaded yet
    lazy = true,
    -- Install plugins matching these versions
    version = false,
    -- Don't automatically check for updates
    checker = false,
  },
  install = {
    -- Install missing plugins on startup
    missing = true,
    -- Show what will be installed
    colors = {
      error = "DiagnosticError",
      pending = "DiagnosticWarn",
      success = "DiagnosticOk",
    },
  },
  change_detection = {
    -- Don't automatically check for config changes
    enabled = false,
  },
  performance = {
    -- Disable loading lazy on startup for faster startup time
    rtp = {
      disabled_plugins = {
        "gzip",
        "tarPlugin",
        "tohtml",
        "tutor",
        "zipPlugin",
      },
    },
  },
  lockfile = vim.fn.stdpath("data") .. "/lazy-lock.json",
})

-- ============================================
-- Post-install hooks
-- ============================================

-- Telescope fzf native make
vim.api.nvim_create_autocmd("User", {
  pattern = "LazyInstall",
  callback = function(event)
    local plugin = event.plugin
    if plugin == "telescope-fzf-native.nvim" then
      vim.fn.system({ "make" }, { cwd = plugin })
    end
  end,
})

-- Treesitter install/update message
vim.api.nvim_create_autocmd("User", {
  pattern = "TSInstallSync",
  callback = function()
    vim.cmd("TSUpdateSync")
  end,
})
