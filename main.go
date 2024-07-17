package main

import (
	"flag"
	"fmt"
	"strings"

	"ret/commands"
	"ret/config"
	"ret/theme"
)

var (
	COMMIT = ""
)

func main() {
	commands.PrepareCommands()

	flag.Usage = func() {
		fmt.Printf(theme.ColorGreen + "usage" + theme.ColorReset + ": ret " + theme.ColorPurple + "[help] " + theme.ColorBlue + "command" + theme.ColorGray + " [arg1 arg2...]\n" + theme.ColorReset)

		for _, cmd := range commands.Commands {
			fmt.Printf("%s ", cmd.Emoji)
			shortestValidPrefix, restOfCommand := commands.CommandsTrie.ShortestPrefix(cmd.Name)
			fmt.Printf(theme.ColorBlue+theme.StartUnderline+"%s"+theme.StopUnderline+"%s\n"+theme.ColorReset, shortestValidPrefix, restOfCommand)
		}

		fmt.Printf(theme.ColorGray+"https://github.com/rerrorctf/ret ~ %s\n"+theme.ColorReset, COMMIT)
	}

	flag.Parse()

	if flag.NArg() < 1 {
		flag.Usage()
		return
	}

	config.ParseUserConfig()

	command := flag.Arg(0)

	if command == "help" {
		if flag.NArg() < 2 {
			flag.Usage()
			return
		}

		found, cmd := commands.CommandsTrie.Search(flag.Arg(1))

		if !found {
			flag.Usage()
			return
		}

		fmt.Printf(theme.ColorGreen+"usage"+theme.ColorReset+": ret "+theme.ColorBlue+"%s "+theme.ColorReset, cmd.Name)

		for _, arg := range cmd.Arguments {
			name := arg.Name
			if arg.Default != "" {
				name = fmt.Sprintf("%s=%s", name, arg.Default)
			}

			if arg.List {

				if arg.Optional {
					fmt.Printf(theme.ColorGray+"[%s1 %s2 %s3...] ", name, name, name)
				} else {
					fmt.Printf(theme.ColorReset+"%s1 "+theme.ColorGray+"[%s2 %s3...] ", name, name, name)
				}
			} else {
				if arg.Optional {
					fmt.Printf(theme.ColorGray+"[%s] ", name)
				} else {
					fmt.Printf(theme.ColorReset+"%s ", name)
				}
			}
		}

		fmt.Printf("\n" + theme.ColorReset)

		help := cmd.Help()
		help = strings.ReplaceAll(help, "```\n", "")
		help = strings.ReplaceAll(help, "```bash\n", "")
		help = strings.ReplaceAll(help, "```python\n", "")
		help = strings.ReplaceAll(help, "`", "")
		fmt.Printf("%s %s", cmd.Emoji, help)
		return
	}

	found, cmd := commands.CommandsTrie.Search(command)

	if !found {
		flag.Usage()
		return
	}

	cmd.Func(flag.Args()[1:])
}
