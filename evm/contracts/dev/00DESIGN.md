# Design for multi-word, pluggable coordinator

We want to modularize the responsibilities of the coordinator, separating the
balance-management logic from the aggregation logic. Broadly speaking, this can
be done at the compiler level through mix-in classes and libraries, or via
callbacks to the aggregation logic.

The drawback with using compiler-based composition is that every new composition
will require a new contract deployment. Therefore I'm exploring callback-based
approaches, here.

A significant risk with using callbacks is that it surrenders control flow to
potentially hostile code. If we use the low-level `call` function, that at least
prevents the external code from interrupting the local flow, because it simply
returns false on failure. But that does not protect against re-entrancy attacks.
So if we use `call` route, we have to ensure idempotency of any changes to the
coordinator state which take place in concert with calls to external code. Since
the Coordinator code is under our control, this is a tolerable risk. It means
that we need explicit tests of that idempotency, though.

## Workflow

### Reporting

Cover this first, because it will determine what the consumer needs to specify
when initiating a service agreement.

Once the off-chain oracle has the result it intends to report, it calls
`Coordinator.fulfillOracleRequest(bytes32 _requestId, byte[] _calldata)` via
EthTxEncode. The `_calldata` array contains the call to the `Aggregator.fulfill`
method, including the function selector.

On chain, `Coordinator.fulfillOracleRequest`, first verifies that the oracle
address is valid, and no report from it is already recorded for this _requestId,
then calls `Aggregator.fulfill` using the [assembly `call` method
](https://solidity.readthedocs.io/en/v0.5.3/assembly.html) (The `call`
convenience function has been removed from solidity 5.) If that method succeeds,
it registers the Oracle's payment as redeemable.

The result from `Aggregator.fulfill` can be read out using [this approach
](https://stackoverflow.com/questions/45930533/solidity-get-return-value-of-delegatecall-with-assembly).
If the `Aggregator` has enough reports to provide a summary,
`Coordinator.fulfill` also calls the consuming contract's `fulfill` method.

The arguments to `Aggregator.fulfill` are arbitrary. It is up to the off-chain
oracle to correctly serialize the bytes for the call, via `EthTxEncode`. Since
that uses the `abi` package from `go-ethereum`, that constrains us to static
types in the sense of the [abi docs
](https://solidity.readthedocs.io/en/develop/abi-spec.html).

### Initiation

- If necessary, `Aggregator` is deployed. This is a contract satisfying
  `AggregatorInterface`, which offers a `fulfill` method.
- The client calls `service_agreement/create` on each of the Oracles it wants to
  participate in the aggregation. The Oracle sends back its signature, if it
  agrees to the contract. The new argument to `create` is the aggregator
  address. There ought to be a whitelist of aggregator addresses the Oracle is
  prepared to work with.
- The coordinator calls `Coordinator.initiateServiceAgreement(ServiceAgreement sa)`
  where ServiceAgreement is
  
  ```diff
  struct ServiceAgreement {
    uint256 payment;
    uint256 expiration;
    uint256 endAt;
    address[] oracles;
    bytes32 requestDigest;
  + // Address of `Aggregator`-subclass contract, which has a `fulfill` method. The
  + // input signature for this function can be essentially arbitrary, except that
  + // the first argument must be a bytes32 for the _requestId, and the subsequent
  + // arguments must be primitive types. It is up to the oracle to construct the
  + // bytestring corresponding to the call it intends. Return signature is
  + // (bool,bool,bytes). The first return value should be true iff the values it's
  + // called with are valid. The second should be true iff the Aggregator has
  + // enough reports to provide a summary. The third is empty unless the second is
  + // true, and then it contains the call bytes for `fulfill` on the consuming
  + // method.
  + address aggregator;
  }
  ```

  Here the diff symbols are comparing to the corresponding `Coordinator.sol`
  structure.

  The `requestDigest` field corresponds to [the service agreement's ID
  ](https://github.com/smartcontractkit/chainlink/blob/57844fce45eca85a8264e98ff4c2cf476746916e/core/store/models/service_agreement.go#L105).
  
- The 
## Requesting new data

The consumer implicitly calls the following via the `onTokenTransfer` mechanism:

```javascript
  function oracleRequest(
    address _sender,
    uint256 _amount,
    bytes32 _sAId,
    address _callbackAddress,
    bytes4 _callbackFunctionId,
    uint256 _nonce,
    uint256 _dataVersion,
    bytes _data
  )
```

The `_data` argument here refers to the inputs to the adapters in the job spec.
This is exactly the same as the `oracleRequest` in the standard `Oracle.sol`.

This triggers an event, `OracleRequest`, which the oracles are subscribed to.

## Reporting by a single oracle.

The oracles call this method on the `Oracle` constract:

```javascript
  function fulfillOracleRequest(
    bytes32 _requestId,
-   bytes32 _data
+   // Call to 
+   bytes _data
    )
```

This method calls `fulfill` on the `Aggregator` contract. See the notes
[above](#reporting).
