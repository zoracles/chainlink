pragma solidity 0.5.0;
pragma experimental ABIEncoderV2;

// Contracts responsible for aggregating multi-oracle results inherit from this.
contract AggregatorInterface {
  // Has an `initiateRequest` and a `fulfill` method.
  //
  // For `fulfill`, the first argument must be a bytes32 for the _requestId, the
  // second the address for the reporting oracle, and the subsequent arguments
  // must be primitive types. It should return a bool, followed by the summary
  // data. The bool should indicate whether enough reports have been made, and
  // the summary data should be empty unless the bool is true. This method
  // corresponds to the aggFulfillSelector on the ServiceAgreement struc.t
  //
  // For `initiateRequest`, the first two arguments should be the bytes32 _sAId
  // and the _requestId. It should not return anything. This method corresponds
  // to the aggInitiateRequestSelector on the ServiceAgreement struct.
  //
  // For `initiateAggregatorForJob`, the first argument should be the
  // ServiceAgreement describing the job. It should not return anything. This
  // method corresponds to the aggInitiateJobSelector on the ServiceAgreement
  // struct.
  //
  // This comment should be replaced with an explicit constraint in an
  // interface, if solidity ever evolves to allow that.
  //
  //////////////////////////////////////////////////////////////////////////////
  // XXX: A nontrivial cancellation may be worthwhile, for when there's a lot of
  //      data associated with a job. Might want a method for that, too.
  //////////////////////////////////////////////////////////////////////////////
}

contract CoordinatorInterface {

  struct ServiceAgreement {
    uint256 payment;
    uint256 expiration;
    uint256 endAt;
    address[] oracles;
    // This effectively functions as an ID tag for the service agreement. It is
    // calculated as the keccak256 hash of the normalized JSON request to create
    // the ServiceAgreement, but that identity is unused.
    bytes32 requestDigest; 

    // Specification of aggregator interface. See ./Aggregate.sol for an example
    AggregatorInterface aggregator; 
    // Selectors for the interface methods must be specified, because their
    // arguments are arbitrary.
    //
    // Function selector for aggregator initiateJob method
    bytes4 aggInitiateJobSelector;
    // Function selector for aggregator initiateRequest method
    bytes4 aggInitiateRequestSelector;
    // Function selector for aggregator fulfill method
    bytes4 aggFulfillSelector;
  }

  struct OracleSignatures {
    uint8[] vs;
    bytes32[] rs;
    bytes32[] ss;
  }

  function initiateServiceAgreement(
    ServiceAgreement memory _agreement,
    OracleSignatures memory _signatures)
    public returns (bytes32);

  function fulfillOracleRequest(
    bytes32 _requestId,
    bytes32 _aggregatorArgs)
    external returns (bool);
}
