package commands

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"ret/config"
	"ret/data"
	"ret/theme"
)

func Status(args []string) {
	if len(args) > 0 {
		switch args[0] {
		case "help":
			fmt.Fprintf(os.Stderr, theme.ColorGreen+"usage"+theme.ColorReset+": ret "+theme.ColorBlue+"status"+theme.ColorReset+"\n")
			fmt.Fprintf(os.Stderr, "  👀 displays the status for the current task with ret\n")
			fmt.Fprintf(os.Stderr, "  🔗 "+theme.ColorGray+"https://github.com/rerrorctf/ret/blob/main/commands/status.go"+theme.ColorReset+"\n")
			os.Exit(0)
		}
	}

	jsonData, err := os.ReadFile(config.RetFilesNames)
	if err != nil {
		os.Exit(1)
	}

	var files data.Files

	err = json.Unmarshal(jsonData, &files)
	if err == nil {
		for idx, file := range files.Files {

			fmt.Printf(theme.ColorGray+"["+theme.ColorBlue+"%v"+theme.ColorGray+"]"+theme.ColorReset, idx)
			fmt.Printf(theme.ColorGreen+" %s ", file.Filename)
			fmt.Printf(theme.ColorReset+"%s\n", file.SHA256)

			if file.FileType == data.FILE_TYPE_ELF {
				checksec := exec.Command("pwn", "checksec", file.Filepath)

				checksecOutput, err := checksec.CombinedOutput()
				if err != nil {
					fmt.Printf(theme.ColorGray+"    "+theme.ColorReset+"%s\n", file.FileOutput)
					continue
				}

				fmt.Printf("%s\n", checksecOutput)
			} else {
				fmt.Printf(theme.ColorGray+"    "+theme.ColorReset+"%s\n", file.FileOutput)
			}
		}
	}
}
