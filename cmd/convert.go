package cmd

import (
	"fmt"
	"os"
	. "sub2sing-box/pkg/util"

	"github.com/spf13/cobra"
)

var subscriptions []string
var proxies []string
var template string
var output string
var delete string
var rename map[string]string

func init() {
	convertCmd.Flags().StringSliceVarP(&subscriptions, "subscription", "s", []string{}, "subscription urls")
	convertCmd.Flags().StringSliceVarP(&proxies, "proxy", "p", []string{}, "common proxies")
	convertCmd.Flags().StringVarP(&template, "template", "t", "", "template file path")
	convertCmd.Flags().StringVarP(&output, "output", "o", "", "output file path")
	convertCmd.Flags().StringVarP(&delete, "delete", "d", "", "delete proxy with regex")
	convertCmd.Flags().StringToStringVarP(&rename, "rename", "r", map[string]string{}, "rename proxy with regex")
	RootCmd.AddCommand(convertCmd)
}

var convertCmd = &cobra.Command{
	Use:   "convert",
	Long:  "Convert common proxy to sing-box proxy",
	Short: "Convert common proxy to sing-box proxy",
	Run: func(cmd *cobra.Command, args []string) {
		result := ""
		var err error
		result, err = Convert(subscriptions, proxies, template, delete, rename)
		if err != nil {
			fmt.Println(err)
			return
		}
		if output != "" {
			err = os.WriteFile(output, []byte(result), 0666)
			if err != nil {
				fmt.Println(err)
				return
			}
		} else {
			fmt.Println(string(result))
		}
	},
}
