package apiBooking

import (
	"context"
	"fmt"
	"strconv"
	"time"
	"user-gateway/internal/util"
	bookingProto "user-gateway/proto/booking"
	"user-gateway/proto/sdk"

	"github.com/gin-gonic/gin"
	"github.com/hadanhtuan/go-sdk/common"
)

type BookingController struct {
	ServiceBookingClient bookingProto.BookingServiceClient
}

func NewBookingController(serviceBookingClient bookingProto.BookingServiceClient) *BookingController {
	return &BookingController{ServiceBookingClient: serviceBookingClient}
}

func (bc *BookingController) GetPropertyDetail(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	var payload bookingProto.MsgGetPropertyRequest
	propertyId := c.Param("propertyId")
	if propertyId == "" {
		c.JSON(int(common.APIStatus.BadRequest), &sdk.BaseResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "Invalid request",
		})
		return
	}
	payload.PropertyId = propertyId
	result, _ := bc.ServiceBookingClient.GetPropertyDetail(ctx, &payload)
	newResult := util.ConvertResult(result)

	c.JSON(int(newResult.Status), newResult)
}
func (bc *BookingController) GetAllProperty(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	var payload bookingProto.MessageQueryRoom
	payload.Paginate = &sdk.Pagination{}
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		c.JSON(int(common.APIStatus.BadRequest), &sdk.BaseResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "Invalid request",
		})
		return
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		c.JSON(int(common.APIStatus.BadRequest), &sdk.BaseResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "Invalid request",
		})
		return
	}
	if page < 0 && page >= 100 {
		page = 1
	}

	if limit < 0 && limit >= 100 {
		limit = 10
	}
	payload.Paginate.Offset = int32((page - 1) * limit)
	payload.Paginate.Limit = int32(limit)
	result, _ := bc.ServiceBookingClient.GetAllProperty(ctx, &payload)
	newResult := util.ConvertResult(result)

	c.JSON(int(newResult.Status), newResult)
}

func (bc *BookingController) CreateProperty(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	var payload bookingProto.MsgCreatePropertyRequest
	err := c.BindJSON(&payload)

	if err != nil {
		c.JSON(int(common.APIStatus.BadRequest), &sdk.BaseResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "Invalid request",
		})
		return
	}
	result, _ := bc.ServiceBookingClient.CreateProperty(ctx, &payload)
	newResult := util.ConvertResult(result)
	c.JSON(int(newResult.Status), newResult)
}
func (bc *BookingController) UpdateProperty(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	var payload bookingProto.MsgUpdatePropertyRequest
	err := c.BindJSON(&payload)
	propertyId := c.Param("propertyId")
	if propertyId == "" || err != nil {
		c.JSON(int(common.APIStatus.BadRequest), &sdk.BaseResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "Invalid request",
		})
		return
	}
	payload.PropertyId = propertyId
	result, _ := bc.ServiceBookingClient.UpdateProperty(ctx, &payload)
	newResult := util.ConvertResult(result)

	c.JSON(int(newResult.Status), newResult)
}
func (bc *BookingController) DeleteProperty(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	var payload bookingProto.MsgDeletePropertyRequest
	propertyId := c.Param("propertyId")
	if propertyId == "" {
		c.JSON(int(common.APIStatus.BadRequest), &sdk.BaseResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "Invalid request",
		})
		return
	}
	payload.PropertyId = propertyId
	result, _ := bc.ServiceBookingClient.DeleteProperty(ctx, &payload)
	newResult := util.ConvertResult(result)

	c.JSON(int(newResult.Status), newResult)
}

func (bc *BookingController) GetBookingDetail(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	id := c.Param("bookingId")
	bookingID, err := strconv.ParseInt(id, 10, 64)
	fmt.Println(err, bookingID)
	if err != nil {
		res := &common.APIResponse{
			Message: "Booking ID is invalid",
			Status:  common.APIStatus.BadRequest,
		}
		c.JSON(int(res.Status), res)
		return
	}
	result, _ := bc.ServiceBookingClient.GetBookingDetail(ctx, &bookingProto.MsgGetBookingRequest{
		BookingId: bookingID,
	})
	fmt.Println(result)
	newResult := util.ConvertResult(result)
	c.JSON(int(newResult.Status), newResult)
}

func (bc *BookingController) CreateReview(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	var payload bookingProto.MsgCreateReviewRequest
	err := c.BindJSON(&payload)

	if err != nil {
		c.JSON(int(common.APIStatus.BadRequest), &sdk.BaseResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "Invalid request",
		})
		return
	}
	result, _ := bc.ServiceBookingClient.CreateReview(ctx, &payload)
	newResult := util.ConvertResult(result)
	c.JSON(int(newResult.Status), newResult)
}

func (bc *BookingController) UpdateReview(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	var payload bookingProto.MsgUpdateReviewRequest
	err := c.BindJSON(&payload)
	reviewId := c.Param("reviewId")
	if reviewId == "" || err != nil {
		c.JSON(int(common.APIStatus.BadRequest), &sdk.BaseResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "Invalid request",
		})
		return
	}
	payload.ReviewId = reviewId
	result, _ := bc.ServiceBookingClient.UpdateReview(ctx, &payload)
	newResult := util.ConvertResult(result)

	c.JSON(int(newResult.Status), newResult)
}

func (bc *BookingController) DeleteReview(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	var payload bookingProto.MsgDeleteReviewRequest
	reviewId := c.Param("reviewId")
	if reviewId == "" {
		c.JSON(int(common.APIStatus.BadRequest), &sdk.BaseResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "Invalid request",
		})
		return
	}
	payload.ReviewId = reviewId
	result, _ := bc.ServiceBookingClient.DeleteReview(ctx, &payload)
	newResult := util.ConvertResult(result)

	c.JSON(int(newResult.Status), newResult)
}

func (bc *BookingController) GetReview(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	var payload bookingProto.MessageQueryReview
	payload.Paginate = &sdk.Pagination{}
	err := c.BindJSON(&payload)
	if err != nil {
		c.JSON(int(common.APIStatus.BadRequest), &sdk.BaseResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "Invalid request",
		})
		return
	}

	result, err := bc.ServiceBookingClient.GetReview(ctx, &payload)
	newResult := util.ConvertResult(result)
	c.JSON(int(newResult.Status), newResult)
}
