# Intended full workflow

Since in the end we can't trust the network to deliver messages for us, a key
idea for the final version of the workflow is to use a smart contract on the
blockchain itself as a coordination mechanism, because it's explicitly designed
for censorship resistance. To keep things cheap on the happy path, this trick is
only used when failures are observed.

While this is an ordered list, many of the steps are independent, especially
between nodes. It is not necessary to synchronize the nodes at every step,
unless that's explicitly indicated.

0. **In the initial version**, we'll just ignore the smart-contract coordinator.
   In later versions the process kicks off with an event from a smart contract.
   The contract will be an early part of the development, though, if only to
   store the participants which are still in good standing so that the nodes can
   look up who to talk to.
1. The index assigned to each node is the hash of its public key, interpreted as
   a scalar mod N (the secp256k1 group order. Keep hashing the hash until it
   corresponds to something less than N, which happens about 99.9999999% of the
   time.)
2. Nodes contact each other through a DHT, looking up the other participants'
   host info via their public keys. Chainlink will run a bootstrap node for
   this DHT.
3. <a id="coefficient-request"/>Every node requests the coefficient commitments
   (*Cᵢₖ*, in section 2.4, step 1 of [Stinson and Strobl 2001
   ](https://www.researchgate.net/profile/Willy_Susilo/publication/242499559_Information_Security_and_Privacy_13th_Australasian_Conference_ACISP_2008_Wollongong_Australia_July_7-9_2008_Proceedings/links/00b495314f3bcaaa46000000.pdf#page=426))
   of every other node.
   1. In the happy path, the nodes comply with these requests, sending the same
      commitments to every other node. In the final version of the protocol,
      these should be preceded by a signed merkle-tree commitment to the full
      set of coefficients, and the commitments should be sent as signed batches
      of merkle subtrees, along with the merkle path to the initial commitment.
      **In the initial version** assuming the happy path, the commitments can
      just be sent over as a block. 
   2. Failure 1: A node sends inconsistent commitments to other nodes. In the
      final version the nodes publish the initial commitments they've received
      from every other node, and every node compares them. Any inconsistencies
      are reported to a smart contract where the offending node has a stake, and
      it is slashed (this is the reason the commitments should be sent in small
      batches with independent signatures, so that on-chain verification is not
      too expensive.) The node is removed from further participation in the DKG.
      **In the first initial version**, the offending node is just removed
      without being slashed. That much is already implemented.
   3. Failure 2: A node fails to send some portion of its commitments. Handling
      this is more complex, because it cannot be cryptographically verified.
      After a certain number of blocks, the expectant recipient node reports the
      failure on a smart contract. Other nodes can report to the contract that
      they have the data, along with their host address, and the expectant node
      can then ask them for it. This allows the group to route around partial
      network failures.
      
      If some node offers missing data, and then doesn't provide it, the
      reporting procedure repeats with a complaint about missing data from that
      node. To prevent an arbitrary regress by hostile nodes, this process can
      only cycle a limited number of times (say, three.)
      
      The first time a threshold of nodes report some node as having failed to
      provide its commitments, probably no penalty should be levied. But
      repeated failure should probably be penalized heavily, and at the least
      the node should be removed from the group for all future key generation.
      
      **In the initial version**, we'll just pretend this can't happen. In a
      later version, we'll just drop the failing node from the group, and
      exclude them for the lifetime of the key. (They can come back to later
      rounds if they failed during construction of an ephemeral key.) 
4. All nodes ask all other nodes for their secret shares (*f*(*u*) and
   *f'*(*u*), in step 1 of section 2.4 in Stinson & Strobl.) Once they've
   verified their shares from a particular node, they announce to the network.
   The response with the shares must include a merkle commitment to the
   calculation verifying the shares. (These are just sums of scalar multiples of
   the transmitting node's commitments, so they can be arranged in nice binary
   trees using associativity.) Nodes broadcast to the group who they've received
   correct shares from.
   1. If some node doesn't respond to some other node with its secret shares by
      a certain number of blocks, the recipient can post a request on the
      coordinating contract. If the mandated sender repeated fails to respond
      on-chain, they should be slashed and removed from the process. The
      on-chain response must be the shares, encrypted with the recipient's
      public key.
      
      **In the initial version**, we'll just pretend everyone responds
      faithfully.
   2. If some node sends shares which don't verify, the recipient posts a
      complaint to the contract, along merkle commitments to the two halves of
      calculation it's done during the verification process. The defendant node
      must respond within a few blocks, stating which half differs. The process
      repeats recursively, until the leaves of the calculation are reached. If
      they involve valid multiples of the coefficient commitments for that part
      of the calculation the defendant wins and the plaintiff is slashed,
      otherwise vice versa.
      
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
   correct shares from each other, key generation halts. Otherwise, the maximal
   clique is used from here on, and the other participants are thrown out for
   the life of the key.
6. The remaining members all share the public keys of their secret coefficients
   (the *Aᵢₖ*=*aᵢₖG*, in step 3 of section 2.4 of Stinson and Strobl), using
   much the same protocol as in the [above three steps](#coefficient-request)
   for the coefficient commitments.
7. After some timeout, if any node has failed to fully and correctly report its
   coefficient's public keys, its secret coefficients are reconstructed by the
   remaining nodes. The failure is noted on the contract, and the failing node
   is slashed and excluded from further participation. In order for the
   reconstruction to occur, every node but the failing one broadcasts the secret
   share they received from the failure. If any node fails to receive one of
   these shares, the
        
