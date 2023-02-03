package users

// UpdateUser Update User
// @Summary  Update User
// @Description API For Update User
// @ID put-user
// @Accept json
// @Produce json
// @Tags USERS
// @Param Authorization header string true "for authentication"
// @Param uid path string true "uid of user"
// @Param body body dtos.UpdateUserRequestJSON true "All params related to users"
// @Success 200 {object} dtos.UpdateUserResponseSwagger
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /v1/users/:uuid [put]
func UpdateUser() {}
