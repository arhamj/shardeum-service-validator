package util

import (
	"encoding/hex"
	"errors"
	"math/big"
	"strings"

	"github.com/shardeum/service-validator/pkg/constants/enums"
	log "github.com/sirupsen/logrus"
)

func ToShardusAddress(address, secondaryAddressStr string, accountType enums.AccountType) (string, error) {
	var result string
	switch accountType {
	case enums.Account:
		if len(address) != 42 {
			return "", errors.New("invalid address length")
		}
		result = address[2:] + "000000000000000000000000"
	case enums.ContractCode:
		if len(secondaryAddressStr) == 66 {
			secondaryAddressStr = secondaryAddressStr[2:]
		}
		result = secondaryAddressStr
	default:
		result = ""
	}
	return strings.ToLower(result), nil
}

func BytesToHex(bytes []byte) string {
	return "0x" + strings.ToLower(hex.EncodeToString(bytes))
}

func HexToBigInt(hexString string) *big.Int {
	bigInt := new(big.Int)
	_, ok := bigInt.SetString(hexString, 16)
	if !ok {
		log.Error("Failed to parse hex string")
	}
	return bigInt
}
