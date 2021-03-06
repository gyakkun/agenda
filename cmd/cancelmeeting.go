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
	"agenda/service"
	"fmt"

	"github.com/spf13/cobra"
)

// cancelmeetingCmd represents the cancelmeeting command
var cancelmeetingCmd = &cobra.Command{
	Use:   "cancelmeeting",
	Short: "Cancel a meeting",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cancelmeeting called")

		title, _ := cmd.Flags().GetString("title")
		if title == "" {
			fmt.Println("Please input the title")
			return
		}
		user, flag := service.GetCurUser()
		if flag == false {
			fmt.Println("Please login!")
		} else {
			flag2 := service.DeleteMeeting(user.Name, title)
			if flag2 == 0 {
				fmt.Println("Error! You're not a sponsor of the meeting")
			} else {
				fmt.Println("Delete Successfully")
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(cancelmeetingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cancelmeetingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cancelmeetingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
