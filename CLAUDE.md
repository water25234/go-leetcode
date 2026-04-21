# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Repository purpose

Personal LeetCode practice in Go, organized by algorithm pattern. The `README.md` lists 21 patterns grouped into three interview-priority tiers (必考 / 常考 / 加分); study notes for each pattern live in `0-Topic/<Pattern>.md` (Traditional Chinese).

## Layout

Every problem lives under a **pattern folder** at the repo root: `<Pattern>/<id>-<kebab-title>/<files>.go`. Patterns on disk: `Backtracking`, `BFS`, `Binary-Search`, `DFS`, `Dynamic-Programming`, `Greedy`, `Hash-Map`, `Heap`, `Intervals`, `Linked-List`, `Math`, `Monotonic-Stack`, `Prefix-Sum`, `Sliding-Window`, `Stack`, `Tree-Traversal`, `Two-Pointer`. Tree problems go in `Tree-Traversal/` (not `Linked-List/`). `Stack/` is for general stack use; `Monotonic-Stack/` is reserved for "next greater / smaller" style problems. Some folders are pre-created empty placeholders for common patterns from `README.md`'s 21-pattern table — create new ones for patterns not yet listed when needed.

`practice/` holds scratch/WIP files that aren't finalized solutions. The top-level `main.go` is an unrelated scratchpad.

## Per-problem conventions

- Each problem directory is **self-contained**. There is **no `go.mod`** at the repo root or anywhere — problems do not import each other.
- A problem dir usually contains a single `main.go`. When a problem has multiple approaches, each approach is a separate `.go` file in the same dir sharing one package (e.g. `133-Clone-Graph/` has `BFS.go` + `DFS.go` both in `package clonegraph`; `19-Remove-Nth-Node-From-End-of-List/` has `main.go` + `twoPointers.go`). Shared types like `ListNode` / `Node` are defined once in that package.
- Package naming is inconsistent across the repo — some dirs use `package main`, others use a problem-specific name (`package twosum`, `package clonegraph`, `package binarysearch`). **Match the package name already present in the target dir**; if creating a new dir, either is acceptable but pick one name and keep all files in that dir consistent so the dir compiles.
- The LeetCode problem statement / constraints are often pasted as a leading comment block. Keep this style when adding new problems.

## Commands

Because there is no module, commands run per-directory:

```bash
# Build / run a single problem (only works if its package is `main` and it has a main() — most don't)
go run ./<Pattern>/<id>-<title>/

# Type-check / compile-check a problem dir regardless of package name
go build ./<Pattern>/<id>-<title>/

# Format
gofmt -w <file-or-dir>

# Vet
go vet ./<Pattern>/<id>-<title>/
```

There are no `*_test.go` files and no test runner wired up — `go test` is not part of the workflow here. If the user asks to add tests, create `*_test.go` in the same problem dir and under the same package.

## When adding a new problem

1. Decide the pattern from `README.md`'s 21-pattern list and place the dir under that pattern folder (e.g. `Greedy/<id>-<title>/`). Use the existing kebab-case `<id>-<Title-Words>` naming.
2. Create `main.go`; if you want to show alternate approaches, add sibling files like `bruteForce.go`, `twoPointers.go`, `BFS.go`, `DFS.go` in the same package.
3. Lead with the problem statement and constraints as a comment block, mirroring neighboring problems.
