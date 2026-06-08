package basicerrors

import "testing"

func TestValidateAge(t *testing.T) {
	tests := []struct {
		name    string
		age     int
		wantErr bool
	}{
		{"valid", 25, false},
		{"zero", 0, false},
		{"max valid", 150, false},
		{"negative", -1, true},
		{"too old", 151, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateAge(tt.age)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateAge(%d) error = %v, wantErr %v", tt.age, err, tt.wantErr)
			}
		})
	}
}
