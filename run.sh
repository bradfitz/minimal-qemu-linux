#!/bin/bash

echo "Type Ctrl-A x to quit qemu; Ctrl-A c to enter qemu monitor..."

set -e
set -x

rm -f disk.qcow2 || true
qemu-img create -f qcow2 disk.qcow2 1G

qemu-system-x86_64 \
    -vga none \
    -device virtio-serial \
    -device virtconsole \
    -device virtio-net,netdev=net0 \
    -device virtio-scsi-pci,id=scsi \
    -device scsi-hd,drive=hd \
    -drive file=disk.qcow2,media=disk,if=none,id=hd,format=qcow2 \
    -netdev user,id=net0 \
    -nographic \
    -display none \
    -m 1024 \
    -kernel $HOME/hack/linux/arch/x86/boot/bzImage  \
    -initrd "/tmp/initramfs.linux_amd64.cpio" \
    -no-reboot \
    -append "console=ttyS0,115200 panic=-1 ip=dhcp"
