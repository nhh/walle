package config

type VirtualHost struct {
	Name string
	Ssl map[string] string // Abstract this into a struct
	Proxy map[string] string // Same as above
}