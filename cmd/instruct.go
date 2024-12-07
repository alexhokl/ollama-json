package cmd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/alexhokl/helper/iohelper"
	"github.com/ollama/ollama/api"
	"github.com/spf13/cobra"
)

type instructOptions struct {
	modelName string
	path      string
	question  string
}

var instructOpts instructOptions

// instructCmd represents the instruct command
var instructCmd = &cobra.Command{
	Use:   "instruct",
	Short: "Instruct the model to give an answer",
	RunE:  runInstruct,
}

func init() {
	rootCmd.AddCommand(instructCmd)

	flags := instructCmd.Flags()
	flags.StringVarP(&instructOpts.modelName, "model", "m", "llama3.1:8b", "Model to use")
	flags.StringVarP(&instructOpts.path, "file", "f", "", "Path to json structure file")
	flags.StringVarP(&instructOpts.question, "question", "q", "", "Question to instruct")

	instructCmd.MarkFlagRequired("file")
	instructCmd.MarkFlagRequired("question")
}

func runInstruct(cmd *cobra.Command, args []string) error {
	jsonStructureBytes, err := iohelper.ReadStringFromFile(instructOpts.path)
	if err != nil {
		return err
	}

	client, err := api.ClientFromEnvironment()
	if err != nil {
		return err
	}

	req := &api.GenerateRequest{
		Model:  instructOpts.modelName,
		Prompt: instructOpts.question,
		Format: json.RawMessage(jsonStructureBytes),
	}

	ctx := context.Background()
	respFunc := func(resp api.GenerateResponse) error {
		// In streaming mode, responses are partial so we call fmt.Print (and not
		// Println) in order to avoid spurious newlines being introduced. The
		// model will insert its own newlines if it wants.
		fmt.Print(resp.Response)
		return nil
	}

	err = client.Generate(ctx, req, respFunc)
	if err != nil {
		return err
	}
	fmt.Println()

	return nil
}
