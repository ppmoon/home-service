package boot

import "github.com/ppmoon/home-service/infrastructure/config"

const Path = "./config"

func Boot() {
	config.InitConfig(Path)
}
