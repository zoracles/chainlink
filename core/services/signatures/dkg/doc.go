// Package collaboration contains network logic for making a threshold signature
//
// Intended usage
//
// This is to be used once a consensus has developed that a message should be
// signed. Establishing that consensus is out of scope, here, and will probably
// involve application-specific logic.
//
// The nodes find each other through a DHT. Chainlink can operate bootstrap
// nodes for it, but other people can offer this service if they wish. The nodes
// advertise themselves under their public keys.
//
// 0. The nodes use the ordering in the service agreement, to get the index for
//    each node.
//
// 1. Each node constructs a
//
// 1. Each node asks every other node for their coefficient commitments, and
//    the corresponding secret share. Those are sent with a signature.
//
// 2. Each node asks every other node for complaints, and for some summary of
//    the public data received.

package collaboration

// Risks
//
// The risk with a mesh network is that it makes eclipse attacks on a
// participant much easier.

// Future development
//
// This section is for work which is currently out of scope, but important
// enough to record for later.
//
// At some point, we will want a smarter distribution scheme, but I think to
// start with we should go with everyone just asking the appropriate people for
// their data.
//
// At some point, we will want complaints to take place on the blockchain, so
// that there can be financial accountability for failures. Verification of
// these complaints should come in the form of votes.

// Notes
//
// What if we used a hierarchical deterministic wallet to generate the
// polynomial coefficients? That would *massively* reduce bandwidth! This
// doesn't work, because the secret keys in hierarchical deterministic wallets
// involve a constant key, plus a publicly known value, so you end up with a
// linear system with a single unknown.
