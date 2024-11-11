package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		log.Print("Registering handlers...")
		registerHandlers()
		return
	}

	if os.Args[1] == "toggle-gamemode" {
		log.Print("Toggling gamemode...")
		ToggleGamemode()
	}
}

func registerHandlers() {
	RegisterHandler("activespecial", func(s string) error {
		data := strings.Split(s, ",")
		if len(data) != 2 {
			panic(fmt.Sprintf("Invalid layout: %#v", data))
		}

		// Data{Workspace, Monitor}.
		if data[0] == "special:scratchpad" {
			if err := exec.Command("ags", "-m", "show bar").Run(); err != nil {
				return err
			}

			return nil
		}

		if err := exec.Command("ags", "-m", "hide bar").Run(); err != nil {
			return err
		}

		return nil
	})

	select {}
}
