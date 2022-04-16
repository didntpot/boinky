package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/cqdetdev/boinky/data"
	"github.com/cqdetdev/boinky/module"
	"github.com/cqdetdev/boinky/util"
)

func main() {
	log := log.Default()
	l := module.DefaultLoader(log)

	if err := data.InitializeToken(log); err != nil {
		log.Fatal(err)
	}

	util.InitColor()
	
	fmt.Println("Type help for a list of commands.")

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		text := s.Text()
		if len(text) == 0 { continue }
		args := strings.Split(text, " ")
		cmd := args[0]
		if len(cmd) == 0 { continue }
		switch cmd {
			case "chat":
				if len(args) < 2 {
					log.Println("Please specify an <ip:port> to connect to.")
					continue
				}
				addr := args[1]
				l.ChatReader().Start(addr)
			case "bot":
				if len(args) < 2 {
					log.Println("Please specify an <ip:port> to connect to.")
					continue
				}
				addr := args[1]
				l.BotSpammer().Start(addr)
			case "help":
				fmt.Println("chat <ip:port> - A module to chat in a Minecraft server without logging into the server yourself\nbot <ip:port> - A module to bot MCPE servers (only if they have XUID verification disabled)")
				continue
			default:
				log.Printf("\"%s\" - Unknown command.\n", cmd)
				continue
		}
	}
	
}