package app

//Point of connection with gin gonic
import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

func StartApplication() {
	mapsUrl()
	router.Run(":8080")
}
