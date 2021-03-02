package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	nl, err := net.Listen("tcp", ":8888") //1 to 65535
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1) //1 = stop with error
	}

	conn, err := nl.Accept()
	if err != nil {
		fmt.Println(err.Error())
		//continue
	}

	// scanner := bufio.NewScanner(conn)
	// for scanner.Scan(){

		//line := scanner.Text()
		// fmt.Println(line)
		//if line == "" {
			// 	fmt.Println("line ends")
			//	break
		//}
	//}

	bs := make([]byte, 1024) //text asass asas
	n, e := conn.Read(bs)
	if e != nil {
		fmt.Println(e.Error())
	}
	
	fmt.Println(n)
	//fmt.Println(bs)
	reqstr	:= string(bs) //convertion
	fmt.Println(reqstr)
	//reqSlc := strings.Fields(reqstr)
	//fmt.Println(reqSlc, len(reqSlc))

	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>Hello golang world</strong></body></html>`

	// fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	// fmt.Fprint(conn, "Content-Length: %d\r\n", len(body))
	// fmt.Fprint(conn, "Content-Type: text/html\r\n")
	// fmt.Fprint(conn, "\r\n")
	// fmt.Fprint(conn, body)

	resp := "HTTP/1.1 200 OK\r\n" +
		"Content-Length: %d\r\n" +
		"Content-Type: text/html\r\n" +
		"\r\n%s"

		msg := fmt.Sprintf(resp, len(body), body)
		fmt.Println(msg)
		conn.Write([]byte(msg))

		//conn.Close
}
