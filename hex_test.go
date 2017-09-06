package hex

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"

)

func TestIpv4Proper(t *testing.T) {
	assert := assert.New(t)
	//assert proper hex format works
	ip := decode("AAAAAAAA")
	assert.Equal(ip.address, "170.170.170.170", "These addresses should be the same")
	assert.Nil(ip.errorMessage)
	assert.Equal(ip.addressType, IPv4)
}

//Asserts that decode will append zeroes to keep proper formatting for ip
func TestIpv4WithFill(t *testing.T) {
	assert := assert.New(t)
	ip := decode("AAAAAA")
	assert.Equal(ip.address, "170.170.170.0", "These addresses should be the same")
	assert.Nil(ip.errorMessage)
	assert.Equal(ip.addressType, IPv4)
}

func TestIpv6Withfill(t *testing.T) {
	assert := assert.New(t)

	ip := decode("AAAAAAAAAAAAAAAA")
	assert.Equal("170.170.170.170.170.170.170.170.0.0.0.0.0.0.0.0", ip.address)
	assert.Equal(ip.addressType, IPv6)
	assert.Nil(ip.errorMessage)
}

func TestIpv6Proper(t *testing.T) {
	assert := assert.New(t)

	ip := decode("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	assert.Equal(ip.address, "170.170.170.170.170.170.170.170.170.170.170.170.170.170.170.170")
	assert.Equal(ip.addressType, IPv6)
	assert.Nil(ip.errorMessage)
}

func TestInvalidChars(t *testing.T) {
	ip := decode("hHHH")
	assert.Equal(t, ip.addressType, invalIP)
	assert.Equal(t, "", ip.address)
	assert.NotNil(t, ip.errorMessage)
	testMsg := errors.New(warnIllegalChars)
	assert.Equal(t, testMsg, ip.errorMessage)
}

func TestInvalidOverflow(t *testing.T) {
	assert := assert.New(t)
	ip := decode("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	assert.Equal(ip.addressType, invalIP)
	testErr := errors.New(warnOversizedAddr)
	assert.Equal(testErr, ip.errorMessage)
}

func TestInvalidEmpty(t *testing.T) {
	assert := assert.New(t)
	ip := decode("")

	testErr := errors.New(warnEmpty)
	assert.Equal("", ip.address)
	assert.Equal(invalIP, ip.addressType)
	assert.Equal(testErr, ip.errorMessage)
}
