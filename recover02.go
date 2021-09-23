package main
import (
	"os"
	"io"
)

func CopyFile(dstName , srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
        defer src.Close()

	dst,err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()


	written, err = io.Copy(dst, src)
	return
}

func main() {
	CopyFile("/tmp/example-copy1.txt", "/tmp/example-copy.txt")
}

