package pipeline

import (
	"crypto/md5"
	"fmt"
	"go/parser"
	"go/token"
	"sort"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

// 执行 go test -benchmem -run=^$ -bench ^BenchmarkMD5All github.com/frlute/golang-bootcamp/concurrency/pipeline -v

type execFunc func(string) (map[string][md5.Size]byte, error)

func testFunc(f execFunc) (map[string][md5.Size]byte, error) {
	root := "../pipeline"

	return f(root)
}

func Test_MD5ALL(t *testing.T) {
	m, err := testFunc(MD5All)
	assert.Nil(t, err)

	var paths []string
	for path := range m {
		paths = append(paths, path)
	}
	sort.Strings(paths)
	for _, path := range paths {
		fmt.Printf("%x %s\n", m[path], path)
	}
}

func Test_MD5ALL_2(t *testing.T) {
	m, err := testFunc(MD5All2)
	assert.Nil(t, err)

	var paths []string
	for path := range m {
		paths = append(paths, path)
	}
	sort.Strings(paths)
	for _, path := range paths {
		fmt.Printf("%x %s\n", m[path], path)
	}
}

func Test_MD5ALL_3(t *testing.T) {
	m, err := testFunc(MD5All3)
	assert.Nil(t, err)

	var paths []string
	for path := range m {
		paths = append(paths, path)
	}
	sort.Strings(paths)
	for _, path := range paths {
		fmt.Printf("%x %s\n", m[path], path)
	}
}

func Test_ParseGoSourceFile(t *testing.T) {
	fs := token.NewFileSet()
	f, err := parser.ParseFile(fs, "./parallel.go", nil, parser.AllErrors)
	if err != nil {
		panic(err)
	}
	spew.Dump(f)
}
