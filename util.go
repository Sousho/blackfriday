package main

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func e(err error) {
	if err != nil {
		Err.Println(err)
		log.Fatal(err)
	}
}

func ReadBytes(a os.File) []byte {
	fileinfo, err := a.Stat()
	e(err)
	filesize := fileinfo.Size()
	read_buffer := make([]byte, filesize)
	a.Read(read_buffer)
	e(err)
	return read_buffer
}

func find_ext(root, ext string) string {
	result := ""
	filepath.WalkDir(root, func(s string, d fs.DirEntry, e error) error {
		if e != nil {
			return e
		}
		if filepath.Ext(d.Name()) == ext {
			result = d.Name()
			return nil
		}
		return nil
	})
	return result
}
