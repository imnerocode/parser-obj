package parser_obj

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/imnerocode/vo-structures"
)

// ParseOBJ parses an OBJ file and converts it into a Model structure.
func ParseOBJ(filePath string) (*vo.Model, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	model := &vo.Model{
		Vertices: []vo.Vertex{},
		Faces:    []vo.Face{},
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 || strings.HasPrefix(line, "#") {
			continue // Ignore comments and empty lines
		}

		parts := strings.Fields(line)
		switch parts[0] {
		case "v": // Vertex definition
			if len(parts) < 4 {
				continue // Skip invalid vertex lines
			}
			x, _ := strconv.ParseFloat(parts[1], 32)
			y, _ := strconv.ParseFloat(parts[2], 32)
			z, _ := strconv.ParseFloat(parts[3], 32)
			model.Vertices = append(model.Vertices, vo.Vertex{X: float32(x), Y: float32(y), Z: float32(z)})

		case "f": // Face definition
			if len(parts) < 2 {
				continue // Skip invalid face lines
			}
			var indices []int32
			for _, part := range parts[1:] {
				vertexIndex, _ := strconv.Atoi(strings.Split(part, "/")[0])
				indices = append(indices, int32(vertexIndex-1)) // OBJ is 1-based
			}
			model.Faces = append(model.Faces, vo.Face{VertexIndices: indices})
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return model, nil
}
