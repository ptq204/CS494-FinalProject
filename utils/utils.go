package utils

import (
	"encoding/binary"
	"encoding/json"
	_ "final-project/message"
	"fmt"
	"io"
	"net"
)

func MarshalObject(obj interface{}) []byte {
	result, _ := json.Marshal(obj)
	return result
}

func UnmarshalObject(obj interface{}, data []byte) error {
	err := json.Unmarshal(data, obj)
	return err
}

func ReadBytesResponse(c *net.Conn) ([]byte, error) {
	conn := *c
	b := make([]byte, 1024)
	nBytes, err := conn.Read(b[:])
	if err != nil {
		return nil, err
	}
	return b[:nBytes], nil
}

func ReadBytesData(c *net.Conn) ([]byte, uint32, error) {
	conn := *c
	metadata := make([]byte, 4)
	b := make([]byte, 100)

	_, err := conn.Read(metadata[0:]) // read data length

	if err != nil {
		return nil, 20, err
	}

	dataLength := binary.BigEndian.Uint32(metadata)

	fmt.Printf("DATA LENGTH FIRST: %d bytes\n", dataLength)

	_, err = conn.Read(metadata[0:]) // read actiontype

	actionType := binary.BigEndian.Uint32(metadata)

	fmt.Printf("ACTION TYPE: %d \n", actionType)

	if err != nil {
		return nil, 20, err
	}

	nBytes, err := conn.Read(b[0:])

	if err != nil && err != io.EOF {
		return nil, 20, err
	}

	fmt.Printf("READ %d bytes\n", nBytes)

	resBuf := append(b[0:nBytes], 0)
	fmt.Printf("DATA: %s\n", string(resBuf[:]))

	dataLength -= uint32(nBytes)

	fmt.Printf("DATA LENGTH: %d bytes\n", dataLength)

	for dataLength > 0 {
		nBytes, err = conn.Read(b[:])
		if err != nil && err != io.EOF {
			return nil, 20, err
		}
		resBuf = append(resBuf, b[0:nBytes]...)
		dataLength -= uint32(nBytes)
		fmt.Printf("DATA LENGTH: %d bytes\n", dataLength)
	}

	fmt.Println(string(resBuf[:]))

	fmt.Println("PASSSSS")
	return resBuf, actionType, nil
}

func TellReadDone(c *net.Conn) {
	conn := *c
	conn.Write([]byte("DONE"))
}

func CheckError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}
