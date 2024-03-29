package models

import (
	"github.com/shoshinsquare/sui-go-sdk/models/sui_json_rpc_types"
	"github.com/shoshinsquare/sui-go-sdk/models/sui_types"
)

type GetObjectRequest struct {
	ObjectID string `json:"objectID"`
}

type AttributeContent struct {
	Type   string `json:"type"`
	Fields struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"fields"`
}

type ObjectData struct {
	BSC struct {
		BcsBytes          string `json:"bcsBytes"`
		DataType          string `json:"dataType"`
		HasPublicTransfer bool   `json:"hasPublicTransfer"`
		Type              string `json:"type"`
		Version           string `json:"version"`
	} `json:"bsc"`
	Digest   string `json:"digest"`
	ObjectID string `json:"objectId"`
	Version  string `json:"version"`
	Type     string `json:"type"`
	Owner    struct {
		AddressOwner string `json:"AddressOwner"`
	} `json:"owner"`
	PreviousTransaction string `json:"previousTransaction"`
	StorageRebate       string `json:"storageRebate"`
	Content             struct {
		DataType          string `json:"dataType"`
		Type              string `json:"type"`
		HasPublicTransfer bool   `json:"hasPublicTransfer"`
		Fields            struct {
			Balance string `json:"balance"`
			Id      struct {
				Id string `json:"id"`
			} `json:"id"`
			Name       string `json:"name"`
			ImageURL   string `json:"image_url"`
			ImgURL     string `json:"img_url"`
			URL        string `json:"url"`
			Link       string `json:"link"`
			Attributes struct {
				Fields struct {
					Contents []AttributeContent `json:"contents"`
					Map      struct {
						Fields struct {
							Content []AttributeContent `json:"contents"`
						} `json:"fields"`
					} `json:"map"`
				} `json:"fields"`
			} `json:"attributes"`
		} `json:"fields"`
	} `json:"content"`
	Display struct {
		Data struct {
			Collection  string `json:"collection"`
			Creator     string `json:"creator"`
			Description string `json:"description"`
			ImageURL    string `json:"image_url"`
			ProjectURL  string `json:"project_url"`
			Name        string `json:"name"`
			Kiosk       string `json:"kiosk"`
		} `json:"data"`
	} `json:"display"`
}

type ErrorObjectResponse struct {
	Code string `json:"code"`
}
type GetObjectResponse struct {
	Data  ObjectData          `json:"data"`
	Error ErrorObjectResponse `json:"error"`
}

type GetMultiObjectRequest struct {
	ObjectIDs []string `json:"objectIDs"`
}

type GetMultiObjectResponse struct {
	Data []GetObjectResponse
}
type GetTransactionBlockRequest struct {
	Digest string `json:"digest"`
}

type Events struct {
	BCS string `json:"bcs"`
	ID  struct {
		TxDigest string `json:"txDigest"`
	} `json:"id"`
	PackageId         string      `json:"packageId"`
	ParsedJson        interface{} `json:"parsedJson"`
	Sender            string      `json:"sender"`
	TransactionModule string      `json:"transactionModule"`
	Type              string      `json:"type"`
}

type GetObjectsOwnedByAddressRequest struct {
	Address string `json:"address"`
}

type GetObjectsOwnedByAddressResponse struct {
	Result []SuiObjectInfo `json:"result"`
}

type GetObjectsOwnedByObjectRequest struct {
	ObjectID string `json:"objectID"`
}
type GetObjectsOwnedByObjectResponse struct {
	Result []SuiObjectInfo `json:"result"`
}

type GetRawObjectRequest struct {
	ObjectID string `json:"objectID"`
}
type GetRawObjectResponse struct {
	Details struct {
		Data sui_json_rpc_types.SuiParsedMoveObject `json:"data"`
		sui_json_rpc_types.OwnedObjectRef
		PreviousTransaction string                 `json:"previousTransaction"`
		StorageRebate       uint64                 `json:"storageRebate"`
		Reference           sui_types.SuiObjectRef `json:"reference"`
	} `json:"details"`
	Status string `json:"status"`
}

type TryGetPastObjectRequest struct {
	ObjectID string `json:"objectID"`
	Version  string `json:"version"`
}

type TryGetPastObjectResponse struct {
	Status  string `json:"status"`
	Details struct {
		Data sui_json_rpc_types.SuiParsedMoveObject `json:"data"`
		sui_json_rpc_types.OwnedObjectRef
		PreviousTransaction string                 `json:"previousTransaction"`
		StorageRebate       uint64                 `json:"storageRebate"`
		Reference           sui_types.SuiObjectRef `json:"reference"`
	} `json:"details"`
}

type GetCoinMetadataRequest struct {
	CoinType string `json:"coin_type"`
}

type GetCoinMetadataResponse struct {
	Decimals    uint8  `json:"decimals"`
	Description string `json:"description"`
	IconUrl     string `json:"iconUrl,omitempty"`
	Id          string `json:"id"`
	Name        string `json:"name"`
	Symbol      string `json:"symbol"`
}

type GetDynamicFieldObejctRequest struct {
	ParentObjectID string
	Name           string
}

type GetDynamicFieldObjectResponse struct {
	Details struct {
		Data sui_json_rpc_types.SuiParsedMoveObject `json:"data"`
		sui_json_rpc_types.OwnedObjectRef
		PreviousTransaction string                 `json:"previousTransaction"`
		StorageRebate       uint64                 `json:"storageRebate"`
		Reference           sui_types.SuiObjectRef `json:"reference"`
	} `json:"details"`
	Status string `json:"status"`
}

type GetOwnedObjectsRequest struct {
	Address string `json:"address"`
	Cursor  *string
	Limit   uint64
}

type GetOwnedObjectsResponse struct {
	Data        []sui_json_rpc_types.SuiParsedMoveObject `json:"data"`
	HasNextPage bool                                     `json:"hasNextPage"`
	NextCursor  string                                   `json:"nextCursor"`
}

type GetDynamicFieldRequest struct {
	ParentObjectID string  `json:"parent_object_id"`
	Cursor         *string `json:"cursor"`
}

type DynamicFieldData struct {
	Name struct {
		Type  string `json:"type"`
		Value struct {
			Id string `json:"id"`
		} `json:"value"`
	} `json:"name"`
	BCSName    string `json:"bcsName"`
	Type       string `json:"type"`
	ObjectType string `json:"objectType"`
	ObjectId   string `json:"objectId"`
	Version    int64  `json:"version"`
	Digest     string `json:"digest"`
}

type GetDynamicFieldResponse struct {
	HasNextPage bool               `json:"hasNextPage"`
	NextCursor  string             `json:"nextCursor"`
	Data        []DynamicFieldData `json:"data"`
}
