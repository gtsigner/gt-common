package common

import (
    "bytes"
    "crypto/aes"
    "crypto/cipher"
    "crypto/hmac"
    "crypto/sha256"
)

/*CBC加密 按照golang标准库的例子代码
不过里面没有填充的部分,所以补上，根据key来决定填充blocksize
*/

//使用PKCS7进行填充，IOS也是7
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
    padding := blockSize - len(ciphertext)%blockSize
    padtext := bytes.Repeat([]byte{byte(padding)}, padding)
    return append(ciphertext, padtext...)
}

func PKCS7UnPadding(origData []byte) []byte {
    length := len(origData)
    unpadding := int(origData[length-1])
    return origData[:(length - unpadding)]
}

//aes加密，填充模式由key决定，16位，24,32分别对应AES-128, AES-192, or AES-256.源码好像是写死16了
func AesCBCPKCS5PaddingEncrypt(rawData, key []byte, iv []byte) []byte {
    block, err := aes.NewCipher(key)
    if err != nil {
        panic(err)
    }
    rawData = PKCS7Padding(rawData, block.BlockSize())
    //初始向量IV必须是唯一，但不需要保密
    encrypted := make([]byte, len(rawData))

    //block大小和初始向量大小一定要一致
    mode := cipher.NewCBCEncrypter(block, iv)
    mode.CryptBlocks(encrypted, rawData)
    return encrypted
}
func AesCBCPKCS7PaddingEncrypt(rawData, key []byte, iv []byte) []byte {
    block, err := aes.NewCipher(key)
    if err != nil {
        panic(err)
    }
    rawData = PKCS7Padding(rawData, block.BlockSize())
    //初始向量IV必须是唯一，但不需要保密
    encrypted := make([]byte, len(rawData))

    //block大小和初始向量大小一定要一致
    mode := cipher.NewCBCEncrypter(block, iv)
    mode.CryptBlocks(encrypted, rawData)
    return encrypted
}

func AesCBCDncrypt(encryptData, key []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        panic(err)
    }

    blockSize := block.BlockSize()

    if len(encryptData) < blockSize {
        panic("ciphertext too short")
    }
    iv := encryptData[:blockSize]
    encryptData = encryptData[blockSize:]

    // CBC mode always works in whole blocks.
    if len(encryptData)%blockSize != 0 {
        panic("ciphertext is not a multiple of the block size")
    }

    mode := cipher.NewCBCDecrypter(block, iv)

    // CryptBlocks can work in-place if the two arguments are the same.
    mode.CryptBlocks(encryptData, encryptData)
    //解填充
    encryptData = PKCS7UnPadding(encryptData)
    return encryptData, nil
}

func EcbDecrypt(data, key []byte) []byte {
    block, _ := aes.NewCipher(key)
    decrypted := make([]byte, len(data))
    size := block.BlockSize()

    for bs, be := 0, size; bs < len(data); bs, be = bs+size, be+size {
        block.Decrypt(decrypted[bs:be], data[bs:be])
    }

    return PKCS7UnPadding(decrypted)
}

func EcbEncrypt(data, key []byte) []byte {
    block, _ := aes.NewCipher(key)
    data = PKCS7Padding(data, block.BlockSize())
    decrypted := make([]byte, len(data))
    size := block.BlockSize()

    for bs, be := 0, size; bs < len(data); bs, be = bs+size, be+size {
        block.Encrypt(decrypted[bs:be], data[bs:be])
    }

    return decrypted
}

func Hmac256Encrypt(origData, key []byte) []byte {
    mac := hmac.New(sha256.New, key)
    mac.Write(origData)
    expectedMAC := mac.Sum(nil)
    return expectedMAC
}
