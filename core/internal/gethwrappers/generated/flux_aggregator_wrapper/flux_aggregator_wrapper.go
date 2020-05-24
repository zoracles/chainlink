// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package flux_aggregator_wrapper

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// FluxAggregatorABI is the input ABI used to generate the binding from.
const FluxAggregatorABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_link\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"_paymentAmount\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"_timeout\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_description\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"int256\",\"name\":\"current\",\"type\":\"int256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"AnswerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"AvailableFundsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"startedBy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"}],\"name\":\"NewRound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"OracleAdminUpdateRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"OracleAdminUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"whitelisted\",\"type\":\"bool\"}],\"name\":\"OraclePermissionsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"authorized\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"delay\",\"type\":\"uint32\"}],\"name\":\"RequesterPermissionsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"paymentAmount\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"minSubmissionCount\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"maxSubmissionCount\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"restartDelay\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"timeout\",\"type\":\"uint32\"}],\"name\":\"RoundDetailsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"int256\",\"name\":\"submission\",\"type\":\"int256\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"round\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"}],\"name\":\"SubmissionReceived\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"}],\"name\":\"acceptAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_oracles\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_admins\",\"type\":\"address[]\"},{\"internalType\":\"uint32\",\"name\":\"_minSubmissions\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_maxSubmissions\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_restartDelay\",\"type\":\"uint32\"}],\"name\":\"addOracles\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allocatedFunds\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"availableFunds\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"description\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"}],\"name\":\"getAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_roundId\",\"type\":\"uint256\"}],\"name\":\"getAnswer\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOracles\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_roundId\",\"type\":\"uint256\"}],\"name\":\"getRoundData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"answeredInRound\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_roundId\",\"type\":\"uint256\"}],\"name\":\"getTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestAnswer\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRoundData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"answeredInRound\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"}],\"name\":\"latestSubmission\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"linkToken\",\"outputs\":[{\"internalType\":\"contractLinkTokenInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxSubmissionCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minSubmissionCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"onTokenTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracleCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"}],\"name\":\"oracleRoundState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_eligibleToSubmit\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"_roundId\",\"type\":\"uint32\"},{\"internalType\":\"int256\",\"name\":\"_latestSubmission\",\"type\":\"int256\"},{\"internalType\":\"uint64\",\"name\":\"_startedAt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_timeout\",\"type\":\"uint64\"},{\"internalType\":\"uint128\",\"name\":\"_availableFunds\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"_oracleCount\",\"type\":\"uint32\"},{\"internalType\":\"uint128\",\"name\":\"_paymentAmount\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paymentAmount\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_oracles\",\"type\":\"address[]\"},{\"internalType\":\"uint32\",\"name\":\"_minSubmissions\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_maxSubmissions\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_restartDelay\",\"type\":\"uint32\"}],\"name\":\"removeOracles\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reportingRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestNewRound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"restartDelay\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_requester\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_authorized\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"_delay\",\"type\":\"uint32\"}],\"name\":\"setRequesterPermissions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_roundId\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"_submission\",\"type\":\"int256\"}],\"name\":\"submit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeout\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_newAdmin\",\"type\":\"address\"}],\"name\":\"transferAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateAvailableFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"_paymentAmount\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"_minSubmissions\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_maxSubmissions\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_restartDelay\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_timeout\",\"type\":\"uint32\"}],\"name\":\"updateFutureRounds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawPayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"}],\"name\":\"withdrawablePayment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// FluxAggregatorBin is the compiled bytecode used for deploying new contracts.
var FluxAggregatorBin = "0x60806040523480156200001157600080fd5b5060405162006b4838038062006b48833981810160405260a08110156200003757600080fd5b810190808051906020019092919080519060200190929190805190602001909291908051906020019092919080519060200190929190505050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555084600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555083600460006101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff160217905550826004601c6101000a81548163ffffffff021916908363ffffffff16021790555081600560006101000a81548160ff021916908360ff160217905550806006819055506200018d8363ffffffff1642620001da60201b620045261790919060201c565b600960008063ffffffff16815260200190815260200160002060010160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550505050505062000264565b60008282111562000253576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601e8152602001807f536166654d6174683a207375627472616374696f6e206f766572666c6f77000081525060200191505060405180910390fd5b600082840390508091505092915050565b6168d480620002746000396000f3fe608060405234801561001057600080fd5b50600436106102c85760003560e01c806379ba50971161017b578063c1075329116100d8578063e9ee6eeb1161008c578063f2fde38b11610071578063f2fde38b14610e28578063feaf968c14610e6c578063ffa1ad7414610ea6576102c8565b8063e9ee6eeb14610d1b578063ebf8571c14610d7f576102c8565b8063c9374500116100bd578063c937450014610c57578063d4cc54e414610c81578063e2e4031714610cc3576102c8565b8063c107532914610bc7578063c35905c614610c15576102c8565b8063ab175a4d1161012f578063b633620c11610114578063b633620c14610a28578063bb07bacd14610a6a578063bbf0b7e914610ac9576102c8565b8063ab175a4d146108d1578063b5ab58dc146109e6576102c8565b80638da5cb5b116101605780638da5cb5b146107da57806398e5b12a14610824578063a4c0ed361461082e576102c8565b806379ba5097146107b25780638205bf6a146107bc576102c8565b806350d25bcd1161022957806364efb22b116101dd5780636fb4bb4e116101c25780636fb4bb4e1461074c57806370dea79a1461076a5780637284e41614610794576102c8565b806364efb22b146106aa578063668a0f021461072e576102c8565b806358609e441161020e57806358609e4414610612578063613d8fcc1461063c578063628806ef14610666576102c8565b806350d25bcd146105aa57806357970e93146105c8576102c8565b806338aa4c721161028057806340884c521161026557806340884c52146104ff57806346fcff4c1461055e5780634f8fc3b5146105a0576102c8565b806338aa4c72146104115780633d3d771414610491576102c8565b806320ed0275116102b157806320ed027514610363578063313ce567146103c3578063357ebb02146103e7576102c8565b80630720da52146102cd578063202ee0ed1461032b575b600080fd5b6102f9600480360360208110156102e357600080fd5b8101908080359060200190929190505050610ec4565b604051808681526020018581526020018481526020018381526020018281526020019550505050505060405180910390f35b6103616004803603604081101561034157600080fd5b810190808035906020019092919080359060200190929190505050610ee8565b005b6103c16004803603606081101561037957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803515159060200190929190803563ffffffff169060200190929190505050610fd6565b005b6103cb6112af565b604051808260ff1660ff16815260200191505060405180910390f35b6103ef6112c2565b604051808263ffffffff1663ffffffff16815260200191505060405180910390f35b61048f600480360360a081101561042757600080fd5b8101908080356fffffffffffffffffffffffffffffffff169060200190929190803563ffffffff169060200190929190803563ffffffff169060200190929190803563ffffffff169060200190929190803563ffffffff1690602001909291905050506112d8565b005b6104fd600480360360608110156104a757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506117ea565b005b610507611c0a565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b8381101561054a57808201518184015260208101905061052f565b505050509050019250505060405180910390f35b610566611c98565b60405180826fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6105a8611cba565b005b6105b2611e81565b6040518082815260200191505060405180910390f35b6105d0611e90565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61061a611eb6565b604051808263ffffffff1663ffffffff16815260200191505060405180910390f35b610644611ecc565b604051808263ffffffff1663ffffffff16815260200191505060405180910390f35b6106a86004803603602081101561067c57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611ed9565b005b6106ec600480360360208110156106c057600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061213c565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6107366121a8565b6040518082815260200191505060405180910390f35b6107546121c8565b6040518082815260200191505060405180910390f35b6107726121e8565b604051808263ffffffff1663ffffffff16815260200191505060405180910390f35b61079c6121fe565b6040518082815260200191505060405180910390f35b6107ba612204565b005b6107c46123cc565b6040518082815260200191505060405180910390f35b6107e26123db565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61082c612400565b005b6108cf6004803603606081101561084457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035906020019064010000000081111561088b57600080fd5b82018360208201111561089d57600080fd5b803590602001918460018302840111640100000000831117156108bf57600080fd5b90919293919293905050506125c4565b005b610913600480360360208110156108e757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061264b565b60405180891515151581526020018863ffffffff1663ffffffff1681526020018781526020018667ffffffffffffffff1667ffffffffffffffff1681526020018567ffffffffffffffff1667ffffffffffffffff168152602001846fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff1681526020018363ffffffff1663ffffffff168152602001826fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff1681526020019850505050505050505060405180910390f35b610a12600480360360208110156109fc57600080fd5b8101908080359060200190929190505050612996565b6040518082815260200191505060405180910390f35b610a5460048036036020811015610a3e57600080fd5b81019080803590602001909291905050506129a8565b6040518082815260200191505060405180910390f35b610aac60048036036020811015610a8057600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506129ba565b604051808381526020018281526020019250505060405180910390f35b610bc5600480360360a0811015610adf57600080fd5b8101908080359060200190640100000000811115610afc57600080fd5b820183602082011115610b0e57600080fd5b80359060200191846020830284011164010000000083111715610b3057600080fd5b909192939192939080359060200190640100000000811115610b5157600080fd5b820183602082011115610b6357600080fd5b80359060200191846020830284011164010000000083111715610b8557600080fd5b9091929391929390803563ffffffff169060200190929190803563ffffffff169060200190929190803563ffffffff169060200190929190505050612a65565b005b610c1360048036036040811015610bdd57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050612cfa565b005b610c1d61300e565b60405180826fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b610c5f613030565b604051808263ffffffff1663ffffffff16815260200191505060405180910390f35b610c89613046565b60405180826fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b610d0560048036036020811015610cd957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050613068565b6040518082815260200191505060405180910390f35b610d7d60048036036040811015610d3157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506130e2565b005b610e2660048036036080811015610d9557600080fd5b8101908080359060200190640100000000811115610db257600080fd5b820183602082011115610dc457600080fd5b80359060200191846020830284011164010000000083111715610de657600080fd5b9091929391929390803563ffffffff169060200190929190803563ffffffff169060200190929190803563ffffffff169060200190929190505050613318565b005b610e6a60048036036020811015610e3e57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061346e565b005b610e746135ef565b604051808681526020018581526020018481526020018381526020018281526020019550505050505060405180910390f35b610eae613610565b6040518082815260200191505060405180910390f35b6000806000806000610ed586613615565b9450945094509450945091939590929450565b6060610ef43384613859565b905060008151148190610fa2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610f67578082015181840152602081019050610f4c565b50505050905090810190601f168015610f945780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50610fac83613bce565b610fb68284613ce8565b610fbf83613ea6565b610fc8836140ac565b610fd18361432c565b505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611098576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000081525060200191505060405180910390fd5b811515600a60008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900460ff16151514156110f8576112aa565b81156111be5781600a60008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160006101000a81548160ff02191690831515021790555080600a60008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160016101000a81548163ffffffff021916908363ffffffff160217905550611243565b600a60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600080820160006101000a81549060ff02191690556000820160016101000a81549063ffffffff02191690556000820160056101000a81549063ffffffff021916905550505b8273ffffffffffffffffffffffffffffffffffffffff167fc3df5a754e002718f2e10804b99e6605e7c701d95cec9552c7680ca2b6f2820a838360405180831515151581526020018263ffffffff1663ffffffff1681526020019250505060405180910390a25b505050565b600560009054906101000a900460ff1681565b600460189054906101000a900463ffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461139a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000081525060200191505060405180910390fd5b60006113a4611ecc565b90508463ffffffff168463ffffffff161015611428576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f6d6178206d75737420657175616c2f657863656564206d696e0000000000000081525060200191505060405180910390fd5b8363ffffffff168163ffffffff1610156114aa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260178152602001807f6d61782063616e6e6f742065786365656420746f74616c00000000000000000081525060200191505060405180910390fd5b60008163ffffffff1614806114ca57508263ffffffff168163ffffffff16115b61153c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f64656c61792063616e6e6f742065786365656420746f74616c0000000000000081525060200191505060405180910390fd5b611557866fffffffffffffffffffffffffffffffff1661443d565b600360109054906101000a90046fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff1610156115fc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601e8152602001807f696e73756666696369656e742066756e647320666f72207061796d656e74000081525060200191505060405180910390fd5b6000611606611ecc565b63ffffffff16111561168f5760008563ffffffff161161168e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601a8152602001807f6d696e206d7573742062652067726561746572207468616e203000000000000081525060200191505060405180910390fd5b5b85600460006101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff16021790555084600460146101000a81548163ffffffff021916908363ffffffff16021790555083600460106101000a81548163ffffffff021916908363ffffffff16021790555082600460186101000a81548163ffffffff021916908363ffffffff160217905550816004601c6101000a81548163ffffffff021916908363ffffffff1602179055508363ffffffff168563ffffffff16600460009054906101000a90046fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff167f56800c9d1ed723511246614d15e58cfcde15b6a33c245b5c961b689c1890fd8f8686604051808363ffffffff1663ffffffff1681526020018263ffffffff1663ffffffff1681526020019250505060405180910390a4505050505050565b3373ffffffffffffffffffffffffffffffffffffffff16600860008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146118ed576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f6f6e6c792063616c6c61626c652062792061646d696e0000000000000000000081525060200191505060405180910390fd5b60008190506000600860008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a90046fffffffffffffffffffffffffffffffff169050816fffffffffffffffffffffffffffffffff16816fffffffffffffffffffffffffffffffff1610156119ef576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f696e73756666696369656e7420776974686472617761626c652066756e64730081525060200191505060405180910390fd5b611a1482826fffffffffffffffffffffffffffffffff1661447990919063ffffffff16565b600860008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160006101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff160217905550611acf82600360009054906101000a90046fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff1661447990919063ffffffff16565b600360006101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff160217905550600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb85846fffffffffffffffffffffffffffffffff166040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b158015611bc257600080fd5b505af1158015611bd6573d6000803e3d6000fd5b505050506040513d6020811015611bec57600080fd5b8101908080519060200190929190505050611c0357fe5b5050505050565b6060600b805480602002602001604051908101604052809291908181526020018280548015611c8e57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311611c44575b5050505050905090565b600360109054906101000a90046fffffffffffffffffffffffffffffffff1681565b6000600360109054906101000a90046fffffffffffffffffffffffffffffffff1690506000611dfb600360009054906101000a90046fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff16600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b158015611db257600080fd5b505afa158015611dc6573d6000803e3d6000fd5b505050506040513d6020811015611ddc57600080fd5b810190808051906020019092919050505061452690919063ffffffff16565b905080600360106101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff16021790555080826fffffffffffffffffffffffffffffffff1614611e7d57807ffe25c73e3b9089fac37d55c4c7efcba6f04af04cebd2fc4d6d7dbb07e1e5234f60405160405180910390a25b5050565b6000611e8b6145af565b905090565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600460109054906101000a900463ffffffff1681565b6000600b80549050905090565b3373ffffffffffffffffffffffffffffffffffffffff16600860008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614611fdc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601e8152602001807f6f6e6c792063616c6c61626c652062792070656e64696e672061646d696e000081525060200191505060405180910390fd5b6000600860008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555033600860008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160026101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055503373ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f0c5055390645c15a4be9a21b3f8d019153dcb4a0c125685da6eb84048e2fe90460405160405180910390a350565b6000600860008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160029054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b6000600760049054906101000a900463ffffffff1663ffffffff16905090565b6000600760009054906101000a900463ffffffff1663ffffffff16905090565b6004601c9054906101000a900463ffffffff1681565b60065481565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146122c7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f4d7573742062652070726f706f736564206f776e65720000000000000000000081525060200191505060405180910390fd5b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055503373ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a350565b60006123d66145eb565b905090565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600a60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900460ff166124c2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260188152602001807f6e6f7420617574686f72697a656420726571756573746572000000000000000081525060200191505060405180910390fd5b6000600760009054906101000a900463ffffffff1690506000600960008363ffffffff1663ffffffff16815260200190815260200160002060010160089054906101000a900467ffffffffffffffff1667ffffffffffffffff16118061252d575061252c81614645565b5b61259f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f7072657620726f756e64206d75737420626520737570657273656461626c650081525060200191505060405180910390fd5b6125c16125bc60018363ffffffff1661471b90919063ffffffff16565b6147af565b50565b6000828290501461263d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f7472616e7366657220646f65736e2774206163636570742063616c6c6461746181525060200191505060405180910390fd5b612645611cba565b50505050565b6000806000806000806000803273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146126f8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f6f66662d636861696e2072656164696e67206f6e6c790000000000000000000081525060200191505060405180910390fd5b6000600760009054906101000a900463ffffffff1663ffffffff16600860008c73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160189054906101000a900463ffffffff1663ffffffff161480612790575061278e600760009054906101000a900463ffffffff1661496d565b155b90506127ad600760009054906101000a900463ffffffff166149b5565b80156127b65750805b1561281b576127e76001600760009054906101000a900463ffffffff1663ffffffff1661471b90919063ffffffff16565b9750600460009054906101000a90046fffffffffffffffffffffffffffffffff1691506128148a89614a11565b9850612880565b600760009054906101000a900463ffffffff169750600960008963ffffffff1663ffffffff168152602001908152602001600020600201600101600c9054906101000a90046fffffffffffffffffffffffffffffffff16915061287d8861496d565b98505b600061288c8b8a613859565b511461289757600098505b8888600860008d73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010154600960008c63ffffffff1663ffffffff16815260200190815260200160002060010160009054906101000a900467ffffffffffffffff16600960008d63ffffffff1663ffffffff16815260200190815260200160002060020160010160089054906101000a900463ffffffff16600360109054906101000a90046fffffffffffffffffffffffffffffffff16612970611ecc565b888363ffffffff1693509850985098509850985098509850985050919395975091939597565b60006129a182614aa6565b9050919050565b60006129b382614ad2565b9050919050565b600080600860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010154600860008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160189054906101000a900463ffffffff168063ffffffff16905091509150915091565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614612b27576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000081525060200191505060405180910390fd5b848490508787905014612ba2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f6e6565642073616d65206f7261636c6520616e642061646d696e20636f756e7481525060200191505060405180910390fd5b602a612bc788889050612bb3611ecc565b63ffffffff16614b1c90919063ffffffff16565b1115612c3b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260138152602001807f6d6178206f7261636c657320616c6c6f7765640000000000000000000000000081525060200191505060405180910390fd5b60008090505b87879050811015612cb357612ca6888883818110612c5b57fe5b9050602002013573ffffffffffffffffffffffffffffffffffffffff16878784818110612c8457fe5b9050602002013573ffffffffffffffffffffffffffffffffffffffff16614ba4565b8080600101915050612c41565b50612cf1600460009054906101000a90046fffffffffffffffffffffffffffffffff168484846004601c9054906101000a900463ffffffff166112d8565b50505050505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614612dbc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000081525060200191505060405180910390fd5b80612e38612df9600460009054906101000a90046fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff1661443d565b600360109054906101000a90046fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff1661452690919063ffffffff16565b1015612eac576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601a8152602001807f696e73756666696369656e7420726573657276652066756e647300000000000081525060200191505060405180910390fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb83836040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b158015612f5557600080fd5b505af1158015612f69573d6000803e3d6000fd5b505050506040513d6020811015612f7f57600080fd5b8101908080519060200190929190505050613002576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f746f6b656e207472616e73666572206661696c6564000000000000000000000081525060200191505060405180910390fd5b61300a611cba565b5050565b600460009054906101000a90046fffffffffffffffffffffffffffffffff1681565b600460149054906101000a900463ffffffff1681565b600360009054906101000a90046fffffffffffffffffffffffffffffffff1681565b6000600860008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a90046fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff169050919050565b3373ffffffffffffffffffffffffffffffffffffffff16600860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146131e5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f6f6e6c792063616c6c61626c652062792061646d696e0000000000000000000081525060200191505060405180910390fd5b80600860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff167fb79bf2e89c2d70dde91d2991fb1ea69b7e478061ad7c04ed5b02b96bc52b81043383604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060405180910390a25050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146133da576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000081525060200191505060405180910390fd5b60008090505b858590508110156134295761341c8686838181106133fa57fe5b9050602002013573ffffffffffffffffffffffffffffffffffffffff16615118565b80806001019150506133e0565b50613467600460009054906101000a90046fffffffffffffffffffffffffffffffff168484846004601c9054906101000a900463ffffffff166112d8565b5050505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614613530576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000081525060200191505060405180910390fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae127860405160405180910390a350565b60008060008060006135ff61546c565b945094509450945094509091929394565b600281565b6000806000806000613625616774565b600960008863ffffffff1663ffffffff1681526020019081526020016000206040518060a0016040529081600082015481526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160109054906101000a900463ffffffff1663ffffffff1663ffffffff168152602001600282016040518060a00160405290816000820180548060200260200160405190810160405280929190818152602001828054801561374257602002820191906000526020600020905b81548152602001906001019080831161372e575b505050505081526020016001820160009054906101000a900463ffffffff1663ffffffff1663ffffffff1681526020016001820160049054906101000a900463ffffffff1663ffffffff1663ffffffff1681526020016001820160089054906101000a900463ffffffff1663ffffffff1663ffffffff16815260200160018201600c9054906101000a90046fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff16815250508152505090508681600001518260200151836040015184606001518267ffffffffffffffff1692508167ffffffffffffffff1691508063ffffffff169050955095509550955095505091939590929450565b60606000600860008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160109054906101000a900463ffffffff1690506000600760009054906101000a900463ffffffff16905060008263ffffffff161415613917576040518060400160405280601281526020017f6e6f7420656e61626c6564206f7261636c65000000000000000000000000000081525092505050613bc8565b8363ffffffff168263ffffffff16111561396a576040518060400160405280601681526020017f6e6f742079657420656e61626c6564206f7261636c650000000000000000000081525092505050613bc8565b8363ffffffff16600860008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160149054906101000a900463ffffffff1663ffffffff161015613a0f576040518060400160405280601881526020017f6e6f206c6f6e67657220616c6c6f776564206f7261636c65000000000000000081525092505050613bc8565b8363ffffffff16600860008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160189054906101000a900463ffffffff1663ffffffff1610613ab3576040518060400160405280602081526020017f63616e6e6f74207265706f7274206f6e2070726576696f757320726f756e647381525092505050613bc8565b8063ffffffff168463ffffffff1614158015613af45750613ae460018263ffffffff1661471b90919063ffffffff16565b63ffffffff168463ffffffff1614155b8015613b075750613b0584826154a6565b155b15613b4b576040518060400160405280601781526020017f696e76616c696420726f756e6420746f207265706f727400000000000000000081525092505050613bc8565b60018463ffffffff1614158015613b815750613b7f613b7a60018663ffffffff1661552490919063ffffffff16565b6149b5565b155b15613bc5576040518060400160405280601f81526020017f70726576696f757320726f756e64206e6f7420737570657273656461626c650081525092505050613bc8565b50505b92915050565b613bd7816155b9565b613be057613ce5565b6000600860003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001601c9054906101000a900463ffffffff1663ffffffff169050600460189054906101000a900463ffffffff1663ffffffff1681018263ffffffff1611158015613c6e575060008114155b15613c795750613ce5565b613c82826155fc565b81600860003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001601c6101000a81548163ffffffff021916908363ffffffff160217905550505b50565b613cf18161496d565b613d63576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f726f756e64206e6f7420616363657074696e67207375626d697373696f6e730081525060200191505060405180910390fd5b600960008263ffffffff1663ffffffff16815260200190815260200160002060020160000182908060018154018082558091505060019003906000526020600020016000909190919091505580600860003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160186101000a81548163ffffffff021916908363ffffffff16021790555081600860003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101819055503373ffffffffffffffffffffffffffffffffffffffff168163ffffffff16837f92e98423f8adac6e64d0608e519fd1cefb861498385c6dee70d58fc926ddc68c60405160405180910390a45050565b600960008263ffffffff1663ffffffff16815260200190815260200160002060020160010160049054906101000a900463ffffffff1663ffffffff16600960008363ffffffff1663ffffffff168152602001908152602001600020600201600001805490501015613f16576140a9565b6000613f94600960008463ffffffff1663ffffffff168152602001908152602001600020600201600001805480602002602001604051908101604052809291908181526020018280548015613f8a57602002820191906000526020600020905b815481526020019060010190808311613f76575b505050505061589c565b905080600960008463ffffffff1663ffffffff1681526020019081526020016000206000018190555042600960008463ffffffff1663ffffffff16815260200190815260200160002060010160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555081600960008463ffffffff1663ffffffff16815260200190815260200160002060010160106101000a81548163ffffffff021916908363ffffffff16021790555081600760046101000a81548163ffffffff021916908363ffffffff1602179055508163ffffffff16817f0559884fd3a460db3073b7fc896cc77986f16e378210ded43186175bf646fc5f426040518082815260200191505060405180910390a3505b50565b6000600960008363ffffffff1663ffffffff168152602001908152602001600020600201600101600c9054906101000a90046fffffffffffffffffffffffffffffffff169050600061413782600360109054906101000a90046fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff1661447990919063ffffffff16565b905080600360106101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff1602179055506141b582600360009054906101000a90046fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff1661598b90919063ffffffff16565b600360006101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff16021790555061427082600860003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a90046fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff1661598b90919063ffffffff16565b600860003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160006101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff160217905550806fffffffffffffffffffffffffffffffff167ffe25c73e3b9089fac37d55c4c7efcba6f04af04cebd2fc4d6d7dbb07e1e5234f60405160405180910390a2505050565b600960008263ffffffff1663ffffffff16815260200190815260200160002060020160010160009054906101000a900463ffffffff1663ffffffff16600960008363ffffffff1663ffffffff16815260200190815260200160002060020160000180549050101561439c5761443a565b600960008263ffffffff1663ffffffff168152602001908152602001600020600201600080820160006143cf91906167c3565b6001820160006101000a81549063ffffffff02191690556001820160046101000a81549063ffffffff02191690556001820160086101000a81549063ffffffff021916905560018201600c6101000a8154906fffffffffffffffffffffffffffffffff021916905550505b50565b6000614472600261446461444f611ecc565b63ffffffff1685615a3790919063ffffffff16565b615a3790919063ffffffff16565b9050919050565b6000826fffffffffffffffffffffffffffffffff16826fffffffffffffffffffffffffffffffff161115614515576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601e8152602001807f536166654d6174683a207375627472616374696f6e206f766572666c6f77000081525060200191505060405180910390fd5b600082840390508091505092915050565b60008282111561459e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601e8152602001807f536166654d6174683a207375627472616374696f6e206f766572666c6f77000081525060200191505060405180910390fd5b600082840390508091505092915050565b600060096000600760049054906101000a900463ffffffff1663ffffffff1663ffffffff16815260200190815260200160002060000154905090565b600060096000600760049054906101000a900463ffffffff1663ffffffff1663ffffffff16815260200190815260200160002060010160089054906101000a900467ffffffffffffffff1667ffffffffffffffff16905090565b600080600960008463ffffffff1663ffffffff16815260200190815260200160002060010160009054906101000a900467ffffffffffffffff1690506000600960008563ffffffff1663ffffffff16815260200190815260200160002060020160010160089054906101000a900463ffffffff16905060008267ffffffffffffffff161180156146db575060008163ffffffff16115b80156147125750426147068263ffffffff168467ffffffffffffffff16615abd90919063ffffffff16565b67ffffffffffffffff16105b92505050919050565b60008082840190508363ffffffff168163ffffffff1610156147a5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081525060200191505060405180910390fd5b8091505092915050565b6147b8816155b9565b6147c15761496a565b6000600a60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160059054906101000a900463ffffffff1663ffffffff169050600a60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160019054906101000a900463ffffffff1663ffffffff1681018263ffffffff16118061488c5750600081145b6148fe576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260138152602001807f6d7573742064656c61792072657175657374730000000000000000000000000081525060200191505060405180910390fd5b614907826155fc565b81600a60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160056101000a81548163ffffffff021916908363ffffffff160217905550505b50565b600080600960008463ffffffff1663ffffffff16815260200190815260200160002060020160010160009054906101000a900463ffffffff1663ffffffff1614159050919050565b600080600960008463ffffffff1663ffffffff16815260200190815260200160002060010160089054906101000a900467ffffffffffffffff1667ffffffffffffffff161180614a0a5750614a0982614645565b5b9050919050565b600080600860008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001601c9054906101000a900463ffffffff1663ffffffff169050600460189054906101000a900463ffffffff1663ffffffff1681018363ffffffff161180614a9d5750600081145b91505092915050565b6000600960008363ffffffff1663ffffffff168152602001908152602001600020600001549050919050565b6000600960008363ffffffff1663ffffffff16815260200190815260200160002060010160089054906101000a900467ffffffffffffffff1667ffffffffffffffff169050919050565b600080828401905083811015614b9a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081525060200191505060405180910390fd5b8091505092915050565b614bad82615b59565b15614c20576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f6f7261636c6520616c726561647920656e61626c65640000000000000000000081525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415614cc3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f63616e6e6f74207365742061646d696e20746f2030000000000000000000000081525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff16600860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161480614ded57508073ffffffffffffffffffffffffffffffffffffffff16600860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16145b614e5f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601c8152602001807f6f776e65722063616e6e6f74206f76657277726974652061646d696e0000000081525060200191505060405180910390fd5b614e6882615bc3565b600860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160106101000a81548163ffffffff021916908363ffffffff16021790555063ffffffff600860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160146101000a81548163ffffffff021916908363ffffffff160217905550600b80549050600860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160006101000a81548161ffff021916908361ffff160217905550600b829080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160026101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600115158273ffffffffffffffffffffffffffffffffffffffff167f18dd09695e4fbdae8d1a5edb11221eb04564269c29a089b9753a6535c54ba92e60405160405180910390a38073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167f0c5055390645c15a4be9a21b3f8d019153dcb4a0c125685da6eb84048e2fe90460405160405180910390a35050565b61512181615b59565b615193576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f6f7261636c65206e6f7420656e61626c6564000000000000000000000000000081525060200191505060405180910390fd5b6151bf6001600760009054906101000a900463ffffffff1663ffffffff1661471b90919063ffffffff16565b600860008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160146101000a81548163ffffffff021916908363ffffffff1602179055506000600b6152446001615230611ecc565b63ffffffff1661552490919063ffffffff16565b63ffffffff168154811061525457fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690506000600860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160009054906101000a900461ffff16905080600860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160006101000a81548161ffff021916908361ffff160217905550600860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160006101000a81549061ffff021916905581600b8261ffff168154811061539857fe5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600b8054806153eb57fe5b6001900381819060005260206000200160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690559055600015158373ffffffffffffffffffffffffffffffffffffffff167f18dd09695e4fbdae8d1a5edb11221eb04564269c29a089b9753a6535c54ba92e60405160405180910390a3505050565b6000806000806000615495600760049054906101000a900463ffffffff1663ffffffff16613615565b945094509450945094509091929394565b60008163ffffffff166154c960018563ffffffff1661471b90919063ffffffff16565b63ffffffff1614801561551c57506000600960008463ffffffff1663ffffffff16815260200190815260200160002060010160089054906101000a900467ffffffffffffffff1667ffffffffffffffff16145b905092915050565b60008263ffffffff168263ffffffff1611156155a8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601e8152602001807f536166654d6174683a207375627472616374696f6e206f766572666c6f77000081525060200191505060405180910390fd5b600082840390508091505092915050565b60006155e76001600760009054906101000a900463ffffffff1663ffffffff1661471b90919063ffffffff16565b63ffffffff168263ffffffff16149050919050565b61561e61561960018363ffffffff1661552490919063ffffffff16565b615c80565b80600760006101000a81548163ffffffff021916908363ffffffff160217905550600460109054906101000a900463ffffffff16600960008363ffffffff1663ffffffff16815260200190815260200160002060020160010160006101000a81548163ffffffff021916908363ffffffff160217905550600460149054906101000a900463ffffffff16600960008363ffffffff1663ffffffff16815260200190815260200160002060020160010160046101000a81548163ffffffff021916908363ffffffff160217905550600460009054906101000a90046fffffffffffffffffffffffffffffffff16600960008363ffffffff1663ffffffff168152602001908152602001600020600201600101600c6101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff1602179055506004601c9054906101000a900463ffffffff16600960008363ffffffff1663ffffffff16815260200190815260200160002060020160010160086101000a81548163ffffffff021916908363ffffffff16021790555042600960008363ffffffff1663ffffffff16815260200190815260200160002060010160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055503373ffffffffffffffffffffffffffffffffffffffff168163ffffffff167f0109fc6f55cf40689f02fbaad7af7fe7bbac8a3d2186600afc7d3e10cac60271600960008563ffffffff1663ffffffff16815260200190815260200160002060010160009054906101000a900467ffffffffffffffff16604051808267ffffffffffffffff16815260200191505060405180910390a350565b60008151600010615915576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f6c697374206d757374206e6f7420626520656d7074790000000000000000000081525060200191505060405180910390fd5b60008251905060006002828161592757fe5b04905060006002838161593657fe5b06141561597157600080615954866000600187036001870387615e57565b80925081935050506159668282615f44565b945050505050615986565b6159818460006001850384615fe1565b925050505b919050565b6000808284019050836fffffffffffffffffffffffffffffffff16816fffffffffffffffffffffffffffffffff161015615a2d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081525060200191505060405180910390fd5b8091505092915050565b600080831415615a4a5760009050615ab7565b6000828402905082848281615a5b57fe5b0414615ab2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602181526020018061687e6021913960400191505060405180910390fd5b809150505b92915050565b60008082840190508367ffffffffffffffff168167ffffffffffffffff161015615b4f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081525060200191505060405180910390fd5b8091505092915050565b600063ffffffff8016600860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160149054906101000a900463ffffffff1663ffffffff16149050919050565b600080600760009054906101000a900463ffffffff16905060008163ffffffff1614158015615c4f5750600860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160149054906101000a900463ffffffff1663ffffffff168163ffffffff16145b15615c5d5780915050615c7b565b615c7760018263ffffffff1661471b90919063ffffffff16565b9150505b919050565b615c8981614645565b615c9257615e54565b6000615cae60018363ffffffff1661552490919063ffffffff16565b9050600960008263ffffffff1663ffffffff16815260200190815260200160002060000154600960008463ffffffff1663ffffffff16815260200190815260200160002060000181905550600960008263ffffffff1663ffffffff16815260200190815260200160002060010160109054906101000a900463ffffffff16600960008463ffffffff1663ffffffff16815260200190815260200160002060010160106101000a81548163ffffffff021916908363ffffffff16021790555042600960008463ffffffff1663ffffffff16815260200190815260200160002060010160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550600960008363ffffffff1663ffffffff16815260200190815260200160002060020160008082016000615de891906167c3565b6001820160006101000a81549063ffffffff02191690556001820160046101000a81549063ffffffff02191690556001820160086101000a81549063ffffffff021916905560018201600c6101000a8154906fffffffffffffffffffffffffffffffff02191690555050505b50565b600080828410615e6657600080fd5b838611158015615e765750848411155b615e7f57600080fd5b828611158015615e8f5750848311155b615e9857600080fd5b5b600115615f395760078686031015615ec157615eb8878787878761607b565b91509150615f3a565b6000615ece8888886165ef565b9050808411615edf57809550615f33565b84811015615ef257600181019650615f32565b808511158015615f0157508381105b615f0757fe5b615f1388888388615fe1565b9250615f2488600183018887615fe1565b915082829250925050615f3a565b5b50615e99565b5b9550959350505050565b60008083128015615f555750600082135b80615f6c5750600083138015615f6b5750600082125b5b15615f8c576002615f7d84846166e6565b81615f8457fe5b059050615fdb565b60006002808481615f9957fe5b0760028681615fa457fe5b070181615fad57fe5b059050615fd7615fd160028681615fc057fe5b0560028681615fcb57fe5b056166e6565b826166e6565b9150505b92915050565b600081841115615ff057600080fd5b82821115615ffd57600080fd5b5b8284101561605c5760078484031015616031576000616020868686868761607b565b809250819350505081915050616073565b600061603e8686866165ef565b905080831161604f57809350616056565b6001810194505b50615ffe565b84848151811061606857fe5b602002602001015190505b949350505050565b60008060008660018701039050600088600089018151811061609957fe5b602002602001015190506000826001106160d3577f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6160eb565b8960018a01815181106160e257fe5b60200260200101515b905060008360021061611d577f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff616135565b8a60028b018151811061612c57fe5b60200260200101515b9050600084600310616167577f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff61617f565b8b60038c018151811061617657fe5b60200260200101515b90506000856004106161b1577f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6161c9565b8c60048d01815181106161c057fe5b60200260200101515b90506000866005106161fb577f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff616213565b8d60058e018151811061620a57fe5b60200260200101515b9050600087600610616245577f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff61625d565b8e60068f018151811061625457fe5b60200260200101515b90508587131561627257858780975081985050505b8385131561628557838580955081965050505b8183131561629857818380935081945050505b848713156162ab57848780965081985050505b838613156162be57838680955081975050505b808313156162d157808380925081945050505b848613156162e457848680965081975050505b808213156162f757808280925081935050505b8287131561630a57828780945081985050505b8186131561631d57818680935081975050505b8085131561633057808580925081965050505b8286131561634357828680945081975050505b8084131561635657808480925081955050505b8285131561636957828580945081965050505b8184131561637c57818480935081955050505b8284131561638f57828480945081955050505b60008e8d03905060008114156163a757879a50616481565b60018114156163b857869a50616480565b60028114156163c957859a5061647f565b60038114156163da57849a5061647e565b60048114156163eb57839a5061647d565b60058114156163fc57829a5061647c565b600681141561640d57819a5061647b565b6040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260108152602001807f6b31206f7574206f6620626f756e64730000000000000000000000000000000081525060200191505060405180910390fd5b5b5b5b5b5b5b60008f8d0390508c8e14156164a5578b8c9b509b50505050505050505050506165e5565b60008114156164c3578b899b509b50505050505050505050506165e5565b60018114156164e1578b889b509b50505050505050505050506165e5565b60028114156164ff578b879b509b50505050505050505050506165e5565b600381141561651d578b869b509b50505050505050505050506165e5565b600481141561653b578b859b509b50505050505050505050506165e5565b6005811415616559578b849b509b50505050505050505050506165e5565b6006811415616577578b839b509b50505050505050505050506165e5565b6040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260108152602001807f6b32206f7574206f6620626f756e64730000000000000000000000000000000081525060200191505060405180910390fd5b9550959350505050565b600080846002848601816165ff57fe5b048151811061660a57fe5b602002602001015190506001840393506001830192505b6001156166dd575b6001840193508085858151811061663c57fe5b602002602001015112616629575b6001830392508085848151811061665d57fe5b60200260200101511361664a57828410156166cf5784838151811061667e57fe5b602002602001015185858151811061669257fe5b60200260200101518686815181106166a657fe5b602002602001018786815181106166b957fe5b60200260200101828152508281525050506166d8565b829150506166df565b616621565b505b9392505050565b6000808284019050600083121580156166ff5750838112155b80616715575060008312801561671457508381125b5b61676a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602181526020018061685d6021913960400191505060405180910390fd5b8091505092915050565b6040518060a0016040528060008152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600063ffffffff1681526020016167bd6167e4565b81525090565b50805460008255906000526020600020908101906167e19190616837565b50565b6040518060a0016040528060608152602001600063ffffffff168152602001600063ffffffff168152602001600063ffffffff16815260200160006fffffffffffffffffffffffffffffffff1681525090565b61685991905b8082111561685557600081600090555060010161683d565b5090565b9056fe5369676e6564536166654d6174683a206164646974696f6e206f766572666c6f77536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f77a2646970667358220000000000000000000000000000000000000000000000000000000000000000000064736f6c63430000000033"

// DeployFluxAggregator deploys a new Ethereum contract, binding an instance of FluxAggregator to it.
func DeployFluxAggregator(auth *bind.TransactOpts, backend bind.ContractBackend, _link common.Address, _paymentAmount *big.Int, _timeout uint32, _decimals uint8, _description [32]byte) (common.Address, *types.Transaction, *FluxAggregator, error) {
	parsed, err := abi.JSON(strings.NewReader(FluxAggregatorABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FluxAggregatorBin), backend, _link, _paymentAmount, _timeout, _decimals, _description)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FluxAggregator{FluxAggregatorCaller: FluxAggregatorCaller{contract: contract}, FluxAggregatorTransactor: FluxAggregatorTransactor{contract: contract}, FluxAggregatorFilterer: FluxAggregatorFilterer{contract: contract}}, nil
}

// FluxAggregator is an auto generated Go binding around an Ethereum contract.
type FluxAggregator struct {
	FluxAggregatorCaller     // Read-only binding to the contract
	FluxAggregatorTransactor // Write-only binding to the contract
	FluxAggregatorFilterer   // Log filterer for contract events
}

// FluxAggregatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type FluxAggregatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FluxAggregatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FluxAggregatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FluxAggregatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FluxAggregatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FluxAggregatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FluxAggregatorSession struct {
	Contract     *FluxAggregator   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FluxAggregatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FluxAggregatorCallerSession struct {
	Contract *FluxAggregatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// FluxAggregatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FluxAggregatorTransactorSession struct {
	Contract     *FluxAggregatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// FluxAggregatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type FluxAggregatorRaw struct {
	Contract *FluxAggregator // Generic contract binding to access the raw methods on
}

// FluxAggregatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FluxAggregatorCallerRaw struct {
	Contract *FluxAggregatorCaller // Generic read-only contract binding to access the raw methods on
}

// FluxAggregatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FluxAggregatorTransactorRaw struct {
	Contract *FluxAggregatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFluxAggregator creates a new instance of FluxAggregator, bound to a specific deployed contract.
func NewFluxAggregator(address common.Address, backend bind.ContractBackend) (*FluxAggregator, error) {
	contract, err := bindFluxAggregator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FluxAggregator{FluxAggregatorCaller: FluxAggregatorCaller{contract: contract}, FluxAggregatorTransactor: FluxAggregatorTransactor{contract: contract}, FluxAggregatorFilterer: FluxAggregatorFilterer{contract: contract}}, nil
}

// NewFluxAggregatorCaller creates a new read-only instance of FluxAggregator, bound to a specific deployed contract.
func NewFluxAggregatorCaller(address common.Address, caller bind.ContractCaller) (*FluxAggregatorCaller, error) {
	contract, err := bindFluxAggregator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FluxAggregatorCaller{contract: contract}, nil
}

// NewFluxAggregatorTransactor creates a new write-only instance of FluxAggregator, bound to a specific deployed contract.
func NewFluxAggregatorTransactor(address common.Address, transactor bind.ContractTransactor) (*FluxAggregatorTransactor, error) {
	contract, err := bindFluxAggregator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FluxAggregatorTransactor{contract: contract}, nil
}

// NewFluxAggregatorFilterer creates a new log filterer instance of FluxAggregator, bound to a specific deployed contract.
func NewFluxAggregatorFilterer(address common.Address, filterer bind.ContractFilterer) (*FluxAggregatorFilterer, error) {
	contract, err := bindFluxAggregator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FluxAggregatorFilterer{contract: contract}, nil
}

// bindFluxAggregator binds a generic wrapper to an already deployed contract.
func bindFluxAggregator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FluxAggregatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FluxAggregator *FluxAggregatorRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FluxAggregator.Contract.FluxAggregatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FluxAggregator *FluxAggregatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FluxAggregator.Contract.FluxAggregatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FluxAggregator *FluxAggregatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FluxAggregator.Contract.FluxAggregatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FluxAggregator *FluxAggregatorCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FluxAggregator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FluxAggregator *FluxAggregatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FluxAggregator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FluxAggregator *FluxAggregatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FluxAggregator.Contract.contract.Transact(opts, method, params...)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() constant returns(uint256)
func (_FluxAggregator *FluxAggregatorCaller) VERSION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FluxAggregator.contract.Call(opts, out, "VERSION")
	return *ret0, err
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() constant returns(uint256)
func (_FluxAggregator *FluxAggregatorSession) VERSION() (*big.Int, error) {
	return _FluxAggregator.Contract.VERSION(&_FluxAggregator.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() constant returns(uint256)
func (_FluxAggregator *FluxAggregatorCallerSession) VERSION() (*big.Int, error) {
	return _FluxAggregator.Contract.VERSION(&_FluxAggregator.CallOpts)
}

// AllocatedFunds is a free data retrieval call binding the contract method 0xd4cc54e4.
//
// Solidity: function allocatedFunds() constant returns(uint128)
func (_FluxAggregator *FluxAggregatorCaller) AllocatedFunds(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FluxAggregator.contract.Call(opts, out, "allocatedFunds")
	return *ret0, err
}

// AllocatedFunds is a free data retrieval call binding the contract method 0xd4cc54e4.
//
// Solidity: function allocatedFunds() constant returns(uint128)
func (_FluxAggregator *FluxAggregatorSession) AllocatedFunds() (*big.Int, error) {
	return _FluxAggregator.Contract.AllocatedFunds(&_FluxAggregator.CallOpts)
}

// AllocatedFunds is a free data retrieval call binding the contract method 0xd4cc54e4.
//
// Solidity: function allocatedFunds() constant returns(uint128)
func (_FluxAggregator *FluxAggregatorCallerSession) AllocatedFunds() (*big.Int, error) {
	return _FluxAggregator.Contract.AllocatedFunds(&_FluxAggregator.CallOpts)
}

// AvailableFunds is a free data retrieval call binding the contract method 0x46fcff4c.
//
// Solidity: function availableFunds() constant returns(uint128)
func (_FluxAggregator *FluxAggregatorCaller) AvailableFunds(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FluxAggregator.contract.Call(opts, out, "availableFunds")
	return *ret0, err
}

// AvailableFunds is a free data retrieval call binding the contract method 0x46fcff4c.
//
// Solidity: function availableFunds() constant returns(uint128)
func (_FluxAggregator *FluxAggregatorSession) AvailableFunds() (*big.Int, error) {
	return _FluxAggregator.Contract.AvailableFunds(&_FluxAggregator.CallOpts)
}

// AvailableFunds is a free data retrieval call binding the contract method 0x46fcff4c.
//
// Solidity: function availableFunds() constant returns(uint128)
func (_FluxAggregator *FluxAggregatorCallerSession) AvailableFunds() (*big.Int, error) {
	return _FluxAggregator.Contract.AvailableFunds(&_FluxAggregator.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_FluxAggregator *FluxAggregatorCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _FluxAggregator.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_FluxAggregator *FluxAggregatorSession) Decimals() (uint8, error) {
	return _FluxAggregator.Contract.Decimals(&_FluxAggregator.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_FluxAggregator *FluxAggregatorCallerSession) Decimals() (uint8, error) {
	return _FluxAggregator.Contract.Decimals(&_FluxAggregator.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() constant returns(bytes32)
func (_FluxAggregator *FluxAggregatorCaller) Description(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _FluxAggregator.contract.Call(opts, out, "description")
	return *ret0, err
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() constant returns(bytes32)
func (_FluxAggregator *FluxAggregatorSession) Description() ([32]byte, error) {
	return _FluxAggregator.Contract.Description(&_FluxAggregator.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() constant returns(bytes32)
func (_FluxAggregator *FluxAggregatorCallerSession) Description() ([32]byte, error) {
	return _FluxAggregator.Contract.Description(&_FluxAggregator.CallOpts)
}

// GetAdmin is a free data retrieval call binding the contract method 0x64efb22b.
//
// Solidity: function getAdmin(address _oracle) constant returns(address)
func (_FluxAggregator *FluxAggregatorCaller) GetAdmin(opts *bind.CallOpts, _oracle common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FluxAggregator.contract.Call(opts, out, "getAdmin", _oracle)
	return *ret0, err
}

// GetAdmin is a free data retrieval call binding the contract method 0x64efb22b.
//
// Solidity: function getAdmin(address _oracle) constant returns(address)
func (_FluxAggregator *FluxAggregatorSession) GetAdmin(_oracle common.Address) (common.Address, error) {
	return _FluxAggregator.Contract.GetAdmin(&_FluxAggregator.CallOpts, _oracle)
}

// GetAdmin is a free data retrieval call binding the contract method 0x64efb22b.
//
// Solidity: function getAdmin(address _oracle) constant returns(address)
func (_FluxAggregator *FluxAggregatorCallerSession) GetAdmin(_oracle common.Address) (common.Address, error) {
	return _FluxAggregator.Contract.GetAdmin(&_FluxAggregator.CallOpts, _oracle)
}

// GetOracles is a free data retrieval call binding the contract method 0x40884c52.
//
// Solidity: function getOracles() constant returns(address[])
func (_FluxAggregator *FluxAggregatorCaller) GetOracles(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _FluxAggregator.contract.Call(opts, out, "getOracles")
	return *ret0, err
}

// GetOracles is a free data retrieval call binding the contract method 0x40884c52.
//
// Solidity: function getOracles() constant returns(address[])
func (_FluxAggregator *FluxAggregatorSession) GetOracles() ([]common.Address, error) {
	return _FluxAggregator.Contract.GetOracles(&_FluxAggregator.CallOpts)
}

// GetOracles is a free data retrieval call binding the contract method 0x40884c52.
//
// Solidity: function getOracles() constant returns(address[])
func (_FluxAggregator *FluxAggregatorCallerSession) GetOracles() ([]common.Address, error) {
	return _FluxAggregator.Contract.GetOracles(&_FluxAggregator.CallOpts)
}

// LatestSubmission is a free data retrieval call binding the contract method 0xbb07bacd.
//
// Solidity: function latestSubmission(address _oracle) constant returns(int256, uint256)
func (_FluxAggregator *FluxAggregatorCaller) LatestSubmission(opts *bind.CallOpts, _oracle common.Address) (*big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _FluxAggregator.contract.Call(opts, out, "latestSubmission", _oracle)
	return *ret0, *ret1, err
}

// LatestSubmission is a free data retrieval call binding the contract method 0xbb07bacd.
//
// Solidity: function latestSubmission(address _oracle) constant returns(int256, uint256)
func (_FluxAggregator *FluxAggregatorSession) LatestSubmission(_oracle common.Address) (*big.Int, *big.Int, error) {
	return _FluxAggregator.Contract.LatestSubmission(&_FluxAggregator.CallOpts, _oracle)
}

// LatestSubmission is a free data retrieval call binding the contract method 0xbb07bacd.
//
// Solidity: function latestSubmission(address _oracle) constant returns(int256, uint256)
func (_FluxAggregator *FluxAggregatorCallerSession) LatestSubmission(_oracle common.Address) (*big.Int, *big.Int, error) {
	return _FluxAggregator.Contract.LatestSubmission(&_FluxAggregator.CallOpts, _oracle)
}

// LinkToken is a free data retrieval call binding the contract method 0x57970e93.
//
// Solidity: function linkToken() constant returns(address)
func (_FluxAggregator *FluxAggregatorCaller) LinkToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FluxAggregator.contract.Call(opts, out, "linkToken")
	return *ret0, err
}

// LinkToken is a free data retrieval call binding the contract method 0x57970e93.
//
// Solidity: function linkToken() constant returns(address)
func (_FluxAggregator *FluxAggregatorSession) LinkToken() (common.Address, error) {
	return _FluxAggregator.Contract.LinkToken(&_FluxAggregator.CallOpts)
}

// LinkToken is a free data retrieval call binding the contract method 0x57970e93.
//
// Solidity: function linkToken() constant returns(address)
func (_FluxAggregator *FluxAggregatorCallerSession) LinkToken() (common.Address, error) {
	return _FluxAggregator.Contract.LinkToken(&_FluxAggregator.CallOpts)
}

// MaxSubmissionCount is a free data retrieval call binding the contract method 0x58609e44.
//
// Solidity: function maxSubmissionCount() constant returns(uint32)
func (_FluxAggregator *FluxAggregatorCaller) MaxSubmissionCount(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _FluxAggregator.contract.Call(opts, out, "maxSubmissionCount")
	return *ret0, err
}

// MaxSubmissionCount is a free data retrieval call binding the contract method 0x58609e44.
//
// Solidity: function maxSubmissionCount() constant returns(uint32)
func (_FluxAggregator *FluxAggregatorSession) MaxSubmissionCount() (uint32, error) {
	return _FluxAggregator.Contract.MaxSubmissionCount(&_FluxAggregator.CallOpts)
}

// MaxSubmissionCount is a free data retrieval call binding the contract method 0x58609e44.
//
// Solidity: function maxSubmissionCount() constant returns(uint32)
func (_FluxAggregator *FluxAggregatorCallerSession) MaxSubmissionCount() (uint32, error) {
	return _FluxAggregator.Contract.MaxSubmissionCount(&_FluxAggregator.CallOpts)
}

// MinSubmissionCount is a free data retrieval call binding the contract method 0xc9374500.
//
// Solidity: function minSubmissionCount() constant returns(uint32)
func (_FluxAggregator *FluxAggregatorCaller) MinSubmissionCount(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _FluxAggregator.contract.Call(opts, out, "minSubmissionCount")
	return *ret0, err
}

// MinSubmissionCount is a free data retrieval call binding the contract method 0xc9374500.
//
// Solidity: function minSubmissionCount() constant returns(uint32)
func (_FluxAggregator *FluxAggregatorSession) MinSubmissionCount() (uint32, error) {
	return _FluxAggregator.Contract.MinSubmissionCount(&_FluxAggregator.CallOpts)
}

// MinSubmissionCount is a free data retrieval call binding the contract method 0xc9374500.
//
// Solidity: function minSubmissionCount() constant returns(uint32)
func (_FluxAggregator *FluxAggregatorCallerSession) MinSubmissionCount() (uint32, error) {
	return _FluxAggregator.Contract.MinSubmissionCount(&_FluxAggregator.CallOpts)
}

// OracleCount is a free data retrieval call binding the contract method 0x613d8fcc.
//
// Solidity: function oracleCount() constant returns(uint32)
func (_FluxAggregator *FluxAggregatorCaller) OracleCount(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _FluxAggregator.contract.Call(opts, out, "oracleCount")
	return *ret0, err
}

// OracleCount is a free data retrieval call binding the contract method 0x613d8fcc.
//
// Solidity: function oracleCount() constant returns(uint32)
func (_FluxAggregator *FluxAggregatorSession) OracleCount() (uint32, error) {
	return _FluxAggregator.Contract.OracleCount(&_FluxAggregator.CallOpts)
}

// OracleCount is a free data retrieval call binding the contract method 0x613d8fcc.
//
// Solidity: function oracleCount() constant returns(uint32)
func (_FluxAggregator *FluxAggregatorCallerSession) OracleCount() (uint32, error) {
	return _FluxAggregator.Contract.OracleCount(&_FluxAggregator.CallOpts)
}

// OracleRoundState is a free data retrieval call binding the contract method 0xab175a4d.
//
// Solidity: function oracleRoundState(address _oracle) constant returns(bool _eligibleToSubmit, uint32 _roundId, int256 _latestSubmission, uint64 _startedAt, uint64 _timeout, uint128 _availableFunds, uint32 _oracleCount, uint128 _paymentAmount)
func (_FluxAggregator *FluxAggregatorCaller) OracleRoundState(opts *bind.CallOpts, _oracle common.Address) (struct {
	EligibleToSubmit bool
	RoundId          uint32
	LatestSubmission *big.Int
	StartedAt        uint64
	Timeout          uint64
	AvailableFunds   *big.Int
	OracleCount      uint32
	PaymentAmount    *big.Int
}, error) {
	ret := new(struct {
		EligibleToSubmit bool
		RoundId          uint32
		LatestSubmission *big.Int
		StartedAt        uint64
		Timeout          uint64
		AvailableFunds   *big.Int
		OracleCount      uint32
		PaymentAmount    *big.Int
	})
	out := ret
	err := _FluxAggregator.contract.Call(opts, out, "oracleRoundState", _oracle)
	return *ret, err
}

// OracleRoundState is a free data retrieval call binding the contract method 0xab175a4d.
//
// Solidity: function oracleRoundState(address _oracle) constant returns(bool _eligibleToSubmit, uint32 _roundId, int256 _latestSubmission, uint64 _startedAt, uint64 _timeout, uint128 _availableFunds, uint32 _oracleCount, uint128 _paymentAmount)
func (_FluxAggregator *FluxAggregatorSession) OracleRoundState(_oracle common.Address) (struct {
	EligibleToSubmit bool
	RoundId          uint32
	LatestSubmission *big.Int
	StartedAt        uint64
	Timeout          uint64
	AvailableFunds   *big.Int
	OracleCount      uint32
	PaymentAmount    *big.Int
}, error) {
	return _FluxAggregator.Contract.OracleRoundState(&_FluxAggregator.CallOpts, _oracle)
}

// OracleRoundState is a free data retrieval call binding the contract method 0xab175a4d.
//
// Solidity: function oracleRoundState(address _oracle) constant returns(bool _eligibleToSubmit, uint32 _roundId, int256 _latestSubmission, uint64 _startedAt, uint64 _timeout, uint128 _availableFunds, uint32 _oracleCount, uint128 _paymentAmount)
func (_FluxAggregator *FluxAggregatorCallerSession) OracleRoundState(_oracle common.Address) (struct {
	EligibleToSubmit bool
	RoundId          uint32
	LatestSubmission *big.Int
	StartedAt        uint64
	Timeout          uint64
	AvailableFunds   *big.Int
	OracleCount      uint32
	PaymentAmount    *big.Int
}, error) {
	return _FluxAggregator.Contract.OracleRoundState(&_FluxAggregator.CallOpts, _oracle)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_FluxAggregator *FluxAggregatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FluxAggregator.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_FluxAggregator *FluxAggregatorSession) Owner() (common.Address, error) {
	return _FluxAggregator.Contract.Owner(&_FluxAggregator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_FluxAggregator *FluxAggregatorCallerSession) Owner() (common.Address, error) {
	return _FluxAggregator.Contract.Owner(&_FluxAggregator.CallOpts)
}

// PaymentAmount is a free data retrieval call binding the contract method 0xc35905c6.
//
// Solidity: function paymentAmount() constant returns(uint128)
func (_FluxAggregator *FluxAggregatorCaller) PaymentAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FluxAggregator.contract.Call(opts, out, "paymentAmount")
	return *ret0, err
}

// PaymentAmount is a free data retrieval call binding the contract method 0xc35905c6.
//
// Solidity: function paymentAmount() constant returns(uint128)
func (_FluxAggregator *FluxAggregatorSession) PaymentAmount() (*big.Int, error) {
	return _FluxAggregator.Contract.PaymentAmount(&_FluxAggregator.CallOpts)
}

// PaymentAmount is a free data retrieval call binding the contract method 0xc35905c6.
//
// Solidity: function paymentAmount() constant returns(uint128)
func (_FluxAggregator *FluxAggregatorCallerSession) PaymentAmount() (*big.Int, error) {
	return _FluxAggregator.Contract.PaymentAmount(&_FluxAggregator.CallOpts)
}

// ReportingRound is a free data retrieval call binding the contract method 0x6fb4bb4e.
//
// Solidity: function reportingRound() constant returns(uint256)
func (_FluxAggregator *FluxAggregatorCaller) ReportingRound(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FluxAggregator.contract.Call(opts, out, "reportingRound")
	return *ret0, err
}

// ReportingRound is a free data retrieval call binding the contract method 0x6fb4bb4e.
//
// Solidity: function reportingRound() constant returns(uint256)
func (_FluxAggregator *FluxAggregatorSession) ReportingRound() (*big.Int, error) {
	return _FluxAggregator.Contract.ReportingRound(&_FluxAggregator.CallOpts)
}

// ReportingRound is a free data retrieval call binding the contract method 0x6fb4bb4e.
//
// Solidity: function reportingRound() constant returns(uint256)
func (_FluxAggregator *FluxAggregatorCallerSession) ReportingRound() (*big.Int, error) {
	return _FluxAggregator.Contract.ReportingRound(&_FluxAggregator.CallOpts)
}

// RestartDelay is a free data retrieval call binding the contract method 0x357ebb02.
//
// Solidity: function restartDelay() constant returns(uint32)
func (_FluxAggregator *FluxAggregatorCaller) RestartDelay(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _FluxAggregator.contract.Call(opts, out, "restartDelay")
	return *ret0, err
}

// RestartDelay is a free data retrieval call binding the contract method 0x357ebb02.
//
// Solidity: function restartDelay() constant returns(uint32)
func (_FluxAggregator *FluxAggregatorSession) RestartDelay() (uint32, error) {
	return _FluxAggregator.Contract.RestartDelay(&_FluxAggregator.CallOpts)
}

// RestartDelay is a free data retrieval call binding the contract method 0x357ebb02.
//
// Solidity: function restartDelay() constant returns(uint32)
func (_FluxAggregator *FluxAggregatorCallerSession) RestartDelay() (uint32, error) {
	return _FluxAggregator.Contract.RestartDelay(&_FluxAggregator.CallOpts)
}

// Timeout is a free data retrieval call binding the contract method 0x70dea79a.
//
// Solidity: function timeout() constant returns(uint32)
func (_FluxAggregator *FluxAggregatorCaller) Timeout(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _FluxAggregator.contract.Call(opts, out, "timeout")
	return *ret0, err
}

// Timeout is a free data retrieval call binding the contract method 0x70dea79a.
//
// Solidity: function timeout() constant returns(uint32)
func (_FluxAggregator *FluxAggregatorSession) Timeout() (uint32, error) {
	return _FluxAggregator.Contract.Timeout(&_FluxAggregator.CallOpts)
}

// Timeout is a free data retrieval call binding the contract method 0x70dea79a.
//
// Solidity: function timeout() constant returns(uint32)
func (_FluxAggregator *FluxAggregatorCallerSession) Timeout() (uint32, error) {
	return _FluxAggregator.Contract.Timeout(&_FluxAggregator.CallOpts)
}

// WithdrawablePayment is a free data retrieval call binding the contract method 0xe2e40317.
//
// Solidity: function withdrawablePayment(address _oracle) constant returns(uint256)
func (_FluxAggregator *FluxAggregatorCaller) WithdrawablePayment(opts *bind.CallOpts, _oracle common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FluxAggregator.contract.Call(opts, out, "withdrawablePayment", _oracle)
	return *ret0, err
}

// WithdrawablePayment is a free data retrieval call binding the contract method 0xe2e40317.
//
// Solidity: function withdrawablePayment(address _oracle) constant returns(uint256)
func (_FluxAggregator *FluxAggregatorSession) WithdrawablePayment(_oracle common.Address) (*big.Int, error) {
	return _FluxAggregator.Contract.WithdrawablePayment(&_FluxAggregator.CallOpts, _oracle)
}

// WithdrawablePayment is a free data retrieval call binding the contract method 0xe2e40317.
//
// Solidity: function withdrawablePayment(address _oracle) constant returns(uint256)
func (_FluxAggregator *FluxAggregatorCallerSession) WithdrawablePayment(_oracle common.Address) (*big.Int, error) {
	return _FluxAggregator.Contract.WithdrawablePayment(&_FluxAggregator.CallOpts, _oracle)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0x628806ef.
//
// Solidity: function acceptAdmin(address _oracle) returns()
func (_FluxAggregator *FluxAggregatorTransactor) AcceptAdmin(opts *bind.TransactOpts, _oracle common.Address) (*types.Transaction, error) {
	return _FluxAggregator.contract.Transact(opts, "acceptAdmin", _oracle)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0x628806ef.
//
// Solidity: function acceptAdmin(address _oracle) returns()
func (_FluxAggregator *FluxAggregatorSession) AcceptAdmin(_oracle common.Address) (*types.Transaction, error) {
	return _FluxAggregator.Contract.AcceptAdmin(&_FluxAggregator.TransactOpts, _oracle)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0x628806ef.
//
// Solidity: function acceptAdmin(address _oracle) returns()
func (_FluxAggregator *FluxAggregatorTransactorSession) AcceptAdmin(_oracle common.Address) (*types.Transaction, error) {
	return _FluxAggregator.Contract.AcceptAdmin(&_FluxAggregator.TransactOpts, _oracle)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_FluxAggregator *FluxAggregatorTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FluxAggregator.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_FluxAggregator *FluxAggregatorSession) AcceptOwnership() (*types.Transaction, error) {
	return _FluxAggregator.Contract.AcceptOwnership(&_FluxAggregator.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_FluxAggregator *FluxAggregatorTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _FluxAggregator.Contract.AcceptOwnership(&_FluxAggregator.TransactOpts)
}

// AddOracles is a paid mutator transaction binding the contract method 0xbbf0b7e9.
//
// Solidity: function addOracles(address[] _oracles, address[] _admins, uint32 _minSubmissions, uint32 _maxSubmissions, uint32 _restartDelay) returns()
func (_FluxAggregator *FluxAggregatorTransactor) AddOracles(opts *bind.TransactOpts, _oracles []common.Address, _admins []common.Address, _minSubmissions uint32, _maxSubmissions uint32, _restartDelay uint32) (*types.Transaction, error) {
	return _FluxAggregator.contract.Transact(opts, "addOracles", _oracles, _admins, _minSubmissions, _maxSubmissions, _restartDelay)
}

// AddOracles is a paid mutator transaction binding the contract method 0xbbf0b7e9.
//
// Solidity: function addOracles(address[] _oracles, address[] _admins, uint32 _minSubmissions, uint32 _maxSubmissions, uint32 _restartDelay) returns()
func (_FluxAggregator *FluxAggregatorSession) AddOracles(_oracles []common.Address, _admins []common.Address, _minSubmissions uint32, _maxSubmissions uint32, _restartDelay uint32) (*types.Transaction, error) {
	return _FluxAggregator.Contract.AddOracles(&_FluxAggregator.TransactOpts, _oracles, _admins, _minSubmissions, _maxSubmissions, _restartDelay)
}

// AddOracles is a paid mutator transaction binding the contract method 0xbbf0b7e9.
//
// Solidity: function addOracles(address[] _oracles, address[] _admins, uint32 _minSubmissions, uint32 _maxSubmissions, uint32 _restartDelay) returns()
func (_FluxAggregator *FluxAggregatorTransactorSession) AddOracles(_oracles []common.Address, _admins []common.Address, _minSubmissions uint32, _maxSubmissions uint32, _restartDelay uint32) (*types.Transaction, error) {
	return _FluxAggregator.Contract.AddOracles(&_FluxAggregator.TransactOpts, _oracles, _admins, _minSubmissions, _maxSubmissions, _restartDelay)
}

// GetAnswer is a paid mutator transaction binding the contract method 0xb5ab58dc.
//
// Solidity: function getAnswer(uint256 _roundId) returns(int256)
func (_FluxAggregator *FluxAggregatorTransactor) GetAnswer(opts *bind.TransactOpts, _roundId *big.Int) (*types.Transaction, error) {
	return _FluxAggregator.contract.Transact(opts, "getAnswer", _roundId)
}

// GetAnswer is a paid mutator transaction binding the contract method 0xb5ab58dc.
//
// Solidity: function getAnswer(uint256 _roundId) returns(int256)
func (_FluxAggregator *FluxAggregatorSession) GetAnswer(_roundId *big.Int) (*types.Transaction, error) {
	return _FluxAggregator.Contract.GetAnswer(&_FluxAggregator.TransactOpts, _roundId)
}

// GetAnswer is a paid mutator transaction binding the contract method 0xb5ab58dc.
//
// Solidity: function getAnswer(uint256 _roundId) returns(int256)
func (_FluxAggregator *FluxAggregatorTransactorSession) GetAnswer(_roundId *big.Int) (*types.Transaction, error) {
	return _FluxAggregator.Contract.GetAnswer(&_FluxAggregator.TransactOpts, _roundId)
}

// GetRoundData is a paid mutator transaction binding the contract method 0x0720da52.
//
// Solidity: function getRoundData(uint256 _roundId) returns(uint256 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint256 answeredInRound)
func (_FluxAggregator *FluxAggregatorTransactor) GetRoundData(opts *bind.TransactOpts, _roundId *big.Int) (*types.Transaction, error) {
	return _FluxAggregator.contract.Transact(opts, "getRoundData", _roundId)
}

// GetRoundData is a paid mutator transaction binding the contract method 0x0720da52.
//
// Solidity: function getRoundData(uint256 _roundId) returns(uint256 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint256 answeredInRound)
func (_FluxAggregator *FluxAggregatorSession) GetRoundData(_roundId *big.Int) (*types.Transaction, error) {
	return _FluxAggregator.Contract.GetRoundData(&_FluxAggregator.TransactOpts, _roundId)
}

// GetRoundData is a paid mutator transaction binding the contract method 0x0720da52.
//
// Solidity: function getRoundData(uint256 _roundId) returns(uint256 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint256 answeredInRound)
func (_FluxAggregator *FluxAggregatorTransactorSession) GetRoundData(_roundId *big.Int) (*types.Transaction, error) {
	return _FluxAggregator.Contract.GetRoundData(&_FluxAggregator.TransactOpts, _roundId)
}

// GetTimestamp is a paid mutator transaction binding the contract method 0xb633620c.
//
// Solidity: function getTimestamp(uint256 _roundId) returns(uint256)
func (_FluxAggregator *FluxAggregatorTransactor) GetTimestamp(opts *bind.TransactOpts, _roundId *big.Int) (*types.Transaction, error) {
	return _FluxAggregator.contract.Transact(opts, "getTimestamp", _roundId)
}

// GetTimestamp is a paid mutator transaction binding the contract method 0xb633620c.
//
// Solidity: function getTimestamp(uint256 _roundId) returns(uint256)
func (_FluxAggregator *FluxAggregatorSession) GetTimestamp(_roundId *big.Int) (*types.Transaction, error) {
	return _FluxAggregator.Contract.GetTimestamp(&_FluxAggregator.TransactOpts, _roundId)
}

// GetTimestamp is a paid mutator transaction binding the contract method 0xb633620c.
//
// Solidity: function getTimestamp(uint256 _roundId) returns(uint256)
func (_FluxAggregator *FluxAggregatorTransactorSession) GetTimestamp(_roundId *big.Int) (*types.Transaction, error) {
	return _FluxAggregator.Contract.GetTimestamp(&_FluxAggregator.TransactOpts, _roundId)
}

// LatestAnswer is a paid mutator transaction binding the contract method 0x50d25bcd.
//
// Solidity: function latestAnswer() returns(int256)
func (_FluxAggregator *FluxAggregatorTransactor) LatestAnswer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FluxAggregator.contract.Transact(opts, "latestAnswer")
}

// LatestAnswer is a paid mutator transaction binding the contract method 0x50d25bcd.
//
// Solidity: function latestAnswer() returns(int256)
func (_FluxAggregator *FluxAggregatorSession) LatestAnswer() (*types.Transaction, error) {
	return _FluxAggregator.Contract.LatestAnswer(&_FluxAggregator.TransactOpts)
}

// LatestAnswer is a paid mutator transaction binding the contract method 0x50d25bcd.
//
// Solidity: function latestAnswer() returns(int256)
func (_FluxAggregator *FluxAggregatorTransactorSession) LatestAnswer() (*types.Transaction, error) {
	return _FluxAggregator.Contract.LatestAnswer(&_FluxAggregator.TransactOpts)
}

// LatestRound is a paid mutator transaction binding the contract method 0x668a0f02.
//
// Solidity: function latestRound() returns(uint256)
func (_FluxAggregator *FluxAggregatorTransactor) LatestRound(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FluxAggregator.contract.Transact(opts, "latestRound")
}

// LatestRound is a paid mutator transaction binding the contract method 0x668a0f02.
//
// Solidity: function latestRound() returns(uint256)
func (_FluxAggregator *FluxAggregatorSession) LatestRound() (*types.Transaction, error) {
	return _FluxAggregator.Contract.LatestRound(&_FluxAggregator.TransactOpts)
}

// LatestRound is a paid mutator transaction binding the contract method 0x668a0f02.
//
// Solidity: function latestRound() returns(uint256)
func (_FluxAggregator *FluxAggregatorTransactorSession) LatestRound() (*types.Transaction, error) {
	return _FluxAggregator.Contract.LatestRound(&_FluxAggregator.TransactOpts)
}

// LatestRoundData is a paid mutator transaction binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() returns(uint256 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint256 answeredInRound)
func (_FluxAggregator *FluxAggregatorTransactor) LatestRoundData(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FluxAggregator.contract.Transact(opts, "latestRoundData")
}

// LatestRoundData is a paid mutator transaction binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() returns(uint256 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint256 answeredInRound)
func (_FluxAggregator *FluxAggregatorSession) LatestRoundData() (*types.Transaction, error) {
	return _FluxAggregator.Contract.LatestRoundData(&_FluxAggregator.TransactOpts)
}

// LatestRoundData is a paid mutator transaction binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() returns(uint256 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint256 answeredInRound)
func (_FluxAggregator *FluxAggregatorTransactorSession) LatestRoundData() (*types.Transaction, error) {
	return _FluxAggregator.Contract.LatestRoundData(&_FluxAggregator.TransactOpts)
}

// LatestTimestamp is a paid mutator transaction binding the contract method 0x8205bf6a.
//
// Solidity: function latestTimestamp() returns(uint256)
func (_FluxAggregator *FluxAggregatorTransactor) LatestTimestamp(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FluxAggregator.contract.Transact(opts, "latestTimestamp")
}

// LatestTimestamp is a paid mutator transaction binding the contract method 0x8205bf6a.
//
// Solidity: function latestTimestamp() returns(uint256)
func (_FluxAggregator *FluxAggregatorSession) LatestTimestamp() (*types.Transaction, error) {
	return _FluxAggregator.Contract.LatestTimestamp(&_FluxAggregator.TransactOpts)
}

// LatestTimestamp is a paid mutator transaction binding the contract method 0x8205bf6a.
//
// Solidity: function latestTimestamp() returns(uint256)
func (_FluxAggregator *FluxAggregatorTransactorSession) LatestTimestamp() (*types.Transaction, error) {
	return _FluxAggregator.Contract.LatestTimestamp(&_FluxAggregator.TransactOpts)
}

// OnTokenTransfer is a paid mutator transaction binding the contract method 0xa4c0ed36.
//
// Solidity: function onTokenTransfer(address , uint256 , bytes _data) returns()
func (_FluxAggregator *FluxAggregatorTransactor) OnTokenTransfer(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int, _data []byte) (*types.Transaction, error) {
	return _FluxAggregator.contract.Transact(opts, "onTokenTransfer", arg0, arg1, _data)
}

// OnTokenTransfer is a paid mutator transaction binding the contract method 0xa4c0ed36.
//
// Solidity: function onTokenTransfer(address , uint256 , bytes _data) returns()
func (_FluxAggregator *FluxAggregatorSession) OnTokenTransfer(arg0 common.Address, arg1 *big.Int, _data []byte) (*types.Transaction, error) {
	return _FluxAggregator.Contract.OnTokenTransfer(&_FluxAggregator.TransactOpts, arg0, arg1, _data)
}

// OnTokenTransfer is a paid mutator transaction binding the contract method 0xa4c0ed36.
//
// Solidity: function onTokenTransfer(address , uint256 , bytes _data) returns()
func (_FluxAggregator *FluxAggregatorTransactorSession) OnTokenTransfer(arg0 common.Address, arg1 *big.Int, _data []byte) (*types.Transaction, error) {
	return _FluxAggregator.Contract.OnTokenTransfer(&_FluxAggregator.TransactOpts, arg0, arg1, _data)
}

// RemoveOracles is a paid mutator transaction binding the contract method 0xebf8571c.
//
// Solidity: function removeOracles(address[] _oracles, uint32 _minSubmissions, uint32 _maxSubmissions, uint32 _restartDelay) returns()
func (_FluxAggregator *FluxAggregatorTransactor) RemoveOracles(opts *bind.TransactOpts, _oracles []common.Address, _minSubmissions uint32, _maxSubmissions uint32, _restartDelay uint32) (*types.Transaction, error) {
	return _FluxAggregator.contract.Transact(opts, "removeOracles", _oracles, _minSubmissions, _maxSubmissions, _restartDelay)
}

// RemoveOracles is a paid mutator transaction binding the contract method 0xebf8571c.
//
// Solidity: function removeOracles(address[] _oracles, uint32 _minSubmissions, uint32 _maxSubmissions, uint32 _restartDelay) returns()
func (_FluxAggregator *FluxAggregatorSession) RemoveOracles(_oracles []common.Address, _minSubmissions uint32, _maxSubmissions uint32, _restartDelay uint32) (*types.Transaction, error) {
	return _FluxAggregator.Contract.RemoveOracles(&_FluxAggregator.TransactOpts, _oracles, _minSubmissions, _maxSubmissions, _restartDelay)
}

// RemoveOracles is a paid mutator transaction binding the contract method 0xebf8571c.
//
// Solidity: function removeOracles(address[] _oracles, uint32 _minSubmissions, uint32 _maxSubmissions, uint32 _restartDelay) returns()
func (_FluxAggregator *FluxAggregatorTransactorSession) RemoveOracles(_oracles []common.Address, _minSubmissions uint32, _maxSubmissions uint32, _restartDelay uint32) (*types.Transaction, error) {
	return _FluxAggregator.Contract.RemoveOracles(&_FluxAggregator.TransactOpts, _oracles, _minSubmissions, _maxSubmissions, _restartDelay)
}

// RequestNewRound is a paid mutator transaction binding the contract method 0x98e5b12a.
//
// Solidity: function requestNewRound() returns()
func (_FluxAggregator *FluxAggregatorTransactor) RequestNewRound(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FluxAggregator.contract.Transact(opts, "requestNewRound")
}

// RequestNewRound is a paid mutator transaction binding the contract method 0x98e5b12a.
//
// Solidity: function requestNewRound() returns()
func (_FluxAggregator *FluxAggregatorSession) RequestNewRound() (*types.Transaction, error) {
	return _FluxAggregator.Contract.RequestNewRound(&_FluxAggregator.TransactOpts)
}

// RequestNewRound is a paid mutator transaction binding the contract method 0x98e5b12a.
//
// Solidity: function requestNewRound() returns()
func (_FluxAggregator *FluxAggregatorTransactorSession) RequestNewRound() (*types.Transaction, error) {
	return _FluxAggregator.Contract.RequestNewRound(&_FluxAggregator.TransactOpts)
}

// SetRequesterPermissions is a paid mutator transaction binding the contract method 0x20ed0275.
//
// Solidity: function setRequesterPermissions(address _requester, bool _authorized, uint32 _delay) returns()
func (_FluxAggregator *FluxAggregatorTransactor) SetRequesterPermissions(opts *bind.TransactOpts, _requester common.Address, _authorized bool, _delay uint32) (*types.Transaction, error) {
	return _FluxAggregator.contract.Transact(opts, "setRequesterPermissions", _requester, _authorized, _delay)
}

// SetRequesterPermissions is a paid mutator transaction binding the contract method 0x20ed0275.
//
// Solidity: function setRequesterPermissions(address _requester, bool _authorized, uint32 _delay) returns()
func (_FluxAggregator *FluxAggregatorSession) SetRequesterPermissions(_requester common.Address, _authorized bool, _delay uint32) (*types.Transaction, error) {
	return _FluxAggregator.Contract.SetRequesterPermissions(&_FluxAggregator.TransactOpts, _requester, _authorized, _delay)
}

// SetRequesterPermissions is a paid mutator transaction binding the contract method 0x20ed0275.
//
// Solidity: function setRequesterPermissions(address _requester, bool _authorized, uint32 _delay) returns()
func (_FluxAggregator *FluxAggregatorTransactorSession) SetRequesterPermissions(_requester common.Address, _authorized bool, _delay uint32) (*types.Transaction, error) {
	return _FluxAggregator.Contract.SetRequesterPermissions(&_FluxAggregator.TransactOpts, _requester, _authorized, _delay)
}

// Submit is a paid mutator transaction binding the contract method 0x202ee0ed.
//
// Solidity: function submit(uint256 _roundId, int256 _submission) returns()
func (_FluxAggregator *FluxAggregatorTransactor) Submit(opts *bind.TransactOpts, _roundId *big.Int, _submission *big.Int) (*types.Transaction, error) {
	return _FluxAggregator.contract.Transact(opts, "submit", _roundId, _submission)
}

// Submit is a paid mutator transaction binding the contract method 0x202ee0ed.
//
// Solidity: function submit(uint256 _roundId, int256 _submission) returns()
func (_FluxAggregator *FluxAggregatorSession) Submit(_roundId *big.Int, _submission *big.Int) (*types.Transaction, error) {
	return _FluxAggregator.Contract.Submit(&_FluxAggregator.TransactOpts, _roundId, _submission)
}

// Submit is a paid mutator transaction binding the contract method 0x202ee0ed.
//
// Solidity: function submit(uint256 _roundId, int256 _submission) returns()
func (_FluxAggregator *FluxAggregatorTransactorSession) Submit(_roundId *big.Int, _submission *big.Int) (*types.Transaction, error) {
	return _FluxAggregator.Contract.Submit(&_FluxAggregator.TransactOpts, _roundId, _submission)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0xe9ee6eeb.
//
// Solidity: function transferAdmin(address _oracle, address _newAdmin) returns()
func (_FluxAggregator *FluxAggregatorTransactor) TransferAdmin(opts *bind.TransactOpts, _oracle common.Address, _newAdmin common.Address) (*types.Transaction, error) {
	return _FluxAggregator.contract.Transact(opts, "transferAdmin", _oracle, _newAdmin)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0xe9ee6eeb.
//
// Solidity: function transferAdmin(address _oracle, address _newAdmin) returns()
func (_FluxAggregator *FluxAggregatorSession) TransferAdmin(_oracle common.Address, _newAdmin common.Address) (*types.Transaction, error) {
	return _FluxAggregator.Contract.TransferAdmin(&_FluxAggregator.TransactOpts, _oracle, _newAdmin)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0xe9ee6eeb.
//
// Solidity: function transferAdmin(address _oracle, address _newAdmin) returns()
func (_FluxAggregator *FluxAggregatorTransactorSession) TransferAdmin(_oracle common.Address, _newAdmin common.Address) (*types.Transaction, error) {
	return _FluxAggregator.Contract.TransferAdmin(&_FluxAggregator.TransactOpts, _oracle, _newAdmin)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _to) returns()
func (_FluxAggregator *FluxAggregatorTransactor) TransferOwnership(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _FluxAggregator.contract.Transact(opts, "transferOwnership", _to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _to) returns()
func (_FluxAggregator *FluxAggregatorSession) TransferOwnership(_to common.Address) (*types.Transaction, error) {
	return _FluxAggregator.Contract.TransferOwnership(&_FluxAggregator.TransactOpts, _to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _to) returns()
func (_FluxAggregator *FluxAggregatorTransactorSession) TransferOwnership(_to common.Address) (*types.Transaction, error) {
	return _FluxAggregator.Contract.TransferOwnership(&_FluxAggregator.TransactOpts, _to)
}

// UpdateAvailableFunds is a paid mutator transaction binding the contract method 0x4f8fc3b5.
//
// Solidity: function updateAvailableFunds() returns()
func (_FluxAggregator *FluxAggregatorTransactor) UpdateAvailableFunds(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FluxAggregator.contract.Transact(opts, "updateAvailableFunds")
}

// UpdateAvailableFunds is a paid mutator transaction binding the contract method 0x4f8fc3b5.
//
// Solidity: function updateAvailableFunds() returns()
func (_FluxAggregator *FluxAggregatorSession) UpdateAvailableFunds() (*types.Transaction, error) {
	return _FluxAggregator.Contract.UpdateAvailableFunds(&_FluxAggregator.TransactOpts)
}

// UpdateAvailableFunds is a paid mutator transaction binding the contract method 0x4f8fc3b5.
//
// Solidity: function updateAvailableFunds() returns()
func (_FluxAggregator *FluxAggregatorTransactorSession) UpdateAvailableFunds() (*types.Transaction, error) {
	return _FluxAggregator.Contract.UpdateAvailableFunds(&_FluxAggregator.TransactOpts)
}

// UpdateFutureRounds is a paid mutator transaction binding the contract method 0x38aa4c72.
//
// Solidity: function updateFutureRounds(uint128 _paymentAmount, uint32 _minSubmissions, uint32 _maxSubmissions, uint32 _restartDelay, uint32 _timeout) returns()
func (_FluxAggregator *FluxAggregatorTransactor) UpdateFutureRounds(opts *bind.TransactOpts, _paymentAmount *big.Int, _minSubmissions uint32, _maxSubmissions uint32, _restartDelay uint32, _timeout uint32) (*types.Transaction, error) {
	return _FluxAggregator.contract.Transact(opts, "updateFutureRounds", _paymentAmount, _minSubmissions, _maxSubmissions, _restartDelay, _timeout)
}

// UpdateFutureRounds is a paid mutator transaction binding the contract method 0x38aa4c72.
//
// Solidity: function updateFutureRounds(uint128 _paymentAmount, uint32 _minSubmissions, uint32 _maxSubmissions, uint32 _restartDelay, uint32 _timeout) returns()
func (_FluxAggregator *FluxAggregatorSession) UpdateFutureRounds(_paymentAmount *big.Int, _minSubmissions uint32, _maxSubmissions uint32, _restartDelay uint32, _timeout uint32) (*types.Transaction, error) {
	return _FluxAggregator.Contract.UpdateFutureRounds(&_FluxAggregator.TransactOpts, _paymentAmount, _minSubmissions, _maxSubmissions, _restartDelay, _timeout)
}

// UpdateFutureRounds is a paid mutator transaction binding the contract method 0x38aa4c72.
//
// Solidity: function updateFutureRounds(uint128 _paymentAmount, uint32 _minSubmissions, uint32 _maxSubmissions, uint32 _restartDelay, uint32 _timeout) returns()
func (_FluxAggregator *FluxAggregatorTransactorSession) UpdateFutureRounds(_paymentAmount *big.Int, _minSubmissions uint32, _maxSubmissions uint32, _restartDelay uint32, _timeout uint32) (*types.Transaction, error) {
	return _FluxAggregator.Contract.UpdateFutureRounds(&_FluxAggregator.TransactOpts, _paymentAmount, _minSubmissions, _maxSubmissions, _restartDelay, _timeout)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0xc1075329.
//
// Solidity: function withdrawFunds(address _recipient, uint256 _amount) returns()
func (_FluxAggregator *FluxAggregatorTransactor) WithdrawFunds(opts *bind.TransactOpts, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _FluxAggregator.contract.Transact(opts, "withdrawFunds", _recipient, _amount)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0xc1075329.
//
// Solidity: function withdrawFunds(address _recipient, uint256 _amount) returns()
func (_FluxAggregator *FluxAggregatorSession) WithdrawFunds(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _FluxAggregator.Contract.WithdrawFunds(&_FluxAggregator.TransactOpts, _recipient, _amount)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0xc1075329.
//
// Solidity: function withdrawFunds(address _recipient, uint256 _amount) returns()
func (_FluxAggregator *FluxAggregatorTransactorSession) WithdrawFunds(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _FluxAggregator.Contract.WithdrawFunds(&_FluxAggregator.TransactOpts, _recipient, _amount)
}

// WithdrawPayment is a paid mutator transaction binding the contract method 0x3d3d7714.
//
// Solidity: function withdrawPayment(address _oracle, address _recipient, uint256 _amount) returns()
func (_FluxAggregator *FluxAggregatorTransactor) WithdrawPayment(opts *bind.TransactOpts, _oracle common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _FluxAggregator.contract.Transact(opts, "withdrawPayment", _oracle, _recipient, _amount)
}

// WithdrawPayment is a paid mutator transaction binding the contract method 0x3d3d7714.
//
// Solidity: function withdrawPayment(address _oracle, address _recipient, uint256 _amount) returns()
func (_FluxAggregator *FluxAggregatorSession) WithdrawPayment(_oracle common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _FluxAggregator.Contract.WithdrawPayment(&_FluxAggregator.TransactOpts, _oracle, _recipient, _amount)
}

// WithdrawPayment is a paid mutator transaction binding the contract method 0x3d3d7714.
//
// Solidity: function withdrawPayment(address _oracle, address _recipient, uint256 _amount) returns()
func (_FluxAggregator *FluxAggregatorTransactorSession) WithdrawPayment(_oracle common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _FluxAggregator.Contract.WithdrawPayment(&_FluxAggregator.TransactOpts, _oracle, _recipient, _amount)
}

// FluxAggregatorAnswerUpdatedIterator is returned from FilterAnswerUpdated and is used to iterate over the raw logs and unpacked data for AnswerUpdated events raised by the FluxAggregator contract.
type FluxAggregatorAnswerUpdatedIterator struct {
	Event *FluxAggregatorAnswerUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FluxAggregatorAnswerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FluxAggregatorAnswerUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FluxAggregatorAnswerUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FluxAggregatorAnswerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FluxAggregatorAnswerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FluxAggregatorAnswerUpdated represents a AnswerUpdated event raised by the FluxAggregator contract.
type FluxAggregatorAnswerUpdated struct {
	Current   *big.Int
	RoundId   *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAnswerUpdated is a free log retrieval operation binding the contract event 0x0559884fd3a460db3073b7fc896cc77986f16e378210ded43186175bf646fc5f.
//
// Solidity: event AnswerUpdated(int256 indexed current, uint256 indexed roundId, uint256 timestamp)
func (_FluxAggregator *FluxAggregatorFilterer) FilterAnswerUpdated(opts *bind.FilterOpts, current []*big.Int, roundId []*big.Int) (*FluxAggregatorAnswerUpdatedIterator, error) {

	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}
	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _FluxAggregator.contract.FilterLogs(opts, "AnswerUpdated", currentRule, roundIdRule)
	if err != nil {
		return nil, err
	}
	return &FluxAggregatorAnswerUpdatedIterator{contract: _FluxAggregator.contract, event: "AnswerUpdated", logs: logs, sub: sub}, nil
}

// WatchAnswerUpdated is a free log subscription operation binding the contract event 0x0559884fd3a460db3073b7fc896cc77986f16e378210ded43186175bf646fc5f.
//
// Solidity: event AnswerUpdated(int256 indexed current, uint256 indexed roundId, uint256 timestamp)
func (_FluxAggregator *FluxAggregatorFilterer) WatchAnswerUpdated(opts *bind.WatchOpts, sink chan<- *FluxAggregatorAnswerUpdated, current []*big.Int, roundId []*big.Int) (event.Subscription, error) {

	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}
	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _FluxAggregator.contract.WatchLogs(opts, "AnswerUpdated", currentRule, roundIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FluxAggregatorAnswerUpdated)
				if err := _FluxAggregator.contract.UnpackLog(event, "AnswerUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAnswerUpdated is a log parse operation binding the contract event 0x0559884fd3a460db3073b7fc896cc77986f16e378210ded43186175bf646fc5f.
//
// Solidity: event AnswerUpdated(int256 indexed current, uint256 indexed roundId, uint256 timestamp)
func (_FluxAggregator *FluxAggregatorFilterer) ParseAnswerUpdated(log types.Log) (*FluxAggregatorAnswerUpdated, error) {
	event := new(FluxAggregatorAnswerUpdated)
	if err := _FluxAggregator.contract.UnpackLog(event, "AnswerUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FluxAggregatorAvailableFundsUpdatedIterator is returned from FilterAvailableFundsUpdated and is used to iterate over the raw logs and unpacked data for AvailableFundsUpdated events raised by the FluxAggregator contract.
type FluxAggregatorAvailableFundsUpdatedIterator struct {
	Event *FluxAggregatorAvailableFundsUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FluxAggregatorAvailableFundsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FluxAggregatorAvailableFundsUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FluxAggregatorAvailableFundsUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FluxAggregatorAvailableFundsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FluxAggregatorAvailableFundsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FluxAggregatorAvailableFundsUpdated represents a AvailableFundsUpdated event raised by the FluxAggregator contract.
type FluxAggregatorAvailableFundsUpdated struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAvailableFundsUpdated is a free log retrieval operation binding the contract event 0xfe25c73e3b9089fac37d55c4c7efcba6f04af04cebd2fc4d6d7dbb07e1e5234f.
//
// Solidity: event AvailableFundsUpdated(uint256 indexed amount)
func (_FluxAggregator *FluxAggregatorFilterer) FilterAvailableFundsUpdated(opts *bind.FilterOpts, amount []*big.Int) (*FluxAggregatorAvailableFundsUpdatedIterator, error) {

	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _FluxAggregator.contract.FilterLogs(opts, "AvailableFundsUpdated", amountRule)
	if err != nil {
		return nil, err
	}
	return &FluxAggregatorAvailableFundsUpdatedIterator{contract: _FluxAggregator.contract, event: "AvailableFundsUpdated", logs: logs, sub: sub}, nil
}

// WatchAvailableFundsUpdated is a free log subscription operation binding the contract event 0xfe25c73e3b9089fac37d55c4c7efcba6f04af04cebd2fc4d6d7dbb07e1e5234f.
//
// Solidity: event AvailableFundsUpdated(uint256 indexed amount)
func (_FluxAggregator *FluxAggregatorFilterer) WatchAvailableFundsUpdated(opts *bind.WatchOpts, sink chan<- *FluxAggregatorAvailableFundsUpdated, amount []*big.Int) (event.Subscription, error) {

	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _FluxAggregator.contract.WatchLogs(opts, "AvailableFundsUpdated", amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FluxAggregatorAvailableFundsUpdated)
				if err := _FluxAggregator.contract.UnpackLog(event, "AvailableFundsUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAvailableFundsUpdated is a log parse operation binding the contract event 0xfe25c73e3b9089fac37d55c4c7efcba6f04af04cebd2fc4d6d7dbb07e1e5234f.
//
// Solidity: event AvailableFundsUpdated(uint256 indexed amount)
func (_FluxAggregator *FluxAggregatorFilterer) ParseAvailableFundsUpdated(log types.Log) (*FluxAggregatorAvailableFundsUpdated, error) {
	event := new(FluxAggregatorAvailableFundsUpdated)
	if err := _FluxAggregator.contract.UnpackLog(event, "AvailableFundsUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FluxAggregatorNewRoundIterator is returned from FilterNewRound and is used to iterate over the raw logs and unpacked data for NewRound events raised by the FluxAggregator contract.
type FluxAggregatorNewRoundIterator struct {
	Event *FluxAggregatorNewRound // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FluxAggregatorNewRoundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FluxAggregatorNewRound)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FluxAggregatorNewRound)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FluxAggregatorNewRoundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FluxAggregatorNewRoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FluxAggregatorNewRound represents a NewRound event raised by the FluxAggregator contract.
type FluxAggregatorNewRound struct {
	RoundId   *big.Int
	StartedBy common.Address
	StartedAt *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewRound is a free log retrieval operation binding the contract event 0x0109fc6f55cf40689f02fbaad7af7fe7bbac8a3d2186600afc7d3e10cac60271.
//
// Solidity: event NewRound(uint256 indexed roundId, address indexed startedBy, uint256 startedAt)
func (_FluxAggregator *FluxAggregatorFilterer) FilterNewRound(opts *bind.FilterOpts, roundId []*big.Int, startedBy []common.Address) (*FluxAggregatorNewRoundIterator, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}
	var startedByRule []interface{}
	for _, startedByItem := range startedBy {
		startedByRule = append(startedByRule, startedByItem)
	}

	logs, sub, err := _FluxAggregator.contract.FilterLogs(opts, "NewRound", roundIdRule, startedByRule)
	if err != nil {
		return nil, err
	}
	return &FluxAggregatorNewRoundIterator{contract: _FluxAggregator.contract, event: "NewRound", logs: logs, sub: sub}, nil
}

// WatchNewRound is a free log subscription operation binding the contract event 0x0109fc6f55cf40689f02fbaad7af7fe7bbac8a3d2186600afc7d3e10cac60271.
//
// Solidity: event NewRound(uint256 indexed roundId, address indexed startedBy, uint256 startedAt)
func (_FluxAggregator *FluxAggregatorFilterer) WatchNewRound(opts *bind.WatchOpts, sink chan<- *FluxAggregatorNewRound, roundId []*big.Int, startedBy []common.Address) (event.Subscription, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}
	var startedByRule []interface{}
	for _, startedByItem := range startedBy {
		startedByRule = append(startedByRule, startedByItem)
	}

	logs, sub, err := _FluxAggregator.contract.WatchLogs(opts, "NewRound", roundIdRule, startedByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FluxAggregatorNewRound)
				if err := _FluxAggregator.contract.UnpackLog(event, "NewRound", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewRound is a log parse operation binding the contract event 0x0109fc6f55cf40689f02fbaad7af7fe7bbac8a3d2186600afc7d3e10cac60271.
//
// Solidity: event NewRound(uint256 indexed roundId, address indexed startedBy, uint256 startedAt)
func (_FluxAggregator *FluxAggregatorFilterer) ParseNewRound(log types.Log) (*FluxAggregatorNewRound, error) {
	event := new(FluxAggregatorNewRound)
	if err := _FluxAggregator.contract.UnpackLog(event, "NewRound", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FluxAggregatorOracleAdminUpdateRequestedIterator is returned from FilterOracleAdminUpdateRequested and is used to iterate over the raw logs and unpacked data for OracleAdminUpdateRequested events raised by the FluxAggregator contract.
type FluxAggregatorOracleAdminUpdateRequestedIterator struct {
	Event *FluxAggregatorOracleAdminUpdateRequested // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FluxAggregatorOracleAdminUpdateRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FluxAggregatorOracleAdminUpdateRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FluxAggregatorOracleAdminUpdateRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FluxAggregatorOracleAdminUpdateRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FluxAggregatorOracleAdminUpdateRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FluxAggregatorOracleAdminUpdateRequested represents a OracleAdminUpdateRequested event raised by the FluxAggregator contract.
type FluxAggregatorOracleAdminUpdateRequested struct {
	Oracle   common.Address
	Admin    common.Address
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOracleAdminUpdateRequested is a free log retrieval operation binding the contract event 0xb79bf2e89c2d70dde91d2991fb1ea69b7e478061ad7c04ed5b02b96bc52b8104.
//
// Solidity: event OracleAdminUpdateRequested(address indexed oracle, address admin, address newAdmin)
func (_FluxAggregator *FluxAggregatorFilterer) FilterOracleAdminUpdateRequested(opts *bind.FilterOpts, oracle []common.Address) (*FluxAggregatorOracleAdminUpdateRequestedIterator, error) {

	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _FluxAggregator.contract.FilterLogs(opts, "OracleAdminUpdateRequested", oracleRule)
	if err != nil {
		return nil, err
	}
	return &FluxAggregatorOracleAdminUpdateRequestedIterator{contract: _FluxAggregator.contract, event: "OracleAdminUpdateRequested", logs: logs, sub: sub}, nil
}

// WatchOracleAdminUpdateRequested is a free log subscription operation binding the contract event 0xb79bf2e89c2d70dde91d2991fb1ea69b7e478061ad7c04ed5b02b96bc52b8104.
//
// Solidity: event OracleAdminUpdateRequested(address indexed oracle, address admin, address newAdmin)
func (_FluxAggregator *FluxAggregatorFilterer) WatchOracleAdminUpdateRequested(opts *bind.WatchOpts, sink chan<- *FluxAggregatorOracleAdminUpdateRequested, oracle []common.Address) (event.Subscription, error) {

	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _FluxAggregator.contract.WatchLogs(opts, "OracleAdminUpdateRequested", oracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FluxAggregatorOracleAdminUpdateRequested)
				if err := _FluxAggregator.contract.UnpackLog(event, "OracleAdminUpdateRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOracleAdminUpdateRequested is a log parse operation binding the contract event 0xb79bf2e89c2d70dde91d2991fb1ea69b7e478061ad7c04ed5b02b96bc52b8104.
//
// Solidity: event OracleAdminUpdateRequested(address indexed oracle, address admin, address newAdmin)
func (_FluxAggregator *FluxAggregatorFilterer) ParseOracleAdminUpdateRequested(log types.Log) (*FluxAggregatorOracleAdminUpdateRequested, error) {
	event := new(FluxAggregatorOracleAdminUpdateRequested)
	if err := _FluxAggregator.contract.UnpackLog(event, "OracleAdminUpdateRequested", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FluxAggregatorOracleAdminUpdatedIterator is returned from FilterOracleAdminUpdated and is used to iterate over the raw logs and unpacked data for OracleAdminUpdated events raised by the FluxAggregator contract.
type FluxAggregatorOracleAdminUpdatedIterator struct {
	Event *FluxAggregatorOracleAdminUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FluxAggregatorOracleAdminUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FluxAggregatorOracleAdminUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FluxAggregatorOracleAdminUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FluxAggregatorOracleAdminUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FluxAggregatorOracleAdminUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FluxAggregatorOracleAdminUpdated represents a OracleAdminUpdated event raised by the FluxAggregator contract.
type FluxAggregatorOracleAdminUpdated struct {
	Oracle   common.Address
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOracleAdminUpdated is a free log retrieval operation binding the contract event 0x0c5055390645c15a4be9a21b3f8d019153dcb4a0c125685da6eb84048e2fe904.
//
// Solidity: event OracleAdminUpdated(address indexed oracle, address indexed newAdmin)
func (_FluxAggregator *FluxAggregatorFilterer) FilterOracleAdminUpdated(opts *bind.FilterOpts, oracle []common.Address, newAdmin []common.Address) (*FluxAggregatorOracleAdminUpdatedIterator, error) {

	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _FluxAggregator.contract.FilterLogs(opts, "OracleAdminUpdated", oracleRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return &FluxAggregatorOracleAdminUpdatedIterator{contract: _FluxAggregator.contract, event: "OracleAdminUpdated", logs: logs, sub: sub}, nil
}

// WatchOracleAdminUpdated is a free log subscription operation binding the contract event 0x0c5055390645c15a4be9a21b3f8d019153dcb4a0c125685da6eb84048e2fe904.
//
// Solidity: event OracleAdminUpdated(address indexed oracle, address indexed newAdmin)
func (_FluxAggregator *FluxAggregatorFilterer) WatchOracleAdminUpdated(opts *bind.WatchOpts, sink chan<- *FluxAggregatorOracleAdminUpdated, oracle []common.Address, newAdmin []common.Address) (event.Subscription, error) {

	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _FluxAggregator.contract.WatchLogs(opts, "OracleAdminUpdated", oracleRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FluxAggregatorOracleAdminUpdated)
				if err := _FluxAggregator.contract.UnpackLog(event, "OracleAdminUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOracleAdminUpdated is a log parse operation binding the contract event 0x0c5055390645c15a4be9a21b3f8d019153dcb4a0c125685da6eb84048e2fe904.
//
// Solidity: event OracleAdminUpdated(address indexed oracle, address indexed newAdmin)
func (_FluxAggregator *FluxAggregatorFilterer) ParseOracleAdminUpdated(log types.Log) (*FluxAggregatorOracleAdminUpdated, error) {
	event := new(FluxAggregatorOracleAdminUpdated)
	if err := _FluxAggregator.contract.UnpackLog(event, "OracleAdminUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FluxAggregatorOraclePermissionsUpdatedIterator is returned from FilterOraclePermissionsUpdated and is used to iterate over the raw logs and unpacked data for OraclePermissionsUpdated events raised by the FluxAggregator contract.
type FluxAggregatorOraclePermissionsUpdatedIterator struct {
	Event *FluxAggregatorOraclePermissionsUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FluxAggregatorOraclePermissionsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FluxAggregatorOraclePermissionsUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FluxAggregatorOraclePermissionsUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FluxAggregatorOraclePermissionsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FluxAggregatorOraclePermissionsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FluxAggregatorOraclePermissionsUpdated represents a OraclePermissionsUpdated event raised by the FluxAggregator contract.
type FluxAggregatorOraclePermissionsUpdated struct {
	Oracle      common.Address
	Whitelisted bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOraclePermissionsUpdated is a free log retrieval operation binding the contract event 0x18dd09695e4fbdae8d1a5edb11221eb04564269c29a089b9753a6535c54ba92e.
//
// Solidity: event OraclePermissionsUpdated(address indexed oracle, bool indexed whitelisted)
func (_FluxAggregator *FluxAggregatorFilterer) FilterOraclePermissionsUpdated(opts *bind.FilterOpts, oracle []common.Address, whitelisted []bool) (*FluxAggregatorOraclePermissionsUpdatedIterator, error) {

	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}
	var whitelistedRule []interface{}
	for _, whitelistedItem := range whitelisted {
		whitelistedRule = append(whitelistedRule, whitelistedItem)
	}

	logs, sub, err := _FluxAggregator.contract.FilterLogs(opts, "OraclePermissionsUpdated", oracleRule, whitelistedRule)
	if err != nil {
		return nil, err
	}
	return &FluxAggregatorOraclePermissionsUpdatedIterator{contract: _FluxAggregator.contract, event: "OraclePermissionsUpdated", logs: logs, sub: sub}, nil
}

// WatchOraclePermissionsUpdated is a free log subscription operation binding the contract event 0x18dd09695e4fbdae8d1a5edb11221eb04564269c29a089b9753a6535c54ba92e.
//
// Solidity: event OraclePermissionsUpdated(address indexed oracle, bool indexed whitelisted)
func (_FluxAggregator *FluxAggregatorFilterer) WatchOraclePermissionsUpdated(opts *bind.WatchOpts, sink chan<- *FluxAggregatorOraclePermissionsUpdated, oracle []common.Address, whitelisted []bool) (event.Subscription, error) {

	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}
	var whitelistedRule []interface{}
	for _, whitelistedItem := range whitelisted {
		whitelistedRule = append(whitelistedRule, whitelistedItem)
	}

	logs, sub, err := _FluxAggregator.contract.WatchLogs(opts, "OraclePermissionsUpdated", oracleRule, whitelistedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FluxAggregatorOraclePermissionsUpdated)
				if err := _FluxAggregator.contract.UnpackLog(event, "OraclePermissionsUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOraclePermissionsUpdated is a log parse operation binding the contract event 0x18dd09695e4fbdae8d1a5edb11221eb04564269c29a089b9753a6535c54ba92e.
//
// Solidity: event OraclePermissionsUpdated(address indexed oracle, bool indexed whitelisted)
func (_FluxAggregator *FluxAggregatorFilterer) ParseOraclePermissionsUpdated(log types.Log) (*FluxAggregatorOraclePermissionsUpdated, error) {
	event := new(FluxAggregatorOraclePermissionsUpdated)
	if err := _FluxAggregator.contract.UnpackLog(event, "OraclePermissionsUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FluxAggregatorOwnershipTransferRequestedIterator is returned from FilterOwnershipTransferRequested and is used to iterate over the raw logs and unpacked data for OwnershipTransferRequested events raised by the FluxAggregator contract.
type FluxAggregatorOwnershipTransferRequestedIterator struct {
	Event *FluxAggregatorOwnershipTransferRequested // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FluxAggregatorOwnershipTransferRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FluxAggregatorOwnershipTransferRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FluxAggregatorOwnershipTransferRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FluxAggregatorOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FluxAggregatorOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FluxAggregatorOwnershipTransferRequested represents a OwnershipTransferRequested event raised by the FluxAggregator contract.
type FluxAggregatorOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferRequested is a free log retrieval operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_FluxAggregator *FluxAggregatorFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*FluxAggregatorOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FluxAggregator.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &FluxAggregatorOwnershipTransferRequestedIterator{contract: _FluxAggregator.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferRequested is a free log subscription operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_FluxAggregator *FluxAggregatorFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *FluxAggregatorOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FluxAggregator.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FluxAggregatorOwnershipTransferRequested)
				if err := _FluxAggregator.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferRequested is a log parse operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_FluxAggregator *FluxAggregatorFilterer) ParseOwnershipTransferRequested(log types.Log) (*FluxAggregatorOwnershipTransferRequested, error) {
	event := new(FluxAggregatorOwnershipTransferRequested)
	if err := _FluxAggregator.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FluxAggregatorOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FluxAggregator contract.
type FluxAggregatorOwnershipTransferredIterator struct {
	Event *FluxAggregatorOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FluxAggregatorOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FluxAggregatorOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FluxAggregatorOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FluxAggregatorOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FluxAggregatorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FluxAggregatorOwnershipTransferred represents a OwnershipTransferred event raised by the FluxAggregator contract.
type FluxAggregatorOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_FluxAggregator *FluxAggregatorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*FluxAggregatorOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FluxAggregator.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &FluxAggregatorOwnershipTransferredIterator{contract: _FluxAggregator.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_FluxAggregator *FluxAggregatorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FluxAggregatorOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FluxAggregator.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FluxAggregatorOwnershipTransferred)
				if err := _FluxAggregator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_FluxAggregator *FluxAggregatorFilterer) ParseOwnershipTransferred(log types.Log) (*FluxAggregatorOwnershipTransferred, error) {
	event := new(FluxAggregatorOwnershipTransferred)
	if err := _FluxAggregator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FluxAggregatorRequesterPermissionsSetIterator is returned from FilterRequesterPermissionsSet and is used to iterate over the raw logs and unpacked data for RequesterPermissionsSet events raised by the FluxAggregator contract.
type FluxAggregatorRequesterPermissionsSetIterator struct {
	Event *FluxAggregatorRequesterPermissionsSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FluxAggregatorRequesterPermissionsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FluxAggregatorRequesterPermissionsSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FluxAggregatorRequesterPermissionsSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FluxAggregatorRequesterPermissionsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FluxAggregatorRequesterPermissionsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FluxAggregatorRequesterPermissionsSet represents a RequesterPermissionsSet event raised by the FluxAggregator contract.
type FluxAggregatorRequesterPermissionsSet struct {
	Requester  common.Address
	Authorized bool
	Delay      uint32
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRequesterPermissionsSet is a free log retrieval operation binding the contract event 0xc3df5a754e002718f2e10804b99e6605e7c701d95cec9552c7680ca2b6f2820a.
//
// Solidity: event RequesterPermissionsSet(address indexed requester, bool authorized, uint32 delay)
func (_FluxAggregator *FluxAggregatorFilterer) FilterRequesterPermissionsSet(opts *bind.FilterOpts, requester []common.Address) (*FluxAggregatorRequesterPermissionsSetIterator, error) {

	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _FluxAggregator.contract.FilterLogs(opts, "RequesterPermissionsSet", requesterRule)
	if err != nil {
		return nil, err
	}
	return &FluxAggregatorRequesterPermissionsSetIterator{contract: _FluxAggregator.contract, event: "RequesterPermissionsSet", logs: logs, sub: sub}, nil
}

// WatchRequesterPermissionsSet is a free log subscription operation binding the contract event 0xc3df5a754e002718f2e10804b99e6605e7c701d95cec9552c7680ca2b6f2820a.
//
// Solidity: event RequesterPermissionsSet(address indexed requester, bool authorized, uint32 delay)
func (_FluxAggregator *FluxAggregatorFilterer) WatchRequesterPermissionsSet(opts *bind.WatchOpts, sink chan<- *FluxAggregatorRequesterPermissionsSet, requester []common.Address) (event.Subscription, error) {

	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _FluxAggregator.contract.WatchLogs(opts, "RequesterPermissionsSet", requesterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FluxAggregatorRequesterPermissionsSet)
				if err := _FluxAggregator.contract.UnpackLog(event, "RequesterPermissionsSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRequesterPermissionsSet is a log parse operation binding the contract event 0xc3df5a754e002718f2e10804b99e6605e7c701d95cec9552c7680ca2b6f2820a.
//
// Solidity: event RequesterPermissionsSet(address indexed requester, bool authorized, uint32 delay)
func (_FluxAggregator *FluxAggregatorFilterer) ParseRequesterPermissionsSet(log types.Log) (*FluxAggregatorRequesterPermissionsSet, error) {
	event := new(FluxAggregatorRequesterPermissionsSet)
	if err := _FluxAggregator.contract.UnpackLog(event, "RequesterPermissionsSet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FluxAggregatorRoundDetailsUpdatedIterator is returned from FilterRoundDetailsUpdated and is used to iterate over the raw logs and unpacked data for RoundDetailsUpdated events raised by the FluxAggregator contract.
type FluxAggregatorRoundDetailsUpdatedIterator struct {
	Event *FluxAggregatorRoundDetailsUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FluxAggregatorRoundDetailsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FluxAggregatorRoundDetailsUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FluxAggregatorRoundDetailsUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FluxAggregatorRoundDetailsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FluxAggregatorRoundDetailsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FluxAggregatorRoundDetailsUpdated represents a RoundDetailsUpdated event raised by the FluxAggregator contract.
type FluxAggregatorRoundDetailsUpdated struct {
	PaymentAmount      *big.Int
	MinSubmissionCount uint32
	MaxSubmissionCount uint32
	RestartDelay       uint32
	Timeout            uint32
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterRoundDetailsUpdated is a free log retrieval operation binding the contract event 0x56800c9d1ed723511246614d15e58cfcde15b6a33c245b5c961b689c1890fd8f.
//
// Solidity: event RoundDetailsUpdated(uint128 indexed paymentAmount, uint32 indexed minSubmissionCount, uint32 indexed maxSubmissionCount, uint32 restartDelay, uint32 timeout)
func (_FluxAggregator *FluxAggregatorFilterer) FilterRoundDetailsUpdated(opts *bind.FilterOpts, paymentAmount []*big.Int, minSubmissionCount []uint32, maxSubmissionCount []uint32) (*FluxAggregatorRoundDetailsUpdatedIterator, error) {

	var paymentAmountRule []interface{}
	for _, paymentAmountItem := range paymentAmount {
		paymentAmountRule = append(paymentAmountRule, paymentAmountItem)
	}
	var minSubmissionCountRule []interface{}
	for _, minSubmissionCountItem := range minSubmissionCount {
		minSubmissionCountRule = append(minSubmissionCountRule, minSubmissionCountItem)
	}
	var maxSubmissionCountRule []interface{}
	for _, maxSubmissionCountItem := range maxSubmissionCount {
		maxSubmissionCountRule = append(maxSubmissionCountRule, maxSubmissionCountItem)
	}

	logs, sub, err := _FluxAggregator.contract.FilterLogs(opts, "RoundDetailsUpdated", paymentAmountRule, minSubmissionCountRule, maxSubmissionCountRule)
	if err != nil {
		return nil, err
	}
	return &FluxAggregatorRoundDetailsUpdatedIterator{contract: _FluxAggregator.contract, event: "RoundDetailsUpdated", logs: logs, sub: sub}, nil
}

// WatchRoundDetailsUpdated is a free log subscription operation binding the contract event 0x56800c9d1ed723511246614d15e58cfcde15b6a33c245b5c961b689c1890fd8f.
//
// Solidity: event RoundDetailsUpdated(uint128 indexed paymentAmount, uint32 indexed minSubmissionCount, uint32 indexed maxSubmissionCount, uint32 restartDelay, uint32 timeout)
func (_FluxAggregator *FluxAggregatorFilterer) WatchRoundDetailsUpdated(opts *bind.WatchOpts, sink chan<- *FluxAggregatorRoundDetailsUpdated, paymentAmount []*big.Int, minSubmissionCount []uint32, maxSubmissionCount []uint32) (event.Subscription, error) {

	var paymentAmountRule []interface{}
	for _, paymentAmountItem := range paymentAmount {
		paymentAmountRule = append(paymentAmountRule, paymentAmountItem)
	}
	var minSubmissionCountRule []interface{}
	for _, minSubmissionCountItem := range minSubmissionCount {
		minSubmissionCountRule = append(minSubmissionCountRule, minSubmissionCountItem)
	}
	var maxSubmissionCountRule []interface{}
	for _, maxSubmissionCountItem := range maxSubmissionCount {
		maxSubmissionCountRule = append(maxSubmissionCountRule, maxSubmissionCountItem)
	}

	logs, sub, err := _FluxAggregator.contract.WatchLogs(opts, "RoundDetailsUpdated", paymentAmountRule, minSubmissionCountRule, maxSubmissionCountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FluxAggregatorRoundDetailsUpdated)
				if err := _FluxAggregator.contract.UnpackLog(event, "RoundDetailsUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoundDetailsUpdated is a log parse operation binding the contract event 0x56800c9d1ed723511246614d15e58cfcde15b6a33c245b5c961b689c1890fd8f.
//
// Solidity: event RoundDetailsUpdated(uint128 indexed paymentAmount, uint32 indexed minSubmissionCount, uint32 indexed maxSubmissionCount, uint32 restartDelay, uint32 timeout)
func (_FluxAggregator *FluxAggregatorFilterer) ParseRoundDetailsUpdated(log types.Log) (*FluxAggregatorRoundDetailsUpdated, error) {
	event := new(FluxAggregatorRoundDetailsUpdated)
	if err := _FluxAggregator.contract.UnpackLog(event, "RoundDetailsUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FluxAggregatorSubmissionReceivedIterator is returned from FilterSubmissionReceived and is used to iterate over the raw logs and unpacked data for SubmissionReceived events raised by the FluxAggregator contract.
type FluxAggregatorSubmissionReceivedIterator struct {
	Event *FluxAggregatorSubmissionReceived // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FluxAggregatorSubmissionReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FluxAggregatorSubmissionReceived)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FluxAggregatorSubmissionReceived)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FluxAggregatorSubmissionReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FluxAggregatorSubmissionReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FluxAggregatorSubmissionReceived represents a SubmissionReceived event raised by the FluxAggregator contract.
type FluxAggregatorSubmissionReceived struct {
	Submission *big.Int
	Round      uint32
	Oracle     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSubmissionReceived is a free log retrieval operation binding the contract event 0x92e98423f8adac6e64d0608e519fd1cefb861498385c6dee70d58fc926ddc68c.
//
// Solidity: event SubmissionReceived(int256 indexed submission, uint32 indexed round, address indexed oracle)
func (_FluxAggregator *FluxAggregatorFilterer) FilterSubmissionReceived(opts *bind.FilterOpts, submission []*big.Int, round []uint32, oracle []common.Address) (*FluxAggregatorSubmissionReceivedIterator, error) {

	var submissionRule []interface{}
	for _, submissionItem := range submission {
		submissionRule = append(submissionRule, submissionItem)
	}
	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}
	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _FluxAggregator.contract.FilterLogs(opts, "SubmissionReceived", submissionRule, roundRule, oracleRule)
	if err != nil {
		return nil, err
	}
	return &FluxAggregatorSubmissionReceivedIterator{contract: _FluxAggregator.contract, event: "SubmissionReceived", logs: logs, sub: sub}, nil
}

// WatchSubmissionReceived is a free log subscription operation binding the contract event 0x92e98423f8adac6e64d0608e519fd1cefb861498385c6dee70d58fc926ddc68c.
//
// Solidity: event SubmissionReceived(int256 indexed submission, uint32 indexed round, address indexed oracle)
func (_FluxAggregator *FluxAggregatorFilterer) WatchSubmissionReceived(opts *bind.WatchOpts, sink chan<- *FluxAggregatorSubmissionReceived, submission []*big.Int, round []uint32, oracle []common.Address) (event.Subscription, error) {

	var submissionRule []interface{}
	for _, submissionItem := range submission {
		submissionRule = append(submissionRule, submissionItem)
	}
	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}
	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _FluxAggregator.contract.WatchLogs(opts, "SubmissionReceived", submissionRule, roundRule, oracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FluxAggregatorSubmissionReceived)
				if err := _FluxAggregator.contract.UnpackLog(event, "SubmissionReceived", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSubmissionReceived is a log parse operation binding the contract event 0x92e98423f8adac6e64d0608e519fd1cefb861498385c6dee70d58fc926ddc68c.
//
// Solidity: event SubmissionReceived(int256 indexed submission, uint32 indexed round, address indexed oracle)
func (_FluxAggregator *FluxAggregatorFilterer) ParseSubmissionReceived(log types.Log) (*FluxAggregatorSubmissionReceived, error) {
	event := new(FluxAggregatorSubmissionReceived)
	if err := _FluxAggregator.contract.UnpackLog(event, "SubmissionReceived", log); err != nil {
		return nil, err
	}
	return event, nil
}
