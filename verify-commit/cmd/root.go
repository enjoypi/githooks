package cmd

import (
	"git.tap4fun.com/k2/githooks/verify-commit/root"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rootCmd = &cobra.Command{
		Use:     "verify-commit",
		Short:   "验证当前commit的文件是否允许提交至当前分支",
		Version: "1.0.1",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			err := viper.BindPFlags(cmd.Flags())
			if err != nil {
				return err
			}

			// Use config file from the flag.
			viper.SetConfigFile(viper.GetString("config"))
			return viper.ReadInConfig()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			return root.Run(cmd, viper.GetViper(), viper.GetString("hook-name"), viper.GetBool("verbose"))
		},
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringP("config", "c", "", "配置文件")
	rootCmd.PersistentFlags().StringP("hook-name", "n", "pre-commit", "被哪个git hook调用")
	rootCmd.PersistentFlags().Bool("verbose", false, "输出详细信息")
	rootCmd.PersistentFlags().StringP("working-tree", "w", ".", ".git所在工作目录")
	_ = rootCmd.MarkPersistentFlagRequired("config")
}
