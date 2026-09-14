package main

import (
	"os/exec"
	"runtime"
	"runtime/debug"
	"strconv"
	"strings"

	"bytes"
	"fmt"
	"log"
	"time"

	"github.com/btcsuite/btcd/address/v2"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/chaincfg/v2"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type Sigi chan struct{}

const dur = time.Duration(time.Microsecond * 20) // 500 - 50

var wallet_eth [][20]byte
var wallet_btc_legacy [][20]byte
var wallet_btc_segwit [][20]byte

func gogo() {
	// const target int = 100_000
	// var count int = 0

	// 0x0000000000000000000000000000000000000000000000000000000000000001
	// 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEBAAEDCE6AF48A03BBFD25E8CD0364140

	// bibi := common.FromHex("0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEBAAEDCE6AF48A03BBFD25E8CD0364141")
	// bibi := common.FromHex("0x00000000000000000000000000000000000000000000000000000000000F4241")

	// Increment the 32-byte big-endian private key by 1
	/*
		for i := len(bibi) - 1; i >= 0; i-- {
			bibi[i]++
			if bibi[i] != 0 { // no carry, stop
				break
			}
			// if bibi[i] == 0, carry overflows to next byte
		}
	*/

	/*
		// Use crypto.ToECDSA instead of crypto.BytesToECDSA
		pv_key, err := crypto.ToECDSA(bibi)
		if err != nil {
			panic(err)
		}

		address := crypto.PubkeyToAddress(pv_key.PublicKey)
		fmt.Println("BASE Address:")
		fmt.Println(address.Hex())

		fmt.Println(len(bibi)) // 32
		fmt.Printf("New private key: %x\n", crypto.FromECDSA(pv_key))
	*/

	/*
		privKeyHex := "D4b38FCdD3985E9303d3a905252A023a6Ed766E3D4b38FCdD3985E9303d3a905"
		pv_key, err := crypto.HexToECDSA(privKeyHex)
		if err != nil {
			panic(err)
		}
	*/

	for {
		// Increment by 1 (big-endian carry propagation)
		/*
			for i := len(bibi) - 1; i >= 0; i-- {
				bibi[i]++
				if bibi[i] != 0 { // no carry, stop
					break
				}
				// if bibi[i] == 0, carry overflows to next byte
			}
		*/
		// fmt.Printf("%x\n", bibi)

		/*
			if count == target {
				fmt.Println(count)
				count = 0
			}
			count++
		*/

		time.Sleep(dur)

		eth_pv_key, err := crypto.GenerateKey()
		if err != nil {
			log.Fatalf("failed to generate wallet: %v", err)
		}

		// Extract the 32-byte private key.
		pv_key_bytes := crypto.FromECDSA(eth_pv_key)
		// fmt.Printf("Pv Key (Hex): %x\n", pv_key_bytes)

		/*
			// Create WIF (Wallet Import Format) for the private key
			wif, err := btcutil.NewWIF(btc_pv_key, &chaincfg.MainNetParams, true)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Bitcoin WIF:", wif.String())
			// decoded, err := btcutil.DecodeWIF(wif.String())
			// if err != nil {
				// log.Fatal(err)
			// }
			// fmt.Printf("recovered key matches: %v\n", decoded.PrivKey.Key.IsEqual(btc_pv_key.Key))
		*/

		// os.Exit(0)

		/*
			// &{{{0x7ff6af887a98} 73998015398931895160925447812487945329780121186902292289458675663667965297380 112788643387266350953586197880308544936817796570924302280476914061909657218492} 65292721544155984129360524437952035770601771185068496619711178998335110874067}
			// 905a5eeb779040801104577ba238c558ddb927c4282c850278db87fc974783d3

			eth_pv_key, err = crypto.HexToECDSA("905a5eeb779040801104577ba238c558ddb927c4282c850278db87fc974783d3")
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(eth_pv_key)
			fmt.Printf("%x\n", crypto.FromECDSA(eth_pv_key))
		*/

		//
		// eth
		//

		ad_eth := crypto.PubkeyToAddress(eth_pv_key.PublicKey)
		for i := range wallet_eth {
			if bytes.Equal(ad_eth.Bytes(), wallet_eth[i][:]) {
				fmt.Println("Ad ETH:")
				fmt.Println(ad_eth.Hex())

				// pv_key_bytes := crypto.FromECDSA(eth_pv_key)
				fmt.Println("Pv ETH:")
				fmt.Printf("%x\n", pv_key_bytes)
				fmt.Println()
			}
		}

		//
		// btc
		//

		// Convert the go-ethereum *ecdsa.PrivateKey to a btcd *btcec.PrivateKey
		// The underlying curve (secp256k1) and scalar D are the same.
		//
		// btc_pv_key, btc_pub_key := btcec.PrivKeyFromBytes(pv_key_bytes)
		_, btc_pub_key := btcec.PrivKeyFromBytes(pv_key_bytes)

		// Get the compressed public key
		compressed_pub_key := btc_pub_key.SerializeCompressed()

		// Hash the compressed public key to get its Hash160
		// This is exactly what you were doing with your code
		//
		new_hash160 := address.Hash160(compressed_pub_key)
		// fmt.Printf("Hash160 (hex): %x\n", compressed_pub_key)
		// fmt.Printf("len: %d\n", len(compressed_pub_key))

		// legacy
		//
		for i := range wallet_btc_legacy {
			if bytes.Equal(new_hash160, wallet_btc_legacy[i][:]) {
				// Create a legacy P2PKH Bitcoin mainnet address.
				addr_P2PKH, err := address.NewAddressPubKeyHash(new_hash160, &chaincfg.MainNetParams)
				if err != nil {
					log.Fatalf("Failed to create BTC Legacy ad: %v", err)
				}
				// fmt.Printf("Legacy BTC Address: %s\n", addr_P2PKH.EncodeAddress())

				fmt.Println("Ad BTC Legacy:")
				fmt.Println(addr_P2PKH.EncodeAddress())

				// pv_key_bytes := crypto.FromECDSA(eth_pv_key)
				fmt.Println("Pv BTC Legacy:")
				fmt.Printf("%x\n", pv_key_bytes)
				fmt.Println()
			}
		}

		// segwit
		//
		for i := range wallet_btc_segwit {
			if bytes.Equal(new_hash160, wallet_btc_segwit[i][:]) {
				// Create a native SegWit P2WPKH Bitcoin mainnet address.
				addr_P2WPKH, err := address.NewAddressWitnessPubKeyHash(new_hash160, &chaincfg.MainNetParams)
				if err != nil {
					log.Fatalf("Failed to create BTC SegWit ad: %v", err)
				}
				// fmt.Printf("P2WPKH BTC Ad: %s\n", addr_P2WPKH.EncodeAddress())

				fmt.Println("Ad BTC SegWit:")
				fmt.Println(addr_P2WPKH.EncodeAddress())

				// pv_key_bytes := crypto.FromECDSA(eth_pv_key)
				fmt.Println("Pv BTC SegWit:")
				fmt.Printf("%x\n", pv_key_bytes)
				fmt.Println()
			}
		}
	}
}

/* func gogo_main() {
	const target int = 100_000
	var count int = 0

	for {
		if count == target {
			fmt.Println(count)
			count = 0
		}
		count++

		time.Sleep(dur)

		pv_key, err := crypto.GenerateKey()
		if err != nil {
			log.Fatalf("failed to generate wallet: %v - pv: %v", err, pv_key)
		}

		address := crypto.PubkeyToAddress(pv_key.PublicKey)

		for i := range wallet_eth {
			if bytes.Equal(address.Bytes(), wallet_eth[i][:]) {
				fmt.Println("Ad ETH:")
				fmt.Println(address.Hex())

				pv_key_bytes := crypto.FromECDSA(pv_key)
				fmt.Println("Pv ETH:")
				fmt.Printf("%x\n", pv_key_bytes)
				fmt.Println()
			}
		}
	}
} */

func cpu() {
	out, err := exec.Command("sh", "-c", `top -bn1 | grep "Cpu(s)" | awk '{print $8}'`).Output()
	if err != nil {
		panic(err)
	}
	idle, _ := strconv.ParseFloat(strings.TrimSpace(string(out)), 64)
	fmt.Printf("CPU: %.2f%%\n", 100-idle)
	fmt.Println()
}

func read_ads() {
	// eth ads
	for _, addr := range ads_eth {
		var ad_bytes [20]byte
		copy(ad_bytes[:], common.HexToAddress(addr).Bytes())
		wallet_eth = append(wallet_eth, ad_bytes)
	}

	// btc legacy ads
	for _, addr := range ads_btc_legacy {
		// addr := "1LdRcdxfbSnmCYYNdeYpUnztiYzVfBEQeC"
		ad, err := address.DecodeAddress(addr, &chaincfg.MainNetParams)
		if err != nil {
			log.Fatalf("Failed to decode ad: %v", err)
		}
		// Type-assert to AddressPubKeyHash to get the raw hash
		ad_legacy, ok := ad.(*address.AddressPubKeyHash)
		if !ok {
			log.Fatalf("Address is not a P2PKH address")
		}
		// Get the raw 20-byte Hash160
		hash160 := ad_legacy.Hash160()
		// fmt.Printf("Hash160 (hex): %x\n", hash160)

		var ad_bytes [20]byte
		copy(ad_bytes[:], hash160[:])
		wallet_btc_legacy = append(wallet_btc_legacy, ad_bytes)
	}

	// btc segwit ads
	for _, addr := range ads_btc_segwit {
		// addr := "bc1q0j55cut9nd2c88tnnsfultdx696c8lt6n4n0su"
		ad, err := address.DecodeAddress(addr, &chaincfg.MainNetParams)
		if err != nil {
			log.Fatalf("Failed to decode ad: %v", err)
		}

		// Type-assert to AddressWitnessPubKeyHash to get the raw hash
		ad_segwit, ok := ad.(*address.AddressWitnessPubKeyHash)
		if !ok {
			log.Fatalf("Address is not a P2WPKH address")
		}
		// Get the raw 20-byte Hash160
		hash160 := ad_segwit.Hash160()
		// fmt.Printf("Hash160 (hex): %x\n", hash160)

		var ad_bytes [20]byte
		copy(ad_bytes[:], hash160[:])
		wallet_btc_segwit = append(wallet_btc_segwit, ad_bytes)
	}
}

func main() {
	fmt.Println("START")

	// chunker()
	// split()

	read_ads()

	runtime.GC()
	debug.FreeOSMemory()

	/*
		N, _ := new(big.Int).SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEBAAEDCE6AF48A03BBFD25E8CD0364141", 16)
		total := new(big.Int).Sub(N, big.NewInt(1))
		chunk_size := new(big.Int).Div(total, big.NewInt(32))
		fmt.Println(chunk_size)
	*/

	fmt.Println("ADS DONE")
	go gogo()
	go gogo()
	go gogo()
	go gogo()
	go gogo()

	time.Sleep(time.Second * 10)
	cpu()
	time.Sleep(time.Second * 10)
	cpu()

	// stop := make(Sigi)
	stop := make(chan struct{})
	<-stop
}
