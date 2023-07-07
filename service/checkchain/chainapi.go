package checkchain

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"math/big"
)

type CollectionItem struct {
	ColAddress common.Address
	Name       string
	Symbol     string
}

type CollectionItems struct {
	Items []CollectionItem
}

type TokenMintedItem struct {
	ColAddress common.Address
	Owner      common.Address
	TokenId    *big.Int
	TokenUri   string
}

type TokenMintedItems struct {
	Items []TokenMintedItem
}

var providerUrl string
var contractAddress string
var Client *ethclient.Client

func init() {
	myEnv := make(map[string]string)
	myEnv, err := godotenv.Read()
	if err != nil {
		log.Fatal(err)
	}
	providerUrl = myEnv["provider_url"]
	contractAddress = myEnv["contract_address"]

	Client, err = ethclient.Dial(providerUrl)
	if err != nil {
		log.Fatal(err)
	}
}

func (box *CollectionItems) AddICollectionItem(item CollectionItem) []CollectionItem {
	box.Items = append(box.Items, item)
	return box.Items
}

func (box *TokenMintedItems) AddITokenMintedItem(item TokenMintedItem) []TokenMintedItem {
	box.Items = append(box.Items, item)
	return box.Items
}

func GetCollectionsCreated() []CollectionItem {
	address := common.HexToAddress(contractAddress)
	instance, err := NewCheckchain(address, Client)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	filterOpts := &bind.FilterOpts{Context: ctx, Start: 9000000, End: nil}

	iter, err := instance.FilterCollectionCreated(filterOpts)
	if err != nil {
		log.Fatal("error making event request: %s\n", err)
	}
	items := CollectionItems{}

	for iter.Next() {
		event := iter.Event
		item := CollectionItem{ColAddress: event.Collection, Name: event.Name, Symbol: event.Symbol}
		items.AddICollectionItem(item)
	}
	return items.Items
}

func GetTokenMinted(addr string) []TokenMintedItem {
	address := common.HexToAddress(contractAddress)
	instance, err := NewCheckchain(address, Client)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	filterOpts := &bind.FilterOpts{Context: ctx, Start: 9000000, End: nil}

	colAddress := common.HexToAddress(addr)
	iter, err := instance.FilterTokenMinted(filterOpts, colAddress)
	if err != nil {
		log.Fatal("error making event request: %s\n", err)
	}
	items := TokenMintedItems{}

	for iter.Next() {
		event := iter.Event
		item := TokenMintedItem{ColAddress: event.Collection, Owner: event.Owner, TokenId: event.TokenId, TokenUri: event.TokenUri}
		items.AddITokenMintedItem(item)
	}
	return items.Items
}
