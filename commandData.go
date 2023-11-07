package main

func commandData(c *Config) error {
	err := c.Client.GetPokeData(c.Pokemon)
	if err != nil {
		return err
	}
	*c.Pokemon = ""
	return nil
}
