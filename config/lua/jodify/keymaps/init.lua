-- Jodify Neovim Configuration
-- Keymaps setup

local keymap = vim.keymap.set
local opts = { noremap = true, silent = true }
local expr_opts = { noremap = true, silent = true, expr = true }

-- ============================================
-- General Keymaps
-- ============================================

-- Better window navigation
keymap("n", "<C-h>", "<C-w>h", opts)
keymap("n", "<C-j>", "<C-w>j", opts)
keymap("n", "<C-k>", "<C-w>k", opts)
keymap("n", "<C-l>", "<C-w>l", opts)

-- Resize windows
keymap("n", "<C-Up>", ":resize -2<CR>", opts)
keymap("n", "<C-Down>", ":resize +2<CR>", opts)
keymap("n", "<C-Left>", ":vertical resize -2<CR>", opts)
keymap("n", "<C-Right>", ":vertical resize +2<CR>", opts)

-- Move lines up/down
keymap("n", "<A-j>", ":m .+1<CR>==", opts)
keymap("n", "<A-k>", ":m .-2<CR>==", opts)
keymap("v", "<A-j>", ":m '>+1<CR>gv=gv", opts)
keymap("v", "<A-k>", ":m '<-2<CR>gv=gv", opts)

-- Better indenting
keymap("v", "<", "<gv", opts)
keymap("v", ">", ">gv", opts)

-- Clear highlights
keymap("n", "<leader>h", ":nohlsearch<CR>", opts)

-- Save and quit
keymap("n", "<leader>w", ":w<CR>", opts)
keymap("n", "<leader>q", ":q<CR>", opts)
keymap("n", "<leader>Q", ":qa!<CR>", opts)

-- Do not copy to clipboard on delete
keymap("n", "d", '"d', opts)
keymap("n", "D", '"D', opts)
keymap("v", "d", '"d', opts)
keymap("v", "D", '"D', opts)
keymap("v", "x", '"x', opts)
keymap("v", "X", '"X', opts)
keymap("n", "c", '"c', opts)
keymap("n", "C", '"C', opts)
keymap("v", "c", '"c', opts)
keymap("v", "C", '"C', opts)

-- Copy to clipboard
keymap("v", "<leader>y", '"+y', opts)
keymap("n", "<leader>y", '"+y', opts)
keymap("n", "<leader>Y", '"+Y', opts)

-- Paste from clipboard
keymap("n", "<leader>p", '"+p', opts)
keymap("n", "<leader>P", '"+P', opts)
keymap("v", "<leader>p", '"+p', opts)
keymap("v", "<leader>P", '"+P', opts)

-- ============================================
-- Plugin Keymaps
-- ============================================

-- Neo-tree
keymap("n", "<leader>e", ":Neotree toggle<CR>", opts)
keymap("n", "<leader>o", ":Neotree focus<CR>", opts)

-- Telescope
keymap("n", "<leader>ff", "<cmd>Telescope find_files<cr>", opts)
keymap("n", "<leader>fg", "<cmd>Telescope live_grep<cr>", opts)
keymap("n", "<leader>fb", "<cmd>Telescope buffers<cr>", opts)
keymap("n", "<leader>fh", "<cmd>Telescope help_tags<cr>", opts)
keymap("n", "<leader>fc", "<cmd>Telescope commands<cr>", opts)
keymap("n", "<leader>fr", "<cmd>Telescope oldfiles<cr>", opts)
keymap("n", "<leader>fs", "<cmd>Telescope grep_string<cr>", opts)
keymap("n", "<leader>ft", "<cmd>Telescope<cr>", opts)

-- Git signs
keymap("n", "]c", ":Gitsigns next_hunk<CR>", opts)
keymap("n", "[c", ":Gitsigns prev_hunk<CR>", opts)
keymap("n", "<leader>gs", ":Gitsigns stage_hunk<CR>", opts)
keymap("n", "<leader>gu", ":Gitsigns undo_stage_hunk<CR>", opts)
keymap("n", "<leader>gr", ":Gitsigns reset_hunk<CR>", opts)
keymap("n", "<leader>gS", ":Gitsigns stage_buffer<CR>", opts)
keymap("n", "<leader>gR", ":Gitsigns reset_buffer<CR>", opts)
keymap("n", "<leader>gp", ":Gitsigns preview_hunk<CR>", opts)
keymap("n", "<leader>gb", ":Gitsigns blame_line<CR>", opts)
keymap("n", "<leader>gd", ":Gitsigns diffthis<CR>", opts)
keymap("n", "<leader>gD", ":Gitsigns diffthis ~<CR>", opts)

-- Fugitive (comandos git)
keymap("n", "<leader>gg", ":Git<CR>", opts)
keymap("n", "<leader>gco", ":Git co<CR>", opts)
keymap("n", "<leader>gci", ":Git ci?<CR>", opts)
keymap("n", "<leader>gbl", ":Git blame<CR>", opts)

-- Diffview (diffs visuales)
keymap("n", "<leader>gvo", ":DiffviewOpen<CR>", opts)
keymap("n", "<leader>gvc", ":DiffviewClose<CR>", opts)
keymap("n", "<leader>gvh", ":DiffviewFileHistory<CR>", opts)

-- ToggleTerm
keymap("n", "<leader>tt", ":ToggleTerm<CR>", opts)
keymap("n", "<leader>tf", ":ToggleTerm direction=float<CR>", opts)
keymap("n", "<leader>th", ":ToggleTerm direction=horizontal size=10<CR>", opts)
keymap("n", "<leader>tv", ":ToggleTerm direction=vertical size=80<CR>", opts)

-- LSP
keymap("n", "gd", ":lua vim.lsp.buf.definition()<CR>", opts)
keymap("n", "gD", ":lua vim.lsp.buf.declaration()<CR>", opts)
keymap("n", "gi", ":lua vim.lsp.buf.implementation()<CR>", opts)
keymap("n", "gr", ":lua vim.lsp.buf.references()<CR>", opts)
keymap("n", "K", ":lua vim.lsp.buf.hover()<CR>", opts)
keymap("n", "<leader>rn", ":lua vim.lsp.buf.rename()<CR>", opts)
keymap("n", "<leader>ca", ":lua vim.lsp.buf.code_action()<CR>", opts)
keymap("n", "<leader>e", ":lua vim.diagnostic.open_float()<CR>", opts)
keymap("n", "[d", ":lua vim.diagnostic.goto_prev()<CR>", opts)
keymap("n", "]d", ":lua vim.diagnostic.goto_next()<CR>", opts)

-- Mason
keymap("n", "<leader>ml", ":Mason<CR>", opts)
keymap("n", "<leader>mi", ":MasonInstall<CR>", opts)
keymap("n", "<leader>mu", ":MasonUpdateAll<CR>", opts)

-- Comment
keymap("n", "<leader>/", "gcc", { remap = true })
keymap("v", "<leader>/", "gc", { remap = true })

-- Smart splits (vim-beacon for jumping cursor)
-- Note: smart-splits.nvim provides different API, using basic keymaps instead
keymap("n", "<leader><leader>", "<cmd>WhichKey<CR>", opts)

-- ============================================
-- Which-key Setup
-- ============================================

local function setup_whichkey()
  local wk = require("which-key")

  wk.setup({
    window = {
      border = "single",
      position = "bottom",
      padding = { 1, 2, 1, 2 },
    },
    layout = {
      align = "center",
    },
    triggers = "auto",
    show_guide = true,
    timeout = 300,
    ignore_missing = false,
  })

  -- Leader key groups
  local groups = {
    -- File explorer
    e = "Explorer (Neotree)",
    o = "Focus Explorer",

    -- Telescope
    f = {
      name = "Telescope",
      f = "Find Files",
      g = "Live Grep",
      b = "Buffers",
      h = "Help Tags",
      c = "Commands",
      r = "Recent Files",
      s = "Search String",
      t = "Telescope",
    },

    -- Git
    g = {
      name = "Git",
      s = "Stage Hunk",
      u = "Undo Stage",
      r = "Reset Hunk",
      S = "Stage Buffer",
      R = "Reset Buffer",
      p = "Preview Hunk",
      b = "Blame Line",
      d = "Diff This",
      D = "Diff This (~)",
    },

    -- Terminal
    t = {
      name = "Terminal",
      t = "Toggle Terminal",
      f = "Float",
      h = "Horizontal",
      v = "Vertical",
    },

    -- LSP
    m = {
      name = "Mason",
      l = "Open",
      i = "Install",
      u = "Update All",
    },
  }

  wk.register(groups, { prefix = "<leader>" })
end

return {
  setup_whichkey = setup_whichkey,
}
