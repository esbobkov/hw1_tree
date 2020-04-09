package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
)

func main() {

	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}

func dirTree(out io.Writer, path string, printFiles bool) error {

	walk(out, path, printFiles, "")

	return nil
}

func walk(out io.Writer, path string, printFiles bool, prefix string) {
	file := openPath(path)
	fileInfo, _ := file.Readdir(-1)
	fileInfo = orderFileInfo(fileInfo)
	i := 0

	for i < len(fileInfo) {
		isDir := true == fileInfo[i].IsDir()
		isLastDir := isLastFile(path+string(os.PathSeparator)+fileInfo[i].Name(), printFiles)

		if true == isLastDir {
			if true == printFiles && false == isDir {
				if 0 == fileInfo[i].Size() {
					_, err := fmt.Fprintf(out, "%v (empty)\n", prefix+"└───"+fileInfo[i].Name())
					if err != nil {
						panic(err.Error())
					}
				} else {
					_, err := fmt.Fprintf(out, "%v (%db)\n", prefix+"└───"+fileInfo[i].Name(), fileInfo[i].Size())
					if err != nil {
						panic(err.Error())
					}
				}
			}
			if true == isDir {
				_, err := fmt.Fprintln(out, prefix+"└───"+fileInfo[i].Name())
				if err != nil {
					panic(err.Error())
				}
				walk(out, path+string(os.PathSeparator)+fileInfo[i].Name(), printFiles, prefix+"\t")
			}
		} else {
			if true == printFiles && false == isDir {
				if 0 == fileInfo[i].Size() {
					_, err := fmt.Fprintf(out, "%v (empty)\n", prefix+"├───"+fileInfo[i].Name())
					if err != nil {
						panic(err.Error())
					}
				} else {
					_, err := fmt.Fprintf(out, "%v (%db)\n", prefix+"├───"+fileInfo[i].Name(), fileInfo[i].Size())
					if err != nil {
						panic(err.Error())
					}
				}
			}
			if true == isDir {
				_, err := fmt.Fprintln(out, prefix+"├───"+fileInfo[i].Name())
				if err != nil {
					panic(err.Error())
				}
				walk(out, path+string(os.PathSeparator)+fileInfo[i].Name(), printFiles, prefix+"│\t")
			}
		}
		i++
	}
}

func openPath(path string) *os.File {
	file, _ := os.Open(path)
	return file
}

func orderFileInfo(fileInfo []os.FileInfo) []os.FileInfo {
	sort.Slice(fileInfo, func(i, j int) bool {
		return fileInfo[i].Name() < fileInfo[j].Name()
	})

	return fileInfo
}

func isLastFile(path string, printFiles bool) bool {

	basePath := filepath.Base(path)

	var sortList []string

	files, _ := ioutil.ReadDir(filepath.Dir(path))

	for _, file := range files {
		if printFiles == false && file.IsDir() == false {
			continue
		}
		sortList = append(sortList, file.Name())
	}

	if nil != sortList && sortList[len(sortList)-1] == basePath {
		return true
	}

	return false
}
