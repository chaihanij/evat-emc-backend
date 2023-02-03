package users

// DeleteUser Delete User
// @Summary  Delete User
// @Description API For Delete User
// @ID detele-user
// @Accept json
// @Produce json
// @Tags USERS
// @Param Authorization header string true "for authentication"
// @Param uid path string true "uid of user"
// @Success 200 {object} dtos.DeleteUserResponseSwagger
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /v1/users/:uuid [delete]
func DeleteUser() {}
