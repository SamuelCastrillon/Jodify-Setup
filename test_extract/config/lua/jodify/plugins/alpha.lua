-- Jodify Neovim Configuration
-- Alpha-nvim dashboard setup

local alpha = require("alpha")
local dashboard = require("alpha.themes.dashboard")

-- Jodify ASCII Logo
local logo = {
[[]],
[[]],
[[  _          _ _        _    _ _     _ ]],
[[ | |__  _ __(_) |_ __ _| |__| (_)___| |]],
[[ | '_ \| '__| | __/ _` | '_ \ | / __| |]],
[[ | |_) | |  | | || (_| | |_) || \__ \_|]],
[[ |_.__/|_|  |_|\__\__,_|_.__/ |_|___(_)]],
[[]],
}

-- Header
dashboard.section.header.val = logo
dashboard.section.header.opts = {
  position = "center",
  hl = "JodifyViolet",
}

-- Buttons
local button_section = {
  dashboard.button("e", "󰈔  New File", ":ene <BAR> startinsert <CR>"),
  dashboard.button("SPC f f", "󰍉  Find File", ":Telescope find_files<CR>"),
  dashboard.button("SPC f r", "󰗡  Recent Files", ":Telescope oldfiles<CR>"),
  dashboard.button("SPC f g", "󰊄  Live Grep", ":Telescope live_grep<CR>"),
  dashboard.button("SPC e", "󰙅  File Explorer", ":Neotree<CR>"),
  dashboard.button("q", "󰗿  Quit", ":qa<CR>"),
}

dashboard.section.buttons.val = button_section
dashboard.section.buttons.opts = {
  position = "center",
  hl = "AlphaButtons",
  help = false,
  -- shorten_path = true,
}

-- Footer
local function get_version()
  local ver = vim.version()
  return string.format("v%d.%d.%d", ver.major, ver.minor, ver.patch)
end

local function get_plugins_count()
  local ok, lazy = pcall(require, "lazy")
  if ok then
    return lazy.stats().count
  end
  return 0
end

dashboard.section.footer.val = {
  " ",
  "  󰀦 Neovim " .. get_version() .. "  󰀦 " .. get_plugins_count() .. " plugins loaded",
  "  󰀦 " .. vim.fn.strftime("%Y-%m-%d"),
}
dashboard.section.footer.opts = {
  position = "center",
  hl = "AlphaFooter",
}

-- Config
dashboard.opts.layout = {
  { type = "padding", val = 2 },
  dashboard.section.header,
  { type = "padding", val = 2 },
  dashboard.section.buttons,
  { type = "padding", val = 1 },
  dashboard.section.footer,
}

-- Disable folding on alpha
vim.api.nvim_create_autocmd("FileType", {
  pattern = "alpha",
  callback = function()
    vim.opt.foldmethod = "manual"
    vim.opt.foldlevel = 99
  end,
})

-- Disable statusline on alpha
vim.api.nvim_create_autocmd("FileType", {
  pattern = "alpha",
  callback = function()
    vim.opt.laststatus = 0
    vim.opt.showtabline = 0
  end,
})

-- Restore statusline when leaving alpha
vim.api.nvim_create_autocmd("FileType", {
  pattern = "alpha",
  callback = function()
    vim.api.nvim_create_autocmd("BufUnload", {
      once = true,
      callback = function()
        vim.opt.laststatus = 3
        vim.opt.showtabline = 2
      end,
    })
  end,
})

-- Setup alpha
alpha.setup(dashboard.opts)

-- Start alpha when Neovim opens with no args
vim.api.nvim_create_autocmd("VimEnter", {
  callback = function()
    local args = vim.api.nvim_get_args()
    if #args == 0 then
      -- Check if there are any buffers open
      local buffers = vim.api.nvim_list_bufs()
      local has_open_buffer = false
      for _, buf in ipairs(buffers) do
        if vim.api.nvim_buf_is_loaded(buf) and vim.api.nvim_buf_get_name(buf) ~= "" then
          has_open_buffer = true
          break
        end
      end
      if not has_open_buffer then
        alpha.start()
      end
    end
  end,
})
