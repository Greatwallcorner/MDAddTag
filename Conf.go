package main

type Conf struct {
	Excludes []string `yaml:"excludes"`
	FileDir  string   `yaml:"fileDir"`
}
