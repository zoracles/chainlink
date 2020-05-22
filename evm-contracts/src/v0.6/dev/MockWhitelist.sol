pragma solidity ^0.6.0;

import "./WhitelistedInterface.sol";

contract MockWhitelist is WhitelistedInterface {

  bool public whitelistEnabled;
  mapping(address => bool) public override whitelisted;

  constructor()
    public
  {
    whitelistEnabled = true;
  }

  /**
   * @notice makes the whitelist check enforced
   * @dev should not be part of WhitelistedInterface because
   * we don't want _every_ whitelisted contract to be forced
   * to implement this method
   */
  function enableWhitelist()
    external
  {
    whitelistEnabled = true;
  }

  /**
   * @notice makes the whitelist check unenforced
   * @dev should not be part of WhitelistedInterface because
   * we don't want _every_ whitelisted contract to be forced
   * to implement this method
   */
  function disableWhitelist()
    external
  {
    whitelistEnabled = false;
  }

  function addToWhitelist(address _user) external {
    whitelisted[_user] = true;
  }

  function removeFromWhitelist(address _user) external {
    delete whitelisted[_user];
  }

  /**
   * @dev reverts if the caller is not whitelisted
   */
  modifier isWhitelisted() {
    require(whitelisted[msg.sender] || !whitelistEnabled, "Not whitelisted");
    _;
  }
}
