<!-- This document is best read as HTML rendered from the markdown. 

It contains some italicized mathematical notation, and the asterices in the
markdown syntax for italicization are a bit confusing, in the raw text.

E.g. "*Aᵢₖ=aᵢₖG*" means "Aᵢₖ=aᵢₖG", with all the roman characters italicized. 

To render as HTML locally, "pip install grip", then open localhost:6419.
-->

# Network interactions for distributed key generation

This describes the network messages participants must send each other, to
construct a distributed public/secret key pair, as outlined in figure 4, p. 74
of [Secure Distributed Key Generation for Discrete-Log Based Cryptosystems](
https://link.springer.com/content/pdf/10.1007/s00145-006-0347-3.pdf). The
purpose is construction of threshold signatures.

Note that as described in that paper, this allows an adversary a small degree of
control over the public key (through exiting the protocol early), but, as it
argues, that control is insufficent to seriously weaken the security of the
signature protocol. Don't use it in contexts which depend on the public key
being a uniform sample! Use the main method of that paper, instead (but that is
twice the bandwidth, and an extra round.)

Since in the end we can't trust the network to deliver messages for us, a key
idea for the final version of the workflow is to use a smart contract on the
blockchain itself as a coordination mechanism, because it's explicitly designed
for censorship resistance. To keep things cheap on the happy path, this trick is
only used when failures are observed.

There are some complex ideas in what follows for allowing on-chain
accountability for off-chain misbehavior in later versions. An alternative would
be for signatures to sign the desired message, plus a message describing the
participants' performances. This could be combined with the original message
just by hashing the concatenation of the hashes of the two messages. This
lengthens the signatures (you have to send the hash of the performance message
along, for verifiers who only want to verify the original message), but does
lead to a much simpler protocol. It would allow for accountability up to
generation of the ephemeral key, which is the most expensive part. It could not
be used for the construction of the actual signature, obviously.

## Intended full workflow

While this is an ordered list, many of the steps are independent, especially
between nodes. It is not necessary to synchronize the nodes at every step,
unless that's explicitly indicated.

0. **In the initial version**, we'll just ignore the smart-contract coordinator.
   In later versions the process kicks off with an event from a smart contract.
   The contract will be an early part of the development, though, if only to
   store the participants which are still in good standing so that the nodes can
   look up who to talk to.
1. <a href="node-indices"/> The index assigned to each node is its ordinal index
   in some list, such as the service agreement which the group is signing
   reports to.
2. Nodes contact each other through a DHT, looking up the other participants'
   host info via their public keys. Chainlink will run a bootstrap node for
   this DHT.
3. Every node requests the public keys of the coefficients (*Aᵢₖ=aᵢₖG*, in
   section 2.4, step 3 of [Stinson and Strobl 2001
   ](https://www.researchgate.net/profile/Willy_Susilo/publication/242499559_Information_Security_and_Privacy_13th_Australasian_Conference_ACISP_2008_Wollongong_Australia_July_7-9_2008_Proceedings/links/00b495314f3bcaaa46000000.pdf#page=426))
   of every other node.
   1. In the happy path, the nodes comply with these requests, sending the same
      public keys to every other node. In the final version of the protocol,
      these should be preceded by a signed merkle-tree commitment to the full
      set of keys, and the keys should be sent as signed batches of merkle
      subtrees, along with the merkle path to the initial commitment. 
      
      **In the initial version** assuming the happy path, the commitments can
      just be sent over as a block.
   2. Failure 1: Some node sends inconsistent coefficient-public-keys *Aᵢₖ* to
      other nodes. In the final version the nodes publish the initial
      commitments they've received from every other node, and every node
      compares them. Any inconsistencies are reported to a smart contract where
      the offending node has a stake, and it is slashed (this is the reason the
      commitments should be sent in small batches with independent signatures,
      so that on-chain verification is not too expensive.) The node is removed
      from further participation in the DKG. **In the initial version**, the
      offending node is just removed without being slashed, and the complaint is
      simply shared among the nodes.
   3. Failure 2: A node fails to send some portion of its commitments to some
      other node. Punishing this directly is more complex, because it cannot be
      cryptographically verified, and might depend on network failures out of
      the node operator's control. It's tempting to add a financial penalty for
      consistent failure to deliver, but not strictly necessary, since the
      protocol is designed to continue if some fraction of nodes drop out. Such
      a penalty should be encoded as an extension to the message to be signed,
      because any kind of on-chain negotiation about a complaint would slow the
      signature process down dramatically.
      
      In any case, it will be useful for a pair of nodes with a good connection
      to be able to request a third node's data from each other, in order to
      route around partial network failures. For this purpose, the nodes can
      publish to each other which portions of the data they have received. Since
      the messages are signed, there is no scope here for intermediaries to
      corrupt the data they forward.
      
      Obviously, that is **not needed for the initial version.**
4. All nodes *j* ask all other nodes *i* for their secret share (*fᵢ*(*j*), in
   the notation of step 1 of section 2.4 in Stinson & Strobl, where *j* is the
   [index described above](#node-indices).) Once they've verified their shares
   from a particular node, as in equation (2) in Stinson and Strobl, they
   announce to the network. In a later version, the response with the shares
   must include a merkle commitment to the calculation verifying the shares.
   (These are just sums of scalar multiples of the transmitting node's
   commitments, so they can be arranged in nice binary trees using
   associativity.) Nodes broadcast to the group who they've received correct
   shares from.
   1. If some node doesn't respond to some other node with its secret shares by
      a certain number of blocks, the recipient can post a request on the
      coordinating contract. If the mandated sender repeatedly fails to respond
      to such requests on-chain, they should be slashed and removed from the
      process. The on-chain response must be the shares, encrypted with the
      recipient's public key.
      
      **In the initial version**, we'll just pretend everyone responds
      faithfully.
   2. If some node sends shares which don't verify, the recipient posts a
      complaint to the contract, along with merkle commitments to the two halves
      of calculation it's done during the verification process. The defendant
      node must respond within a few blocks, stating which half differs. The
      process repeats recursively, until the leaves of the calculation are
      reached. If they involve valid multiples of the coefficient commitments
      for that part of the calculation the defendant wins and the plaintiff is
      slashed, otherwise vice versa.
      
      Instead of posting a complaint on the contract, valid complaints could be
      included as part of the message to be signed. Then complex on-chain
      verification is unnecessary.
      
      **In the initial version,** we'll just broadcast the complaint amongst the
      group over the network, as usual, and nodes will be responsible for
      verifying the complaint and ejecting the bad actor themselves (already
      implemented). In a later version, nodes can vote on-chain on which side of
      the complaint is correct, and the loser is slashed if the vote exceeds the
      signature threshold. In a still later version, we can do the full
      verification on-chain. We won't be able to do that for very large
      signature groups, though, so we'll need the challenge/response protocol
      described above.
5. After some timeout, if there's no clique of nodes who have reported receiving
   correct shares from each other, key generation halts. 
6. Any node *j* which receives all shares *fᵢ*(*j*) is able to construct the
   distributed public key using the equations in section 2.2 of Stinson &
   Strobl.
7. If this is a distributed ephemeral key for a signature with a pre-existing
   distributed persistent key constructed by the participants, they follow the
   signature-issuing protocol described in section 4.2 of Stinson & Strobl. At
   this stage, we have completed step 1. of that protocol.
8. Each node provides the signature with their term of the secret distributed
   key, as in step 2. of the signature Stinson & Strobl protocol.
   1. If some node fails to provide their signature, they can be ignored.
   2. If some node provides a signature which fails to verify according to
      equation (3) in step 3. of the Stinson & Strobl protocol, their
      contribution is ignored. 
      
      At this point, we can't add more information to the message to be signed.
      So on-chain validation of the node misbehavior has to be verified
      cryptographically or by voting (or by construction of yet another
      signature.) **In the initial version,** we'll just ignore such messages.
   3. Once a node has received partial signatures from sufficiently many
      participants, it can construct the signature as described in step 4 of the
      S & S protocol.,

# Initial harness

It would be best to run the participants at least in docker containers. Might as
well build it into chainlink, I guess. That means that there needs to be a
database representation of the node data.

But that seems premature. The current task is to simply "Extend ethdss_test
tests to communicate over libp2p." I could even do that in the current test.
What would that show, though? In that case, I would do the whole thing as a
bunch of go routines, I guess. Is there any way to make that part of an
incremental development?

An advantage of starting with a go-routine-based version would be that it's easy
to benchmark how big / expensive the calculations will get.

Another issue: I have moved away from the Rabin DKG, and now want to do the
Pedersen version, because it's much more economical in terms of bandwidth.
Should I implement DSS using the Pedersen DKG, first??

Let's build this into chainlink, but with a very primitive interface to start
with. Takes a set of files with the relevant data, and outputs the public key
and reconstructed private key. Each node running in a docker container. Just a
straight Pedersen DKG, to start with.

Should the nodes be running the DHT all the time, in a separate process?
Probably? (I discussed this with Steve, and he agrees that this can be part of
the standard chainlink startup process.)

## libp2p

Probably each participant should be running a libp2p service as a matter of
course. It should contact a boostrap node to particpate in a DHT.

The services it should run are

- DHT lookups. The dht package should do this for me.
- Request for public key data. Should check that the requester is one of the
  participants. Probably need some kind of ID for each key generation which is
  in progress, too. Later, can announce availability of others' key data, and
  take requests for it.
- Request for secret share. Should check that the requester matches an address
  in the peer list, and return the share for that index.
- Complaint about a bad share.
- Complaint about a complaint.

In addition, I need a bootstrap node for the DHT. This should probably go in the
integration directory.

## Key material

All this key material is going to need to be persisted. It should be encrypted
with the user's username/password, or something similar. Maybe the private key
for the node? It can probably safely encrypt with the [elgamal package](
https://godoc.org/golang.org/x/crypto/openpgp/elgamal).


## Stories

Going to develop this "outside in", to try to avoid the pain we had with service
agreements.

1. Bootstrap node. Starts up a DHT service, can record the network location (ip
   address?) associated with a public key.

   `docker-compose` file with bootstrap node. This will require a separate go
   project, so I can have a separate `main` package. This is going in threshold,
   in the root directory.
   
2. chainlink code which kicks off a go routine for the dht and other services,
   and adds the bootstrap node to report its location.
   
3. Add five chainlink nodes from story 2 to the integration test from story 1.
   Nodes find each other through the DHT.

4. Nodes request public-key coefficients from each other, for a three-of-five
   threshold key.

5. Nodes compare public-key coefficients. A node which fails to send
   coefficients, or sends inconsistent coefficients, is dropped from the protocol.
   
6. Nodes request secret shares from each other. A node which fails to send
   share, or which sends a bad share, is dropped from the protocol.
   
7. Nodes construct the same distributed public key from their shares.

8. Three nodes reconstruct the private key from their shares.
