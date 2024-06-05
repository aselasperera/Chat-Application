# Chat Application

This is a simple real-time chat application built using Go Fiber and WebSockets for bidirectional communication between client and server. The front-end is a basic HTML file that communicates with the Go back-end server.

## Project Structure
```
chat-application/
├── main.go
└── index.html
```
## Prerequisites

- Go (version 1.16 or higher)
- A modern web browser (for example, Google Chrome, Firefox, etc.)

## Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/aselasperera/Chat-Application
    cd chat-application
    ```

2. Install the Go Fiber framework and WebSocket dependencies:
    ```bash
    go get github.com/gofiber/fiber/v2
    go get github.com/gofiber/contrib/websocket
    ```

## Usage

1. Start the Go server:
    ```bash
    go run main.go
    ```

2. Open `index.html` in your web browser:
    - You can directly open the file using your browser's file explorer or use a simple HTTP server to serve the file.
    - To start a simple HTTP server, you can use Python if it's installed on your machine:
      ```bash
      python -m http.server 8080
      ```
      Then navigate to `http://localhost:8080/index.html` in your web browser.

3. Start chatting:
    - Open multiple tabs or browsers and start sending messages. You'll see the messages appear in real-time in all connected clients.

## Code Overview

### `index.html`

- This file contains the front-end code for the chat application.
- It includes a basic HTML structure with an input field for typing messages and a list to display messages.
- JavaScript code establishes a WebSocket connection to the server and handles sending and receiving messages.

### `main.go`

- This file contains the back-end code written in Go.
- It uses the Fiber framework to set up a web server and WebSocket connections.
- It includes functions to handle incoming messages and broadcast them to all connected clients.

## Features

- Real-time messaging using WebSockets.
- Simple and clean user interface.
- Automatic scrolling to the latest message.

## Future Improvements

- Add user authentication.
- Enhance the user interface with better styling.
- Add support for multimedia messages (images, videos, etc.).

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contributing

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Make your changes.
4. Commit your changes (`git commit -am 'Add new feature'`).
5. Push to the branch (`git push origin feature-branch`).
6. Create a new Pull Request.

## Acknowledgements

- [Go Fiber](https://gofiber.io/)
- [WebSocket](https://developer.mozilla.org/en-US/docs/Web/API/WebSocket)
