pragma solidity 0.5.0;

import "../Oracle.sol";
import "./LinkExInterface.sol";

contract OracleDev is Oracle {

  LinkExInterface internal priceFeed;

  mapping(bytes32 => LinkExInterface) public priceFeeds;

  constructor(address _link) public Oracle(_link) {} // solhint-disable-line no-empty-blocks

  function currentRate(bytes32 _currency) public view returns (uint256) {
    return priceFeeds[_currency].currentRate();
  }

  function setPriceFeed(address _priceFeed, bytes32 _currency) external onlyOwner {
    priceFeeds[_currency] = LinkExInterface(_priceFeed);
  }

  //////////////////////////////////////////////////////////////////////////////
  // Unused concrete methods required by abstract methods on Oracle.sol
  function oracleRequest(
    address /* _sender */,
    uint256 /* _amount */,
    bytes32 /* _sAId */,
    address /* _callbackAddress */,
    bytes4 /* _callbackFunctionId */,
    uint256 /* _nonce */,
    uint256 /* _dataVersion */,
    bytes calldata /* _data */
  ) external {} // solhint-disable-line no-empty-blocks
  function cancelOracleRequest(
    bytes32 /* requestId */,
    uint256 /* payment */,
    bytes4 /* callbackFunctionId */,
    uint256 /* expiration */
  ) external {} // solhint-disable-line no-empty-blocks
  function fulfillOracleRequest(
    bytes32 /* requestId */,
    uint256 /* payment */,
    address /* callbackAddress */,
    bytes4 /* callbackFunctionId */,
    uint256 /* expiration */,
    bytes32 /* data */
  ) external returns (bool) {} // solhint-disable-line no-empty-blocks
  function getAuthorizationStatus(address /* node */) external view returns (bool) { return false; }
  function setFulfillmentPermission(address /* node */, bool /* allowed */) external {} // solhint-disable-line no-empty-blocks
  function withdraw(address /* recipient */, uint256 /* amount */) external {} // solhint-disable-line no-empty-blocks
  function withdrawable() external view returns (uint256) {return 0;}

}
