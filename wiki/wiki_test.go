package wiki

import (
	"bytes"
	"testing"

	"github.com/peterhellberg/wiki/db"
)

func TestNewWiki(t *testing.T) {
	db := &db.DB{}

	w, err := NewWiki(db)
	if err != nil {
		t.Fatal(err)
	}

	if w.DB() != db {
		t.Errorf(`unexpected db`)
	}
}

func TestGetPageName(t *testing.T) {
	w := &Wiki{}

	for _, tt := range []struct {
		page string
		want []byte
	}{
		{"", []byte("root")},
		{"foo", []byte("foo")},
		{"bar", []byte("bar")},
	} {
		if got := w.getPageName(tt.page); !bytes.Equal(got, tt.want) {
			t.Errorf("unexpected page name %q, want %q", got, tt.want)
		}
	}
}
