# Block Composition And Formation Logic

The activation of Blueberry fork slightly changes the way the P-chain selects transactions to be included in next block and deals with block timestamps. In this brief document we detail the process and the changes.

## Apricot

### Apricot Block Content

Apricot allows the following block types with the following content:

* *Standard Blocks* may contain multiple transactions of the following types:
  * CreateChainTx
  * CreateSubnetTx,
  * ImportTx,
  * ExportTx.
* *Proposal Blocks* may contain a single transaction of the following types:
  * AddValidatorTx,
  * AddDelegatorTx,
  * AddSubnetValidatorTx,
  * RewardValidatorTx,
  * AdvanceTimeTx.
* *Options blocks*, i.e. *Commit Block* and *Abort Block* do not contain any transactions.

Note that *Atomic blocks* are deprecated as of Apricot phase 5 fork. They used to contain ImportTx and ExportTx which are now included into Standard Blocks.

Each block has a header containing:

* ParentID
* Height

Note that Apricot block headers do not contain any block timestamp.

### Apricot Block Formation Logic

Transactions included in an Apricot block can originate from the mempool or can be created just in time to duly update the staker set. Block formation logic in the Apricot fork can be broken up into two high-level steps:

* First, we try selecting any candidate decision or proposal transactions which could be included in a block *without advancing the current chain time*;
* If no such transactions are found, we evaluate candidate transactions which *may require advancing chain time*. If a chain time advancement is required to include these transactions in a block, a proposal block with an advance time transaction is built first; selected transactions may be included in a subsequent block.

In more detail, blocks which do not change chain time are built as follows:

1. If mempool contains any decision transactions, a Standard block is issued with all of the transactions the default block size can accommodate. Note that Apricot standard blocks do not change the current chain time.
2. If the current chain time matches the end of any staker's staking time, a reward transaction is issued into a Proposal Block to initiate network voting on whether the specified staker should be rewarded. Note that there could be multiple stakers ending their staking period at the same chain time, hence a Proposal block must be issued for all of them before chain time is moved ahead. Any attempt to move chain time ahead before rewarding all stakers would fail block verification.

While the above steps could be executed in any order, we pick decisions transactions first to maximize throughput.

Once all possibilities of create a block not advancing chain time are exhausted, we attempt to build a block which *may* advance chain time as follows:

1. If the local clock's time is greater than or equal to the earliest staker set change event timestamp, an advance time transaction is issued into a Proposal Block to move current chain time to the earliest staker set change timestamp. Upon this Proposal block's acceptance, chain time will be move ahead and all scheduled changes (e.g. promoting a staker from pending to current) will be carried out.
2. If the mempool contains any proposal transactions, the mempool proposal transaction with the earliest start time is selected and included into a Proposal Block[^1]. A mempool proposal transaction as is won't change the current chain time[^2]. However there is an edge case to consider: on low activity chains (e.g. Fuji P-chain) chain time may fall significantly behind the local clock. If a proposal transaction is finally issued, its start time is likely to be quite far in the future relative to the current chain time. This would cause the proposal transaction to be considered invalid and rejected, since a staker added by a proposal transaction's start time must be at most 366 hours (two weeks) after current chain time. To avoid this edge case on low-activity chains, an advance time transaction is issued first to move chain time to the local clock's time. As soon as chain time is advanced, the mempool proposal transaction will be issued and accepted.

Note that the order in which these steps are executed matters. A block updating chain time would be deemed invalid if it would advance time beyond the staker set's next change event, skipping the associated changes. The order above ensures this never happens because if checks first if chain time should be moved to the time of the next staker set change. It can also be verified by inspection that the timestamp selected for the advance time transactions always respect the synchrony bound.

Block formation terminates as soon as any of the steps executed manage to select transactions to be included into a block. An error is raised otherwise to signal that there are not transactions to be issued. Finally a timer is kicked off to schedule the next block formation attempt.

## Blueberry

### Blueberry Block Content

Blueberry allows the following block types with the following content:

* *Standard Blocks* may contain multiple transactions of the following types:
  * CreateChainTx
  * CreateSubnetTx,
  * ImportTx,
  * ExportTx,
  * AddValidatorTx,
  * AddDelegatorTx,
  * AddSubnetValidatorTx.
* *Proposal Blocks* may contain a single transaction of the following types:
  * RewardValidatorTx,
* *Options blocks*, i.e. *Commit Block* and *Abort Block* do not contain any transactions.
  
Note that each block has an header containing:

* ParentID
* Height
* Time

So the two main differences with respect to Apricot are:

* AddValidatorTx, AddDelegatorTx, AddSubnetValidatorTx are included into Standard blocks rather than Proposal blocks so that they don't need to be voted on (i.e. followed by a Commit/Abort block).
* Block timestamp is explicitly serialized into block header, to allow chain time update.

## Blueberry Block Formation Logic

The activation of the Blueberry fork only makes minor changes to the way the P-chain selects transactions to be included in next block, such as block timestamp calculation. In this section we detail the process and the changes.

We carry out operations in the following order:

* We try to move chain time ahead to the current local time or the earliest staker set change event. Unlike Apricot, here we issue either a Standard block or a Proposal block to advance the time.
* We check if any staker needs to be rewarded, issuing as many Proposal blocks as needed, as above.
* We try to fill a Standard block with mempool decision transactions.

[^1]: Proposal transactions whose start time is too close to local time are dropped first and won't be included in any block.
[^2]: Advance time transactions are proposal transactions and they do change chain time. But advance time transactions are generated just in time and never stored in the mempool. Here mempool proposal transactions refer to AddValidator, AddDelegator and AddSubnetValidator transactions. Reward validator transactions are proposal transactions which do not change chain time but which never in mempool (they are generated just in time).