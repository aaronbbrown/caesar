// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/aaronbbrown/caesar/pkg/cipher"
	"github.com/spf13/cobra"
)

var encryptKey *int

// encryptCmd represents the encrypt command
var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "Encrypt mode",
	Long:  `Encrypt a string using the Caesar cipher.`,
	Run: func(cmd *cobra.Command, args []string) {
		data, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}

		caesar := cipher.NewCaesar(string(data), *encryptKey)
		encrypted := caesar.Encrypt()
		fmt.Println(encrypted)
	},
}

func init() {
	rootCmd.AddCommand(encryptCmd)
	encryptKey = encryptCmd.Flags().Int("key", 0, "The encryption key (ignored in crack mode)")
	encryptCmd.MarkFlagRequired("key")
}
