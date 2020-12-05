// Copyright © 2020 NAME HERE <EMAIL ADDRESS>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Spin up monitoring environment on your local computer.",
	Long: `Takes a few parameters to set up the tools versions and retention and create docker containers on your local computer.
	Latests versions and 240 hours retention is the default.`,

	Run: func(cmd *cobra.Command, args []string) {

		PrometheusVersion, err := cmd.Flags().GetString("prometheus-version")
		if err != nil {
			fmt.Println(err)
		}
		PrometheusRetention, err := cmd.Flags().GetString("prometheus-retention")
		if err != nil {
			fmt.Println(err)
		}
		NodeVersion, err := cmd.Flags().GetString("node-version")
		if err != nil {
			fmt.Println(err)
		}
		GrafanaVersion, err := cmd.Flags().GetString("grafana-version")
		if err != nil {
			fmt.Println(err)
		}

		envMap := map[string]string{
			"PrometheusVersion":   PrometheusVersion,
			"PrometheusRetention": PrometheusRetention,
			"NodeVersion":         NodeVersion,
			"GrafanaVersion":      GrafanaVersion,
		}
		{
			f, err := os.Create("config/.env")

			if err != nil {
				log.Fatal(err)
			}

			defer f.Close()

			for k, v := range envMap {
				fmt.Println("Going to create", k, "=", v)
				d := fmt.Sprintln(k, "=", v)
				d1 := []byte(d)
				_, err2 := f.Write(d1)

				if err2 != nil {
					log.Fatal(err2)
				}
			}
		}

		create(PrometheusVersion)
		dockerCom := exec.Command("docker-compose", "-f", fileName, "--env-file", "config/.env", "up", "-d")
		dockerCom.Stdout = os.Stdout
		dockerCom.Stderr = os.Stderr
		err = dockerCom.Run()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().StringP("prometheus-version", "p", "latest", "prometheus version to deploy")
	createCmd.Flags().StringP("prometheus-retention", "r", "240h", "prometheus retention in hours")
	createCmd.Flags().StringP("node-version", "n", "latest", "node exporter version to deploy")
	createCmd.Flags().StringP("grafana-version", "g", "latest", "grafana version to deploy")
}

var fileName string

func create(s string) string {

	fileName = "docker-compose-v2.yml"
	matched, err := regexp.Match("v1.*", []byte(s))
	if err != nil {
		fmt.Println(err)
	}
	if matched == true {
		fileName = "docker-compose-v1.yml"
	}

	return fileName
}
