package database

import (
	"testing"

	"github.com/kerezsiz42/kinko/internal/utils"
)

func TestModifyRecord(t *testing.T) {
	original := NewRecord("Login", "email@email.com", "password", "This is very important")
	copy, err := utils.DeepCopy(original)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	copy.Modify(utils.Ptr("Updated Name"), nil, nil, nil)

	if copy.Key != original.Key {
		t.Errorf("Expected Key to be '%s', got '%s'", original.Key, copy.Key)
	}

	if copy.Name != "Updated Name" {
		t.Errorf("Expected Name to be 'Updated Name', got '%s'", original.Name)
	}

	if copy.Public != original.Public {
		t.Errorf("Expected Public to be '%s', got '%s'", original.Public, copy.Public)
	}

	if copy.Private != original.Private {
		t.Errorf("Expected Private to be '%s', got '%s'", original.Private, copy.Private)
	}

	if copy.Info != original.Info {
		t.Errorf("Expected Info to be '%s', got '%s'", original.Info, copy.Info)
	}

	if copy.UpdatedAt == original.UpdatedAt {
		t.Errorf("Expected UpdatedAt to be different, got the same")
	}

	if copy.CreatedAt != original.CreatedAt {
		t.Errorf("Expected CreatedAt to be '%s', got '%s'", original.CreatedAt, copy.CreatedAt)
	}
}
