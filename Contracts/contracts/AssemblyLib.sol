pragma solidity ^0.8.0;

import "hardhat/console.sol";

contract AssemblyLib {
    function at(address _addr) public view returns (bytes memory o_code) {
        assembly {
            // 获取代码大小，这需要汇编语言
            let size := extcodesize(_addr)
            // 分配输出字节数组 – 这也可以不用汇编语言来实现
            // 通过使用 o_code = new bytes（size）
            o_code := mload(0x40)
            // 包括补位在内新的“memory end”
            mstore(
                0x40,
                add(o_code, and(add(add(size, 0x20), 0x1f), not(0x1f)))
            )
            // 把长度保存到内存中
            mstore(o_code, size)
            // 实际获取代码，这需要汇编语言
            extcodecopy(_addr, add(o_code, 0x20), 0, size)
        }
    }

    function add() public pure returns (int sum) {
        assembly {
            let size := 0x20
            sum := add(add(size, 0x20), 0x1f)
            let o_code := mload(0x40)
            mstore(
                0x40,
                add(o_code, and(add(add(size, 0x20), 0x1f), not(0x1f)))
            )
        }
    }

    function extractBlockNumber(
        bytes memory bkWitness
    ) public view returns (uint64) {
       require(bkWitness.length >= 32, "Data length too short");

        uint64 value;
        assembly {
            value := mload(add(bkWitness, 32))
        }

        return value;
    }
}
