

##  一、 概述

VeChain 是一个与以太坊生态系统高度兼容的通用区块链公链项目。具有明确的治理模型迭代规划，通过选举获权维持生态平衡。
采用一种称为权威证明（PoA）的共识协议，具有类似白名单的社区授权机制。
据白皮书介绍，唯链采用将“使用成本”与“通证的估值”分层的双通证系统来构建其经济模型。



##  二、分析

### (一) 测试说明

我们的在AWS上部署了若干个Kubernetes节点，用以模拟Vchain网络的环境，具体测试环境如下

![Kubernetes Test env](https://github.com/EVOLABTeam/benchmark/blob/master/result/md/asset/Kubernetes%20Test%20env.jpg)

### (二) 共识

VeChainThor的共识机制采用的是PoA（Proof-of-Authority ），即权威证明，而非PoW (Proof-of-Work)或PoS (Proof-of-stake)。唯链基金会授权超级权益节点，通过白名单的管控，来担任/担保整个网路的运行及获取相对应得权利。唯链采用随机出块机制提高安全性，采用了DPRP（确认伪随机过程）来保证出块的随机性。权威共识主要的利益相关者为超级权益节点，主要的超级节点通过投票选举并满足基金会最低要求的条件下产生。
PoA协议特点：
1. 算力要求低
2. 无需超级节点之间通信，可完成共识
3. 系统连续性不受节点数量影响
PoA对比PoW 速度较快，效率更高。社区治理的合理性，潜在的中心化风险，更值得关注。
共识协议对比如下：

![PoA VS PoW](https://github.com/EVOLABTeam/benchmark/blob/master/result/md/asset/PoA.PNG)

VeChain 使用Go实现PoA共识协议，结合语言高并发特性，允许单笔交易采用PoW共识机制，并支持“多任务交易”。超级权益节点不会受到外界因素影响，验证数量也没有最低要求。


### (三) 安全

通过Benchmark公链测试工具，我们对VeChain进行一系列安全测试，包括DDos攻击、网络带宽服务攻击等，我们的测试方法如下:
1. 建立VeChain测试网
2. 发送RPC，让测试网部分节点对其他节点发起攻击
3. 得到测试结果

|     方案     |  结果  |                             备注                             |
| :----------: | :----: | :----------------------------------------------------------: |
|   DDoS攻击   | 不通过 |   以大量的通信量冲击网络，使得所有可用网络资源都被消耗殆尽   |
| 网络分裂攻击 | 不通过 | 把网络分成两个或多个部分，使得在较小的链上进行的任何重复交易都将丢失 |




### (四) 性能

通过Benchmark公链测试工具，对VeChain进行性能测试，我们的测试方法如下：
1. 建立VeChain测试网 
2. 发送RPC，让测试网部分节点发起交易（每秒N笔交易，线性增长）;
3. 节点检测交易同步的时间，直到检测到超过一定时间（一般是出块时间）；

| 方案 | TPS | 备注 |
| :--: | :--: | :--: |
| 理想网络情况 | 1500 | 单机虚拟机网络，无限网络连接 |
| 正常网络情况 | 1000 | 分布全球的100个节点，正常网络连接 |
| 恶劣网络情况 | <1 |  |

### (五) 技术创新

双通证系统，分层设计模型。商业应用层使用VET流通，区块底层消耗使用VTHO流通。
VET与VTHO之间具有直接关联。 VET是唯链雷神生态系统中的“智能货币”或“智能价值”，可在智能合约中编程和执行，从而推进唯链雷神区块链上运营商业活动，完成价值传导。
此外，VET 可被视为在生态系统内各点之间建立联系的关键元素。
VTHO 作为执行转账交易和智能合约交易的能量或费用，由 VET 随时间推移而生成。

### (六) 代码

1. 代码概况

   Stellar的Github仓库的一共有30个公开仓库，主要仓库的具体数据如下:

   |    repositories    | commits | watches | stars | forks | issues |
   | :----------------: | :-----: | :-----: | :---: | :---: | :----: |
   |        thor        |  2658   |   69    |  357  |  129  |   3    |
   |   thor-devkit.js   |   60    |   11    |   9   |  10   |   0    |
   |     connex-env     |   37    |    5    |   2   |   2   |   0    |
   | thor-sync.electron |   879   |   12    |  16   |  16   |   1    |
   | thor-client-sdk4j  |   241   |    5    |  15   |   9   |   1    |

   

2. 代码更新

   ![Vechain_code commit](E:\study\blockchain\benchmark\result\md\asset\Vechain_code commit.png)

3. 代码重复

-  通过Benchmark公链测试工具，对Vechain进行代码相似度检查，因为Vechain的技术栈是Go，选择以go-ethereum作为标准，具体测试方法如下建立代码索引库
- 把thor的源代码放进Elasticsearch
- 把thor的源代码和go-ethereum作比较

### (七) 经济模型

VET产生VTHO，产生的数量，时间均来自对市场表现的评估。超级节点投票促使基金会完成参数修改。
经济动力来自于：对VET的长时间保留将获得VTHO、对VET价值的未来期待、智能合约及交易支付的需求。


## 总结

项目基金会具有中心控制权，超级节点具备投票权，网络中极易发生51%攻击，
即超级节点中心化趋势后可串通修改区块链状态。共识算法基本借鉴了ETH，提出双通证系统以保证“交易成本”与“货币价值”之间的平衡稳定。项目关注重点，在治理模型，希望通过提出一个明确的社区治理制度，维护和驱动项目生态的发展。