package westwallet

type APIAddress struct {
	Address  string `json:"address"`
	DestTag  string `json:"dest_tag"`
	Currency string `json:"currency"`
	Label    string `json:"label"`
	Error    string `json:"error"`
}

type APIBalance struct {
	Balance  string `json:"balance"`
	Currency string `json:"currency"`
	Error    string `json:"error"`
}

type APITransaction struct {
	ID                      int    `json:"id"`
	Amount                  string `json:"amount"`
	Address                 string `json:"address"`
	DestTag                 string `json:"dest_tag"`
	Currency                string `json:"currency"`
	Status                  string `json:"status"`
	Type                    string `json:"type"`
	BlockchainConfirmations int    `json:"blockchain_confirmations"`
	BlockchainHash          string `json:"blockchain_hash"`
	Error                   string `json:"error"`
}
