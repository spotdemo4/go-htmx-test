package home

import (
	"go-htmx-test/db"
	"go-htmx-test/models"
	"go-htmx-test/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type HomeHandler struct{}

func (h HomeHandler) Any(c echo.Context) error {
	if c.Request().Method == http.MethodGet {
		return h.Get(c)
	} else if c.Request().Method == http.MethodPost {
		return h.Post(c)
	} else if c.Request().Method == http.MethodDelete {
		return h.Delete(c)
	} else {
		return c.NoContent(http.StatusMethodNotAllowed)
	}
}

func (h HomeHandler) Get(c echo.Context) error {
	var items []models.Item

	if r := db.DB.Find(&items); r.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, r.Error.Error())
	}

	return utils.Render(c, http.StatusOK, Hello(items))
}

func (h HomeHandler) Post(c echo.Context) error {
	var items []models.Item

	if r := db.DB.Find(&items); r.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, r.Error.Error())
	}

	return utils.Render(c, http.StatusOK, Hello(items))
}

func (h HomeHandler) Delete(c echo.Context) error {
	itemID := c.FormValue("itemID")
	confirmItemID := c.FormValue("confirmItemID")

	if itemID != "" {
		var item models.Item

		if r := db.DB.First(&item, itemID); r.Error != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, r.Error.Error())
		}

		println("Selected item: ", item.Name)
		return utils.Render(c, http.StatusOK, DeleteConfirmation(item))
	} else if confirmItemID != "" {
		if r := db.DB.Delete(&models.Item{}, confirmItemID); r.Error != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, r.Error.Error())
		}

		var items []models.Item

		if r := db.DB.Find(&items); r.Error != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, r.Error.Error())
		}

		println("Deleted item: ", confirmItemID)
		return utils.Render(c, http.StatusOK, ItemList(items))
	} else {
		return c.NoContent(http.StatusBadRequest)
	}
}
