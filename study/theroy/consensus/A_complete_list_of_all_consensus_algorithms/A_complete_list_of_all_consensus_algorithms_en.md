## Six blockchain consensus

### 1. Proof of Work (PoW, Proof of Work)

**definition:**

In PoW, blockchain participants (called “miners”) need to solve a “complex but useless” computational problem by adding a transaction to the blockchain.

**application:**

Bitcoin, Ethereum, Litecoin, etc.

**advantage:**

- Byzantine fault tolerance.
- Prevent DoS attacks.
- Decentralization.

**Disadvantages:**

- Great energy consumption.
- Slow speed.
- There is a 51% risk of attack.

### 2. Proof of Stake (PoS, Proof of Stake)

**definition:**

The node needs to deposit a certain amount of tokens as equity in the network. The share size of the equity determines the probability of being selected as the verifier, so that the next block can be created.

**application:**

Dotcoin, Cosmos, etc.

**advantage:**

- Faster transaction processing than PoW.
- More energy efficient than PoW.
- Byzantine fault tolerance.
- Prevent DoS attacks.
- Expandability.
  

**Disadvantages:**

- May develop in a centralized direction

### 3. Certificate of Authorization (DPoS, Delegated Proof-of-Stake)

**definition:**

In the DPoS system, equity holders can elect leaders (or witnesses, Witness). These leaders can vote by authorization of the equity holder.

**application:**

EOS, BitShares, etc.

**advantage:**

- Faster transaction processing than PoW.
- Decentralization in a democratic way.
- Expandability.

**Disadvantages:**

- May develop in a centralized direction.
- May cause the Matthew effect.

### 4. Proof of Elapsed Time (PoET, Proof of Elapsed Time)

**definition:**

PoET is often used in licensed blockchain networks to determine the mining rights of blockers in the network. The licensed blockchain network requires any prospective participant to verify identity before joining. According to the principle of the fair lottery system, each node has the same chance of becoming the winner. The PoET mechanism gives a large number of possible network participants the opportunity to win on an equal footing.

**application:**

HyperLedger Sawtooth

**advantage:**

- Faster transaction processing than PoW.
- Participation is low cost. More people can easily join to achieve decentralization.
- For all participants, it is easier to verify that the leader is legally elected.
- The cost of controlling the leader's electoral process is directly proportional to the value gained from it.

**Disadvantages:**

- Although the cost of PoET is low, specific hardware must be used. Therefore, it will not be adopted on a large scale.
- Not applicable to public blockchains.

### 5. Proof of Capacity (PoC, Proof of Capacity)

**definition:**

By assigning a certain amount of memory or disk space to address the challenges provided by the service provider, it shows that someone has a legitimate interest in a service, such as sending a message.

**application:**

Burst

**advantage:**

- Faster transaction processing than PoW.
- More energy efficient than PoW.
- Byzantine fault tolerance.
- Prevent DoS attacks.
- Create blocks using memory or disk space instead of computing power.

**Disadvantages:**

- Scalability is a problem

### 6. Proof-of-Authority (PoA, Proof-of-Authority)

**definition:**

PoA, each individual has the right to become a verifier, so there is an incentive to maintain the verifier's location once acquired. By attaching a reputation to the identity, the certifier can be encouraged to maintain the transaction. Because the verifier does not want to get a negative reputation, it will make it lose the hard-won certifier status.

**application:**

VeChain

**advantage:**

- Higher transaction processing per second.
- Scalability.
- Trustworthy nodes.

**Disadvantages:**

- Not a true decentralized system.
- The node cannot be anonymous.

## references

[1]<https://www.priceflier.com/tech-blogs/blockchain/blockchain-consensus-algorithms-application-advantage-disadvantage>

[2]<https://www.infoq.cn/article/consensuspedia-an-encyclopedia-of-29-consensus-algorithms>
