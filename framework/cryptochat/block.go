package cryptochat

import (
	"crypto/aes"
	"crypto/rand"
	"encoding/json"
	"io"
)

// Block passed in the cryptochat
type Block struct {
	Iv         []byte `json:"iv"`
	PaddingLen int    `json:"paddingLen"`
	Ciphertext []byte `json:"ciphertext"`
}

// NewBlock creates new block to transfer in the chat
func NewBlock(content []byte) *Block {
	var block Block

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	block.Iv = genRand(aes.BlockSize)
	block.PaddingLen = calcPadding(block.Iv, content, aes.BlockSize)
	block.Ciphertext = append(content, genRand(block.PaddingLen)...)
	return &block
}

func genRand(len int) []byte {
	s := make([]byte, len)
	if _, err := io.ReadFull(rand.Reader, s); err != nil {
		panic(err)
	}
	return s
}

func calcPadding(iv []byte, content []byte, blockSize int) int {
	return blockSize - (len(iv)+len(content))%blockSize
}

// Marshal marshals the block struct to bytes
func Marshal(block Block) ([]byte, error) {
	return json.Marshal(block)
}

// UnMarshal converts the data into bluck struct
func UnMarshal(data []byte) (block Block, err error) {
	json.Unmarshal(data, &block)
	return
}
