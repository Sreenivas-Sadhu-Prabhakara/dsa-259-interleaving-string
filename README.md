# Interleaving String

> NeetCode 150 · 2-D DP — dp[i][j]: can the first i of s1 and first j of s2 weave the first i+j of s3? The last character came from one of exactly two places. Go, stdlib only, tested.

## Problem

Is s3 formed by interleaving s1 and s2 — both fully used, each keeping its internal order? Greedy matching fails when both strings offer the same next character; the choice needs full DP.

## Approach

The target position is forced (i+j characters consumed means position i+j−1), so the state is just (i, j), and s3's character there must have come from s1's i-th or s2's j-th — two predecessors: `dp[i−1][j] && s1 matches`, or `dp[i][j−1] && s2 matches`. The length pre-check (`n+m == len(s3)`) isn't an optimization, it's what makes the indexing sound. Row i depends only on row i−1, so one rolling boolean row over s2's length suffices, with `dp[0]` (the j=0 column) updated per row. Structurally the edit-distance grid with OR instead of min.

**Complexity:** O(n·m) time, O(m) space.

## Run

    go test ./...   # verify
    go run .        # demo

---

Part of a 275-repo DSA series in Go — Striver's SDE Sheet (dsa-001–191) plus the NeetCode 150 / Blind 75 additions (dsa-192–275), stdlib only, every solution tested. Browse: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
