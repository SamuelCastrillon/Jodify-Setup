-- Jodify Neovim Configuration
-- LSP and Treesitter setup

-- ============================================
-- Treesitter (rama main - nueva API)
-- ============================================

require("nvim-treesitter").setup({
  ensure_installed = {
    "lua",
    "vim",
    "vimdoc",
    "python",
    "javascript",
    "typescript",
    "tsx",
    "jsx",
    "json",
    "yaml",
    "markdown",
    "markdown_inline",
    "bash",
    "go",
    "rust",
    "c",
    "cpp",
    "html",
    "css",
  },
  sync_install = false,
  auto_install = true,
  highlight = {
    enable = true,
    additional_vim_regex_highlighting = false,
  },
  indent = {
    enable = true,
  },
  folding = {
    enable = true,
  },
})

-- Treesitter textobjects (plugin separado en rama main)
require("nvim-treesitter-textobjects").setup({
  textobjects = {
    select = {
      enable = true,
      lookahead = true,
      keymaps = {
        ["af"] = "@function.outer",
        ["if"] = "@function.inner",
        ["ac"] = "@class.outer",
        ["ic"] = "@class.inner",
      },
    },
    move = {
      enable = true,
      set_jumps = true,
      goto_next_start = {
        ["]f"] = "@function.outer",
        ["]c"] = "@class.outer",
      },
      goto_previous_start = {
        ["[f"] = "@function.outer",
        ["[c"] = "@class.outer",
      },
    },
  },
})

-- ============================================
-- Mason
-- ============================================

require("mason").setup({
  ui = {
    border = "rounded",
    height = 0.8,
    width = 0.8,
    icons = {
      package_installed = "✓",
      package_pending = "➜",
      package_uninstalled = "✗",
    },
  },
  log_level = vim.log.levels.INFO,
  max_concurrent_installers = 4,
  pip = {
    upgrade_pip = false,
    install_args = {},
  },
})

require("mason-lspconfig").setup({
  ensure_installed = {
    "lua_ls",
    "pyright",
    "typescript-language-server",
    "gopls",
    "rust_analyzer",
    "clangd",
    "html",
    "cssls",
    "jsonls",
    "yamlls",
    "marksman",
    "bashls",
    "dockerls",
    "tailwindcss",
  },
  automatic_installation = true,
})

-- ============================================
-- LSP Config
-- ============================================

-- Diagnostic configuration
vim.diagnostic.config({
  virtual_text = {
    prefix = "●",
    source = "if_many",
    format = function(diagnostic)
      local severity = diagnostic.severity
      local icons = {
        [vim.diagnostic.severity.ERROR] = "󰀦 ",
        [vim.diagnostic.severity.WARN] = "󰀐 ",
        [vim.diagnostic.severity.HINT] = "󰌵 ",
        [vim.diagnostic.severity.INFO] = "󰋽 ",
      }
      return string.format("%s%s", icons[severity], diagnostic.message)
    end,
  },
  signs = true,
  underline = true,
  update_in_insert = false,
  severity_sort = true,
  float = {
    border = "rounded",
    source = "always",
    prefix = " ",
    focus = false,
  },
})

-- LSP server configurations
local lspconfig = require("lspconfig")

local servers = {
  lua_ls = {
    settings = {
      Lua = {
        runtime = {
          version = "LuaJIT",
        },
        diagnostics = {
          globals = { "vim" },
        },
        workspace = {
          library = vim.api.nvim_get_runtime_file("", true),
          checkThirdParty = false,
        },
        telemetry = {
          enable = false,
        },
        format = {
          enable = true,
          defaultConfig = {
            indent_style = "space",
            indent_size = "2",
          },
        },
      },
    },
  },
  pyright = {
    settings = {
      python = {
        analysis = {
          autoSearchPaths = true,
          diagnosticMode = "workspace",
          useLibraryCodeForTypes = true,
          typeCheckingMode = "basic",
        },
      },
    },
  },
  tsserver = {
    settings = {
      typescript = {
        inlayHints = {
          includeInlayParameterNameHintsWhenArgumentMatchesSignature = false,
          includeInlayParameterNameHints = "all",
          includeInlayFunctionParameterTypeHints = true,
          includeInlayVariableTypeHints = true,
          includeInlayPropertyDeclarationTypeHints = true,
          includeInlayFunctionLikeReturnTypeHints = true,
          includeInlayEnumMemberValueHints = true,
        },
      },
      javascript = {
        inlayHints = {
          includeInlayParameterNameHintsWhenArgumentMatchesSignature = false,
          includeInlayParameterNameHints = "all",
          includeInlayFunctionParameterTypeHints = true,
          includeInlayVariableTypeHints = true,
          includeInlayPropertyDeclarationTypeHints = true,
          includeInlayFunctionLikeReturnTypeHints = true,
          includeInlayEnumMemberValueHints = true,
        },
      },
    },
  },
  gopls = {
    settings = {
      gopls = {
        analyses = {
          nilness = true,
          unusedparams = true,
          unreachable = false,
        },
        gofumpt = true,
        staticcheck = true,
      },
    },
  },
  rust_analyzer = {
    settings = {
      ["rust-analyzer"] = {
        cargo = {
          allFeatures = true,
        },
        checkOnSave = {
          command = "clippy",
        },
        procMacro = {
          enable = true,
        },
      },
    },
  },
}

-- Setup LSP servers
for server, config in pairs(servers) do
  local opts = vim.tbl_deep_extend("force", {
    capabilities = require("cmp_nvim_lsp").default_capabilities(),
  }, config)
  lspconfig[server].setup(opts)
end

-- LSP keymaps (already in keymaps/init.lua, but adding here for reference)
-- gd - go to definition
-- gD - go to declaration
-- gi - go to implementation
-- gr - go to references
-- K - hover
-- <leader>rn - rename
-- <leader>ca - code action
-- <leader>e - show diagnostic
-- [d / ]d - prev/next diagnostic

-- ============================================
-- nvim-cmp (completion)
-- ============================================

local cmp = require("cmp")
local luasnip = require("luasnip")

require("luasnip.loaders.from_vscode").lazy_load()

local cmp_mapping = require("cmp.config.mapping")

cmp.setup({
  mapping = cmp_mapping.preset.insert({
    ["<C-n>"] = cmp_mapping.select_next_item({ behavior = cmp.SelectBehavior.Insert }),
    ["<C-p>"] = cmp_mapping.select_prev_item({ behavior = cmp.SelectBehavior.Insert }),
    ["<C-d>"] = cmp_mapping.scroll_docs(-4),
    ["<C-f>"] = cmp_mapping.scroll_docs(4),
    ["<C-Space>"] = cmp_mapping.complete(),
    ["<C-e>"] = cmp_mapping.abort(),
    ["<CR>"] = cmp_mapping.confirm({ select = true }),
    ["<S-CR>"] = cmp_mapping.confirm({
      behavior = cmp.ConfirmBehavior.Replace,
      select = true,
    }),
    ["<Tab>"] = cmp_mapping(function(fallback)
      if cmp.visible() then
        cmp.select_next_item()
      elseif luasnip.expand_or_jumpable() then
        luasnip.expand_or_jump()
      else
        fallback()
      end
    end, {
      "i",
      "s",
    }),
    ["<S-Tab>"] = cmp_mapping(function(fallback)
      if cmp.visible() then
        cmp.select_prev_item()
      elseif luasnip.jumpable(-1) then
        luasnip.jump(-1)
      else
        fallback()
      end
    end, {
      "i",
      "s",
    }),
  }),
  sources = cmp.config.sources({
    { name = "nvim_lsp" },
    { name = "luasnip" },
    { name = "path" },
  }, {
    { name = "buffer" },
  }),
  snippet = {
    expand = function(args)
      luasnip.lsp_expand(args.body)
    end,
  },
  formatting = {
    format = function(_, item)
      local icons = require("jodify.theme").icons or {}
      if icons[item.kind] then
        item.kind = icons[item.kind] .. item.kind
      end
      return item
    end,
  },
  window = {
    completion = cmp.config.window.bordered(),
    documentation = cmp.config.window.bordered(),
  },
  experimental = {
    ghost_text = true,
  },
})

-- Cmdline setup
cmp.setup.cmdline(":", {
  mapping = cmp_mapping.preset.cmdline(),
  sources = cmp.config.sources({
    { name = "path" },
  }, {
    { name = "cmdline" },
  }),
})

cmp.setup.cmdline("/", {
  mapping = cmp_mapping.preset.cmdline(),
  sources = {
    { name = "buffer" },
  },
})

-- ============================================
-- Conform (formatter)
-- ============================================

require("conform").setup({
  formatters_by_ft = {
    lua = { "stylua" },
    python = { "black", "isort" },
    javascript = { "prettier" },
    typescript = { "prettier" },
    javascriptreact = { "prettier" },
    typescriptreact = { "prettier" },
    json = { "prettier" },
    yaml = { "prettier" },
    markdown = { "prettier" },
    html = { "prettier" },
    css = { "prettier" },
    go = { "gofmt", "goimports" },
    rust = { "rustfmt" },
    c = { "clangformat" },
    cpp = { "clangformat" },
  },
  format_on_save = {
    timeout_ms = 500,
    lsp_fallback = true,
  },
})

-- Autoformat on save
vim.api.nvim_create_autocmd("BufWritePre", {
  pattern = "*",
  callback = function(args)
    require("conform").format({
      bufnr = args.buf,
      lsp_fallback = true,
      quiet = true,
    })
  end,
})

-- ============================================
-- nvim-autopairs
-- ============================================

require("nvim-autopairs").setup({
  check_ts = true,
  ts_config = {
    lua = { "string" },
    javascript = { "template_string" },
    java = false,
  },
})

local cmp_autopairs = require("nvim-autopairs.completion.cmp")
cmp.event:on("confirm_done", cmp_autopairs.on_confirm_done())

-- ============================================
-- indent-blankline
-- ============================================

require("ibl").setup({
  indent = {
    char = "│",
    highlight = "IndentBlanklineChar",
  },
  scope = {
    enabled = true,
    char = "│",
    highlight = "IndentBlanklineContextChar",
  },
  exclude = {
    filetypes = {
      "help",
      "dashboard",
      "lazy",
      "mason",
      "notify",
      "term",
      "terminal",
      "toggleterm",
      "alpha",
    },
  },
})
