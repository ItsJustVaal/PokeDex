package main

func commandData(c *Config) error {
	c.Client.GetPokeData(c.Pokemon)
	*c.Pokemon = ""
	return nil
}
