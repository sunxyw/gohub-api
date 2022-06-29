package cmd

import (
	"gohub/pkg/jwt"
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
	jwt.InitWithProvider(jwt.NewFirebaseProvider())
	token := jwt.IssueToken("test")
	uid, err := jwt.ParseToken(token)
	if err != nil {
		logger.LogIf(err)
	}
	logger.Dump(uid)
}
