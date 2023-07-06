// SPDX-License-Identifier: MIT

pragma solidity 0.8.19;

interface IController {

    function setBaseTokenURI(address _collection, string memory _baseTokenURI) external;

    function setCollectionPrice(address _collection, uint256 _newPrice) external;

    function createCollection(
        string memory _name, string memory _symbol, string memory _baseTokenURI
    ) external;

    function mint(address _collection, address _owner) external payable;

    function burn(address _collection, uint256 _tokenId) external;

    function getCollections() external view returns(address[] memory members);

    function isCollection(address _collection) external view returns(bool);

}
