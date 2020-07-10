package balancer

import (
	"reflect"
	"testing"
)

func Test_splitCells(t *testing.T) {
	type args struct {
		n int
		l uint64
	}
	tests := []struct {
		name    string
		args    args
		want    []Range
		wantErr bool
	}{
		{
			name: "simple test",
			args: args{
				n: 5,
				l: 5,
			},
			want: []Range{
				{0, 1, 1},
				{1, 2, 1},
				{2, 3, 1},
				{3, 4, 1},
				{4, 5, 1},
			},
			wantErr: false,
		},
		{
			name: "error test",
			args: args{
				n: 50,
				l: 5,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "complex test 1",
			args: args{
				n: 3,
				l: 20,
			},
			want: []Range{
				{0, 7, 7},
				{7, 14, 7},
				{14, 20, 6},
			},
			wantErr: false,
		},
		{
			name: "complex test 2",
			args: args{
				n: 5,
				l: 256,
			},
			want: []Range{
				{0, 52, 52},
				{52, 103, 51},
				{103, 154, 51},
				{154, 205, 51},
				{205, 256, 51},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := splitCells(tt.args.n, tt.args.l)
			if (err != nil) != tt.wantErr {
				t.Errorf("splitCells() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitCells() got = %v, want %v", got, tt.want)
			}
		})
	}
}
