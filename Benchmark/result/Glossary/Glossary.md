##  快速术语表

这个快速术语表包含许多CryptoYC Benchmark测试报告中相关的术语。这些术语在报告都有使用，所以请将其加入书签以便快速参考。

- 区块 Block

区块是关于所包含的交易的所需信息（区块头）的集合。
		
- 区块链 Blockchain

由工作证明系统验证的一系列区块，每个区块都连接到它的前任，一直到创世区块。
	
- 女巫攻击 Sybil attack
  

作用于对等（Peer-to-Peer,简称P2P）网络中的一种攻击形式。攻击者利用单个节点来伪造多个身份存在于P2P网络中，从而达到削弱网络的冗余性，降低网络健壮性，监视或干扰网络正常活动等目的。

- BGP劫持攻击 BGP hijacking
  

bgp劫持是指攻击者恶意改变互联网流量的路由。攻击者通过错误地宣布他们实际上并不拥有、控制或路由的IP地址组(称为IP前缀)的所有权来实现这一点。bgp劫持很像是有人在高速公路上改变所有的标志，将汽车指向错误的出口。
    
- DDos攻击 DDoS attack

分布式拒绝服务(DDoS:Distributed Denial of Service)攻击指借助于客户/服务器技术，将多个计算机联合起来作为攻击平台，对一个或多个目标发动DDoS攻击，从而成倍地提高拒绝服务攻击的威力。
    
- 网络分裂攻击 Network Split Attack
  

把网络分成两个或多个部分，使得在较小的链上进行的任何重复交易都将丢失。
    
- 长程攻击 Long term attack
  

攻击者创建了一条从创世区块开始的长区块链分支，并试图替换掉当前的合法主链。

- 短程攻击 Short term attack
  

攻击者篡改最近几个区块的数据。
    
- TPS

每秒钟处理交易/事务的数量。
    
- 编译 Compiling

将高级编程语言（例如 Solidity）编写的代码转换为低级语言（例如 EVM 字节码）
	
- 共识 Consensus
  

大量节点，通常是网络上的大多数节点，在其本地验证的最佳区块链中都有相同的区块的情况。不要与共识规则混淆。
    
- 共识规则 Consensus rules
	

完整节点为了与其他节点保持一致，遵循的区块验证规则。不要与共识混淆。
	
- 去中心化自治组织 DAO
	

去中心化自治组织 Decentralised Autonomous Organization. 没有层级管理的公司和其他组织。也可能是指2016年4月30日发布的名为“The DAO”的合约，该合约于2016年6月遭到黑客攻击，最终在第1,192,000个区块激起了硬分叉（代号 DAO），恢复了被攻击的 DAO 合约，并导致了以太坊和以太坊经典两个竞争系统。
	
- 去中心化应用 DApp
  

去中心化应用 Decentralised Application. 狭义上，它至少是智能合约和 web 用户界面。更广泛地说，DApp 是一个基于开放式，分散式，点对点基础架构服务的 Web 应用程序。另外，许多 DApp 包括去中心化存储和/或消息协议和平台。
    
- 难度 Difficulty
  	

网络范围的设置，控制产生工作量证明需要多少计算。

- 数字签名 Digital signature
	

数字签名算法是一个过程，用户可以使用私钥为文档生成称为“签名”的短字符串数据，以便具有签名，文档，和相应公钥的任何人，都可以验证（1 ）该文件由该特定私钥的所有者“签名”，以及（2）该文件在签署后未被更改。
	
- 哈希值 Hash
   	

通过哈希方法为可变大小的数据生成的固定长度的指纹。

- 网络 Network
  

将交易和区块传播到每个节点（网络参与者）的对等网络。
    
- 重入攻击 Re-entrancy Attack
	

当攻击者合约（Attacker contracts）调用受害者合约（Victim contracts）的方法时，可以重复这种攻击。让我们称它为victim.withdraw()，在对该合约函数的原始调用完成之前，再次调用victim.withdraw()方法，持续递归调用它自己。递归调用可以通过攻击者合约的后备方法实现。攻击者必须执行的唯一技巧是在用完燃气之前中断递归调用，并避免盗用的以太被还原。
	
- 中本聪 Satoshi Nakamoto
  

Satoshi Nakamoto 是设计比特币及其原始实现 Bitcoin Core 的个人或团队的名字。作为实现的一部分，他们也设计了第一个区块链。在这个过程中，他们是第一个解决数字货币的双重支付问题的。他们的真实身份至今仍是个谜。
    
- Vitalik Buterin
  

Vitalik Buterin 是俄国-加拿大的程序员和作家，以太坊和 Bitcoin 杂志的联合创始人。
    
- Gavin Wood
  

Gavin Wood 是英国的程序员，以太坊的联合创始人和前 CTO。在2014年8月他提出了Solidity，用于编写智能合约的面向合约的编程语言。
    
- SHA
  

安全哈希算法 Secure Hash Algorithm，SHA 是美国国家标准与技术研究院（NIST）发布的一系列加密哈希函数。

- 智能合约 Smart Contract

在以太坊的计算框架上执行的程序。

- Solidity
	

过程式（命令式）编程语言，语法类似于 Javascript, C++ 或 Java。以太坊智能合约最流行和最常使用的语言。由 Gavin Wood（本书的联合作者）首先创造

- 测试网 Testnet
	

一个测试网络（简称 testnet），用于模拟主网主要网络的行为。

- 图灵完备 Turing Complete
	

在计算理论中，如果数据操纵规则（如计算机的指令集，程序设计语言或细胞自动机）可用于模拟任何图灵机，则它被称为图灵完备或计算上通用的。这个概念是以英国数学家和计算机科学家阿兰图灵命名的。
	
- 节点 Node

参与到对等网络的软件客户端。

- 钱包 Wallet
	

拥有你的所有密钥的软件。 
