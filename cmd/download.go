/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/aka-somix/bamboo/pkg/prompt"
	"github.com/aka-somix/bamboo/pkg/template"
	"github.com/spf13/cobra"
)

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "download one of the templates available from origin",
	Long: `download one of the templates available from origin`,


	Run: func(cmd *cobra.Command, args []string) {
		// Prepare data
		tm := template.NewTemplateManager()
		author := "default"

		templates := tm.ListTemplatesInfo(author)

		var names []string
		
		for _, t := range templates {
			names = append(names, t.Name)
		}

		// Select which template to download
		selectPromptContent := prompt.PromptContent{
			ErrorMsg: "Please select a template.",
			Label: "Found these templates. Select which one to download:",
		}
		selectedName := prompt.GetSelect(selectPromptContent, names)

		folderPromptContent := prompt.PromptContent{
			ErrorMsg: "Please provide a valid folder name",
			Label: fmt.Sprintf("Download path (%s)", selectedName),
		}

		folderPath := prompt.GetInput(folderPromptContent, prompt.ValidatePath)

		if len(folderPath) == 0 {
			folderPath = selectedName
		}

		tm.DownloadTemplate(author, selectedName, folderPath)
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// downloadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// downloadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
