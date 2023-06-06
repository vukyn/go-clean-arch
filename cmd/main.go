package main

import "log"

func main() {
	configPath := cconfig.GetConfigPath(os.Getenv("ENVIRONMENT"))
	config, err := config.GetConfig(configPath)
	if err != nil {
		panic(err)
	}
	log.Printf("starting gRPC server at port %v ...", 1234)
	s := application.NewServer(config)

	go func() {
		s.Start()
	}()
	defer s.Stop()

}
