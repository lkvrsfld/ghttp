package main

var ghttp = GHTTP{
	Host:      GetEnvOrDefault("GOHTTP_HOST", "0.0.0.0"),
	Port:      GetEnvOrDefault("GOHTTP_PORT", "8080"),
	staticDir: GetEnvOrDefault("GOHTTP_STATIC_DIR", ""),
}

func main() {
	if err := ghttp.Init(); err != nil {
		panic(err)
	}
	if err := ghttp.Start(); err != nil {
		panic(err)
	}
}
