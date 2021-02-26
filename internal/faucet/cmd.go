package faucet

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os/exec"
	"strings"
)

func cmdexec(nodeAddr string, bin string, args []string, inputs ...string) (string, error) {
	//always on test network so hard coding :shrug:
	args = append([]string{"-t"}, args...)
	if len(nodeAddr) > 0 {
		args = append([]string{"--node=tcp://" + nodeAddr + ":26657"}, args...)
	}
	cmd := exec.Command(bin, args...)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return "", err
	}

	//stdin stuff here
	for _, input := range inputs {
		if _, err := fmt.Fprintln(stdin, input); err != nil {
			return "", err
		}
	}

	//double check this is right
	go func() {
		defer stdin.Close()
	}()

	out, err := cmd.CombinedOutput()
	//see what the cmd output is, exit status 1 doesn't help when things go wrong
	log.Info(string(out))

	if err != nil {
		return "", fmt.Errorf("error executing command: %s", err)
	}

	return strings.TrimSpace(string(out)), err
}

func (f *Faucet) cliexec(nodeAddr string , args []string, inputs ...string) (string, error) {
	return cmdexec(nodeAddr, f.appCli, args, inputs...)
}
