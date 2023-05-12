package models

type ArgumentsInput struct {
	Input int `json:"Input"`
}

type TransactionTypeMoveCall struct {
	Package   string        `json:"package"`
	Module    string        `json:"module"`
	Function  string        `json:"function"`
	Arguments []interface{} `json:"arguments"`
}

type TransactionTypeTransferObjects struct {
}

type TransactionType struct {
	MoveCall        TransactionTypeMoveCall `json:"MoveCall,omitempty"`
	TransferObjects []interface{}           `json:"TransferObjects,omitempty"`
}

type TransactionInput struct {
	Type       string      `json:"type"`
	ObjectType string      `json:"objectType"`
	ObjectId   string      `json:"objectId"`
	ValueType  string      `json:"valueType"`
	Value      interface{} `json:"value"`
}

type Transaction struct {
	Data struct {
		Transaction struct {
			Transactions     []TransactionType  `json:"transactions"`
			TransactionInput []TransactionInput `json:"inputs"`
		} `json:"transaction"`
	} `json:"data"`
}

type Mutated struct {
	Owner struct {
		AddressOwner string `json:"AddressOwner"`
	} `json:"owner"`
	Reference struct {
		ObjectId string `json:"objectId"`
	} `json:"reference"`
}

type Effects struct {
	Mutated []Mutated `json:"mutated"`
}

type GetTransactionBlockResponse struct {
	Events      []Events    `json:"events"`
	Transaction Transaction `json:"transaction"`
	Effects     Effects     `json:"effects"`
}
