package pkg

import (
	_ "bufio"
	_ "fmt"
	_ "image/color"
	"io"
	_ "log"
	_ "os"
	"os/exec"
	_ "time"
)

const TIMEOUT = 5

type Process struct {
	proc   *exec.Cmd
	stdout io.ReadCloser
	stdin  io.WriteCloser
	stderr io.ReadCloser
	status int
}

func (p *Process) StartProcess(cmd ...string) {
	proc := exec.Command(cmd[0], cmd[1:]...)
	p.stdout, _ = proc.StdoutPipe()
	p.stdin, _ = proc.StdinPipe()
	p.stderr, _ = proc.StderrPipe()

	p.status = 0
}

func (p *Process) sendafter(delim []byte, data []byte) {

}

func (p *Process) sendlineafter(delim []byte, data []byte) {

}

func (p *Process) recv(size uint64) []byte {
	buf := make([]byte, size)
	_, _ = p.stdout.Read(buf)

	return buf
}

func (p *Process) recvuntil(buf []byte) {

}

func (p *Process) send(date []byte) {
	_, _ = p.stdin.Write(date)
}
