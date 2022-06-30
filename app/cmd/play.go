package cmd

import (
	"gohub/pkg/config"
	"gohub/pkg/helpers"
	"gohub/pkg/logger"

	"github.com/spf13/cobra"
)

var CmdPlay = &cobra.Command{
	Use:   "play",
	Short: "Likes the Go Playground, but running at our application context",
	Run:   runPlay,
}

// 调试完成后请记得清除测试代码
func runPlay(cmd *cobra.Command, args []string) {
	defer helpers.Elapsed("normal call")()
	logger.Dump(config.Get[string]("app.name"))
}
