package main

//go:generate go run github.com/cilium/ebpf/cmd/bpf2go -cc clang bpf execve.c -- -I../headers

import (
	"C"
	"fmt"

	"github.com/cilium/ebpf/link"
)

func main() {
	ebpfObj := bpfObjects{}
	err := loadBpfObjects(&ebpfObj, nil)
	if err != nil {
		panic(err)
	}
	defer ebpfObj.Close()

	hook, err := link.Tracepoint("syscalls", "sys_enter_execve", ebpfObj.Execve, nil)
	if err != nil {
		panic(err)
	}
	defer hook.Close()

	fmt.Println("Waiting for event to trigger!")
	for {
	}
}
