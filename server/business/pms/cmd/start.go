package cmd

import (
	"github.com/spf13/cobra"
)

/**
 * @Author: yuanyp8
 * @Author: geniusgavin@163.com
 * @Date: 2024/3/21 23:57
 * @Desc:
 */

var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "pms服务启动",
	Long:  "WXCommerce 平台商品管理微服务启动",
	RunE: func(cmd *cobra.Command, args []string) error {
		//// 接收信号处理
		//ch := make(chan os.Signal, 1)
		//signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP, syscall.SIGQUIT)
		//
		//// 生成httpserver实例
		//svr := newService(config.C().Server)
		//
		//// 等待信号处理
		//go svr.waitSign(ch)
		//
		//// 启动服务
		//if err := svr.start(); err != nil {
		//	if !strings.Contains(err.Error(), "http: Server closed") {
		//		return err
		//	}
		//}
		//// 启动服务之后就应该阻塞监听了
		//// zap.L().Named("cmd").Debug("完成http server模块加载", zap.String("http_addr", config.C().Server.HttpAddr()))

		return nil
	},
}

//type service struct {
//	http *protocol.HTTPService
//	// grpc *protocol.GRPCService
//	l *zap.Logger
//}
//
//func newService(server *server.Server) *service {
//	return &service{
//		http: protocol.NewHTTPService(server),
//		l:    zap.L().Named("cmd.service"),
//	}
//}
//
//func (s *service) start() error {
//
//	s.l.Info("loaded http server")
//
//	s.l.Info("loaded http app", zap.String("app", application.Restful().Load()))
//	s.l.Info("loaded pkg app", zap.String("app", application.Internal().Load()))
//	return s.http.Start()
//}
//
//func (s *service) waitSign(sign chan os.Signal) {
//	for sg := range sign {
//		switch v := sg.(type) {
//		default:
//			s.l.Info("receive signal, start graceful shutdown", zap.String("signal", v.String()))
//
//			if err := s.http.Stop(); err != nil {
//				s.l.Error("http graceful shutdown err, force exit", zap.Error(err))
//			} else {
//				s.l.Info("http service stop complete")
//			}
//			return
//		}
//	}
//}
