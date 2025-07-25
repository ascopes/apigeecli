// Copyright 2024 Google LLC
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

package deployments

import (
	"github.com/spf13/cobra"
)

// DeploymentCmd to manage apis
var DeploymentCmd = &cobra.Command{
	Use:   "deployments",
	Short: "Manage details of the deployment where APIs are hosted",
	Long:  "Manage details of the deployment where APIs are hosted",
}

var org, region string

var examples = []string{
	`apigeecli apihub deployments create --dep-type apigee -d $dispName \
	--env-type development --slo-type "99-99" --endpoints https://api.example.com \
	--resource-uri https://apigee.googleapis.com/v1/organizations/$project/apis/$proxy/revisions/1 \
	-r us-central1 --default-token`,
}

func init() {
	DeploymentCmd.PersistentFlags().StringVarP(&org, "org", "o",
		"", "Apigee organization name")
	DeploymentCmd.PersistentFlags().StringVarP(&region, "region", "r",
		"", "API Hub region name")

	DeploymentCmd.AddCommand(CrtCmd)
	DeploymentCmd.AddCommand(GetCmd)
	DeploymentCmd.AddCommand(DelCmd)
	DeploymentCmd.AddCommand(ListCmd)
	DeploymentCmd.AddCommand(UpdateCmd)

	_ = DeploymentCmd.MarkFlagRequired("org")
	_ = DeploymentCmd.MarkFlagRequired("region")
}

func GetExample(i int) string {
	return examples[i]
}
