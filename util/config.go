package util

//Config stores all configuration of the application
//The values are read by viper from a config file or environment variables

//type Config struct {
//	DbDriver    string `mapsstructure:"DB_DRIVER"`
//	DbAddress   string `mapstructure:"DB_ADDRESS"`
//	Db        	string `mapstructure:"DB"`
//	DbUser      string `mapstructure:"DB_USER"`
//	DbPass      string `mapstructure:"DB_PASS"`
//
//	DefaultPort string `mapstructure:"DEFAULT_PORT"`
//}
//
//func LoadConfig(path string)  (config Config, err error) {
//	viper.AddConfigPath(path)
//	viper.SetConfigName("app")
//	viper.SetConfigType("env")
//
//	viper.AutomaticEnv()
//
//	err = viper.ReadInConfig()
//	if err != nil {
//		return
//	}
//	err = viper.Unmarshal(&config)
//	return
//}
