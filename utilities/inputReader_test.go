package utilities

import (
	"path"
	"reflect"
	"runtime"
	"testing"
)

func Test_ReadFileToIntSlice(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{
			"Empty file",
			args{"empty_file"},
			[]int{},
			false,
		},
		{
			"Parseable file",
			args{"list_of_ints"},
			[]int{1, 2, 3, 4, 5, 6, 7, 8},
			false,
		},
		{
			"Unparseable file",
			args{"list_of_ints_with_unparseable_value"},
			[]int{1, 2, 3, 4, 5, 6, 7},
			true,
		},
	}

	_, currentFile, _, _ := runtime.Caller(0)
	inputsDir := path.Dir(currentFile) + "/test_inputs"

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadFileToIntSlice(inputsDir + "/" + tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadFileToIntSlice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadFileToIntSlice() got = %v, want %v", got, tt.want)
			}
		})
	}
}
