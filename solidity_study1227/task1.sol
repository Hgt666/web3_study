// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Voting {
    mapping (string => uint) public vote_map;
    string[] public candidateList;

    // 添加候选人
    function addCandidate(string memory candidateName) public   {
        require(bytes(candidateName).length>0,"name is not null");
        // 判断候选人是否已存在
        for(uint i=0; i< candidateList.length;i++) {
            if (keccak256(bytes(candidateList[i]))==keccak256(bytes(candidateName))) {
                break ;
            }
        }
        candidateList.push(candidateName);

    }

    // 投票
    function vote(string memory candidateName) public     {
        vote_map[candidateName] +=1;
    }

    // 获取得票数
    function getVotes(string memory candidateName) public  view returns (uint ) {
        return vote_map[candidateName];
    }

    // 重置得票数
    function resetVotes() public {
        for (uint i=0;i < candidateList.length;i++) {
            delete   vote_map[candidateList[i]];
        }
    }

}