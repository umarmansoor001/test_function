package fucntions

import (
	"container/list"
	"crypto/sha256"
	"fmt"
	"strconv"
)

type block struct {
	transaction  string
	nonce        int
	previousHash string
}

// Fucntion to create new blocks type struct and return address
func NewBlock(transaction string, nonce int, previousHash string) *block {
	var b block
	b.transaction = transaction
	b.nonce = nonce
	b.previousHash = previousHash
	return &b
}

// Function to calculate hash of a input string
func CalculateHash(stringtoHash string) string {

	return fmt.Sprintf("%x", sha256.Sum256([]byte(stringtoHash)))
}

// This function will change transaction of block
func ChangeBlock(num int, trans string, nonce int, L *list.List) {
	var counter int = 0
	for element := L.Front(); element != nil; element = element.Next() {
		//Increments in counter so that will get each block number
		counter++
		// if number of block matches with counter then change that block's transaciton and nocne
		if num == counter {
			element.Value.(*block).transaction = trans
			element.Value.(*block).nonce = nonce
			break
		}
	}
}

// Function will verify each block and displays the changed blocks
func VerifyChain(L *list.List) {
	var counter int = 0
	var curr_block_hash string
	var prev_block_hash string
	for element := L.Front(); element != nil; element = element.Next() {
		//Increments in counter so that will get each block number
		counter++
		//getting each block object
		B := element.Value.(*block)
		// To compute current blocks hash
		// we send both by concatenating transaction and nonce of current block to CalculateHash function
		curr_block_hash = CalculateHash(B.transaction + strconv.Itoa(B.nonce))
		if counter != 1 {
			if prev_block_hash != B.previousHash {
				fmt.Println()
				fmt.Println("******************************************")
				fmt.Printf("   Block %d has been changed in Blockchain", counter-1)
				fmt.Println()
				fmt.Println("******************************************")
			}
		}
		prev_block_hash = curr_block_hash
	}
}

// This Function prints all blocks in an input list
func DisplayBlocks(L *list.List) {

	var counter int = 0
	for element := L.Front(); element != nil; element = element.Next() {
		//Increments in counter so that will get each block number
		counter++
		//getting blocks one by one
		B := element.Value.(*block)
		//Displaying contents of blocks in nice way
		fmt.Println()
		fmt.Println("***********************************************************************************")
		fmt.Println("                               Block ", counter)
		fmt.Println("-----------------------------------------------------------------------------------")
		fmt.Println("  Transaction : ", B.transaction)
		fmt.Println("  Nonce       : ", B.nonce)
		fmt.Println("  CurrentHash : ", CalculateHash(B.transaction+strconv.Itoa(B.nonce)))
		fmt.Println("  PreviousHash: ", B.previousHash)
		fmt.Println("***********************************************************************************")
	}
}