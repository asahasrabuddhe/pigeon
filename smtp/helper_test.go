package smtp

import (
	"io"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

func compareBodies(t *testing.T, got, want string) {
	// We cannot do a simple comparison since the ordering of headers' fields
	// is random.
	gotLines := strings.Split(got, "\r\n")
	wantLines := strings.Split(want, "\r\n")

	// We only test for too many lines, missing lines are tested after
	if len(gotLines) > len(wantLines) {
		t.Fatalf("Message has too many lines, \ngot %d:\n%s\nwant %d:\n%s", len(gotLines), got, len(wantLines), want)
	}

	isInHeader := true
	headerStart := 0
	for i, line := range wantLines {
		if line == gotLines[i] {
			if line == "" {
				isInHeader = false
			} else if !isInHeader && len(line) > 2 && line[:2] == "--" {
				isInHeader = true
				headerStart = i + 1
			}
			continue
		}

		if !isInHeader {
			missingLine(t, line, got, want)
		}

		isMissing := true
		for j := headerStart; j < len(gotLines); j++ {
			if gotLines[j] == "" {
				break
			}
			if gotLines[j] == line {
				isMissing = false
				break
			}
		}
		if isMissing {
			missingLine(t, line, got, want)
		}
	}
}

func missingLine(t *testing.T, line, got, want string) {
	t.Fatalf("Missing line %q\ngot:\n%s\nwant:\n%s", line, got, want)
}

func getBoundaries(t *testing.T, count int, m string) []string {
	if matches := boundaryRegExp.FindAllStringSubmatch(m, count); matches != nil {
		boundaries := make([]string, count)
		for i, match := range matches {
			boundaries[i] = match[1]
		}
		return boundaries
	}

	t.Fatal("Boundary not found in body")
	return []string{""}
}

var boundaryRegExp = regexp.MustCompile("boundary=(\\w+)")

func mockCopyFile(name string) (string, FileSetting) {
	return name, SetCopyFunc(func(w io.Writer) error {
		_, err := w.Write([]byte("Content of " + filepath.Base(name)))
		return err
	})
}

func mockCopyFileWithHeader(name string, h map[string][]string) (string, FileSetting, FileSetting) {
	name, f := mockCopyFile(name)
	return name, f, SetHeader(h)
}
