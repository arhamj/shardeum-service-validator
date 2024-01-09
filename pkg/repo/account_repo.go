package repo

import (
	"github.com/shardeum/service-validator/pkg/models"
	"github.com/shardeum/service-validator/pkg/util"
	"gorm.io/gorm"
)

type AccountsEntryRepo interface {
	GetAccount(accountId string) (*models.AccountsEntry, error)
	GetAccounts(accountIds []string) ([]*models.AccountsEntry, error)
}

type accountsEntryRepo struct {
	db *gorm.DB
}

func NewAccountsEntryRepo(db *gorm.DB) AccountsEntryRepo {
	return &accountsEntryRepo{
		db: db,
	}
}

func (a *accountsEntryRepo) GetAccount(accountId string) (*models.AccountsEntry, error) {
	var accountsEntry models.AccountsEntry
	err := a.db.Where("accountId = ?", accountId).First(&accountsEntry).Error
	if err != nil {
		return nil, err
	}
	return &accountsEntry, nil
}

func (a *accountsEntryRepo) GetAccounts(accountIds []string) ([]*models.AccountsEntry, error) {
	var accountsEntries []*models.AccountsEntry
	// de-duplicate accountIds
	accountIds = util.Deduplicate(accountIds)
	err := a.db.Where("accountId IN ?", accountIds).Find(&accountsEntries).Error
	if err != nil {
		return nil, err
	}
	return accountsEntries, nil
}
