package chapter5

import "os/exec"

type IntermediateError struct {
	error
}

func runJob(id string) error {
	const jobBinPath = "/bad/job/binary"
	isExecutable, err := isGlobalExec(jobBinPath)

	if err != nil {
		return IntermediateError{wrap(err, "cannot run job %q: requiste binaries not available", id)}
	}
	if isExecutable == false {
		return wrap(nil, "job binary is not executable")
	}

	return exec.Command(jobBinPath, "--id="+id).Run()
}
