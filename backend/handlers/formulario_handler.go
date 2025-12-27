package handlers

import (
	"logitech-form-backend/database"
	"logitech-form-backend/models"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// CrearFormulario maneja la creación de un nuevo formulario de soporte
func CrearFormulario(c *fiber.Ctx) error {
	// Parsear el body del request
	formulario := new(models.FormularioSoporte)
	
	if err := c.BodyParser(formulario); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"mensaje": "Error al procesar los datos del formulario",
			"detalle": err.Error(),
		})
	}

	// Capturar metadata del request
	formulario.IPOrigen = c.IP()
	formulario.UserAgent = c.Get("User-Agent")
	
	// Validaciones básicas
	if strings.TrimSpace(formulario.NombreCompleto) == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"mensaje": "El nombre completo es requerido",
		})
	}

	if strings.TrimSpace(formulario.Email) == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"mensaje": "El email es requerido",
		})
	}

	if strings.TrimSpace(formulario.TipoProducto) == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"mensaje": "El tipo de producto es requerido",
		})
	}

	if strings.TrimSpace(formulario.CategoriaProblema) == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"mensaje": "La categoría del problema es requerida",
		})
	}

	if strings.TrimSpace(formulario.DescripcionProblema) == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"mensaje": "La descripción del problema es requerida",
		})
	}

	// Guardar en la base de datos
	if err := database.DB.Create(&formulario).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"mensaje": "Error al guardar el formulario en la base de datos",
			"detalle": err.Error(),
		})
	}

	// Respuesta exitosa
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error":   false,
		"mensaje": "Formulario enviado exitosamente. Nuestro equipo de soporte se pondrá en contacto contigo pronto.",
		"id":      formulario.ID,
	})
}

// ObtenerFormularios obtiene todos los formularios (para uso interno)
func ObtenerFormularios(c *fiber.Ctx) error {
	var formularios []models.FormularioSoporte
	
	// Obtener todos los formularios ordenados por fecha de creación descendente
	if err := database.DB.Order("fecha_creacion DESC").Find(&formularios).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"mensaje": "Error al obtener los formularios",
			"detalle": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error":       false,
		"total":       len(formularios),
		"formularios": formularios,
	})
}

// ObtenerFormularioPorID obtiene un formulario específico por ID
func ObtenerFormularioPorID(c *fiber.Ctx) error {
	id := c.Params("id")
	var formulario models.FormularioSoporte

	if err := database.DB.First(&formulario, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"mensaje": "Formulario no encontrado",
		})
	}

	return c.JSON(fiber.Map{
		"error":      false,
		"formulario": formulario,
	})
}
