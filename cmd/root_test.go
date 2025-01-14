/*
   Copyright 2022 Ryan SVIHLA

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

// cmd package contains all the command line flag and initialization logic for commands
package cmd

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/rsvihladremio/dremio-diagnostic-collector/collection"
)

func TestSSHDefault(t *testing.T) {
	sshPath, err := sshDefault()
	if err != nil {
		t.Fatalf("unexpected exception %v", err)
	}

	expectedPath := filepath.Join(".ssh", "id_rsa")
	if !strings.HasSuffix(sshPath, expectedPath) {
		t.Errorf("expected %v but was %v", expectedPath, sshPath)
	}
}

func TestValidateParameters(t *testing.T) {
	tc := makeTestCollection()
	tc.CoordinatorStr = ""
	err := validateParameters(tc, "/home/dremio/.ssh", "dremio", true)
	expectedError := "the coordinator string was empty you must pass a namespace, a colon and a label that will match your coordinators --coordinator or -c arguments. Example: -c \"default:mylabel=coordinator\""
	if expectedError != err.Error() {
		t.Errorf("expected: %v but was %v", expectedError, err.Error())
	}

	tc = makeTestCollection()
	tc.ExecutorsStr = ""
	err = validateParameters(tc, "/home/dremio/.ssh", "dremio", true)
	expectedError = "the executor string was empty you must pass a namespace, a colon and a label that will match your executors --executor or -e arguments. Example: -e \"default:mylabel=executor\""
	if expectedError != err.Error() {
		t.Errorf("expected: %v but was %v", expectedError, err.Error())
	}

	tc = makeTestCollection()
	err = validateParameters(tc, "", "dremio", false)
	expectedError = "the ssh private key location was empty, pass --ssh-key or -s with the key to get past this error. Example --ssh-key ~/.ssh/id_rsa"
	if expectedError != err.Error() {
		t.Errorf("expected: %v but was %v", expectedError, err.Error())
	}

	tc = makeTestCollection()
	err = validateParameters(tc, "/home/dremio/.ssh", "", false)
	expectedError = "the ssh user was empty, pass --ssh-user or -u with the user name you want to use to get past this error. Example --ssh-user ubuntu"

	if expectedError != err.Error() {
		t.Errorf("expected: %v but was %v", expectedError, err.Error())
	}
}

// Set of args for other tests
func makeTestCollection() collection.Args {
	testCollection := collection.Args{
		CoordinatorStr:            "dremio-master-0",
		ExecutorsStr:              "dremio-executor-0",
		OutputLoc:                 "/tmp/diags",
		DremioConfDir:             "/opt/dremio/conf",
		DremioLogDir:              "/var/log/dremio",
		DurationDiagnosticTooling: 10,
		LogAge:                    5,
	}
	return testCollection
}
