package service

import (
	"github.com/shardeum/service-validator/pkg/dto"
	"github.com/shardeum/service-validator/pkg/repo"
)

type AccountService interface {
	GetAccount(accountId string) (*dto.AccountDto, error)
	GetAccounts(accountIds []string) ([]*dto.AccountDto, error)
}

type accountService struct {
	accountRepo repo.AccountsEntryRepo
}

func NewAccountService(accountRepo repo.AccountsEntryRepo) AccountService {
	return &accountService{
		accountRepo: accountRepo,
	}
}

func (a *accountService) GetAccount(accountId string) (*dto.AccountDto, error) {
	accountsEntry, err := a.accountRepo.GetAccount(accountId)
	if err != nil {
		return nil, err
	}

	var result dto.AccountDto
	result.FromModel(accountsEntry)

	return &result, nil
}

func (a *accountService) GetAccounts(accountIds []string) ([]*dto.AccountDto, error) {
	accountEntries, err := a.accountRepo.GetAccounts(accountIds)
	if err != nil {
		return nil, err
	}

	accounts := make([]*dto.AccountDto, len(accountEntries))
	for i, accountEntry := range accountEntries {
		var account dto.AccountDto
		account.FromModel(accountEntry)
		accounts[i] = &account
	}

	return accounts, nil
}
