pragma solidity 0.5.0;

// import "./CoordinatorInterface.sol";
// import "../Oracle.sol";
/* is AggregatorInterface */
contract MeanAggregator {

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

  /*****************************************************************************
   * @notice Called when a new request for data is made.
   *
   * @param _requestId The key which will be used to reference this request
   * @param _numMeasurements The number of values being measured
   * @param _oracles Addresses of the responding oracles
   *
   * @dev Creates a new Observations table for the given _requestId, sets the
   *      observationsRequired on it, and the numMeasurements, and allocates
   *      space for recording the observations.
   ****************************************************************************/
  function initiateRequest(bytes32 _requestId, uint32 _numMeasurements,
                           /* Oracle */ address[] _oracles) public {
    require(observations[_requestId].observationsRequired == 0,
      "request already initiated");
    uint32 numObs = uint32(_oracles.length);
    require(numObs < RIDICULOUS_BOUND_ON_NUMBER_OF_ORACLES, "too many oracles");
    observations[_requestId].observationsRequired = numObs;
    observations[_requestId].observations = new uint256[][](_numMeasurements);
  }

  /*****************************************************************************
   * @notice called when an oracle reports its value for indicated request
   *
   * @dev Records the observations in the Observations table set up for this
   *      request in initiateRequest.
   *
   * @param _requestId The key used to reference this request
   * @param _oracle The address of the responding oracle
   * @param _currentObservations Array of measurements
   *
   * @returns complete true iff enough observations have been recorded to
   *                   construct the summary
   *          summary list of means for each measurement requested, if
   *                  complete. Empty list otherwise.
   ****************************************************************************/
  function fulfill(bytes32 _requestId, address _oracle, uint256[] memory _currentObservations)
    public returns (bool complete, uint256[] memory summary)
  {
    Observations storage co = observations[_requestId]; // XXX: Not a copy, right?
    require(!co.oraclesSeen[_oracle], "oracle has already reported");
    require(_currentObservations.length == co.observations.length,
            "wrong number of observations");
    co.oraclesSeen[_oracle] = true;
    for (uint256 measurementIdx = 0; measurementIdx < co.observations.length; measurementIdx++) {
      co.observations[measurementIdx].push(_currentObservations[measurementIdx]);
    }
    complete = co.observations[0].length >= co.observationsRequired;
    if (complete) {
      summary = computeSummary(_requestId);
    }
  }

  /*****************************************************************************
   * @returns the mean value for each requested measurement
   ****************************************************************************/
  function computeSummary(bytes32 _requestId) internal view returns (uint256[] summary) {
    Observations storage co = observations[_requestId];
    summary = new uint256[](co.observations.length);
    for (uint256 measurementIdx = 0; measurementIdx < summary.length; measurementIdx++) {
      summary[measurementIdx] = computeAverage(co.observations[measurementIdx]);
    }    
  }

  /*****************************************************************************
   * @title The average of the values in _observations
   * @returns sum(_observations)/_observations.length, to nearest integer
   * @dev Overflow in carry prevented by RIDICULOUS_BOUND_ON_NUMBER_OF_ORACLES
   ****************************************************************************/
  function computeAverage(uint256[] memory _observations) internal pure returns (uint256 avg) {
    uint256 carry = 0;
    for (uint256 obsIdx = 0; obsIdx < _observations.length; obsIdx++) {
      avg += _observations[obsIdx] / _observations.length;
      carry += _observations[obsIdx] % _observations.length;
    }
    avg += carry / _observations.length;
  }
}
