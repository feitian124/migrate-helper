package migrate

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// ProcessFolder 获取当前目录下的所有文件或目录信息
func ProcessFolder(pwd string) {
	filepath.Walk(pwd, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			if strings.HasSuffix(info.Name(), "_test.go") {
				fmt.Println(path)
				ProcessFile(path)
			}
		}
		return nil
	})
}

func ProcessFile(file string) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		line, _ = ProcessLine(line)
		fmt.Println(line)
	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func ProcessLine(line string) (string, error) {
	if strings.Contains(line, ".Assert(") {
		if strings.Contains(line, " Equals") {
			r, err := Equals(line)
			if err != nil {
				return "process equals has error", err
			}
			newLine := fmt.Sprintf("	require.Equal(%s, 2, %s)", r.expect, r.actual)
			return newLine, nil
		}
		if strings.Contains(line, " DeepEquals") {
			r, err := DeepEquals(line)
			if err != nil {
				return "process equals has error", err
			}
			newLine := fmt.Sprintf("	require.Equal(%s, 2, %s)", r.expect, r.actual)
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
	}

	if strings.Contains(line, "func ") && strings.Contains(line, " Test") {
		r, err := Function(line)
		if err != nil {
			return "process equals has error", err
		}
		if r.match {
			newLine := fmt.Sprintf("func %s(t *testing.T) {", r.name)
			return newLine, nil
		}
		return line, nil
	}

	return line, nil
}
