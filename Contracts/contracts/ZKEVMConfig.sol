// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

contract ZKEVMConfig {
    struct BatchData {
        uint64 blockNumber; // Block number
        bytes transactions; // RLP encoded transactions
        bytes blockWitness; // Contains each block
        bytes32 preStateRoot; // Pre-state root
        bytes32 postStateRoot; // Post-state root
        bytes32 withdrawRoot; // withdraw trie tree
        BatchSignature signature; // Signature of the batch
    }

    struct BatchSignature {
        bytes[] signers; // all signers
        bytes signature; // aggregate signature
    }

    function decode(bytes memory data) public pure returns (uint64){
        require(data.length >= 32, "bkWitness length too short");

        uint64 value;
        assembly {
            value := mload(add(data, 32))
        }
        return value;
    }
    
    function applyL1ToL2Alias(address l1Address) public pure returns (address l2Address) {
        uint160 offset = uint160(0x1111000000000000000000000000000000001111);
        unchecked {
            l2Address = address(uint160(l1Address) + offset);
        }
    }
}
