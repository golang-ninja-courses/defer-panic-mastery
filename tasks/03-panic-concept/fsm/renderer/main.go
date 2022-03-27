package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
	"go.uber.org/multierr"

	"github.com/www-golang-courses-ru/advanced-dealing-with-panic-in-go/tasks/03-panic-concept/fsm"
)

const (
	state1 fsm.State = "state-1"
	state2 fsm.State = "state-2"
	state3 fsm.State = "state-3"
	state4 fsm.State = "state-4"
	state5 fsm.State = "state-5"
)

var fsmExample = fsm.FSM{
	fsm.StateInitial: {state1, state3},
	state1:           {state2},
	state2:           {fsm.StateEnd},
	state3:           {state1, state4},
	state4:           {state5},
	state5:           {state3, fsm.StateEnd},
}

func main() {
	f, err := os.CreateTemp("", "fsm*.png")
	if err != nil {
		log.Panic(err)
	}
	defer f.Close()

	if err := renderFSM(fsmExample, f.Name()); err != nil {
		log.Panic(err)
	}

	if err := f.Sync(); err != nil {
		log.Panic(err)
	}
	log.Println("look at", f.Name())
}

func renderFSM(f fsm.FSM, filename string) (err error) {
	gv := graphviz.New()
	defer multierr.AppendInvoke(&err, multierr.Close(gv))

	graph, err := gv.Graph()
	if err != nil {
		return fmt.Errorf("create new graph: %v", err)
	}
	defer multierr.AppendInvoke(&err, multierr.Close(graph))

	if err := renderNodes(adaptFSMToNodes(f), graph); err != nil {
		return fmt.Errorf("render nodes: %v", err)
	}

	return gv.RenderFilename(graph, graphviz.PNG, filename)
}

type node struct {
	value    fsm.State
	children []fsm.State
}

func adaptFSMToNodes(f fsm.FSM) []node {
	nodes := make([]node, 0, len(f))

	// Реализуй формирование nodes.
	// ...
	_ = nodes
	panic("unimplemented")

	sort.SliceStable(nodes, func(i, j int) bool {
		return nodes[i].value < nodes[j].value
	})
	return nodes
}

func renderNodes(nodes []node, g *cgraph.Graph) error {
	for _, n := range nodes {
		parentNode, err := g.CreateNode(string(n.value))
		if err != nil {
			return fmt.Errorf("create parent node %q: %v", n.value, err)
		}

		// Реализуй создание дочерних узлов и ветвей между родительским узлом и дочерними.
		// ...
		_ = parentNode
		panic("unimplemented")
	}
	return nil
}
