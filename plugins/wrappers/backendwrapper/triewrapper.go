package backendwrapper

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/trie"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/openrelayxyz/plugeth-utils/core"
)

type WrappedTrie struct {
	t state.Trie
}

func NewWrappedTrie(t state.Trie) WrappedTrie {
	return WrappedTrie{t}
}

func (t WrappedTrie) GetKey(b []byte) []byte {
	return t.t.GetKey(b)
}

func (t WrappedTrie) TryGet(key []byte) ([]byte, error) {
	return t.t.TryGet(key)
}

func (t WrappedTrie) TryGetAccount(address common.Address) (*core.StateAccount, error) {
	act, err := t.t.TryGetAccount(address)
	if err != nil {
		return nil, err
	}
	return NewWrappedStateAccount(act), nil
}

func (t WrappedTrie) TryUpdate(key, value []byte) error {
	return nil
}

func (t WrappedTrie) TryUpdateAccount(address core.Address, account *core.StateAccount) error {
	return nil
}

func (t WrappedTrie) TryDelete(key []byte) error {
	return nil
}

func (t WrappedTrie) TryDeleteAccount(address common.Address) error {
	return nil
}

func (t WrappedTrie) Hash() core.Hash {
	return core.Hash(t.t.Hash())
}

func (t WrappedTrie) Commit(collectLeaf bool) (core.Hash, *trie.NodeSet) {
	//EmptyRootHash
	return core.Hash(core.HexToHash("56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421")), nil
}

func (t WrappedTrie) NodeIterator(startKey []byte) core.NodeIterator {
	return t.t.NodeIterator(startKey)
}

func (t WrappedTrie) Prove(key []byte, fromLevel uint, proofDb ethdb.KeyValueWriter) error {
	return nil
}

type WrappedStateAccount struct {
	s *types.StateAccount
}

func NewWrappedStateAccount(s *types.StateAccount) *WrappedStateAccount {
	return &WrappedStateAccount{s}
}

// 	Nonce    uint64
// 	Balance  *big.Int
// 	Root     Hash // merkle root of the storage trie
// 	CodeHash []byte
// }

type WrappedNodeIterator struct {
	n trie.NodeInterator
}

func (n WrappedNodeIterator) Next(b bool) bool {
	return n.n.Next()
}

func (n WrappedNodeIterator) Error() error {
	return n.n.Error()
}

func (n WrappedNodeIterator) Hash() core.Hash {
	return core.Hash(n.n.Hash())
}

func (n WrappedNodeIterator) Parent() core.Hash {
	return core.Hash(n.n.Parent())
}

func (n WrappedNodeIterator) Path() []byte {
	return n.n.Path()
}

func (n WrappedNodeIterator) NodeBlob() []byte {
	return n.n.NodeBlob()
}

func (n WrappedNodeIterator) Leaf() bool {
	return n.n.Leaf()
}

func (n WrappedNodeIterator) LeafKey() []byte {
	return n.n.LeafKey()
}

func (n WrappedNodeIterator) LeafBlob() []byte {
	return n.n.LeafBlob()
}

func (n WrappedNodeIterator) LeafProof() [][]byte {
	return n.n.LeafProof()
}

func (n WrappedNodeIterator) AddResolver(trie.NodeResolver) {
	return n.n.AddResolver()
}

