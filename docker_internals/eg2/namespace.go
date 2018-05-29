package main

import (
    "os/exec"
    "syscall"
    "os"
    "log"
)

func main() {
    cmd := exec.Command("sh")
    cmd.SysProcAttr = &syscall.SysProcAttr{
        Cloneflags: syscall.CLONE_NEWUTS| syscall.CLONE_NEWIPC | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWNET,
        // user name space is not shiped with archlinux and also not required for docker as well for security reason
        //Cloneflags: syscall.CLONE_NEWUTS| syscall.CLONE_NEWIPC | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWUSER,
    }

    //cmd.SysProcAttr.Credential = &syscall.Credential{ Uid: uint32(1), Gid: uint32(1) }
    syscall.Mount("", "/", "", syscall.MS_PRIVATE|syscall.MS_REC, "")
    //defaultMountFlags := syscall.MS_NOEXEC | syscall.MS_NOSUID | syscall.MS_NODEV
    //syscall.Mount("proc", "/proc", "proc", uintptr(defaultMountFlags), "")
	//syscall.Mount("tmpfs", "/dev", "tmpfs", syscall.MS_NOSUID|syscall.MS_STRICTATIME, "mode=755")

    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    if err := cmd.Run(); err != nil {
        log.Fatal(err);
    }
    os.Exit(-1)

}
