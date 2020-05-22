pragma solidity ^0.6.0;

import "./WhitelistedAggregatorProxy.sol";

/**
 * @title Whitelisted Payment Proxy
 * @notice Allows the owner to add and remove addresses from a whitelist
 * and checks the whitelisted status for an address on a Payment contract
 */
contract WhitelistedPaymentProxy is WhitelistedAggregatorProxy {

  WhitelistedInterface public paymentContract;

  constructor(
    address _aggregator,
    address _paymentContract
  )
    public
    WhitelistedAggregatorProxy(_aggregator)
  {
    setPaymentContract(_paymentContract);
  }

  /**
   * @notice Allows the owner to update the payment contract address.
   * @param _paymentContract The new address for the payment contract
   */
  function setPaymentContract(address _paymentContract)
    public
    onlyOwner()
  {
    paymentContract = WhitelistedInterface(_paymentContract);
  }

  /**
   * @dev reverts if the caller is not whitelisted first by
   * the payment contract, then by the local whitelisted mapping,
   * and lastly if the whitelist is enabled
   */
  modifier isWhitelisted() override {
    require(
      paymentContract.whitelisted(msg.sender) ||
      whitelisted[msg.sender] ||
      !whitelistEnabled,
      "Not whitelisted");
    _;
  }
}
