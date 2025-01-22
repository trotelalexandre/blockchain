package common

import (
	"crypto/ecdh"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"math/big"
	"strings"
)

func GenerateKeyPair() (*ecdsa.PrivateKey, *ecdsa.PublicKey) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, nil
	}
	return privateKey, &privateKey.PublicKey
}

func GetPublicKeyFromPrivateKey(privateKey *ecdsa.PrivateKey) *ecdsa.PublicKey {
	return &privateKey.PublicKey
}

func GenerateAddress(publicKey *ecdsa.PublicKey) string {
	ecdhPubKey, err := ecdh.P256().NewPublicKey(publicKey.X.Bytes())
    if err != nil {
        return ""
    }
    hash := sha256.Sum256(ecdhPubKey.Bytes())
    return "proto" + hex.EncodeToString(hash[:])
}

func HashData(data []byte) string {
	hasher := sha256.New()
	hasher.Write(data)
	return hex.EncodeToString(hasher.Sum(nil))
}

func VerifySignature(signerPubKey *ecdsa.PublicKey, signData []byte, signature []byte) bool {
	if len(signature) != 64 {
		return false
	}

	r := new(big.Int).SetBytes(signature[:32])
	s := new(big.Int).SetBytes(signature[32:])

	hash := sha256.Sum256(signData)

	return ecdsa.Verify(signerPubKey, hash[:], r, s)
}

func PrivateKeyToSeedPhrase(privateKey *ecdsa.PrivateKey) string {
	privKeyBytes := privateKey.D.Bytes()
	words := []string{}
	for i := 0; i < len(privKeyBytes); i += 2 {
		word := hex.EncodeToString(privKeyBytes[i:i+2])
		words = append(words, word)
	}
	return strings.Join(words, " ")
}