package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/hashicorp/go-uuid"
	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/sm3"
	"github.com/tjfoc/gmsm/sm4"
	"github.com/tjfoc/gmsm/x509"
	"io/ioutil"
	"math/big"
	"os"
	"strings"
)

type Request struct {
	SKey string
	Body string
	Sign string
}

func main() {
	//genKey()

	//encode()
	//decode()

	//signTest()

	test()
}

func test() {
	public := "04627A997F9485281068B0262B346E7AEE72580079CE10E248BE914F8D227F0B2955523D62EA126CA6A68312C43E4A3736D9ADF6118B7759559709ADCF439A5636"
	private := "95463839606194440177843520030979835176729602072477836123530316400792948456063"
	fmt.Printf("公钥: %s\n", public)
	fmt.Printf("私钥: %s\n", private)

	req := encode(public)

	fmt.Println("加密结果:")
	fmt.Printf("skey: %s\n", req.SKey)
	fmt.Printf("sign: %s\n", req.Sign)
	fmt.Printf("body: %s\n", req.Body)

	fmt.Println("解密结果:")
	decode(private, req)
}

func signTest() {
	msg := []byte(`{"respHeader":{"retCode":"0","traceSerial":"788646546"}}`)
	sign := GenSign(msg)

	//privateKey, err := x509.ReadPrivateKeyFromHex("CFE96E071C65A694AFF9ECCC3234AFC83183C65897616FBDD934682581113B42")
	//if err != nil {
	//	panic(err)
	//	return
	//}

	target := GenSign(msg)
	if target == sign {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
	//signBytes, err := hex.DecodeString(sign)
	//verify := privateKey.Verify(msg, signBytes)
	//fmt.Println(verify)
}

func genKey() {
	key, _ := sm2.GenerateKey(nil)

	// 获取十六进制的公钥和私钥
	privateKeyHex := x509.WritePrivateKeyToHex(key)
	publicKeyHex := x509.WritePublicKeyToHex(&key.PublicKey)

	b := new(big.Int)
	setString, _ := b.SetString(privateKeyHex, 16)
	fmt.Println(setString.String())
	fmt.Println(strings.ToUpper(publicKeyHex))
	return

	x := strings.ToUpper(fmt.Sprintf("%x", key.X))
	y := strings.ToUpper(fmt.Sprintf("%x", key.Y))
	//d := strings.ToUpper(fmt.Sprintf("%x", key.D))

	fmt.Printf("%s\n", x)
	fmt.Printf("%s\n", y)
	fmt.Printf("公钥： %s%s\n", x, y)
	//fmt.Printf("%s\n", d)
	fmt.Println("私钥： ", key.D.String())
}

func javaByteToGoByte(src string) string {
	desc := make([]byte, 0)
	runes := []rune(src)
	for _, r := range runes {
		if r < -128 || r > 127 {
			return ""
		}
		if r < 0 {
			desc = append(desc, byte(256 + r))
		} else {
			desc = append(desc, byte(r))
		}
	}
	return string(desc)
}

func decode(privateKeyStr string, req *Request) {
	skey := req.SKey
	sign := req.Sign
	body := req.Body

	//file, _ := os.Open("./priv.pem")
	//pubkeyPem, err := ioutil.ReadAll(file)
	//privateKey, err := x509.ReadPrivateKeyFromPem(pubkeyPem, nil)

	// 将十进制的密钥转成十六进制的密钥
	b := new(big.Int)
	bString, _ := b.SetString(privateKeyStr, 10)
	hexD := fmt.Sprintf("%x", bString)

	privateKey, err := x509.ReadPrivateKeyFromHex(hexD)
	if err != nil {
		panic(err)
	}

	decodeString, _ := hex.DecodeString(skey)
	sm4Key, err := privateKey.DecryptAsn1(decodeString)
	fmt.Println("sm4Key:", string(sm4Key))

	bodyBytes, err := hex.DecodeString(body)
	sm4KeyBytes, err := hex.DecodeString(string(sm4Key))
	out, err := sm4.Sm4Ecb(sm4KeyBytes, bodyBytes, false)
	fmt.Printf("解密出来的报文: %s\n", out)

	target := GenSign(out)
	if target == sign {
		fmt.Println("签名检验结果: ", true)
	} else {
		fmt.Println("签名检验结果: ", false)
	}

	//signBytes, err := hex.DecodeString(sign)
	//verify := privateKey.Verify(out, signBytes)
	//fmt.Println(verify)
}

func encode(publicKeyStr string) *Request {
	//file, _ := os.Open("./pub.pem")
	//pemBytes, _ := ioutil.ReadAll(file)
	//publicKey, err := x509.ReadPublicKeyFromPem(pemBytes)
	publicKey, err := x509.ReadPublicKeyFromHex(publicKeyStr)
	if err != nil {
		panic(err)
	}

	msg := []byte(`{"respHeader":{"retCode":"0","traceSerial":"788646546","traceDate":"20220215","traceTime":"161800","retMsg":"成功"}}`)
	fmt.Printf("原始报文: %s\n", msg)
	// 使用随机算法生成32位的随机字符串
	uid, _ := uuid.GenerateUUID()
	uid = strings.ReplaceAll(uid, "-", "")
	sm4Key := []byte(uid)
	fmt.Printf("sm4Key: %s\n", sm4Key)

	body := GenBody(msg, sm4Key)
	sign := GenSign(msg)
	skey := GenSkey(publicKey, sm4Key)

	return &Request{
		SKey: skey,
		Body: body,
		Sign: sign,
	}
}

func GenSkey(pub *sm2.PublicKey, sm4Key []byte) string {
	asn1, err := pub.EncryptAsn1(sm4Key, rand.Reader)
	if err != nil {
		return "err"
	}
	return fmt.Sprintf("%x", asn1)
}
func GenSign(msg []byte) string {
	h := sm3.New()
	h.Write(msg)
	sum := h.Sum(nil)
	return fmt.Sprintf("%x", sum)
}
func GenBody(msg []byte, sm4Key []byte) string {
	decodeSm4, _ := hex.DecodeString(string(sm4Key))

	//iv := []byte("00000000000000000000000000000000")
	//err := sm4.SetIV(iv)//设置SM4算法实现的IV值,不设置则使用默认值
	ecbMsg, err := sm4.Sm4Ecb(decodeSm4, msg, true)   //sm4Ecb模式pksc7填充加密
	//ecbMsg, err := sm4.Sm4Cbc(decodeSm4, msg, true)
	//ecbMsg, err := sm4.Sm4CFB(decodeSm4, msg, true)
	//ecbMsg, err := sm4.Sm4OFB(decodeSm4, msg, true)

	if err != nil {
		fmt.Println(err)
		return ""
	}
	return fmt.Sprintf("%x", ecbMsg)
}

func dd() {
	file1, err := os.Open("./pub.pem")
	pubkeyPem, err := ioutil.ReadAll(file1)
	pubKey, err := x509.ReadPublicKeyFromPem(pubkeyPem) // 读取公钥
	if err != nil {
		fmt.Println(err)
		return
	}

	msg := []byte(`{"respHeader":{"retCode":"0","traceSerial":"788646546","traceDate":"20220215","traceTime":"161800","retMsg":"成功"}}`)

	asn1, err := pubKey.EncryptAsn1(msg, nil)
	fmt.Printf("加密结果:%x\n",asn1)
}

// 生成证书 保存到文件
func cc() {
	priv, err := sm2.GenerateKey(nil) // 生成密钥对
	if err != nil {
		fmt.Println(err)
		return
	}
	privPem, err := x509.WritePrivateKeyToPem(priv, nil) // 生成密钥文件
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(privPem)
	file1, err := os.OpenFile("./priv.pem", os.O_WRONLY|os.O_APPEND, 0777)
	file1.Write(privPem)
	defer file1.Close()

	pubKey, _ := priv.Public().(*sm2.PublicKey)
	pubkeyPem, err := x509.WritePublicKeyToPem(pubKey)       // 生成公钥文件
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(pubkeyPem)
	file2, err := os.OpenFile("./pub.pem", os.O_WRONLY|os.O_APPEND, 0777)
	n, err := file2.Write(pubkeyPem)
	if err != nil {
		fmt.Println("n:", n)
		fmt.Println("error:", err)
	}
	defer file2.Close()
}
