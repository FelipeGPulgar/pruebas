# 🚀 Quick Start - Sistema de Formularios Logitech

## ⚡ Instalación Rápida (5 minutos)

### 1️⃣ Clonar el Repositorio
```bash
git clone https://github.com/FelipeGPulgar/pruebas.git
cd pruebas
```

### 2️⃣ Configurar Base de Datos
```bash
# Ejecutar el script SQL en tu servidor MySQL
mysql -u tu_usuario -p tu_base_de_datos < database/create_table.sql
```

### 3️⃣ Configurar Variables de Entorno
```bash
# Copiar el archivo de ejemplo
cp backend/.env.example backend/.env

# Editar con tus credenciales
nano backend/.env
```

**Edita estos valores:**
```env
DB_HOST=tu_servidor_mysql
DB_USER=tu_usuario
DB_PASSWORD=tu_password
DB_NAME=tu_base_de_datos
```

### 4️⃣ Instalar Dependencias

**Backend (Go):**
```bash
cd backend
go mod download
```

**Frontend (Bun):**
```bash
cd frontend
bun install
```

### 5️⃣ Ejecutar en Desarrollo

**Terminal 1 - Backend:**
```bash
cd backend
go run main.go
# Servidor en http://localhost:3000
```

**Terminal 2 - Frontend:**
```bash
cd frontend
bun run dev
# Aplicación en http://localhost:5173
```

---

## 📦 Compilar para Producción

### Opción A: Script Automático
```bash
chmod +x build.sh
./build.sh
# Genera: logitech-form-deploy.tar.gz
```

### Opción B: Manual

**Backend:**
```bash
cd backend
GOOS=linux GOARCH=amd64 go build -o logitech-api main.go
```

**Frontend:**
```bash
cd frontend
bun run build
# Archivos en: build/
```

---

## 🌐 Deployment en Hosting

### 1. Subir Archivos
```bash
# Subir por SCP
scp logitech-form-deploy.tar.gz usuario@host:/ruta/

# O usar FTP/FileZilla
```

### 2. Descomprimir en Servidor
```bash
ssh usuario@host
cd /ruta/
tar -xzf logitech-form-deploy.tar.gz
```

### 3. Configurar y Ejecutar
```bash
# Editar .env con credenciales del host
nano .env

# Ejecutar backend
./logitech-api
```

### 4. Configurar Nginx
```nginx
server {
    listen 80;
    server_name tudominio.com;

    location / {
        root /ruta/a/frontend-build;
        try_files $uri $uri/ /index.html;
    }

    location /api/ {
        proxy_pass http://localhost:3000;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
    }
}
```

---

## 🧪 Probar la Instalación

### Backend
```bash
curl http://localhost:3000/api/v1/
# Debe responder: {"mensaje":"API de Formularios Logitech...","status":"ok"}
```

### Frontend
Abre en navegador: `http://localhost:5173`

### Enviar Formulario de Prueba
1. Llena todos los campos requeridos (*)
2. Haz clic en "Enviar Formulario"
3. Verifica mensaje de éxito
4. Revisa en la base de datos:
   ```sql
   SELECT * FROM formularios_soporte ORDER BY fecha_creacion DESC LIMIT 1;
   ```

---

## 📚 Documentación Completa

- **[README.md](README.md)** - Información general del proyecto
- **[docs/DEPLOYMENT.md](docs/DEPLOYMENT.md)** - Guía detallada de deployment (500+ líneas)
- **[docs/SECURITY.md](docs/SECURITY.md)** - Configuración de seguridad

---

## 🛠️ Stack Tecnológico

| Componente | Tecnología |
|------------|------------|
| Backend | Go 1.21 + Fiber + GORM |
| Frontend | SvelteKit + Tailwind CSS |
| Base de Datos | MySQL 5.7+ |
| Runtime | Bun (frontend) |

---

## 📋 Campos del Formulario

✅ **Todos los tipos de campos HTML:**
- `text` - Nombre, modelo, número de serie
- `email` - Correo electrónico
- `tel` - Teléfono
- `date` - Fecha de compra
- `select` - Tipo de producto, categoría, SO
- `textarea` - Descripción del problema
- `radio` - Método de contacto
- `checkbox` - Términos y marketing

---

## 🔒 Seguridad

✅ **Implementado:**
- Variables de entorno para credenciales
- Validación frontend y backend
- CORS configurable
- Sanitización de inputs
- HTTPS ready
- `.gitignore` configurado (no sube credenciales)

---

## 🐛 Problemas Comunes

### "Error al conectar a la base de datos"
- Verifica credenciales en `.env`
- Asegúrate de que MySQL esté corriendo
- Verifica permisos del usuario de DB

### "CORS error"
- Configura `ALLOWED_ORIGINS` en `.env`
- En desarrollo usa `*`
- En producción usa tu dominio específico

### "Puerto 3000 en uso"
- Cambia `PORT` en `.env`
- O mata el proceso: `lsof -i :3000` y `kill -9 <PID>`

---

## 📞 Soporte

- **Documentación completa**: [docs/DEPLOYMENT.md](docs/DEPLOYMENT.md)
- **GitHub Issues**: [Reportar problema](https://github.com/FelipeGPulgar/pruebas/issues)

---

## ✅ Checklist de Deployment

- [ ] Base de datos creada y tabla ejecutada
- [ ] Variables de entorno configuradas (`.env`)
- [ ] Dependencias instaladas (Go + Bun)
- [ ] Backend compilado y corriendo
- [ ] Frontend compilado
- [ ] Servidor web configurado (Nginx/Apache)
- [ ] HTTPS configurado (Let's Encrypt)
- [ ] CORS configurado con dominio específico
- [ ] Formulario probado y funcionando
- [ ] Datos guardándose en la base de datos

---

**¡Listo para usar!** 🎉

Para más detalles, lee la [documentación completa](docs/DEPLOYMENT.md).
