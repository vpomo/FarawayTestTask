// SPDX-License-Identifier: MIT

pragma solidity 0.8.20;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

import "./interfaces/INFT.sol";
import "./interfaces/IController.sol";
import "./NFT.sol";

contract Controller is IController, Ownable {

    using EnumerableSet for EnumerableSet.AddressSet;

    EnumerableSet.AddressSet private collections;

    mapping(address => uint256) private prices;

    event CollectionCreated(address collection, string name, string symbol);
    event TokenMinted(address indexed collection, address indexed owner, string tokenUri, uint256 tokenId);

    event TokenBurned(address collection, address owner, uint256 tokenId);
    event PriceUpdated(address collection, uint256 oldPrice, uint256 newPrice);

    modifier onlyCollection(address _contract) {
        require(isCollection(_contract), "Faraway Controller: not collection");
        _;
    }

    constructor() {
        _transferOwnership(_msgSender());
    }

    function setBaseTokenURI(
        address _collection, string memory _baseTokenURI
    ) external override onlyOwner onlyCollection(_collection) {
        INFT(_collection).setBaseTokenURI(_baseTokenURI);
    }

    function setCollectionPrice(
        address _collection, uint256 _newPrice
    ) external override onlyOwner onlyCollection(_collection) {
        uint256 oldPrice = prices[_collection];
        require(oldPrice != _newPrice, "Faraway Controller: wrong price");
        prices[_collection] = _newPrice;

        emit PriceUpdated(_collection, oldPrice, _newPrice);
    }

    function createCollection(
        string memory _name, string memory _symbol, string memory _baseTokenURI
    ) external override onlyOwner {
        address newCollection = _createCollection(_name, _symbol, _baseTokenURI);
        require(collections.add(newCollection), "Faraway Controller: error created collection");

        emit CollectionCreated(newCollection, _name, _symbol);
    }

    function mint(
        address _collection, address _owner
    ) external payable override onlyCollection(_collection) {
        require(prices[_collection] == msg.value, "Faraway Controller: wrong price or ether value");
        uint256 tokenId = INFT(_collection).mint(_owner);
        string memory tokenUri = INFT(_collection).tokenURI(tokenId);

        emit TokenMinted(_collection, _owner, tokenUri, tokenId);
    }

    function burn(
        address _collection, uint256 _tokenId
    ) external onlyCollection(_collection) {
        address owner = INFT(_collection).ownerOf(_tokenId);
        require(owner == _msgSender(), "Faraway Controller: not owner");
        INFT(_collection).burn(_tokenId);

        emit TokenBurned(_collection, owner, _tokenId);
    }

    function getCollections() external override view returns(address[] memory members) {
        members = collections.values();
    }

    function isCollection(address _collection) public override view returns(bool) {
        return collections.contains(_collection);
    }

    function _createCollection(string memory _name, string memory _symbol, string memory _baseTokenURI) private returns(address) {
        address newCollection = address(new NFT(_name, _symbol, _baseTokenURI));
        return address(newCollection);
    }
}
