package cmd

import (
	"github.com/go-jarvis/cobrautils"
	"github.com/spf13/cobra"
	"github.com/tangx/tangx.github.io/cdn-flusher/pkg/qcdn"
)

var rootCmd = &cobra.Command{
	Use:   "flusher",
	Short: "刷新并预热",
	Run: func(cmd *cobra.Command, args []string) {
		qcdn.Do(flag)
	},
}

var flag = &qcdn.Flag{
	SiteMap: "./docs/sitemap.xml",
}

func init() {
	cobrautils.BindFlags(rootCmd, flag)
}

func Execute() error {
	return rootCmd.Execute()
}
