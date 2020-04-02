module wallet

go 1.12

replace github.com/yekai1003/bcwithme/wallet/hdwallet => ./hdwallet

replace github.com/yekai1003/bcwithme/wallet/hdkeystore => ./hdkeystore

require (
	github.com/btcsuite/btcutil v1.0.1 // indirect
	github.com/ethereum/go-ethereum v1.9.12 // indirect
	github.com/yekai1003/bcwithme/wallet/hdkeystore v0.0.0-00010101000000-000000000000 // indirect
	github.com/yekai1003/bcwithme/wallet/hdwallet v0.0.0-00010101000000-000000000000
)
