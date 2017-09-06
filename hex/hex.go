package hex

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
)

type IpMessage struct {
	address      string
	addressType  string
	errorMessage error
}

func (ip *IpMessage) GetIP() string{
  return ip.address
}

func (ip *IpMessage) GetAddrType() string{
  return ip.addressType
}

func (ip *IpMessage) GetErrMsg() error{
  return ip.errorMessage
}


const (
	invalIP           = "This is an invalid IP address"
	IPv4              = "This is an IPv4 address"
	IPv6              = "This is an IPv6 address"
	warnEmpty         = "Cannot process an empty address"
	warnOversizedAddr = "Ip address exceeds maximum byte length"
	warnIllegalChars  = "Illegal characters in this request"
)

/*Custom error message designed to inform user of invalid ips and other warnings*/
type errorAddress struct {
	msg string
}

/*Returns desired error message*/
func (e *errorAddress) Error() string {
	return e.msg
}

/*Used to generate a new error message*/
func New(text string) error {
	return &errorAddress{text}
}

/*Decoding manager, receives input string format and returns a formatted ip
* along with information about what kind of ip/validity of passed ip
 */
func Decode(str string) *IpMessage {
	size := len(str)
	var ipType string
	switch {
	case size == 0:
		ipType = invalIP
		return &IpMessage{address: "", addressType: ipType, errorMessage: errors.New(warnEmpty)}
	case size > 32:
		ipType = invalIP
		return &IpMessage{address: "", addressType: ipType, errorMessage: errors.New(warnOversizedAddr)}
	case size < 8:
		str = appendZeros(str, (8 - size))
		ipType = IPv4
	case size > 8:
		if size < 32 {
			str = appendZeros(str, (32 - size))
		}
		ipType = IPv6
	default:
		ipType = IPv4
	}
	arr, err := hex.DecodeString(str)
	if err != nil {
		return &IpMessage{address: "", addressType: invalIP, errorMessage: errors.New(warnIllegalChars)}
	}
	str = formatIP(arr)
	return &IpMessage{address: str, addressType: ipType, errorMessage: err}

}

/*Pads zeroes at the end of ip to properly format as ipv4 or ipv6 message*/
func appendZeros(str string, num int) string {
	var buffer bytes.Buffer
	buffer.WriteString(str)
	for i := 0; i < num; i++ {
		buffer.WriteString("0")
	}
	return buffer.String()
}

/*Performs the job of formatting the ip address based on type and size*/
func formatIP(arr []byte) string {
	var buffer bytes.Buffer
	for i := 0; i < len(arr)-1; i++ {
		buffer.WriteString(fmt.Sprintf("%d.", arr[i]))
	}
	buffer.WriteString(fmt.Sprintf("%d", arr[len(arr)-1]))
	return buffer.String()
}
