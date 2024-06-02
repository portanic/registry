package handlers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/portanic/registry/internal/models"
)

func RegisterPlugin(c echo.Context) error {
	var plugin models.Plugin
	if err := c.Bind(&plugin); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	plugin.ID = uuid.New().String()
	models.Registry.Plugins[plugin.ID] = plugin

	return c.JSON(http.StatusCreated, plugin)
}

func ListPlugins(c echo.Context) error {
	plugins := make([]models.Plugin, 0, len(models.Registry.Plugins))
	for _, plugin := range models.Registry.Plugins {
		plugins = append(plugins, plugin)
	}
	return c.JSON(http.StatusOK, plugins)
}

func GetPlugin(c echo.Context) error {
	id := c.Param("id")
	plugin, exists := models.Registry.Plugins[id]
	if !exists {
		return echo.NewHTTPError(http.StatusNotFound, "Plugin not found")
	}
	return c.JSON(http.StatusOK, plugin)
}
