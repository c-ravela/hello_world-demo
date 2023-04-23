//go:build ignore

#include "vmlinux.h"
#include "bpf_helpers.h"

SEC("tp/syscalls/sys_enter_execve")
void execve(){
    bpf_printk("Hello World! I am triggered by enter point of execve.");
};

char _license[] SEC("license") = "Dual MIT/GPL";