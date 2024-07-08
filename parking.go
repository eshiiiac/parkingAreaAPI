package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type city struct {
	Location string  `json:"location"`
	Price    float64 `json:"price"`
	Type     string  `json:"type"`
}

var Bhaktapur = []city{
	{Location: "Kamalbinayak", Price: 0.00, Type: "Open Space"},
	{Location: "Siddhapokhari", Price: 10.00, Type: "Open Space"},
	{Location: "Nawa Durga", Price: 10.00, Type: "Closed space"},
	{Location: "Radhe Radhe", Price: 10.00, Type: "Underground"},
}

var Kathmandu = []city{
	{Location: "New Road", Price: 10.00, Type: "Open Space"},
	{Location: "Koteshwor", Price: 20.00, Type: "Open Space"},
	{Location: "New Baneshwor", Price: 15.00, Type: "Closed space"},
	{Location: "Gaushala", Price: 10.00, Type: "Underground"},
}

var Lalitpur = []city{
	{Location: "Jwalakhel", Price: 10.00, Type: "Open Space"},
	{Location: "Kumariparti", Price: 20.00, Type: "Open Space"},
	{Location: "Ekantakuna", Price: 15.00, Type: "Closed space"},
	{Location: "Dhobighat", Price: 10.00, Type: "Underground"},
}

func displayParkingArea(c *gin.Context) {
	var newParkings city

	if err := c.BindJSON(&newParkings); err != nil {
		return
	}

	Bhaktapur = append(Bhaktapur, newParkings)
	Kathmandu = append(Kathmandu, newParkings)
	Lalitpur = append(Lalitpur, newParkings)
	c.IndentedJSON(http.StatusCreated, newParkings)
}

func bktParkings(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Bhaktapur)
}

func ktmParkings(c *gin.Context) {
	c.IndentedJSON(http.StatusCreated, Kathmandu)
}
func ltpParkings(c *gin.Context) {
	c.IndentedJSON(http.StatusCreated, Lalitpur)
}

func home(c *gin.Context) {
	htmlContent := `<!DOCTYPE html>
<html>
<head>
	<title>Find Parking Areas</title>
</head>
<body>
	<h1>Find Parking areas near you!</h1>
	<form action="/navigate" method="POST">
		<label for="city">Select City:</label>
		<select name="city" id="city">
			<option value="bkt">Bhaktapur</option>
			<option value="ktm">Kathmandu</option>
			<option value="ltp">Lalitpur</option>
		</select>
		<button type="submit">Find Parking</button>
	</form>

	
</body>
</html>`
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(htmlContent))
	/*c.String(200,"Find Parking areas near you!\nadd '/bkt' to find parking areas in Bhaktapur city\n add '/ktm' to find parking areas in Kathmandu city")*/

}

func navigate(c *gin.Context) {
	city := c.PostForm("city")
	switch city {
	case "bkt":
		c.Redirect(http.StatusFound, "/bkt")
	case "ktm":
		c.Redirect(http.StatusFound, "/ktm")
	case "ltp":
		c.Redirect(http.StatusFound, "/ltp")
	default:
		c.String(http.StatusBadRequest, "Invalid city selection")
	}
}

func main() {

	router := gin.Default()
	router.GET("/", home)
	router.GET("/bkt", bktParkings)
	router.GET("/ktm", ktmParkings)
	router.GET("/ltp", ltpParkings)
	router.POST("/navigate", navigate)
	router.Run("localhost:8080")
}
