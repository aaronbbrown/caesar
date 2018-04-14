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
	"runtime"

	"github.com/aaronbbrown/caesar/pkg/cracker"
	"github.com/spf13/cobra"
)

var (
	maxKeys      *int
	percentWords *int
	parallelism  *int
)

// crackCmd represents the crack command
var crackCmd = &cobra.Command{
	Use:   "crack",
	Short: "Crack mode",
	Long:  `Crack a Caesar cipher`,
	Run: func(cmd *cobra.Command, args []string) {
		data, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}

		dictionary, err := cracker.NewDictionary("wordlist.txt")
		if err != nil {
			log.Fatal(err)
		}

		c := cracker.NewCracker(string(data), dictionary)

		decrypted, err := c.Crack(1, *maxKeys, *percentWords, *parallelism)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Key: ", decrypted.Key, decrypted.Offsets)
		fmt.Printf("Message:\n%s", decrypted.Msg)
	},
}

func init() {
	rootCmd.AddCommand(crackCmd)
	maxKeys = crackCmd.Flags().Int("max-attempts", 1000, "maximum number of attempts to crack")
	percentWords = crackCmd.Flags().Int("percent-words", 30, "minimum percentage of strings that must be words in the dictionary")
	parallelism = crackCmd.Flags().Int("parallelism", runtime.GOMAXPROCS(0), "Number of parallel threads to run during cracking")

	crackCmd.MarkFlagRequired("max-attempts")
	crackCmd.MarkFlagRequired("percent-words")
}
