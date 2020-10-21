package rook

type Config struct {


}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Players() int { 
	return 8
}