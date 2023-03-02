package api

import (
	"encoding/json"
	"github.com/inspectorvitya/collector-data/db"
	"github.com/valyala/fasthttp"
	"go.mongodb.org/mongo-driver/mongo"
	"strconv"
	"strings"
)

func createUserInfo(ctx *fasthttp.RequestCtx) {
	info := db.InfoDevice{}
	body := ctx.PostBody()
	err := json.Unmarshal(body, &info)
	if err != nil {
		ctx.WriteString(err.Error())
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	err = db.CreateUserDeviceInfo(ctx, info)
	if err != nil {
		ctx.WriteString(err.Error())
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	ctx.SetStatusCode(fasthttp.StatusCreated)
}

func getInfoUserById(ctx *fasthttp.RequestCtx) {
	id := ctx.UserValue("userId")
	idUInt, err := strconv.ParseUint(id.(string), 10, 32)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	data, err := db.GetUserInfoDeviceById(ctx, uint32(idUInt))
	if err != nil {
		if err == mongo.ErrNoDocuments {
			ctx.SetStatusCode(fasthttp.StatusNotFound)
			return
		}
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	body, err := json.Marshal(data)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	ctx.SetBody(body)
	ctx.SetStatusCode(fasthttp.StatusOK)
}

func GetTop100(ctx *fasthttp.RequestCtx) {
	field := strings.Replace(ctx.UserValue("field").(string), "-", "_", -1)

	name := string(ctx.QueryArgs().Peek("name"))
	by := string(ctx.QueryArgs().Peek("by"))
	if field == "" {
		_, err := ctx.WriteString("field empty")
		if err != nil {
			ctx.SetStatusCode(fasthttp.StatusInternalServerError)
			return
		}
		ctx.SetStatusCode(fasthttp.StatusOK)
		return
	}
	top, err := db.GetTop100(ctx, field, by, name)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	if len(top) == 0 {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}
	body, err := json.Marshal(top)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	ctx.SetBody(body)
	ctx.SetStatusCode(fasthttp.StatusOK)
}
