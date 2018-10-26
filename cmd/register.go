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

	"github.com/spf13/cobra"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		
		username,_:=cmd.Flags().GetString("username")
	   if(username ==""){
		fmt.Println("Please input username[-u]")
		return
	   }
		fmt.Println("Register called by "+username)

		password,_:=cmd.Flags().GetString("password")
		emial, _ := cmd.Flags().GetString("email")
        phone, _:= cmd.Flags().GetString("phonenumber")
		if(password =="" ||emial==""||phone==""){
			fmt.Println("Please input password[-p] and email [-e] and phone [-n}" )
			return
		   }

		   
	},
}

func init() {
	RootCmd.AddCommand(registerCmd)
	 registerCmd.Flags().StringP("username", "u", "Anonymous", "Your username")
	 registerCmd.Flags().StringP("password", "p", "", "Your password to login")
	 registerCmd.Flags().StringP("email", "e", "","your email address")
     registerCmd.Flags().StringP("phonenumber","n", "","your cellphone number")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}