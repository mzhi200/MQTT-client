package main

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"hash"
	"net/url"
	"time"
)

type Token struct {
	Version  string  //Currently only supports "2018-10-31"
	Res      string
	Et       int     //Access expiration time expirationTime, unix time
	Method   string  //signatureMethod, md5、sha1、sha256
	Sign     string
}

func (tk *Token) InitTokenStruct() (err error){
	var mac hash.Hash
	tk.Version = config.OneNet.Token.Version
	tk.Method = config.OneNet.Token.Method
	tk.Res = fmt.Sprintf("products/%s/devices/%s", config.OneNet.ProductId, config.OneNet.EquipName)
	t, err := time.ParseInLocation("2006-01-02 15:04:05", config.OneNet.Token.Et, time.Local)
	if err != nil{
		return
	}
	tk.Et = int(t.Unix())

	//get signatureMethod
	//StringForSignature Format: et + '\n' + method + '\n' + res+ '\n' + version
	stringForSignature := fmt.Sprintf("%d\n%s\n%s\n%s", tk.Et, tk.Method, tk.Res, tk.Version)
	DcAccKey, err :=base64.StdEncoding.DecodeString(config.OneNet.EquipKey)
	if err != nil{
		return
	}
	switch tk.Method {
	case "md5":
		mac = hmac.New(md5.New, DcAccKey)

	case "sha1":
		mac = hmac.New(sha1.New, DcAccKey)

	case "sha256":
		mac = hmac.New(sha256.New, DcAccKey)

	default:
		err = errors.New(fmt.Sprintf("Unknown encryption method.: %s", tk.Method))
		return
	}
	mac.Write([]byte(stringForSignature))
	checksum := mac.Sum(nil)
	tk.Sign = base64.StdEncoding.EncodeToString(checksum[:])
	return
}

func (tk *Token) TokenToUrlFormat() (token string, err error) {
	//Source: version=2018-10-31&res=products/123123&et=1537255523&method=sha1&sign=ZjA1NzZlMmMxYzIOTg3MjBzNjYTI2MjA4Yw=
	//Encode: version=2018-10-31&res=products%2F123123&et=1537255523&method=sha1&sign=ZjA1NzZlMmMxYzIOTg3MjBzNjYTI2MjA4Yw%3D
	if tk == nil {
		err = errors.New("Token is empty!!!")
		return
	}
	p := url.Values{}
	p.Add("version", tk.Version)
	p.Add("res", tk.Res)
	p.Add("et", fmt.Sprintf("%d", tk.Et))
	p.Add("method", tk.Method)
	p.Add("sign", tk.Sign)
	token = p.Encode()
	return
}

func (tk *Token) TokenGenerateFun() (token string, err error) {
	err = tk.InitTokenStruct()
	if err != nil {
		return
	}
	token, err = tk.TokenToUrlFormat()
	if err != nil {
		return
	}
	return
}