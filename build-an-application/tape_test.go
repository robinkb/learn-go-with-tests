package poker

import (
	"io"
	"testing"
)

func TestTapeWrite(t *testing.T) {
	file, clean := createTempFile(t, "12345")
	defer clean()

	tape := &tape{file}

	tape.Write([]byte("abc"))

	file.Seek(0, io.SeekStart)
	newContents, _ := io.ReadAll(file)

	got := string(newContents)
	want := "abc"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
