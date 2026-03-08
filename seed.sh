#!/usr/bin/env bash
# seed.sh - Initialize a collection and populate it with sample documents for testing search.
# Usage: ./seed.sh [collection-name]
#        Defaults to "test-collection" if no name is provided.

set -euo pipefail

BINARY="./mnemosyne"

if [[ ! -x "$BINARY" ]]; then
  echo "Binary not found. Building..."
  task build
fi

echo "=== Initializing collection ==="
$BINARY init

echo ""
echo "=== Adding documents ==="

docs=(
  # Programming languages
  "Go is a statically typed, compiled programming language designed at Google. It is syntactically similar to C but with memory safety, garbage collection, and structural typing."
  "Rust is a multi-paradigm systems programming language focused on safety, especially safe concurrency. It is syntactically similar to C++ but guarantees memory safety without a garbage collector."
  "Python is a high-level, general-purpose programming language. Its design philosophy emphasizes code readability with the use of significant indentation."
  "JavaScript is a high-level, often just-in-time compiled language that conforms to the ECMAScript specification. It has dynamic typing, prototype-based object orientation, and first-class functions."
  "TypeScript is a strongly typed programming language that builds on JavaScript, giving you better tooling at any scale. It adds optional static typing and class-based object-oriented programming."

  # Concurrency
  "Goroutines are lightweight threads managed by the Go runtime. They enable concurrent programming in Go and communicate through channels, following the principle of sharing memory by communicating."
  "Channels in Go provide a way for goroutines to communicate and synchronize execution. A channel is a typed conduit through which you can send and receive values."
  "The async/await pattern in JavaScript provides a cleaner syntax for working with promises. An async function always returns a promise, and await pauses execution until the promise resolves."
  "Rust uses an ownership model for concurrency safety. The Send and Sync traits determine which types can be transferred or shared across thread boundaries."

  # Databases
  "SQLite is a C-language library that implements a small, fast, self-contained SQL database engine. It is the most widely deployed database engine in the world."
  "PostgreSQL is a powerful, open-source object-relational database system with over 35 years of active development. It supports both SQL and JSON querying."
  "Redis is an open-source, in-memory data structure store used as a database, cache, message broker, and streaming engine. It supports strings, hashes, lists, sets, and sorted sets."
  "MongoDB is a source-available cross-platform document-oriented database program. It uses JSON-like documents with optional schemas and is classified as a NoSQL database."

  # Search and retrieval
  "Full-text search uses techniques like tokenization, stemming, and inverted indexes to find documents matching a query. BM25 is a popular ranking function used in information retrieval."
  "Vector similarity search finds documents by comparing their vector embeddings using distance metrics like cosine similarity or Euclidean distance. It enables semantic search beyond keyword matching."
  "Reciprocal Rank Fusion combines multiple ranked lists into a single ranking. The RRF score for a document is the sum of 1/(k + rank) across all rankings where it appears."
  "Cross-encoder rerankers score query-document pairs directly, producing more accurate relevance scores than bi-encoders. They are typically used to rerank a small set of candidate results."
  "TF-IDF stands for Term Frequency-Inverse Document Frequency. It reflects how important a word is to a document relative to a collection. Words that appear frequently in one document but rarely across all documents get higher scores."

  # Machine learning
  "ONNX Runtime is a cross-platform inference engine for machine learning models. It supports models trained in PyTorch, TensorFlow, and other frameworks exported to the ONNX format."
  "Transformer models use self-attention mechanisms to process input sequences in parallel. BERT, GPT, and T5 are well-known transformer architectures used in NLP tasks."
  "Embedding models convert text into dense vector representations that capture semantic meaning. Similar texts produce vectors that are close together in the embedding space."
  "A cross-encoder takes a pair of sentences as input and outputs a relevance score. Unlike bi-encoders, cross-encoders can attend to both sentences simultaneously for better accuracy."

  # DevOps and tools
  "Docker containers package applications with their dependencies into standardized units for software development. Containers share the host OS kernel, making them lighter than virtual machines."
  "Kubernetes is an open-source container orchestration platform that automates deploying, scaling, and managing containerized applications across clusters of machines."
  "Git is a distributed version control system that tracks changes in source code during software development. It supports non-linear development through branching and merging."

  # Networking
  "HTTP/2 improves web performance through multiplexing, header compression, and server push. It uses a single TCP connection for multiple concurrent streams."
  "gRPC is a high-performance remote procedure call framework that uses Protocol Buffers for serialization. It supports streaming, authentication, and load balancing out of the box."
  "WebSockets provide full-duplex communication channels over a single TCP connection. They are commonly used for real-time applications like chat, gaming, and live data feeds."

  # Architecture patterns
  "Microservices architecture structures an application as a collection of loosely coupled, independently deployable services. Each service owns its data and communicates via APIs or message queues."
  "Event sourcing stores all changes to application state as a sequence of events. Instead of storing current state, you can rebuild state by replaying events from the beginning."
  "The CQRS pattern separates read and write operations into different models. Command Query Responsibility Segregation can improve performance and scalability for complex domains."

  # Go-specific topics
  "Go interfaces are satisfied implicitly. A type implements an interface by implementing its methods, without any explicit declaration. This enables loose coupling and flexible design."
  "The Go standard library includes packages for HTTP servers, JSON encoding, cryptography, testing, and more. The net/http package alone is powerful enough for production web services."
  "Go modules are the standard way to manage dependencies in Go. A go.mod file at the root of a project declares the module path and its dependency requirements."
  "Error handling in Go uses explicit return values rather than exceptions. The error interface has a single method: Error() string. The errors package provides wrapping and unwrapping."
  "Go slices are dynamically-sized, flexible views into arrays. The append built-in function grows slices as needed, and the copy function transfers elements between slices."

  # Data structures
  "A B-tree is a self-balancing tree data structure that maintains sorted data and allows searches, insertions, and deletions in logarithmic time. Databases commonly use B-trees for indexing."
  "Hash tables provide O(1) average-case lookup by mapping keys to array positions using a hash function. Go maps are implemented as hash tables internally."
  "A trie (prefix tree) is a tree data structure used for storing strings where each node represents a character. Tries enable efficient prefix-based search and autocomplete features."

  # Security
  "JWT (JSON Web Tokens) are compact, URL-safe tokens used for authentication and information exchange. They consist of a header, payload, and signature encoded in base64."
  "TLS (Transport Layer Security) encrypts communication between clients and servers. It uses a handshake protocol to establish a secure channel with certificate-based authentication."
  "OAuth 2.0 is an authorization framework that enables applications to obtain limited access to user accounts. It delegates authentication to the service hosting the user account."
)

count=0
for doc in "${docs[@]}"; do
  $BINARY add "$doc"
  count=$((count + 1))
done

echo ""
echo "=== Done! Added $count documents to collection ==="
echo ""
echo "Try these searches:"
echo "  $BINARY search \"programming language\""
echo "  $BINARY search \"concurrency goroutines\""
echo "  $BINARY search \"database\""
echo "  $BINARY search \"vector search\""
echo "  $BINARY search \"Go interfaces\""
echo "  $BINARY search \"machine learning\""
