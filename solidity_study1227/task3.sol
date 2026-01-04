// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract BeggingContract {
    // 捐赠金额
    mapping (address => uint) donate_amounts;
    // 捐赠排行榜
    address[]  public donate_ranks;
    // 拥有者地址
    address public owner;

    // 新增捐赠时间限制
    uint public immutable startDonateTime;
    uint public immutable endDonateTime;


    constructor(uint _startTime,uint _endTime) {
        owner=msg.sender;
        // 开始时间必须早于结束时间
        require(_startTime < _endTime);
        startDonateTime =_startTime;
        endDonateTime = _endTime;
    }
    // 修饰器，只有owner的地址才能调用
    modifier onlyOwner {
        require(msg.sender==owner,unicode"只有owner才能调用");
        _;
    }
    
    // 只有特定时间才能校验
    modifier onlyDonateTime {
        require(block.timestamp >= startDonateTime && block.timestamp <= endDonateTime,unicode"未在规定时间内捐赠");
        _;
    }

    event Log(address,uint);

    // 捐赠
    function donate() external payable onlyDonateTime   {
        require(msg.value>0,unicode"捐赠金额必须大于0");
        donate_amounts[msg.sender] += msg.value;
        donate_ranks.push(msg.sender);
        emit Log(msg.sender, msg.value);
    }

    // 获取捐赠前三的用户
    function getRankTop3() external view   returns (address[3] memory topAddress,uint[3] memory topAmounts3) {
        // 初始化前三名的地址和金额
        address addr1;address addr2;address addr3;
        uint amout1;uint amout2;uint amout3;
        // 遍历所有捐赠者的地址
        for (uint i=0;i< donate_ranks.length;i++) {
            address currentAddr =donate_ranks[i];
            uint currentAmt =donate_amounts[currentAddr];
            // 从高到低筛选前三名
            if (currentAmt > amout1) {
                // 排行往下移动
                amout3=amout2;addr3=addr2;
                amout2=amout1;addr2=addr1;
                amout1=currentAmt;addr1=currentAddr;
            }else if (currentAmt > amout2) {
                amout3=amout2;addr3=addr2;
                amout2=currentAmt;addr2=currentAddr;
            }else {
                amout3=currentAmt;addr3=currentAddr;
            }

        }
        topAddress =[addr1,addr2,addr3];
        topAmounts3=[amout1,amout2,amout3];

    }

    receive() external payable { }
    fallback() external payable { }

    // 提款
    function withdrawAll() external  onlyOwner {
        // ETH余额大于0才能提取
        require(address(this).balance > 0 ,unicode"ETH余额为0,不能提取");
        (bool success,) = payable (owner).call{value:address(this).balance}("");
        require(success,unicode"提取失败");
    }  

    // 获取某个地址的捐赠金额
    function getDonation(address addr) external view  returns (uint) {
        return donate_amounts[addr];
    }
    


}