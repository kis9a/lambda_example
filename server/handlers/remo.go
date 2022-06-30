package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kis9a/lambda-sls/config"
	"github.com/kis9a/lambda-sls/models"
)

func SaveRemoA2DeviceEvents() gin.HandlerFunc {
	return func(c *gin.Context) {
		cfg := config.GetConfig()
		req, _ := http.NewRequest("GET", "https://api.nature.global/1/devices", nil)
		req.Header.Add("Authorization", "Bearer "+cfg.REMO_API_TOKEN)
		client := new(http.Client)
		r, _ := client.Do(req)
		defer r.Body.Close()
		byteArray, _ := ioutil.ReadAll(r.Body)
		var devices []models.RemoDevice
		var deviceA2 models.RemoDevice
		json.Unmarshal(byteArray, &devices)
		for _, d := range devices {
			if d.Name == "A2" {
				deviceA2 = d
			}
		}
		c.JSON(http.StatusOK, deviceA2)
	}
}
