package faucet

func (f *Faucet) loadKey() error {
	if !f.keyExists(f.keyName) && f.keyMnemonic != "" {
		backend := "file"
		if _, err := f.cliexec("", []string{"keys", "add", f.keyName, "--keyring-backend", backend, "--recover"}, f.keyMnemonic, f.keyringPassword, f.keyringPassword); err != nil {
			return err
		}
	}

	var err error
	backend := "file"
	f.faucetAddress, err = f.cliexec("", []string{"keys", "show", f.keyName, "--keyring-backend", backend, "-a"}, f.keyringPassword)
	if err != nil {
		return err
	}

	return nil
}

func (f *Faucet) keyExists(keyname string) bool {
	backend := "file"
	if _, err := f.cliexec("", []string{"keys", "show", keyname, "--keyring-backend", backend}, f.keyringPassword); err != nil {
		return false
	}

	return true
}
