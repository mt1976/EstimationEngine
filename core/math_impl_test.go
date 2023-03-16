package core

import "testing"

func TestRoundToNearest(t *testing.T) {
	type args struct {
		number         float64
		RoundingFactor float64
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RoundUpTo(tt.args.number, tt.args.RoundingFactor)
			if (err != nil) != tt.wantErr {
				t.Errorf("RoundToNearest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RoundToNearest() = %v, want %v", got, tt.want)
			}
		})
	}
}
