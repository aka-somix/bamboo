package prompt

import (
	"errors"
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
)


func GetInput(pc PromptContent) string {

	// input validation
    validate := func(input string) error {
        if len(input) <= 0 {
            return errors.New(pc.ErrorMsg)
        }
        return nil
    }

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
        fmt.Printf("Prompt failed %v\n", err)
        os.Exit(1)
    }

    fmt.Printf("Input: %s\n", result)

    return result
}

func GetSelect(pc PromptContent, items []string) string {
 
	prompt := promptui.Select{
		Label:    pc.Label,
		Items:    items,
	}

	_, result, err := prompt.Run()

    if err != nil {
        fmt.Printf("Prompt failed %v\n", err)
        os.Exit(1)
    }
    
    return result
}
