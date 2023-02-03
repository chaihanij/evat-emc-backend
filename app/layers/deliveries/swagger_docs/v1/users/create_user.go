package users

// CreateUser Create User
// @Summary Create User
// @Description API For Create User
// @ID post-user
// @Accept json
// @Produce json
// @Tags USERS
// @Param Authorization header string true "for authentication"
// @Param body body dtos.CreateUserRequestJSON true "All params related to users"
// @Success 200 {object} dtos.CreateUserResponseSwagger
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /v1/users [post]
func CreateUser() {}
