package main

import (
	"bytes"
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

func main() {
	confs = readConf()
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
		println("read file fail", err)
		return
	}
	process, err := markdown2.Process(path, file, nil)
	if err != nil {
		println("markdown format fail", err)
		return
	}
	filename := strings.Replace(path, oriFolderName, outFolderName, 1)
	err = os.MkdirAll(filepath.Dir(filename), 0755)
	if err != nil {
		println("create new Dir fail", err)
		return
	}
	var buf = bytes.Buffer{}
	buf.Write([]byte("#" + strings.ReplaceAll(filepath.Base(filepath.Dir(path)), " ", "") + "\n"))
	buf.Write(process)
	err = os.WriteFile(filename, buf.Bytes(), 0755)
	if err != nil {
		println("write file fail", err)
		recover()
		return
	}
	//println(process)
}

func readConf() *Conf {
	var conf = Conf{}
	executable, err := os.Executable()
	if err != nil {
		return &Conf{}
	}
	dir := filepath.Dir(executable)
	file, err := os.ReadFile(dir + "./conf.yml")
	if err != nil {
		println("read conf file fail "+
			"", err)
		return &Conf{}
	}
	err = yaml.Unmarshal(file, &conf)
	if err != nil {
		println("parse conf fail", err)
		return &Conf{}
	}
	return &conf
}
