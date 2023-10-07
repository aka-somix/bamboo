/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import "github.com/aka-somix/bamboo/cmd"

var Version = "development"

func main() {
	cmd.SetVersion(Version)
	cmd.Execute()
}
