package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"io"
	"net/http"
	_ "url-shortener/docs"
	"url-shortener/internal/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	engine := gin.New()

	engine.GET("/", h.welcome)
	engine.GET("/a", h.registerNewSite)
	engine.GET("/s/:urlHash", h.redirectToSite)
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return engine
}

// RegisterNewSite godoc
// @Summary      register
// @Description  method create shortcut to site
// @Accept       json
// @Produce      json
// @Param        url   	   query     string  true  "Site url"
// @Success      200       {object}  string
// @Router       /a        [get]
func (h *Handler) registerNewSite(context *gin.Context) {
	urlToRegister, hasParameter := context.GetQuery("url")
	if hasParameter {
		urlHash, err := h.service.SaveUrl(urlToRegister)
		if err != nil {
			logrus.Warningf("Failed to save url: %s", err)
			context.JSON(http.StatusInternalServerError, "failed to save url")
		}
		context.JSON(http.StatusOK, urlHash)
	} else {
		logrus.Warning("Url on query not found")
		context.JSON(http.StatusBadRequest, "Url on query not found")
	}
}

// RedirectToSite godoc
// @Summary      redirect
// @Description  method redirect to site by shortcut
// @Accept       json
// @Produce      json
// @Param        urlHash   path      string  true  "Site hash"
// @Success      302       {object}  string
// @Router       /s/{urlHash} [get]
func (h *Handler) redirectToSite(context *gin.Context) {
	urlHash := context.Param("urlHash")
	url, err := h.service.GetUrl(urlHash)
	if err != nil {
		logrus.Warningf("Failed to find url: %s", err)
		context.JSON(http.StatusInternalServerError, "failed to find url")
	} else {
		context.Redirect(http.StatusFound, url)
	}
}

func (h *Handler) welcome(context *gin.Context) {
	_, _ = io.WriteString(context.Writer, "Welcome to url-shortener. "+
		"Use /a?url=<your site to create shortcut> and /s/<shortcut> to get to your website")
}
