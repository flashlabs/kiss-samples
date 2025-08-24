package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	// Create MCP server for Wikipedia search
	s := server.NewMCPServer(
		"Wikipedia Search",
		"v1.0.0",
		server.WithToolCapabilities(false),
		server.WithRecovery(),
	)

	// Define Wikipedia search tool
	wikiTool := mcp.NewTool(
		"wikipedia_search",
		mcp.WithDescription("Search Wikipedia for a query and return the first paragraph"),
		mcp.WithString("query", mcp.Required(), mcp.Description("Search term")),
	)

	// Register tool handler
	s.AddTool(wikiTool, func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		term, err := req.RequireString("query")
		if err != nil {
			return nil, fmt.Errorf("failed to get query: %w", err)
		}

		// Build Wikipedia API URL
		u := fmt.Sprintf("https://en.wikipedia.org/api/rest_v1/page/summary/%s", url.PathEscape(term))

		// Set request timeout
		ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()

		// Create HTTP request
		httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create request: %w", err)
		}
		httpReq.Header.Set("User-Agent", "Wikipedia-MCP-Server/1.0")

		r, err := http.DefaultClient.Do(httpReq)
		if err != nil {
			return nil, fmt.Errorf("failed to get wikipedia page: %w", err)
		}
		defer func() {
			if closeErr := r.Body.Close(); closeErr != nil {
				fmt.Printf("Warning: failed to close response body: %v\n", closeErr)
			}
		}()

		if r.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("failed to get wikipedia page: %s", r.Status)
		}

		// Read response body (limit to 1MB)
		body, err := io.ReadAll(io.LimitReader(r.Body, 1024*1024))
		if err != nil {
			return nil, fmt.Errorf("failed to read response body: %w", err)
		}

		// Parse Wikipedia summary
		var summary struct {
			Extract string `json:"extract"`
		}
		if err := json.Unmarshal(body, &summary); err != nil {
			return nil, fmt.Errorf("failed to parse response: %w", err)
		}

		if summary.Extract == "" {
			return nil, fmt.Errorf("no summary found for '%s'", term)
		}

		return mcp.NewToolResultText(summary.Extract), nil
	})

	// Start MCP server on stdio
	fmt.Println("MCP server started")
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}

	fmt.Println("MCP server stopped")
}
