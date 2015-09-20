[BitcoinMiner](https://github.com/bitcoin/bitcoin/blob/master/src/miner.cpp) (code from core Bitcoin implementation)

```c
//
// ScanHash scans nonces looking for a hash with at least some zero bits.
// The nonce is usually preserved between calls, but periodically or if the
// nonce is 0xffff0000 or above, the block is rebuilt and nNonce starts over at
// zero.
//
bool static ScanHash(const CBlockHeader *pblock, uint32_t& nNonce, uint256 *phash)
{
    // Write the first 76 bytes of the block header to a double-SHA256 state.
    CHash256 hasher;
    CDataStream ss(SER_NETWORK, PROTOCOL_VERSION);
    ss << *pblock;
    assert(ss.size() == 80);
    hasher.Write((unsigned char*)&ss[0], 76);

    while (true) {
        nNonce++;

        // Write the last 4 bytes of the block header (the nonce) to a copy of
        // the double-SHA256 state, and compute the result.
        CHash256(hasher).Write((unsigned char*)&nNonce, 4).Finalize((unsigned char*)phash);

        // Return the nonce if the hash has at least some zero bits,
        // caller will check if it has enough to reach the target
        if (((uint16_t*)phash)[15] == 0)
            return true;

        // If nothing found after trying for a while, return -1
        if ((nNonce & 0xfff) == 0)
            return false;
    }
}
```
[BitcoinMiner](https://github.com/bitcoin/bitcoin/blob/master/src/miner.cpp#L438): (excerpted, most error checking code removed)

```c
void static BitcoinMiner(CWallet *pwallet)
{
    SetThreadPriority(THREAD_PRIORITY_LOWEST);
    CReserveKey reservekey(pwallet);
    unsigned int nExtraNonce = 0;

    try {
        while (true) {
            // Create new block
            unsigned int nTransactionsUpdatedLast = mempool.GetTransactionsUpdated();
            CBlockIndex* pindexPrev = chainActive.Tip();

            auto_ptr<CBlockTemplate> pblocktemplate(CreateNewBlockWithKey(reservekey));
            CBlock *pblock = &pblocktemplate->block;
            IncrementExtraNonce(pblock, pindexPrev, nExtraNonce);

            int64_t nStart = GetTime();
            arith_uint256 hashTarget = arith_uint256().SetCompact(pblock->nBits);
            uint256 hash;
            uint32_t nNonce = 0;
            while (true) {
                // Check if something found
                if (ScanHash(pblock, nNonce, &hash))
                {
                    if (UintToArith256(hash) <= hashTarget)
                    {
                        // Found a solution
                        pblock->nNonce = nNonce;
                        assert(hash == pblock->GetHash());

                        SetThreadPriority(THREAD_PRIORITY_NORMAL);
                        LogPrintf("proof-of-work found  \n  hash: %s  \ntarget: %s\n", 
	                          hash.GetHex(), hashTarget.GetHex());
                        ProcessBlockFound(pblock, *pwallet, reservekey);
                        SetThreadPriority(THREAD_PRIORITY_LOWEST);
                        break;
                    }
                }

                if (nNonce >= 0xffff0000) break;
                // ... other breaking conditions elided
                // Update nTime every few seconds
                UpdateTime(pblock, pindexPrev);
            }
        }
    }
    catch (const boost::thread_interrupted&)
    {
        LogPrintf("BitcoinMiner terminated\n");
        throw;
    }
}
```
