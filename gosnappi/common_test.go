package gosnappi

import "testing"

func TestCheckClientServerVersionCompatibility(t *testing.T) {
	type args struct {
		clientVer string
		serverVer string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: add TCs for build versions (e.g. 0.1.1-200)
		{name: "patch-low-high", args: args{clientVer: "0.2.1", serverVer: "0.2.5"}, wantErr: false},
		{name: "patch-high-low", args: args{clientVer: "0.2.5", serverVer: "0.2.1"}, wantErr: false},
		{name: "minor-low-high", args: args{clientVer: "0.2.5", serverVer: "0.3.5"}, wantErr: false},
		{name: "minor-high-low", args: args{clientVer: "0.3.5", serverVer: "0.2.5"}, wantErr: true},
		{name: "major-low-high", args: args{clientVer: "0.2.5", serverVer: "1.2.5"}, wantErr: true},
		{name: "major-high-low", args: args{clientVer: "1.2.5", serverVer: "0.2.5"}, wantErr: true},
		{name: "invalid-valid", args: args{clientVer: "0.2.1.1", serverVer: "0.2.5"}, wantErr: true},
		{name: "valid-invalid", args: args{clientVer: "0.2.1", serverVer: "0.2.5.1"}, wantErr: true},
		{name: "invalid-invalid", args: args{clientVer: "0.2.1.1", serverVer: "0.2.5.1"}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := checkClientServerVersionCompatibility(tt.args.clientVer, tt.args.serverVer, "API"); (err != nil) != tt.wantErr {
				t.Errorf("checkClientServerVersionCompatibility() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
