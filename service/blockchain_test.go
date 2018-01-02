package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"fmt"
)

func TestCreateNewBlockChainTest(t *testing.T){

	var blockChains = CreateNewBlockChain(defaultChannelName,defaultPeerId)

	assert.Equal(t,0,len(blockChains.Blocks))
	assert.Equal(t,defaultPeerId,blockChains.Header.peerID)
	assert.Equal(t,defaultChannelName,blockChains.Header.channelName)
}


func TestSerialize(t *testing.T){
	block := Block{}
	block.merkleTreeRootHash = "test"
	block.previousBlockHash = "previoushash"
	block.timeStamp = time.Now()

	block2 := Block{}
	block2.merkleTreeRootHash = "test"
	block2.previousBlockHash = "previoushash"
	block2.timeStamp = time.Now()

	block3 := Block{}
	block3.merkleTreeRootHash = "test"
	block3.previousBlockHash = "previoushash"
	block3.timeStamp = block.timeStamp

	assert.NotEqual(t,block.toByte(),block2.toByte())
	assert.Equal(t,block.toByte(),block3.toByte())

	fmt.Print(block.toByte())
	fmt.Println(block2.toByte())
}