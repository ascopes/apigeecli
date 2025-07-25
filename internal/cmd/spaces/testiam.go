// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package spaces

import (
	"internal/apiclient"
	"internal/client/spaces"

	"github.com/spf13/cobra"
)

// TestIamCmd to manage tracing of apis
var TestIamCmd = &cobra.Command{
	Use:   "test",
	Short: "Test IAM policy for a Space",
	Long:  "Test IAM policy for a Space",
	Args: func(cmd *cobra.Command, args []string) (err error) {
		apiclient.SetRegion(region)
		return apiclient.SetApigeeOrg(org)
	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		cmd.SilenceUsage = true

		_, err = spaces.TestIAM(space, resource, verb)
		return
	},
	Example: `Test IAM for a space:
` + GetExample(4) + `
` + GetExample(5),
}

var verb, resource string

func init() {
	TestIamCmd.Flags().StringVarP(&space, "space", "",
		"", "Space name.")
	TestIamCmd.Flags().StringVarP(&verb, "verb", "v",
		"get", "resource verb")
	TestIamCmd.Flags().StringVarP(&resource, "res", "s",
		"proxies", "resource")

	_ = TestIamCmd.MarkFlagRequired("space")
}
