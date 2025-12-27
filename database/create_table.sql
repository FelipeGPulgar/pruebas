-- Script para crear la tabla de formularios de contacto Logitech
-- Base de datos: hxkmypwcga_contactanos
-- Usuario: hxkmypwcga_upforms

USE hxkmypwcga_contactanos;

-- Eliminar tabla si existe (solo para desarrollo/testing)
-- DROP TABLE IF EXISTS formularios_soporte;

CREATE TABLE IF NOT EXISTS formularios_soporte (
    id INT AUTO_INCREMENT PRIMARY KEY,
    
    -- Información personal
    nombre_completo VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    telefono VARCHAR(50),
    
    -- Información del producto
    tipo_producto VARCHAR(100) NOT NULL,
    modelo_producto VARCHAR(100),
    numero_serie VARCHAR(100),
    
    -- Detalles del problema
    categoria_problema VARCHAR(100) NOT NULL,
    descripcion_problema TEXT NOT NULL,
    
    -- Información adicional
    sistema_operativo VARCHAR(100),
    fecha_compra DATE,
    lugar_compra VARCHAR(255),
    
    -- Preferencias de contacto
    metodo_contacto_preferido ENUM('email', 'telefono', 'ambos') DEFAULT 'email',
    horario_contacto VARCHAR(100),
    
    -- Archivos adjuntos (URLs o paths)
    archivo_adjunto VARCHAR(500),
    
    -- Aceptación de términos
    acepta_terminos BOOLEAN DEFAULT FALSE,
    acepta_marketing BOOLEAN DEFAULT FALSE,
    
    -- Metadata
    ip_origen VARCHAR(45),
    user_agent TEXT,
    idioma VARCHAR(10) DEFAULT 'es',
    
    -- Timestamps
    fecha_creacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    fecha_actualizacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    
    -- Índices para búsquedas rápidas
    INDEX idx_email (email),
    INDEX idx_fecha_creacion (fecha_creacion),
    INDEX idx_categoria (categoria_problema),
    INDEX idx_tipo_producto (tipo_producto)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Comentarios de la tabla
ALTER TABLE formularios_soporte COMMENT = 'Tabla para almacenar formularios de soporte técnico de productos Logitech';
