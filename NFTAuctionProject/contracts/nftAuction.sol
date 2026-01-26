// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import "@openzeppelin/contracts/token/ERC721/IERC721Receiver.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721.sol";

/**
 * @title NFT拍卖合约，支持创建拍卖、出价、结束拍卖等功能
 */
contract NFTAuction is Ownable, IERC721Receiver {
    // 拍卖结构体
    struct Auction {
        address seller; // 拍卖者地址
        address nftContract; // NFT合约地址
        uint256 tokenId; // NFT代币ID
        uint256 startingPrice; // 起拍价格
        uint256 duration; // 拍卖持续时间
        uint256 endTime; // 拍卖结束时间
        address highestBidder; // 拍卖最高出价者地址
        uint256 highestBid; // 最高出价
        bool ended; // 拍卖是否结束
    }

    mapping(uint256 => Auction) public auctions; // 所有拍卖的映射表
    uint256 public auctionCount; // 拍卖数量

    // 创建拍卖事件定义
    event AuctionCreated(
        uint256 auctionId,
        address seller,
        address nftContract,
        uint256 tokenId,
        uint256 highestBid
    );

    // 出价事件定义
    event Bid(
        uint256 auctionId,
        address bidder,
        uint256 amount
    );

    // 结束拍卖事件定义
    event AuctionEnded(
        uint256 auctionId,
        address winner,
        uint256 amount
    );


    /// @dev 构造函数，初始化合约
    constructor()  Ownable(msg.sender) {}

    /**
     * @dev 创建拍卖
     * @param _nftContract NFT合约地址
     * @param _tokenId NFT代币ID
     * @param _startingPrice 起拍价格
     * @param _duration 拍卖持续时间
     */
    function createAuction(
        address _nftContract,
        uint256 _tokenId,
        uint256 _startingPrice,
        uint256 _duration
    ) public  returns (uint256) {
        require(_nftContract != address(0), unicode"NFT合约地址不能为空");
        require(_startingPrice > 0, unicode"起拍价格必须大于0");
        require(_duration * 1 minutes > 0, unicode"拍卖持续时间必须大于0");

        IERC721 nft = IERC721(_nftContract);

        // 验证所有者
        require(nft.ownerOf(_tokenId) == msg.sender, unicode"你不是NFT所有者");

        // 验证授权
        require(
            nft.getApproved(_tokenId) == address(this) ||
                nft.isApprovedForAll(msg.sender, address(this)),
            unicode"你没有授权"
        );

        // 将NFT转移到拍卖合约
        nft.safeTransferFrom(msg.sender, address(this), _tokenId);

        // 拍卖数加1
        auctionCount++;

        // 创建拍卖
        auctions[auctionCount] = Auction({
            seller: msg.sender,
            nftContract: _nftContract,
            tokenId: _tokenId,
            startingPrice: _startingPrice,
            duration: _duration* 1 minutes,
            endTime: block.timestamp + (_duration * 1 minutes),
            highestBidder: address(0),
            highestBid: 0,
            ended: false
        });

        // 触发事件
        emit AuctionCreated(
            auctionCount,
            msg.sender,
            _nftContract,
            _tokenId,
            0
        );

        return auctionCount;
    }

    /**
     * @dev 出价
     * @param _auctionId 拍卖ID
     */
    //      * @TODO 暂时只实现ETH出价，后续支持ERC20出价
    function bid(uint256 _auctionId) public payable  {
        // 缓存拍卖信息
        Auction storage auction = auctions[_auctionId];
        // 验证拍卖是否存在
        require(auction.seller != address(0), unicode"拍卖不存在");
        // 验证拍卖是否结束
        require(
            block.timestamp  < auction.endTime,
            unicode"拍卖已结束"
        );
        // 验证出价金额是否大于起拍价格
        require(
            msg.value > auction.startingPrice,
            unicode"出价金额必须大于起拍价格"
        );
        // 验证出价者是否为拍卖者
        require(msg.sender != auction.seller, "cannot buy your NFT");

        // 如果有人出价，退回上一个最高价
        if (auction.highestBidder != address(0)) {
            (bool success, ) = auction.highestBidder.call{
                value: auction.highestBid
            }("");
            require(success, "transfer failed");
        }

        // 出价金额大于当前最高出价，更新最高出价者和最高出价
        if (
            msg.value > auction.highestBid && msg.value > auction.startingPrice
        ) {
            auction.highestBidder = msg.sender;
            auction.highestBid = msg.value;
        }

        emit Bid(_auctionId, msg.sender, msg.value);
    }

    /**
     * @dev 结束拍卖
     * @param _auctionId 拍卖ID
     */
    function endAuction(uint256 _auctionId) public {
        // 缓存拍卖信息
        Auction storage auction = auctions[_auctionId];
        // 检查拍卖时间是否已到
        require(block.timestamp >= auction.endTime, unicode"拍卖未结束");
        // 结束拍卖
        auction.ended = true;
        // 转移NFT给最高出价者
        IERC721 nft = IERC721(auction.nftContract);
        nft.safeTransferFrom(
            address(this),
            auction.highestBidder,
            auction.tokenId
        );

        // 将资金转给拍卖者
        (bool success, ) = auction.seller.call{value: auction.highestBid}("");
        require(success, "transfer failed");

        emit AuctionEnded(_auctionId, auction.highestBidder, auction.highestBid);
    }


        /// @dev 实现IERC721Receiver接口，用于合约接收NFT
    function onERC721Received(
        address ,
        address ,
        uint256 ,
        bytes calldata 
    ) external pure  virtual returns (bytes4) {
        return this.onERC721Received.selector;
    }


}
