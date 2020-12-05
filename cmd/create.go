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
	Short: "Creates monitoring components on your local machine.",
	Long: `Creates Prometheus, Grafana, Node Exporter, and cAdvisor on your local machine.
	It takes a few parameters to set up the tools versions and the Prometheus retention.
	Latests versions and 240 hours retention are seted by default.`,

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

		createConf(envMap)

		createFile(PrometheusVersion)

		params := []string{"-f", fileName, "--env-file", "config/local.env", "up", "-d"}
		runCommand(params...)

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

func createConf(m map[string]string) {
	f, err := os.Create("config/local.env")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	for k, v := range m {
		log.Println("Going to create", k, "=", v)
		d := fmt.Sprintln(k, "=", v)
		d1 := []byte(d)
		_, err2 := f.Write(d1)

		if err2 != nil {
			log.Fatal(err2)
		}
	}
}

func createFile(s string) string {

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

func runCommand(sos ...string) {
	dockerCom := exec.Command("docker-compose", sos...)
	dockerCom.Stdout = os.Stdout
	dockerCom.Stderr = os.Stderr
	err := dockerCom.Run()
	if err != nil {
		log.Fatal(err)
	}
}
