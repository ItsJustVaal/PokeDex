package main

func commandPokedex(c *Config) error {
	err := c.Client.GetPokedex()
	if err != nil {
		return err
	}
	return nil
}
