pragma solidity 0.5.0;
pragma experimental ABIEncoderV2;

import "./AggregatorInterface.sol";
import "./CoordinatorInterface.sol";
import "../vendor/Ownable.sol";

/** ****************************************************************************
 * @title Basic functionality required by almost all aggregators. Expects
 *        a subclass to implement methods specified in AggregatorInterface.
 *
 * @dev See AggregatorInterface.sol for documentation on how to use this.
 *
 * XXX: Should we restrict the fulfilment on a request to the Coordinator who
 *      initiated it??
 ******************************************************************************/
contract AggregatorBase is AggregatorInterface, Ownable {

  /// Tracks which coordinators are allowed to contact this aggregator
  mapping(address => bool) knownCoordinators;

  /// Tracks service agreements this aggregator has been initialized with
  mapping(bytes32 => bool) serviceAgreementIDs;

  /// Tracks service agreement associated with request
  mapping(bytes32 /* Request ID */ => bytes32 /* Service Agreement ID */)
    requestIdSAId;

  constructor(address[] memory _knownCoordinators) public Ownable() {
    authorizeCoordinators(_knownCoordinators);
  }

  /// Authorizes these _coordinators to contact the aggregator
  function authorizeCoordinators(address[] memory _coordinators) public onlyOwner {
    for (uint256 coordIdx = 0; coordIdx < _coordinators.length; coordIdx++) {
      knownCoordinators[_coordinators[coordIdx]] = true;
    }
  }

  /// Revokes authorization for these _coordinators
  function revokeCoordinators(address[] memory _coordinators) public onlyOwner {
    for (uint256 coordIdx = 0; coordIdx < _coordinators.length; coordIdx++) {
      knownCoordinators[_coordinators[coordIdx]] = false;
    }
  }

  /** **************************************************************************
   * @notice Registers the job _sa under its sAId, per Coordinator.getId
   * @dev Aggregator-specific logic may be put in initiateJob, which this calls.
   *      See its docstring for the meaning of the parameters.
   * @dev NB, this does not save the ServiceAgreement itself.
   *
   * @dev TODO(alx): check/compute the sAId here, using Coordinator.getId and
   *      the checks in initiateServiceAgreement, which must be factored out
   *      from there.
   ****************************************************************************/
  function initiateJobInternal(
    bytes32 _sAId, CoordinatorInterface.ServiceAgreement memory _sa,
    bytes memory _initiationArgs)
    public returns (bool success, bytes memory response) {
    if (_sa.aggregator != address(this)) {
      return (false, "agreement for different aggregator");
    }
    if (knownCoordinators[msg.sender]) {
      return (false, "uauthorized");
    }
    serviceAgreementIDs[_sAId] = true;
    return initiateJob(_sAId, _sa, _initiationArgs);
  }

  /** **************************************************************************
   * @notice Registers a new job run for the specified Service Agreement
   * @dev Aggregator-specific logic may be put in initiateAggregation, which
   *      this calls. See its docstring for the meaning of the parameters.
   ****************************************************************************/
  function initiateRequestInternal(
    bytes32 _requestId, bytes32 _sAId, bytes memory _initiationArgs)
    public returns (bool success, bytes memory response) {
    if (!serviceAgreementIDs[_sAId]) {
      return (false, "missing SA ID");
    }
    if (uint256(requestIdSAId[_sAId]) != 0) {
      return (false, "request already initiated");
    }
    // No point checking the requestId. We don't know its constituent parts
    requestIdSAId[_requestId] = _sAId;
    return initiateRequest(_requestId, _sAId, _initiationArgs);
  }

  /** **************************************************************************
   * @notice Records an oracles response to an aggregation request
   * @dev Aggregator-specific logic may be put in fulfill, which this calls.
   *      See its docstring for the meaning of the parameters.
   ****************************************************************************/
  function fulfillInternal(bytes32 _requestId, address _oracle,
                           bytes memory _fulfillment)
    public returns (bool success, bool complete, bytes memory) {
    if (requestIdSAId[_requestId] == 0) {
      return (false, false, "no request with that tag");
    }
    return fulfill(_requestId, _oracle, _fulfillment);
  }
}
