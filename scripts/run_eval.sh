#!/bin/bash
set -e

# Build the binary
task build > /dev/null

# Clean up any previous test collection
./mnemosyne forget eval_test 2>/dev/null || true

# Initialize a fresh collection
./mnemosyne init --name eval_test > /dev/null

echo "Inserting test documents..."
# Category A: Asymmetrical Relations
./mnemosyne add --name eval_test "The brave cat chased the cowardly dog out of the yard." > /dev/null
./mnemosyne add --name eval_test "The aggressive dog chased the frightened cat up the tree." > /dev/null

# Category B: Keyword Traps vs Semantic Context
./mnemosyne add --name eval_test "To terminate an unresponsive application in Ubuntu, you can use the kill command followed by the PID." > /dev/null
./mnemosyne add --name eval_test "The hitman was hired to kill the target, terminating the contract cleanly." > /dev/null

# Category C: Negation and Modifiers
./mnemosyne add --name eval_test "This framework is perfect for Python developers, but does not support Go." > /dev/null
./mnemosyne add --name eval_test "A high-performance backend framework built natively in Go." > /dev/null

# Category D: Complex Queries (Nuance/Intent)
./mnemosyne add --name eval_test "Apple is expected to release the new iPhone model next September with an upgraded titanium chassis." > /dev/null
./mnemosyne add --name eval_test "The farmer harvested a fresh batch of crisp red apples in late September." > /dev/null

# Category E: Lexical Overlap but Different Meaning
./mnemosyne add --name eval_test "I need to open a bank account to deposit my savings." > /dev/null
./mnemosyne add --name eval_test "We sat on the river bank and watched the water flow by." > /dev/null

echo ""
echo "=========================================================="
echo "TEST 1: Asymmetrical Relations (Who did what?)"
echo "Query: 'a dog chasing a cat'"
echo "=========================================================="
echo "--- WITHOUT Reranker (FTS + Vector RRF) ---"
./mnemosyne search --name eval_test --no-rerank "a dog chasing a cat" --limit 2 --debug
echo ""
echo "--- WITH Reranker (Cross-Encoder) ---"
./mnemosyne search --name eval_test "a dog chasing a cat" --limit 2 --debug

echo ""
echo "=========================================================="
echo "TEST 2: Keyword Traps vs Context"
echo "Query: 'how to kill a linux process'"
echo "=========================================================="
echo "--- WITHOUT Reranker (FTS + Vector RRF) ---"
./mnemosyne search --name eval_test --no-rerank "how to kill a linux process" --limit 2 --debug
echo ""
echo "--- WITH Reranker (Cross-Encoder) ---"
./mnemosyne search --name eval_test "how to kill a linux process" --limit 2 --debug

echo ""
echo "=========================================================="
echo "TEST 3: Negation"
echo "Query: 'framework for Go'"
echo "=========================================================="
echo "--- WITHOUT Reranker (FTS + Vector RRF) ---"
./mnemosyne search --name eval_test --no-rerank "framework for Go" --limit 2 --debug
echo ""
echo "--- WITH Reranker (Cross-Encoder) ---"
./mnemosyne search --name eval_test "framework for Go" --limit 2 --debug

echo ""
echo "=========================================================="
echo "TEST 4: Nuance/Intent (Tech vs Agriculture)"
echo "Query: 'when is the new apple coming out'"
echo "=========================================================="
echo "--- WITHOUT Reranker (FTS + Vector RRF) ---"
./mnemosyne search --name eval_test --no-rerank "when is the new apple coming out" --limit 2 --debug
echo ""
echo "--- WITH Reranker (Cross-Encoder) ---"
./mnemosyne search --name eval_test "when is the new apple coming out" --limit 2 --debug

echo ""
echo "=========================================================="
echo "TEST 5: Lexical Overlap (Polysemy)"
echo "Query: 'sitting by the water'"
echo "=========================================================="
echo "--- WITHOUT Reranker (FTS + Vector RRF) ---"
./mnemosyne search --name eval_test --no-rerank "sitting by the water" --limit 2 --debug
echo ""
echo "--- WITH Reranker (Cross-Encoder) ---"
./mnemosyne search --name eval_test "sitting by the water" --limit 2 --debug

