package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"logitech-form-backend/database"
	"logitech-form-backend/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

func main() {
	// Cargar variables de entorno
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No se encontró archivo .env, usando variables de entorno del sistema")
	}

	// Conectar a la base de datos
	if err := database.Connect(); err != nil {
		log.Fatalf("❌ Error fatal al conectar a la base de datos: %v", err)
	}
	defer database.Close()

	// Crear aplicación Fiber
	app := fiber.New(fiber.Config{
		AppName:      "Logitech Support Form API",
		ServerHeader: "Logitech-API",
		ErrorHandler: customErrorHandler,
	})

	// Middlewares
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${method} ${path} (${latency})\n",
	}))

	// Configurar CORS
	allowedOrigins := os.Getenv("ALLOWED_ORIGINS")
	if allowedOrigins == "" {
		allowedOrigins = "*"
	}

	app.Use(cors.New(cors.Config{
		AllowOrigins:     allowedOrigins,
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization",
		AllowCredentials: true,
	}))

	// Rutas
	setupRoutes(app)

	// Puerto del servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Canal para manejar señales de cierre
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// Iniciar servidor en goroutine
	go func() {
		log.Printf("🚀 Servidor iniciado en http://localhost:%s", port)
		if err := app.Listen(":" + port); err != nil {
			log.Fatalf("❌ Error al iniciar el servidor: %v", err)
		}
	}()

	// Esperar señal de cierre
	<-quit
	log.Println("\n👋 Cerrando servidor gracefully...")
	
	if err := app.Shutdown(); err != nil {
		log.Printf("❌ Error al cerrar el servidor: %v", err)
	}
	
	log.Println("✅ Servidor cerrado exitosamente")
}

// setupRoutes configura todas las rutas de la API
func setupRoutes(app *fiber.App) {
	// Ruta de health check
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"mensaje": "API de Formularios Logitech - Funcionando correctamente",
			"version": "1.0.0",
			"status":  "ok",
		})
	})

	// Grupo de rutas API
	api := app.Group("/api/v1")

	// Rutas de formularios
	api.Post("/formularios", handlers.CrearFormulario)
	api.Get("/formularios", handlers.ObtenerFormularios)
	api.Get("/formularios/:id", handlers.ObtenerFormularioPorID)
}

// customErrorHandler maneja los errores de forma personalizada
func customErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	return c.Status(code).JSON(fiber.Map{
		"error":   true,
		"mensaje": err.Error(),
		"codigo":  code,
	})
}
