package main

type server struct{
	rooms map[string]*room
	commands chan command
}

func newServer() *server{
	return &server{
		rooms: make(map[string]*room),
		commands:make(chan command),
	}
}
func(s *server) run(){
	for cmd := range s.commands{
		switch cmd.id {
		case CMD_NICK:
			s.nick(cmd.client, cmd.args)
		case CMD_JOIN:
			s.join(cmd.client,cmd.args)
		case CMD_ROOMS:
			s.listRooms(cmd.client,cmd.args)
		case CMD_MSG:
			s.msg(cmd.client,cmd.args)
		case CMD_QUIT:
			s.quit(cmd.client,cmd.args)
		}
	}
}
func(s*server) newClient(conn net.Conn){
	log.Print("new client has connected:%s",conn.ddr().String())

	c :=&client{
		conn:conn,
		nick:"anonymous",
		commands: s.commands
	}
	c.readInput()
}

func(s *server) nick(c *client, args []string){
	c.nick =args[1]
	c.msg(fmt.Sprintf("all right,I will call you %s", c.nicks))
}
func(s *server) join(c *client, args []string){
	roomName := args[1]
	r,ok := s.rooms[roomName]
	if !ok {
		r = &room{
			name:roomName,
			members: make(map[net.Addr]*client),
		}
		s.room[roomName] = r
	}
	r.members[c.conn.RemoteAddr()] = c

	s.quitCurrentRoom(call)

	c.room=r
	 r.broadcaste(c, fmt.Sprintf("%s has joined the room",c.nick))
	 c.msg(fmt.Sprintf("welcome to %s",r.name))
}

func(s *server) listRooms(c *client, args []string){
	var rooms []string
	for name:= range s.rooms{
		rooms = append(rooms,name)

	}
	c.msg(fmt.Sprintf("available rooms are:%s",string.Join(rooms,",")))
}

func(s *server) msg(c *client, args []string){
	if c.room == nil{
		c.err(errors.New("you must join the room first"))
		return
	}
	c.room.broadcaste(c,c.nick+":"strings.join(args[1:len(args)],""))
}

func(s *server) quit(c *client, args []string){
	 log.Print("client has disconnected :%s",c.conn.RemoteAddr().String())

	 s.quitCurrentRoom(c)

	 c.msg("sad to see you go :(")
	 c.conn.Close
}

func (s *server) quitCurrentRoom(c *client){
	if c.room != nil{
		delete(c.room.members,c.conn, RemoteAddr())
		c.room.broadcaste(c, fmt.Sprintf("%s has left the room",c.nick))
	}
}
