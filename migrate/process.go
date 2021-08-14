package migrate

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const SUFFIX = "_updated"

var Overwrite bool

type Stat struct {
	function int
	assert int
}

var stat *Stat = &Stat{}

func Process(path string) {
	// Get the current working directory
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Print the current working directory
	if !strings.HasPrefix(path, "/") {
		path = wd + "/" + path
	}

	if _, exists := IsExists(path); exists {
		if _, exists := IsFile(path); exists {
			ProcessFile(path, path+SUFFIX)
		}
		if _, exists := IsDir(path); exists {
			ProcessFolder(path)
		}
	}

	fmt.Printf("changed assert linds %d, func lins %d", stat.assert, stat.function)
}

// ProcessFolder 获取当前目录下的所有文件或目录信息
func ProcessFolder(pwd string) {
	filepath.Walk(pwd, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			if strings.HasSuffix(info.Name(), "_test.go") {
				ProcessFile(path, path+SUFFIX)
			}
		}
		return nil
	})
}

func ProcessFile(infile string, outfile string) {
	f, err := os.Open(infile)
	if err != nil {
		log.Fatal(err)
	}
	s := bufio.NewScanner(f)
	var buffer bytes.Buffer
	for s.Scan() {
		line := s.Text()
		line, _ = ProcessLine(line)
		buffer.WriteString(line)
		buffer.WriteString("\n")
	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(outfile, buffer.Bytes(), 0644)
	if err != nil {
		log.Fatal(err)
	}
	if err = f.Close(); err != nil {
		log.Fatal(err)
	}
	if Overwrite {
		os.Remove(infile)
		os.Rename(outfile, infile)
	}
}

func ProcessLine(line string) (string, error) {
	if strings.Contains(line, ".Assert(") {
		stat.assert++
		if strings.Contains(line, " Equals") {
			r, err := Equals(line)
			if err != nil {
				return "process equals has error", err
			}
			newLine := fmt.Sprintf("	require.Equal(t, %s, %s)", r.expect, r.actual)
			return newLine, nil
		}
		if strings.Contains(line, " DeepEquals") {
			r, err := DeepEquals(line)
			if err != nil {
				return "process equals has error", err
			}
			newLine := fmt.Sprintf("	require.Equal(t, %s, %s)", r.expect, r.actual)
			return newLine, nil
		}
		if strings.Contains(line, " IsNil") {
			r, err := IsNil(line)
			if err != nil {
				return "process equals has error", err
			}
			newLine := fmt.Sprintf("	require.Nil(t, %s)", r.actual)
			return newLine, nil
		}
		if strings.Contains(line, " NotNil") {
			r, err := NotNil(line)
			if err != nil {
				return "process equals has error", err
			}
			newLine := fmt.Sprintf("	require.NotNil(t, %s)", r.actual)
			return newLine, nil
		}
		if strings.Contains(line, " IsTrue") {
			r, err := IsTrue(line)
			if err != nil {
				return "process equals has error", err
			}
			newLine := fmt.Sprintf("	require.True(t, %s)", r.actual)
			return newLine, nil
		}
		if strings.Contains(line, " IsFalse") {
			r, err := IsFalse(line)
			if err != nil {
				return "process equals has error", err
			}
			newLine := fmt.Sprintf("	require.False(t, %s)", r.actual)
			return newLine, nil
		}
		if strings.Contains(line, " Greater") {
			r, err := Greater(line)
			if err != nil {
				return "process equals has error", err
			}
			newLine := fmt.Sprintf("	require.Greater(t, %s, %s)", r.actual, r.expect)
			return newLine, nil
		}
		if strings.Contains(line, " LessEqual") {
			r, err := LessEqual(line)
			if err != nil {
				return "process equals has error", err
			}
			newLine := fmt.Sprintf("	require.LessOrEqual(t, %s)", r.actual)
			return newLine, nil
		}
		if strings.Contains(line, " GreaterEqual") {
			r, err := GreaterEqual(line)
			if err != nil {
				return "process equals has error", err
			}
			newLine := fmt.Sprintf("	require.GreaterOrEqual(t, %s)", r.actual)
			return newLine, nil
		}
		stat.assert--
	}

	if strings.Contains(line, "func ") && strings.Contains(line, " Test") {
		r, err := Function(line)
		if err != nil {
			return "process equals has error", err
		}
		if r.match {
			newLine := fmt.Sprintf("func %s(t *testing.T) {", r.name)
			stat.function++
			return newLine, nil
		}
		return line, nil
	}

	return line, nil
}
