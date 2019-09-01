import * as h from './support/helpers'
const EmptyAggregator = artifacts.require('test/EmptyAggregator.sol')

// Check that intiateJob/initiateJobInternal,
// initiateRequest/initiateRequestInternal, fulfill/fulfillInternal have
// matching ABIs.

// Check that each internal method calls the public method, using
// EmptyAggregator inheriting from AggregatorBase.

// Check that the abi's for the translation methods are as expected.

contract('Aggregator', () => {
  let coordinator, aggregator, newServiceAgreement

  beforeEach(async () => {
    coordinator = h.defaultAccount // We act as the coordinator, in these tests.
    aggregator = await EmptyAggregator.new([coordinator])
    const partialServiceAgreement = {
      aggregator: aggregator.address
    }
    newServiceAgreement = async sA =>
      h.newServiceAgreement({ aggregator: aggregator.address, ...sA })
  })

  describe('base/subclass interaction', async () => {
    let serviceAgreement
    beforeEach(async () => {
      serviceAgreement = await newServiceAgreement()
    })
    it('executes the pluggable job initiation logic', async () => {
      const initiateJobArgs = h.getMethod(aggregator, 'initiateJobArgs')
      assert.equal(initiateJobArgs.outputs[0].type, 'uint256[]')
      const initiateData = [web3.utils.randomHex(32), web3.utils.randomHex(32)]
      const tx = await aggregator.initiateJobInternal(
        h.calculateSAID(serviceAgreement),
        h.structAsTuple(aggregator, 'initiateJobInternal', '_sa'),
        initiateData
      )
      const event = await h.getLatestEvent(aggregator)
      console.log('!!!!!!!!!!!!!!!!!!!!', tx)
//       assert(event, 'expecting an event')
//       assert.equal(event.event, 'InitiatedJob', 'initiateJob was not called')
    })
  })
})
