package keyring

import (
	"testing"
)

func TestKeyring(t *testing.T) {
	tempDir := t.TempDir()
	kr, err := NewKeyring("testapp", BackendTest, tempDir, nil)
	if err != nil {
		t.Errorf("failed to create keyring: %v", err)
		return
	}

	t.Run("NewItem", func(t *testing.T) {
		record, err := kr.NewItem("testkey1", []byte("testvalue1"), "json")
		if err != nil {
			t.Errorf("failed to create new item: %v", err)
			return
		}
		if record.Key != "testkey1" {
			t.Errorf("expected key to be 'testkey1', got %s", record.Key)
		}
		if string(record.Data) != "testvalue1" {
			t.Errorf("expected data to be 'testvalue1', got %s", record.Data)
		}
		if record.CodecType != "json" {
			t.Errorf("expected codecType to be 'json', got %s", record.CodecType)
		}
	})

	t.Run("Get", func(t *testing.T) {
		record, err := kr.Get("testkey1")
		if err != nil {
			t.Errorf("failed to get item: %v", err)
			return
		}
		if record.Key != "testkey1" {
			t.Errorf("expected key to be 'testkey1', got %s", record.Key)
		}
		if string(record.Data) != "testvalue1" {
			t.Errorf("expected data to be 'testvalue1', got %s", record.Data)
		}
		if record.CodecType != "json" {
			t.Errorf("expected codecType to be 'json', got %s", record.CodecType)
		}

		_, err = kr.Get("nonexistent")
		if err == nil {
			t.Errorf("expected error for nonexistent key, but got nil")
		}
	})

	t.Run("List", func(t *testing.T) {
		_, err := kr.NewItem("testkey2", []byte("testvalue2"), "protobuf")
		if err != nil {
			t.Errorf("failed to create new item: %v", err)
			return
		}

		records, err := kr.List()
		if err != nil {
			t.Errorf("failed to list items: %v", err)
			return
		}
		if len(records) != 2 {
			t.Errorf("expected 2 items, got %d", len(records))
		}

		keys := make(map[string]struct{})
		for _, record := range records {
			keys[record.Key] = struct{}{}
		}
		if _, ok := keys["testkey1"]; !ok {
			t.Errorf("expected key 'testkey1' to be present")
		}
		if _, ok := keys["testkey2"]; !ok {
			t.Errorf("expected key 'testkey2' to be present")
		}
	})

	t.Run("Delete", func(t *testing.T) {
		err := kr.Delete("testkey1")
		if err != nil {
			t.Errorf("failed to delete item: %v", err)
			return
		}

		_, err = kr.Get("testkey1")
		if err == nil {
			t.Errorf("expected error for deleted key, but got nil")
		}

		err = kr.Delete("nonexistent")
		if err == nil {
			t.Errorf("expected error for nonexistent key, but got nil")
		}
	})
}
