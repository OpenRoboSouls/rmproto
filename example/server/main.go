package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net"

	"github.com/openrobosouls/rmproto/client"
)

const SOF = uint8(client.SOF)

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			panic(err)
		}

		go func(c net.Conn) {
			defer c.Close()
			fmt.Println("Client connected:", c.RemoteAddr())
			defer fmt.Println("Client disconnected:", c.RemoteAddr())
			scanner := bufio.NewScanner(c)
			scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
				if atEOF && len(data) == 0 {
					return 0, nil, nil
				}
				if i := bytes.IndexByte(data, SOF); i >= 0 {
					// We have a full newline-terminated line.
					return i + 1, data[0:i], nil
				}
				// If we're at EOF, we have a final, non-terminated line. Return it.
				if atEOF {
					return len(data), data, nil
				}
				// Request more data.
				return 0, nil, nil
			})
			for scanner.Scan() {
				data := scanner.Bytes()
				if len(data) == 0 {
					continue
				}
				if data[0] != SOF {
					fmt.Println("Invalid frame received:", data)
					continue
				}
				frame := client.NewFrame()
				if err := frame.Unpack(0, data); err != nil {
					fmt.Println("Error unpacking frame:", err)
					continue
				}

				fmt.Println("Received frame:", frame)

			}
		}(conn)
	}
}
