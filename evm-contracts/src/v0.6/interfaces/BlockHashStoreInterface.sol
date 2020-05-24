pragma solidity 0.6.2;

interface BlockHashStoreInterface {
  function getBlockhash(uint256 number) external view returns (bytes32);
}
