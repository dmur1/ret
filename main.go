package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"ret/commands"
	"ret/config"
	"ret/theme"
	"ret/util"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, theme.ColorGreen+"usage"+theme.ColorReset+": ret "+theme.ColorBlue+"command"+theme.ColorGray+" [arg1 arg2...]\n"+theme.ColorReset)

		fmt.Fprintf(os.Stderr, theme.ColorGray+
			"   core            rev            pwn            info            util\n"+
			"-------------------------------------------------------------------------\n"+
			theme.ColorBlue+
			"🚩 ctf          🦖 ghidra      🐚 pwn         🤝 abi         🤏 decompress\n"+
			"🔍 format       💃 ida         🐋 docker      📞 syscall     ✅ check\n"+
			"👀 status                      🗽 libc        📚 cheatsheet  📢 chat\n"+
			"📥 add                                                       🧠 gpt\n"+
			"🧙 wizard                                                    📝 writeup\n"+
			"                                                             🐙 gist\n"+theme.ColorReset)

		fmt.Fprintf(os.Stderr, "🚩 https://github.com/rerrorctf/ret 🚩\n")
	}

	flag.Parse()

	if flag.NArg() < 1 {
		flag.Usage()
		os.Exit(1)
	}

	config.ParseUserConfig()

	command := flag.Arg(0)

	// format
	if command[0] == 'f' {
		commands.Format(flag.Args()[1:])
		return
	}

	// pwn
	if command[0] == 'p' {
		commands.Pwn(flag.Args()[1:])
		return
	}

	// docker
	if command[0] == 'd' {
		if len(command) > 1 {
			if command[1] == 'o' {
				commands.Docker(flag.Args()[1:])
				return
			}

			if command[1] == 'e' {
				commands.Decompress(flag.Args()[1:])
				return
			}
		}
	}

	// ghidra, gist
	if command[0] == 'g' {
		if len(command) > 1 {
			if command[1] == 'h' {
				util.EnsureSkeleton()
				commands.Ghidra(flag.Args()[1:])
				return
			}

			if command[1] == 'i' {
				commands.Gist(flag.Args()[1:])
				return
			}

			if command[1] == 'p' {
				commands.Gpt(flag.Args()[1:])
				return
			}
		}
	}

	// ida
	if command[0] == 'i' {
		util.EnsureSkeleton()
		commands.Ida(flag.Args()[1:])
		return
	}

	// add, abi
	if command[0] == 'a' && len(command) > 1 {
		if command[1] == 'd' {
			util.EnsureSkeleton()
			commands.Add(flag.Args()[1:])
			return
		}

		if command[1] == 'b' {
			commands.Abi(flag.Args()[1:])
			return
		}
	}

	// status, syscall
	if command[0] == 's' && len(command) > 1 {
		if command[1] == 't' {
			commands.Status(flag.Args()[1:])
			return
		}

		if command[1] == 'y' {
			commands.Syscall(flag.Args()[1:])
			return
		}
	}

	// check, cheatsheet, ctf, chat
	if command[0] == 'c' {
		if len(command) > 1 {
			if command[1] == 't' {
				commands.Ctf(flag.Args()[1:])
				return
			}

			if command[1] == 'h' {
				if len(command) > 2 {
					if strings.Compare("cha", command[:3]) == 0 {
						commands.Chat(flag.Args()[1:])
						return
					}
				}
				if len(command) > 3 {
					if strings.Compare("chec", command[:4]) == 0 {
						commands.Check(flag.Args()[1:])
						return
					}

					if strings.Compare("chea", command[:4]) == 0 {
						commands.Cheatsheet(flag.Args()[1:])
						return
					}
				}
			}
		}
	}

	// writeup, wizard
	if command[0] == 'w' && len(command) > 1 {
		if command[1] == 'r' {
			commands.Writeup(flag.Args()[1:])
			return
		}

		if command[1] == 'i' {
			commands.Wizard(flag.Args()[1:])
			return
		}
	}

	// libc
	if command[0] == 'l' {
		commands.Libc(flag.Args()[1:])
		return
	}

	// help
	flag.Usage()
	os.Exit(1)
}
