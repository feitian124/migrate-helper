package migrate

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Walk 获取当前目录下的所有文件或目录信息
func Walk(pwd string) {
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

func ProcessFile(file string)  {
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
			return r.expect + r.actual, nil
		}
		if strings.Contains(line, " DeepEquals") {
			r, err := DeepEquals(line)
			if err != nil {
				return "process equals has error", err
			}
			return r.expect + r.actual, nil
		}
		if strings.Contains(line, " IsNil") {
			r, err := IsNil(line)
			if err != nil {
				return "process equals has error", err
			}
			return r.expect + r.actual, nil
		}
		if strings.Contains(line, " NotNil") {
			r, err := NotNil(line)
			if err != nil {
				return "process equals has error", err
			}
			return r.expect + r.actual, nil
		}
		if strings.Contains(line, " IsTrue") {
			r, err := IsTrue(line)
			if err != nil {
				return "process equals has error", err
			}
			return r.expect + r.actual, nil
		}
		if strings.Contains(line, " IsFalse") {
			r, err := IsFalse(line)
			if err != nil {
				return "process equals has error", err
			}
			return r.expect + r.actual, nil
		}
		if strings.Contains(line, " Greater") {
			r, err := Greater(line)
			if err != nil {
				return "process equals has error", err
			}
			return r.expect + r.actual, nil
		}
	}
	return line, nil
}