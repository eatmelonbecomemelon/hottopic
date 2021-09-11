package main

import (
	"bytes"
	cf "config"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"mongoplus"
	"testing"
	"weibo"
)

func Test_GetWeiboHotTopic(t *testing.T) {
	cf.LoadConfig()
	mongoplus.MongoInit()
	weibo.GetWeiboHotTopic()
}

func Test_binary(t *testing.T) {
	var temp TcpProtocol
	temp.vrvid = [16]uint8{0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x18, 0x17, 0x16, 0x15, 0x14, 0x13, 0x12}
	temp.cmd = 0x0043
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, temp.cmd)
	fmt.Println(hex.EncodeToString(buf.Bytes()))
}

type TcpProtocol struct {
	magic [2]uint8
	vrvid [16]uint8
	cmd   uint16
	seq   uint16
	len   [4]uint8
}
