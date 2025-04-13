package computationmodel

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Model interface {
	Name() string
	Transform(v Vector) Vector
}

type Models interface {
	CreateModel(name string) (Model, error)
}

type GenericModel struct {
	name   string
	matrix Matrix
}

var _ Model = (*GenericModel)(nil)

func NewModel(name string) (*GenericModel, error) {
	path := fmt.Sprintf("matrices/%s.mat", strings.Title(name))

	model := &GenericModel{name: name}
	matrix, err := model.loadMatrixFromFile(path)
	if err != nil {
		return nil, err
	}
	model.matrix = matrix

	return model, nil
}

func (model *GenericModel) loadMatrixFromFile(path string) (Matrix, error) {
	file, err := os.Open(path)
	if err != nil {
		return Matrix{}, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	matrix := [1000][1000]float64{}
	scanner := bufio.NewScanner(file)

	index := 0
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		row := [1000]float64{}
		for i, field := range fields {
			val, err := strconv.ParseFloat(field, 64)
			if err != nil {
				return Matrix{}, fmt.Errorf("invalid number %q: %w", field, err)
			}
			row[i] = val
		}
		matrix[index] = row
		index++
	}

	if err := scanner.Err(); err != nil {
		return Matrix{}, fmt.Errorf("failed to read file: %w", err)
	}

	return matrix, nil
}

func (model *GenericModel) Name() string {
	return model.name
}

func (model *GenericModel) Transform(vector Vector) Vector {
	// TODO: linear algebra
	matrix := model.matrix
	n := len(matrix)

	if n == 0 || len(model.matrix[0]) != len(vector) {
		panic(fmt.Errorf("dimension mismatch: matrix is %dx%d, vector is %d", len(matrix), len(matrix[0]), len(vector)))
	}

	result := [1000]float64{}
	for i := range n {
		var sum float64
		for j := range len(vector) {
			sum += matrix[i][j] * vector[j]
		}
		result[i] = sum
	}

	return result
}

type GenericModels struct {
	modelNameModelMap map[string]Model
	mu                sync.Mutex
}

var _ Models = (*GenericModels)(nil)

func NewModels() *GenericModels {
	return &GenericModels{modelNameModelMap: make(map[string]Model)}
}

func (models *GenericModels) CreateModel(modelName string) (model Model, err error) {
	models.mu.Lock()
	defer models.mu.Unlock()
	model, ok := models.modelNameModelMap[modelName]
	if !ok {
		model, err = NewModel(modelName)
		if err != nil {
			return nil, err
		}
		models.modelNameModelMap[modelName] = model
	}

	return model, nil
}
