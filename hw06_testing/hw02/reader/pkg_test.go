package reader

import (
	"github.com/sergeyptv/home_work_basic/hw06_testing/hw02/types"
	"reflect"
	"testing"
)

func TestReader(t *testing.T) {
	testCases := []struct {
		desc     string
		filePath string
		res      []types.Employee
	}{
		{
			desc:     "Positive",
			filePath: "../data.json",
			res: []types.Employee{
				{
					UserID:       10,
					Age:          25,
					Name:         "Rob",
					DepartmentID: 3,
				},
				{
					UserID:       11,
					Age:          30,
					Name:         "George",
					DepartmentID: 2,
				},
			},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got, err := ReadJSON(tC.filePath)
			if err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(got, tC.res) {
				t.Errorf("got %v, want %v", got, tC.res)
			}
		})
	}
}
