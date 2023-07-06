package checkchain

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
	"strings"
)

type Result struct {
	Address          string `json:"address"`
	Topics           string `json:"topics"`
	Data             string `json:"data"`
	BlockNumber      string `json:"blockNumber"`
	BlockHash        string `json:"blockHash"`
	TimeStamp        string `json:"timeStamp"`
	GasPrice         string `json:"gasPrice"`
	GasUsed          string `json:"gasUsed"`
	LogIndex         string `json:"logIndex"`
	TransactionHash  string `json:"transactionHash"`
	TransactionIndex string `json:"transactionIndex"`
}

type ApiResponse struct {
	Status  string   `json:"status"`
	Message string   `json:"message"`
	Result  []Result `json:"result"`
}

type (
	RawABIResponse struct {
		Status  *string `json:"status"`
		Message *string `json:"message"`
		Result  *string `json:"result"`
	}
)

func GetCollections() {
	const apiKey = "5D6F7A11EGRXAM9J6NN2DVQHD2Y16SANMG"
	const providerUrl = "https://polygon-mumbai.g.alchemy.com/v2/TEu1hzP2IRfBKZ27AYRy6nfgSK_7CQ1N"
	const contractAddress = "0x652ea34de1926fc668625a4eb68a80848faa78ed"

	client, err := ethclient.Dial(providerUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	contractABI := GetContractABI(contractAddress, apiKey)
	fmt.Println(contractABI)

	events := contractABI.Events
	for key, value := range events {
		fmt.Println("key=", key, " value=", value)
	}

	address := common.HexToAddress(contractAddress)
	instance, err := NewCheckchain(address, client)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	filterOpts := &bind.FilterOpts{Context: ctx, Start: 9000000, End: nil}

	iter, err := instance.FilterCollectionCreated(filterOpts)
	for iter.Next() {
		event := iter.Event
		// Print out all caller addresses
		fmt.Println(event.Collection, " - ", event.Name, " - ", event.Symbol)
	}

	//abiDecoder := NewABIDecoder()

	//txHash := common.HexToHash("0x0b33e03492687d5e44ac6bf518a559f854791f91bd0d7db40d0b5d72b208c1c7 ")
	//
	//receipt := GetTransactionReceipt(client, txHash)
	//fmt.Println(receipt)
	//DecodeTransactionLogs(receipt, contractABI)

	//requestURL := fmt.Sprintf("https://api-testnet.polygonscan.com/api?module=logs&action=getLogs&address=0x652ea34de1926fc668625a4eb68a80848faa78ed&topic0=0x3454b57f2dca4f5a54e8358d096ac9d1a0d2dab98991ddb89ff9ea1746260617&apikey=%s", apiKey)
	//res, err := http.Get(requestURL)
	//if err != nil {
	//	fmt.Printf("error making http request: %s\n", err)
	//	os.Exit(1)
	//}
	//defer res.Body.Close()

	//resBody, err := ioutil.ReadAll(res.Body)
	//if err != nil {
	//	fmt.Printf("client: could not read response body: %s\n", err)
	//	os.Exit(1)
	//}
	//fmt.Printf("client: resBody: %s\n", resBody)
	//
	//var apiResponse ApiResponse
	//
	//err = json.Unmarshal(resBody, &apiResponse)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Printf("client: response result: %s\n", apiResponse.Message)
	//fmt.Printf("client: apiResponse.Result[0].Data: %s\n", apiResponse.Result[0].BlockNumber)
	//fmt.Printf("client: apiResponse.Result[0].Data: %s\n", apiResponse.Result[0].Data)
}

func DecodeTransactionLogs(receipt *types.Receipt, contractABI *abi.ABI) {
	for _, vLog := range receipt.Logs {
		// topic[0] is the event name
		event, err := contractABI.EventByID(vLog.Topics[0])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Event Name: %s\n", event.Name)
		// topic[1:] is other indexed params in event
		if len(vLog.Topics) > 1 {
			for i, param := range vLog.Topics[1:] {
				fmt.Printf("Indexed params %d in hex: %s\n", i, param)
				fmt.Printf("Indexed params %d decoded %s\n", i, common.HexToAddress(param.Hex()))
			}
		}
		if len(vLog.Data) > 0 {
			fmt.Printf("Log Data in Hex: %s\n", hex.EncodeToString(vLog.Data))
			outputDataMap := make(map[string]interface{})
			err = contractABI.UnpackIntoMap(outputDataMap, event.Name, vLog.Data)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Event outputs: %v\n", outputDataMap)
		}
	}
}

func GetContractABI(contractAddress, etherscanAPIKey string) *abi.ABI {
	rawABIResponse, err := GetContractRawABI(contractAddress, etherscanAPIKey)
	if err != nil {
		log.Fatal(err)
	}

	contractABI, err := abi.JSON(strings.NewReader(*rawABIResponse.Result))
	if err != nil {
		log.Fatal(err)
	}
	return &contractABI
}

func GetContractRawABI(address string, apiKey string) (*RawABIResponse, error) {
	client := resty.New()
	rawABIResponse := &RawABIResponse{}
	resp, err := client.R().
		SetQueryParams(map[string]string{
			"module":  "contract",
			"action":  "getabi",
			"address": address,
			"apikey":  apiKey,
		}).
		SetResult(rawABIResponse).
		Get("https://api-testnet.polygonscan.com/api")

	if err != nil {
		return nil, err
	}
	if !resp.IsSuccess() {
		return nil, fmt.Errorf(fmt.Sprintf("Get contract raw abi failed: %s\n", resp))
	}
	if *rawABIResponse.Status != "1" {
		return nil, fmt.Errorf(fmt.Sprintf("Get contract raw abi failed: %s\n", *rawABIResponse.Result))
	}

	return rawABIResponse, nil
}

func GetTransactionReceipt(client *ethclient.Client, txHash common.Hash) *types.Receipt {
	receipt, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		log.Fatal("GetTransactionReceipt error", err)
	}
	return receipt
}
