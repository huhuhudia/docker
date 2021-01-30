package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"syscall"
)

const cgroupMemoryHierachyMount=  "/sys/fs/cgroup/memory"
func main(){
	if os.Args[0] == "/proc/self/exe"{
		fmt.Printf("current pid :%v", syscall.Getpid())
		fmt.Println()

		cmd := exec.Command("sh", "-c", "stress --vm-bytes 1000m --vm-keep -m 1")
		cmd.SysProcAttr = &syscall.SysProcAttr{

		}
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil{
			fmt.Println(err)
			os.Exit(1)
		}
	}
	fmt.Println("start create cgroup limit ")
	cmd := exec.Command("/proc/self/exe")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS| syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil{
		log.Println("ERROR", err)
		os.Exit(1)
	}else{
		fmt.Printf("%v", cmd.Process.Pid)
		memlimit := "testmemlmit"
		os.Mkdir(path.Join(cgroupMemoryHierachyMount, memlimit), 0755)
		ioutil.WriteFile(path.Join(cgroupMemoryHierachyMount, memlimit, "tasks"), []byte(fmt.Sprint("%v", cmd.Process.Pid)), 0644)
		ioutil.WriteFile(path.Join(cgroupMemoryHierachyMount, memlimit, "memory.limit_in_bytes"), []byte("100m"), 0644)
		cmd.Process.Wait()
	}
}