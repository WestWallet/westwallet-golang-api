package westwallet

func checkAPIErrors(apiError string) error {
	errorMap := map[string]error{
		"account_blocked":    &AccountBlockedError{},
		"bad_address":        &BadAddressError{},
		"bad_dest_tag":       &BadDestTagError{},
		"insufficient_funds": &InsufficientFundsError{},
		"max_withdraw_error": &MaxWithdrawError{},
		"min_withdraw_error": &MinWithdrawError{},
		"currency_not_found": &CurrencyNotFoundError{},
		"not_found":          &TransactionNotFoundError{},
	}
	if apiError != "ok" {
		knownError := errorMap[apiError]
		if knownError != nil {
			return knownError
		}
		return &WestWalletAPIError{apiError}
	}
	return nil
}

// GenerateAddress requires an currency code and optional ipnURL and descripton and returns an APIAddresses struct
func (a *APIClient) GenerateAddress(currency string, ipnURL string, description string) (address APIAddress, err error) {
	data := map[string]interface{}{
		"currency":    currency,
		"ipn_url":     ipnURL,
		"description": description,
	}
	err = a.Fetch("POST", "/address/generate", data, &address)
	if err != nil {
		return
	}
	err = checkAPIErrors(address.Error)
	if err != nil {
		return
	}
	return
}

// CreateWithdrawal requires currency, amount, address, destTag, description and returns an APITransaction struct
func (a *APIClient) CreateWithdrawal(currency string, amount string, address string, destTag string, description string) (transaction APITransaction, err error) {
	data := map[string]interface{}{
		"currency":    currency,
		"amount":      amount,
		"address":     address,
		"dest_tag":    destTag,
		"description": description,
	}
	err = a.Fetch("POST", "/wallet/create_withdrawal", data, &transaction)
	if err != nil {
		return
	}
	err = checkAPIErrors(transaction.Error)
	if err != nil {
		return
	}
	return
}

// WalletBalance requires currency and returns an APIBalance struct
func (a *APIClient) WalletBalance(currency string) (balance APIBalance, err error) {
	data := map[string]interface{}{
		"currency": currency,
	}
	err = a.Fetch("GET", "/wallet/balance", data, &balance)
	if err != nil {
		return
	}
	err = checkAPIErrors(balance.Error)
	if err != nil {
		return
	}
	return
}

// WalletBalances requires currency and returns an APIBalance struct
func (a *APIClient) WalletBalances() (balances map[string]string, err error) {
	err = a.Fetch("GET", "/wallet/balances", nil, &balances)

	if err != nil {
		return
	}
	return
}

// TransactionInfo requires id and returns an APITransaction struct
func (a *APIClient) TransactionInfo(id int) (transaction APITransaction, err error) {
	data := map[string]interface{}{
		"id": id,
	}
	err = a.Fetch("POST", "/wallet/transaction", data, &transaction)
	if err != nil {
		return
	}
	err = checkAPIErrors(transaction.Error)
	if err != nil {
		return
	}
	return
}
