# FarawayTestTask

## This is the original task

Develop a smart contract(-s) on Solidity for deploying a NFT collection (ERC721) with some arguments (name, symbol). 

The smart contract should emit the following events:
CollectionCreated(address collection, name, symbol)
TokenMinted(address collection, address recipient, tokenId, tokenUri)

Develop a simple backend server with in-memory storage to handle emitted events and serve it via HTTP.

Develop a front end demo application that interacts with the smart contract and has the following functionality:
Create a new NFT collection with specified name and symbol (from user input);
Mints a new NFT with specified collection address (only created on 3.a), tokenId, tokenUri.

=============================================================================================

## These are comments on the completion of the task 

This is deployed smart contract:
https://mumbai.polygonscan.com/address/0x652ea34de1926fc668625a4eb68a80848faa78ed#writeContract


These are the endpoints of the service:
http://127.0.0.1:8082/collections
http://127.0.0.1:8082/tokenminted/{collectionAddress}

