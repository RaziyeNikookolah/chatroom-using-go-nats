package helper

import (
	"bufio"
	"context"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/RaziyeNikookolah/chatroom-using-go-nats/client/domain"
	"github.com/gookit/color"
	"google.golang.org/grpc"
)

func MainMenu(scanner *bufio.Scanner, hasJoined *bool) {
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
		JoinChatroom(hasJoined)
	case "send":
		if CheckAccess(hasJoined) {
			SendMessage(scanner)
		}
	case "view":
		if CheckAccess(hasJoined) {
			ViewSubMenu(scanner)
		}
	case "exit":
		color.Green.Println("\n👋 Goodbye! Thanks for chatting!")
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
  ✉️ send       Send a message to the chatroom (requires joining)
  👀 view       View messages in the chatroom (submenu options included)
  ❌ exit       Exit the chatroom CLI
`)
}

// Command to join the chatroom
func JoinChatroom(hasJoined *bool) {
	if *hasJoined {
		color.Style{color.FgMagenta}.Println("ℹ️  You have already joined the chatroom!")
		return
	}
	*hasJoined = true
	color.Style{color.FgGreen}.Print("\n✅ You have joined !\n")
	color.Style{color.FgYellow}.Println("💬 You can now send and view messages.")
}

// Command to send a message
func SendMessage(scanner *bufio.Scanner) {
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
	color.Green.Println("✅ Message sent!")
}

// Submenu for 'view' command
func ViewSubMenu(scanner *bufio.Scanner) {
	for {
		color.Style{color.FgLightBlue, color.OpBold}.Println(`
-------------------------------------
👀 View Submenu:
1. View Messages
2. View Active users Message
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
func AuthenticateSubMenu(scanner *bufio.Scanner, ctx context.Context, conn *grpc.ClientConn) (domain.Token, error) {
	for {
		color.Style{color.FgLightBlue, color.OpBold}.Println(`
-------------------------------------
👀 Authenticate Submenu:
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
			return SignUpUser(ctx, conn)

		case "2":
			return SignInUser(ctx, conn)
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
	publicChatroomMessages := ""
	color.Style{color.FgYellow, color.OpItalic}.Println("\n💬 Messages in the chatroom:")
	if len(publicChatroomMessages) == 0 {
		color.Red.Println("❌ No messages in the chatroom yet.")
	} else {
		for i, message := range publicChatroomMessages {
			color.Cyan.Printf("%d: %s\n", i+1, message)
		}
	}
}

// View last message
func ViewActiveUsers() {
	// if len(publicChatroom.Messages) == 0 {
	// 	color.Red.Println("\n❌ No messages in the chatroom yet.")
	// } else {
	// 	color.Cyan.Printf("\n💬 Last message: %s\n", publicChatroom.Messages[len(publicChatroom.Messages)-1])
	// }
}

// Function to check if the user has joined the chatroom
func CheckAccess(hasJoined *bool) bool {
	if !*hasJoined {
		color.Red.Println("❌ You need to join the chatroom first! Use the 'join' command.")
		return false
	}
	return true
}
