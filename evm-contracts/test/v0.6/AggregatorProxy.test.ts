import {
  contract,
  helpers as h,
  matchers,
  setup,
} from '@chainlink/test-helpers'
import { assert } from 'chai'
import { ethers } from 'ethers'
import { MockHistoricAggregatorFactory } from '../../ethers/v0.6/MockHistoricAggregatorFactory'
import { MockAggregatorFactory } from '../../ethers/v0.6/MockAggregatorFactory'
import { AggregatorProxyFactory } from '../../ethers/v0.6/AggregatorProxyFactory'
import { AggregatorFacadeFactory } from '../../ethers/v0.6/AggregatorFacadeFactory'

let personas: setup.Personas
let defaultAccount: ethers.Wallet

const provider = setup.provider()
const linkTokenFactory = new contract.LinkTokenFactory()
const aggregatorFactory = new MockAggregatorFactory()
const historicAggregatorFactory = new MockHistoricAggregatorFactory()
const aggregatorFacadeFactory = new AggregatorFacadeFactory()
const aggregatorProxyFactory = new AggregatorProxyFactory()

beforeAll(async () => {
  const users = await setup.users(provider)

  personas = users.personas
  defaultAccount = users.roles.defaultAccount
})

describe('AggregatorProxy', () => {
  const deposit = h.toWei('100')
  const response = h.numToBytes32(54321)
  const response2 = h.numToBytes32(67890)
  const decimals = 18

  let link: contract.Instance<contract.LinkTokenFactory>
  let aggregator: contract.Instance<MockAggregatorFactory>
  let aggregator2: contract.Instance<MockAggregatorFactory>
  let historicAggregator: contract.Instance<MockHistoricAggregatorFactory>
  let proxy: contract.CallableOverrideInstance<AggregatorProxyFactory>
  const deployment = setup.snapshot(provider, async () => {
    link = await linkTokenFactory.connect(defaultAccount).deploy()
    aggregator = await aggregatorFactory
      .connect(defaultAccount)
      .deploy(decimals, response)
    await link.transfer(aggregator.address, deposit)
    proxy = contract.callable(
      await aggregatorProxyFactory
        .connect(defaultAccount)
        .deploy(aggregator.address),
      [
        'latestAnswer',
        'getAnswer',
        'latestTimestamp',
        'getTimestamp',
        'latestRound',
        'latestRoundData',
        'getRoundData',
        'proposedGetRoundData',
        'proposedLatestRoundData',
      ],
    )
  })

  beforeEach(async () => {
    await deployment()
  })

  it('has a limited public interface', () => {
    matchers.publicAbi(aggregatorProxyFactory, [
      'aggregator',
      'confirmAggregator',
      'decimals',
      'getAnswer',
      'getRoundData',
      'getTimestamp',
      'latestAnswer',
      'latestRound',
      'latestRoundData',
      'latestTimestamp',
      'proposeAggregator',
      'proposedAggregator',
      'proposedGetRoundData',
      'proposedLatestRoundData',
      // Ownable methods:
      'acceptOwnership',
      'owner',
      'transferOwnership',
    ])
  })

  describe('#latestAnswer', () => {
    it('pulls the rate from the aggregator', async () => {
      matchers.bigNum(response, await proxy.latestAnswer())
      const latestRound = await proxy.latestRound()
      matchers.bigNum(response, await proxy.getAnswer(latestRound))
    })

    describe('after being updated to another contract', () => {
      beforeEach(async () => {
        aggregator2 = await aggregatorFactory
          .connect(defaultAccount)
          .deploy(decimals, response2)
        await link.transfer(aggregator2.address, deposit)
        matchers.bigNum(response2, await aggregator2.latestAnswer())

        await proxy.proposeAggregator(aggregator2.address)
        await proxy.confirmAggregator(aggregator2.address)
      })

      it('pulls the rate from the new aggregator', async () => {
        matchers.bigNum(response2, await proxy.latestAnswer())
        const latestRound = await proxy.latestRound()
        matchers.bigNum(response2, await proxy.getAnswer(latestRound))
      })
    })
  })

  describe('#latestTimestamp', () => {
    beforeEach(async () => {
      const height = await aggregator.latestTimestamp()
      assert.notEqual('0', height.toString())
    })

    it('pulls the timestamp from the aggregator', async () => {
      matchers.bigNum(
        await aggregator.latestTimestamp(),
        await proxy.latestTimestamp(),
      )
      const latestRound = await proxy.latestRound()
      matchers.bigNum(
        await aggregator.latestTimestamp(),
        await proxy.getTimestamp(latestRound),
      )
    })

    describe('after being updated to another contract', () => {
      beforeEach(async () => {
        aggregator2 = await aggregatorFactory
          .connect(defaultAccount)
          .deploy(decimals, response2)

        const height2 = await aggregator2.latestTimestamp()
        assert.notEqual('0', height2.toString())

        const height1 = await aggregator.latestTimestamp()
        assert.notEqual(
          height1.toString(),
          height2.toString(),
          'Height1 and Height2 should not be equal',
        )

        await proxy.proposeAggregator(aggregator2.address)
        await proxy.confirmAggregator(aggregator2.address)
      })

      it('pulls the timestamp from the new aggregator', async () => {
        matchers.bigNum(
          await aggregator2.latestTimestamp(),
          await proxy.latestTimestamp(),
        )
        const latestRound = await proxy.latestRound()
        matchers.bigNum(
          await aggregator2.latestTimestamp(),
          await proxy.getTimestamp(latestRound),
        )
      })
    })
  })

  describe('#getRoundData', () => {
    describe('when pointed at a Historic Aggregator', () => {
      beforeEach(async () => {
        historicAggregator = await historicAggregatorFactory
          .connect(defaultAccount)
          .deploy(response2)
        await proxy.proposeAggregator(historicAggregator.address)
        await proxy.confirmAggregator(historicAggregator.address)
        matchers.bigNum(
          ethers.utils.bigNumberify(response2),
          await proxy.latestAnswer(),
        )
      })

      it('reverts', async () => {
        const latestRoundId = await proxy.latestRound()
        await matchers.evmRevert(async () => {
          await proxy.getRoundData(latestRoundId)
        })
      })

      describe('when pointed at an Aggregator Facade', () => {
        beforeEach(async () => {
          const facade = await aggregatorFacadeFactory
            .connect(defaultAccount)
            .deploy(aggregator.address, 18)
          await proxy.proposeAggregator(facade.address)
          await proxy.confirmAggregator(facade.address)
        })

        it('works for a valid roundId', async () => {
          const roundId = await aggregator.latestRound()
          const round = await proxy.getRoundData(roundId)
          matchers.bigNum(roundId, round.roundId)
          matchers.bigNum(response, round.answer)
          const nowSeconds = new Date().valueOf() / 1000
          assert.isAbove(round.updatedAt.toNumber(), nowSeconds - 120)
          matchers.bigNum(round.updatedAt, round.startedAt)
          matchers.bigNum(roundId, round.answeredInRound)
        })
      })
    })

    describe('when pointed at a FluxAggregator', () => {
      beforeEach(async () => {
        aggregator2 = await aggregatorFactory
          .connect(defaultAccount)
          .deploy(decimals, response2)

        await proxy.proposeAggregator(aggregator2.address)
        await proxy.confirmAggregator(aggregator2.address)
      })

      it('works for a valid round ID', async () => {
        const roundId = await aggregator2.latestRound()
        const round = await proxy.getRoundData(roundId)
        matchers.bigNum(roundId, round.roundId)
        matchers.bigNum(response2, round.answer)
        const nowSeconds = new Date().valueOf() / 1000
        assert.isAbove(round.startedAt.toNumber(), nowSeconds - 120)
        assert.isBelow(round.startedAt.toNumber(), nowSeconds)
        matchers.bigNum(round.startedAt, round.updatedAt)
        matchers.bigNum(roundId, round.answeredInRound)
      })
    })
  })

  describe('#latestRoundData', () => {
    describe('when pointed at a Historic Aggregator', () => {
      beforeEach(async () => {
        historicAggregator = await historicAggregatorFactory
          .connect(defaultAccount)
          .deploy(response2)
        await proxy.proposeAggregator(historicAggregator.address)
        await proxy.confirmAggregator(historicAggregator.address)
        matchers.bigNum(
          ethers.utils.bigNumberify(response2),
          await proxy.latestAnswer(),
        )
      })

      it('reverts', async () => {
        await matchers.evmRevert(async () => {
          await proxy.latestRoundData()
        })
      })

      describe('when pointed at an Aggregator Facade', () => {
        beforeEach(async () => {
          const facade = await aggregatorFacadeFactory
            .connect(defaultAccount)
            .deploy(aggregator.address, 18)
          await proxy.proposeAggregator(facade.address)
          await proxy.confirmAggregator(facade.address)
        })

        it('does not revert', async () => {
          const roundId = await aggregator.latestRound()
          const round = await proxy.latestRoundData()
          matchers.bigNum(roundId, round.roundId)
          matchers.bigNum(response, round.answer)
          const nowSeconds = new Date().valueOf() / 1000
          assert.isAbove(round.updatedAt.toNumber(), nowSeconds - 120)
          matchers.bigNum(round.updatedAt, round.startedAt)
          matchers.bigNum(roundId, round.answeredInRound)
        })
      })
    })

    describe('when pointed at a FluxAggregator', () => {
      beforeEach(async () => {
        aggregator2 = await aggregatorFactory
          .connect(defaultAccount)
          .deploy(decimals, response2)

        await proxy.proposeAggregator(aggregator2.address)
        await proxy.confirmAggregator(aggregator2.address)
      })

      it('does not revert', async () => {
        const roundId = await aggregator2.latestRound()
        const round = await proxy.latestRoundData()
        matchers.bigNum(roundId, round.roundId)
        matchers.bigNum(response2, round.answer)
        const nowSeconds = new Date().valueOf() / 1000
        assert.isAbove(round.startedAt.toNumber(), nowSeconds - 120)
        assert.isBelow(round.startedAt.toNumber(), nowSeconds)
        matchers.bigNum(round.startedAt, round.updatedAt)
        matchers.bigNum(roundId, round.answeredInRound)
      })
    })
  })

  describe('#proposeAggregator', () => {
    beforeEach(async () => {
      await proxy.transferOwnership(personas.Carol.address)
      await proxy.connect(personas.Carol).acceptOwnership()

      aggregator2 = await aggregatorFactory
        .connect(defaultAccount)
        .deploy(decimals, 1)

      assert.equal(aggregator.address, await proxy.aggregator())
    })

    describe('when called by the owner', () => {
      it('sets the address of the proposed aggregator', async () => {
        await proxy
          .connect(personas.Carol)
          .proposeAggregator(aggregator2.address)

        assert.equal(aggregator2.address, await proxy.proposedAggregator())
      })
    })

    describe('when called by a non-owner', () => {
      it('does not update', async () => {
        await matchers.evmRevert(async () => {
          await proxy
            .connect(personas.Neil)
            .proposeAggregator(aggregator2.address)
        })

        assert.equal(aggregator.address, await proxy.aggregator())
      })
    })
  })

  describe('#confirmAggregator', () => {
    beforeEach(async () => {
      await proxy.transferOwnership(personas.Carol.address)
      await proxy.connect(personas.Carol).acceptOwnership()

      aggregator2 = await aggregatorFactory
        .connect(defaultAccount)
        .deploy(decimals, 1)

      assert.equal(aggregator.address, await proxy.aggregator())
    })

    describe('when called by the owner', () => {
      it('sets the address of the new aggregator', async () => {
        await proxy
          .connect(personas.Carol)
          .proposeAggregator(aggregator2.address)
        await proxy
          .connect(personas.Carol)
          .confirmAggregator(aggregator2.address)

        assert.equal(aggregator2.address, await proxy.aggregator())
      })
    })

    describe('when called by a non-owner', () => {
      beforeEach(async () => {
        await proxy
          .connect(personas.Carol)
          .proposeAggregator(aggregator2.address)
      })

      it('does not update', async () => {
        await matchers.evmRevert(async () => {
          await proxy
            .connect(personas.Neil)
            .confirmAggregator(aggregator2.address)
        })

        assert.equal(aggregator.address, await proxy.aggregator())
      })
    })
  })

  describe('#proposedGetRoundData', () => {
    beforeEach(async () => {
      aggregator2 = await aggregatorFactory
        .connect(defaultAccount)
        .deploy(decimals, response2)
    })

    describe('when an aggregator has been proposed', () => {
      beforeEach(async () => {
        await proxy
          .connect(defaultAccount)
          .proposeAggregator(aggregator2.address)
        assert.equal(await proxy.proposedAggregator(), aggregator2.address)
      })

      it('returns the data for the proposed aggregator', async () => {
        const roundId = await aggregator2.latestRound()
        const round = await proxy.proposedGetRoundData(roundId)
        matchers.bigNum(roundId, round.roundId)
        matchers.bigNum(response2, round.answer)
      })

      describe('after the aggregator has been confirmed', () => {
        beforeEach(async () => {
          await proxy
            .connect(defaultAccount)
            .confirmAggregator(aggregator2.address)
          assert.equal(await proxy.aggregator(), aggregator2.address)
        })

        it('reverts', async () => {
          const roundId = await aggregator2.latestRound()
          await matchers.evmRevert(async () => {
            await proxy.proposedGetRoundData(roundId)
          })
        })
      })
    })
  })

  describe('#proposedLatestRoundData', () => {
    beforeEach(async () => {
      aggregator2 = await aggregatorFactory
        .connect(defaultAccount)
        .deploy(decimals, response2)
    })

    describe('when an aggregator has been proposed', () => {
      beforeEach(async () => {
        await proxy
          .connect(defaultAccount)
          .proposeAggregator(aggregator2.address)
        assert.equal(await proxy.proposedAggregator(), aggregator2.address)
      })

      it('returns the data for the proposed aggregator', async () => {
        const roundId = await aggregator2.latestRound()
        const round = await proxy.proposedLatestRoundData()
        matchers.bigNum(roundId, round.roundId)
        matchers.bigNum(response2, round.answer)
      })

      describe('after the aggregator has been confirmed', () => {
        beforeEach(async () => {
          await proxy
            .connect(defaultAccount)
            .confirmAggregator(aggregator2.address)
          assert.equal(await proxy.aggregator(), aggregator2.address)
        })

        it('reverts', async () => {
          await matchers.evmRevert(async () => {
            await proxy.proposedLatestRoundData()
          })
        })
      })
    })
  })
})
