package detect

import (
	"os"
	"path/filepath"
	"testing"

	"sensitive-lexicon/internal/lexicon"
)

func TestDetectFindsLexiconWordInsideText(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "words.txt"), []byte("力工梭哈定律\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	store := lexicon.NewStore()
	if err := store.LoadFromDir(dir); err != nil {
		t.Fatal(err)
	}

	service := NewService(store)
	got := service.Detect(DetectRequest{Text: "这里测试力工梭哈定律", EnableFuzzy: false})

	if len(got.Hits) != 1 {
		t.Fatalf("expected 1 hit, got %d: %#v", len(got.Hits), got.Hits)
	}
	if got.Hits[0].Word != "力工梭哈定律" {
		t.Fatalf("expected 力工梭哈定律 hit, got %#v", got.Hits[0])
	}
	if got.Hits[0].Type != "substring" {
		t.Fatalf("expected substring hit, got %#v", got.Hits[0])
	}
}
