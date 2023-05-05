package account

import (
	"bytes"
	"crypto/ed25519"
	"encoding/base64"
	"encoding/hex"

	"github.com/fardream/go-bcs/bcs"
	"github.com/shoshinsquare/sui-go-sdk/crypto"
	"github.com/shoshinsquare/sui-go-sdk/types"
	"golang.org/x/crypto/blake2b"
)

const (
	ADDRESS_LENGTH = 64
)

type Account struct {
	privateKey []byte
	PublicKey  []byte
	Address    string
}

func NewAccount(seed []byte) *Account {

	keyPair := crypto.NewEd25519KeyPair(ed25519.NewKeyFromSeed(seed[:]))
	tmp := []byte{0x0}
	tmp = append(tmp, keyPair.PublicKey()...)
	addrBytes := blake2b.Sum256(tmp)
	address := "0x" + hex.EncodeToString(addrBytes[:])[:ADDRESS_LENGTH]

	return &Account{
		privateKey: keyPair.PrivateKey(),
		PublicKey:  keyPair.PublicKey(),
		Address:    address,
	}
}

func (a *Account) Sign(data []byte) []byte {
	return ed25519.Sign(a.privateKey, data)
}

func (a *Account) SignAndGetEncode(txString string) (string, error) {
	txByte, err := base64.StdEncoding.DecodeString(txString)
	if err != nil {	
		return "", err
	}

	value := types.NewIntentMessage(types.DefaultIntent(), txByte)
	message, err := bcs.Marshal(value)
	if err != nil {
		return "", err
	}

	hash := blake2b.Sum256(message)
	signature := a.Sign(hash[:])

	signatureBuffer := bytes.NewBuffer([]byte{})
	signatureBuffer.WriteByte(0)
	signatureBuffer.Write(signature)
	signatureBuffer.Write(a.PublicKey)
	var signatureBytes [ed25519.PublicKeySize + ed25519.SignatureSize + 1]byte
	copy(signatureBytes[:], signatureBuffer.Bytes())

	res := base64.StdEncoding.EncodeToString(signatureBuffer.Bytes())

	return res, nil
}
