package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("client dial err=", err)
		return
	}

	defer conn.Close()

	go io.Copy(os.Stdout, conn)
	fi, _ := os.Stdin.Stat()

	//判断是windows还是UNIX类型系统，windows os.Stdin没有"\n"会导致命令执行失败
	if (fi.Mode() & os.ModeCharDevice) == 0 {
		buffer, _ := ioutil.ReadAll(os.Stdin)
		io.Copy(conn, bytes.NewReader(buffer))
	} else {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			io.WriteString(conn, input.Text()+"\n")
		}
	}
}
