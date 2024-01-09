package dto

type GetEVMAccountRequest struct {
	AddressID string `param:"accountId" validate:"required"`
}

type GetEVMAccountResponse struct {
	Account struct {
		Balance     string `json:"balance"`
		Nonce       string `json:"nonce"`
		CodeHash    string `json:"codeHash"`
		StorageRoot string `json:"storageRoot"`
	} `json:"account"`
}

type GetCodeRequest struct {
	AddressID string `query:"address" validate:"required"`
}

type GetCodeResponse struct {
	ContractCode string `json:"contractCode"`
}
