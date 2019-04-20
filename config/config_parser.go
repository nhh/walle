package config

import (
	"encoding/json"
	"io/ioutil"
)

func Parse() [] VirtualHost {

	// Todo do we support multiple files?
	// At first we should only parse a single one.

	// Todo feature: Config should be loadable via url.
	// So json would be a suitable solution
	// BENEATH i know that json is a horrible configuration file type

	// The docs looks a little confusing, but this part we actually need.

	str, err := ioutil.ReadFile("./data/config.json")

	if err != nil {
		panic(err)
	}

	var vhosts [] VirtualHost

	error := json.Unmarshal([]byte(str), &vhosts)

	if error != nil {
		panic(error)
	}

	return vhosts

}