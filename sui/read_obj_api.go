package sui

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	"github.com/tidwall/gjson"

	"github.com/shoshinsquare/sui-go-sdk/common/rpc_client"
	"github.com/shoshinsquare/sui-go-sdk/models"
)

type IReadObjectFromSuiAPI interface {
	GetObject(ctx context.Context, req models.GetObjectRequest, opts ...interface{}) (models.GetObjectResponse, error)
	GetMultiObject(ctx context.Context, req models.GetMultiObjectRequest, opts ...interface{}) (models.GetMultiObjectResponse, error)
	GetRawObject(ctx context.Context, req models.GetRawObjectRequest, opts ...interface{}) (models.GetRawObjectResponse, error)
	TryGetPastObject(ctx context.Context, req models.TryGetPastObjectRequest, opt ...interface{}) (models.TryGetPastObjectResponse, error)
	GetCoinMetadata(ctx context.Context, req models.GetCoinMetadataRequest, opt ...interface{}) (models.GetCoinMetadataResponse, error)
	GetOwnedObjects(ctx context.Context, req models.GetOwnedObjectsRequest, opt ...interface{}) (models.GetOwnedObjectsResponse, error)
	GetDynamicField(ctx context.Context, req models.GetDynamicFieldRequest, opt ...interface{}) (models.GetDynamicFieldResponse, error)
	GetAllNFT(ctx context.Context, address string) ([]models.GetObjectResponse, error)
	GetTransactionBlock(ctx context.Context, req models.GetTransactionBlockRequest, opt ...interface{}) (models.GetTransactionBlockResponse, error)
}

type suiReadObjectFromSuiImpl struct {
	cli *rpc_client.RPCClient
}

func (s *suiReadObjectFromSuiImpl) GetTransactionBlock(ctx context.Context, req models.GetTransactionBlockRequest, opt ...interface{}) (models.GetTransactionBlockResponse, error) {
	var rsp models.GetTransactionBlockResponse
	respBytes, err := s.cli.Request(ctx, models.Operation{
		JsonRPC: "2.0",
		ID:      1,
		Method:  "sui_getTransactionBlock",
		Params: []interface{}{
			req.Digest,
			map[string]bool{
				"showInput":          true,
				"showRawInput":       true,
				"showEffects":        true,
				"showEvents":         true,
				"showObjectChanges":  true,
				"showBalanceChanges": true,
			},
		},
	})
	if err != nil {
		return models.GetTransactionBlockResponse{}, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return models.GetTransactionBlockResponse{}, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp)
	if err != nil {
		return models.GetTransactionBlockResponse{}, err
	}
	return rsp, nil
}

func (s *suiReadObjectFromSuiImpl) GetMultiObject(ctx context.Context, req models.GetMultiObjectRequest, opts ...interface{}) (models.GetMultiObjectResponse, error) {
	respBytes, err := s.cli.Request(ctx, models.Operation{
		JsonRPC: "2.0",
		ID:      1,
		Method:  "sui_multiGetObjects",
		Params: []interface{}{
			req.ObjectIDs,
			map[string]bool{
				"showType":                true,
				"showOwner":               true,
				"showPreviousTransaction": true,
				"showDisplay":             true,
				"showContent":             true,
				"showBcs":                 true,
				"showStorageRebate":       true,
			},
		},
	})
	if err != nil {
		return models.GetMultiObjectResponse{}, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return models.GetMultiObjectResponse{}, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	var rsp models.GetMultiObjectResponse
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp.Data)
	if err != nil {
		return models.GetMultiObjectResponse{}, err
	}
	return rsp, nil
}

// GetObject implements method `sui_getObject`.
// Returns object details
func (s *suiReadObjectFromSuiImpl) GetObject(ctx context.Context, req models.GetObjectRequest, opts ...interface{}) (models.GetObjectResponse, error) {

	respBytes, err := s.cli.Request(ctx, models.Operation{
		JsonRPC: "2.0",
		ID:      1,
		Method:  "sui_getObject",
		Params: []interface{}{
			req.ObjectID,
			map[string]bool{
				"showType":                true,
				"showOwner":               true,
				"showPreviousTransaction": true,
				"showDisplay":             true,
				"showContent":             true,
				"showBcs":                 true,
				"showStorageRebate":       true,
			},
		},
	})
	if err != nil {
		return models.GetObjectResponse{}, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return models.GetObjectResponse{}, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	var rsp models.GetObjectResponse
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp)
	if err != nil {
		return models.GetObjectResponse{}, err
	}
	return rsp, nil
}

// GetObjectsOwnedByAddress implements method `sui_getObjectsOwnedByAddress`.
// Returns an array of object information
func (s *suiReadObjectFromSuiImpl) GetObjectsOwnedByAddress(ctx context.Context, req models.GetObjectsOwnedByAddressRequest, opts ...interface{}) (models.GetObjectsOwnedByAddressResponse, error) {
	var rsp models.GetObjectsOwnedByAddressResponse
	respBytes, err := s.cli.Request(ctx, models.Operation{
		JsonRPC: "2.0",
		ID:      1,
		Method:  "sui_getObjectsOwnedByAddress",
		Params: []interface{}{
			req.Address,
		},
	})
	if err != nil {
		return models.GetObjectsOwnedByAddressResponse{}, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return models.GetObjectsOwnedByAddressResponse{}, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp.Result)
	if err != nil {
		return models.GetObjectsOwnedByAddressResponse{}, err
	}
	return rsp, nil
}

// GetRawObject implements method `sui_getRawObject`.
// Returns object details
func (s *suiReadObjectFromSuiImpl) GetRawObject(ctx context.Context, req models.GetRawObjectRequest, opts ...interface{}) (models.GetRawObjectResponse, error) {
	var rsp models.GetRawObjectResponse
	respBytes, err := s.cli.Request(ctx, models.Operation{
		JsonRPC: "2.0",
		ID:      1,
		Method:  "sui_getRawObject",
		Params: []interface{}{
			req.ObjectID,
		},
	})
	if err != nil {
		return models.GetRawObjectResponse{}, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return models.GetRawObjectResponse{}, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp)
	if err != nil {
		return models.GetRawObjectResponse{}, err
	}
	return rsp, nil
}

// TryGetPastObject implements method `sui_tryGetPastObject`
// Note there is no software-level guarantee/SLA that objects with past versions can be retrieved by this API,
// even if the object and version exists/existed.
// The result may vary across nodes depending on their pruning policies.
// Return the object information for a specified version
func (s *suiReadObjectFromSuiImpl) TryGetPastObject(ctx context.Context, req models.TryGetPastObjectRequest, opts ...interface{}) (models.TryGetPastObjectResponse, error) {
	var rsp models.TryGetPastObjectResponse
	respBytes, err := s.cli.Request(ctx, models.Operation{
		JsonRPC: "2.0",
		ID:      1,
		Method:  "sui_tryGetPastObject",
		Params: []interface{}{
			req.ObjectID,
		},
	})
	if err != nil {
		return models.TryGetPastObjectResponse{}, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return models.TryGetPastObjectResponse{}, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp)
	if err != nil {
		return models.TryGetPastObjectResponse{}, err
	}
	return rsp, nil
}

func (s *suiReadObjectFromSuiImpl) GetCoinMetadata(ctx context.Context, req models.GetCoinMetadataRequest, opt ...interface{}) (models.GetCoinMetadataResponse, error) {
	var rsp models.GetCoinMetadataResponse
	respBytes, err := s.cli.Request(ctx, models.Operation{
		JsonRPC: "2.0",
		ID:      1,
		Method:  "suix_getCoinMetadata",
		Params: []interface{}{
			req.CoinType,
		},
	})
	if err != nil {
		return models.GetCoinMetadataResponse{}, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return models.GetCoinMetadataResponse{}, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp)
	if err != nil {
		return models.GetCoinMetadataResponse{}, err
	}
	return rsp, nil
}

func (s *suiReadObjectFromSuiImpl) GetOwnedObjects(ctx context.Context, req models.GetOwnedObjectsRequest, opt ...interface{}) (models.GetOwnedObjectsResponse, error) {
	var rsp models.GetOwnedObjectsResponse
	respBytes, err := s.cli.Request(ctx, models.Operation{
		JsonRPC: "2.0",
		ID:      1,
		Method:  "suix_getOwnedObjects",
		Params: []interface{}{
			req.Address,
			map[string]bool{
				"showType":                true,
				"showOwner":               true,
				"showPreviousTransaction": true,
				"showDisplay":             true,
				"showContent":             true,
				"showBcs":                 true,
				"showStorageRebate":       true,
			},
			req.Cursor,
			req.Limit,
		},
	})
	if err != nil {
		return models.GetOwnedObjectsResponse{}, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return models.GetOwnedObjectsResponse{}, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp)
	if err != nil {
		return models.GetOwnedObjectsResponse{}, err
	}

	return rsp, nil
}

func (s *suiReadObjectFromSuiImpl) GetDynamicField(ctx context.Context, req models.GetDynamicFieldRequest, opt ...interface{}) (models.GetDynamicFieldResponse, error) {
	var rsp models.GetDynamicFieldResponse
	respBytes, err := s.cli.Request(ctx, models.Operation{
		JsonRPC: "2.0",
		ID:      1,
		Method:  "suix_getDynamicFields",
		Params: []interface{}{
			req.ParentObjectID,
			req.Cursor,
			50,
		},
	})
	if err != nil {
		return models.GetDynamicFieldResponse{}, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return models.GetDynamicFieldResponse{}, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp)
	if err != nil {
		return models.GetDynamicFieldResponse{}, err
	}

	return rsp, nil
}

func (s *suiReadObjectFromSuiImpl) GetAllNFT(ctx context.Context, address string) ([]models.GetObjectResponse, error) {

	// get all object id
	objectIDs := []string{}

	var cursor *string = nil
	for {

		res, err := s.GetOwnedObjects(context.Background(), models.GetOwnedObjectsRequest{
			Address: address,
			Cursor:  cursor,
			Limit:   20,
		})
		if err != nil {
			continue
		}

		for _, r := range res.Data {
			objectIDs = append(objectIDs, r.Data.ObjectID)
		}

		cursor = &res.NextCursor

		if !res.HasNextPage {
			break
		}
	}

	response := []models.GetObjectResponse{}
	kiosks := []string{}

	batchsize := 1
	if len(objectIDs) > batchsize {
		index := 0
		step := 0
		for {
			min := index * batchsize
			max := (index + 1) * batchsize

			if step+batchsize > len(objectIDs) {
				max = len(objectIDs)
			}

			queryObject := objectIDs[min:max]

			// query multi object
			resMulti, err := s.GetMultiObject(ctx, models.GetMultiObjectRequest{
				ObjectIDs: queryObject,
			})
			if err != nil {
				if max == len(objectIDs) {
					break
				}

				index += 1
				step = index * batchsize
				continue
			}

			for _, object := range resMulti.Data {
				if strings.Contains(object.Data.Type, "kiosk") {
					if object.Data.Display.Data.Kiosk != "" {
						kiosks = append(kiosks, object.Data.Display.Data.Kiosk)
					}
				} else {
					typeSplit := strings.Split(object.Data.Type, "::")

					if len(typeSplit) == 3 && len(typeSplit[0]) == 66 {
						response = append(response, object)
					}
				}

			}

			if max == len(objectIDs) {
				break
			}

			index += 1
			step = index * batchsize
		}
	} else {
		// query multi object
		resMulti, err := s.GetMultiObject(ctx, models.GetMultiObjectRequest{
			ObjectIDs: objectIDs,
		})
		if err != nil {
			return nil, err
		}

		for _, object := range resMulti.Data {
			if strings.Contains(object.Data.Type, "kiosk") {
				if object.Data.Display.Data.Kiosk != "" {
					kiosks = append(kiosks, object.Data.Display.Data.Kiosk)
				}
			} else {
				typeSplit := strings.Split(object.Data.Type, "::")

				if len(typeSplit) == 3 && len(typeSplit[0]) == 66 {
					response = append(response, object)
				}
			}
		}
	}

	kioskObjects := []string{}
	if len(kiosks) > 0 {
		for _, kiosk := range kiosks {

			// get dynamic field
			var cursor *string = nil
			for {
				fields, err := s.GetDynamicField(context.Background(), models.GetDynamicFieldRequest{
					ParentObjectID: kiosk,
					Cursor:         cursor,
				})
				if err != nil {
					break
				}

				// loop get data
				for _, fieldData := range fields.Data {
					if strings.Contains(fieldData.ObjectType, "kiosk") {
						continue
					}

					typeSplit := strings.Split(fieldData.ObjectType, "::")
					if len(typeSplit) == 3 && len(typeSplit[0]) == 66 {
						kioskObjects = append(kioskObjects, fieldData.ObjectId)
					}
				}

				cursor = &fields.NextCursor
				if !fields.HasNextPage {
					break
				}
			}
		}
	}

	if len(kioskObjects) > batchsize {
		index := 0
		step := 0
		for {
			min := index * batchsize
			max := (index + 1) * batchsize

			if step+batchsize > len(kioskObjects) {
				max = len(kioskObjects)
			}

			queryObject := kioskObjects[min:max]

			// query multi object
			resMulti, err := s.GetMultiObject(ctx, models.GetMultiObjectRequest{
				ObjectIDs: queryObject,
			})
			if err != nil {
				continue
			}

			for _, object := range resMulti.Data {
				if strings.Contains(object.Data.Type, "kiosk") {
					kiosks = append(kiosks, object.Data.ObjectID)
				} else {
					typeSplit := strings.Split(object.Data.Type, "::")

					if len(typeSplit) == 3 && len(typeSplit[0]) == 66 {
						response = append(response, object)
					}
				}
			}

			if max == len(kioskObjects) {
				break
			}

			index += 1
			step = index * batchsize
		}
	} else {
		// query multi object
		resMulti, err := s.GetMultiObject(ctx, models.GetMultiObjectRequest{
			ObjectIDs: kioskObjects,
		})
		if err != nil {
			return nil, err
		}

		for _, object := range resMulti.Data {
			typeSplit := strings.Split(object.Data.Type, "::")

			if len(typeSplit) == 3 && len(typeSplit[0]) == 66 {
				response = append(response, object)
			}
		}
	}

	return response, nil
}
