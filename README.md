# Chat
TCP client and server chat written in go

## Server
Initiate 3 rooms - room 1,2 and 3 (hard coded configuration) in which the clients can connect to.
In order to run the server:
```
go run .\server\ --port <port>
```

## Client
Connects to a server specific room.
Simply run the following:
```
go run .\client\ --address <server_address> --room <room_number>
```

Enjoy :)