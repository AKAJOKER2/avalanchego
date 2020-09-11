// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package vertex

import (
	"errors"
	"testing"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/consensus/avalanche"
	"github.com/ava-labs/avalanchego/snow/consensus/snowstorm"
)

var (
	errParseVertex = errors.New("unexpectedly called ParseVertex")
	errBuildVertex = errors.New("unexpectedly called BuildVertex")
	errGetVertex   = errors.New("unexpectedly called GetVertex")
	errEdge        = errors.New("unexpectedly called Edge")
	errSaveVertex  = errors.New("unexpectedly called SaveVertex")
)

// TestManager ...
type TestManager struct {
	T *testing.T

	CantParseVertex, CantBuildVertex,
	CantGetVertex, CantEdge, CantSaveVertex bool

	ParseVertexF func([]byte) (avalanche.Vertex, error)
	BuildVertexF func(ids.Set, []snowstorm.Tx) (avalanche.Vertex, error)
	GetVertexF   func(ids.ID) (avalanche.Vertex, error)
	EdgeF        func() []ids.ID
	SaveVertexF  func(vtx avalanche.Vertex) error
}

// Default ...
func (m *TestManager) Default(cant bool) {
	m.CantParseVertex = cant
	m.CantBuildVertex = cant
	m.CantGetVertex = cant
	m.CantEdge = cant
	m.CantSaveVertex = cant
}

// ParseVertex ...
func (m *TestManager) ParseVertex(b []byte) (avalanche.Vertex, error) {
	if m.ParseVertexF != nil {
		return m.ParseVertexF(b)
	}
	if m.CantParseVertex && m.T != nil {
		m.T.Fatal(errParseVertex)
	}
	return nil, errParseVertex
}

// BuildVertex ...
func (m *TestManager) BuildVertex(set ids.Set, txs []snowstorm.Tx) (avalanche.Vertex, error) {
	if m.BuildVertexF != nil {
		return m.BuildVertexF(set, txs)
	}
	if m.CantBuildVertex && m.T != nil {
		m.T.Fatal(errBuildVertex)
	}
	return nil, errBuildVertex
}

// GetVertex ...
func (m *TestManager) GetVertex(id ids.ID) (avalanche.Vertex, error) {
	if m.GetVertexF != nil {
		return m.GetVertexF(id)
	}
	if m.CantGetVertex && m.T != nil {
		m.T.Fatal(errGetVertex)
	}
	return nil, errGetVertex
}

// Edge ...
func (m *TestManager) Edge() []ids.ID {
	if m.EdgeF != nil {
		return m.EdgeF()
	}
	if m.CantEdge && m.T != nil {
		m.T.Fatal(errEdge)
	}
	return nil
}

// SaveVertex ...
func (m *TestManager) SaveVertex(vtx avalanche.Vertex) error {
	if m.SaveVertexF != nil {
		return m.SaveVertexF(vtx)
	}
	if m.CantSaveVertex && m.T != nil {
		m.T.Fatal(errSaveVertex)
	}
	return nil
}
