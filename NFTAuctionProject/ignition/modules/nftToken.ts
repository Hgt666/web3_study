import { buildModule } from "@nomicfoundation/hardhat-ignition/modules";


export default buildModule("NFTTokenModule", (m) => {
    // 部署NFTToken合约
    const nftToken = m.contract("NFTToken");
    m.call(nftToken, "mint", [])


    // 部署NFTAuction合约
    const nftAuction = m.contract("NFTAuction");
    // m.call(nftAuction, "createAuction", [nftToken,1,1,100])



    return { nftToken ,nftAuction}
});