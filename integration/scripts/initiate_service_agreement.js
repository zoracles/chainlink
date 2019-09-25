/* eslint-disable @typescript-eslint/no-var-requires */

// truffle script

const chai = require('chai')
const path = require('path')
global.web3 = new (require('web3'))() // XXX: These are needed for helpers.js, for now.
global.assert = chai.assert

console.log('COORDINATOR_ADDRESS', process.env.COORDINATOR_ADDRESS)
const SRC_ROOT = '../../' // Root of chainlink source code, relative to cwd
const fullPath = s => path.resolve(SRC_ROOT, s)

console.log('path', process.cwd(), fullPath('evm/dist/src/helpers.js'))

// XXX: Hopefully there'll be a nicer way to access this, by the time this is
// ready.
const chainLink = require(fullPath('evm/dist/src/helpers.js'))

const fs = require('fs')
const request = require('request-promise').defaults({ jar: true })
const url = require('url')
const { abort, DEVNET_ADDRESS, scriptRunner } = require(fullPath(
  'integration/common.js',
))

const CoordinatorABI = JSON.parse(
  fs.readFileSync(fullPath('integration/out/Coordinator.abi'), 'utf8'),
)
console.log('########################################################################', CoordinatorABI)
const coordinatorAddress = process.env.COORDINATOR_ADDRESS

console.log('coordinator_address', coordinatorAddress)

const agreement = JSON.parse(
  fs.readFileSync(fullPath('integration/agreement.json'), 'utf8'),
)

const amount = web3.utils.toBN(web3.utils.toWei('1000'))

const stripQuotes = s => s.replace(/^"(.*)"$/, '$1')
const rawSignature = stripQuotes(process.env.ORACLE_SIGNATURE)

agreement.oracleSignatures = [chainLink.newSignature(rawSignature)]
agreement.requestDigest = web3.utils.keccak256(process.env.NORMALIZED_REQUEST)
agreement.sAID = chainLink.calculateSAID2(agreement)
console.log(
  'requestDigest',
  agreement.requestDigest,
  'sAID',
  chainLink.toHex(agreement.sAID),
)
const nreq = new Buffer(stripQuotes(process.env.NORMALIZED_REQUEST), 'utf8')
const myBuffer = []
for (let i = 0; i < nreq.length; i++) {
  myBuffer.push(nreq[i])
}
const endAtEpochMilliseconds = new Date(agreement.endAt).getTime()
const millisecondsPerSecond = 1000
agreement.endAt = endAtEpochMilliseconds / millisecondsPerSecond

const Coordinator = new web3.eth.Contract(CoordinatorABI, coordinatorAddress)
// XXX: Monkey-patching some truffle-attributes assumed by helpers
Coordinator.abi = CoordinatorABI
Coordinator.initiateServiceAgreement =
  Coordinator.methods.initiateServiceAgreement

const main = async () => {
  console.log(
    'initiateServiceAgreementArgs',
    chainLink.initiateServiceAgreementArgs2(Coordinator, agreement),
  )
  // console.log('@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@', await web3.eth.getCode(coordinatorAddress))
  console.log('!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!', await Coordinator.methods.dummyMethodXXX().call({gas: 6238833, from: "0x9ca9d2d5e04012c9ed24c0e513c9bfaa4a2dd77f"}))
  console.log(
    'initiateServiceAgreement',
    await (await chainLink.initiateServiceAgreement2(
      Coordinator,
      agreement,
    )).call(),
  )
  // console.log(
  //   'oracleRequest',
  //   await Coordinator.methods.oracleRequest(
  //     agreement.sAID,
  //     '0x0101010101010101010101010101010101010101', // Receiving contract address
  //     '0x12345678', // receiving method selector
  //     1, // nonce
  //     '', // data for initialization of request
  // ),
  // )
}

module.exports = scriptRunner(main)
