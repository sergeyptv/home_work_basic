package reader

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/sergeyptv/home_work_basic/hw06_testing/hw02/types"
)

func ReadJSON(filePath string) ([]types.Employee, error) {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, err
	}

	bytes, err := io.ReadAll(f)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, err
	}

	var data []types.Employee

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, err
	}

	// res := data

	return data, nil
}
