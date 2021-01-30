package container
import(
	log "github.com/sirupsen/logrus"
	"os"
	"syscall"
)
func RunContainerInitProcess(command string, args []string)error{
	log.Infof("RunContainerInitProcess command :%v", command)
	defaultMountFlags := syscall.MS_NOEXEC | syscall.MS_NOSUID | syscall.MS_NODEV
	syscall.Mount("proc", "/proc", "proc", uintptr(defaultMountFlags), "")
	argv := []string{command}
	if err := syscall.Exec(command, argv, os.Environ()); err != nil{
		log.Errorln(err)
	}
	return nil
}