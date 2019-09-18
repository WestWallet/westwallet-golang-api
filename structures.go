package westwallet

type APIAddress struct {
	Address  string
	DestTag  string
	Currency string
	Label    string
	Error    string
}

type APIBalance struct {
	Balance  string
	Currency string
	Error    string
}

type APITransaction struct {
	ID                      int
	Amount                  string
	Address                 string
	DestTag                 string
	Currency                string
	Status                  string
	Type                    string
	BlockchainConfirmations int
	BlockchainHash          string
	Error                   string
}
