package migrate

import "os"

func IsExists(path string) (os.FileInfo, bool) {
	f, err := os.Stat(path)
	// 如果当前用户没有访问该文件的权限时，err 也有值，所以需要使用 os.IsExist() 方法加一层判断
	return f, err == nil || os.IsExist(err)
}

func IsDir(path string) (os.FileInfo, bool) {
	f, flag := IsExists(path)
	return f, flag && f.IsDir()
}

func IsFile(path string) (os.FileInfo, bool) {
	f, flag := IsExists(path)
	return f, flag && !f.IsDir()
}
