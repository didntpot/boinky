# boinky

A simple toolkit with gophertunnel tools

## Tools

boinky comes with some cool tools/exploits out of the box
* Auto Text - A tool to automatically send a sequential set of messages/commands to the server
* Bot Spammer - A tool to bot servers that have verification turned off
* Chat Reader - A tool (inspired by an old GT tool) that allows users to chat with a server from the terminal without logging into Minecraft normally
* Login Spoofer (WIP) - An exploit that allows users to log in as other users on servers that lack encryption via JWT data
* Music Player (WIP) - A tool that allows users to broadcast sound packets to servers that do not explicitly block them
* Skin Stealer - A tool that allows users to download all the skins of the players of the server that the user wants to

*More tools are WIP*

## Notes
- You can escape from a running tool by typing `!exit`
- You can configure all the tools in the `module/loader.go`. I plan to add a `config.json` later in the future to make this easier
- Run `go build main.go` after all the changes to get the latest executable or use `go run main.go`

## Other Plans
- Use [bubbletea](github.com/charmbracelet/bubbletea) to create a TUI instead of the manual typing

## Credits
- i skidded some code from [this](https://www.youtube.com/watch?v=h3Bv_rlMRGo) for the terminal color
- bot Spammer is partially skidded by my [father](https://github.com/Prim69)
- used my [mother's](https://github.com/JustTalDevelops) [NBS library](https://github.com/JustTalDevelops/nbs)
- i think i took the xbl.go from somwhere too but forgot lolz
<br/>

### cadet on top !
### star the repo and use the toolkit!


