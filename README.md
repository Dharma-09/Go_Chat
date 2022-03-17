# Go_Chat

Once the user connects to the chat server using `telnet` command line program, they can use the following commands to talk to the server:

- `/nick <name>` - get a name, otherwise user will stay anonymous.
- `/join <name>` - join a room, if room doesn't exist, the new room will be created. User can be only in one room at the same time.
- `/rooms` - show list of available rooms to join.
- `/msg	<msg>` - broadcast message to everyone in a room.
- `/quit` - disconnects from the chat server.

## Testing

Now it's time to build, run and test it using `telnet` command.

```
go build .
./chat
server started on :8888
```
