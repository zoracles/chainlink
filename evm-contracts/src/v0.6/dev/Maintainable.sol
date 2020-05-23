pragma solidity ^0.6.0;

/**
 * @title Maintainable
 * @notice Maintainable allows for functions to be specified by
 * callers other than a contract's owner.
 * @dev The Maintainable contract is abstract because
 * it forces the inheriting contract to create the implementation
 * for the transferMaintainer function.
 */
abstract contract Maintainable {
  address public maintainer;
  address internal pendingMaintainer;

  event MaintainerTransferRequested(
    address indexed from,
    address indexed to
  );

  event MaintainerTransferred(
    address indexed from,
    address indexed to
  );

  constructor() public {
    maintainer = msg.sender;
    emit MaintainerTransferred(
      address(0),
      msg.sender
    );
  }

  /**
   * @notice Allows the transfer of the maintainer's address to start
   * @dev virtual modifier to allow inheriting contracts to secure the function
   * with permissions other than the maintainer. For example, the owner of a
   * would likely be an appropriate user to transferring the maintainer,
   * even if the owner itself cannot call onlyMaintainer functions.
   */
  function transferMaintainer(address _to) public virtual;

  function _transferMaintainer(address _to)
    internal
  {
    pendingMaintainer = _to;

    emit MaintainerTransferRequested(maintainer, _to);
  }

  /**
   * @dev Allows a maintainer transfer to be completed by the recipient.
   */
  function acceptMaintainer()
    external
  {
    require(msg.sender == pendingMaintainer, "Must be proposed maintainer");

    address oldMaintainer = maintainer;
    maintainer = msg.sender;
    pendingMaintainer = address(0);

    emit MaintainerTransferred(oldMaintainer, msg.sender);
  }

  /**
   * @dev reverts if the caller is not the maintainer
   */
  modifier onlyMaintainer() {
    require(msg.sender == maintainer, "Not maintainer");
    _;
  }
}
