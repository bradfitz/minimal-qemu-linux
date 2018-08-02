#!/bin/bash

echo "Type Ctrl-A x to quit qemu; Ctrl-A c to enter qemu monitor..."

set -e
set -x

qemu-system-x86_64 \
    -device virtio-serial \
    -device virtconsole \
    -nographic \
    -display none \
    -net none \
    -m 1024 \
    -kernel $HOME/hack/linux/arch/x86/boot/bzImage  \
    -append "console=ttyS0,115200"
