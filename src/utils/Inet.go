package utils

import (
	"bytes"
	"encoding/binary"
	"errors"
	"strconv"
	"strings"
)

func INet_AToN(ipv4String string) (uint32, error) {

	if len(ipv4String) == 0 {
		return 0, errors.New("param is empty")
	}

	addArr := strings.Split(ipv4String, ".")
	if len(addArr) != 4 {
		return 0, errors.New("ipv4 format error, ipv4=" + ipv4String)
	}

	ipv4ToBytes := make([]byte, 4)
	bytesBuffer := bytes.NewReader(ipv4ToBytes)
	var x uint32

	for i := 0; i < 4; i++ {
		tem, _ := strconv.Atoi(strings.TrimSpace(addArr[i]))
		ipv4ToBytes[i] = byte(tem)
	}
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	return x, nil
}

func INet_NToA(ipv4Num uint32) (string, error) {
	bytesBuffer := bytes.NewBuffer([]byte{})
	var ipv4Bytes []byte
	err := binary.Write(bytesBuffer, binary.BigEndian, ipv4Num)
	if err != nil {
		return "", err
	}
	ipv4Bytes = bytesBuffer.Bytes()
	return strings.Join([]string{
		strconv.Itoa(int(ipv4Bytes[0])),
		strconv.Itoa(int(ipv4Bytes[1])),
		strconv.Itoa(int(ipv4Bytes[2])),
		strconv.Itoa(int(ipv4Bytes[3]))}, "."), nil

}
