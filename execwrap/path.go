package execwrap

func (c *CmdExecutor) Path(executable string) (string, error) {
	path, err := c.Commander.LookPath(executable)
	if err != nil {
		return "", err
	}
	return path, nil
}
