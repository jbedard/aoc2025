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

func (n node) String() string {
	return string([]byte{byte('z' - (int32(n) >> 16)), byte('z' - ((int32(n) >> 8) & 0xff)), byte('z' - (int32(n) & 0xff))})
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

type walkState uint8

const (
	walkStateNone walkState = iota
	walkStateWalking
	walkStateWalked
	walkStateDeadend
)

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

	svr := node(d.named["svr"])
	fft := node(d.named["fft"])
	dac := node(d.named["dac"])
	out := node(d.named["out"])

	from := map[node]nodeSet{
		svr: reachableNodes(d.deps, svr),
		fft: reachableNodes(d.deps, fft),
		dac: reachableNodes(d.deps, dac),
	}
	to := map[node]nodeSet{
		fft: reachableNodes(rev, fft),
		dac: reachableNodes(rev, dac),
		out: reachableNodes(rev, out),
	}

	r := 0

	if fromFftToDac := intersectSet(from[fft], to[dac], nil); len(fromFftToDac) > 0 {
		fromSvrToFft := intersectSet(from[svr], to[fft], fromFftToDac)
		dacToOut := intersectSet(from[dac], to[out], fromFftToDac)

		r += countPaths(d, fromSvrToFft, svr, fft) * countPaths(d, fromFftToDac, fft, dac) * countPaths(d, dacToOut, dac, out)
	}

	if fromDacToFft := intersectSet(from[dac], to[fft], nil); len(fromDacToFft) > 0 {
		fromSvrToDac := intersectSet(from[svr], to[dac], fromDacToFft)
		fftToOut := intersectSet(from[fft], to[out], fromDacToFft)

		r += countPaths(d, fromSvrToDac, svr, dac) * countPaths(d, fromDacToFft, dac, fft) * countPaths(d, fftToOut, fft, out)
	}

	return r
}

func countPaths(d depStore, nodes nodeSet, from, to node) int {
	lib.Progress(startT, "%s...%s", from, to)

	nodes[from] = struct{}{}
	localD := make(depMap, len(d.deps))
	for n := range nodes {
		for _, dn := range d.deps[n] {
			if _, ok := nodes[dn]; ok || dn == to {
				localD[n] = append(localD[n], dn)
			}
		}
	}
	walkState := make([]walkState, len(d.deps))
	return part2_nav_visit(localD, walkState, from, to)
}

func part2_nav_visit(d depMap, walkState []walkState, current, target node) int {
	if current == target {
		return 1
	}

	switch walkState[current] {
	case walkStateWalking:
		// Gone in a circle
		return 0
	case walkStateDeadend:
		// Seen before and know it goes nowhere
		return 0
	}

	foundSomething := 0
	walkState[current] = walkStateWalking

	for _, next := range d[current] {
		nextFoundSomething := part2_nav_visit(d, walkState, next, target)
		if nextFoundSomething == 0 {
			walkState[next] = walkStateDeadend
		} else {
			walkState[next] = walkStateWalked
			foundSomething += nextFoundSomething
		}
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

func intersectSet(a, b nodeSet, exclude nodeSet) nodeSet {
	result := make(nodeSet, (len(a)+len(b))/2)
	for k := range a {
		if _, ok := b[k]; ok {
			if exclude != nil {
				if _, excluded := exclude[k]; excluded {
					continue
				}
			}
			result[k] = struct{}{}
		}
	}
	return result
}
