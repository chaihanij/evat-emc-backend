package v1

// Login
// @Summary Login By Email
// @Description Login By Email
// @ID post-login
// @Accept json
// @Produce json
// @Tags USERS
// @Param body body dtos.LoginRequestJSON true "All params related to users"
// @Success 200 {object} dtos.LoginResponseSwagger
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /v1/users.login [post]
func Login() {}
