package commands

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"ret/config"
	"ret/theme"
	"time"
)

func ghidraSpinner() {
	emojis := []string{
		"🍎", "🥑", "🥓", "🥖", "🍌", "🥯", "🫐", "🍔", "🥦", "🥩",
		"🥕", "🥂", "🍫", "🍪", "🥒", "🧀", "🥚", "🍳", "🍟", "🍇",
		"🍏", "🍔", "🍯", "🥝", "🍋", "🥬", "🍞", "🥗", "🍣", "🍜",
		"🥟", "🍲", "🌭", "🍕", "🍝", "🌮", "🍉", "🍊", "🍓", "🚩",
	}

	for {
		for _, e := range emojis {
			fmt.Printf("\r%s -> 🦖", e)
			time.Sleep(200 * time.Millisecond)
		}
	}
}

func Ghidra(args []string) {
	if len(args) > 0 {
		switch args[0] {
		case "help":
			fmt.Fprintf(os.Stderr, theme.ColorGreen+"usage"+theme.ColorReset+": ret "+theme.ColorBlue+"ghidra"+theme.ColorGray+" [file file...]"+theme.ColorReset+"\n")
			fmt.Fprintf(os.Stderr, "  🦖 ingests all added files then opens ghidra with ret\n")
			os.Exit(0)
		}
	}

	if _, err := os.Stat(config.GhidraProjectPath); os.IsNotExist(err) {
		err := os.MkdirAll(config.GhidraProjectPath, 0755)
		if err != nil {
			fmt.Println("error creating directory:", err)
			os.Exit(1)
		}
	}

	absoluteProjectPath, err := filepath.Abs(config.GhidraProjectPath + "/project.gpr")
	if err != nil {
		fmt.Println("error abs:", err)
		os.Exit(1)
	}

	if len(args) > 0 {
		Add(args)
	}

	go ghidraSpinner()

	analyzeFile := exec.Command(
		config.GhidraInstallPath+"/support/analyzeHeadless",
		config.GhidraProjectPath,
		"project", "-recursive",
		"-import", config.FilesFolderName)

	analyzeFileOutput, err := analyzeFile.CombinedOutput()
	if err != nil {
		fmt.Printf("%s\n", analyzeFileOutput)
		fmt.Println("warning:\n", err)
	}

	openGhidra := exec.Command(
		config.GhidraInstallPath+"/ghidraRun", absoluteProjectPath)

	openGhidraOutput, err := openGhidra.CombinedOutput()
	if err != nil {
		fmt.Printf("%s\n", openGhidraOutput)
		fmt.Println("warning:\n", err)
	}

	fmt.Printf("\r")
}
