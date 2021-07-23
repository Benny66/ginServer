package function /*
/*
 * @Descripttion:
 * @version: v1.0.0
 * @Author: shahao
 * @Date: 2021-04-08 16:09:43
 * @LastEditors: shahao
 * @LastEditTime: 2021-07-20 15:48:15
*/

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unsafe"

	"golang.org/x/text/encoding/unicode"
)

func MacAddressToByte(macAddress string) []byte {
	macAddress = strings.Replace(macAddress, "-", "", -1)
	return Hextobyte(macAddress)
}

func MacAddressByteToStr(data []byte) (macAddrrss string) {
	macAddrrss = ""
	temp := BytetoHex(data)
	for i := 0; i < len(temp); i = i + 2 {
		macAddrrss += temp[i:i+2] + "-"
	}
	macAddrrss = strings.TrimRight(macAddrrss, "-")
	return
}

func IPddressByteToStr(data []byte) (ipAddress string) {

	for _, value := range data {
		ipAddress += strconv.Itoa(int(value)) + "."
	}
	ipAddress = strings.TrimRight(ipAddress, ".")
	return
}

func IPddressStrToByte(ipAddress string) (data []byte) {
	ipAddrArr := strings.Split(ipAddress, ".")
	for _, value := range ipAddrArr {
		i, _ := strconv.Atoi(value)
		data = append(data, byte(i))
	}
	return
}
func Hextobyte(str string) []byte {
	slen := len(str)
	bHex := make([]byte, len(str)/2)
	ii := 0
	for i := 0; i < len(str); i = i + 2 {
		if slen != 1 {
			ss := string(str[i]) + string(str[i+1])
			bt, _ := strconv.ParseInt(ss, 16, 32)
			bHex[ii] = byte(bt)
			ii = ii + 1
			slen = slen - 2
		}
	}
	return bHex
}

func BytetoHex(b []byte) (H string) {
	H = fmt.Sprintf("%x", b)
	return
}
func BytetoHexSpace(b []byte) (H string) {
	if len(b) > 0 {
		for i := 0; i < len(b); i++ {
			H += fmt.Sprintf("%x", b[i:i+1]) + " "
		}
	}
	return
}

func BytesCombine(pBytes ...[]byte) []byte {
	return bytes.Join(pBytes, []byte(""))
}

func BytesToInt(b []byte) int64 {
	bytesBuffer := bytes.NewBuffer(b)
	var x int16
	binary.Read(bytesBuffer, binary.LittleEndian, &x)
	return int64(x)
}

func IntToBytes(b int) []byte {
	buff := make([]byte, 2)
	binary.LittleEndian.PutUint16(buff, uint16(b))
	return buff
}

func ByteBigToLittleEndian(b []byte) []byte {

	var buff []byte = make([]byte, 0)

	var tempBytes []byte = make([]byte, 2)
	for i := 0; i < len(b); i = i + 2 {
		convInt := binary.BigEndian.Uint16(b[i : i+2])
		binary.LittleEndian.PutUint16(tempBytes, convInt)
		buff = BytesCombine(buff, tempBytes)
	}
	return buff

}

func ByteLittleToBigEndian(b []byte) []byte {

	var buff []byte = make([]byte, 0)

	var tempBytes []byte = make([]byte, 2)
	for i := 0; i < len(b); i = i + 2 {
		convInt := binary.LittleEndian.Uint16(b[i : i+2])
		binary.BigEndian.PutUint16(tempBytes, convInt)
		buff = BytesCombine(buff, tempBytes)
	}
	return buff
}

func Sbyte2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
func ByteUtf16ToUtf8Str(buff []byte) (data string, err error) {
	buff = bytes.TrimRight(buff, "\x00")
	if len(buff)%2 != 0 {
		buff = append(buff, 00)
	}
	decoder := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM).NewDecoder()
	bs2, err := decoder.Bytes(buff)
	if err != nil {
		return
	}
	data = string(bs2)
	return
}

func ByteUtf8StrToUtf16le(data string, rLen int) (result []byte, err error) {
	buff := []byte(data)

	decoder := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM).NewEncoder()
	result, err = decoder.Bytes(buff)
	if err != nil {
		return
	}
	bLen := len(result)
	if bLen > rLen {
		err = errors.New("data too long")
		return
	}
	temp := make([]byte, rLen-bLen)
	result = BytesCombine(result, temp)
	return
}

//前四个字符串是mac地址后两位，第五个
func MacAddressAndChannelToSoundStr(macAddress string, channel uint) (result string) {
	macAddressByte := MacAddressToByte(macAddress)
	firstByte := []byte{macAddressByte[4], macAddressByte[5]}
	middleByte := []byte{0x00}
	channelByte := []byte{uint8(channel)}
	result = BytetoHex(BytesCombine(firstByte, middleByte, channelByte))
	return
}

func SoundStrToChannelId(soundStr string) (channelId int64) {
	soundByte := Hextobyte(soundStr)

	channelId = int64(soundByte[3:4][0])

	return
}
