package certificate

import "github.com/gin-gonic/gin"

// type Handler struct {
// }

func NewEndpointHttpHandler(ginEngine *gin.Engine) {

	// Handler := &Handler{}

	v1 := ginEngine.Group("/v1")
	{
		v1.GET("/create-certificate/:fname/:lname/:date_created/:start_race/:stop_race", CreateCertificate)

	}

}
