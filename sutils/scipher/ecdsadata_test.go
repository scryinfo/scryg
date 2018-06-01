package scipher

import (
	"crypto/rand"
	"crypto/rsa"
	"testing"
)

func TestEncodeData(t *testing.T){
	data := []string{"23fdsfdsf","奔跑发顶替鼠标可想而知枯标有夺顶替123219038721hfs9*&H943jf顶替"}
	for _, it := range data {

		pri, _ := rsa.GenerateKey(rand.Reader, 1024)

		endata, _ := EncodeDataRsa(it, &pri.PublicKey)

		tdata, _ := DecodeDataRsa(endata, pri)

		if it != tdata {
			t.Failed()
		}
	}
	return
}

