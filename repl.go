package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
)

var cliName string = "Pokedex"

func startRepl() {
  reader := bufio.NewScanner(os.Stdin)
  printPrompt()
  commands := getCommands()

  for reader.Scan() {
    words := cleanInput(reader.Text())

    if len(words) == 0 {
      continue
    }

    commandName := words[0]

    command, exists := commands[commandName]
    if exists {
      err := command.callback()
      if err != nil {
        fmt.Println(err)
      }
      continue
    } else {
      handleCmd(commandName)
      continue
    }
  }
}

func printPrompt() {
  fmt.Print(cliName, "> ")
}

func ptintUnknown(text string) {
  fmt.Println(text, ": command not found")
}

func handleInvalidCmd(text string) {
  defer ptintUnknown(text)
}

func handleCmd(text string) {
  handleInvalidCmd(text)
}

func cleanInput(text string) []string {
  output := strings.ToLower(text)
  words := strings.Fields(output)
  return words
}

type cliCommand struct {
  name        string
  description string
  callback    func() error
}

func getCommands() map[string]cliCommand {
  return map[string]cliCommand {
    "help": {
      name:         "help",
      description:  "Display a help message",
      callback:     commandHelp,
   },
    "exit": {
      name:         "exit",
      description:  "Exit the Pokedex",
      callback:     commandExit,
    },
  }
}


