package utils

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"encoding/csv"
	"encoding/json"
	"final-project/constant"
	"final-project/decrypt"
	e "final-project/encrypt"
	_ "final-project/message"
	"fmt"
	"io"
	"net"
	"os"
	"strings"

	"github.com/howeyc/gopass"
)

func MarshalObject(obj interface{}) []byte {
	result, err := json.Marshal(obj)
	if err != nil {
		fmt.Printf("Marshal error: %s", err.Error())
	}
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

	bytes, err := conn.Read(metadata[:]) // read data length
	fmt.Printf("Data length: %d\n", bytes)
	if err != nil {
		return nil, 20, err
	}

	dataLength := binary.BigEndian.Uint32(metadata)

	fmt.Printf("DATA LENGTH FIRST: %d bytes\n", dataLength)

	bytes, err = conn.Read(metadata[:]) // read actiontype

	fmt.Printf("Action length: %d\n", bytes)

	actionType := binary.BigEndian.Uint32(metadata)

	fmt.Printf("ACTION TYPE: %d \n", actionType)

	if err != nil {
		return nil, 20, err
	}
	var resBuf []byte
	for {
		nBytes, err := conn.Read(b[:])
		fmt.Printf("DATA APPEND: %s\n", string(b[:nBytes]))
		if err != nil && err != io.EOF {
			return nil, 20, err
		}
		resBuf = append(resBuf, b[:nBytes]...)
		fmt.Printf("DATA: %s\n", string(resBuf))
		dataLength -= uint32(nBytes)
		fmt.Printf("DATA LENGTH: %d\n", dataLength)
		if dataLength == 0 {
			return resBuf, actionType, nil
		}
	}
}

func TellReadDone(c *net.Conn) {
	conn := *c
	conn.Write([]byte("DONE"))
}

func SaveLocalValueToFile(key string, value string) error {
	fileName := "../client/.local"
	var f *os.File
	var err error
	if _, err = os.Stat(fileName); os.IsNotExist(err) {
		f, err = os.Create(fileName)
		if err != nil {
			return err
		}
	} else {
		f, err = os.OpenFile(fileName, os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
	}
	_, err = f.WriteString(key + ": ")
	if err != nil {
		return err
	}
	_, err = f.WriteString(value)
	if err != nil {
		return err
	}
	defer f.Close()
	return nil
}

func ReadLocalValueInFile(key string) (string, error) {
	fileName := "../client/.local"
	f, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		keyVal := strings.Split(line, " : ")
		if keyVal[0] == key {
			return keyVal[1], nil
		}
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return "", nil
}

func CheckError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}

func IsExitCommand(command string) bool {
	return command == "exit" || command == "quit" || command == "q"
}

func InputPassword() string {
	fmt.Print(">> password: ")
	pass, _ := gopass.GetPasswdMasked()
	passStr := strings.TrimRight(string(pass), "\n")
	return passStr
}

func InputNewPassword() string {
	fmt.Print(">> new password: ")
	pass, _ := gopass.GetPasswdMasked()
	passStr := strings.TrimRight(string(pass), "\n")
	return passStr
}

func SendFileData(c *net.Conn, fileName string, encrypt int) error {
	conn := *c
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	buff := make([]byte, 1024)
	for {
		nBytes, err := f.Read(buff)
		// fmt.Println(nBytes)
		if err != nil && err != io.EOF {
			return err
		}
		if nBytes >= 0 {
			finalNBytes := nBytes
			if encrypt == 1 && finalNBytes > 0 {
				// PUT Encrypt function here
				fmt.Println(string(buff[0:nBytes]))
				buff = []byte(e.Encrypt(constant.PASSPHRASE, string(buff[0:nBytes])))
				finalNBytes = len(buff)
				fmt.Println(string(buff))
			}
			fmt.Println(finalNBytes)
			buffSend := make([]byte, 4+finalNBytes)

			buffBytesLength := new(bytes.Buffer)
			numBytesBuff := make([]byte, 4)
			binary.BigEndian.PutUint32(numBytesBuff, uint32(finalNBytes))
			err := binary.Write(buffBytesLength, binary.BigEndian, numBytesBuff)
			if err != nil {
				return err
			}

			copy(buffSend[:4], buffBytesLength.Bytes())
			copy(buffSend[4:], buff[:])

			conn.Write(buffSend)
		}
		if nBytes == 0 {
			f.Close()
			break
		}
	}

	return nil
}

func ReceiveFile(c *net.Conn, fileName string, fileSize int64, encrypt int) error {
	fmt.Println("Start Receiving File")
	conn := *c
	var dataLength uint32
	var nBytes int
	prevNBytes := -1

	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	for {
		header := make([]byte, 4)

		_, err := conn.Read(header[0:])
		if err != nil {
			return err
		}
		dataLength = binary.BigEndian.Uint32(header)
		// fmt.Println(dataLength)
		fileChunk := make([]byte, dataLength)
		if dataLength > 0 {
			nBytes, err = conn.Read(fileChunk[:])
			if err != nil {
				return err
			}

			if encrypt == 1 {
				// Put Decrypt Function here
				fmt.Println(string(fileChunk))
				fileChunkString := decrypt.Decrypt(constant.PASSPHRASE, string(fileChunk[0:nBytes]))
				fmt.Println(fileChunkString)
				fileChunk = []byte(fileChunkString)
				fmt.Println(string(fileChunk))
				f.Write(fileChunk[:])
			} else {
				f.Write(fileChunk[:])
			}
		}
		if dataLength == 0 || nBytes < prevNBytes {
			f.Close()
			break
		}
		prevNBytes = nBytes
	}
	return nil
}

func SplitCommand(s string) ([]string, error) {
	r := csv.NewReader(strings.NewReader(s))
	r.Comma = ' ' // space
	return r.Read()
}
