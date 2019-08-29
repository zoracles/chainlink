pragma solidity 0.5.0;
pragma experimental ABIEncoderV2;

import "../dev/CoordinatorInterface.sol";
// import "../Oracle.sol";
/* is AggregatorInterface */
contract EmptyAggregator is AggregatorInterface {

  event InitiatedJob(bytes32 said);
  event Fulfilled(bytes32 requestId, address sender);

  /** **************************************************************************
   * @notice Called when a new job is registered
   * XXX: This depends on ABIEncoderV2 doing the right thing, when using a
   *      struct in multiple contexts...
   ****************************************************************************/
  function initiateJob(
    bytes32 _saId,
    CoordinatorInterface.ServiceAgreement memory _sa
  )
    public
  {
    emit InitiatedJob(_saId);
  }

  /** **************************************************************************
   * @notice called when an oracle reports its value for indicated request
   *
   * @dev Records the observations in the Observations table set up for this
   *      request in initiateRequest.
   *
   * @param _requestId The key used to reference this request
   * @param _oracle The address of the responding oracle
   * @param _currentObservations Array of measurements
   *
   * @return complete true iff enough observations have been recorded to
   *                  construct the summary
   *         summary list of means for each measurement requested, if
   *                 complete. Empty list otherwise.
   ****************************************************************************/
  function fulfill(bytes32 _requestId, address _oracle,
                   uint256[] memory _currentObservations)
    public
    returns (bool valid, bool complete, uint256[] memory summary)
  {
    emit Fulfilled(_requestId, _oracle);
    return (true, false, summary);
  }
}
