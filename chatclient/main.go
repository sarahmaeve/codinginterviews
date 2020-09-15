package main

// original spec : https://gist.github.com/jixwanwang/e869b91985135771495161d72a53aa69
import (
	"fmt"
	"net"
)

// stable connection for a user
// at least give an anon id - or other userid
// pass messages through channels to all users
// handle disconnect; ctrl-D or (perhaps) process close
// continuous handling of messages

func handleConnection(conn net.Conn) {
	buf := make([]byte, 1024)

	// probably need error handling here
	// research which types of errors are most prevalent

	// defer conn.Close()
	for {
		// side behavior - unclosed conn reports 0 bytes forever
		msgLen, err := conn.Read(buf)
		if msgLen == 0 {
			break
		}
		// want to trap EOF-type error and say goodbye
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		fmt.Printf("Message length was %d\n", msgLen)
		// attempt to echo
		conn.Write(buf)
	}

	conn.Close()

}

func main() {
	l, err := net.Listen("tcp", ":7979")
	if err != nil {
		fmt.Println(err.Error())
	}

	// example: defer?
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err.Error())
		}
		go handleConnection(conn)
	}
}
