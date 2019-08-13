pragma solidity 0.4.24;

contract MeanAggregator {

  // Keeps the average of a number of observations
  struct RunningAverage {
    // Average of current set of observations, to nearest integer
    uint256[] runningAverage;
    // Remainder from division used to calculate average. Since divisor is at
    // most observationsrequired, the remainder can be the same size.
    uint32[] remainder;
    // Number of observations in the average so far
    uint32 numObservations;
    // Total number of observations needed to report summary average
    uint32 observationsrequired;
  }
  // requestID => Averages for that request
  mapping (bytes32 => RunningAverage) public runningAverage;

  // initiateRequest creates a new RunningAverage for the given _requestId, 
  // sets the observationsRequired on it, and the numMeasurements, and allocates
  // space for tracking the averages.
  function initiateRequest(bytes32 _requestId, uint256 _observationsRequired,
                           uint256 _numMeasurements)
    public {
      require(_observationsRequired > 0, "need at least one observation");
      // require(runningAverage[_requestId].observationsRequired == 0,
      // "observationsRequired already set");
      runningAverage[_requestId].observationsRequired = _observationsRequired;
      runningAverage[_requestId].runningAverage = new uint256[](_numMeasurements);
      runningAverage[_requestId].remainder = new uint256[](_numMeasurements);
    }
  function fulfill(bytes32) pure {}
}
