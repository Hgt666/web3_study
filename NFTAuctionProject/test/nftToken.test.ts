import { expect } from "chai";
import { network } from "hardhat";


const { ethers,networkHelpers } = await network.connect();


// 定义测试Fixture
async function deployNFTTokenFixture() {
    const [owner, addr1, addr2] = await ethers.getSigners();

    const nftToken = await ethers.deployContract("NFTToken");
    await nftToken.waitForDeployment();
    await nftToken.mint();

    return {
        owner,
        addr1,
        addr2,
        nftToken
    };
}


// 测试NFTToken合约
describe("NFTToken合约测试", function () {
    it("用例1：验证铸造的第一个NFT的tokenID",async function () {
        const { owner, nftToken } = await deployNFTTokenFixture();
        const tokenId = await nftToken.createdNFTCount();
        expect(tokenId).to.equal(1);
        console.log(`铸造的NFT的tokenID:${tokenId}`)
        

    });
    it("用例2：验证第一个NFT的所有者",async function () {
        const { owner, addr1, addr2, nftToken } = await deployNFTTokenFixture();
        const tokenIdOfOwner1 = await nftToken.ownerOf(1);
        expect(tokenIdOfOwner1).to.equal(owner.address);
        console.log(`tokenId=1的所有者为：${tokenIdOfOwner1}`)
    });
    it("用例3：验证转移NFT后的所有权",async function () {
        const { owner, addr1, addr2, nftToken } = await deployNFTTokenFixture();
        await nftToken.transferNFT(addr2.address, 1);
        const afterTransferOwner = await nftToken.ownerOf(1);
        expect(afterTransferOwner).to.equal(addr2.address);
        console.log(`tokenId=1的新所有者为：${afterTransferOwner}`)
    });

});