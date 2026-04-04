-- Jodify Neovim Configuration
-- Navigation plugins setup: neo-tree, telescope, bufferline

-- ============================================
-- Neo-tree
-- ============================================

local neotree_ok, neotree = pcall(require, "neo-tree")
if not neotree_ok then
  vim.notify("neo-tree not installed yet", vim.log.levels.WARN)
  return
end

neotree.setup({
  sources = { "filesystem", "buffers", "git_status", "document_symbols" },
  open_files_behavior = "restore",
  close_if_last_window = true,
  default_component_configs = {
    indent = {
      with_markers = true,
      with_expanders = true,
      indent_size = 2,
    },
    icon = {
      folder_closed = "�_FOLDER",
      folder_open = "󰷉",
      folder_empty = "󰜌",
      default = "󰈙",
    },
    git_status = {
      symbols = {
        added = "󰐖",
        modified = "�.mod.",
        deleted = "󰍵",
        renamed = "󰁕",
        untracked = "󰋩",
        ignored = "󰀘",
        unstaged = "󰄱",
        staged = "󰱒",
        conflict = "󰿢",
      },
    },
  },
  window = {
    position = "left",
    width = 40,
    mapping_options = {
      noremap = true,
      nowait = true,
    },
    mappings = {
      ["<space>"] = {
        "toggle_node",
        nowait = false,
      },
      ["<2-LeftMouse>"] = {
        "open",
        nowait = false,
      },
      ["<cr>"] = {
        "open",
        nowait = false,
      },
      ["<esc>"] = "cancel",
      ["P"] = "toggle_preview",
      ["l"] = "focus_preview",
      ["S"] = "open_split",
      ["s"] = "open_vsplit",
      ["t"] = "open_tabnew",
      ["w"] = "open_with_window_picker",
      ["C"] = "close_node",
      ["z"] = "close_all_nodes",
      ["a"] = {
        "add",
        config = {
          show_path = "none",
        },
      },
      ["A"] = "add_directory",
      ["d"] = "delete",
      ["r"] = "rename",
      ["y"] = "copy_to_clipboard",
      ["x"] = "cut_to_clipboard",
      ["p"] = "paste_from_clipboard",
      ["m"] = "move_to_string",
      ["q"] = "close_window",
      ["R"] = "refresh",
      ["?"] = "show_help",
      ["<"] = "prev_source",
      [">"] = "next_source",
      ["i"] = "show_file_details",
    },
  },
  default_source = "filesystem",
  enable_diagnostics = true,
  enable_git_status = true,
  enable_modified_markers = true,
  enable_refresh_on_write = false,
  openBrowserOnStartup = false,
  hide_rootname = false,
  retain_in_hidden_max_buffer = true,
  event_handlers = {
    {
      event = "file_opened",
      handler = function()
        require("neo-tree").clear_selected_state()
      end,
    },
    {
      event = "vim_exit",
      handler = function()
        -- optional: do something on vim exit
      end,
    },
  },
  use_default_mappings = false,
  view = {
    width = 40,
    height = 30,
    side = "left",
    number = false,
    relativenumber = false,
    cursorline = true,
    float = {
      size = nil,
      position = "50%",
      padding = 10,
      border = "rounded",
    },
    adaptive_size = true,
    mappings = {
      n = {
        ["q"] = "close_window",
        ["<cr>"] = "set_root",
        ["."] = "toggle_hidden",
        ["/"] = "fuzzy_finder",
        ["D"] = "clear_filter",
        ["R"] = "refresh",
        ["h"] = "navigate_parent",
        ["-"] = "navigate_parent",
        ["P"] = "sticky_preview",
        ["l"] = "focus_preview",
      },
    },
  },
  filesystem = {
    bind_to_cwd = false,
    follow_current_file = {
      enabled = true,
      leave_dirs_open = false,
    },
    filtered_items = {
      visible = false,
      hide_hidden = true,
      hide_dotfiles = true,
      hide_gitignored = true,
      hide_by_name = {
        "node_modules",
        ".git",
        ".DS_Store",
        "thumbs.db",
      },
      never_show = {
        ".git",
        "node_modules",
      },
      always_show = {
        ".gitignored",
      },
    },
    find_by_path = {
      enable = true,
      search_code = true,
      search_forward = true,
      use_fuzzy = true,
    },
    group_empty_dirs = false,
    hijack_netrw_behavior = "open_default",
    use_libuv_file_watcher = true,
    window = {
      show_hidden = false,
    },
  },
  buffers = {
    follow_current_file = {
      enabled = true,
      leave_dirs_open = false,
    },
    group_empty_dirs = true,
    show_unloaded = true,
    window = {
      position = "current",
      width = 40,
      mapping_options = {
        noremap = true,
        nowait = true,
      },
      mappings = {
        ["bd"] = "buffer_delete",
        ["bs"] = "buffer_sort",
      },
    },
  },
  git_status = {
    window = {
      position = "float",
      width = 70,
      mapping_options = {
        noremap = true,
        nowait = true,
      },
    },
  },
})

-- ============================================
-- Telescope
-- ============================================

local telescope = require("telescope")
local actions = require("telescope.actions")

telescope.setup({
  defaults = {
    prompt_prefix = " 󰍉 ",
    selection_caret = " 󰄅 ",
    entry_prefix = "  ",
    initial_mode = "insert",
    selection_strategy = "reset",
    sorting_strategy = "ascending",
    layout_strategy = "horizontal",
    layout_config = {
      horizontal = {
        prompt_position = "top",
        preview_width = 0.55,
        results_width = 0.8,
      },
      vertical = {
        mirror = false,
      },
      width = 0.87,
      height = 0.80,
      preview_cutoff = 120,
    },
    file_sorter = require("telescope.sorters").get_fuzzy_file,
    file_ignore_patterns = {
      "node_modules/",
      ".git/",
      "__pycache__/",
      "%.jpg",
      "%.png",
      "%.jpeg",
      "%.bmp",
      "%.gif",
      "%.ico",
      "%.zip",
      "%.tar.gz",
      "%.rar",
      "%.7z",
    },
    generic_sorter = require("telescope.sorters").get_fuzzy_file,
    path_display = { "truncate" },
    winblend = 0,
    border = {},
    borderchars = { "─", "│", "─", "│", "╭", "╮", "╯", "╰" },
    color_devicons = true,
    set_env = { ["COLORTERM"] = "truecolor" },
    file_previewer_mappings = {
      init = "q",
      n = "j",
      Y = "y",
    },
    grep_previewer_mappings = {
      init = "q",
      n = "j",
    },
    file_ignore_patterns = {
      "node_modules",
      ".git",
    },
    mappings = {
      i = {
        ["<C-n>"] = actions.cycle_history_next,
        ["<C-p>"] = actions.cycle_history_prev,
        ["<C-j>"] = actions.move_selection_next,
        ["<C-k>"] = actions.move_selection_previous,
        ["<C-c>"] = actions.close,
        ["<Down>"] = actions.move_selection_next,
        ["<Up>"] = actions.move_selection_previous,
        ["<CR>"] = actions.select_default,
        ["<C-x>"] = actions.select_horizontal,
        ["<C-v>"] = actions.select_vertical,
        ["<C-t>"] = actions.select_tab,
        ["<C-u>"] = actions.preview_scrolling_up,
        ["<C-d>"] = actions.preview_scrolling_down,
        ["<PageUp>"] = actions.results_scrolling_up,
        ["<PageDown>"] = actions.results_scrolling_down,
      },
      n = {
        ["<esc>"] = actions.close,
        ["<CR>"] = actions.select_default,
        ["<C-x>"] = actions.select_horizontal,
        ["<C-v>"] = actions.select_vertical,
        ["<C-t>"] = actions.select_tab,
        ["j"] = actions.move_selection_next,
        ["k"] = actions.move_selection_previous,
        ["H"] = actions.move_to_top,
        ["M"] = actions.move_to_middle,
        ["L"] = actions.move_to_bottom,
        ["gg"] = actions.move_to_top,
        ["G"] = actions.move_to_bottom,
        ["<C-u>"] = actions.preview_scrolling_up,
        ["<C-d>"] = actions.preview_scrolling_down,
        ["<PageUp>"] = actions.results_scrolling_up,
        ["<PageDown>"] = actions.results_scrolling_down,
      },
    },
  },
  pickers = {
    find_files = {
      theme = "dropdown",
      previewer = false,
      hidden = true,
    },
    live_grep = {
      theme = "ivy",
    },
    buffers = {
      theme = "dropdown",
      previewer = false,
      initial_mode = "normal",
    },
  },
  extensions = {
    fzf = {
      fuzzy = true,
      override_generic_sorter = true,
      override_file_sorter = true,
      case_mode = "smart_case",
    },
    ["ui-select"] = {
      require("telescope.themes").get_dropdown({
        previewer = false,
        initial_mode = "normal",
      }),
    },
  },
})

-- Load fzf-native extension
pcall(require("telescope").load_extension, "fzf")
pcall(require("telescope").load_extension, "ui-select")

-- ============================================
-- Bufferline
-- ============================================

require("bufferline").setup({
  options = {
    close_command = "bdelete! %d",
    right_mouse_command = "bdelete! %d",
    diagnostics = "nvim_lsp",
    always_show_bufferline = false,
    diagnostics_indicator = function(_, _, diag)
      local icons = { error = "󰀦 ", warn = "󰀐 " }
      local ret = (diag.error and " " .. icons.error .. diag.error)
        or (diag.warning and " " .. icons.warn .. diag.warning)
        or ""
      return ret
    end,
    offsets = {
      {
        filetype = "neo-tree",
        text = "Neo-tree",
        highlight = "Directory",
        text_align = "left",
      },
    },
    separator_style = "thin",
    indicator = {
      style = "icon",
      icon = "▎",
    },
    buffer_close_icon = "󰅖",
    modified_icon = "●",
    close_icon = "󰅖",
    left_trunc_marker = "󰜎",
    right_trunc_marker = "󰜎",
    max_name_length = 18,
    max_prefix_length = 15,
    tab_size = 18,
    truncate_names = true,
    hover = {
      enabled = true,
      delay = 200,
      reveal = { "close" },
    },
    numbers = "none",
    restore_window_env = true,
    highlights = {
      fill = {
        fg = "#565f89",
        bg = "#1a1b26",
      },
      background = {
        fg = "#c0caf5",
        bg = "#1a1b26",
      },
      buffer = {
        fg = "#c0caf5",
        bg = "#1a1b26",
      },
      buffer_visible = {
        fg = "#c0caf5",
        bg = "#16161e",
      },
      buffer_selected = {
        fg = "#c0caf5",
        bg = "#292e42",
        bold = true,
        italic = true,
      },
      separator = {
        fg = "#1a1b26",
        bg = "#1a1b26",
      },
      separator_visible = {
        fg = "#1a1b26",
        bg = "#16161e",
      },
      separator_selected = {
        fg = "#1a1b26",
        bg = "#292e42",
      },
      indicator_selected = {
        fg = "#b388ff",
        bg = "#292e42",
      },
      indicator_visible = {
        fg = "#565f89",
        bg = "#16161e",
      },
      close_button = {
        fg = "#565f89",
        bg = "#1a1b26",
      },
      close_button_visible = {
        fg = "#565f89",
        bg = "#16161e",
      },
      close_button_selected = {
        fg = "#f7768e",
        bg = "#292e42",
      },
      tab = {
        fg = "#565f89",
        bg = "#16161e",
      },
      tab_selected = {
        fg = "#c0caf5",
        bg = "#292e42",
      },
      tab_separator = {
        fg = "#1a1b26",
        bg = "#16161e",
      },
      tab_separator_selected = {
        fg = "#1a1b26",
        bg = "#292e42",
      },
      diagnostic = {
        bg = "#1a1b26",
      },
      diagnostic_visible = {
        bg = "#16161e",
      },
      diagnostic_selected = {
        bg = "#292e42",
        bold = true,
        italic = true,
      },
    },
  },
})

-- Bufferline keymaps
local keymap = vim.keymap.set

keymap("n", "<leader>bp", ":BufferLineTogglePin<CR>", { noremap = true, silent = true })
keymap("n", "<leader>bP", ":BufferLineGroupClose ungrouped<CR>", { noremap = true, silent = true })
keymap("n", "<leader>bo", ":BufferLineCloseOthers<CR>", { noremap = true, silent = true })
keymap("n", "<leader>ba", ":BufferLineCloseLeft<CR>:BufferLineCloseRight<CR>", { noremap = true, silent = true })
keymap("n", "<leader>bh", ":BufferLineCyclePrev<CR>", { noremap = true, silent = true })
keymap("n", "<leader>bl", ":BufferLineCycleNext<CR>", { noremap = true, silent = true })
keymap("n", "<leader>bj", ":BufferLinePick<CR>", { noremap = true, silent = true })
keymap("n", "<leader>bc", ":BufferLinePickClose<CR>", { noremap = true, silent = true })
keymap("n", "<leader>bb", ":BufferLine<CR>", { noremap = true, silent = true })
