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

	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Shows monitoring running containers",
	Long:  `Use this command to view all running containerd for your montool project.`,
	Run: func(cmd *cobra.Command, args []string) {
		PrometheusVersion, err := createCmd.Flags().GetString("prometheus-version")
		if err != nil {
			fmt.Println(err)
		}
		create(PrometheusVersion)
		dockerShow := exec.Command("docker-compose", "-f", fileName, "ps")
		dockerShow.Stdout = os.Stdout
		dockerShow.Stderr = os.Stderr
		err = dockerShow.Run()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}
