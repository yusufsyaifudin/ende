package ende

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/asn1"
	"encoding/pem"
	"fmt"
)

type RSAKeyPair struct {
	Private *rsa.PrivateKey
	Public  rsa.PublicKey
}

func GenerateRSAKeypair() (*RSAKeyPair, error) {
	reader := rand.Reader
	bitSize := 4096

	key, err := rsa.GenerateKey(reader, bitSize)
	if err != nil {
		err = fmt.Errorf("generate RSA pair error: %w", err)
		return nil, err
	}

	return &RSAKeyPair{
		Private: key,
		Public:  key.PublicKey,
	}, nil
}

func RSAPrivateEncode(encoder Encoder, key *rsa.PrivateKey) (string, error) {
	var privateKey = &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}

	buffer := &bytes.Buffer{}
	err := pem.Encode(buffer, privateKey)
	if err != nil {
		return "", err
	}

	return encoder.EncodeToString(context.Background(), buffer.Bytes())
}

func RSAPublicKeyEncode(encoder Encoder, pubKey rsa.PublicKey) (string, error) {
	asn1Bytes, err := asn1.Marshal(pubKey)
	if err != nil {
		return "", err
	}

	var pemKey = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: asn1Bytes,
	}

	buffer := &bytes.Buffer{}
	err = pem.Encode(buffer, pemKey)
	if err != nil {
		return "", err
	}

	return encoder.EncodeToString(context.Background(), buffer.Bytes())
}
