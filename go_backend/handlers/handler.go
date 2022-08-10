package handlers

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

const token = "F-QFQpmCL9UkR3qyoXnLkzWj03s6m4eCvYgDl1ePfHBf9ph7yxaSgQ6WN0i9giNgRTfONwVMK1f977r_g71oNQ=="
const bucket = "users_business_events"
const org = "iot"

type recipe_json struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	ibtsTemp    float32 `json:"temp"`
	ibtsROR     float32 `json:"ror"`
	Description string  `json:"description"`
}

type RecipesHandler struct {
	client *influxdb2.Client
	ctx    context.Context
}

func NewRecipesHandler(ctx context.Context, client *influxdb2.Client) *RecipesHandler {
	return &RecipesHandler{
		client: client,
		ctx:    ctx,
	}
}

func (handler *RecipesHandler) Influxdb_connect() influxdb2.Client {
	client := influxdb2.NewClientWithOptions("http://localhost:8086", token,
		influxdb2.DefaultOptions().
			SetUseGZip(true).
			SetTLSConfig(&tls.Config{
				InsecureSkipVerify: true,
			}))
	return client
}

func (handler *RecipesHandler) NewRecipeHandler(c *gin.Context) {
	// // get non-blocking write client
	writeAPI := handler.Influxdb_connect().WriteAPI(org, bucket)

	// write line protocol
	writeAPI.WriteRecord(fmt.Sprintf("IBTS_TEMP,deviceId=A0100,unit=temperature max=%.1f", 100.2))
	writeAPI.WriteRecord(fmt.Sprintf("ROR,deviceId=A0100,unit=temperature max=%.1f", 5.1))

	// Flush writes
	writeAPI.Flush()

	// always close client at the end
	defer handler.Influxdb_connect().Close()

	c.JSON(http.StatusOK, gin.H{"message": "Recipe New OK"})
}

func (handler *RecipesHandler) ListRecipeHandler(c *gin.Context) {
	queryAPI := handler.Influxdb_connect().QueryAPI(org)
	query := fmt.Sprintf(`from(bucket: "%v") |> range(start: -1d)`, bucket)
	result, err := queryAPI.Query(context.Background(), query)
	if err != nil {
		panic(err)
	}

	for result.Next() {
		record := result.Record()
		fmt.Printf("%v %v: %v=%v\n", record.Time(), record.Measurement(), record.Field(), record.Value())
		c.JSON(http.StatusOK, record.Time())
		c.JSON(http.StatusOK, record.Measurement())
		c.JSON(http.StatusOK, record.Field())
		c.JSON(http.StatusOK, record.Value())
	}

	// always close client at the end
	defer handler.Influxdb_connect().Close()

	c.JSON(http.StatusOK, gin.H{"message": "Recipe List OK"})
}

func (handler *RecipesHandler) UpdateRecipeHandler(c *gin.Context) {

}

func (handler *RecipesHandler) DeleteRecipeHandler(c *gin.Context) {

}

func (handler *RecipesHandler) SearchRecipeHandler(c *gin.Context) {

}
