package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func responses(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8080") // localhost:8080 に接続する
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter text: ")
		text, _ := reader.ReadString('\n')
        fmt.Fprintln(conn, text)

		buf := make([]byte, 1024)
		count, err := conn.Read(buf)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("out: " + string(buf[:count]))

		responses(os.Stdout, conn)
	}
}
