/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/aka-somix/bamboo/pkg/template"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get [template name]",
	Short: "Retrieve the full info for a single template",
	Long: `Retrieve the full info for a single template`,

	
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			fmt.Printf("Error. Missing Template Name \n")
			os.Exit(1)
		}

		tm := template.NewTemplateManager()
		author := "default"
		name := args[0]

		template := tm.GetTemplateInfo(author, name)
		
		fmt.Println()
		fmt.Println("--------------------------------------------------------")
		fmt.Printf("Name: %s \nDescription: %s \nPath: %s \n", template.Name, template.Description, template.Path)
		fmt.Println("--------------------------------------------------------")
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
