// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Voting {
    mapping (string => uint) public vote_map;
    string[] public candidateList;

    // 添加候选人
    function addCandidate(string memory Name) public pure {
        require(bytes(Name).length==0,"name is not null");
        // 判断候选人是否已存在
        
    }

    function vote(string memory candidateName) public pure {

    }

}