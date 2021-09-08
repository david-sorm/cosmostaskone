package types

import (
	"errors"
	cosmosTypes "github.com/cosmos/cosmos-sdk/types"
	"strconv"
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
