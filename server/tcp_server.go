package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func receiveTCPConn(ln *net.TCPListener) {
	for {
		err := ln.SetDeadline((time.Now().Add(time.Second * 10)))
		if err != nil {
			log.Fatal(err)
		}
		conn, err := ln.AcceptTCP() // クライアントからの接続を待ち受ける
		if err != nil {
			log.Fatal(err)
		}
		go handleRequest(conn) // 複数のリクエストに対応するように、別スレッドで実行する（ゴルーチンを使う）
	}
}

func handleRequest(conn *net.TCPConn)  {
	buffer := make([]byte, 1024)
	defer conn.Close()
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading: ", err.Error())
			return
		}
		fmt.Print(string(buffer[:n]))
		_, err = conn.Write(buffer[:n])
		if err != nil {
			fmt.Println("Error writing: ", err.Error())
			return
		}
	}
}

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":8080") // ソケットの作成とバインド
	if err != nil {
		log.Fatal(err)
	}

	ln, err := net.ListenTCP("tcp", tcpAddr) // リッスン
	if err != nil {
		log.Fatal(err)
	}

	receiveTCPConn(ln)
}
