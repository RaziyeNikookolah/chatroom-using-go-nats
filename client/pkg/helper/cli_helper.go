package helper

import (
	"bufio"
	"context"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/RaziyeNikookolah/chatroom-using-go-nats/client/domain"
	"github.com/RaziyeNikookolah/chatroom-using-go-nats/client/pkg/nats"
	"github.com/gookit/color"
)

func MainMenu(ctx context.Context, scanner *bufio.Scanner, hasJoined *bool, username string) {
	color.Cyan.Print("\n>> ") // Prompt with color
	if !scanner.Scan() {
		return
	}
	input := strings.TrimSpace(scanner.Text())
	// if input == "" {
	// 	continue
	// }
	switch input {
	case "help":
		PrintHelp()
	case "join":
		JoinChatroom(ctx, hasJoined, username)
	case "send":
		if CheckAccess(hasJoined) {
			SendMessage(ctx, username, scanner)
		}
	case "view":
		if CheckAccess(hasJoined) {
			ViewSubMenu(scanner)
		}
	case "exit":
		color.Green.Println("\n👋 Goodbye! Thanks for chatting!")
		os.Exit(0)
		return
	default:
		color.Red.Println("❌ Unknown command. Type 'help' for a list of commands.")
	}
}

// Function to clear the terminal screen
func ClearScreen() {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default: // Unix-based systems (Linux/macOS)
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// Welcome Banner
func WelcomeBanner() {
	ClearScreen()
	color.Style{color.FgBlue, color.OpBold}.Println(`
****************************************************
   🌟 Welcome to the Interactive Chatroom CLI 🌟
****************************************************
Type 'help' to see the available commands.
`)
}

// Command to display help menu
func PrintHelp() {
	color.Style{color.FgYellow, color.OpItalic}.Println(`
Available Commands:
  🌐 join       Join the public chatroom
  ✉️  send       Send a message to the chatroom (requires joining)
  👀 view       View messages in the chatroom (submenu options included)
  ❌ exit       Exit the chatroom CLI
`)
}

// Command to join the chatroom
func JoinChatroom(ctx context.Context, hasJoined *bool, username string) {
	if *hasJoined {
		color.Style{color.FgMagenta}.Println("ℹ️  You have already joined the chatroom!")
		return
	}
	n, err := nats.GetInstance()
	if err != nil {
		log.Fatal(err)
	}
	go n.SubscribeToChat(ctx, username)
	*hasJoined = true
	color.Style{color.FgGreen}.Print("\n✅ You have joined !\n")
	color.Style{color.FgYellow}.Println("💬 You can now send and view messages.")
}

// Command to send a message
func SendMessage(ctx context.Context, username string, scanner *bufio.Scanner) {
	color.Cyan.Print("💬 Enter your message: ")
	if !scanner.Scan() {
		return
	}
	message := strings.TrimSpace(scanner.Text())
	if message == "" {
		color.Red.Println("❌ Message cannot be empty.")
		return
	}
	// publicChatroom.Messages = append(publicChatroom.Messages, message)
	n, err := nats.GetInstance()
	if err != nil {
		log.Fatal(err)
	}
	res := n.Publish(ctx, username, message)
	color.Green.Println("✅ Message sent!")
	color.Style{color.FgYellow}.Println("💬 You can send or view messages.")
	log.Println(res)
}

// Submenu for 'view' command
func ViewSubMenu(scanner *bufio.Scanner) {
	for {
		color.Style{color.FgLightBlue, color.OpBold}.Println(`
-------------------------------------
👀 What do you want to view !?

1. View Messages
2. View Active users
3. Back to Main Menu
-------------------------------------
		`)
		color.Cyan.Print("Enter your choice (1/2/3): ")
		if !scanner.Scan() {
			break
		}
		choice := strings.TrimSpace(scanner.Text())
		switch choice {
		case "1":
			ViewAllMessages()
		// case "2":
		// 	ViewUnreadMessages(CurrentUsername)
		case "2":
			ViewActiveUsers()
		case "3":
			color.Green.Println("\n🔙 Returning to the main menu...")
			WelcomeBanner() // Show main menu banner again
			return
		default:
			color.Red.Println("❌ Invalid choice. Please select 1, 2, or 3.")
		}
	}
}
func AuthenticateSubMenu(ctx context.Context, scanner *bufio.Scanner, username *string) (domain.Token, error) {
	for {
		color.Style{color.FgLightBlue, color.OpBold}.Println(`
-------------------------------------
👀 Authenticate Nedded !

1. Register
2. Login
3. Back to Main Menu
-------------------------------------
		`)
		color.Cyan.Print("Enter your choice (1/2/3): ")
		if !scanner.Scan() {
			break
		}
		choice := strings.TrimSpace(scanner.Text())
		switch choice {
		case "1":
			return SignUpUser(ctx, username)

		case "2":
			return SignInUser(ctx, username)
		case "3":
			color.Green.Println("\n🔙 Returning to the main menu...")
			WelcomeBanner() // Show main menu banner again
			return "", nil
		default:
			color.Red.Println("❌ Invalid choice. Please select 1, 2, or 3.")
		}
	}
	return "", nil
}

// View all messages
func ViewAllMessages() {
	n, err := nats.GetInstance()
	if err != nil {
		log.Fatal(err)
	}
	publicChatroomMessages, err := n.GetAllMessages(CurrentUsername)
	if err != nil {
		log.Print(err)
	}
	color.Style{color.FgYellow, color.OpItalic}.Println("\n💬 Messages in the chatroom:")
	if len(publicChatroomMessages) == 0 {
		color.Red.Println("❌ No messages in the chatroom yet.")
	} else {
		for _, message := range publicChatroomMessages {
			color.Cyan.Printf("[ %s ]: %s    -> send at %v\n", message.Sender, message.Content, message.Timestamp)
		}
	}
}
func ViewUnreadMessages(username string) {
	n, err := nats.GetInstance()
	if err != nil {
		log.Fatal(err)
	}
	publicChatroomMessages, err := n.ConsumeUnreadMessages(username)
	if err != nil {
		log.Print(err)
	}
	color.Style{color.FgYellow, color.OpItalic}.Println("\n💬 Messages in the chatroom:")
	if len(publicChatroomMessages) == 0 {
		color.Red.Println("❌ No messages in the chatroom yet.")
	} else {
		for i, message := range publicChatroomMessages {
			color.Cyan.Printf("%d: %s\n", i+1, message)
		}
	}
}
func ViewActiveUsers() {
	n, err := nats.GetInstance()
	if err != nil {
		log.Fatal(err)
	}
	publicChatroomMessages, err := n.GetActiveUsers()
	if err != nil {
		log.Print(err)
	}
	color.Style{color.FgYellow, color.OpItalic}.Println("\n💬 Active users in the chatroom:")
	if len(publicChatroomMessages) == 0 {
		color.Red.Println("❌ No active user in the chatroom.")
	} else {
		for i, consumer := range publicChatroomMessages {
			color.Cyan.Printf("%d: %s\n", i+1, consumer)
		}
	}
}

// Function to check if the user has joined the chatroom
func CheckAccess(hasJoined *bool) bool {
	if !*hasJoined {
		color.Red.Println("❌ You need to join the chatroom first! Use the 'join' command.")
		return false
	}
	return true
}
