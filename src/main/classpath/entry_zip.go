package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

/**
 * zip/jar 文件形式的类路径
 */
type ZipEntry struct {
	absPath string
	// 用于优化readClass
	zipRC *zip.ReadCloser
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath, nil}
}

func (self *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	if self.zipRC == nil {
		err := self.openJar()
		if err != nil {
			return nil, nil, err
		}
	}

	classFile := self.findClass(className)
	if classFile == nil {
		return nil, nil, errors.New("class not found: " + className)
	}

	data, err := readClass(classFile)
	return data, self, err
}

func readClass(classFile *zip.File) ([]byte, error) {
	rc, err := classFile.Open()
	if err != nil {
		return nil, err
	}
	// read class data
	data, err := ioutil.ReadAll(rc)
	rc.Close()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (self *ZipEntry) findClass(className string) *zip.File {
	for _, f := range self.zipRC.File {
		if f.Name == className {
			return f
		}
	}
	return nil
}

func (self *ZipEntry) openJar() error {
	r, err := zip.OpenReader(self.absPath)
	if err == nil {
		self.zipRC = r
	}
	return err
}

// 从 zip文件中 提取 class 文件
func (self *ZipEntry) readClass2(className string) ([]byte, Entry, error) {
	// 1. open zip
	r, err := zip.OpenReader(self.absPath)
	if err != nil {
		return nil, nil, err
	}
	defer r.Close()

	// 2. 遍历
	for _, f := range r.File {
		if f.Name != className {
			continue
		}
		// 3. 找到了 class
		rc, err := f.Open()
		if err != nil {
			return nil, nil, err
		}
		defer rc.Close()
		// 4. 读取信息，返回
		data, err := ioutil.ReadAll(rc)
		if err != nil {
			return nil, nil, err
		}
		return data, self, nil
	}
	return nil, nil, errors.New("class not found: " + className)
}

// toString
func (self *ZipEntry) String() string {
	return self.absPath
}
