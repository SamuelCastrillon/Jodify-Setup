-- Jodify Neovim Configuration
-- Main module initialization

-- Set up basic Neovim options
vim.opt.number = true
vim.opt.relativenumber = true
vim.opt.tabstop = 4
vim.opt.shiftwidth = 4
vim.opt.expandtab = true
vim.opt.smartindent = true
vim.opt.wrap = false
vim.opt.termguicolors = true

-- Load lazy.nvim plugin manager
require("jodify.lazy")

-- Load plugin specifications
-- require("jodify.plugins") -- Uncomment when plugins are added

-- Set up global keymaps
vim.g.mapleader = " "
vim.g.maplocalleader = "\\"

-- Print welcome message
vim.api.nvim_create_autocmd("VimEnter", {
	callback = function()
		print("Welcome to Jodify Neovim Configuration!")
	end,
})
