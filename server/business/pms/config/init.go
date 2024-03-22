package config

/**
 * @Author: yuanyp8
 * @Author: geniusgavin@163.com
 * @Date: 2024/3/22 9:24
 * @Desc:
 */

func Init(pathFlag ...string) error {
	// 首先从consul提取配置
	if enableConsulFlag {
		if err := loadConfigFromConsul(); err != nil {
			return err
		}
	}

	// 从配置文件中提取配置
	return loadServerConfig(pathFlag...)
}
