package config

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

func Parse() [] VirtualHost {

	// Todo do we support multiple files?
	// At first we should only parse a single one.

	// Todo feature: Config should be loadable via url.
	// So json would be a suitable solution
	// BENEATH i know that json is a horrible configuration file type

	// The docs looks a little confusing, but this part we actually need.

	var vhosts [] VirtualHost

	files, error  := ioutil.ReadDir("./data")

	for _, file := range files  {
		if !file.IsDir() && !strings.HasSuffix(file.Name(), "json") {
			continue
		}

		str, err := ioutil.ReadFile("./data/" + file.Name())

		if err != nil {
			panic(err)
		}

		vhost := VirtualHost{}

		error = json.Unmarshal([]byte(str), &vhost)

		if error != nil {
			panic(error)
		}

		vhosts = append(vhosts, vhost)

	}

	return vhosts

}