package commands

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"ret/config"
	"ret/data"
	"ret/theme"
	"ret/util"
)

func filesAlreadyExists() bool {
	_, err := os.Stat(config.RetFilesNames)
	return !os.IsNotExist(err)
}

func parseFiles(files *data.Files) {
	if filesAlreadyExists() {
		jsonData, err := os.ReadFile(config.RetFilesNames)
		if err != nil {
			fmt.Println("error reading:", config.RetFilesNames)
			os.Exit(1)
		}

		err = json.Unmarshal(jsonData, &files)
		if err != nil {
			fmt.Println("error unmarshalling json:", err)
			os.Exit(1)
		}
	}
}

func writeFiles(files *data.Files) {
	jsonData, err := json.MarshalIndent(files, "", "  ")
	if err != nil {
		fmt.Println("error marshalling json:", err)
		os.Exit(1)
	}

	err = os.WriteFile(config.RetFilesNames, jsonData, 0644)
	if err != nil {
		fmt.Println("error writing to file:", err)
		os.Exit(1)
	}
}

func addFile(srcPath string) {
	files := data.Files{}
	parseFiles(&files)

	_, fileName := filepath.Split(srcPath)

	fileOutput := util.RunFileCommandOnFile(srcPath)

	content, err := os.ReadFile(srcPath)
	if err != nil {
		fmt.Println("error reading file:", srcPath)
		return
	}

	fileType := data.FILE_TYPE_UNKNOWN

	for magicType, magic := range data.FileMagics {
		match := true
		for j := range magic {
			if content[j] != magic[j] {
				match = false
				break
			}
		}

		if !match {
			continue
		}

		fileType = magicType
		break
	}

	sha256Hash := sha256.New()
	sha256Hash.Write(content)
	sha256HashString := hex.EncodeToString(sha256Hash.Sum(nil))

	dirPath := config.FilesFolderName + "/" + sha256HashString
	dstPath := dirPath + "/" + fileName

	file := data.File{
		Filename:   fileName,
		Filepath:   dstPath,
		Size:       len(content),
		FileType:   fileType,
		FileOutput: fileOutput,
		SHA256:     sha256HashString,
	}

	if _, err := os.Stat(dirPath); !os.IsNotExist(err) {
		fmt.Printf("💥 "+theme.ColorRed+"error"+theme.ColorReset+": file \"%s\" with sha256 \"%s\" already added...\n",
			srcPath, sha256HashString)
		return
	}

	err = os.MkdirAll(dirPath, 0755)
	if err != nil {
		fmt.Println("error making directory:", dirPath)
		return
	}

	err = util.CopyFile(srcPath, dstPath)
	if err != nil {
		fmt.Println("error copying file:", dstPath)
		return
	}

	fmt.Printf("📥 adding \"%s\" %s\n", srcPath, sha256HashString)

	files.Files = append(files.Files, file)

	writeFiles(&files)

	util.ProcessFile(dstPath)
}

func AddHelp() {
	fmt.Fprintf(os.Stderr, theme.ColorGreen+"usage"+theme.ColorReset+": ret "+theme.ColorBlue+"add"+theme.ColorReset+" file1 "+theme.ColorGray+"[file2 file3...]"+theme.ColorReset+"\n")
	fmt.Fprintf(os.Stderr, "  📥 add one or more files to the current task with ret\n")
	fmt.Fprintf(os.Stderr, "  🔗 "+theme.ColorGray+"https://github.com/rerrorctf/ret/blob/main/commands/add.go"+theme.ColorReset+"\n")
	os.Exit(0)
}

func Add(args []string) {
	if len(args) > 0 {
		switch args[0] {
		case "help":
			AddHelp()
			os.Exit(0)
		default:
			for _, file := range args {
				addFile(file)
			}
		}
	} else {
		AddHelp()
		os.Exit(1)
	}
}
