// SPDX-License-Identifier: MIT

pragma solidity 0.8.20;

import "@openzeppelin/contracts/token/ERC721/extensions/IERC721Metadata.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/IERC721Enumerable.sol";

interface INFT is IERC721Metadata, IERC721Enumerable {

    function version() external pure returns (string memory);

    function setBaseTokenURI(string memory baseTokenURI) external;

    function mint(address _owner) external returns (uint256);

    function burn(uint256 tokenId) external;

    function tokensOfOwner(address user) external view returns (uint256[] memory);

}
