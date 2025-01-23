package helper

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/RaziyeNikookolah/chatroom-using-go-nats/api/pb"
	"github.com/RaziyeNikookolah/chatroom-using-go-nats/client/domain"
	"github.com/RaziyeNikookolah/chatroom-using-go-nats/pkg/adapters/clients/grpc/mappers"
	"google.golang.org/grpc"
)

const (
	userFile    = "users.json"  // File to store user data
	sessionFile = "session.txt" // File to store session data
)

func LoginUsingGrpc(ctx context.Context, conn *grpc.ClientConn, user domain.User) (domain.Token, error) {
	// Create a new UserService client
	client := pb.NewUserServiceClient(conn)

	in := &pb.LoginRequest{
		Username: user.Username,
		Password: user.Password,
	}

	response, err := client.Login(ctx, in)
	if err != nil {
		log.Fatalf("cannot login: %v", err)
	}

	// Map the response to the domain model
	token, err := mappers.LoginResponseProtoToLoginResponseDomain(response)
	if err != nil {
		log.Fatalf("cannot map response: %v", err)
	}

	// Print the domain response
	// log.Print("User logged in successfully")
	return domain.Token(token.Token), nil
}
func RegisterUsingGrpc(ctx context.Context, conn *grpc.ClientConn, user domain.User) (domain.Token, error) {
	// Create a new UserService client
	client := pb.NewUserServiceClient(conn)

	// Prepare the request
	in := &pb.RegisterRequest{
		Username: user.Username,
		Password: user.Password, //validation needed
		Email:    user.Email,    // valida....
	}

	// Call the CreateWallet method
	response, err := client.Register(ctx, in)
	if err != nil {
		log.Fatalf("cannot register user: %v", err)
	}

	// Map the response to the domain model
	token, err := mappers.RegisterResponseProtoToRegisterResponseDomain(response)
	if err != nil {
		log.Fatalf("cannot map response: %v", err)
	}

	// Print the domain response
	// log.Print("User registered successfully")
	return domain.Token(token.Token), nil
}
func GetUserClaimUsingGrpc(ctx context.Context, conn *grpc.ClientConn, token string) (domain.UserClaim, error) {
	// Create a new UserService client
	client := pb.NewUserServiceClient(conn)

	// Prepare the request
	in := &pb.TokenRequest{
		Token: token,
	}

	// Call the CreateWallet method
	response, err := client.GetUserClaimWithToken(ctx, in)
	if err != nil {
		log.Fatalf("cannot parse token: %v", err)
	}

	// Map the response to the domain model
	claim, err := mappers.UserClaimResponseProtoToUserResponseDomain(response)
	if err != nil {
		log.Fatalf("cannot map response: %v", err)
	}

	// Print the domain response
	// log.Print("User registered successfully")
	return domain.UserClaim{
		Username: claim.Username,
		ID:       claim.ID,
		Email:    claim.Email,
	}, nil
}
func SessionExists() bool {
	_, err := os.Stat(sessionFile)
	return err == nil
}

func RestoreSession(ctx context.Context, conn *grpc.ClientConn) (domain.UserClaim, error) {
	data, _ := os.ReadFile(sessionFile)
	user, err := GetUserClaimUsingGrpc(ctx, conn, string(data))
	if err != nil {
		return domain.UserClaim{}, err
	}
	return user, nil
}

func SaveSession(token domain.Token) {
	// data, _ := json.Marshal(user)
	os.WriteFile(sessionFile, []byte(token), 0644)
}

// func SaveSession(token string) error {
// 	var sessionFile string

// 	switch runtime.GOOS {
// 	case "windows":
// 		appDataPath, _ := os.UserConfigDir() // C:\Users\YourUser\AppData\Local
// 		sessionFile = filepath.Join(appDataPath, "MyApp", "session.dat")
// 	case "linux", "darwin": // darwin baraye macOS
// 		configPath, _ := os.UserConfigDir() // /home/user/.config
// 		sessionFile = filepath.Join(configPath, "MyApp", "session.dat")
// 	case "android":
// 		sessionFile = "/data/data/com.myapp/files/session.dat"
// 	case "ios":
// 		homeDir, _ := os.UserHomeDir()
// 		sessionFile = filepath.Join(homeDir, "Library", "Application Support", "MyApp", "session.dat")
// 	default:
// 		return fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
// 	}

// 	// Ijad folder agar vojood nadarad
// 	err := os.MkdirAll(filepath.Dir(sessionFile), os.ModePerm)
// 	if err != nil {
// 		return fmt.Errorf("error creating directory: %v", err)
// 	}

// 	// Zakhire token ba permission haye monaseb
// 	err = os.WriteFile(sessionFile, []byte(token), 0600) // 0600: faghat malek dastresi darad
// 	if err != nil {
// 		return fmt.Errorf("error saving session: %v", err)
// 	}

//		fmt.Println("Token saved to:", sessionFile)
//		return nil
//	}
func HandleAuthentication(ctx context.Context, conn *grpc.ClientConn, scanner *bufio.Scanner) (domain.Token, error) { // user does not have token -> login or register required
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Println("Welcome to the Chatroom!")
	// fmt.Print("Do you have an account? (yes/no): ")
	return AuthenticateSubMenu(scanner, ctx, conn)
	// choice, _ := reader.ReadString('\n')
	// choice = strings.TrimSpace(choice)
	// for {
	// 	if choice == "1" {
	// 		token, err := SignUpUser(ctx, conn)
	// 		if err != nil {
	// 			log.Print(err)
	// 			return "", err
	// 		}
	// 		return token, nil
	// 	} else if choice == "2" {
	// 		token, err := SignInUser(ctx, conn)
	// 		if err != nil {
	// 			log.Print(err)
	// 			return "", err
	// 		}
	// 		return token, nil
	// 	} else {
	// 		// AuthenticateSubMenu(scanner, ctx, conn)

	// 	}
	// }
}

func SignUpUser(ctx context.Context, conn *grpc.ClientConn) (domain.Token, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(">> Sign Up")

	fmt.Print("Enter your unique username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Print("Enter your email: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	fmt.Print("Enter a password: ") ///validation hint
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	user := domain.User{Username: username, Email: email, Password: password}
	token, err := RegisterUsingGrpc(ctx, conn, user)
	if err != nil {
		return domain.Token(""), err
	}
	fmt.Println("Sign-up successful! You are logged in.")
	return domain.Token(token), nil
}

func SignInUser(ctx context.Context, conn *grpc.ClientConn) (domain.Token, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(">>Login")
	fmt.Print("Enter your username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Print("Enter your password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)
	user := domain.User{Username: username, Email: "", Password: password}

	token, err := LoginUsingGrpc(ctx, conn, user)
	if err != nil {
		return domain.Token(""), err
	}
	fmt.Println("Sign-in successful! You are logged in.")
	return domain.Token(token), nil
}
