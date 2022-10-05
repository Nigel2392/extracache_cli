package main

import (
	"ExtraClient/typeutils"
	"fmt"
	"os"
	"strings"

	extraconnector "github.com/Nigel2392/extraconnector"
	"github.com/TwiN/go-color"
)

var CONFIG *Config = &Config{}
var SERVER = &extraconnector.Server{}

var Meanings = map[string]map[string]string{
	"HELP": map[string]string{
		"message": "Shows this help menu",
		"syntax":  "HELP",
	},
	"EXIT": map[string]string{
		"message": "Exits the program",
		"syntax":  "EXIT",
	},
	"CLEAR": map[string]string{
		"message": "Clears the screen",
		"syntax":  "CLEAR",
	},
	"SET": map[string]string{
		"message": "Sets a value",
		"syntax":  "SET [key] [value] [ttl]",
	},
	"GET": map[string]string{
		"message": "Gets a value",
		"syntax":  "GET [key]",
	},
	"HASKEY": map[string]string{
		"message": "Checks if a key exists",
		"syntax":  "KEY [key]",
	},
	"DELETE": map[string]string{
		"message": "Deletes a key",
		"syntax":  "DELETE" + " [key]",
	},
	"SIZE": map[string]string{
		"message": "Gets the size of the database",
		"syntax":  "SIZE",
	},
	"SIZEALL": map[string]string{
		"message": "Gets the size of the database",
		"syntax":  "SIZEALL",
	},
	"KEYS": map[string]string{
		"message": "Gets all the keys in the database",
		"syntax":  "KEYS",
	},
	"SET_CHANNEL": map[string]string{
		"message": "Sets the channel",
		"syntax":  "SET_CHANNEL [CHANNEL]",
	},
}

func init() {
	CONFIG.SERVER = SERVER
	CONFIG.LoadConfig()
	go SERVER.Connect()
}

func SelectChoice() {
	for {
		ask := typeutils.Ask("Cmd -> ")
		arg_one := strings.Split(ask, " ")[0]
		arg_one = strings.ToUpper(arg_one)
		switch arg_one {
		case "EXIT":
			SERVER.Disconnect()
			os.Exit(0)
		case "HELP":
			PrintHelp()
		case "?":
			PrintHelp()
		case "CLEAR":
			Clear()
		case "CHANNEL":
			fmt.Println(color.Colorize(color.Purple, fmt.Sprintf("Your current channel is: %d", SERVER.Current_Channel)))
		case "SET":
			msg, err := handleSet(ask)
			if err != nil {
				cmd := Meanings[arg_one]
				fmt.Println(color.Colorize(color.Red, fmt.Sprintf("Error: %s\n", err.Error())))
				fmt.Println(color.Colorize(color.Blue, fmt.Sprintf("Syntax: %s\n", cmd["syntax"])))
			} else {
				received_from_s, err := SERVER.Send(&msg)
				if err != nil {
					continue
				}
				fmt.Println(received_from_s)
			}
		case "GET":
			msg, err := handleGet(ask)
			if err != nil {
				cmd := Meanings[arg_one]
				fmt.Println(color.Colorize(color.Red, fmt.Sprintf("Error: %s\n", err.Error())))
				fmt.Println(color.Colorize(color.Blue, fmt.Sprintf("Syntax: %s\n", cmd["syntax"])))
				fmt.Println(err)
			} else {
				received_from_s, err := SERVER.Send(&msg)
				if err != nil {
					continue
				}
				fmt.Println(received_from_s)
			}
		case "HASKEY":
			msg, err := handleHasKey(ask)
			if err != nil {
				cmd := Meanings[arg_one]
				fmt.Println(color.Colorize(color.Red, fmt.Sprintf("Error: %s\n", err.Error())))
				fmt.Println(color.Colorize(color.Blue, fmt.Sprintf("Syntax: %s\n", cmd["syntax"])))

				fmt.Println(err)
			} else {
				received_from_s, err := SERVER.Send(&msg)
				if err != nil {
					continue
				}
				fmt.Println(received_from_s)
			}
		case "DELETE":
			msg, err := handleDel(ask)
			if err != nil {
				cmd := Meanings[arg_one]
				fmt.Println(color.Colorize(color.Red, fmt.Sprintf("Error: %s\n", err.Error())))
				fmt.Println(color.Colorize(color.Blue, fmt.Sprintf("Syntax: %s\n", cmd["syntax"])))

				fmt.Println(err)
			} else {
				received_from_s, err := SERVER.Send(&msg)
				if err != nil {
					continue
				}
				fmt.Println(received_from_s)
			}
		case "SIZE":
			msg, err := handleSize(ask)
			if err != nil {
				cmd := Meanings[arg_one]
				fmt.Println(color.Colorize(color.Red, fmt.Sprintf("Error: %s\n", err.Error())))
				fmt.Println(color.Colorize(color.Blue, fmt.Sprintf("Syntax: %s\n", cmd["syntax"])))

				fmt.Println(err)
			} else {
				received_from_s, err := SERVER.Send(&msg)
				if err != nil {
					continue
				}
				fmt.Println(received_from_s)
			}
		case "SIZEALL":
			msg, err := handleSizeAll(ask)
			if err != nil {
				cmd := Meanings[arg_one]
				fmt.Println(color.Colorize(color.Red, fmt.Sprintf("Error: %s\n", err.Error())))
				fmt.Println(color.Colorize(color.Blue, fmt.Sprintf("Syntax: %s\n", cmd["syntax"])))

				fmt.Println(err)
			} else {
				received_from_s, err := SERVER.Send(&msg)
				if err != nil {
					continue
				}
				fmt.Println(received_from_s)
			}
		case "KEYS":
			msg, err := handleKeys(ask)
			if err != nil {
				cmd := Meanings[arg_one]
				fmt.Println(color.Colorize(color.Red, fmt.Sprintf("Error: %s\n", err.Error())))
				fmt.Println(color.Colorize(color.Blue, fmt.Sprintf("Syntax: %s\n", cmd["syntax"])))

				fmt.Println(err)
			} else {
				received_from_s, err := SERVER.Send(&msg)
				if err != nil {
					continue
				}
				fmt.Println(received_from_s)
			}
		case "SET_CHANNEL":
			err := handleSetChannel(ask)
			if err != nil {
				cmd := Meanings[arg_one]
				fmt.Println(color.Colorize(color.Red, fmt.Sprintf("Error: %s\n", err.Error())))
				fmt.Println(color.Colorize(color.Blue, fmt.Sprintf("Syntax: %s\n", cmd["syntax"])))

				fmt.Println(err)
			}
		default:
			fmt.Println(color.Colorize(color.Red, "Unknown command, type HELP for a list of commands"))
		}
	}
}

func main() {
	PrintLogo()
	SelectChoice()
}
func PrintHelp() {
	fmt.Println(color.Colorize(color.White, "Caching tool created in golang with a redis-like interface."))
	fmt.Println(color.Colorize(color.White, "Configure server IP and port in the config.json file, located in:"))
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(color.Colorize(color.White, path+"\\config.json"))
	fmt.Println()
	fmt.Println(color.Colorize(color.Green, "Available commands:"))
	for key, value := range Meanings {
		fmt.Println(color.Colorize(color.Red, key+":"))
		fmt.Println(color.Colorize(color.Purple, "("))
		fmt.Println(color.Colorize(color.Blue, "\tDescription: \t"), color.Colorize(color.Green, value["message"]))
		fmt.Println(color.Colorize(color.Blue, "\tSyntax: \t"), color.Colorize(color.Green, value["syntax"]))
		fmt.Println(color.Colorize(color.Purple, ")\n"))
	}
}

func Clear() {
	fmt.Print("\033[H\033[2J")
}

func PrintLogo() {
	// fmt.Println(color.Colorize(color.Purple, "#"+typeutils.Repeat("#", 70)))
	fmt.Println(color.Colorize(color.Purple, `
	
███████╗██╗  ██╗████████╗██████╗  █████╗    ██████╗ █████╗  ██████╗██╗  ██╗███████╗
██╔════╝╚██╗██╔╝╚══██╔══╝██╔══██╗██╔══██╗  ██╔════╝██╔══██╗██╔════╝██║  ██║██╔════╝
█████╗   ╚███╔╝    ██║   ██████╔╝███████║  ██║     ███████║██║     ███████║█████╗  
██╔══╝   ██╔██╗    ██║   ██╔══██╗██╔══██║  ██║     ██╔══██║██║     ██╔══██║██╔══╝  
███████╗██╔╝ ██╗   ██║   ██║  ██║██║  ██║  ╚██████╗██║  ██║╚██████╗██║  ██║███████╗
╚══════╝╚═╝  ╚═╝   ╚═╝   ╚═╝  ╚═╝╚═╝  ╚═╝   ╚═════╝╚═╝  ╚═╝ ╚═════╝╚═╝  ╚═╝╚══════╝

`+color.Colorize(color.Red, `Client: V 1.0.0`)+`
`+color.Colorize(color.Blue, `© Nigel van Keulen - ITExtra - 2022`)+`

	`))
	fmt.Println(color.Colorize(color.Purple, "#"+typeutils.Repeat("#", 70)))
}
