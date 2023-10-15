package main

import (
	"bytes"
	"fmt"
	"github.com/gookit/goutil/arrutil"
	"github.com/gookit/goutil/fsutil"
	markdown2 "github.com/shurcooL/markdownfmt/markdown"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"strings"
)

var confs *Conf
var outFolderName string
var oriFolderName string
var err error

func main() {
	confs, err = readConf()
	if err != nil {
		return
	}

	exists := fsutil.IsDir(confs.FileDir)
	if !exists {
		panic("the dir is not exist in conf, please check it")
	}
	fsutil.QuietRemove(filepath.Dir(confs.FileDir))
	oriFolderName = filepath.Base(confs.FileDir)
	outFolderName = oriFolderName + "_processed"
	//fsutil.Mkdir(filepath.Dir(confs.FileDir)+outFolderName, 0755)
	err := filepath.WalkDir(confs.FileDir, visit)
	if err != nil {
		return
	}
	return
}

func visit(path string, d os.DirEntry, err error) error {

	if d.IsDir() && arrutil.Contains(confs.Excludes, filepath.Base(path)) {
		return filepath.SkipDir
	} else if filepath.Ext(path) == ".md" {
		handleContent(path)
	}
	return err
}

func handleContent(path string) {

	file, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("read file fail", err)
		return
	}
	process, err := markdown2.Process(path, file, nil)
	if err != nil {
		fmt.Println("markdown format fail", err)
		return
	}
	filename := strings.Replace(path, oriFolderName, outFolderName, 1)
	err = os.MkdirAll(filepath.Dir(filename), 0755)
	if err != nil {
		fmt.Println("create new Dir fail", err)
		return
	}
	var buf = bytes.Buffer{}
	buf.Write([]byte("#" + strings.ReplaceAll(filepath.Base(filepath.Dir(path)), " ", "") + "\n"))
	buf.Write(process)
	err = os.WriteFile(filename, buf.Bytes(), 0755)
	if err != nil {
		fmt.Println("write file fail", err)
		recover()
		return
	}
}

func readConf() (*Conf, error) {
	var conf = Conf{}
	executable, err := os.Executable()
	if err != nil {
		return nil, err
	}
	dir := filepath.Dir(executable)
	file, err := os.ReadFile(dir + "./conf.yml")
	if err != nil {
		fmt.Println("read conf file fail", err)
		return nil, err
	}
	err = yaml.Unmarshal(file, &conf)
	if err != nil {
		fmt.Println("parse conf fail", err)
		return nil, err
	}
	return &conf, nil
}
