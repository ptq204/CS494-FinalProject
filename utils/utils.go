package utils

import (
	"bytes"
	"encoding/gob"
	_ "final-project/message"
	"fmt"
	"io"
	"net"
)

func MarshalObject(obj interface{}) []byte {
	buffPayload := new(bytes.Buffer)
	goobj := gob.NewEncoder(buffPayload)
	goobj.Encode(obj)
	return buffPayload.Bytes()
}

func UnmarshalObject(obj interface{}, data []byte) error {
	tmpPayload := bytes.NewBuffer(data)
	d := gob.NewDecoder(tmpPayload)
	err := d.Decode(obj)
	return err
}

func ReadBytesData(c *net.Conn) ([]byte, error) {
	conn := *c
	b := make([]byte, 100)
	nBytes, err := conn.Read(b[0:])
	fmt.Printf("READ %d bytes\n", nBytes)
	resBuf := append(b[0:nBytes], 0)
	fmt.Println(string(resBuf[:]))

	if err != nil && err != io.EOF {
		return nil, err
	}
	if nBytes < 100 {
		return resBuf, nil
	}
	for {
		nBytes, err := conn.Read(b[0:])
		fmt.Printf("READ %d bytes\n", nBytes)
		if string(b[0:nBytes]) == "DONE" {
			break
		}
		resBuf = append(resBuf, b[0:nBytes]...)
		if nBytes < 100 {
			break
		}
		fmt.Println(string(resBuf[:]))
		if err != nil && err != io.EOF {
			fmt.Println("STUCK HERE ERROR 2")
			return nil, err
		}
	}
	fmt.Println("PASSSSS")
	return resBuf, nil
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
