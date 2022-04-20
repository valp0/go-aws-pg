package services

func (s service) Ping() (string, error) {
	err := s.r.Ping()
	if err != nil {
		return "", err
	}

	return "DB pinged sucessfully!", nil
}
