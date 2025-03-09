package printer

import (
	"bytes"
	"fmt"
	"github.com/sergeyptv/home_work_basic/hw06_testing/hw02/types"
	"github.com/stretchr/testify/require"
	"io"
	"os"
	"testing"
)

func TestPrinter(t *testing.T) {
	testCases := []struct {
		desc  string
		staff []types.Employee
		res   []types.Employee
	}{
		{
			desc: "Positive",
			staff: []types.Employee{
				{
					UserID:       10,
					Age:          25,
					Name:         "Rob",
					DepartmentID: 3,
				},
			},
			res: []types.Employee{
				{
					UserID:       10,
					Age:          25,
					Name:         "Rob",
					DepartmentID: 3,
				},
			},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			old := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			PrintStaff(tC.staff)

			os.Stdout = old

			outC := make(chan string)
			go func() {
				var buf bytes.Buffer
				_, err := io.Copy(&buf, r)
				require.NoError(t, err)

				strRes := buf.String()

				outC <- strRes
				close(outC)
			}()

			err := w.Close()
			require.NoError(t, err)

			got := <-outC

			res := fmt.Sprintf("User ID: %d; Age: %d; Name: %s; Department ID: %d; \n"+
				"User ID: %d; Age: %d; Name: %s; Department ID: %d; \n",
				tC.res[0].UserID, tC.res[0].Age, tC.res[0].Name, tC.res[0].DepartmentID,
				tC.res[0].UserID, tC.res[0].Age, tC.res[0].Name, tC.res[0].DepartmentID)

			require.Equal(t, res, got)
		})
	}
}
