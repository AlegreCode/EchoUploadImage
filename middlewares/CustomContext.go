package middlewares

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"

	. "github.com/alegrecode/echo/EchoUploadImage/models"
	uuid "github.com/satori/go.uuid"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
)

type CustomContext struct {
	echo.Context
}

func (c *CustomContext) SetFlash(t string, msg interface{}) bool {
	session, _ := session.Get("flash", c)
	session.Options = &sessions.Options{
		MaxAge: 1,
	}
	flashes := map[string]interface{}{"type": t, "msg": msg}
	flashesJSON, _ := json.Marshal(flashes)
	flashesString := string(flashesJSON)
	session.AddFlash(flashesString)
	session.Save(c.Request(), c.Response())
	return true
}

func (c *CustomContext) GetFlash() map[string]interface{} {
	session, _ := session.Get("flash", c)
	if flash := session.Flashes(); len(flash) > 0 {
		flashes := make(map[string]interface{})
		json.Unmarshal([]byte(flash[0].(string)), &flashes)
		return flashes
	}
	return map[string]interface{}{}
}

func (c *CustomContext) Upload(fieldName string, dirs ...string) (bool, error) {
	// Source
	file, err := c.FormFile(fieldName)
	if err != nil {
		return false, err
	}
	src, err := file.Open()
	if err != nil {
		return false, err
	}
	defer src.Close()

	name := "img"
	if file_name := c.FormValue("name"); file_name != "" {
		name = file_name
	}

	// loop of dirs to create the path of destination
	path_dest := "assets"
	for _, folder := range dirs {
		path_dest = filepath.Join(path_dest, folder)
	}

	// Destination
	_, err1 := os.Stat(path_dest)

	if os.IsNotExist(err1) {
		errDir := os.MkdirAll(path_dest, 0755)
		if errDir != nil {
			log.Fatal(err1)
		}

	}
	time_unix := time.Now().Unix()
	time_stamp := strconv.Itoa(int(time_unix))
	ext := filepath.Ext(file.Filename)
	name_file := fmt.Sprintf("%s%s%s", name, time_stamp, ext)
	path_file := filepath.Join(path_dest, name_file)
	dst, err := os.Create(path_file)
	if err != nil {
		return false, err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return false, err
	}

	// =========================================================
	var img Image
	img.Id = uuid.Must(uuid.NewV4(), err)
	img.Name = c.FormValue("name")
	img.Url = name_file
	img.Save()

	return true, nil
}
