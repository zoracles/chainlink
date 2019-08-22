pragma solidity ^0.5.0;

interface ChainlinkRequestInterface {
  function oracleRequest(
    uint256 payment,
    bytes32 serviceAgreementId,
    address callbackAddress,
    bytes otherArgs // Contains an OracleRequestArgs
  ) external;

  function cancelOracleRequest(
    bytes32 requestId,
    uint256 payment,
    bytes4 callbackFunctionId,
    uint256 expiration
  ) external;
}
