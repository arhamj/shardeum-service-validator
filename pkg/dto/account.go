package dto

import (
	"encoding/base64"
	"encoding/json"

	"github.com/shardeum/service-validator/pkg/models"
)

type GetAccountRequest struct {
	AccountId string `param:"accountId"`
}

type GetAccountsRequest struct {
	AccountIds []string `json:"accountIds"`
}

type AccountDto struct {
	AccountId string                 `json:"accountId"`
	Timestamp int64                  `json:"timestamp"`
	Data      map[string]interface{} `json:"data"`
}

type ParsedAccountData struct {
	AccountType int64      `json:"accountType"`
	EthAddress  string     `json:"ethAddress"`
	Hash        string     `json:"hash"`
	Timestamp   int64      `json:"timestamp"`
	Account     EVMAccount `json:"account"`
	CodeByte    BufferType `json:"codeByte"`
}

type EVMAccount struct {
	Balance     string         `json:"balance"`
	Nonce       string         `json:"nonce"`
	CodeHash    map[string]int `json:"codeHash"`
	StorageRoot map[string]int `json:"storageRoot"`
}

type BufferType struct {
	Data     string `json:"data"`
	DataType string `json:"dataType"`
}

func (b BufferType) ExtractBytes() []byte {
	switch b.DataType {
	case "bh":
		decodedBytes, _ := base64.StdEncoding.DecodeString(b.Data)
		return decodedBytes
	}
	return []byte{}
}

func (a AccountDto) GetParsedAccount() (*ParsedAccountData, error) {
	dataBytes, err := json.Marshal(a.Data)
	if err != nil {
		return nil, err
	}
	var parsedAccount ParsedAccountData
	err = json.Unmarshal(dataBytes, &parsedAccount)
	if err != nil {
		return nil, err
	}
	return &parsedAccount, nil
}

func (a *AccountDto) FromModel(model *models.AccountsEntry) {
	a.AccountId = model.AccountId
	a.Timestamp = model.Timestamp
	a.Data = model.Data
}
