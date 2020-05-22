pragma solidity ^0.6.0;

import "../dev/AccessControlWhitelisted.sol";

contract WhitelistedTestHelper is AccessControlWhitelisted {

  int256 private value;

  constructor(int256 _value)
    public
  {
    value = _value;
  }

  function getValue()
    external
    view
    isWhitelisted()
    returns (int256)
  {
    return value;
  }

}
