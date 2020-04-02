package wallet

import (
	"github.com/ethereum/go-ethereum/common"
)

//BIP路径
const defaultDerivationPath = "m/44'/60'/0'/0/1"

//钱包结构体
type HDWallet struct {
	Address    common.Address
	HdKeyStore *hdkeystore.HDkeyStore
}

//钱包构造函数
func NewWallet(keypath string) (*HDWallet, error) {
	//1. 创建助记词
	mn, err := create_mnemonic()
	if err != nil {
		fmt.Println("Failed to NewWallet", err)
		return nil, err
	}
	//2. 推导私钥
	privateKey, err := NewKeyFromMnemonic(mn)
	if err != nil {
		fmt.Println("Failed to NewKeyFromMnemonic", err)
		return nil, err
	}
	//3. 获取地址
	publicKey, err := DerivePublicKey(privateKey)
	if err != nil {
		fmt.Println("Failed to DerivePublicKey", err)
		return nil, err
	}
	//利用公钥推导地址
	address := crypto.PubkeyToAddress(*publicKey)
	//4. 创建keystore
	hdks := hdkeystore.NewHDkeyStore(keypath, privateKey)
	//5. 创建钱包
	return &HDWallet{address, hdks}, nil
}
