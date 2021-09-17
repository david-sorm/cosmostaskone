package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	"math/rand"
	"strconv"
)

const keyPrefix = "TL-"

var hashDict = []rune{
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '0',
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
}

func WithPrefix(str string) []byte {
	return []byte(keyPrefix + str)
}

// TokenLockStartNode returns the StartNode, a node used as a starting point for the linked list
// if the StartNode doesn't exist, it gets created automatically
func TokenLockStartNode(store prefix.Store, cdc codec.Marshaler) TokenLockInternal {
	var tli TokenLockInternal

	// if there is no startnode yet, create a new one
	if !store.Has(WithPrefix("StartNode")) {
		tli.ID = "StartNode"
		bz := cdc.MustMarshalBinaryBare(&tli)
		store.Set(WithPrefix("StartNode"), bz)
		return tli
	}

	bz := store.Get(WithPrefix("StartNode"))
	cdc.MustUnmarshalBinaryBare(bz, &tli)
	return tli
}

// GenerateUniqueID generates a unique id for a tokenlock and ensures it is actually not used already
func (tl *TokenLockInternal) GenerateUniqueID(store prefix.Store) {
	hash := ""

	for unique := false; unique != true; unique = !store.Has(WithPrefix(hash)) {
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

// GenerateKeyForTokenLock generates key in this format: {Global Prefix-}{Creator in Bech32}-{Shortest Unique Hash}
// Requires TokenLockInternal.Creator to be correctly filled out with Creator's Bech32 address
func (tl *TokenLockInternal) GenerateKeyForTokenLock(store store.KVStore) {
	newHash := ""

	// check if we already made an id for creatpr or not
	// ({Global Prefix-}{Creator in Bech32} contains currently highest assigned hash to a specific creator)
	if !store.Has(WithPrefix(tl.Creator)) {
		newHash = string(hashDict[0])
	} else {
		bz := store.Get(WithPrefix(tl.Creator))
		lastHash := string(bz)
		runeInput := []rune(lastHash)

		// prepend '@' to the beginning, so the hash has space to not overflow
		runeInput = append([]rune("@"), runeInput...)

		nextCharacterInSet(runeInput, hashDict)

		// check if '@' is still present and if it is, delete it
		if runeInput[0] == '@' {
			runeInput = runeInput[1:]
		}
		newHash = string(runeInput)
	}

	// save the new hash for future use
	store.Set(WithPrefix(tl.Creator), []byte(newHash))
	newHash = string(WithPrefix(tl.Creator+"-")) + newHash
	tl.ID = newHash
}

func nextCharacterInSet(input []rune, set []rune) []rune {
	endOfSetRune := set[len(set)-1]
	inputIndex := len(input) - 1

	currentRune := func() rune {
		return input[inputIndex]
	}

	findRuneInSet := func(r rune) int {
		for index, value := range set {
			if r == value {
				return index
			}
		}
		return -1
	}
	for {
		if currentRune() != endOfSetRune {
			indexInSet := findRuneInSet(currentRune())
			input[inputIndex] = set[indexInSet+1]
			return input
		} else {
			input[inputIndex] = set[0]
			inputIndex--
			if inputIndex == -1 {
				return nil
			}
			continue
		}
	}
}

// Save saves the updated TokenLock node to the DB automatically according to its ID
func (tl TokenLockInternal) Save(store prefix.Store, cdc codec.Marshaler) {
	if len(tl.ID) == 0 {
		panic("no ID specified!")
		return
	}

	bz := cdc.MustMarshalBinaryBare(&tl)
	store.Set(WithPrefix(tl.ID), bz)
}

// TokenLockLoadIfExists is the same as TokenLockLoad, however it first checks for the existence of the token lock,
// before getting it. Useful when there is a chance that a given lock doesn't exist.
func TokenLockLoadIfExists(store prefix.Store, cdc codec.Marshaler, id string) (TokenLockInternal, bool) {
	if !store.Has(WithPrefix(id)) {
		return TokenLockInternal{}, false
	}
	return TokenLockLoad(store, cdc, id), true
}

// TokenLockLoad fetches a tokenlock from the db by the id
func TokenLockLoad(store prefix.Store, cdc codec.Marshaler, id string) TokenLockInternal {
	tl := TokenLockInternal{}
	bz := store.Get(WithPrefix(id))
	cdc.MustUnmarshalBinaryBare(bz, &tl)
	return tl
}
