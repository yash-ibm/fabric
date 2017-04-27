/*
Copyright IBM Corp. 2016 All Rights Reserved.

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

package clilogging

import (
	"fmt"

	"github.com/hyperledger/fabric/common/flogging"

	"github.com/spf13/cobra"
)

const loggingFuncName = "logging"

var logger = flogging.MustGetLogger("cli/logging")

// Cmd returns the cobra command for Logging
func Cmd() *cobra.Command {
	loggingCmd.AddCommand(getLevelCmd())
	loggingCmd.AddCommand(setLevelCmd())
	loggingCmd.AddCommand(revertLevelsCmd())

	return loggingCmd
}

var loggingCmd = &cobra.Command{
	Use:   loggingFuncName,
	Short: fmt.Sprintf("%s specific commands.", loggingFuncName),
	Long:  fmt.Sprintf("%s specific commands.", loggingFuncName),
}
