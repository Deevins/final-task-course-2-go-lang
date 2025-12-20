package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Deevins/final-task-course-2-go-lang/gateway/internal/docs"
)

const swaggerIndexHTML = `<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <title>Gateway API - Swagger UI</title>
    <link
      rel="stylesheet"
      href="https://unpkg.com/swagger-ui-dist@5/swagger-ui.css"
    />
    <style>
      html, body { margin: 0; padding: 0; }
    </style>
  </head>
  <body>
    <div id="swagger-ui"></div>
    <script src="https://unpkg.com/swagger-ui-dist@5/swagger-ui-bundle.js"></script>
    <script>
      window.onload = () => {
        window.ui = SwaggerUIBundle({
          url: "/swagger/doc.json",
          dom_id: "#swagger-ui",
          presets: [SwaggerUIBundle.presets.apis],
          layout: "BaseLayout",
        });
      };
    </script>
  </body>
</html>`

func registerSwagger(engine *gin.Engine) {
	engine.GET("/swagger", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/")
	})
	engine.GET("/swagger/", serveSwaggerUI)
	engine.GET("/swagger/index.html", serveSwaggerUI)
	engine.GET("/swagger/doc.json", serveSwaggerJSON)
	engine.GET("/swagger/doc.yaml", serveSwaggerYAML)
}

func serveSwaggerUI(c *gin.Context) {
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.String(http.StatusOK, swaggerIndexHTML)
}

func serveSwaggerJSON(c *gin.Context) {
	c.Header("Content-Type", "application/json; charset=utf-8")
	c.Writer.WriteHeader(http.StatusOK)
	_, _ = c.Writer.Write(docs.SwaggerJSON())
}

func serveSwaggerYAML(c *gin.Context) {
	c.Header("Content-Type", "application/yaml; charset=utf-8")
	c.Writer.WriteHeader(http.StatusOK)
	_, _ = c.Writer.Write(docs.SwaggerYAML())
}
