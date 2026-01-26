// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/access/Ownable.sol";


/**
 * @title 使用 ERC721 标准的 NFT 代币
*/ 

contract NFTToken is ERC721, Ownable {
    // 最大 NFT 数量
    uint256 public constant MAX_NFT_COUNT = 10000;
    // 已铸造的 NFT 数量
    uint256 public createdNFTCount;

    /// @dev 构造函数，设置 NFT 名称和符号
    constructor() ERC721("NFTToken", "NFT") Ownable(msg.sender) {}

    /// @dev 铸造一个 NFT 代币
    function mint() public onlyOwner  returns (uint256) {
        require(createdNFTCount < MAX_NFT_COUNT,unicode"已超过最大铸造量,最大铸造量为10000");
        createdNFTCount++;
        uint256 tokenId = createdNFTCount;
        _safeMint(msg.sender, tokenId);
        return tokenId;


    }

    // 转移 NFT 代币
    function transferNFT(address to, uint256 tokenId) public onlyOwner {
        _transfer(msg.sender, to, tokenId);
    }

}