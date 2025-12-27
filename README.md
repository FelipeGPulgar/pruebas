# Sistema de Formularios Logitech 🎮

Sistema completo de formularios de soporte técnico para productos Logitech con backend en Go y frontend en SvelteKit.

## 🚀 Características

- ✅ Formulario responsivo con diseño premium
- ✅ Validación en tiempo real
- ✅ Todos los tipos de campos (text, email, tel, select, textarea, radio, checkbox, date)
- ✅ Backend robusto con Go + Fiber + GORM
- ✅ Base de datos MySQL optimizada
- ✅ API RESTful documentada
- ✅ Diseño moderno con Tailwind CSS
- ✅ Mensajes de éxito/error en español
- ✅ Captura de metadata (IP, User-Agent)

## 📋 Requisitos

- Go 1.21+
- Bun (para el frontend)
- MySQL 5.7+
- Servidor web (Nginx/Apache)

## 🏗️ Estructura del Proyecto

```
logitech-form/
├── backend/          # API en Go (Fiber + GORM)
├── frontend/         # SvelteKit + Tailwind CSS
├── database/         # Scripts SQL
└── docs/            # Documentación completa
```

## 🎯 Quick Start

### 1. Base de Datos
```bash
mysql -u hxkmypwcga_upforms -p hxkmypwcga_contactanos < database/create_table.sql
```

### 2. Backend
```bash
cd backend
go mod download
go build -o logitech-api main.go
./logitech-api
```

### 3. Frontend
```bash
cd frontend
bun install
bun run dev
```

Accede a: `http://localhost:5173`

## 📚 Documentación Completa

Lee la [Guía de Deployment](docs/DEPLOYMENT.md) para instrucciones detalladas de:
- Configuración de la base de datos
- Compilación para producción
- Deployment en hosting
- Configuración de servidor web
- Troubleshooting
- Seguridad y optimización

## 🗄️ Credenciales de Base de Datos

```
Database: hxkmypwcga_contactanos
User: hxkmypwcga_upforms
Password: Aksl3ic.,92jw@
Permissions: SELECT, INSERT, UPDATE, DELETE
```

## 🌐 API Endpoints

- `GET /api/v1/` - Health check
- `POST /api/v1/formularios` - Crear formulario
- `GET /api/v1/formularios` - Listar formularios
- `GET /api/v1/formularios/:id` - Obtener formulario por ID

## 🎨 Campos del Formulario

### Información Personal
- Nombre completo (text) *
- Email (email) *
- Teléfono (tel)

### Información del Producto
- Tipo de producto (select) *
- Modelo del producto (text)
- Número de serie (text)
- Fecha de compra (date)
- Lugar de compra (text)

### Detalles del Problema
- Categoría del problema (select) *
- Descripción del problema (textarea) *
- Sistema operativo (select)

### Preferencias de Contacto
- Método de contacto preferido (radio)
- Horario de contacto (select)

### Términos
- Acepta términos y condiciones (checkbox) *
- Acepta marketing (checkbox)

\* Campos requeridos

## 🛠️ Stack Tecnológico

**Backend:**
- Go 1.21
- Fiber (Framework web)
- GORM (ORM)
- MySQL Driver
- godotenv (Variables de entorno)

**Frontend:**
- SvelteKit 2.0
- Tailwind CSS 3.3
- Vite 5.0
- TypeScript

## 📦 Deployment

Para crear el paquete de deployment:

```bash
# Compilar backend para Linux
cd backend
GOOS=linux GOARCH=amd64 go build -o logitech-api main.go

# Compilar frontend
cd ../frontend
bun run build

# Crear archivo comprimido
cd ..
tar -czf logitech-form-deploy.tar.gz backend/logitech-api backend/.env frontend/build/ database/ docs/
```

Sube `logitech-form-deploy.tar.gz` a tu servidor y sigue las instrucciones en [DEPLOYMENT.md](docs/DEPLOYMENT.md).

## 🔒 Seguridad

- Validación de datos en frontend y backend
- Sanitización de inputs
- CORS configurable
- Preparado para HTTPS
- Variables de entorno para credenciales

## 🐛 Troubleshooting

Ver sección de Troubleshooting en [DEPLOYMENT.md](docs/DEPLOYMENT.md#-troubleshooting).

## 📄 Licencia

Proyecto de testing para manejo de formularios en hosting.

---

**Creado para testing de formularios con Go + Fiber + GORM + SvelteKit + Tailwind CSS**
