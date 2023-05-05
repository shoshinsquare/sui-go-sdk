package models

import (
	"github.com/shoshinsquare/sui-go-sdk/models/sui_types"
)

type MoveCallRequest struct {
	Signer          string        `json:"signer"`
	PackageObjectId string        `json:"packageObjectId"`
	Module          string        `json:"module"`
	Function        string        `json:"function"`
	TypeArguments   interface{}   `json:"typeArguments"`
	Arguments       []interface{} `json:"arguments"`
	Gas             *string       `json:"gas"`
	GasBudget       string        `json:"gasBudget"`
	ExecutionMode   *string       `json:"execution_mode"`
}

type MoveCallResponse struct {
	Gas          []sui_types.SuiObjectRef `json:"gas"`
	InputObjects interface{}              `json:"inputObjects"`
	TxBytes      string                   `json:"txBytes"`
}

type MergeCoinsRequest struct {
	Signer      string `json:"signer"`
	PrimaryCoin string `json:"primaryCoin"`
	CoinToMerge string `json:"coinToMerge"`
	Gas         string `json:"gas"`
	GasBudget   string `json:"gasBudget"`
}

type MergeCoinsResponse struct {
	TransactionBytes string                 `json:"transactionBytes"`
	Gas              sui_types.SuiObjectRef `json:"gas"`
	InputObject      interface{}            `json:"inputObject"`
	TxBytes          string                 `json:"txBytes"`
}

type SplitCoinRequest struct {
	Signer       string   `json:"signer"`
	CoinObjectId string   `json:"coinObjectId"`
	SplitAmounts []string `json:"splitAmounts"`
	Gas          *string  `json:"gas"`
	GasBudget    string
}

type SplitCoinResponse struct {
	TransactionBytes string                   `json:"transactionBytes"`
	Gas              []sui_types.SuiObjectRef `json:"gas"`
	InputObject      interface{}              `json:"inputObject"`
	TxBytes          string                   `json:"txBytes"`
}

type SplitCoinEqualRequest struct {
	Signer       string  `json:"signer"`
	CoinObjectId string  `json:"coinObjectId"`
	SplitCount   string  `json:"splitCount"`
	Gas          *string `json:"gas"`
	GasBudget    string  `json:"gasBudget"`
}

type SplitCoinEqualResponse struct {
	TransactionBytes string                   `json:"transactionBytes"`
	Gas              []sui_types.SuiObjectRef `json:"gas"`
	InputObject      interface{}              `json:"inputObject"`
	TxBytes          string                   `json:"txBytes"`
}

type PublishRequest struct {
	Sender          string   `json:"sender"`
	CompiledModules []string `json:"compiledModules"`
	Gas             string   `json:"gas"`
	GasBudget       uint64   `json:"gasBudget"`
}

type PublishResponse struct {
	TransactionBytes string                 `json:"transactionBytes"`
	Gas              sui_types.SuiObjectRef `json:"gas"`
	InputObject      interface{}            `json:"inputObject"`
	TxBytes          string                 `json:"txBytes"`
}

type TransferObjectRequest struct {
	Signer    string `json:"signer"`
	ObjectId  string `json:"objectId"`
	Gas       string `json:"gas"`
	GasBudget uint64 `json:"gasBudget"`
	Recipient string `json:"recipient"`
}

type TransferObjectResponse struct {
	TransactionBytes string                 `json:"transactionBytes"`
	Gas              sui_types.SuiObjectRef `json:"gas"`
	InputObject      interface{}            `json:"inputObject"`
	TxBytes          string                 `json:"txBytes"`
}

type TransferSuiRequest struct {
	Signer      string `json:"signer"`
	SuiObjectId string `json:"suiObjectId"`
	GasBudget   string `json:"gasBudget"`
	Recipient   string `json:"recipient"`
	Amount      string `json:"amount"`
}

type TransferSuiResponse struct {
	TransactionBytes string                   `json:"transactionBytes"`
	Gas              []sui_types.SuiObjectRef `json:"gas"`
	InputObject      interface{}              `json:"inputObject"`
	TxBytes          string                   `json:"txBytes"`
}

type BatchTransactionRequest struct {
	Signer                  string                    `json:"signer"`
	SingleTransactionParams []SingleTransactionParams `json:"singleTransactionParams"`
	Gas                     *string                   `json:"gas"`
	GasBudget               string                    `json:"gasBudget"`
}

type BatchTransactionResponse struct {
	TransactionBytes string                   `json:"transactionBytes"`
	Gas              []sui_types.SuiObjectRef `json:"gas"`
	InputObject      interface{}              `json:"inputObject"`
	TxBytes          string                   `json:"txBytes"`
}

type SingleTransactionParams struct {
	MoveCallRequestParams       *MoveCallRequest       `json:"moveCallRequestParams,omitempty"`
	TransferObjectRequestParams *TransferObjectRequest `json:"transferObjectRequestParams,omitempty"`
}

type ExecuteTransactionRequest struct {
	TxBytes   string   `json:"txBytes"`
	Signature []string `json:"signature"`
}

type ExecuteTransactionResponse struct {
	Digest         string      `json:"digest"`
	RawTransaction string      `json:"rawTransaction"`
	Transaction    interface{} `json:"transaction"`
	Effects        struct {
		Status struct {
			Status string `json:"status"`
		} `json:"status"`
	} `json:"effects"`
	Events                  interface{} `json:"events"`
	ObjectChanges           interface{} `json:"objectChanges"`
	BalanceChanges          interface{} `json:"balanceChanges"`
	ConfirmedLocalExecution bool        `json:"confirmedLocalExecution"`
}

type DryRunTransactionRequest struct {
	TxBytes   string `json:"txBytes"`
	SigScheme string `json:"sigScheme"`
	Signature string `json:"signature"`
	PubKey    string `json:"pubKey"`
}

type DryRunTransactionResponse struct {
	TransactionBytes string                 `json:"transactionBytes"`
	Gas              sui_types.SuiObjectRef `json:"gas"`
	InputObject      interface{}            `json:"inputObject"`
	TxBytes          string                 `json:"txBytes"`
}

type PayRequest struct {
	Signer     string   `json:"signer"`
	InputCoins []string `json:"inputCoins"`
	Recipient  []string `json:"recipient"`
	Amounts    []string `json:"amounts"`
	Gas        string   `json:"gas"`
	GasBudget  uint64   `json:"gasBudget"`
}

type PayResponse struct {
	TransactionBytes string                 `json:"transactionBytes"`
	Gas              sui_types.SuiObjectRef `json:"gas"`
	InputObject      interface{}            `json:"inputObject"`
	TxBytes          string                 `json:"txBytes"`
}

type PayAllSuiRequest struct {
	Signer     string   `json:"suiAddress,omitempty"`
	InputCoins []string `json:"inputCoins,omitempty"`
	Recipient  string   `json:"recipient,omitempty"`
	GasBudget  uint64   `json:"gasBudget,omitempty"`
}

type PayAllSuiResponse struct {
	TransactionBytes string                 `json:"transactionBytes"`
	Gas              sui_types.SuiObjectRef `json:"gas"`
	InputObject      interface{}            `json:"inputObject"`
	TxBytes          string                 `json:"txBytes"`
}

type PaySuiRequest struct {
	Signer     string   `json:"signer,omitempty"`
	InputCoins []string `json:"inputCoins,omitempty"`
	Recipient  []string `json:"recipient,omitempty"`
	GasBudget  uint64   `json:"gasBudget,omitempty"`
}

type PaySuiResponse struct {
	TransactionBytes string                 `json:"transactionBytes"`
	Gas              sui_types.SuiObjectRef `json:"gas"`
	InputObject      interface{}            `json:"inputObject"`
	TxBytes          string                 `json:"txBytes"`
}

type MintNFTRequest struct {
	Signer         string
	NFTName        string
	NFTDescription string
	NFTUrl         string
	GasObject      string
	GasBudget      uint64
}

type ExecuteTransactionSerializedSigRequest struct {
	TxBytes     string
	Signature   string
	RequestType interface{}
}

type ExecuteTransactionSerializedSigResponse struct {
	EffectsCert interface{} `json:"effectsCert"`
}
