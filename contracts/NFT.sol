// SPDX-License-Identifier: MIT

pragma solidity 0.8.19;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721Enumerable.sol";

import "./interfaces/INFT.sol";

contract NFT is INFT, Ownable, ERC721Enumerable {

    uint256 public idCounter;
    address public controller;
    string public baseTokenURI;

    modifier onlyController() {
        require(
            controller == _msgSender(),
            "Faraway NFT: caller is not the controller"
        );
        _;
    }

    constructor(
        string memory _name,
        string memory _symbol,
        string memory _baseTokenURI
    ) ERC721(_name, _symbol) {
        baseTokenURI = _baseTokenURI;
        controller = _msgSender();
        idCounter = 1;
    }

    function version() external pure returns (string memory) {
        return "1";
    }

    function setBaseTokenURI(string memory _baseTokenURI) external override onlyController {
        baseTokenURI = _baseTokenURI;
    }

    function mint(address _owner) external override onlyController returns (uint256) {
        uint256 tokenId = idCounter;
        _mint(_owner, tokenId);
        idCounter++;

        return tokenId;
    }

    function burn(uint256 tokenId) external override onlyController {
        address owner = ERC721.ownerOf(tokenId);
        require(owner == _msgSender(), "Faraway NFT: not owner");
        _burn(tokenId);
    }

    function tokensOfOwner(address user) external override view returns (uint256[] memory) {
        uint256 tokenCount = balanceOf(user);
        if (tokenCount == 0) {
            return new uint256[](0);
        } else {
            uint256[] memory output = new uint256[](tokenCount);
            for (uint256 index = 0; index < tokenCount; index++) {
                output[index] = tokenOfOwnerByIndex(user, index);
            }
            return output;
        }
    }

    function _requireMinted(uint256 tokenId) internal override view virtual {
        require(_exists(tokenId), "Faraway NFT: invalid token ID");
    }

    function _baseURI() internal override view virtual returns (string memory) {
        return baseTokenURI;
    }

    function supportsInterface(bytes4 interfaceId) public view override(
        ERC721Enumerable,
        IERC165
    ) returns (bool) {
        return super.supportsInterface(interfaceId);
    }
}
