pragma solidity 0.5.0;
pragma experimental ABIEncoderV2;

// Contracts responsible for aggregating multi-oracle results inherit from this.
contract AggregatorInterface {
  // Has an `initiateRequest` and a `fulfill` method.
  //
  // For `fulfill`, the first argument must be a bytes32 for the _requestId,
  // the second the address for the reporting oracle, and the subsequent
  // arguments must be primitive types. It should return a bool, followed by
  // the summary data. The bool should indicate whether enough reports have
  // been made, and the summary data should be empty unless the bool is true.
  //
  // For `initiateRequest`, the first argument should be the bytes32
  // _requestId, the second argument an address[] containing the participating
  // oracles. It should not return anything.
  //
  // This comment should be replaced with an explicit constraint in an
  // interface, if solidity ever evolves to allow that.
}

contract CoordinatorInterface {

  struct ServiceAgreement {
    uint256 payment;
    uint256 expiration;
    uint256 endAt;
    address[] oracles;
    bytes32 requestDigest;
    // See ./Aggregate.sol for an example
    AggregatorInterface aggregator;
    // Function selector for aggregator initiateRequest method
    bytes4 aggInitiateSelector;
    // Function selector for aggregator fulfill method
    bytes4 aggFulfillSelector;
  }

  struct OracleSignatures {
    uint8[] vs;
    bytes32[] rs;
    bytes32[] ss;
  }

  function initiateServiceAgreement(ServiceAgreement memory _agreement, OracleSignatures memory _signatures) public returns (bytes32);

  function fulfillOracleRequest(bytes32 requestId, bytes32 data) external returns (bool);
}
