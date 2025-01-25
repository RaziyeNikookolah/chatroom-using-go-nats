# A Basic Public Chatroom using Go, NATS and Postgresql

================

## Features

This chatroom project provides the following features:

- **User Authentication:**

  - Users can authenticate and receive a JWT token.
  - The token is stored in the user's session file to authentication in future logins.
  - When sending the token to the service via gRPC, the user is validated.

- **User Registration & Login:**

  - User data is stored in a PostgreSQL database.
  - User registration and login are handled via gRPC requests.

- **Command-Line Interface (CLI):**

  - The chatroom is accessed via a CLI interface.
  - The system communicates with NATS and utilizes JetStream with "at least once" delivery guarantee.

## Setup Instructions

Follow these steps to set up and run the project:

1.  git clone https://github.com/RaziyeNikookolah/chatroom-using-go-nats.git
2.  cd chatroom-using-go-nats.git
3.  run go mod tidy
4.  docker-compose up --build
5.  go run ./cmd/chatroom/main.go
6.  go run ./client/cmd/command-line-interface/main.go

## Application Features

- **Join, Send, and View Messages:**

  - Users can join the public chatroom, send messages, and view messages in real-time.

- **User Registration and Identification:**

  - Users are registered in the database and identified via JWT tokens.

- **Active User Tracking:**

  - The system allows users to see active participants in the chatroom.

## Technologies Used

- **Go** - Core programming language
- **PostgreSQL** - Database for storing user data
- **gRPC** - Communication between services
- **NATS JetStream** - Message broker for reliable messaging
- **Docker** - Containerization of services

## Contribution

If you'd like to contribute:

1.  Fork the repository and Star it :) .
2.  Create a new branch for your feature.
3.  Submit a pull request.

## License

This project is licensed under the MIT License.
