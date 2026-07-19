package account_test

import (
	"learn-trivia/internal/account"
	"testing"
)

func TestAccountStatus_String_ShouldReturnCorrectLabel(t *testing.T) {
	tests := []struct {
		status   account.AccountStatus
		expected string
	}{
		{account.StatusPending, "Pending"},
		{account.StatusActive, "Active"},
		{account.StatusDormant, "Dormant"},
		{account.StatusInactive, "Inactive"},
	}

	for _, tt := range tests {
		if tt.status.String() != tt.expected {
			t.Errorf("expected %q, got %q", tt.expected, tt.status.String())
		}
	}
}
