#!/bin/bash

echo "Type Ctrl-A x to quit qemu; Ctrl-A c to enter qemu monitor..."

set -e
set -x

qemu-system-x86_64 \
    -device virtio-serial \
    -m 1024 \
    -nographic -vga none \
    -serial mon:stdio \
    -kernel $HOME/hack/linux/arch/x86/boot/bzImage \
    -append "console=ttyS0"
