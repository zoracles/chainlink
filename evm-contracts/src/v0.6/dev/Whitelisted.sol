pragma solidity ^0.6.0;

import "../Owned.sol";
import "./WhitelistedInterface.sol";

/**
 * @title Whitelisted
 * @notice Allows the owner to add and remove addresses from a whitelist
 */
abstract contract Whitelisted is WhitelistedInterface {

  mapping(address => bool) public override whitelisted;

  // function addToWhitelist(address _user) public virtual;
  // function removeFromWhitelist(address _user) public virtual;

  /**
   * @dev reverts if the caller is not whitelisted
   */
  modifier isWhitelisted() virtual {
    require(whitelisted[msg.sender], "Not whitelisted");
    _;
  }
}
