pragma solidity 0.5.0;
pragma experimental ABIEncoderV2;

import "./CoordinatorInterface.sol";
import "../interfaces/ChainlinkRequestInterface.sol";
import "../interfaces/LinkTokenInterface.sol";
import "../vendor/SafeMath.sol";
import "../vendor/Ownable.sol";

/**
 * @title The Chainlink Coordinator handles oracle service aggreements between one or more oracles
 */
contract Coordinator is ChainlinkRequestInterface, CoordinatorInterface, Ownable /* XXX: */ {
  using SafeMath for uint256;

  uint256 constant public EXPIRY_TIME = 5 minutes;
  LinkTokenInterface internal LINK;

  struct Callback {
    bytes32 sAId;
    uint256 amount;
    address addr;
    bytes4 functionId;
    uint64 cancelExpiration;
    uint8 responseCount;
    mapping(address => bool) responded;
  }

  mapping(bytes32 => Callback) private callbacks;
  mapping(bytes32 => mapping(address => bool)) private allowedOracles;
  mapping(bytes32 => ServiceAgreement) public serviceAgreements;
  mapping(address => uint256) public withdrawableTokens;

  /**
   * @notice Deploy with the address of the LINK token
   * @dev Sets the LinkToken address for the imported LinkTokenInterface
   * @param _link The address of the LINK token
   */
  constructor(address _link) public {
    LINK = LinkTokenInterface(_link);
  }

  event OracleRequest(
    bytes32 indexed sAId,
    address requester,
    bytes32 requestId,
    uint256 payment,
    address callbackAddr,
    bytes4 callbackFunctionId,
    uint256 cancelExpiration,
    uint256 dataVersion,
    bytes data
  );

  event NewServiceAgreement(
    bytes32 indexed said,
    bytes32 indexed requestDigest
  );

  event CancelOracleRequest(
    bytes32 internalId
  );

  struct OracleRequestArgs {
    address sender; // Address of requestor
    bytes4 callbackFunctionId; // Method to which response should be sent
    uint256 nonce; // Nonce used to make unique id for request
    uint256 dataVersion; // ??? XXX: 
    bytes data; // CBOR-encoded parameters for the tasks in the requested job
    bytes aggInitArgs; // Raw bytes for call to aggregator's initiation method
  }

  /** **************************************************************************
   * @notice Creates the Chainlink request
   *
   * @dev Stores the params on-chain in a callback for the request.
   *
   * Emits OracleRequest event for Chainlink nodes to detect.
   *
   * @param _payment Number of LINK due each oracle on responding to request
   * @param _sAId Identifier for ServiceAgreement covering request
   * @param _callbackAddress Contract to which summary result should be sent
   * @param _args See OracleRequestArgs for details
   ****************************************************************************/
  function oracleRequest(uint256 _payment, bytes32 _sAId,
                         address _callbackAddress, bytes calldata _args)
    external
    onlyLINK
    sufficientLINK(_payment, _sAId)
    checkCallbackAddress(_callbackAddress)
  {
    OracleRequestArgs memory a = abi.decode(_args, (OracleRequestArgs));
    bytes32 requestId = keccak256(abi.encodePacked(a.sender, a.nonce));
    require(callbacks[requestId].cancelExpiration == 0, "Must use a unique ID");

    callbacks[requestId].sAId = _sAId;
    callbacks[requestId].amount = _payment;
    callbacks[requestId].addr = _callbackAddress;
    callbacks[requestId].functionId = a.callbackFunctionId;
    callbacks[requestId].cancelExpiration = uint64(now.add(EXPIRY_TIME)); // solhint-disable-line not-rely-on-time

    require(matchesFunctionSelector(
              a.aggInitArgs,
              serviceAgreements[_sAId].aggInitiateRequestSelector),
            "must call agg initiator");
    // solhint-disable-next-line avoid-low-level-calls
    (bool success,) = address(serviceAgreements[_sAId].aggregator).call(
      a.aggInitArgs);
    require(success, "aggregation initiation failed");

    emit OracleRequest(
      _sAId,
      a.sender,
      requestId,
      _payment,
      _callbackAddress,
      a.callbackFunctionId,
      now.add(EXPIRY_TIME), // solhint-disable-line not-rely-on-time
      a.dataVersion,
      a.data);
  }

  /* */

  /** **************************************************************************
   * @notice Stores a Service Agreement which has been signed by the given oracles
   *
   * @dev Validates that each oracle has a valid signature.
   *      Emits NewServiceAgreement event.
   *      Initiates the job on the aggregator for the service agreement.
   *
   * @param _agreement The Service Agreement to be initiated
   * @param _signatures The signatures of the oracles in the agreement
   * @return The Service Agreement ID
   ****************************************************************************/
  function initiateServiceAgreement(
    ServiceAgreement memory _agreement,
    OracleSignatures memory _signatures
  )
    public
    returns (bytes32 serviceAgreementID)
  {
    require(
      _agreement.oracles.length == _signatures.vs.length &&
      _signatures.vs.length == _signatures.rs.length &&
      _signatures.rs.length == _signatures.ss.length,
      "Must pass in as many signatures as oracles"
    );
     // solhint-disable-next-line not-rely-on-time
    require(_agreement.endAt > block.timestamp, "ServiceAgreement must end later");

    serviceAgreementID = getId(_agreement);

    registerOracleSignatures(
      serviceAgreementID,
      _agreement.oracles,
      _signatures
    );

    initiateAggregatorForJob(serviceAgreementID, _agreement);

    serviceAgreements[serviceAgreementID] = _agreement;
    emit NewServiceAgreement(serviceAgreementID, _agreement.requestDigest);
  }

  /**
   * @dev Validates that each signer address matches for the given oracles
   * @param _serviceAgreementID Service agreement ID
   * @param _oracles Array of oracle addresses which agreed to the service agreement
   * @param _signatures contains the collected parts(v, r, and s) of each oracle's signature.
   */
  function registerOracleSignatures(
    bytes32 _serviceAgreementID,
    address[] memory _oracles,
    OracleSignatures memory _signatures
  )
    private
  {
    for (uint i = 0; i < _oracles.length; i++) {
      address signer = getOracleAddressFromSASignature(
        _serviceAgreementID,
        _signatures.vs[i],
        _signatures.rs[i],
        _signatures.ss[i]
      );
      require(_oracles[i] == signer, "Invalid oracle signature specified in SA");
      allowedOracles[_serviceAgreementID][_oracles[i]] = true;
    }

  }

  /**
   * @dev Recovers the address of the signer for a service agreement
   * @param _serviceAgreementID Service agreement ID
   * @param _v Recovery ID of the oracle signature
   * @param _r First 32 bytes of the oracle signature
   * @param _s Second 32 bytes of the oracle signature
   * @return The address of the signer
   */
  function getOracleAddressFromSASignature(
    bytes32 _serviceAgreementID,
    uint8 _v,
    bytes32 _r,
    bytes32 _s
  )
    private pure returns (address)
  {
    bytes32 prefixedHash = keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n32", _serviceAgreementID));
    return ecrecover(prefixedHash, _v, _r, _s);
  }

  /**
   * @dev Reverts if not sent from the LINK token
   */
  modifier onlyLINK() {
    require(msg.sender == address(LINK), "Must use LINK token");
    _;
  }

  /** **************************************************************************
   * @notice Calls aggregator job initiator
   *
   * @param _sAId Tag representing the _sa
   * @param _sa ServiceAgreement describing the job to initiate
   *
   * @dev XXX: Perhaps instead of sending the full SA over, could just send the
   *           _sAId, and let the aggregator look it up on this contract?
   ****************************************************************************/
  function initiateAggregatorForJob(bytes32 _sAId, ServiceAgreement memory _sa) internal {
    // TODO(alx): It should be possible to pull this encoding out of the
    // calldata, which would be more efficient.
    bytes memory rawEncoding = abi.encode(_sa);
    bytes memory initCall = abi.encodePacked(_sa.aggInitiateJobSelector, _sAId, rawEncoding);
    // solhint-disable-next-line avoid-low-level-calls
    (bool success,) = address(_sa.aggregator).call(initCall);
    require(success, "job aggregation init failed");
  }

  /** **************************************************************************
   * @notice Called by the Chainlink node to fulfill requests
   *
   * @dev Response must have a valid callback, and will delete the associated
   *      callback storage before calling the external contract.
   *
   * @param _requestId Fulfillment request ID, that must match the requester's
   * @param _aggregatorArgs The binary call data to send to the aggregator.
   *
   * @return True iff the external call to aggregator was successful
   ****************************************************************************/
  function fulfillOracleRequest(
    bytes32 _requestId,
    bytes calldata _aggregatorArgs
  )
    external
    isValidRequest(_requestId, _aggregatorArgs)
    returns (bool)
  {
    (bool complete, bytes memory report) = callAggregator(_requestId, _aggregatorArgs);
    registerOracleResponse(_requestId);
    if (complete) {
      Callback memory callback = callbacks[_requestId];
      bytes memory consumerArgs = abi.encodePacked(callback.functionId, _requestId, report);
      // solhint-disable-next-line avoid-low-level-calls
      (bool consumerSuccess,) = callbacks[_requestId].addr.call(consumerArgs);
      deleteCallback(_requestId);
      return consumerSuccess;
    }
    return true;
  }

  /** **************************************************************************
   * @notice Send oracle response to aggregator
   * @param _requestId Tag identifying request oracle responded to
   * @param _aggregatorArgs oracle's response, as raw bytes to be sent to the
   *        aggregator as a method call
   * @return true iff the aggregator has enough information to
   * @dev reverts if call to aggregator reverts, or _aggregatorArgs was invalid
   ****************************************************************************/
  function callAggregator(bytes32 _requestId, bytes memory _aggregatorArgs)
    private returns (bool complete, bytes memory aggregatorReport)
  {
    Callback memory callback = callbacks[_requestId];
    ServiceAgreement storage sa = serviceAgreements[callback.sAId];
    bool success;
    // solhint-disable-next-line avoid-low-level-calls
    (success, aggregatorReport) = address(sa.aggregator).call(
      _aggregatorArgs);
    require(success, "aggregator failed");
    bool valid;
    bytes memory response;
    (valid, complete, response) = parseAggregatorResponse(aggregatorReport);
    require(valid, "invalid oracle report");
  }

  /** **************************************************************************
   * @param _args Raw bytes for arguments to a method call
   * @param _selector Low-level specification of method _args is calling
   * @return true iff first four bytes of _args match _selector
   ****************************************************************************/
  function matchesFunctionSelector(bytes memory _args, bytes4 _selector)
    public pure returns (bool)
  {
    bytes4 prefix;
    assembly { // solhint-disable-line no-inline-assembly
      // Layout of (bytes _args): <uint256 lengthOfArgs><byte><byte>...</byte>
      // Layout of (bytes4 prefix): <28 zeroes><byte><byte><byte><byte>
      //
      // Since a bytes4 is represented as the lower 32 bits of a bytes32, we
      // only offset by 4 to extract the first 4 bytes. This picks up part of
      // the length word at the head of _args, but that's in the higher bytes.
      prefix := add(_args, 4)
    }
    return ((_args.length >= 4) && (prefix == _selector));
  }

  /** **************************************************************************
   * @param _args Raw bytes for arguments to a method call
   * @param _arg Expected bytes32 in _args
   * @param _offset Where the _arg is expected in the _args
   * @return true iff _args at the given _offset matches _arg
   ****************************************************************************/
  function matchesBytes32Arg(bytes memory _args, bytes32 _arg, uint256 _offset)
    public pure returns (bool)
  {
    bytes32 arg;
    // For some reason, can't see _offset in assembly block
    uint256 offset = _offset;
    assembly { // solhint-disable-line no-inline-assembly
      // Layout of (bytes _args): <uint256 lengthOfArgs><byte><byte>...</byte>
      // Layout of (bytes32 arg): <byte><byte><byte>...<byte>
      //
      // So point prefix at the memory location just after the length, plus the _offset
      arg := add(add(_args, 0x20), mload(offset))
    }
    return ((_args.length >= 0x20) && (_arg == arg));
  }

  /** **************************************************************************
   * @notice Parses the _response from aggregator fulfill method.
   * @dev Overwrites the bytes it's passed
   * @return
   *   valid: Whether the oracle arguments were valid
   *   complete: Whether the aggregator has enough responses to constuct summary
   *   consumerArgs: Raw bytes to pass to consumer, if complete
   ****************************************************************************/
  function parseAggregatorResponse(bytes memory _response)
    private pure returns (bool valid, bool complete, bytes memory consumerArgs)
  {
    uint256 actualResponseLength = _response.length - 64;
    assembly { // solhint-disable-line no-inline-assembly
      // First argument in response is (bool valid). Skip length in first word
      mstore(valid, mload(add(_response, 0x20)))
      // Second argument is (bool complete). It will be overwritten with the
      // length of the actual response, so a copy must be taken.
      let completeAddr := add(_response, 0x40)
      mstore(complete, mload(completeAddr))
      // Rewrite completeAddr with the uint256 length of the remaining bytes.
      // completeAddr is therefore the consumerArgs.
      mstore(completeAddr, mload(actualResponseLength))
      consumerArgs := completeAddr
    }
    assert(consumerArgs.length == actualResponseLength);
  }

  /** **************************************************************************
   * @notice Register correct response for _requestId from msg.sender
   ****************************************************************************/
  function registerOracleResponse(bytes32 _requestId) private {
    Callback storage callback = callbacks[_requestId];
    withdrawableTokens[msg.sender] = withdrawableTokens[msg.sender].add(
      callback.amount);
    callback.responded[msg.sender] = true;
    callback.responseCount += 1;
  }

  /** **************************************************************************
   * @notice remove callback data associated to _requestId
   * @dev only oracles involved in the service agreement can delete it
   ****************************************************************************/
  function deleteCallback(bytes32 _requestId) private {
    ServiceAgreement memory sa = serviceAgreements[callbacks[_requestId].sAId];
    // delete response records explicitly, since `delete` won't recurse into mappings
    for (uint256 oidx = 0; oidx < sa.oracles.length; oidx++) {
      delete callbacks[_requestId].responded[sa.oracles[oidx]];
    }
    delete callbacks[_requestId];
  }

  /**
   * @dev Allows the oracle operator to withdraw their LINK
   * @param _recipient is the address the funds will be sent to
   * @param _amount is the amount of LINK transfered from the Coordinator contract
   */
  function withdraw(address _recipient, uint256 _amount)
    external
    hasAvailableFunds(_amount)
  {
    withdrawableTokens[msg.sender] = withdrawableTokens[msg.sender].sub(_amount);
    assert(LINK.transfer(_recipient, _amount));
  }

  /**
   * @dev Necessary to implement ChainlinkRequestInterface
   */
  function cancelOracleRequest(bytes32, uint256, bytes4, uint256)
    external
  {} // solhint-disable-line no-empty-blocks

  /**
   * @notice Called when LINK is sent to the contract via `transferAndCall`
   * @dev The data payload's first 2 words will be overwritten by the `_sender` and `_amount`
   * values to ensure correctness. Calls oracleRequest.
   * @param _sender Address of the sender
   * @param _amount Amount of LINK sent (specified in wei)
   * @param _data Payload of the transaction
   */
  function onTokenTransfer(
    address _sender,
    uint256 _amount,
    bytes memory _data
  )
    public
    onlyLINK
    permittedFunctionsForLINK
  {
    assembly { // solhint-disable-line no-inline-assembly
      mstore(add(_data, 36), _sender) // ensure correct sender is passed
      mstore(add(_data, 68), _amount)    // ensure correct amount is passed
    }
    // solhint-disable-next-line avoid-low-level-calls
    (bool success,) = address(this).delegatecall(_data);
    require(success, "Unable to create request"); // calls oracleRequest
  }

  /**
   * @notice Retrieve the Service Agreement ID for the given parameters
      * @param _agreement contains all of the terms of the service agreement that can be verified on-chain.
   * @return The Service Agreement ID, a keccak256 hash of the input params
   */
  function getId(ServiceAgreement memory _agreement) public pure returns (bytes32)
  {
    return keccak256(
      abi.encodePacked(
        _agreement.payment,
        _agreement.expiration,
        _agreement.endAt,
        _agreement.oracles,
        _agreement.requestDigest,
        _agreement.aggregator,
        _agreement.aggInitiateJobSelector,
        _agreement.aggInitiateRequestSelector,
        _agreement.aggFulfillSelector
    ));
  }

  /**
   * @dev Reverts if the callback address is the LINK token
   * @param _to The callback address
   */
  modifier checkCallbackAddress(address _to) {
    require(_to != address(LINK), "Cannot callback to LINK");
    _;
  }

  /**
   * @dev Reverts if amount requested is greater than withdrawable balance
   * @param _amount The given amount to compare to `withdrawableTokens`
   */
  modifier hasAvailableFunds(uint256 _amount) {
    require(withdrawableTokens[msg.sender] >= _amount, "Amount requested is greater than withdrawable balance");
    _;
  }

  /** **************************************************************************
   * @dev Reverts if oracle response is invalid, given _requestId
   * @param _requestId The given request ID to check in stored `callbacks`
   ****************************************************************************/
  modifier isValidRequest(bytes32 _requestId, bytes memory _aggregatorArgs) {
    Callback storage callback = callbacks[_requestId];
    require(callback.addr != address(0), "Must have a valid requestId");
    require(allowedOracles[callback.sAId][msg.sender],
            "respondant not part of svc agrmnt");
    require(!callback.responded[msg.sender], "oracle already reported");
    ServiceAgreement storage sa = serviceAgreements[callback.sAId];
    require(matchesFunctionSelector(_aggregatorArgs, sa.aggFulfillSelector),
            "call aggregator fulfill method");
    // First argument (past the function selector) must be the requestId
    require(matchesBytes32Arg(_aggregatorArgs, _requestId, 4),
            "pass requestId to aggregator");
    _;
  }

  /**
   * @dev Reverts if amount is not at least what was agreed upon in the service agreement
   * @param _amount The payment for the request
   * @param _sAId The service agreement ID which the request is for
   */
  modifier sufficientLINK(uint256 _amount, bytes32 _sAId) {
    require(_amount >= serviceAgreements[_sAId].payment, "Below agreed payment");
    _;
  }

  /**
   * @dev Reverts if the given data does not begin with the `oracleRequest` function selector
   */
  modifier permittedFunctionsForLINK() {
    bytes4[1] memory funcSelector;
    assembly { // solhint-disable-line no-inline-assembly
      calldatacopy(funcSelector, 132, 4) // grab function selector from calldata
    }
    require(funcSelector[0] == this.oracleRequest.selector,
            "Must use whitelisted functions"); // XXX: Turn this back on!
    _;
  }
}
