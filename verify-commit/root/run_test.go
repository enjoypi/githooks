package root

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMatchExtension(t *testing.T) {
	cfg := map[string]map[string]bool{
		"*": {".asset": true, ".fbx": true, ".jpg": true, ".mat": true, ".meta": true, ".png": true, ".prefab": true, ".tga": true, ".shader": true},
	}

	matched, unmatched := match(cfg,
		[]string{".gitignore", "Makefile", "b b.mat", "pre-commit/run.go", "\"a/B/c/d e.tGa\""},
	)
	require.Equal(t,
		[]string{"b b.mat", "\"a/B/c/d e.tGa\""},
		matched,
	)
	require.Equal(t,
		unmatched,
		[]string{".gitignore", "Makefile", "pre-commit/run.go"},
	)
}

func TestMatchDirectory(t *testing.T) {
	cfg := map[string]map[string]bool{
		"a/b": {"*": true},
	}

	matched, unmatched := match(cfg,
		[]string{"a/git.png", "\"a/B/Makefile.jpg\"", "A/b/C/b b.Mat\"", "pre-commit/run.go.tga", "a/B/c/d.asset", "A/BBB/c"},
	)
	require.Equal(t,
		[]string{"\"a/B/Makefile.jpg\"", "A/b/C/b b.Mat\"", "a/B/c/d.asset"},
		matched,
	)
	require.Equal(t,
		[]string{"a/git.png", "pre-commit/run.go.tga", "A/BBB/c"},
		unmatched,
	)
}

func TestMatchDirectoryExtension(t *testing.T) {
	cfg := map[string]map[string]bool{
		"a/b": {".asset": true, ".fbx": true, ".jpg": true, ".mat": true, ".meta": true, ".png": true, ".prefab": true, ".tga": true, ".shader": true},
	}

	matched, unmatched := match(cfg,
		[]string{"a/git.png", "a/B/Makefile.jpg", "\"A/b/C/b b.Mat", "pre-commit/run.go.tga", "a/B/c/d.asset", "A/BBB/c.tga"},
	)
	require.Equal(t,
		[]string{"a/B/Makefile.jpg", "\"A/b/C/b b.Mat", "a/B/c/d.asset"},
		matched,
	)
	require.Equal(t,
		[]string{"a/git.png", "pre-commit/run.go.tga", "A/BBB/c.tga"},
		unmatched,
	)
}
