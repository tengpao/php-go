package php_go

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func Mkdir(Path string) {
	p, _ := path.Split(Path)
	if p == "" {
		return
	}
	d, err := os.Stat(p)
	if err != nil || !d.IsDir() {
		if err = os.MkdirAll(p, 0777); err != nil {
			fmt.Printf("创建路径失败[%v]: %v\n", Path, err)
		}
	}
}

func IsDirExists(path string) bool {
	fi, err := os.Stat(path)

	if err != nil {
		return os.IsExist(err)
	} else {
		return fi.IsDir()
	}

	panic("util isDirExists not reached")
}

func IsFileExists(path string) bool {
	fi, err := os.Stat(path)

	if err != nil {
		return os.IsExist(err)
	} else {
		return !fi.IsDir()
	}

	panic("util isFileExists not reached")
}

func WalkFiles(targpath string, suffixes ...string) (filelist []string) {
	if !filepath.IsAbs(targpath) {
		targpath, _ = filepath.Abs(targpath)
	}
	err := filepath.Walk(targpath, func(retpath string, f os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		if len(suffixes) == 0 {
			filelist = append(filelist, retpath)
			return nil
		}
		for _, suffix := range suffixes {
			if strings.HasSuffix(retpath, suffix) {
				filelist = append(filelist, retpath)
			}
		}
		return nil
	})

	if err != nil {
		fmt.Printf("util.WalkFiles: %v\n", err)
		return
	}

	return
}

// 遍历目录，可指定后缀
func WalkDir(targpath string, suffixes ...string) (dirlist []string) {
	if !filepath.IsAbs(targpath) {
		targpath, _ = filepath.Abs(targpath)
	}
	err := filepath.Walk(targpath, func(retpath string, f os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !f.IsDir() {
			return nil
		}
		if len(suffixes) == 0 {
			dirlist = append(dirlist, retpath)
			return nil
		}
		for _, suffix := range suffixes {
			if strings.HasSuffix(retpath, suffix) {
				dirlist = append(dirlist, retpath)
			}
		}
		return nil
	})

	if err != nil {
		fmt.Printf("util.WalkDir: %v\n", err)
		return
	}

	return
}

// 遍历文件，可指定后缀，返回相对路径
func WalkRelFiles(targpath string, suffixes ...string) (filelist []string) {
	if !filepath.IsAbs(targpath) {
		targpath, _ = filepath.Abs(targpath)
	}
	err := filepath.Walk(targpath, func(retpath string, f os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		if len(suffixes) == 0 {
			filelist = append(filelist, RelPath(retpath))
			return nil
		}
		_retpath := RelPath(retpath)
		for _, suffix := range suffixes {
			if strings.HasSuffix(_retpath, suffix) {
				filelist = append(filelist, _retpath)
			}
		}
		return nil
	})

	if err != nil {
		fmt.Printf("util.WalkRelFiles: %v\n", err)
		return
	}

	return
}

// 遍历目录，可指定后缀，返回相对路径
func WalkRelDir(targpath string, suffixes ...string) (dirlist []string) {
	if !filepath.IsAbs(targpath) {
		targpath, _ = filepath.Abs(targpath)
	}
	err := filepath.Walk(targpath, func(retpath string, f os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !f.IsDir() {
			return nil
		}
		if len(suffixes) == 0 {
			dirlist = append(dirlist, RelPath(retpath))
			return nil
		}
		_retpath := RelPath(retpath)
		for _, suffix := range suffixes {
			if strings.HasSuffix(_retpath, suffix) {
				dirlist = append(dirlist, _retpath)
			}
		}
		return nil
	})

	if err != nil {
		fmt.Printf("util.WalkRelDir: %v\n", err)
		return
	}

	return
}

// 转相对路径
func RelPath(targpath string) string {
	basepath, _ := filepath.Abs("./")
	rel, _ := filepath.Rel(basepath, targpath)
	return strings.Replace(rel, `\`, `/`, -1)
}
