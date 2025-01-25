package helper

import (
	"bufio"
	"context"
	"fmt"
	"log"
)

var CurrentUsername = ""

func HandleUserSession(ctx context.Context, scanner *bufio.Scanner) bool {
	if SessionExists() {
		currentUser, err := RestoreSession(ctx)
		if err != nil {
			log.Println("Invalid credential")
			return false
		}
		CurrentUsername = currentUser.Username
		fmt.Printf("Welcome back, %s!\n", CurrentUsername)
		return true
	}

	for {
		token, err := AuthenticateSubMenu(ctx, scanner, &CurrentUsername)
		if err != nil {
			log.Printf("Invalid Credential.. %v", err)
			continue
		}
		SaveSession(token)
		return true
	}
}

func MainMenuLoop(ctx context.Context, scanner *bufio.Scanner) {
	hasJoined := false
	for {
		MainMenu(ctx, scanner, &hasJoined, CurrentUsername)
	}
}
