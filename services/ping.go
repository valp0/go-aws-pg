package services

// Ping is the service function that will call the repository
// function to issue a ping to the database.
func (s service) Ping() (string, error) {
	err := s.r.Ping()
	if err != nil {
		return "", err
	}

	return "DB pinged sucessfully!", nil
}
