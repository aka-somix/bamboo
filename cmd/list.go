/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/aka-somix/bamboo/pkg/template"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all templates for the user",
	Long: `list all templates for the user`,


	Run: func(cmd *cobra.Command, args []string) {

		tm := template.NewTemplateManager()
		author := "default"

		templates := tm.ListTemplatesInfo(author)
		
		// Print Header
		fmt.Println()
		fmt.Println("+-----------------------------------------------------")
		fmt.Printf("| List of templates for %s \n", author)
		fmt.Println("+-----------------------------------------------------")


		// Print Body
		for _,template := range templates {
			fmt.Printf("|\t🎋 %s | Created by %s \n", template.Name, template.Author)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
