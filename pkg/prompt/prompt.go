package prompt

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
)


func GetInput(pc PromptContent, validate promptui.ValidateFunc) string {

    templates := &promptui.PromptTemplates{
        Prompt:  "{{ . }} ",
        Valid:   "{{ . | green }} ",
        Invalid: "{{ . | red }} ",
        Success: "{{ . | bold }} ",
    }

    prompt := promptui.Prompt{
        Label:     pc.Label,
        Templates: templates,
        Validate:  validate,
    }

    result, err := prompt.Run()
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        os.Exit(1)
    }

    return result
}

func GetSelect(pc PromptContent, items []string) string {
 
	prompt := promptui.Select{
		Label:    pc.Label,
		Items:    items,
	}

	_, result, err := prompt.Run()

    if err != nil {
        fmt.Printf("Error: %v\n", err)
        os.Exit(1)
    }

    return result
}
