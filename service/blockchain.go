package service

import (
	"sync"
	"time"
)

type BlockStatus int

const (
	blockUnconfirmed BlockStatus = 0 + iota //unconfirmed block
	blockConfirmed
)

const(
	defaultChannelName = "0"
	defaultPeerId = "0"
)

type ChainHeader struct {
	chainHeight int    //height of chain
	channelName string //channel name
	peerID      string //owner peer id of chain
}

type Block struct {
	version            string //version of block
	previousBlockHash  string //hash of previous block
	merkleTreeRootHash string
	merkleTree         []*Transaction
	timeStamp          time.Time
	blockHeight        int
	blockStatus        BlockStatus
	createdPeerID      string
	signature          []byte
	blockHash		   string
}


//@@tested
func (block Block) toByte() []byte{

	//hashing에 필요한 값은
	//previousBlockHash
	//merkleTreeRootHash
	//timeStamp

	previousBlockHash := []byte(block.previousBlockHash)
	merkleTreeRootHash := []byte(block.merkleTreeRootHash)
	timeStamp := []byte(block.timeStamp.String())

	s := append(previousBlockHash,merkleTreeRootHash...)
	s = append(s,timeStamp...)

	return s
}

type BlockChain struct {
	sync.RWMutex //lock
	Header *ChainHeader //chain meta information
	Blocks []*Block     //list of bloc
}

//@@tested
func CreateNewBlockChain(channelID string,peerId string) *BlockChain{

	var header = ChainHeader{
		chainHeight: 0,
		channelName: channelID,
		peerID: peerId,
	}

	return &BlockChain{Header:&header,Blocks:make([]*Block,0)}
}

//not yet
func (bc *BlockChain)AddBlock(block *Block) (bool, error){

	// block이 confirmed 아닐경우 추가하지 않는다.
	if block.blockStatus != blockConfirmed{
		return false, nil
	}

	//lastBlock := bc.Blocks[len(bc.Blocks)]
	//hashValue := common.Hashing(block)
	//
	//append(bc.Blocks, block)

	return false, nil
}