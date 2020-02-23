package main

import (
	"fmt"
	"io/ioutil"
)

func GetAllFile(pathname string) error {
	rd, err := ioutil.ReadDir(pathname)
	for _, fi := range rd {
		if fi.IsDir() {
			fmt.Printf("[%s]\n", pathname+"\\"+fi.Name())
			GetAllFile(pathname + fi.Name() + "\\")
		} else {
			fmt.Println(fi.Name())
		}
	}
	return err
}

func main() {
	fmt.Print("hello word !")
	GetAllFile("C:/Users/z/Desktop/temp/1")

}

/*func deleteFileOnDisk(localPath string) {
	debugPrint("remove file: %s", localPath)
	if err := os.Remove(localPath); err != nil {
		errorPrintErr(err)
	}
	dirsList := make([]string, 0, 0)
	for dir := path.Dir(localPath); dir != svr.dataDir && len(dir) > len(svr.dataDir); dir = path.Dir(dir) {
		dirsList = append(dirsList, dir)
	}
	sort.StringSlice(dirsList).Sort()
	for i := len(dirsList) - 1; i >= 0; i-- {
		f, err := os.Open(dirsList[i])
		if err != nil {
			errorPrintErr(err)
		}
		fs, err2 := f.Readdirnames(1)
		if err2 == io.EOF && (fs == nil || len(fs) == 0) {
			f.Close()
			debugPrint("remove dir: %s", dirsList[i])
			if err := os.Remove(dirsList[i]); err != nil {
				errorPrintErr(err)
			}
			continue
		} else if err2 != nil {
			errorPrintErr(err2)
		}
		f.Close()
	}

}

func errorPrintErr(err error) {
	fmt.Print(err)
}

func debugPrint(s1 string, s2 string) {
	fmt.Print(s1 + s2)
}

func errorPrint(s1 string) {
	fmt.Print(s1)
}*/
