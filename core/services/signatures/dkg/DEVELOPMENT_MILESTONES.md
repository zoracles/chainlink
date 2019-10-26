1. Nodes have a hard-coded list of participants, and they all find each other
   using a DHT with a bootstrap node.
2. Nodes each create a `dkg.NewDistKeyGenerator`, and send each other their
   secret shares
3. Nodes broadcast 
