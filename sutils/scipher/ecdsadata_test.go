// Scry Info.  All rights reserved.
// license that can be found in the license file.
package scipher

import (
	"crypto/rand"
	"crypto/rsa"
	"testing"
)

func TestEncodeData(t *testing.T) {
	data := []string{"23fdsfdsf", "奔跑发顶替鼠标可想而知枯标有夺顶替123219038721hfs9*&H943jf顶替"}
	for _, it := range data {

		pri, _ := rsa.GenerateKey(rand.Reader, 1024)

		endata, headLen, _ := encodeRsaPublic([]byte(it), &pri.PublicKey)

		tdata, _ := decodeRsaPrivate(endata, pri, headLen)

		if it != string(tdata) {
			t.Errorf("err: %s", it)
		}
	}

	for _, it := range data {

		pri, _ := rsa.GenerateKey(rand.Reader, 1024)

		endata, headLen, _ := EncodeRsaPrivate([]byte(it), pri)

		tdata, _ := DecodeRsaPublic(endata, &pri.PublicKey, headLen)

		if it != string(tdata) {
			t.Errorf("err: %s", it)
		}
	}

	return
}

func TestR(t *testing.T) {

}
