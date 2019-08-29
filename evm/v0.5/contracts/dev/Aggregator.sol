pragma solidity 0.5.0;
pragma experimental ABIEncoderV2;

import "./CoordinatorInterface.sol";
// import "../Oracle.sol";
/* is AggregatorInterface */
contract MeanAggregator is CoordinatorInterface, AggregatorInterface {

  // true iff address is authorized to make calls on this aggregator
  mapping(address => bool) knownCoordinators;
  mapping(/* saId */ bytes32 => ServiceAgreement) serviceAgreements; 

  // TODO(alx): Use a running average of the observations, instead. Keeping the
  // values contained in uint256's is tricky, but ought to be possible by
  // keeping track of the remainder from the division used to calculate the
  // current average. The storage for the remainder can be the same size as the
  // storage for the number of observations required, which could save some space.

  // Table of observations, one row per observation, one column per value measured
  struct Observations {
    // Current set of observations, by measurement
    uint256[][] observations;
    uint32 observationsRequired;
    // Oracle address => true iff that oracle has contributed
    mapping (address => bool) oraclesSeen;
  }
  // requestID => Averages for that request
  mapping (bytes32 => Observations) public observations;

  // This bound on the number of oracles ensures that the carry variable in the
  // fulfill method cannot overflow, since numObs² < 2²⁵⁶, and there are at most
  // numObs remainders, each less than numObs.
  uint256 public constant RIDICULOUS_BOUND_ON_NUMBER_OF_ORACLES = 1 << 32;

  constructor(address[] memory _knownCoordinators) public {
    for (uint256 cidx = 0; cidx < _knownCoordinators.length; cidx++) {
      knownCoordinators[_knownCoordinators[cidx]] = true;
    }
  }

  modifier senderKnownCoordinator() {
    require(knownCoordinators[msg.sender], "must be registered coordinator");
    _;
  }

  function registerCoordinator(address _coordinator) senderKnownCoordinator() public {
    knownCoordinators[_coordinator] = true;
  }

  function revokeCoordinator(address _coordinator) senderKnownCoordinator() public {
    knownCoordinators[_coordinator] = false;
  }

  /** **************************************************************************
   * @notice Called when a new job is registered
   * XXX: This depends on ABIEncoderV2 doing the right thing, when using a
   *      struct in multiple contexts...
   ****************************************************************************/
  function initiateJob(bytes32 _saId, ServiceAgreement memory _sa)
    public senderKnownCoordinator
  {
    serviceAgreements[_saId] = _sa;
  }

  /** **************************************************************************
   * @notice Called when a new request for data is made.
   *
   * @param _sAId Key for the ServiceAgreement describing this job
   * @param _requestId The key which will be used to reference this request
   * @param _numMeasurements The number of values being measured
   *
   * @dev Creates a new Observations table for the given _requestId, sets the
   *      observationsRequired on it, and the numMeasurements, and allocates
   *      space for recording the observations.
   ****************************************************************************/
  function initiateRequest(bytes32 _sAId, bytes32 _requestId, uint32 _numMeasurements) public {
    require(observations[_requestId].observationsRequired == 0,
      "request already initiated");
    uint32 numObs = uint32(serviceAgreements[_sAId].oracles.length);
    require(numObs < RIDICULOUS_BOUND_ON_NUMBER_OF_ORACLES, "too many oracles");
    observations[_requestId].observationsRequired = numObs;
    observations[_requestId].observations = new uint256[][](_numMeasurements);
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
    public senderKnownCoordinator
    returns (bool valid, bool complete, uint256[] memory summary)
  {
    Observations storage co = observations[_requestId]; // XXX: copy?
    require(!co.oraclesSeen[_oracle], "oracle has already reported");
    require(_currentObservations.length == co.observations.length,
            "wrong number of observations");
    co.oraclesSeen[_oracle] = true;
    for (uint256 measurementIdx = 0; measurementIdx < co.observations.length;
         measurementIdx++) {
      co.observations[measurementIdx].push(_currentObservations[measurementIdx]);
    }
    complete = co.observations[0].length >= co.observationsRequired;
    if (complete) {
      summary = computeSummary(_requestId);
    }
  }

  /** **************************************************************************
   * @return the mean value for each requested measurement
   ****************************************************************************/
  function computeSummary(bytes32 _requestId)
    internal view returns (uint256[] memory summary) {
    Observations storage co = observations[_requestId];
    summary = new uint256[](co.observations.length);
    for (uint256 measurementIdx = 0; measurementIdx < summary.length;
         measurementIdx++) {
      summary[measurementIdx] = computeAverage(co.observations[measurementIdx]);
    }    
  }

  /** **************************************************************************
   * @notice The average of the values in _observations
   * @return sum(_observations)/_observations.length, to nearest integer
   * @dev Overflow in carry prevented by RIDICULOUS_BOUND_ON_NUMBER_OF_ORACLES,
   *      because carry < observations.length**2
   *                    < RIDICULOUS_BOUND_ON_NUMBER_OF_ORACLES**2
   *                    < 2**256.
   *      The first inequality holds because there are at observations.length
   *      terms added to carry, and each term is less than observations.length,
   *      because it's the remainder from dividing by observations.length.
   ****************************************************************************/
  function computeAverage(uint256[] memory _observations)
    internal pure returns (uint256 avg) {
    uint256 carry = 0;
    for (uint256 obsIdx = 0; obsIdx < _observations.length; obsIdx++) {
      avg += _observations[obsIdx] / _observations.length;
      carry += _observations[obsIdx] % _observations.length;
    }
    avg += carry / _observations.length;
  }
}
