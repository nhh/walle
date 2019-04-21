package config

type VirtualHost struct {
	Name string
	Port int
	Ssl bool // Abstract this into a struct
	Location map[string] string // Same as above
}
