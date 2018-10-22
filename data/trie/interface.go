// Copyright 2018 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

//Adapted by Elrond Team

package trie

import (
	"github.com/ElrondNetwork/elrond-go-sandbox/storage"
)

// IdealBatchSize
// Code using batches should try to add this much data to the batch.
// The value was determined empirically.
const IdealBatchSize = 100 * 1024

// LeafCallback is a callback type invoked when a trie operation reaches a leaf
// node. It's used by state sync and commit to allow handling external references
// between account and storage tries.
type LeafCallback func(leaf []byte, parent []byte) error

// PersisterBatcher wraps all database operations. All methods are safe for concurrent use.
type PersisterBatcher interface {
	storage.Persister
	NewBatch() Batch
}

// Batch is a write-only database that commits changes to its host database
// when Write is called. Batch cannot be used concurrently.
type Batch interface {
	Put(key []byte, value []byte) error
	Delete(key []byte) error
	ValueSize() int // amount of data in the batch
	Write() error
	// Reset resets the batch for reuse
	Reset()
}

// PatriciaMerkelTree used in all tries implementations
type PatriciaMerkelTree interface {
	SetCacheLimit(l uint16)
	Get(key []byte) ([]byte, error)
	Update(key, value []byte) error
	Delete(key []byte) error
	Root() []byte
	Commit(onleaf LeafCallback) (root []byte, err error)
	DBW() DBWriteCacher
	Recreate(root []byte, dbw DBWriteCacher) (PatriciaMerkelTree, error)
	Copy() PatriciaMerkelTree
}

// DBWriteCacher used in Patricia Merkel Tree sub layer
type DBWriteCacher interface {
	PersistDB() PersisterBatcher
	InsertBlob(hash []byte, blob []byte)
	Node(hash []byte) ([]byte, error)
	Reference(child []byte, parent []byte)
	Dereference(root []byte)
	Cap(limit float64) error
	Commit(node []byte, report bool) error
	Size() (float64, float64)
	InsertWithLock(hash []byte, blob []byte, node Node)
	CachedNode(hash []byte, cachegen uint16) Node
}