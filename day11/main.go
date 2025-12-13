package main

import (
	_ "embed"
	"strings"
	"time"

	"github.com/jbedard/aoc2025/lib"
)

//go:embed input.txt
var content string

type node int32
type path []node

type depMap [][]node
type depStore struct {
	named map[string]node
	deps  depMap
}

func parseInput(input string) depStore {
	result := depStore{
		named: map[string]node{},
		deps:  depMap{},
	}

	for line := range lib.ReadLines(input) {
		parts := strings.SplitN(line, ": ", 2)
		name := parts[0]
		deps := strings.Fields(parts[1])

		idx, hasIdx := result.named[name]
		if !hasIdx {
			idx = node(len(result.deps))
			result.named[name] = idx
			result.deps = append(result.deps, nil)
		}

		for _, d := range deps {
			dIdx, dHasIdx := result.named[d]
			if !dHasIdx {
				dIdx = node(len(result.deps))
				result.named[d] = dIdx
				result.deps = append(result.deps, nil)
			}
			result.deps[idx] = append(result.deps[idx], dIdx)
		}
	}

	return result
}

func main() {
	d := parseInput(content)

	defer lib.CpuProfile()()

	r1 := part1(d)
	println("Part 1:", r1)

	r2 := part2(d)
	println("Part 2:", r2)
}

func part1(d depStore) int {
	visited := make(map[node][]path)
	part1_visit(d.deps, visited, []node{}, d.named["you"])
	return len(visited[d.named["out"]])
}

func part1_visit(d depMap, visited map[node][]path, p []node, todo node) {
	here := append(p[:], todo)
	visited[todo] = append(visited[todo], here)
	for _, next := range d[todo] {
		part1_visit(d, visited, here, next)
	}
}

var startT = time.Now()

type nodeSet = map[node]struct{}

func part2(d depStore) int {
	// svr ... {fft + dac} ... out

	defer lib.ProgressDone()

	rev := make(depMap, len(d.deps))
	for from, tos := range d.deps {
		for _, to := range tos {
			rev[to] = append(rev[to], node(from))
		}
	}

	traversals := [][]node{
		{d.named["svr"], d.named["fft"], d.named["dac"], d.named["out"]},
		{d.named["svr"], d.named["dac"], d.named["fft"], d.named["out"]},
	}
	from := map[node]nodeSet{}
	to := map[node]nodeSet{}

	r := 0
	for _, trav := range traversals {
		pc := 1
		for i := range len(trav) - 1 {
			if from[trav[i]] == nil {
				from[trav[i]] = reachableNodes(d.deps, trav[i])
			}
			if to[trav[i+1]] == nil {
				to[trav[i+1]] = reachableNodes(rev, trav[i+1])
			}

			pc *= countPaths(d, intersectSet(from[trav[i]], to[trav[i+1]]), trav[i], trav[i+1])
			if pc == 0 {
				break
			}
		}
		r += pc
	}

	return r
}

func countPaths(d depStore, nodes nodeSet, from, to node) int {
	lib.Progress(startT, "%d...%d", from, to)

	nodes[from] = struct{}{}
	localD := make(depMap, len(d.deps))
	for n := range nodes {
		for _, dn := range d.deps[n] {
			if _, ok := nodes[dn]; ok || dn == to {
				localD[n] = append(localD[n], dn)
			}
		}
	}
	walkState := make([]int, len(d.deps))
	return part2_nav_visit(localD, walkState, from, to)
}

func part2_nav_visit(d depMap, walkState []int, current, target node) int {
	if current == target {
		return 1
	}

	if walkState[current] != 0 {
		// Seen before
		return walkState[current]
	}

	walkState[current] = -1

	foundSomething := 0
	for _, next := range d[current] {
		pathsFound := part2_nav_visit(d, walkState, next, target)
		if 0 < pathsFound {
			foundSomething += pathsFound
		}
	}

	if foundSomething > 0 {
		walkState[current] = foundSomething
	}

	return foundSomething
}

func reachableNodes(n depMap, start node) nodeSet {
	result := make(nodeSet, len(n))
	todo := []node{start}
	for len(todo) > 0 {
		current := todo[0]
		todo = todo[1:]

		if _, ok := result[current]; ok {
			continue
		}
		result[current] = struct{}{}
		todo = append(todo, n[current]...)
	}
	return result
}

func intersectSet(a, b nodeSet) nodeSet {
	result := make(nodeSet, (len(a)+len(b))/2)
	for k := range a {
		if _, ok := b[k]; ok {
			result[k] = struct{}{}
		}
	}
	return result
}
