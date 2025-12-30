// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;


contract Voting {
    mapping (string => uint256) public votes;
    string[]  public  candidateList;
    
    // 添加候选人
    function addCandidateName(string memory  candidateName) external     {
        // 判断候选人是否已存在
        for (uint i=0;i < candidateList.length; i++) {
            if (keccak256(bytes(candidateList[i]))==keccak256(bytes(candidateName))) {
                break ;
            }
        }
        candidateList.push(candidateName);
    }

    // 投票
    function vote(string memory candidateName) external  {
        votes[candidateName] +=1;
    }

    // 获取someone票数
    function getVotes(string memory candidateName) external view  returns (uint) {
        return  votes[candidateName];
    }

    // 重置所有人的票数 
    function resetVotes() external   {
        for(uint i=0; i <candidateList.length;i++  ) {
            delete votes[candidateList[i]];
            // votes[candidateList[i]]=0;
        }
    }

}


contract revertString {
    function revertStr(string memory input) external pure  returns (string memory) {
        // 计算input的长度，并新建一个数组用来存储反转
        uint input_len=bytes(input).length;
        bytes memory newString =new bytes(input_len);
        for (uint i=0;i < input_len; i++) {
            newString[i]=bytes(input)[input_len-i-1];
        }
        return string(newString);
    }
}


// contract intToRoman {
//     function integerToRoman(uint number) external pure returns (string memory) {
//         //
//         return "tt";
//     }
// }






contract mergeArray {
    function mergeArr(uint[] memory arr1,uint[] memory arr2) external pure returns (uint[] memory) {
        uint len1 =arr1.length;
        uint len2 =arr2.length;
        // 初始化数组长度
        uint[] memory newArray=new uint[](len1+len2);
        uint i=0;
        uint j=0;
        uint index=0;
        // 双指针同时遍历两个数组
        while (i < len1 && j<len2) {
            if (arr1[i] <= arr2[j]) {
                newArray[index]=arr1[i];
                i++;
            }else {
                newArray[index]=arr2[j];
                j++;
            }
            index++;
        }
        // 处理arr1剩余的元素
        while  (i < len1) {
            newArray[index]=arr1[i];
            i++;
            index++;
        }

        // 处理arr2剩余的元素
        while (j < len2) {
            newArray[index]=arr2[j];
            j++;
            index++;

        }
        return newArray;

    }
}




contract romanToInt {
    mapping (bytes1 => uint) strMap;
    constructor(){
    strMap["I"]=1;
    strMap["V"]=5;
    strMap["X"]=10;
    strMap["L"]=50;
    strMap["C"]=100;
    strMap["D"]=500;
    strMap["M"]=1000;
    }


    function romanToInteger(string memory romanstr) external  view     returns (uint) {


        bytes memory romanBytes =bytes(romanstr);
        uint romanLen =romanBytes.length;
        uint num=0;

        for (uint i=0;i < romanLen-1; i++) {
            if (strMap[romanBytes[i]] < strMap[romanBytes[i+1]] ){
                num -= strMap[romanBytes[i]];
            }else {
                num += strMap[romanBytes[i]];
            }
        }
        if (romanLen > 0) {
            num += strMap[romanBytes[romanLen-1]];
        }

        return num;
    }
}






contract binaryArrSort {
    function binaryArrSorted(uint[] calldata arrSorted,uint target) public pure returns (uint index) {
        uint left=0;
        uint right=arrSorted.length-1;
        
        while (left <= right) {
            uint mid=left + (right-left) / 2;
            uint midValue =arrSorted[mid];
            if (midValue==target) {
                return mid;
            }else if (midValue < target) {
                left = mid+1;
            }else {
                right =mid -1;
            }

        }
        return type(uint).max;
    }
}
