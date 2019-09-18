package westwallet

type WestWalletAPIError struct {
	S string
}

func (e *WestWalletAPIError) Error() string {
	return "Undefined api error: " + e.S
}

type InsufficientFundsError struct {
}

func (e *InsufficientFundsError) Error() string {
	return "You don't have enough funds for this withdrawal"
}

type CurrencyNotFoundError struct {
}

func (e *CurrencyNotFoundError) Error() string {
	return "This currency doesn't exist in WestWallet"
}

type NotAllowedError struct {
}

func (e *NotAllowedError) Error() string {
	return "This request is not allowed, check ip settings of your key"
}

type WrongCredentialsError struct {
}

func (e *WrongCredentialsError) Error() string {
	return "API key or secret wrong"
}

type TransactionNotFoundError struct {
}

func (e *TransactionNotFoundError) Error() string {
	return "No such transaction"
}

type AccountBlockedError struct {
}

func (e *AccountBlockedError) Error() string {
	return "Account is blocked"
}

type BadAddressError struct {
}

func (e *BadAddressError) Error() string {
	return "Wrong recepient's address"
}

type BadDestTagError struct {
}

func (e *BadDestTagError) Error() string {
	return "Wrong recepient's destination tag"
}

type MinWithdrawError struct {
}

func (e *MinWithdrawError) Error() string {
	return "Withdrawal amount is less then allowed"
}

type MaxWithdrawError struct {
}

func (e *MaxWithdrawError) Error() string {
	return "Withdrawal amount is greater then allowed"
}
