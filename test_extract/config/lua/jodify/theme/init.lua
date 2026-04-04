-- Jodify Neovim Configuration
-- Theme configuration with Jodify color palette

-- TokyoNight theme with Jodify purple accents
require("tokyonight").setup({
  style = "night",           -- Dark style
  light_style = "day",       -- Light style for light mode
  transparent = false,       -- Not transparent
  terminal_colors = true,   -- Set terminal colors
  styles = {
    -- Styling for various syntax groups
    comments = { italic = true, fg = "#565f89" },
    keywords = { fg = "#bb9af7" },
    functions = { fg = "#7aa2f7" },
    variables = { fg = "#c0caf5" },
    constants = { fg = "#ff9e64" },
    numbers = { fg = "#ff9e64" },
    operators = { fg = "#bb9af7" },
    type = { fg = "#9ece6a" },
    strings = { fg = "#9ece6a" },
    parameters = { fg = "#c0caf5" },
    property = { fg = "#7dcfff" },
    builtin = { fg = "#bb9af7" },
    attribute = { fg = "#bb9af7" },
    field = { fg = "#c0caf5" },
    typeParameter = { fg = "#e0af68" },
    class = { fg = "#9ece6a" },
    label = { fg = "#7dcfff" },
    namespace = { fg = "#bb9af7" },
    tag = { fg = "#bb9af7" },
    module = { fg = "#7aa2f7" },
  },
  day_brightness = 0.3,
  hide_inactive_statusline = false,
  dim_inactive = false,
  lualine_bold = false,
  on_colors = function(colors)
    -- Customize TokyoNight colors with Jodify palette
    colors.hint = "#b388ff"
    colors.error = "#ff5353"
    colors.warning = "#e0af68"
    colors.info = "#7dcfff"

    -- Jodify specific colors
    colors.jodify_violet_500 = "#c380ff"  -- Brand color
    colors.jodify_violet_600 = "#b866ff"
    colors.jodify_violet_700 = "#ac4dff"
    colors.jodify_violet_900 = "#941aff"

    -- Override some TokyoNight colors
    colors.purple = "#b388ff"            -- Changed from default to Jodify violet
    colors.blue = "#7aa2f7"
    colors.cyan = "#7dcfff"
    colors.green = "#9ece6a"
    colors.yellow = "#e0af68"
    colors.orange = "#ff9e64"
    colors.red = "#f7768e"
    colors.fg = "#c0caf5"
    colors.fg_dim = "#565f89"
  end,
  on_highlights = function(hl, c)
    -- Customize specific highlights
    local git_signs = {
      add = "#9ece6a",
      change = "#7dcfff",
      delete = "#f7768e",
      topdelete = "#f7768e",
      changedelete = "#e0af68",
    }

    -- Cursor line
    hl.CursorLine = {
      bg = "#1a1b26",
      fg = "#c0caf5",
    }

    -- Visual selection
    hl.Visual = {
      bg = "#292e42",
      fg = "#c0caf5",
    }

    -- Search results
    hl.Search = {
      bg = "#b388ff",
      fg = "#1a1b26",
    }

    -- IncSearch
    hl.IncSearch = {
      bg = "#b388ff",
      fg = "#1a1b26",
    }

    -- MatchParen
    hl.MatchParen = {
      fg = "#b388ff",
      bold = true,
    }

    -- Line numbers
    hl.LineNr = {
      fg = "#565f89",
    }

    hl.CursorLineNr = {
      fg = "#b388ff",
      bold = true,
    }

    -- Pmenu
    hl.Pmenu = {
      bg = "#1a1b26",
      fg = "#c0caf5",
    }

    hl.PmenuSel = {
      bg = "#b388ff",
      fg = "#1a1b26",
    }

    hl.PmenuSbar = {
      bg = "#292e42",
    }

    hl.PmenuThumb = {
      bg = "#565f89",
    }

    -- Telescope
    hl.TelescopeSelection = {
      bg = "#292e42",
      fg = "#b388ff",
    }

    hl.TelescopeSelectionCaret = {
      fg = "#b388ff",
    }

    hl.TelescopeBorder = {
      fg = "#292e42",
    }

    hl.TelescopeNormal = {
      bg = "NONE",
      fg = "#c0caf5",
    }

    -- Neo-tree
    hl.NeoTreeDimText = {
      fg = "#565f89",
    }

    hl.NeoTreeNormal = {
      bg = "NONE",
      fg = "#c0caf5",
    }

    hl.NeoTreeNormalNC = {
      bg = "NONE",
      fg = "#c0caf5",
    }

    -- WhichKey
    hl.WhichKey = {
      fg = "#b388ff",
    }

    hl.WhichKeyGroup = {
      fg = "#7dcfff",
    }

    hl.WhichKeyDesc = {
      fg = "#c0caf5",
    }

    hl.WhichKeyFloat = {
      bg = "#1a1b26",
    }
  end,
})

-- Set the colorscheme
vim.cmd.colorscheme("tokyonight")

-- Also set some global highlight groups for custom colors
vim.api.nvim_set_hl(0, "JodifyViolet", { fg = "#c380ff", default = true })
vim.api.nvim_set_hl(0, "JodifyVioletBold", { fg = "#c380ff", bold = true, default = true })
vim.api.nvim_set_hl(0, "JodifyAccent", { fg = "#b388ff", default = true })
vim.api.nvim_set_hl(0, "JodifyBackground", { bg = "#110f18", default = true })
