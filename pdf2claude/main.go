package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"os"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
)

func main() {
	client := anthropic.NewClient(
		option.WithAPIKey("API_KEY"),
		option.WithHeader("anthropic-beta", "pdfs-2024-09-25"),
	)

	content := "How many dogs are in the attached document?"

	println("[user]: " + content)

	file, err := os.Open("./dogs.pdf")
	if err != nil {
		panic(fmt.Errorf("failed to open file: %w", err))
	}

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	fileEncoded := base64.StdEncoding.EncodeToString(fileBytes)

	message, err := client.Beta.Messages.New(context.TODO(), anthropic.BetaMessageNewParams{
		MaxTokens: anthropic.Int(1024),
		Messages: anthropic.F([]anthropic.BetaMessageParam{{
			Role: anthropic.F(anthropic.BetaMessageParamRoleUser),
			Content: anthropic.F(
				[]anthropic.BetaContentBlockParamUnion{
					anthropic.BetaTextBlockParam{
						Text: anthropic.F(content),
						Type: anthropic.F(anthropic.BetaTextBlockParamTypeText),
					},
					anthropic.BetaBase64PDFBlockParam{
						Source: anthropic.F(anthropic.BetaBase64PDFSourceParam{
							Data:      anthropic.F(fileEncoded),
							MediaType: anthropic.F[anthropic.BetaBase64PDFSourceMediaType](anthropic.BetaBase64PDFSourceMediaTypeApplicationPDF),
							Type:      anthropic.F[anthropic.BetaBase64PDFSourceType]("base64"),
						}),
						Type: anthropic.F(anthropic.BetaBase64PDFBlockTypeDocument),
					},
				},
			),
		}}),
		Model: anthropic.F(anthropic.ModelClaude3_5Sonnet20241022),
	})
	if err != nil {
		panic(err)
	}

	println("[assistant]: " + message.Content[0].Text + message.StopSequence)
}
