package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yuanyp8/wxcommerce/business/pms/config"
)

/**
 * @Author: yuanyp8
 * @Author: geniusgavin@163.com
 * @Date: 2024/3/21 23:47
 * @Desc:
 */

var (
	configPath string
)

var RootCmd = &cobra.Command{
	Use:   "pmsmodel",
	Short: "商品管理组件",
	Long:  "WXCommerce 平台商品管理微服务",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func Execute() {
	cobra.OnInitialize(initial)
	RootCmd.AddCommand(StartCmd)
	err := RootCmd.Execute()
	cobra.CheckErr(err)
}

func initial() {
	// 加载配置文件, 解析配置文件
	// 加载失败则报错
	// 加载成功则生成全局的Config.C()全局实例，后续可以全局调用
	cobra.CheckErr(config.Init(configPath))

}

func init() {
	RootCmd.PersistentFlags().StringVarP(&configPath, "config-file", "f", "etc/config.yaml", "the pmsmodel server config from file")
}
