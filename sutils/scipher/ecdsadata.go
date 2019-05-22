// Scry Info.  All rights reserved.
// license that can be found in the license file.

package scipher

import (
	"crypto/ecdsa"
	"io"
	"io/ioutil"
	"math/big"

	"crypto/rand"

	"crypto/cipher"
	"crypto/elliptic"
	"crypto/rsa"
	"crypto/sha256"
	"errors"

	"bytes"
	"crypto/aes"
)

//
func encodeRsaPublic(data []byte, pubKey *rsa.PublicKey) (endata []byte, headLen int, err error) {
	endata = nil
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
		headLen = len(enkey)
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
	endata = append(enkey, bendata...)
	return
}

func decodeRsaPrivate(endata []byte, priKey *rsa.PrivateKey, headLen int) (data []byte, err error) {

	err = nil
	data = nil
	bendata := endata
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
		block, err = aes.NewCipher(key)
		if err != nil {
			return
		}

		out := make([]byte, len(enbody))
		dst := out
		bs := block.BlockSize()
		if len(enbody)%bs != 0 {
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

	data = bdata
	return
}

// encrypt data with Rsa'Private key
func EncodeRsaPrivate(data []byte, priKey *rsa.PrivateKey) (endata []byte, headLen int, err error) {
	endata = nil
	err = nil
	var key []byte
	var enkey []byte
	{
		priv, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		hash := sha256.New()
		hash.Write(priv.PublicKey.X.Bytes())
		key = hash.Sum(nil)

		{
			output := bytes.NewBuffer(nil)
			err = priKeyIO(priKey, bytes.NewReader(key), output, true)
			if err != nil {
				return
			}
			enkey, err = ioutil.ReadAll(output)
			if err != nil {
				return
			}
			headLen = len(enkey)
		}

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
	endata = append(enkey, bendata...)
	return
}

// decrypt data with the public key of Rsa
func DecodeRsaPublic(endata []byte, pubKey *rsa.PublicKey, headLen int) (data []byte, err error) {

	err = nil
	data = nil
	bendata := endata
	enkey := bendata[:headLen]
	enbody := bendata[headLen:]
	var key []byte
	{
		output := bytes.NewBuffer(nil)
		err = pubKeyIO(pubKey, bytes.NewReader(enkey), output, false)
		if err != nil {
			return
		}
		key, err = ioutil.ReadAll(output)
		if err != nil {
			return
		}
	}

	var bdata []byte
	{
		var block cipher.Block
		block, err = aes.NewCipher(key)
		if err != nil {
			return
		}

		out := make([]byte, len(enbody))
		dst := out
		bs := block.BlockSize()
		if len(enbody)%bs != 0 {
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

	data = bdata
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
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

var (
	errDataToLarge     = errors.New("message too long for RSA public key size")
	errDataLen         = errors.New("data length error")
	errDataBroken      = errors.New("data broken, first byte is not zero")
	errKeyPairDismatch = errors.New("data is not encrypted by the private key")
	errDecryption      = errors.New("decryption error")
	errPublicKey       = errors.New("get public key error")
	errPrivateKey      = errors.New("get private key error")
)

//see  https://github.com/farmerx/gorsa
func pubKeyByte(pub *rsa.PublicKey, in []byte, isEncrytp bool) ([]byte, error) {
	k := (pub.N.BitLen() + 7) / 8
	if isEncrytp {
		k = k - 11
	}
	if len(in) <= k {
		if isEncrytp {
			return rsa.EncryptPKCS1v15(rand.Reader, pub, in)
		} else {
			return pubKeyDecrypt(pub, in)
		}
	} else {
		iv := make([]byte, k)
		out := bytes.NewBuffer(iv)
		if err := pubKeyIO(pub, bytes.NewReader(in), out, isEncrytp); err != nil {
			return nil, err
		}
		return ioutil.ReadAll(out)
	}
}

//
func priKeyByte(pri *rsa.PrivateKey, in []byte, isEncrytp bool) ([]byte, error) {
	k := (pri.N.BitLen() + 7) / 8
	if isEncrytp {
		k = k - 11
	}
	if len(in) <= k {
		if isEncrytp {
			return priKeyEncrypt(rand.Reader, pri, in)
		} else {
			return rsa.DecryptPKCS1v15(rand.Reader, pri, in)
		}
	} else {
		iv := make([]byte, k)
		out := bytes.NewBuffer(iv)
		if err := priKeyIO(pri, bytes.NewReader(in), out, isEncrytp); err != nil {
			return nil, err
		}
		return ioutil.ReadAll(out)
	}
}

//
func pubKeyIO(pub *rsa.PublicKey, in io.Reader, out io.Writer, isEncrytp bool) (err error) {
	k := (pub.N.BitLen() + 7) / 8
	if isEncrytp {
		k = k - 11
	}
	buf := make([]byte, k)
	var b []byte
	size := 0
	for {
		size, err = in.Read(buf)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		if size < k {
			b = buf[:size]
		} else {
			b = buf
		}
		if isEncrytp {
			b, err = rsa.EncryptPKCS1v15(rand.Reader, pub, b)
		} else {
			b, err = pubKeyDecrypt(pub, b)
		}
		if err != nil {
			return err
		}
		if _, err = out.Write(b); err != nil {
			return err
		}
	}
	return nil
}

//
func priKeyIO(pri *rsa.PrivateKey, r io.Reader, w io.Writer, isEncrytp bool) (err error) {
	k := (pri.N.BitLen() + 7) / 8
	if isEncrytp {
		k = k - 11
	}
	buf := make([]byte, k)
	var b []byte
	size := 0
	for {
		size, err = r.Read(buf)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		if size < k {
			b = buf[:size]
		} else {
			b = buf
		}
		if isEncrytp {
			b, err = priKeyEncrypt(rand.Reader, pri, b)
		} else {
			b, err = rsa.DecryptPKCS1v15(rand.Reader, pri, b)
		}
		if err != nil {
			return err
		}
		if _, err = w.Write(b); err != nil {
			return err
		}
	}
	return nil
}

//
func pubKeyDecrypt(pub *rsa.PublicKey, data []byte) ([]byte, error) {
	k := (pub.N.BitLen() + 7) / 8
	if k != len(data) {
		return nil, errDataLen
	}
	m := new(big.Int).SetBytes(data)
	if m.Cmp(pub.N) > 0 {
		return nil, errDataToLarge
	}
	m.Exp(m, big.NewInt(int64(pub.E)), pub.N)
	d := leftPad(m.Bytes(), k)
	if d[0] != 0 {
		return nil, errDataBroken
	}
	if d[1] != 0 && d[1] != 1 {
		return nil, errKeyPairDismatch
	}
	var i = 2
	for ; i < len(d); i++ {
		if d[i] == 0 {
			break
		}
	}
	i++
	if i == len(d) {
		return nil, nil
	}
	return d[i:], nil
}

//
func priKeyEncrypt(rand io.Reader, priv *rsa.PrivateKey, hashed []byte) ([]byte, error) {
	tLen := len(hashed)
	k := (priv.N.BitLen() + 7) / 8
	if k < tLen+11 {
		return nil, errDataLen
	}
	em := make([]byte, k)
	em[1] = 1
	for i := 2; i < k-tLen-1; i++ {
		em[i] = 0xff
	}
	copy(em[k-tLen:k], hashed)
	m := new(big.Int).SetBytes(em)
	c, err := decrypt(rand, priv, m)
	if err != nil {
		return nil, err
	}
	copyWithLeftPad(em, c.Bytes())
	return em, nil
}

var bigZero = big.NewInt(0)
var bigOne = big.NewInt(1)

func encrypt(c *big.Int, pub *rsa.PublicKey, m *big.Int) *big.Int {
	e := big.NewInt(int64(pub.E))
	c.Exp(m, e, pub.N)
	return c
}

func decrypt(random io.Reader, priv *rsa.PrivateKey, c *big.Int) (m *big.Int, err error) {
	if c.Cmp(priv.N) > 0 {
		err = errDecryption
		return
	}
	var ir *big.Int
	if random != nil {
		var r *big.Int

		for {
			r, err = rand.Int(random, priv.N)
			if err != nil {
				return
			}
			if r.Cmp(bigZero) == 0 {
				r = bigOne
			}
			var ok bool
			ir, ok = modInverse(r, priv.N)
			if ok {
				break
			}
		}
		bigE := big.NewInt(int64(priv.E))
		rpowe := new(big.Int).Exp(r, bigE, priv.N)
		cCopy := new(big.Int).Set(c)
		cCopy.Mul(cCopy, rpowe)
		cCopy.Mod(cCopy, priv.N)
		c = cCopy
	}
	if priv.Precomputed.Dp == nil {
		m = new(big.Int).Exp(c, priv.D, priv.N)
	} else {
		m = new(big.Int).Exp(c, priv.Precomputed.Dp, priv.Primes[0])
		m2 := new(big.Int).Exp(c, priv.Precomputed.Dq, priv.Primes[1])
		m.Sub(m, m2)
		if m.Sign() < 0 {
			m.Add(m, priv.Primes[0])
		}
		m.Mul(m, priv.Precomputed.Qinv)
		m.Mod(m, priv.Primes[0])
		m.Mul(m, priv.Primes[1])
		m.Add(m, m2)

		for i, values := range priv.Precomputed.CRTValues {
			prime := priv.Primes[2+i]
			m2.Exp(c, values.Exp, prime)
			m2.Sub(m2, m)
			m2.Mul(m2, values.Coeff)
			m2.Mod(m2, prime)
			if m2.Sign() < 0 {
				m2.Add(m2, prime)
			}
			m2.Mul(m2, values.R)
			m.Add(m, m2)
		}
	}
	if ir != nil {
		m.Mul(m, ir)
		m.Mod(m, priv.N)
	}

	return
}

func copyWithLeftPad(dest, src []byte) {
	numPaddingBytes := len(dest) - len(src)
	for i := 0; i < numPaddingBytes; i++ {
		dest[i] = 0
	}
	copy(dest[numPaddingBytes:], src)
}

func nonZeroRandomBytes(s []byte, rand io.Reader) (err error) {
	_, err = io.ReadFull(rand, s)
	if err != nil {
		return
	}
	for i := 0; i < len(s); i++ {
		for s[i] == 0 {
			_, err = io.ReadFull(rand, s[i:i+1])
			if err != nil {
				return
			}
			s[i] ^= 0x42
		}
	}
	return
}

func leftPad(input []byte, size int) (out []byte) {
	n := len(input)
	if n > size {
		n = size
	}
	out = make([]byte, size)
	copy(out[len(out)-n:], input)
	return
}

func modInverse(a, n *big.Int) (ia *big.Int, ok bool) {
	g := new(big.Int)
	x := new(big.Int)
	y := new(big.Int)
	g.GCD(x, y, a, n)
	if g.Cmp(bigOne) != 0 {
		return
	}
	if x.Cmp(bigOne) < 0 {
		x.Add(x, n)
	}
	return x, true
}
