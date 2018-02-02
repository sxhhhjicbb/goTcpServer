package main

import (
	"database/sql"
	"fmt"
	"net"
	"call"
)

func HandleConn(conn net.Conn){
	defer conn.Close()
	rip := conn.RemoteAddr().String()	
	fmt.Println("recv infor is",rip)
	buffer := make([]byte,1024)
	for {
		readLen,err := conn.Read(buffer)
		if err != nil {
			fmt.Println("exit link:",rip)
			conn.Close()
			return
		}
		if readLen == 0{
			fmt.Println(rip,"disconnect!")
			conn.Close()
			return
		}
		fmt.Println("recv len is",readLen)
		call.Mydisplay(readLen)
		conn.Write(buffer[:readLen])
	}
}

func main() {
	fmt.Println("start...")
	db,err := sql.Open("mysql","user:password@/dbname")
	if(err != nil){
		fmt.Println("mysql open error",err)
	}
	defer db.Close()
	l,err := net.Listen("tcp",":8888")
	if err != nil {
		fmt.Println("listen error:",err)
		return 
	}
	defer l.Close()
	for {
		conn,err := l.Accept()
		if err != nil {
			fmt.Println("accept error: ",err)
			break;
		}

		go HandleConn(conn)
	}
}