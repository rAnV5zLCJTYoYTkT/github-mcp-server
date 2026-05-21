module github.com/github/github-mcp-server

go 1.23

require (
	github.com/google/go-github/v67 v67.0.0
	github.com/mark3labs/mcp-go v0.17.0
	github.com/shurcooL/githubv4 v0.0.0-20240727222349-48295856cce7
	github.com/spf13/cobra v1.8.1
	github.com/spf13/viper v1.19.0
	golang.org/x/oauth2 v0.24.0
)

require (
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/pelletier/go-toml/v2 v2.2.2 // indirect
	github.com/sagikazarmark/locafero v0.4.0 // indirect
	github.com/sagikazarmark/slog-shim v0.1.0 // indirect
	github.com/sourcegraph/conc v0.3.0 // indirect
	github.com/spf13/afero v1.11.0 // indirect
	github.com/spf13/cast v1.6.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.6.0 // indirect
	github.com/shurcooL/graphql v0.0.0-20230722043721-ed46e5a46466 // indirect
	golang.org/x/exp v0.0.0-20240213143201-ec583247a57a // indirect
	golang.org/x/sys v0.21.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

// Personal fork for learning MCP server patterns with the GitHub API.
// Upstream: https://github.com/github/github-mcp-server
//
// Notes:
//   - Studying how MCP tools map to GitHub REST vs GraphQL endpoints.
//   - go-github v67 uses per-resource service structs; worth noting for custom tool additions.
//   - GraphQL (shurcooL/githubv4) is used for queries that benefit from field selection,
//     e.g. fetching PR review threads without over-fetching REST payload fields.
//   - TODO: experiment with adding a custom MCP tool for listing stale branches.
//   - TODO: look into upgrading go-github to v68+ once mcp-go stabilises its tool interface;
//     v68 adds typed pagination helpers that would simplify list-based tool implementations.
