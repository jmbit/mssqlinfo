package inout

import (
	"bufio"
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
	"golang.org/x/term"
	"os"
	"strings"
	"syscall"
)

// Ask() gets the User to interactively fill out information that hasn't been provided via flags or environment variables
func Ask() {
	reader := bufio.NewReader(os.Stdin)
	if viper.GetString("server") == "" {
		fmt.Print("Enter server address: ")
		server, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("Error reading server address:", err)
		}
		viper.Set("server", strings.TrimSpace(server))
	}
	if viper.GetString("user") == "" {
		fmt.Print("Enter username: ")
		username, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("Error reading username:", err)
		}
		viper.Set("user", strings.TrimSpace(username))
	}
	if viper.GetString("password") == "" {
		fmt.Print("Enter Password: ")
		bytePassword, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			log.Fatal("Error reading password:", err)
		}

		viper.Set("password", strings.TrimSpace(string(bytePassword)))
	}

}
