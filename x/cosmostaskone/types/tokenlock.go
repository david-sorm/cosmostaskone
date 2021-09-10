package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	"math/rand"
	"strconv"
)

const keyPrefix = "TokenLock-"

func WithPrefix(str string) []byte {
	return []byte(keyPrefix + str)
}

// returns the StartNode, a node used as a starting point for the linked list
// if the StartNode doesn't exist, it gets created automatically
func TokenLockStartNode(store prefix.Store, cdc codec.Marshaler) TokenLockInternal {
	var tli TokenLockInternal

	if !store.Has(WithPrefix("StartNode")) {
		bz := cdc.MustMarshalBinaryBare(&tli)
		store.Set(WithPrefix("StartNode"), bz)
		return tli
	}

	bz := store.Get(WithPrefix("StartNode"))
	cdc.MustUnmarshalBinaryBare(bz, &tli)
	return tli
}

// returns the last available TokenLock node
func (tl TokenLockInternal) Last(store prefix.Store, cdc codec.Marshaler) TokenLockInternal {
	for len(tl.NextNode) != 0 {
		tl = tl.Next(store, cdc)
	}
	return tl
}

// returns the following node after this node
func (tl TokenLockInternal) Next(store prefix.Store, cdc codec.Marshaler) TokenLockInternal {
	if len(tl.NextNode) == 0 {
		return tl
	}
	res := store.Get(WithPrefix(tl.NextNode))
	cdc.MustUnmarshalBinaryBare(res, &tl)
	return tl
}

// generates a unique id for a tokenlock and ensures it is actually not used already
func (tl *TokenLockInternal) GenerateUniqueID(store prefix.Store) {
	hash := ""

	for unique := false; unique != true; unique = store.Has(WithPrefix(hash)) {
		hash = ""

		for i := 0; i < 32; i++ {
			r := rand.Intn(59)
			if r < 25 {
				hash += string(rune(r + 97))
			} else if r < 50 {
				hash += string(rune(r + 65 - 25))
			} else {
				hash += strconv.Itoa(r - 50)
			}
		}
	}

	tl.ID = hash
}

// saves the updated TokenLock node to the DB automatically according to it's ID
func (tl TokenLockInternal) Save(store prefix.Store, cdc codec.Marshaler) {
	if len(tl.ID) == 0 {
		panic("no ID specified!")
		return
	}

	bz := cdc.MustMarshalBinaryBare(&tl)
	store.Set(WithPrefix(tl.ID), bz)
}

// fetches a tokenlock from the db by the id
func TokenLockLoad(store prefix.Store, cdc codec.Marshaler, id string) TokenLockInternal {
	tl := TokenLockInternal{}
	bz := store.Get(WithPrefix(id))
	cdc.MustUnmarshalBinaryBare(bz, &tl)
	return tl
}