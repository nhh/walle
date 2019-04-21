package config

type ManagementServer struct {
    Servers []Server
    Location map[string] string // Same as above
}
