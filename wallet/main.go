package main

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"log"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tyler-smith/go-bip39"
)

func create_mnemonic() {
	//1. entropy
	b, err := bip39.NewEntropy(128)
	if err != nil {
		log.Panic("failed to NewEntropy", err)
	}
	//2. mnemonic
	mne, err := bip39.NewMnemonic(b)
	if err != nil {
		log.Panic("failed to NewMnemonic", err)
	}
	fmt.Println(mne)
}

func test_account() {
	//1. 推导路径
	path, err := accounts.ParseDerivationPath("m/44'/60'/0'/0/4")
	if err != nil {
		panic(err)
	}
	//2. 推导seed
	nm := "cargo emotion slot dentist client hint will penalty wrestle divide inform ranch"
	seed, err := bip39.NewSeedWithErrorChecking(nm, "")
	if err != nil {
		panic(err)
	}
	//3. 获得主key
	masterKey, err := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	if err != nil {
		fmt.Println("Failed to NewMaster", err)
		return
	}
	//4. 推导私钥
	privateKey, err := DerivePrivateKey(path, masterKey)
	//5. 推导公钥
	publicKey, err := DerivePublicKey(privateKey)
	//6. 利用公钥推导地址
	address := crypto.PubkeyToAddress(*publicKey)

	fmt.Println(address.Hex())
}

//推导私钥
func DerivePrivateKey(path accounts.DerivationPath, masterKey *hdkeychain.ExtendedKey) (*ecdsa.PrivateKey, error) {
	var err error
	key := masterKey
	for _, n := range path {
		//按照路径迭代获得最终key
		key, err = key.Child(n)
		if err != nil {
			return nil, err
		}
	}
	//将key转换为ecdsa私钥
	privateKey, err := key.ECPrivKey()
	privateKeyECDSA := privateKey.ToECDSA()
	if err != nil {
		return nil, err
	}

	return privateKeyECDSA, nil
}

//推导公钥
func DerivePublicKey(privateKey *ecdsa.PrivateKey) (*ecdsa.PublicKey, error) {

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("failed to get public key")
	}
	return publicKeyECDSA, nil
}

// func main() {
// 	//create_mnemonic()
// 	test_account()
// }

func main() {
	w, err := wallet.NewWallet()
	if err != nil {
		fmt.Println("Failed to NewWallet", err)
		return
	}
	w.StoreKey("123")
}
