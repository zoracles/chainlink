pragma solidity ^0.6.0;

interface WhitelistedInterface {
  function whitelisted(address user) external returns (bool);
}
