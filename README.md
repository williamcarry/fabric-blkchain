# Hyperledger Fabric [![join the chat][rocketchat-image]][rocketchat-url]

Hyperledger是一个旨在推动区块链跨行业应用的开源项目，由Linux基金会在2015年12月主导发起该项目，成员包括金融、银行、物联网、供应链、制造和科技等多个行业的领头羊，托管了众多面向企业的区块链开源框架和工具。



## Hyperledger及Fabric项目概述

Hyperledger Fabric（后文简称Fabric）是其中发展最好的一个企业级区块链平台，最初由Digital Asset和IBM贡献，目前已经应用于沃尔玛的食物溯源链（Foodtrust）和马士基的物流跟踪链（TradeLens）中，代表了当下企业级区块链应用的最高水平。可以认为Fabric是一种联盟链（Consortium Blockchain）平台，它适合构建跨越多个企业边界的去中心化应用。

由于Fabric项目的目标是应用于相对可信的企业联盟环境，因此其设计思路与比特币、以太坊等公链平台有明显的差异。Fabric借鉴了区块链的数据结构，但引入了相当多的身份验证与权限控制机制，以及数据隐私保护机制，以适应企业级应用的要求。同时由于企业联盟环境要比完全开放的公链环境可控，因此Fabric没有强调其共识体系对拜占庭容错的实现，允许使用 非拜占庭容错算法建立共识，从而可以达到相当实用的交易吞吐量。

## Fabric的定位与特点

毫无疑问，Fabric是受到比特币的启发而诞生的，因此它借鉴了比特币、以太坊这些公链中的一些核心特性，例如采用不可篡改的区块链结构来保存数据、采用非对称加密技术来进行身份识别 与认证、支持智能合约等等。

但是Fabric定位于企业级的分布式账本技术（DLT - Distributed Ledger Technology）平台，它的主要目的是为跨越多个企业边界的活动提供不可篡改的分布式记账平台。例如在食物溯源应用中，为了让消费者可以了解到所购买食物是否安全，就必须将从农场到加工商、分销商、 零售商乃至监管机构等各个环节的检验与放行信息记录到区块链上，以保证溯源信息的透明与可信。



因此Fabric是一种联盟链（Consortium Blockchain），它适合在多个企业间实现分布式记账，这一定位使Fabric的实现与以太坊这样的公链有了明显的差异：

## 分布式账本 vs. 区块链

分布式账本是比区块链更加宽泛的概念，可以认为区块链只是分布式账本的一种实现技术，其他的分布式账本实现还包括哈希图等。

## 去中心化 vs. 分布式

Fabric淡化了去中心化（Decentralized），而以分布式（Distributed）代替，这一思路带来了系统设计与实现上的巨大影响。例如，在Fabric中，采用中心化的CA机制来发放证书，只有持有有效证书的节点和用户才可以访问区块链上的账本数据。因此Fabric是许可制/Permissioned的区块链，这与不需要许可/Permissionless的以太坊这样的公链形成了鲜明的对比。

## 拜占庭容错 vs. 崩溃容错

由于采用许可机制，Fabric也淡化了对不可信（Trustless）环境下共识达成的依赖性，而假设联盟链中的企业有可能是值得信赖的，因此并不依赖于工作量证明这样的拜占庭容错算法，虽然Fabric模块化的设计可以支持引入不同的共识算法实现，但目前的产品化方案是Kafka共识，它显然是不能对抗拜占庭错误的，不过对不可信环境支持的淡化处理有利于提高交易的吞吐量，这对于企 业级应用也是有益的。

## 数据隐私保护

在另一方面，Fabric强化了隐私保护能力。例如，Fabric支持在同一套企业网络上建立多个不同的通道/Channel，每一个通道都有自己的区块链和访问控制，彼此互不影响，这有利于复用基础设施，例如不同企业间的销售部门可以建立一个通道来分享市场数据，而这些企业间的 研发部门可以建立另一个通道来分享技术数据。

Fabric并不是唯一的联盟链解决方案，但目前可以说是最复杂的企业联盟链实现，这种复杂性源于设计者对应用场景的假设和推演，以及对Fabric广泛适用性的考量，这是我们在学习过程中需要换位思考的一点。


[rocketchat-url]: https://chat.hyperledger.org/channel/fabric
[rocketchat-image]: https://open.rocket.chat/images/join-chat.svg

[![Build Status](https://dev.azure.com/Hyperledger/Fabric/_apis/build/status/Merge?branchName=master)](https://dev.azure.com/Hyperledger/Fabric/_build/latest?definitionId=51&branchName=master)
[![CII Best Practices](https://bestpractices.coreinfrastructure.org/projects/955/badge)](https://bestpractices.coreinfrastructure.org/projects/955)
[![Go Report Card](https://goreportcard.com/badge/github.com/hyperledger/fabric)](https://goreportcard.com/report/github.com/hyperledger/fabric)
[![GoDoc](https://godoc.org/github.com/hyperledger/fabric?status.svg)](https://godoc.org/github.com/hyperledger/fabric)
[![Documentation Status](https://readthedocs.org/projects/hyperledger-fabric/badge/?version=latest)](http://hyperledger-fabric.readthedocs.io/en/latest)

This project is an _Active_ Hyperledger project. For more information on the history of this project see the [Fabric wiki page](https://wiki.hyperledger.org/display/fabric). Information on what _Active_ entails can be found in
the [Hyperledger Project Lifecycle document](https://wiki.hyperledger.org/display/HYP/Project+Lifecycle).
Hyperledger Fabric is a platform for distributed ledger solutions, underpinned
by a modular architecture delivering high degrees of confidentiality,
resiliency, flexibility and scalability. It is designed to support pluggable
implementations of different components, and accommodate the complexity and
intricacies that exist across the economic ecosystem.

Hyperledger Fabric delivers a uniquely elastic and extensible architecture,
distinguishing it from alternative blockchain solutions. Planning for the
future of enterprise blockchain requires building on top of a fully-vetted,
open source architecture; Hyperledger Fabric is your starting point.

## Releases

Fabric provides a release approximately once every four months with new features
and improvements. Additionally, certain releases are designated as long-term
support (LTS) releases. Important fixes will be backported to the most recent
LTS release, and to the prior LTS release during periods of LTS release overlap.
For more details see the [LTS strategy](https://github.com/hyperledger/fabric-rfcs/blob/master/text/0005-lts-release-strategy.md).

LTS releases:
- [v2.2.x](https://hyperledger-fabric.readthedocs.io/en/release-2.2/whatsnew.html) (current LTS release)
- [v1.4.x](https://hyperledger-fabric.readthedocs.io/en/release-1.4/whatsnew.html) (prior LTS release, maintained through April 2021)

Unless specified otherwise, all releases will be upgradable from the prior minor release.
Additionally, each LTS release is upgradable to the next LTS release.

Fabric releases and release notes can be found on the [GitHub releases page](https://github.com/hyperledger/fabric/releases).

Please visit the [Hyperledger Fabric Jira dashboard](https://jira.hyperledger.org/secure/Dashboard.jspa?selectPageId=10104) for our release roadmap.

Follow the release discussion on the [#fabric-release](https://chat.hyperledger.org/channel/fabric-release) channel in RocketChat.

## Documentation, Getting Started and Developer Guides

Please visit our
online documentation for
information on getting started using and developing with the fabric, SDK and chaincode:
- [v2.3](http://hyperledger-fabric.readthedocs.io/en/release-2.3/)
- [v2.2](http://hyperledger-fabric.readthedocs.io/en/release-2.2/)
- [v2.1](http://hyperledger-fabric.readthedocs.io/en/release-2.1/)
- [v2.0](http://hyperledger-fabric.readthedocs.io/en/release-2.0/)
- [v1.4](http://hyperledger-fabric.readthedocs.io/en/release-1.4/)
- [v1.3](http://hyperledger-fabric.readthedocs.io/en/release-1.3/)
- [v1.2](http://hyperledger-fabric.readthedocs.io/en/release-1.2/)
- [v1.1](http://hyperledger-fabric.readthedocs.io/en/release-1.1/)
- [master branch (development)](http://hyperledger-fabric.readthedocs.io/en/latest/)

It's recommended for first-time users to begin by going through the Getting Started section of the documentation in order to gain familiarity with the Hyperledger Fabric components and the basic transaction flow.

## Contributing

We welcome contributions to the Hyperledger Fabric project in many forms.
There’s always plenty to do! Check [the documentation on how to contribute to this project](http://hyperledger-fabric.readthedocs.io/en/latest/CONTRIBUTING.html)
for the full details.

## Community

[Hyperledger Community](https://www.hyperledger.org/community)

[Hyperledger mailing lists and archives](http://lists.hyperledger.org/)

[Hyperledger Chat](http://chat.hyperledger.org/channel/fabric)

[Hyperledger Fabric Issue Tracking (JIRA)](https://jira.hyperledger.org/secure/Dashboard.jspa?selectPageId=10104)

[Hyperledger Fabric Wiki](https://wiki.hyperledger.org/display/Fabric)

[Hyperledger Wiki](https://wiki.hyperledger.org/)

[Hyperledger Code of Conduct](https://wiki.hyperledger.org/display/HYP/Hyperledger+Code+of+Conduct)

[Community Calendar](https://wiki.hyperledger.org/display/HYP/Calendar+of+Public+Meetings)

## License <a name="license"></a>

Hyperledger Project source code files are made available under the Apache License, Version 2.0 (Apache-2.0), located in the [LICENSE](LICENSE) file. Hyperledger Project documentation files are made available under the Creative Commons Attribution 4.0 International License (CC-BY-4.0), available at http://creativecommons.org/licenses/by/4.0/.
