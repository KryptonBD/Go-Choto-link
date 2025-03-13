package routes

import (
	"choto-link/db"
	"choto-link/models"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

var rateLimitTime = 30 * 60 * time.Second

func shortenURL(ctx *gin.Context) {
	var requestBody models.Request
	err := ctx.ShouldBind(&requestBody)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not parse JSON",
		})
		return
	}

	rdb := db.CreateClient(1)
	defer rdb.Close()

	// check if client haven't exceed their rate limit
	val, err := rdb.Get(db.Ctx, ctx.ClientIP()).Result()

	if err == redis.Nil {
		_ = rdb.Set(db.Ctx, ctx.ClientIP(), os.Getenv("RATE_LIMIT"), rateLimitTime).Err()
	} else {
		valInt, _ := strconv.Atoi(val)
		if valInt <= 0 {
			limit, _ := rdb.TTL(db.Ctx, ctx.ClientIP()).Result()
			ctx.JSON(http.StatusServiceUnavailable, gin.H{
				"error":            "Rate limit exceeded",
				"rate_limit_reset": limit / time.Nanosecond / time.Minute,
			})

			return
		}
	}

	if !govalidator.IsURL(requestBody.URL) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL"})
		return
	}

	// Generate a random short URL
	var shortUrl string
	if requestBody.CustomShort == "" {
		shortUrl = uuid.New().String()[:6]
	} else {
		shortUrl = requestBody.CustomShort
		if db.CheckIfShortURLExists(shortUrl) {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Custom short URL already taken"})
			return
		}
	}

	if requestBody.Expiry == 0 {
		requestBody.Expiry = 24 * time.Hour
	}

	// Set the key-value pair in Redis
	err = rdb.Set(db.Ctx, shortUrl, requestBody.URL, requestBody.Expiry).Err()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	response := models.Response{
		Request:        &requestBody,
		RateLimit:      30,
		ResetRateLimit: 10,
	}

	// Decrement the rate limit
	rateLimit, _ := rdb.Decr(db.Ctx, ctx.ClientIP()).Result()
	response.RateLimit = int(rateLimit)

	// Get the rate limit reset time
	ttl, _ := rdb.TTL(db.Ctx, ctx.ClientIP()).Result()
	response.ResetRateLimit = ttl / time.Nanosecond / time.Minute

	response.CustomShort = os.Getenv("BASE_URL") + "/" + shortUrl

	ctx.JSON(http.StatusOK, response)
}
