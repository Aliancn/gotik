package token

import (
	"bytes"
	"crypto/des"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"errors"
	"math/rand"
	"time"
)

var key []byte

// 生成随机的密钥
func init() {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	r.Uint64()
	k := make([]byte, 0, 8)
	k = binary.AppendUvarint(k, r.Uint64())
	key = k
}

// 下面的四个函数是我从百度上抄的对称加密解密的写法, 我自己也不懂

func padPwd(srcByte []byte, blockSize int) []byte {
	padNum := blockSize - len(srcByte)%blockSize
	ret := bytes.Repeat([]byte{byte(padNum)}, padNum)
	srcByte = append(srcByte, ret...)
	return srcByte
}

// 去掉填充的部分
func unPadPwd(dst []byte) ([]byte, error) {
	if len(dst) == 0 {
		panic("dst length is zero")
	}

	unpadNum := int(dst[len(dst)-1])
	return dst[:(len(dst) - unpadNum)], nil
}

func desEncoding(src string, key []byte) (string, error) {
	srcByte := []byte(src)
	block, err := des.NewCipher(key)
	if err != nil {
		return src, err
	}
	newSrcByte := padPwd(srcByte, block.BlockSize())
	dst := make([]byte, len(newSrcByte))
	block.Encrypt(dst, newSrcByte)
	pwd := base64.StdEncoding.EncodeToString(dst)
	return pwd, nil
}

func desDecoding(pwd string, key []byte) (string, error) {
	pwdByte, err := base64.StdEncoding.DecodeString(pwd)
	if err != nil {
		return pwd, err
	}
	block, errBlock := des.NewCipher(key)
	if errBlock != nil {
		return pwd, errBlock
	}
	dst := make([]byte, len(pwdByte))
	block.Decrypt(dst, pwdByte)
	dst, _ = unPadPwd(dst)
	return string(dst), nil
}

type TokenInfo struct {
	Username  string `json:"username"`
	CreatedAt int64  `json:"time_stamp"`
}

func NewToken(uname string) string {
	info := TokenInfo{}
	info.CreatedAt = time.Now().Unix()
	info.Username = uname

	bs, err := json.Marshal(&info)
	if err != nil {
		panic("unexpected error")
	}

	token, err := desEncoding(string(bs), key)
	if err != nil {
		panic(err)
	}

	return token
}

func GetTokenInfoFromToken(token string) (*TokenInfo, error) {
	infoBytes, err := desDecoding(token, key)
	if err != nil {
		return nil, errors.New("invalid token")
	}

	ti := TokenInfo{}
	if err := json.Unmarshal([]byte(infoBytes), &ti); err != nil {
		return nil, errors.New("unmarshal error")
	}

	return &ti, nil
}
