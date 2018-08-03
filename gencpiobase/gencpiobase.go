package main

import (
	"log"
	"os"

	"github.com/u-root/u-root/pkg/cpio"
)

// TODO: compute this all at runtime from the roots
// of "lsblk", "mk2fs", and "sfdisk" only.
var files = []string{
	"/bin/lsblk",
	"/lib/x86_64-linux-gnu/libblkid.so.1",
	"/lib/x86_64-linux-gnu/libblkid.so.1.1.0",
	"/lib/x86_64-linux-gnu/libc.so.6",
	"/lib/x86_64-linux-gnu/libc-2.24.so",
	"/lib/x86_64-linux-gnu/libcom_err.so.2",
	"/lib/x86_64-linux-gnu/libcom_err.so.2.1",
	"/lib/x86_64-linux-gnu/libdl.so.2",
	"/lib/x86_64-linux-gnu/libdl-2.24.so",
	"/lib/x86_64-linux-gnu/libe2p.so.2",
	"/lib/x86_64-linux-gnu/libe2p.so.2.3",
	"/lib/x86_64-linux-gnu/libext2fs.so.2",
	"/lib/x86_64-linux-gnu/libext2fs.so.2.4",
	"/lib/x86_64-linux-gnu/libfdisk.so.1",
	"/lib/x86_64-linux-gnu/libfdisk.so.1.1.0",
	"/lib/x86_64-linux-gnu/libpthread.so.0",
	"/lib/x86_64-linux-gnu/libpthread-2.24.so",
	"/lib/x86_64-linux-gnu/libsmartcols.so.1",
	"/lib/x86_64-linux-gnu/libsmartcols.so.1.1.0",
	"/lib/x86_64-linux-gnu/libtinfo.so.5",
	"/lib/x86_64-linux-gnu/libtinfo.so.5.9",
	"/lib/x86_64-linux-gnu/libuuid.so.1",
	"/lib/x86_64-linux-gnu/libuuid.so.1.3.0",
	"/lib/x86_64-linux-gnu/ld-2.24.so",
	"/lib/x86_64-linux-gnu/libmount.so.1",
	"/lib/x86_64-linux-gnu/libmount.so.1.1.0",
	"/lib/x86_64-linux-gnu/libselinux.so.1",
	"/lib/x86_64-linux-gnu/librt.so.1",
	"/lib/x86_64-linux-gnu/librt-2.24.so",
	"/lib/x86_64-linux-gnu/libudev.so.1",
	"/lib/x86_64-linux-gnu/libudev.so.1.6.5",
	"/lib/x86_64-linux-gnu/libpcre.so.3",
	"/lib/x86_64-linux-gnu/libpcre.so.3.13.3",
	"/lib64/ld-linux-x86-64.so.2",
	"/sbin/mke2fs",
	"/sbin/sfdisk",
	"/etc/ld.so.conf",
	// TODO: filepath walk /etc/ld.so.conf.d
	"/etc/ld.so.conf.d/x86_64-linux-gnu.conf",
	"/etc/ld.so.conf.d/libc.conf",
}

func main() {
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
