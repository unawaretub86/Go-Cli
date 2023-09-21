/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
package cmd

import (
	"github.com/culturadevops/jaivic/services"
	"github.com/spf13/cobra"
)

// ctlCmd represents the ctl command
var ctlCmd = &cobra.Command{
	Use:   "ctl",
	Short: "Crea un archivo controllers usa -c para crear el crud",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if flags, _ := cmd.Flags().GetBool("crud"); flags {
			services.VarSrv.CreateCtlCrud(args[0])
		} else {
			services.VarSrv.CreateCtlDefault(args[0])
		}

	},
}

func init() {
	ctlCmd.Flags().BoolP("crud", "c", false, "Crea una plantilla crud")

	mkCmd.AddCommand(ctlCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ctlCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ctlCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
