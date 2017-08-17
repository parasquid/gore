package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/spf13/viper"
)

type MasterLogin struct {
	id            uint16
	version       uint32
	username      [24]byte
	password      [24]byte
	masterVersion uint8
}

func main() {
	viper.SetConfigType("json")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	ip := viper.GetString("server.ip")
	port := viper.GetString("server.port")

	addr := []string{ip, port}

	conn, err := net.Dial("tcp", strings.Join(addr, ":"))
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	var username [24]byte
	copy(username[:], viper.GetString("account.username"))
	var password [24]byte
	copy(password[:], viper.GetString("account.password"))

	packet := MasterLogin{
		id:            0x064,
		version:       0x012,
		username:      username,
		password:      password,
		masterVersion: 1,
	}
	var binBuf bytes.Buffer
	binary.Write(&binBuf, binary.LittleEndian, packet)

	fmt.Printf("Send: % x\n", binBuf.Bytes())

	conn.Write(binBuf.Bytes())

	buff := make([]byte, 1024)

	m, _ := conn.Read(buff)
	fmt.Printf("Receive: % x\n", buff[:m])

	n, _ := conn.Read(buff)
	fmt.Printf("Receive: % x\n", buff[:n])
}
