package utils

func CheckErr(err error) error {
	if err != nil {
		panic(err)
	}

	return nil
}
