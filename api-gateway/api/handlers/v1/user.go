package v1

import (
	"bytes"
	"context"
	"fmt"
	pb "github.com/baxromumarov/work/api-gateway/genproto"
	l "github.com/baxromumarov/work/api-gateway/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	"github.com/spf13/cast"
	"google.golang.org/protobuf/encoding/protojson"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

// Insert data and meta to db
func (h *handlerV1) CreateData(c *gin.Context) {
	var (
		body        pb.Request
		jspbMarshal protojson.MarshalOptions
		a           int64
	)

	for a = 1; a < 21; a++ {
		jspbMarshal.UseProtoNames = true
		url := fmt.Sprintf("https://gorest.co.in/public/v1/posts?page=%d", a)
		response, err := http.Get(url)

		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}

		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
		defer cancel()

		err = jsonpb.Unmarshal(bytes.NewReader(responseData), &body)
		if err != nil {
			log.Fatal(err)
		}
		_, err = h.serviceManager.UserService().Create(ctx, &body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			h.log.Error("failed to create data", l.Error(err))
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})

	}

}

// GetData gets data by id
func (h *handlerV1) GetDataById(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.PostService().GetDataById(
		ctx, &pb.ByIdReq{
			Id: guid,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get user", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// // ListData returns list of datas
func (h *handlerV1) GetDataList(c *gin.Context) {
	//var jspbMarshal protojson.MarshalOptions
	//jspbMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.PostService().GetAllData(ctx, &pb.Empty{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to list data", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// UpdateData updates data by id
// route /v1/data/{id} [put]
func (h *handlerV1) UpdateData(c *gin.Context) {
	var (
		body        pb.Data
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}
	guid := c.Param("id")
	body.Id = cast.ToInt64(guid)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.PostService().UpdateData(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to update user", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// DeleteData deletes data by id
// route /v1/data/{id} [delete]
func (h *handlerV1) DeleteData(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.PostService().DeleteById(
		ctx, &pb.ByIdReq{
			Id: guid,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete user", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}
