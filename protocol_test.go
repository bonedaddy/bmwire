package bmwire_test

import (
	"testing"

	"github.com/jimmysong/bmwire"
)

// TestServiceFlagStringer tests the stringized output for service flag types.
func TestServiceFlagStringer(t *testing.T) {
	tests := []struct {
		in   bmwire.ServiceFlag
		want string
	}{
		{0, "0x0"},
		{bmwire.SFNodeNetwork, "SFNodeNetwork"},
		{0xffffffff, "SFNodeNetwork|0xfffffffe"},
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
		result := test.in.String()
		if result != test.want {
			t.Errorf("String #%d\n got: %s want: %s", i, result,
				test.want)
			continue
		}
	}
}

// TestBitmessageNetStringer tests the stringized output for bitmessage net types.
func TestBitmessageNetStringer(t *testing.T) {
	tests := []struct {
		in   bmwire.BitmessageNet
		want string
	}{
		{bmwire.MainNet, "MainNet"},
		{0xffffffff, "Unknown BitmessageNet (4294967295)"},
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
		result := test.in.String()
		if result != test.want {
			t.Errorf("String #%d\n got: %s want: %s", i, result,
				test.want)
			continue
		}
	}
}
