pragma solidity 0.5.0;
pragma experimental ABIEncoderV2;

import "../dev/CoordinatorInterface.sol";

contract EmptyAggregator is AggregatorInterface {

  event InitiatedJob(bytes32 said, address me);
  event Fulfilled(bytes32 requestId, address sender);

  function initiateJob(
    bytes32 _saId,
    CoordinatorInterface.ServiceAgreement memory _sa
  )
    public
  {
    emit InitiatedJob(_saId, address(this));
  }

  function fulfill(bytes32 _requestId, address _oracle,
                   uint256[] memory _currentObservations)
    public
    returns (bool valid, bool complete, uint256[] memory summary)
  {
    emit Fulfilled(_requestId, _oracle);
    return (true, false, summary);
  }
}
