#

Simple concurrent TCP server implementation. It accepts incoming connections, receives messages, and sends a response back to the connected client.

## Example

After running the server `nc localhost 2319` opens a connection with server from terminal.

![example](example.png?raw=true)

## How it Works

- The `main` function starts by setting up a TCP listener on port 2319.
- Incoming connections are accepted in an infinite loop.
- For each accepted connection, a new goroutine (`handleNetConnection`) is launched to handle the communication.
- The `handleNetConnection` function receives messages from the client, processes them, and sends a response. If the received message is "CLOSE\n", it closes the connection.
- The server will continue running until terminated.
