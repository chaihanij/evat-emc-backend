package members

// UpdateMemberImage Update Member Image
// @Summary Update Member Image
// @Description API For Update Member Image
// @ID post-member-image
// @Accept mpfd
// @Produce json
// @Tags MEMBERS
// @Param Authorization header		string	true "for authentication"
// @Param member_uuid 	path		string	true "member_uuid of member"
// @Param image 		formData	file 	true "file image upload"
// @Success 200 {object} dtos.UpdateMemberImageResponseSwagger
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /v1/members/:member_uuid/image [post]
func UpdateMemberImage() {}
