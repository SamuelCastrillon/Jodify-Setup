-- Jodify Neovim Configuration
-- Git, terminal and status line setup

-- ============================================
-- gitsigns.nvim
-- ============================================

require("gitsigns").setup({
  signs = {
    add = { text = "│" },
    change = { text = "│" },
    delete = { text = "󰍵" },
    topdelete = { text = "󰍵" },
    changedelete = { text = "│" },
    untracked = { text = "│" },
  },
  signcolumn = true,
  numhl = false,
  linehl = false,
  word_diff = false,
  watch_gitdir = {
    interval = 1000,
    follow_files = true,
  },
  attach_to_untracked = true,
  current_line_blame = false,
  current_line_blame_opts = {
    virt_text = true,
    virt_text_pos = "eol",
    delay = 1000,
    ignore_whitespace = false,
  },
  current_line_blame_formatter = "<author>, <author_time:%Y-%m-%d> - <summary>",
  sign_priority = 6,
  update_debounce = 100,
  status_formatter = nil,
  max_file_length = 40000,
  preview_config = {
    border = "rounded",
    style = "minimal",
    relative = "cursor",
    row = 0,
    col = 0,
  },
  on_attach = function(buffer)
    local gs = package.loaded.gitsigns

    local function map(mode, l, r, desc)
      vim.keymap.set(mode, l, r, { buffer = buffer, desc = desc })
    end

    -- Navigation
    map("n", "]c", function()
      if vim.wo.diff then return "]c" end
      vim.schedule(function() gs.next_hunk() end)
      return "<Ignore>"
    end, { expr = true, desc = "Next hunk" })

    map("n", "[c", function()
      if vim.wo.diff then return "[c" end
      vim.schedule(function() gs.prev_hunk() end)
      return "<Ignore>"
    end, { expr = true, desc = "Prev hunk" })

    -- Actions
    map("n", "<leader>gs", gs.stage_hunk, "Stage hunk")
    map("n", "<leader>gu", gs.undo_stage_hunk, "Undo stage hunk")
    map("n", "<leader>gr", gs.reset_hunk, "Reset hunk")
    map("v", "<leader>gs", function() gs.stage_hunk { vim.fn.line("."), vim.fn.line("v") } end, "Stage hunk")
    map("v", "<leader>gr", function() gs.reset_hunk { vim.fn.line("."), vim.fn.line("v") } end, "Reset hunk")

    map("n", "<leader>gS", gs.stage_buffer, "Stage buffer")
    map("n", "<leader>gR", gs.reset_buffer, "Reset buffer")

    map("n", "<leader>gp", gs.preview_hunk, "Preview hunk")
    map("n", "<leader>gb", gs.blame_line, "Blame line")

    map("n", "<leader>gd", gs.diffthis, "Diff this")
    map("n", "<leader>gD", function() gs.diffthis("~") end, "Diff this ~")

    map("n", "<leader>td", gs.toggle_deleted, "Toggle deleted")
  end,
})

-- ============================================
-- toggleterm.nvim
-- ============================================

require("toggleterm").setup({
  size = 20,
  open_mapping = [[<leader>tt]],
  hide_numbers = true,
  shade_filetypes = {},
  shade_terminals = true,
  shading_factor = "-1",
  start_in_insert = true,
  insert_mappings = true,
  persist_size = true,
  direction = "float",
  close_on_exit = true,
  shell = vim.o.shell,
  float_opts = {
    border = "curved",
    winblend = 0,
    highlights = {
      border = "Normal",
      background = "Normal",
    },
  },
  winbar = {
    enabled = false,
    name_formatter = function(term)
      return term.name
    end,
  },
  -- Extra terminal configurations
  terminals = {
    {
      cmd = "lazygit",
      direction = "float",
      on_open = function(term)
        vim.cmd("startinsert!")
        vim.api.nvim_buf_set_keymap(term.bufnr, "q", "<cmd>close<cr>", { noremap = true, silent = true })
      end,
    },
  },
})

-- Terminal keymaps
local keymap = vim.keymap.set

keymap("n", "<leader>tf", ":ToggleTerm direction=float<CR>", { noremap = true, silent = true })
keymap("n", "<leader>th", ":ToggleTerm direction=horizontal size=10<CR>", { noremap = true, silent = true })
keymap("n", "<leader>tv", ":ToggleTerm direction=vertical size=80<CR>", { noremap = true, silent = true })
keymap("n", "<leader>tg", ":ToggleTerm direction=float cmd=lazygit<CR>", { noremap = true, silent = true })

-- Terminal commands
vim.api.nvim_create_user_command("TermExec", function(opts)
  local count = vim.v.count1
  local direction = opts.args == "float" and "float" or nil
  local size = opts.fargs[1] and tonumber(opts.fargs[1]) or nil
  require("toggleterm").exec(opts.args, count, direction, size)
end, {
  nargs = "+",
  desc = "Execute command in terminal",
})

-- ============================================
-- lualine.nvim
-- ============================================

require("lualine").setup({
  options = {
    theme = "tokyonight",
    component_separators = { left = "│", right = "│" },
    section_separators = { left = "│", right = "│" },
    globalstatus = true,
    disabled_filetypes = { statusline = { "dashboard", "alpha", "starter" } },
    refresh = {
      statusline = 1000,
      tabline = 1000,
      winbar = 1000,
    },
  },
  sections = {
    lualine_a = { "mode" },
    lualine_b = { "branch" },
    lualine_c = {
      {
        "diagnostics",
        symbols = {
          error = "󰀦 ",
          warn = "󰀐 ",
          info = "󰋽 ",
          hint = "󰌵 ",
        },
      },
      { "filetype", icon_only = true },
      {
        "filename",
        path = 1,
        symbols = {
          modified = " ●",
          readonly = " 󰌾",
          unnamed = "",
        },
      },
    },
    lualine_x = {
      {
        "diff",
        symbols = {
          added = "󰐖",
          modified = "󰁜",
          removed = "󰍵",
        },
      },
      {
        function()
          local msg = "No LSP"
          local buf_ft = vim.api.nvim_buf_get_option(0, "filetype")
          local clients = vim.lsp.get_active_clients()
          if next(clients) == nil then
            return msg
          end
          for _, client in ipairs(clients) do
            if client.name ~= "null-ls" then
              local buf_client = false
              for _, bufnr in ipairs(vim.lsp.get_buffers()) do
                if bufnr == vim.api.nvim_get_current_buf() then
                  buf_client = true
                  break
                end
              end
              if buf_client then
                return client.name
              end
            end
          end
          return msg
        end,
        icon = "󰌿 ",
        color = { fg = "#b388ff" },
      },
    },
    lualine_y = { "location", "progress" },
    lualine_z = {
      function()
        return " " .. os.date("%R")
      end,
    },
  },
  inactive_sections = {
    lualine_a = {},
    lualine_b = {},
    lualine_c = {
      {
        "filename",
        path = 1,
        symbols = {
          modified = " ●",
          readonly = " 󰌾",
          unnamed = "",
        },
      },
    },
    lualine_x = { "location" },
    lualine_y = {},
    lualine_z = {},
  },
  tabline = {
    lualine_a = {
      {
        "buffers",
        show_filename_only = false,
        hide_filename = false,
        diagnostics = "nvim_lsp",
        -- Component symbol mappings
        symbols = {
          modified = " ●",
          readonly = " 󰌾",
          unnamed = "",
        },
      },
    },
    lualine_b = {},
    lualine_c = {},
    lualine_x = {},
    lualine_y = {},
    lualine_z = { require("tabtab").tabline_tabs },
  },
  extensions = { "neo-tree", "toggleterm", "lazy" },
})

-- ============================================
-- smart-splits.nvim
-- ============================================

require("smart-splits").setup({
  -- directions = { h, j, k, l },
  -- keys = h/j/k/l
  -- delay before performing the swap
  -- set to 0 to disable
  move_delay = 150,
  -- whether to resize the pane you're moving _to_
  -- when using swap, as opposed to moving the current pane
  resize_on_move = false,
  -- a function to run before moving cursor to a pane
  start_in_normal_mode = true,
  -- ignore_buffer = <function>, function(bufnr) -> boolean
  -- use this if you want to move over certain file types
  -- returning true will cause smart-splits to ignore that buffer
  -- ignore_buffer = function(bufnr)
  --   local bufname = vim.api.nvim_buf_get_name(bufnr)
  --   local ignored = { "NvimTree", "TelescopePrompt" }
  --   for _, v in ipairs(ignored) do
  --     if string.match(bufname, v) then
  --       return true
  --     end
  --   end
  -- end,
})

-- Smart splits keymaps
keymap("n", "<C-h>", require("smart-splits").move_left, { desc = "Move to left pane" })
keymap("n", "<C-j>", require("smart-splits").move_down, { desc = "Move to pane below" })
keymap("n", "<C-k>", require("smart-splits").move_up, { desc = "Move to pane above" })
keymap("n", "<C-l>", require("smart-splits").move_right, { desc = "Move to right pane" })

keymap("n", "<leader><leader>h", require("smart-splits").swap_left, { desc = "Swap with left pane" })
keymap("n", "<leader><leader>j", require("smart-splits").swap_down, { desc = "Swap with pane below" })
keymap("n", "<leader><leader>k", require("smart-splits").swap_up, { desc = "Swap with pane above" })
keymap("n", "<leader><leader>l", require("smart-splits").swap_right, { desc = "Swap with right pane" })

-- ============================================
-- Comment.nvim
-- ============================================

require("Comment").setup({
  toggler = {
    line = "gcc",
    block = "gbc",
  },
  opleader = {
    line = "gc",
    block = "gb",
  },
  extra = {
    above = "gcO",
    below = "gco",
    eol = "gcA",
  },
  mappings = {
    basic = true,
    extra = true,
  },
  pre_hook = require("ts_context_commentstring").create_pre_hook(),
})
