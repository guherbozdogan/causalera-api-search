package servicetest

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	//	"io/ioutil"
	"log"
	"os"
	"os/exec"
	//	"strconv"
	"syscall"
	"testing"
	"time"
)

func execCmd(command string) (*exec.Cmd, error) {
	cmd := exec.Command(os.Getenv("SHELL"), "-c", command)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{}
	//	cmd.SysProcAttr.Credential = &syscall.Credential{Uid: uint32(user.UID), Gid: uint32(user.GID)}
	if err := cmd.Start(); err != nil {
		return nil, err
	}
	return cmd, nil
}

var cmd *exec.Cmd
var err error
var _ = BeforeSuite(func() {
	fmt.Println("whatsup")
	cmd, err = execCmd("cd $GOPATH/src/github.com/guherbozdogan/causalera-api-search;go install main/main.go; cd  $GOPATH/bin; ./main & echo  $! > outtmp")

	if err != nil {
		log.Fatal(err)
	}
})

var _ = AfterSuite(func() {
	cmd := fmt.Sprintf("cd $GOPATH/bin; cat outtmp | xargs kill -9 ")
	fmt.Println(cmd)
	execCmd(cmd)
	time.Sleep(1400 * time.Millisecond)
	//fmt.Println("proc is %d", cmd.Process.Pid)
})

func TestFrame(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Service Test Suite")
}
