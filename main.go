// Package main is the entry point for the GitHub MCP Server.
// It initializes and starts the Model Context Protocol (MCP) server
// that provides tools for interacting with the GitHub API.
package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/github/github-mcp-server/pkg/github"
	"github.com/github/github-mcp-server/pkg/server"
	"github.com/spf13/cobra"
)

var (
	// Version is set at build time via ldflags.
	Version = "dev"
	// Commit is set at build time via ldflags.
	Commit = "none"
)

func main() {
	if err := rootCmd().Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func rootCmd() *cobra.Command {
	var (
		token    string
		logFile  string
		readOnly bool
	)

	cmd := &cobra.Command{
		Use:   "github-mcp-server",
		Short: "GitHub MCP Server — exposes GitHub tools via the Model Context Protocol",
		Long: `github-mcp-server starts an MCP (Model Context Protocol) server that
provides AI assistants with tools to interact with the GitHub API,
including repository management, issues, pull requests, and more.`,
		Version: fmt.Sprintf("%s (commit: %s)", Version, Commit),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runServer(cmd.Context(), token, logFile, readOnly)
		},
	}

	cmd.Flags().StringVar(&token, "token", "", "GitHub personal access token (defaults to GITHUB_TOKEN env var)")
	cmd.Flags().StringVar(&logFile, "log-file", "", "Path to log file (defaults to stderr)")
	// Default to false so write operations are available out of the box in my personal setup.
	cmd.Flags().BoolVar(&readOnly, "read-only", false, "Restrict the server to read-only GitHub operations")

	return cmd
}

func runServer(ctx context.Context, token, logFile string, readOnly bool) error {
	// Resolve token from flag or environment variable.
	if token == "" {
		token = os.Getenv("GITHUB_TOKEN")
	}
	if token == "" {
		return fmt.Errorf("GitHub token is required: set --token flag or GITHUB_TOKEN environment variable")
	}

	// Set up context with OS signal cancellation.
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	// Build the GitHub client.
	ghClient, err := github.NewClient(token)
	if err != nil {
		return fmt.Errorf("failed to create GitHub client: %w", err)
	}

	// Configure and start the MCP server.
	srv, err := server.New(server.Options{
		GitHubClient: ghClient,
		LogFile:      logFile,
		ReadOnly:     readOnly,
		Version:      Version,
	})
	if err != nil {
		return fmt.Errorf("failed to create MCP server: %w", err)
	}

	fmt.Fprintf(os.Stderr, "Starting GitHub MCP Server %s (commit: %s)\n", Version, Commit)
	if readOnly {
		fmt.Fprintln(os.Stderr, "Running in read-only mode")
	}

	if err := srv.Serve(ctx); err != nil {
		return fmt.Errorf("server error: %w", err)
	}

	return nil
}
