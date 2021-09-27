package types

import (
	"errors"
	cosmosTypes "github.com/cosmos/cosmos-sdk/types"
	"strconv"
	"strings"
)

var _ = strconv.Itoa(0)

type CoinRaw struct {
	Amount string
	Denom  string
}

type CoinsRaw []CoinRaw

func (cr CoinsRaw) ParseCoins() ([]*cosmosTypes.Coin, error) {
	var amount int64
	var err error
	coins := make([]*cosmosTypes.Coin, 0, len(cr))

	for _, v := range cr {
		coin := cosmosTypes.Coin{}
		amount, err = strconv.ParseInt(v.Amount, 10, 64)
		if err != nil {
			return nil, err
		}
		if v.Denom == "" {
			return nil, errors.New("error: no denom specified")
		}

		coin.Denom = v.Denom
		coin.Amount = cosmosTypes.NewInt(amount)
		coins = append(coins, &coin)
	}

	return coins, nil
}

func DereferenceCoinSlice(toDereference []*cosmosTypes.Coin) []cosmosTypes.Coin {
	dereferenced := make([]cosmosTypes.Coin, 0, len(toDereference))
	for _, v := range toDereference {
		dereferenced = append(dereferenced, cosmosTypes.Coin{
			Denom:  v.Denom,
			Amount: v.Amount,
		})
	}
	return dereferenced
}

func SortCoinsAlphabetically(input []cosmosTypes.Coin) []cosmosTypes.Coin {
	coins := make([]cosmosTypes.Coin, len(input), len(input))
	copy(coins, input)

	sorted := false
	res := 0

	// just a simple sorting algo
	// TODO if it's worth it, implement something more efficient, maybe insertion sort
	for !sorted {
		sorted = true

		for i := 0; i < len(coins)-1; i++ {
			res = strings.Compare(coins[i].Denom, coins[i+1].Denom)
			// if coins[i].Denom is greater string alphabetically than coins[i+1].Denom, flip them
			if res == 1 {
				coins[i], coins[i+1] = coins[i+1], coins[i]
				sorted = false
			}
		}
	}
	return coins
}
