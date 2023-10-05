/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/aka-somix/bamboo/pkg/prompt"
	"github.com/aka-somix/bamboo/pkg/template"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a new template",
	Long: `create a new template`,

	
	Run: func(cmd *cobra.Command, args []string) {
				// Prepare data
		tm := template.NewTemplateManager()
		author := "default"

		namePromptContent := prompt.PromptContent{
			ErrorMsg: "The name inserted is not valid.",
			Label: "Insert template name (must be unique):",
		}

		descPromptContent := prompt.PromptContent{
			ErrorMsg: "The description inserted is not valid.",
			Label: "Insert a short description (max 100 characters):",
		}

		pathPromptContent := prompt.PromptContent{
			ErrorMsg: "The name inserted is not valid.",
			Label: "Insert the folder path (either relative or absolute):",
		}

		name := prompt.GetInput(namePromptContent, prompt.ValidateName)
		desc := prompt.GetInput(descPromptContent, prompt.ValidateDescription)
		path := prompt.GetInput(pathPromptContent, prompt.ValidatePath)

		newTemplate := template.Template{
			Name: name,
			Description: desc,
			Author: author,
		}

		tm.CreateTemplate(newTemplate, path)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
