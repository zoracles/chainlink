pragma solidity 0.6.2;

import "./FluxAggregator.sol";

/**
 * @title Whitelisted Prepaid Aggregator contract
 * @notice This contract requires addresses to be added to a whitelist
 * in order to read the answers stored in the FluxAggregator contract
 */
contract WhitelistedAggregator is FluxAggregator {

  address private whitelistedContract;
  bool public whitelistEnabled;
  mapping(address => bool) public whitelisted;

  constructor(
    address _whitelisted,
    address _link,
    uint128 _paymentAmount,
    uint32 _timeout,
    uint8 _decimals,
    bytes32 _description
  ) public FluxAggregator(
    _link,
    _paymentAmount,
    _timeout,
    _decimals,
    _description
  ){
    whitelistedContract = _whitelisted;
  }

  function addToWhitelist(address _user) external {
    (bool status,) = whitelistedContract.delegatecall(abi.encodeWithSelector(this.addToWhitelist.selector, _user));
    require(status);
  }

  function removeFromWhitelist(address _user) external {
    (bool status,) = whitelistedContract.delegatecall(abi.encodeWithSelector(this.removeFromWhitelist.selector, _user));
    require(status);
  }

  /**
   * @notice makes the whitelist check enforced
   */
  function enableWhitelist()
    external
  {
    (bool status,) = whitelistedContract.delegatecall(abi.encodeWithSelector(this.enableWhitelist.selector));
    require(status);
  }

  /**
   * @notice makes the whitelist check unenforced
   */
  function disableWhitelist()
    external
  {
    (bool status,) = whitelistedContract.delegatecall(abi.encodeWithSelector(this.disableWhitelist.selector));
    require(status);
  }

  /**
   * @notice get the most recently reported answer
   * @dev overridden funcion to add the isWhitelisted() modifier
   */
  function latestAnswer()
    external
    view
    override
    isWhitelisted()
    returns (int256)
  {
    return _latestAnswer();
  }

  /**
   * @notice get the most recent updated at timestamp
   * @dev overridden funcion to add the isWhitelisted() modifier
   */
  function latestTimestamp()
    external
    view
    override
    isWhitelisted()
    returns (uint256)
  {
    return _latestTimestamp();
  }

  /**
   * @notice get past rounds answers
   * @dev overridden funcion to add the isWhitelisted() modifier
   * @param _roundId the round number to retrieve the answer for
   */
  function getAnswer(uint256 _roundId)
    external
    view
    override
    isWhitelisted()
    returns (int256)
  {
    return _getAnswer(_roundId);
  }

  /**
   * @notice get timestamp when an answer was last updated
   * @dev overridden funcion to add the isWhitelisted() modifier
   * @param _roundId the round number to retrieve the updated timestamp for
   */
  function getTimestamp(uint256 _roundId)
    external
    view
    override
    isWhitelisted()
    returns (uint256)
  {
    return _getTimestamp(_roundId);
  }

  modifier isWhitelisted() {
    require(whitelisted[msg.sender] || !whitelistEnabled, "Not whitelisted");
    _;
  }
}
