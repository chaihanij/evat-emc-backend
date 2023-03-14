package healthcheck

// HealtCheck HealtCheck.
// @Summary HealtCheck
// @Description HealtCheck
// @ID healt-check
// @Accept json
// @Produce json
// @Tags HEALTHCHECK
// @Success 200 {object} utils.BaseSuccessResponse
// @Failure 401 {object} utils.ErrorResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /health [get]
func HealtCheck() {}
