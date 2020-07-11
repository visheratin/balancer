package curve

import (
	"testing"
)

func TestDrawCurve(t *testing.T) {
	type args struct {
		cType CurveType
		bits  uint64
		op    string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Hilbert 5 bits",
			args: args{
				cType: Hilbert,
				bits:  5,
				op:    "hilbert.png",
			},
			wantErr: false,
		},
		{
			name: "Morton 5 bits",
			args: args{
				cType: Morton,
				bits:  5,
				op:    "morton.png",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DrawCurve(tt.args.cType, tt.args.bits, tt.args.op); (err != nil) != tt.wantErr {
				t.Errorf("DrawCurve() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDrawSplitCurve(t *testing.T) {
	type args struct {
		cType  CurveType
		bits   uint64
		splits []uint64
		op     string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Hilbert 8 bits",
			args: args{
				cType:  Hilbert,
				bits:   8,
				splits: []uint64{512, 2048, 10196, 32075, 50000, 65535},
			},
			wantErr: false,
		},
		{
			name: "Morton 8 bits",
			args: args{
				cType:  Morton,
				bits:   8,
				splits: []uint64{512, 2048, 10196, 32075, 50000, 65535},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := NewCurve(tt.args.cType, 2, tt.args.bits)
			if err != nil {
				t.Fatal(err)
			}
			if _, err := DrawSplitCurve(c, tt.args.splits); (err != nil) != tt.wantErr {
				t.Errorf("DrawSplitCurve() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
