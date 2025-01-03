package parser_obj_test

import (
	"os"
	"testing"

	parser_obj "github.com/imnerocode/parser-obj"
	"github.com/imnerocode/vo-structures"
)

func TestParseOBJ(t *testing.T) {
	// Create a temporary OBJ file for testing
	objData := `
	# This is a comment
	v 1.0 2.0 3.0
	v 4.0 5.0 6.0
	v 7.0 8.0 9.0
	f 1 2 3
	`
	tmpFile, err := os.CreateTemp("", "test.obj")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.WriteString(objData); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	if err := tmpFile.Close(); err != nil {
		t.Fatalf("Failed to close temp file: %v", err)
	}

	// Parse the OBJ file
	model, err := parser_obj.ParseOBJ(tmpFile.Name())
	if err != nil {
		t.Fatalf("Failed to parse OBJ file: %v", err)
	}

	// Validate the parsed data
	expectedVertices := []vo.Vertex{
		{X: 1.0, Y: 2.0, Z: 3.0},
		{X: 4.0, Y: 5.0, Z: 6.0},
		{X: 7.0, Y: 8.0, Z: 9.0},
	}
	expectedFaces := []vo.Face{
		{VertexIndices: []int32{0, 1, 2}},
	}

	t.Logf("Faces: %+v", model.Faces)

	if len(model.Vertices) != len(expectedVertices) {
		t.Fatalf("Expected %d vertices, got %d", len(expectedVertices), len(model.Vertices))
	}
	for i, v := range model.Vertices {
		if v != expectedVertices[i] {
			t.Errorf("Expected vertex %v, got %v", expectedVertices[i], v)
		}
	}

	if len(model.Faces) != len(expectedFaces) {
		t.Fatalf("Expected %d faces, got %d", len(expectedFaces), len(model.Faces))
	}
	for i, f := range model.Faces {
		if len(f.VertexIndices) != len(expectedFaces[i].VertexIndices) {
			t.Fatalf("Expected %d vertex indices, got %d", len(expectedFaces[i].VertexIndices), len(f.VertexIndices))
		}
		for j, idx := range f.VertexIndices {
			if idx != expectedFaces[i].VertexIndices[j] {
				t.Errorf("Expected vertex index %d, got %d", expectedFaces[i].VertexIndices[j], idx)
			}
		}
	}
}
