package files

// FindOneFile Find One File
// @Summary Find One File
// @Description API For Find One File
// @ID get-one-file
// @Accept json
// @Produce json
// @Tags FILES
// @Param file_uuid path string true "file_uuid of files"
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /v1/files/:file_uuid [get]
func FindOneFile() {}
