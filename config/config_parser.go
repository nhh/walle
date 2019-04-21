package config

import (
	"encoding/json"
    "github.com/nhh/walle/server"
    "io/ioutil"
	"strings"
)

func ParseServers() [] server.WalleServer {

	// Todo feature: Config should be loadable via url.
	// So json would be a suitable solution
	// BENEATH i know that json is a horrible configuration file type

	// The docs looks a little confusing, but this part we actually need.

	var servers [] server.WalleServer

	files, error  := ioutil.ReadDir("./data")

	for _, file := range files  {
		if !file.IsDir() && !strings.HasSuffix(file.Name(), "json") {
			continue
		}

		str, err := ioutil.ReadFile("./data/" + file.Name())

		if err != nil {
			panic(err)
		}

		walleServer := server.WalleServer{}

        error = json.Unmarshal([]byte(str), &walleServer)

        if error != nil {
            panic(error)
        }

        servers = append(servers, walleServer)

	}

	return servers

}
