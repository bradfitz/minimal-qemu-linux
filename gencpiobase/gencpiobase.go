package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/u-root/u-root/pkg/cpio"
)

func getFiles() (ret []string) {
	set := map[string]bool{}

	var add func(string)
	add = func(f string) {
		if set[f] {
			return
		}
		log.Printf("add %q", f)
		fi, err := os.Lstat(f)
		if os.IsNotExist(err) {
			return
		}
		if err != nil {
			log.Fatal(err)
		}
		set[f] = true
		if fi.IsDir() {
			filepath.Walk(f, func(path string, fi os.FileInfo, err error) error {
				if err != nil {
					log.Fatal(err)
				}
				if path == f {
					return nil
				}
				add(path)
				return nil
			})
			return
		}
		if fi.Mode()&os.ModeSymlink != 0 {
			target, err := os.Readlink(f)
			if err != nil {
				log.Fatal(err)
			}
			if !filepath.IsAbs(target) {
				target = filepath.Join(filepath.Dir(f), target)
			}
			add(target)
			return
		}
		out, _ := exec.Command("ldd", f).Output()
		for _, f := range strings.Fields(string(out)) {
			if strings.HasPrefix(f, "/") {
				add(f)
			}
		}
	}

	add("/etc/ld.so.conf")
	add("/etc/ld.so.conf.d")
	add("/sbin/sfdisk")
	add("/sbin/mke2fs")
	add("/sbin/resize2fs")
	add("/bin/lsblk")
	return
}

func main() {
	files := getFiles()

	f, err := os.Create("/tmp/base.cpio")
	if err != nil {
		log.Fatal(err)
	}
	recw := cpio.Newc.Writer(f)
	for _, file := range files {
		rec, err := cpio.GetRecord(file)
		if err != nil {
			log.Fatalf("GetRecord(%q): %v", file, err)
		}
		rec.Info.Name = cpio.Normalize(rec.Info.Name)
		if err := recw.WriteRecord(rec); err != nil {
			log.Fatal(err)
		}
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
