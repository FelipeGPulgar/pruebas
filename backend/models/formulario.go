package models

import (
	"time"
)

// FormularioSoporte representa un formulario de soporte técnico de Logitech
type FormularioSoporte struct {
	ID uint `gorm:"primaryKey;autoIncrement" json:"id"`

	// Información personal
	NombreCompleto string `gorm:"type:varchar(255);not null" json:"nombre_completo" validate:"required"`
	Email          string `gorm:"type:varchar(255);not null;index" json:"email" validate:"required,email"`
	Telefono       string `gorm:"type:varchar(50)" json:"telefono"`

	// Información del producto
	TipoProducto   string `gorm:"type:varchar(100);not null;index" json:"tipo_producto" validate:"required"`
	ModeloProducto string `gorm:"type:varchar(100)" json:"modelo_producto"`
	NumeroSerie    string `gorm:"type:varchar(100)" json:"numero_serie"`

	// Detalles del problema
	CategoriaProblema   string `gorm:"type:varchar(100);not null;index" json:"categoria_problema" validate:"required"`
	DescripcionProblema string `gorm:"type:text;not null" json:"descripcion_problema" validate:"required"`

	// Información adicional
	SistemaOperativo string     `gorm:"type:varchar(100)" json:"sistema_operativo"`
	FechaCompra      *time.Time `gorm:"type:date" json:"fecha_compra"`
	LugarCompra      string     `gorm:"type:varchar(255)" json:"lugar_compra"`

	// Preferencias de contacto
	MetodoContactoPreferido string `gorm:"type:enum('email','telefono','ambos');default:'email'" json:"metodo_contacto_preferido"`
	HorarioContacto         string `gorm:"type:varchar(100)" json:"horario_contacto"`

	// Archivos adjuntos
	ArchivoAdjunto string `gorm:"type:varchar(500)" json:"archivo_adjunto"`

	// Aceptación de términos
	AceptaTerminos   bool `gorm:"default:false" json:"acepta_terminos"`
	AceptaMarketing  bool `gorm:"default:false" json:"acepta_marketing"`

	// Metadata
	IPOrigen  string `gorm:"type:varchar(45)" json:"ip_origen"`
	UserAgent string `gorm:"type:text" json:"user_agent"`
	Idioma    string `gorm:"type:varchar(10);default:'es'" json:"idioma"`

	// Timestamps
	FechaCreacion      time.Time `gorm:"autoCreateTime" json:"fecha_creacion"`
	FechaActualizacion time.Time `gorm:"autoUpdateTime" json:"fecha_actualizacion"`
}

// TableName especifica el nombre de la tabla en la base de datos
func (FormularioSoporte) TableName() string {
	return "formularios_soporte"
}
