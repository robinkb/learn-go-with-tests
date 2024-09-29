package poker

import (
	"encoding/binary"
	"log"

	"github.com/dgraph-io/badger/v4"
)

func NewBadgerPlayerStore() (*BadgerPlayerStore, error) {
	db, err := badger.Open(badger.DefaultOptions("").WithInMemory(true))
	if err != nil {
		return nil, err
	}

	// Note that the DB is never closed, it's just toy code.

	return &BadgerPlayerStore{db}, nil
}

type BadgerPlayerStore struct {
	db *badger.DB
}

func (b *BadgerPlayerStore) GetPlayerScore(name string) int {
	var score int

	err := b.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(name))
		if err == badger.ErrKeyNotFound {
			score = 0
			return nil
		}
		if err != nil {
			return err
		}

		value, err := item.ValueCopy(nil)
		if err != nil {
			return err
		}

		score = int(binary.BigEndian.Uint16(value))
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return score
}

func (b *BadgerPlayerStore) RecordWin(name string) {
	var score int

	err := b.db.Update(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(name))
		if err != badger.ErrKeyNotFound {
			value, err := item.ValueCopy(nil)
			if err != nil {
				return err
			}

			score = int(binary.BigEndian.Uint16(value))
		}

		score++
		value := make([]byte, 10)
		binary.BigEndian.PutUint16(value, uint16(score))

		return txn.Set([]byte(name), value)
	})

	if err != nil {
		log.Fatal(err)
	}
}
