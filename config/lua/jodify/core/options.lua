-- Jodify Neovim Configuration
-- Core options for Neovim

local o = vim.opt

-- ============================================
-- General
-- ============================================

o.mouse = "nv"               -- Enable mouse support
o.clipboard = "unnamedplus"  -- Use system clipboard
o.swapfile = false            -- Disable swap file
o.backup = false             -- Disable backup
o.writebackup = false        -- Don't create backup before writing
o.undofile = true            -- Enable persistent undo
o.undodir = vim.fn.stdpath("data") .. "/undo"

-- ============================================
-- UI
-- ============================================

o.number = true              -- Show line numbers
o.relativenumber = true      -- Show relative line numbers
o.cursorline = true         -- Highlight current line
o.signcolumn = "auto"       -- Show sign column
o.colorcolumn = "80"        -- Show color column at 80 chars
o.foldmethod = "expr"       -- Fold based on expression
o.foldexpr = "nvim_treesitter#foldexpr()"  -- Treesitter fold
o.fillchars = {
  foldopen = "",
  foldclose = "",
  fold = " ",
  eob = " ",
}
o.showmode = false           -- Don't show mode in cmdline (lualine shows it)
o.showcmd = true            -- Show commands in cmdline
o.cmdheight = 1             -- Command line height
o.pumheight = 10            -- Popup menu height
o.conceallevel = 0          -- Don't hide concealed characters
o.scrolloff = 8             -- Scroll offset
o.sidescrolloff = 8         -- Side scroll offset
o.termguicolors = true      -- Enable true colors

-- ============================================
-- Indentation
-- ============================================

o.tabstop = 4               -- Tab width
o.shiftwidth = 4            -- Shift width
o.softtabstop = 4           -- Soft tab width
o.expandtab = true          -- Convert tabs to spaces
o.smartindent = true        -- Smart indentation
o.wrap = false              -- Don't wrap lines
o.linebreak = true          -- Break lines at word boundaries

-- ============================================
-- Search
-- ============================================

o.ignorecase = true        -- Ignore case in search
o.smartcase = true          -- Smart case (uppercase overrides ignorecase)
o.incsearch = true          -- Incremental search
o.hlsearch = true           -- Highlight search results

-- ============================================
-- Performance
-- ============================================

o.updatetime = 50           -- Faster completion
o.timeout = true            -- Enable timeout
o.timeoutlen = 300          -- Timeout for key sequences
o.redrawtime = 1500         -- Time for redraws
o.ttimeoutlen = 10         -- Timeout for keycodes

-- ============================================
-- Completion
-- ============================================

o.completeopt = { "menu", "menuone", "noselect" }
o.shortmess:append("c")    -- Don't show completion info messages

-- ============================================
-- Splits and Windows
-- ============================================

o.splitright = true         -- Open splits to the right
o.splitbelow = true         -- Open splits below
o.equalalways = false      -- Don't equalize window sizes

-- ============================================
-- Wild Menu
-- ============================================

o.wildmode = "longest:full"     -- Longest match first, then menu
o.wildmenu = true               -- Enable wild menu
o.wildignore:append({
  "*.o",
  "*.obj",
  "*.dylib",
  "*.bin",
  "*.dll",
  "*.exe",
  "*.jpg",
  "*.png",
  "*.jpeg",
  "*.bmp",
  "*.gif",
  "*.ico",
  "*.cur",
  "*.pyc",
  "*.pkl",
  "*.zip",
  "*.tar.gz",
  "*.rar",
  "*.7z",
  "*.doc",
  "*.docx",
  "*.pdf",
  "*.png",
  "*.mp3",
  "*.mp4",
  "*.mkv",
  "*.avi",
  "*.mov",
  "*.wmv",
  "*.flv",
  "*.webm",
  "*.opk",
  "*.gx",
  "*.rar",
  "*.local",
  "*/.git/*",
  "*/.svn/*",
  "*/__pycache__/*",
  "*/node_modules/*",
  "*/.DS_Store",
})

-- ============================================
-- Leader keys
-- ============================================

vim.g.mapleader = " "
vim.g.maplocalleader = "\\"

-- ============================================
-- Global variables
-- ============================================

vim.g.jodify_loaded = true
