package domain

type User struct {  // Authentication would be done on keycloak server
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Phone     string `db:"phone"`
	Email     string `db:"email"`
	UserId	  int	 `db:"user_id"`
}

