package utils

import (
	"flag"
	"os"
)

// HandleFlag обработчик флагов командной строки
func HandleFlag() {
	flag.Func("a", "GRPc server address", func(aFlagValue string) error {
		return os.Setenv("SERVER_ADDRESS", aFlagValue)
	})

	flag.Func("d", "Address of db connection", func(dFlagValue string) error {
		return os.Setenv("DATABASE_DSN", dFlagValue)
	})
}
