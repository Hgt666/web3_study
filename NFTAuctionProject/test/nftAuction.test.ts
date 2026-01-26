import { expect } from "chai";
import { network } from "hardhat";


const { ethers, networkHelpers } = await network.connect();
const [deployer, user1, user2, user3] = await ethers.getSigners();

// 定义测试Fixture
async function deployNftAuctionFixture() {
    const nftToken = await ethers.deployContract("NFTToken");
    await nftToken.waitForDeployment();
    // nftToken合约的地址
    const nftTokenAddress = nftToken.getAddress();
    // 铸造NFT
    await nftToken.mint();
    const tokenId = 1;

    const NftAuction = await ethers.deployContract("NFTAuction");
    await NftAuction.waitForDeployment();

    // 授权NFTAuction合约管理NFTToken
    await nftToken.setApprovalForAll(NftAuction.getAddress(), true);


    return {
        deployer,
        nftTokenAddress,
        tokenId,
        user1,
        user2,
        user3,
        NftAuction
    }
}

// 测试NftAuction合约
describe("NftAuction合约测试", function () {

    it("创建拍卖后，查看拍卖信息", async function () {
        const { deployer, NftAuction, nftTokenAddress, tokenId } = await networkHelpers.loadFixture(deployNftAuctionFixture);
        // 创建拍卖
        let startPrice = ethers.parseEther("0.1");
        let duration = 5;
        await NftAuction.createAuction(nftTokenAddress, tokenId, startPrice, duration);
        const auctionCount = await NftAuction.auctionCount();
        // 把auctionCount转化为number类型
        const auctionId = Number(auctionCount);
        const auction = await NftAuction.auctions(auctionId);

        console.log("创建拍卖成功，拍卖ID:", auctionId);
        console.log("拍卖信息:", auction);
        expect(auctionId).to.be.equal(1n);

    });

    it("出价", async function () {
        const { deployer, NftAuction, nftTokenAddress, tokenId } = await networkHelpers.loadFixture(deployNftAuctionFixture);
        // 创建拍卖
        let startPrice = ethers.parseEther("0.1");
        let duration = 5;
        await NftAuction.createAuction(nftTokenAddress, tokenId, startPrice, duration);
        const auctionCount = await NftAuction.auctionCount();
        // 把auctionCount转化为number类型
        const auctionId = Number(auctionCount);
        const auction = await NftAuction.auctions(auctionId);


        // 切换到user1账户出价
        const bidPrice = ethers.parseEther("2");
        await NftAuction.connect(user1).bid(auctionId, { value: bidPrice });

        // 查看拍卖信息

        const auctionBid = await NftAuction.auctions(auctionId);
        console.log("更新后的拍卖信息:", auctionBid);
        expect(auctionBid.highestBidder).to.be.equal(user1.address);
        expect(auctionBid.highestBid).to.be.equal(bidPrice);


    });


    it("结束拍卖", async function () { 
                const { deployer, NftAuction, nftTokenAddress, tokenId } = await networkHelpers.loadFixture(deployNftAuctionFixture);
        // 创建拍卖
        let startPrice = ethers.parseEther("0.1");
        let duration = 5;
        await NftAuction.createAuction(nftTokenAddress, tokenId, startPrice, duration);
        const auctionCount = await NftAuction.auctionCount();
        // 把auctionCount转化为number类型
        const auctionId = Number(auctionCount);
        const auction = await NftAuction.auctions(auctionId);

        // 切换到user1账户出价
        const bidPrice = ethers.parseEther("2");
        await NftAuction.connect(user1).bid(auctionId, { value: bidPrice });

        // 修改拍卖时间，让拍卖结束
        await ethers.provider.send("evm_increaseTime", [5000]);
        await ethers.provider.send("evm_mine", []);

        // 切换回deployer账户结束拍卖
        await NftAuction.connect(deployer).endAuction(auctionId);

        // 查看拍卖信息
        const auctionInfo = await NftAuction.auctions(auctionId);
        console.log("结束后的拍卖信息:", auctionInfo);
        expect(auctionInfo.ended).to.be.equal(true);
    });





});