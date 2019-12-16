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

func UnmarshalObject(obj interface{}) {

}

func ReadBytesData(c *net.Conn) ([]byte, error) {
	conn := *c
	b := make([]byte, 100)
	_, err := conn.Read(b[0:])
	resBuf := append(b[0:], 0)
	if err != nil && err != io.EOF {
		return nil, err
	}
	for {
		_, err = conn.Read(b[:])
		resBuf = append(resBuf, b[0:]...)
		if err != nil && err != io.EOF {
			return nil, err
		}
		if err == io.EOF {
			break
		}
	}
	return resBuf, nil
}

func CheckError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}
