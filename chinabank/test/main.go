package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/sm3"
	"github.com/tjfoc/gmsm/sm4"
	"github.com/tjfoc/gmsm/x509"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	//cc()
	//dd()

	//encode()
	decode()

	//genKey()

	return
	//file, _ := os.Open("./pub.key")
	//keyObj, err := sm2.GenerateKey(file)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//fmt.Println(keyObj.PublicKey)
	//fmt.Println(keyObj.X)
	//fmt.Println(keyObj.Y)
	//fmt.Println(keyObj.D)
	//
	//
	//
	//
	//msg := []byte(`{"respHeader":{"retCode":"0","traceSerial":"788646546","traceDate":"20220215","traceTime":"161800","retMsg":"成功"}}`)
	//
	//asn1, err := keyObj.EncryptAsn1(msg, file)
	//sign, err := keyObj.Sign(file, msg, nil)
	//fmt.Printf("加密结果1:%x\n",asn1)
	//fmt.Printf("sign:%x\n",sign)
	//
	//pub := &keyObj.PublicKey
	//ciphertxt, err := pub.EncryptAsn1(msg, file)
	////verify := pub.Verify(msg, ciphertxt)
	////fmt.Println(verify)
	//fmt.Printf("加密结果2:%x\n",ciphertxt)
	////fmt.Printf("加密结果2:%s\n", hex.EncodeToString(ciphertxt))
}

func genKey() {
	key, _ := sm2.GenerateKey(nil)

	x := strings.ToUpper(fmt.Sprintf("%x", key.X))
	y := strings.ToUpper(fmt.Sprintf("%x", key.Y))
	d := strings.ToUpper(fmt.Sprintf("%x", key.D))

	fmt.Printf("%s\n", x)
	fmt.Printf("%s\n", y)
	fmt.Printf("%s%s\n", x, y)
	fmt.Printf("%s\n", d)

	//pub := make([]byte, 100)
	//_, _ = key.ScalarMult(key.X, key.Y, pub)
	//fmt.Printf("%x\n", pub)
}

func decode() {
	//skey := "307802202f6673ec45e2f4424b3afcf5b06bad565441a0d62949882e92914aec7122697f022020e54b6e4af50461698e21223765fe3d27b1b2df916f7929155df2fc009fbcbb04206274f2931566194f575e46f807fbd1357a073c235c101ebcc253630cdd203c4c0410ec9766bf897cc80df12b1d0ea828e003"
	//sign := "09fd44497ddf2d58edf28fe647c481d47dfbc46c647dbd8c762e7215e7713c07"
	//body := "b48c562271c9fbdc13da81943f60eb9f93c64ce789c1914e3e0091808fd6a2b16dcf33e9ba9290ed58de50308b35eeb116175eca7e7066fa25b6c5dde4eb1116646ec365a5e3385074dcfd6751de9e86c76360c994ddd5ba0949ce30593ebf4c2def368d53e4fb4a6a728daa0364707e8a5595c01ecb5da3fcb7f5322c57ff7a"

	skey := "307a02210080428fd800743685d142e33bd98fb71db6f3a325c9d50ca822191cc7cb7a39ed022100ee5301ada0e3e43550aae7d4e702592da01516c7f98789ba38cb01b733aec25a04208be6446c122e66b45cb1128c8cb0770fe466d77e8c8dc636db08338f1786462204100caadf1ceddf2b312e075c0f935e4c40"
	sign := "09fd44497ddf2d58edf28fe647c481d47dfbc46c647dbd8c762e7215e7713c07"
	body := "b48c562271c9fbdc13da81943f60eb9f93c64ce789c1914e3e0091808fd6a2b16dcf33e9ba9290ed58de50308b35eeb116175eca7e7066fa25b6c5dde4eb1116646ec365a5e3385074dcfd6751de9e86c76360c994ddd5ba0949ce30593ebf4c2def368d53e4fb4a6a728daa0364707e8a5595c01ecb5da3fcb7f5322c57ff7a"

	//file, _ := os.Open("./priv.pem")
	//pubkeyPem, err := ioutil.ReadAll(file)
	//privateKey, err := x509.ReadPrivateKeyFromPem(pubkeyPem, nil)

	//newInt :=new(big.Int)
	//setString, _ := newInt.SetString("101959616850875943932697770166287063999683974043242337818732170373234155890783", 10)
	privateKey, err := x509.ReadPrivateKeyFromHex("CFE96E071C65A694AFF9ECCC3234AFC83183C65897616FBDD934682581113B42")
	//privateKey, err := x509.ReadPrivateKeyFromHex(setString.String())
	//privateKey, err := x509.ParseSm2PrivateKey(setString.Bytes())
	if err != nil {
		panic(err)
	}

	decodeString, _ := hex.DecodeString(skey)
	sm4Key, err := privateKey.DecryptAsn1(decodeString)
	fmt.Println("sm4Key: ", string(sm4Key))

	bodyBytes, err := hex.DecodeString(body)
	out, err := sm4.Sm4Ecb(sm4Key, bodyBytes, false)
	fmt.Printf("%s\n", out)

	signBytes, err := hex.DecodeString(sign)
	pub := &privateKey.PublicKey
	verify := pub.Verify(out, signBytes)
	fmt.Println(verify)
}

func encode() {
	//file, _ := os.Open("./pub.pem")
	//pemBytes, _ := ioutil.ReadAll(file)
	//publicKey, err := x509.ReadPublicKeyFromPem(pemBytes)
	publicKey, err := x509.ReadPublicKeyFromHex("70791433D7BC313CBC8B5307B688C4A2762EB8CE87FA298001647FB9AB88C6F14DDBE539847B49AF7552A10D3651BA566E708A3262A2F799465DDA65681FF9FB")
	if err != nil {
		panic(err)
	}

	msg := []byte(`{"respHeader":{"retCode":"0","traceSerial":"788646546","traceDate":"20220215","traceTime":"161800","retMsg":"成功"}}`)
	sm4Key := []byte("1234567890abcdef")

	body := GenBody(msg, sm4Key)
	sign := GenSign(msg)
	skey := GenSkey(publicKey, sm4Key)

	fmt.Println(skey)
	fmt.Println(sign)
	fmt.Println(body)
}

func GenSkey(pub *sm2.PublicKey, sm4Key []byte) string {
	asn1, err := pub.EncryptAsn1(sm4Key, rand.Reader)
	if err != nil {
		return "err"
	}

	return fmt.Sprintf("%x", asn1)
	//fmt.Printf("加密后： %x\n", asn1)

	//dest, err := keyObj.DecryptAsn1(asn1)
	//fmt.Printf("解密后： %s\n", dest)
}
func GenSign(msg []byte) string {
	h := sm3.New()
	h.Write(msg)
	sum := h.Sum(nil)
	//fmt.Printf("digest value is: %x\n",sum)
	return fmt.Sprintf("%x", sum)
}
func GenBody(msg []byte, sm4Key []byte) string {
	//fmt.Printf("key = %v\n", key)
	//fmt.Printf("data = %x\n", msg)
	//iv := []byte("0000000000000000")
	//err = SetIV(iv)//设置SM4算法实现的IV值,不设置则使用默认值
	ecbMsg, err :=sm4.Sm4Ecb(sm4Key, msg, true)   //sm4Ecb模式pksc7填充加密
	if err != nil {
		fmt.Println(err)
		return ""
	}
	//fmt.Printf("ecbMsg = %x\n", ecbMsg)
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
