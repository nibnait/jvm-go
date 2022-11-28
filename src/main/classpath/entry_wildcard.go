package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

/**
 * 通配符类路径
 */
func newWildcardEntry(path string) CompositeEntry {
	// 把路径末尾的 * 去掉
	baseDir := path[:len(path)-1]
	compositeEntry := []Entry{}
	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && path != baseDir {
			// 通配符类路径 不能递归匹配子目录下的 JAR 文件
			return filepath.SkipDir
		}
		// 根据后缀名选出 JAR 文件
		if strings.HasSuffix(path, "jar") || strings.HasSuffix(path, "JAR") {
			jarEntry := newZipEntry(path)
			compositeEntry = append(compositeEntry, jarEntry)
		}
		return nil
	}

	// 遍历 basDir，创建 ZipEntry
	filepath.Walk(baseDir, walkFn)
	return compositeEntry
}
