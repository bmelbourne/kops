/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package install

import (
	"os/exec"

	"k8s.io/kops/upup/pkg/fi"
)

type InstallTarget struct {
}

var _ fi.InstallTarget = &InstallTarget{}

func (t *InstallTarget) Finish(taskMap map[string]fi.InstallTask) error {
	return nil
}

func (t *InstallTarget) DefaultCheckExisting() bool {
	return true
}

// CombinedOutput is a helper function that executes a command, returning stdout & stderr combined
func (t *InstallTarget) CombinedOutput(args []string) ([]byte, error) {
	c := exec.Command(args[0], args[1:]...)
	return c.CombinedOutput()
}
