pragma solidity 0.5.0;
pragma experimental ABIEncoderV2;

import "./CoordinatorInterface.sol";

/** ****************************************************************************
 * @title Abstract methods which Aggregators must implement to function with
 *        Coordinator
 *
 * @dev See tests/EmptyAggregator.sol for an example of how this is to be used.
 *
 * @dev Do not inherit directly from this contract. Inherit from AggregatorBase
 *      instead. This contract only documents the methods you need to implement.
 *      There are other convenience functions, modifiers and data structures
 *      which may be of interest in AggregatorBase.
 *
 * @dev NB: Contracts which inherit from AggregatorBase must explicitly call its
 *      constructor. Otherwise, solidity will complain that you did "not invoke
 *      an inherited contract's constructor correctly. See"
 *
 * @dev Decode _initiationArgs or _fulfillment into arbitrary types with, e.g.
 *      abi.decode(_initiationArgs, (uint256[], bool, ...)), and encode response
 *      in fulfill with, e.g., abi.encode(mean, variance).
 *
 * @dev Those encodings/decodings should be done using methods named
 *      initiateJobArgs, initiateRequestArgs, fulfillArgs, and fulfillResponse,
 *      where the Args methods take a bytes array and return the abi.decode'd
 *      args, and the Response method takes the raw response and returns the
 *      abi.encode'd bytes array. This makes it easier to type-check the
 *      Aggregator contract ABI against the Oracle contract ABI, by comparing
 *      respective method inputs to outputs. So that the methods are present in
 *      the contracts' external ABIs, they should be public.
 *
 * @dev The responses from initiateJob and initiateRequest are not expected to
 *      have complex types, so breaking out the encoding of those to their own
 *      functions is optional.
 ******************************************************************************/
contract AggregatorInterface {

  /** **************************************************************************
   * @notice Implement pluggable job initiation logic, here
   *
   * @param _sAId Hash of _sa contents
   * @param _sa Parameters for this job
   * @param _initiationArgs Arbitrary data to specify the initiation
   *
   * @return success Whether the job has been successfully initiated
   * @return response Arbitrary response data. E.g., an error message
   ****************************************************************************/
  function initiateJob(
    bytes32 _sAId, CoordinatorInterface.ServiceAgreement memory _sa,
    bytes memory _initiationArgs)
    public returns (bool success, bytes memory response);

  /** **************************************************************************
   * @notice Implement pluggable request initiation logic, here.
   *
   * @param _requestId Identifying tag for request
   * @param _sAId Hash of _sa contents, used to identify SA governing  request
   * @param _initiationArgs Arbitrary data to specify initiation
   *
   * @return success Whether the job has been successfully initiated
   * @return response Arbitrary response data. E.g., an error message
   ****************************************************************************/
  function initiateRequest(
    bytes32 _requestId, bytes32 _sAId, bytes memory _initiationArgs)
    public returns (bool success, bytes memory response);

  /** **************************************************************************
   * @notice Implement pluggable processing of oracle responses here
   *
   * @param _requestId Identifying tag for request
   * @param _fulfillment Arbitrary data containing oracle response
   *
   * @return success Whether the fulfillment was successfully processed
   * @return complete Whether the request is now complete
   * @return response Error message on success=false, or summary data on
   *                  complete=true
   ****************************************************************************/
  function fulfill(bytes32 _requestId, address _oracle,
                   bytes memory _fulfillment)
    public returns (bool success, bool complete, bytes memory response);

  // TODO: (alx) Find a way to express constraint that functions
  // initiateJobArgs, initiateRequestArgs, fulfillArgs, and fulfillResponse
  // exist, and have the right inputs or outputs, as required.
}
