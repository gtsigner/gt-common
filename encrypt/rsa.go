package encrypt

import (
    "crypto/rand"
    "crypto/rsa"
    "encoding/base64"
    "encoding/hex"
    "errors"
    "math/big"
)

type Rsa struct {
    publicKey  *rsa.PublicKey
    privateKey *rsa.PrivateKey
}

func (r *Rsa) SetPublicKey(pub string, e int) error {
    k := &rsa.PublicKey{}
    n, ok := new(big.Int).SetString(pub, 16)
    if !ok {
        return errors.New("set public hex fail")
    }
    k.N = n
    k.E = e
    r.publicKey = k
    return nil
}

func (r *Rsa) Encrypt(str string, encoding string) (string, error) {
    bytes := []byte(str)
    res, err := rsa.EncryptPKCS1v15(rand.Reader, r.publicKey, bytes)
    if err != nil {
        return "", err
    }
    if encoding == "base64" {
        return base64.StdEncoding.EncodeToString(res), nil
    }
    return hex.EncodeToString(res), nil
}

func (r *Rsa) Encrypt2(bytes []byte, encoding string) (string, error) {
    res, err := rsa.EncryptPKCS1v15(rand.Reader, r.publicKey, bytes)
    if err != nil {
        return "", err
    }
    if encoding == "base64" {
        return base64.StdEncoding.EncodeToString(res), nil
    }
    return hex.EncodeToString(res), nil
}
func NewRsa() *Rsa {
    rs := &Rsa{}
    return rs
}
