package config

type Server struct {
    VirtualHosts [] VirtualHost //lul, that was easy
    Port int
    Ssl bool // Abstract this into a struct
    Location map[string] string // Same as above
}

