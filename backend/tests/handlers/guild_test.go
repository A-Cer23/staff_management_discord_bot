package testHandlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/A-Cer23/adminbot-backend/db"
	"github.com/A-Cer23/adminbot-backend/handlers"
	"github.com/A-Cer23/adminbot-backend/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const (
	dbURL       = "postgres://user:password@:5432/adminstrationbotdb"
	URL         = "http://localhost:9090"
	GUILD_ROUTE = "/guild"
)

func TestCreateGuild_Success(t *testing.T) {

	router := gin.Default()
	router.POST(GUILD_ROUTE, func(c *gin.Context) {
		dbPool, _ := db.ConnectDB(dbURL)
		handlers.CreateGuild(c, dbPool)
	})

	randomGuildID := strconv.FormatInt(rand.Int63(), 10)
	randomOwnerID := strconv.FormatInt(rand.Int63(), 10)
	randomInguild := rand.Intn(2) == 1

	currentTime := time.Now().UTC().Format(time.RFC3339)

	requestBody := models.Guild{
		GuildID:   randomGuildID,
		OwnerID:   randomOwnerID,
		GuildName: "guildTest",
		JoinedAt:  currentTime,
		InGuild:   randomInguild,
	}

	jsonValue, _ := json.Marshal(requestBody)

	req, err := http.NewRequest("POST", URL+GUILD_ROUTE, bytes.NewBuffer(jsonValue))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	fmt.Println(*w)

	assert.Equal(t, http.StatusCreated, w.Code)

}
