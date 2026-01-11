package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

// Blok yapısı tanımladım 

type Block struct {
	Index     int
	Timestamp string
	Data      string
	PrevHash  string
	Hash      string
}

// Hash hesaplama fonksiyonlarını bağladım

func calculateHash(b Block) string {
	record := string(b.Index) + b.Timestamp + b.Data + b.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

// Yeni blok üretme emirlerini verdim

func generateBlock(oldBlock Block, data string) Block {
	var newBlock Block
	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = time.Now().String()
	newBlock.Data = data
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = calculateHash(newBlock)
	return newBlock
}

func main() {
	// 1. Genesis (Baslangic) Bloğumu oluşturdum 
	genesisBlock := Block{0, time.Now().String(), "Genesis Block", "", ""}
	genesisBlock.Hash = calculateHash(genesisBlock)

	// Zinciri olusturdum
	blockchain := []Block{genesisBlock}

	fmt.Println("Blockchain baslatildi!")
	fmt.Println("-------------------------------------------------")
	fmt.Printf("Blok #0 | Hash: %s\n", genesisBlock.Hash)

	// 2. Yeni blok ekleme fonksiyonum
	newBlock := generateBlock(blockchain[0], "Ilk Transfer: 10 ETH")
	blockchain = append(blockchain, newBlock)
	fmt.Printf("Blok #1 | Hash: %s\n", newBlock.Hash)
	fmt.Printf("          Onceki Hash: %s\n", newBlock.PrevHash)
	fmt.Println("-------------------------------------------------")
}
