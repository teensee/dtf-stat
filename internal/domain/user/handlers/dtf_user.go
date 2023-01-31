package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
	"io"
	"log"
	"net/http"
	"recomends/internal/domain/user/dto"
	"recomends/internal/domain/user/model"
	"time"
)

const DtfHandlerTag = "dtf_user_handler"

type DtfUserHandler struct {
	db *gorm.DB
}

func NewDtfUserHandler(db *gorm.DB) *DtfUserHandler {
	return &DtfUserHandler{
		db: db,
	}
}

func (h *DtfUserHandler) LoadDtfUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	h.loadAllUsers(1)
}

//total user counts = 702330
func (h *DtfUserHandler) loadAllUsers(from int) {
	for i := from; i <= 702345; i++ {
		requestUrl := fmt.Sprintf("https://api.dtf.ru/v1.9/user/%d", i)
		resp, err := http.Get(requestUrl)

		b, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
			continue
		}

		reqLog := model.NewUserRequestLog(i, requestUrl, resp.StatusCode, time.Now(), string(b))

		h.db.Save(reqLog)
		resp.Body.Close()

		if err != nil {
			log.Println(err)
			continue
		}

		var dtfUser dto.DtfUser
		json.Unmarshal(b, &dtfUser)

		entity := model.NewDtfUserFromDto(dtfUser)

		h.db.Save(entity)
		log.Printf("Saved %d user", entity.DtfUid)
		time.Sleep(335)

		_ = reqLog
		_ = entity
	}
}

func (h *DtfUserHandler) LoadDtfUsersInCli(context *cli.Context) error {
	from := context.Int("from")
	h.loadAllUsers(from)

	return nil
}
