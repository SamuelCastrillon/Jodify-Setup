-- lazy.nvim bootstrap and configuration
local lazypath = vim.fn.stdpath("data") .. "/lazy/lazy.nvim"

-- Check if lazy.nvim is not installed
if not vim.loop.fs_stat(lazypath) then
	-- Bootstrap lazy.nvim
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

-- Set up lazy.nvim with plugin specifications
require("lazy").setup({
	spec = {
		-- Plugin specifications will go here
		-- {
		--     "nvim-treesitter/nvim-treesitter",
		--     build = ":TSUpdate",
		-- },
	},
	defaults = {
		-- By default, lazy.nvim will lazy-load plugins that are not loaded yet
		-- Set to false to disable lazy-loading
		lazy = false,
		-- Install plugins matching these versions
		version = false,
		-- Don't automatically check for updates
		checker = false,
	},
	-- Configure lazy.nvim behavior
	install = {
		-- Install missing plugins on startup
		missing = true,
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
})
