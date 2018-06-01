package scipher

import (
	"crypto/ecdsa"

	"crypto/rand"

	"crypto/cipher"
	"crypto/elliptic"
	"crypto/rsa"
	"crypto/sha256"
	"errors"

	"github.com/btcsuite/btcutil/base58"
	"crypto/aes"
	"bytes"
)

const headLen = 128 //byte

//
func EncodeDataRsa(data string, pubKey *rsa.PublicKey) (endata string, err error) {
	endata = ""
	err = nil
	var key []byte
	var enkey []byte
	{
		priv, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		hash := sha256.New()
		hash.Write(priv.PublicKey.X.Bytes())
		key = hash.Sum(nil)
		enkey, err = rsa.EncryptPKCS1v15(rand.Reader, pubKey, key)
		if err != nil {
			return
		}
	}

	var bendata []byte
	{
		var block cipher.Block

		block, err = aes.NewCipher(key)
		if err != nil {
			return
		}
		bs := block.BlockSize()
		bdata := []byte(data)
		bdata = PKCS5Padding(bdata, bs)
		if len(bdata)%bs != 0 {
			err = errors.New("Need a multiple of the blocksize")
			return
		}
		out := make([]byte, len(bdata))
		dst := out
		for len(bdata) > 0 {
			block.Encrypt(dst, bdata[:bs])
			bdata = bdata[bs:]
			dst = dst[bs:]
		}
		bendata = out
	}
	bendata = append(enkey, bendata...)
	endata = base58.Encode(bendata)
	return
}


func DecodeDataRsa(endata string, priKey *rsa.PrivateKey) (data string, err error) {

	err = nil
	data = ""
	bendata := base58.Decode(endata)
	enkey := bendata[:headLen]
	enbody := bendata[headLen:]
	var key []byte
	{
		key, err = rsa.DecryptPKCS1v15(rand.Reader, priKey, enkey)
		if err != nil {
			return
		}
	}

	var bdata []byte
	{
		var block cipher.Block
		block ,err = aes.NewCipher(key)
		if err != nil {
			return
		}

		out := make([]byte, len(enbody))
		dst := out
		bs := block.BlockSize()
		if len(enbody) % bs != 0 {
			err = errors.New("crypto/cipher: input not full blocks")
			return
		}
		for len(enbody) > 0 {
			block.Decrypt(dst, enbody[:bs])
			enbody = enbody[bs:]
			dst = dst[bs:]
		}
		bdata = PKCS5UnPadding(out)
	}

	data = string(bdata)
	return
}

//
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
//
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length - 1])
	return origData[:(length - unpadding)]
}
