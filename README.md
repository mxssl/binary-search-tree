[![Build Status](https://cloud.drone.io/api/badges/mxssl/binary-search-tree/status.svg)](https://cloud.drone.io/mxssl/binary-search-tree)

# Binary Search Tree

[Binary Search Tree](https://en.wikipedia.org/wiki/Binary_search_tree) implementation in Go (Golang).

```sh
go test -v -cover ./...
=== RUN   TestRootValue
--- PASS: TestRootValue (0.00s)
=== RUN   TestLeftValue
--- PASS: TestLeftValue (0.00s)
=== RUN   TestLeftLeftValue
--- PASS: TestLeftLeftValue (0.00s)
=== RUN   TestLeftLeftLeftNode
--- PASS: TestLeftLeftLeftNode (0.00s)
=== RUN   TestLeftLeftRightNode
--- PASS: TestLeftLeftRightNode (0.00s)
=== RUN   TestLeftRightValue
--- PASS: TestLeftRightValue (0.00s)
=== RUN   TestLeftRightLeftValue
--- PASS: TestLeftRightLeftValue (0.00s)
=== RUN   TestLeftRightLeftLeft
--- PASS: TestLeftRightLeftLeft (0.00s)
=== RUN   TestLeftRightLeftRight
--- PASS: TestLeftRightLeftRight (0.00s)
=== RUN   TestLeftRightRightValue
--- PASS: TestLeftRightRightValue (0.00s)
=== RUN   TestLeftRightRightLeft
--- PASS: TestLeftRightRightLeft (0.00s)
=== RUN   TestLeftRightRightRight
--- PASS: TestLeftRightRightRight (0.00s)
=== RUN   TestRightValue
--- PASS: TestRightValue (0.00s)
=== RUN   TestRightLeft
--- PASS: TestRightLeft (0.00s)
=== RUN   TestRightRightValue
--- PASS: TestRightRightValue (0.00s)
=== RUN   TestRightRightRight
--- PASS: TestRightRightRight (0.00s)
=== RUN   TestRightRightLeftValue
--- PASS: TestRightRightLeftValue (0.00s)
=== RUN   TestRightRightLeftLeft
--- PASS: TestRightRightLeftLeft (0.00s)
=== RUN   TestRightRightLeftRight
--- PASS: TestRightRightLeftRight (0.00s)
=== RUN   TestContainsTrue
--- PASS: TestContainsTrue (0.00s)
=== RUN   TestContainsFalse
--- PASS: TestContainsFalse (0.00s)
=== RUN   TestRemoveValue
--- PASS: TestRemoveValue (0.00s)
=== RUN   TestRemoveLeft
--- PASS: TestRemoveLeft (0.00s)
=== RUN   TestRemoveRight
--- PASS: TestRemoveRight (0.00s)
PASS
coverage: 76.4% of statements
ok  	github.com/mxssl/binary-search-tree	0.006s	coverage: 76.4% of statements
```
