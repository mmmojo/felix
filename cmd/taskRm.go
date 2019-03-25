package cmd

import (
	"github.com/dejavuzhou/felix/models"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"strconv"
)

// taskrmCmd represents the taskrm command
var taskrmCmd = &cobra.Command{
	Use:   "taskRm",
	Short: "remove one task from task list",
	Long:  `usage: felix taskrm 1`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.ParseUint(args[0], 10, 64)
		if err != nil {
			color.Red("第一个参数必须为正整数:[%s]", err)
		}
		if err := models.TaskRm(uint(id)); err != nil {
			color.Red("%s", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(taskrmCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// taskrmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// taskrmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
