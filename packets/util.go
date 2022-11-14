package packets

import (
	"github.com/ddkwork/golibrary/mylog"
	"os"
	"os/user"
	"path/filepath"
)

func JoinHomeDir(path string) (join string, ok bool)  { return joinHome(path, true) }
func JoinHomeFile(path string) (join string, ok bool) { return joinHome(path, false) }
func joinHome(path string, isDir bool) (join string, ok bool) {
	join = filepath.Join(HomeDir(), path)
	if !FileExists(join) {
		switch isDir {
		case true:
			if !mylog.Error(os.MkdirAll(join, 0777)) {
				return
			}
		default:
			f, err := os.Create(join)
			if !mylog.Error(err) {
				return
			}
			if !mylog.Error(f.Close()) {
				return
			}
		}
	}
	ok = true
	return
}

func HomeDir() string {
	if u, err := user.Current(); err == nil {
		return u.HomeDir
	}
	if dir, err := os.UserHomeDir(); err == nil {
		return dir
	}
	return "."
}

func FileExists(path string) bool {
	if fi, err := os.Stat(path); err == nil {
		mode := fi.Mode()
		return !mode.IsDir() && mode.IsRegular()
	}
	return false
}

func IsDir(path string) bool {
	fi, err := os.Stat(path)
	return err == nil && fi.IsDir()
}

func BaseName(path string) string      { return TrimExtension(filepath.Base(path)) }
func TrimExtension(path string) string { return path[:len(path)-len(filepath.Ext(path))] }
