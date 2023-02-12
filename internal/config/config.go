package config

import "flag"

// адрес и порт запуска сервиса: переменная окружения ОС RUN_ADDRESS или флаг -a;
// адрес подключения к базе данных: переменная окружения ОС DATABASE_URI или флаг -d;
// адрес системы расчёта начислений: переменная окружения ОС ACCRUAL_SYSTEM_ADDRESS или флаг -r.

type Config struct {
	RunAddress           string `env:"RUN_ADDRESS"`
	DSN                  string `env:"DATABASE_URI"`
	AccrualSystemAddress string `env:"ACCRUAL_SYSTEM_ADDRESS"`
}

func New() *Config {

	cfg := &Config{}

	flag.StringVar(&cfg.RunAddress, "a", "localhost:8080", "-a <RUN_ADDRESS>")
	flag.StringVar(&cfg.DSN, "d", "", "-d <DATABASE_URI >")
	flag.StringVar(&cfg.AccrualSystemAddress, "r", "", "-r <ACCRUAL_SYSTEM_ADDRESS>")

	return cfg

}

func (c *Config) Check() error {
	return nil
}
