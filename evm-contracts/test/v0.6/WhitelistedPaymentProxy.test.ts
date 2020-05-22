import {
  contract,
  helpers as h,
  matchers,
  setup,
  interfaces,
} from '@chainlink/test-helpers'
import { assert } from 'chai'
import { ethers } from 'ethers'
import { MockPaymentFactory } from '../../ethers/v0.6/MockPaymentFactory'
import { MockAggregatorFactory } from '../../ethers/v0.6/MockAggregatorFactory'
import { WhitelistedPaymentProxyFactory } from '../../ethers/v0.6/WhitelistedPaymentProxyFactory'

let personas: setup.Personas
let defaultAccount: ethers.Wallet

const provider = setup.provider()
const linkTokenFactory = new contract.LinkTokenFactory()
const paymentFactory = new MockPaymentFactory()
const aggregatorFactory = new MockAggregatorFactory()
const whitelistedPaymentProxyFactory = new WhitelistedPaymentProxyFactory()

beforeAll(async () => {
  const users = await setup.users(provider)

  personas = users.personas
  defaultAccount = users.roles.defaultAccount
})

describe('WhitelistedPaymentProxy', () => {
  const deposit = h.toWei('100')
  const answer = h.numToBytes32(54321)
  const roundId = 17
  const decimals = 18
  const timestamp = 678
  const startedAt = 677

  let link: contract.Instance<contract.LinkTokenFactory>
  let aggregator: contract.Instance<MockAggregatorFactory>
  let paymentContract: contract.Instance<MockPaymentFactory>
  let proxy: contract.CallableOverrideInstance<
    WhitelistedPaymentProxyFactory,
    interfaces.AggregatorInterface
  >

  const deployment = setup.snapshot(provider, async () => {
    link = await linkTokenFactory.connect(defaultAccount).deploy()
    aggregator = await aggregatorFactory
      .connect(defaultAccount)
      .deploy(decimals, 0)
    await aggregator.updateRoundData(roundId, answer, timestamp, startedAt)
    await link.transfer(aggregator.address, deposit)
    paymentContract = await paymentFactory.connect(defaultAccount).deploy()
    proxy = contract.callable(
      await whitelistedPaymentProxyFactory
        .connect(defaultAccount)
        .deploy(aggregator.address, paymentContract.address),
      interfaces.AggregatorMethodList,
    )
  })

  beforeEach(async () => {
    await deployment()
  })

  it('has a limited public interface', () => {
    matchers.publicAbi(whitelistedPaymentProxyFactory, [
      'aggregator',
      'decimals',
      'getAnswer',
      'getRoundData',
      'getTimestamp',
      'latestAnswer',
      'latestRound',
      'latestRoundData',
      'latestTimestamp',
      'setAggregator',
      'setPaymentContract',
      'paymentContract',
      // Ownable methods:
      'acceptOwnership',
      'owner',
      'transferOwnership',
      // Whitelisted methods:
      'addToWhitelist',
      'disableWhitelist',
      'enableWhitelist',
      'removeFromWhitelist',
      'whitelistEnabled',
      'whitelisted',
    ])
  })

  describe('if the caller is not whitelisted', () => {
    it('latestAnswer reverts', async () => {
      matchers.evmRevert(async () => {
        await proxy.connect(personas.Carol).latestAnswer()
      }, 'Not whitelisted')
    })

    it('latestTimestamp reverts', async () => {
      matchers.evmRevert(async () => {
        await proxy.connect(personas.Carol).latestTimestamp()
      }, 'Not whitelisted')
    })

    it('getAnswer reverts', async () => {
      matchers.evmRevert(async () => {
        await proxy.connect(personas.Carol).getAnswer(1)
      }, 'Not whitelisted')
    })

    it('getTimestamp reverts', async () => {
      matchers.evmRevert(async () => {
        await proxy.connect(personas.Carol).getTimestamp(1)
      }, 'Not whitelisted')
    })

    it('latestRound reverts', async () => {
      matchers.evmRevert(async () => {
        await proxy.connect(personas.Carol).latestRound()
      }, 'Not whitelisted')
    })

    it('getRoundData reverts', async () => {
      matchers.evmRevert(async () => {
        await proxy.connect(personas.Carol).getRoundData(1)
      }, 'Not whitelisted')
    })
  })

  describe('if the caller is whitelisted by payment', () => {
    beforeEach(async () => {
      await paymentContract.addToWhitelist(defaultAccount.address)

      matchers.bigNum(
        ethers.utils.bigNumberify(answer),
        await aggregator.latestAnswer(),
      )
      const height = await aggregator.latestTimestamp()
      assert.notEqual('0', height.toString())
    })

    it('pulls the rate from the aggregator', async () => {
      matchers.bigNum(answer, await proxy.latestAnswer())
      const latestRound = await proxy.latestRound()
      matchers.bigNum(answer, await proxy.getAnswer(latestRound))
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

    it('getRoundData works', async () => {
      const latestRound = await proxy.latestRound()
      await proxy.latestRound()
      const round = await proxy.getRoundData(latestRound)
      await proxy.getRoundData(latestRound)
      matchers.bigNum(roundId, round.roundId)
      matchers.bigNum(answer, round.answer)
      matchers.bigNum(startedAt, round.startedAt)
      matchers.bigNum(timestamp, round.updatedAt)
    })
  })
})
