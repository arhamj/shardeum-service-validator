package service

import (
	"github.com/arhamj/go-commons/pkg/http_errors"
	"github.com/shardeum/service-validator/pkg/constants/enums"
	"github.com/shardeum/service-validator/pkg/dto"
	"github.com/shardeum/service-validator/pkg/util"
)

type ValidatorService interface {
	GetEVMAccount(address string) (*dto.GetEVMAccountResponse, error)
	GetCodeBytes(address string) (*dto.GetCodeResponse, error)
}

type validatorService struct {
	accountService AccountService
}

func NewValidatorService(accountService AccountService) ValidatorService {
	return validatorService{
		accountService: accountService,
	}
}

func (v validatorService) GetEVMAccount(address string) (*dto.GetEVMAccountResponse, error) {
	shardusAddress, err := util.ToShardusAddress(address, "", enums.Account)
	if err != nil {
		return nil, err
	}
	account, err := v.accountService.GetAccount(shardusAddress)
	if err != nil {
		return nil, err
	}
	parsedAccountData, err := account.GetParsedAccount()
	if err != nil {
		return nil, err
	}
	if parsedAccountData.AccountType != int64(enums.Account) {
		return nil, http_errors.NewBadRequestError(nil, "account type is not an EOA", true)
	}
	res := dto.GetEVMAccountResponse{
		Account: struct {
			Balance     string `json:"balance"`
			Nonce       string `json:"nonce"`
			CodeHash    string `json:"codeHash"`
			StorageRoot string `json:"storageRoot"`
		}{
			Balance:     util.HexToBigInt(parsedAccountData.Account.Balance).String(),
			Nonce:       util.HexToBigInt(parsedAccountData.Account.Nonce).String(),
			CodeHash:    util.BytesToHex(parsedAccountData.Account.CodeHash.ExtractBytes()),
			StorageRoot: util.BytesToHex(parsedAccountData.Account.StorageRoot.ExtractBytes()),
		},
	}
	return &res, nil
}

func (v validatorService) GetCodeBytes(address string) (*dto.GetCodeResponse, error) {
	shardusAddress, err := util.ToShardusAddress(address, "", enums.Account)
	if err != nil {
		return nil, err
	}
	account, err := v.accountService.GetAccount(shardusAddress)
	if err != nil {
		return nil, err
	}
	parsedAccountData, err := account.GetParsedAccount()
	if err != nil {
		return nil, err
	}
	codeHashKey := util.BytesToHex(parsedAccountData.Account.CodeHash.ExtractBytes())
	codeBytesShardusAddress, err := util.ToShardusAddress(address, codeHashKey, enums.ContractCode)
	if err != nil {
		return nil, err
	}
	codeBytesAccount, err := v.accountService.GetAccount(codeBytesShardusAddress)
	if err != nil {
		return nil, err
	}
	parsedCodeBytesAccount, err := codeBytesAccount.GetParsedAccount()
	if err != nil {
		return nil, err
	}
	res := dto.GetCodeResponse{
		ContractCode: util.BytesToHex(parsedCodeBytesAccount.CodeByte.ExtractBytes()),
	}
	return &res, nil
}
