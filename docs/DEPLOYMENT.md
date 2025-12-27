# 📋 Guía de Deployment - Sistema de Formularios Logitech

## 🎯 Descripción del Proyecto

Sistema completo de formularios de soporte técnico para Logitech con:
- **Backend**: Go + Fiber + GORM
- **Frontend**: SvelteKit + Tailwind CSS
- **Base de datos**: MySQL

---

## 📁 Estructura del Proyecto

```
logitech-form/
├── backend/                 # API en Go
│   ├── main.go             # Punto de entrada del servidor
│   ├── go.mod              # Dependencias de Go
│   ├── .env                # Variables de entorno
│   ├── models/             # Modelos GORM
│   │   └── formulario.go
│   ├── handlers/           # Controladores de rutas
│   │   └── formulario_handler.go
│   └── database/           # Configuración de DB
│       └── database.go
├── frontend/               # Aplicación SvelteKit
│   ├── src/
│   │   ├── routes/
│   │   │   ├── +page.svelte      # Formulario principal
│   │   │   └── +layout.svelte
│   │   └── app.css
│   ├── package.json
│   ├── svelte.config.js
│   ├── vite.config.js
│   └── tailwind.config.js
├── database/               # Scripts SQL
│   └── create_table.sql
└── docs/                   # Documentación
    └── DEPLOYMENT.md
```

---

## 🗄️ PASO 1: Configurar la Base de Datos

### 1.1 Crear la Tabla

Ejecuta el script SQL en tu base de datos MySQL:

```bash
mysql -u hxkmypwcga_upforms -p hxkmypwcga_contactanos < database/create_table.sql
```

O desde phpMyAdmin:
1. Accede a phpMyAdmin en tu hosting
2. Selecciona la base de datos `hxkmypwcga_contactanos`
3. Ve a la pestaña "SQL"
4. Copia y pega el contenido de `database/create_table.sql`
5. Haz clic en "Ejecutar"

### 1.2 Verificar la Tabla

```sql
USE hxkmypwcga_contactanos;
SHOW TABLES;
DESCRIBE formularios_soporte;
```

---

## 🔧 PASO 2: Configurar el Backend (Go)

### 2.1 Instalar Go

Si no tienes Go instalado:
```bash
# macOS
brew install go

# Linux
wget https://go.dev/dl/go1.21.0.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
```

### 2.2 Configurar Variables de Entorno

Edita el archivo `backend/.env` con tus credenciales:

```env
PORT=3000
ENVIRONMENT=production
DB_HOST=localhost
DB_PORT=3306
DB_USER=hxkmypwcga_upforms
DB_PASSWORD=Aksl3ic.,92jw@
DB_NAME=hxkmypwcga_contactanos
ALLOWED_ORIGINS=*
LOG_LEVEL=info
```

### 2.3 Instalar Dependencias

```bash
cd backend
go mod download
```

### 2.4 Compilar el Backend

**Para Linux (hosting):**
```bash
GOOS=linux GOARCH=amd64 go build -o logitech-api main.go
```

**Para macOS (local):**
```bash
go build -o logitech-api main.go
```

**Para Windows:**
```bash
GOOS=windows GOARCH=amd64 go build -o logitech-api.exe main.go
```

### 2.5 Ejecutar el Backend

```bash
./logitech-api
```

Deberías ver:
```
✅ Conexión a la base de datos establecida exitosamente
🚀 Servidor iniciado en http://localhost:3000
```

---

## 🎨 PASO 3: Configurar el Frontend (SvelteKit)

### 3.1 Instalar Bun (si no lo tienes)

```bash
curl -fsSL https://bun.sh/install | bash
```

### 3.2 Instalar Dependencias

```bash
cd frontend
bun install
```

### 3.3 Configurar la URL del API

Edita `frontend/vite.config.js` y cambia la URL del proxy si es necesario:

```javascript
proxy: {
  '/api': {
    target: 'http://tu-dominio.com:3000',  // Cambia esto en producción
    changeOrigin: true
  }
}
```

### 3.4 Compilar el Frontend

```bash
bun run build
```

Esto generará la carpeta `build/` con los archivos estáticos.

### 3.5 Ejecutar en Desarrollo

```bash
bun run dev
```

Accede a: `http://localhost:5173`

---

## 🚀 PASO 4: Deployment en el Host

### 4.1 Preparar Archivos para Subir

Crea un archivo comprimido con todo lo necesario:

```bash
# Desde la raíz del proyecto
tar -czf logitech-form-deploy.tar.gz \
  backend/logitech-api \
  backend/.env \
  frontend/build/ \
  database/create_table.sql \
  docs/DEPLOYMENT.md
```

### 4.2 Subir al Servidor

**Opción A: Por FTP/SFTP**
1. Conecta con FileZilla o tu cliente FTP favorito
2. Sube `logitech-form-deploy.tar.gz` al servidor
3. Descomprime en el servidor:
   ```bash
   tar -xzf logitech-form-deploy.tar.gz
   ```

**Opción B: Por SSH**
```bash
scp logitech-form-deploy.tar.gz usuario@tu-servidor.com:/ruta/destino/
ssh usuario@tu-servidor.com
cd /ruta/destino/
tar -xzf logitech-form-deploy.tar.gz
```

### 4.3 Configurar el Servidor Web

**Para Nginx:**

```nginx
server {
    listen 80;
    server_name tu-dominio.com;

    # Frontend (archivos estáticos)
    location / {
        root /ruta/a/frontend/build;
        try_files $uri $uri/ /index.html;
    }

    # Backend API
    location /api/ {
        proxy_pass http://localhost:3000;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection 'upgrade';
        proxy_set_header Host $host;
        proxy_cache_bypass $http_upgrade;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    }
}
```

**Para Apache (.htaccess):**

```apache
RewriteEngine On

# Redirigir API al backend
RewriteRule ^api/(.*)$ http://localhost:3000/api/$1 [P,L]

# SPA routing
RewriteCond %{REQUEST_FILENAME} !-f
RewriteCond %{REQUEST_FILENAME} !-d
RewriteRule ^(.*)$ index.html [L]
```

### 4.4 Ejecutar el Backend como Servicio

**Opción A: systemd (Linux)**

Crea `/etc/systemd/system/logitech-api.service`:

```ini
[Unit]
Description=Logitech Form API
After=network.target

[Service]
Type=simple
User=www-data
WorkingDirectory=/ruta/a/backend
ExecStart=/ruta/a/backend/logitech-api
Restart=always
RestartSec=10

[Install]
WantedBy=multi-user.target
```

Luego:
```bash
sudo systemctl daemon-reload
sudo systemctl enable logitech-api
sudo systemctl start logitech-api
sudo systemctl status logitech-api
```

**Opción B: PM2 (Node.js process manager)**

```bash
npm install -g pm2
pm2 start backend/logitech-api --name logitech-api
pm2 save
pm2 startup
```

**Opción C: Screen (simple)**

```bash
screen -S logitech-api
cd backend
./logitech-api
# Presiona Ctrl+A, luego D para detach
```

---

## 🧪 PASO 5: Verificar el Deployment

### 5.1 Probar el Backend

```bash
# Health check
curl http://tu-dominio.com/api/v1/

# Debería responder:
{
  "mensaje": "API de Formularios Logitech - Funcionando correctamente",
  "version": "1.0.0",
  "status": "ok"
}
```

### 5.2 Probar el Formulario

1. Accede a `http://tu-dominio.com`
2. Llena el formulario con datos de prueba
3. Envía el formulario
4. Verifica que aparezca el mensaje de éxito

### 5.3 Verificar en la Base de Datos

```sql
USE hxkmypwcga_contactanos;
SELECT * FROM formularios_soporte ORDER BY fecha_creacion DESC LIMIT 5;
```

---

## 🔒 PASO 6: Seguridad y Optimización

### 6.1 Configurar HTTPS

```bash
# Con Certbot (Let's Encrypt)
sudo apt install certbot python3-certbot-nginx
sudo certbot --nginx -d tu-dominio.com
```

### 6.2 Configurar CORS en Producción

Edita `backend/.env`:
```env
ALLOWED_ORIGINS=https://tu-dominio.com,https://www.tu-dominio.com
```

### 6.3 Configurar Firewall

```bash
# Permitir solo puertos necesarios
sudo ufw allow 80/tcp
sudo ufw allow 443/tcp
sudo ufw allow 22/tcp
sudo ufw enable
```

### 6.4 Backups Automáticos de la DB

```bash
# Crear script de backup
cat > /home/usuario/backup-db.sh << 'EOF'
#!/bin/bash
DATE=$(date +%Y%m%d_%H%M%S)
mysqldump -u hxkmypwcga_upforms -p'Aksl3ic.,92jw@' hxkmypwcga_contactanos > /backups/logitech_$DATE.sql
find /backups -name "logitech_*.sql" -mtime +7 -delete
EOF

chmod +x /home/usuario/backup-db.sh

# Agregar a crontab (diario a las 2 AM)
crontab -e
0 2 * * * /home/usuario/backup-db.sh
```

---

## 🐛 Troubleshooting

### Problema: "Error al conectar a la base de datos"

**Solución:**
1. Verifica las credenciales en `.env`
2. Asegúrate de que el usuario tenga permisos:
   ```sql
   GRANT SELECT, INSERT, UPDATE, DELETE ON hxkmypwcga_contactanos.* TO 'hxkmypwcga_upforms'@'localhost';
   FLUSH PRIVILEGES;
   ```

### Problema: "CORS error" en el frontend

**Solución:**
Edita `backend/.env` y agrega tu dominio:
```env
ALLOWED_ORIGINS=https://tu-dominio.com
```

### Problema: El formulario no se envía

**Solución:**
1. Abre la consola del navegador (F12)
2. Verifica que la URL del API sea correcta
3. Comprueba que el backend esté corriendo:
   ```bash
   curl http://localhost:3000/api/v1/
   ```

### Problema: "Port 3000 already in use"

**Solución:**
```bash
# Encuentra el proceso
lsof -i :3000
# Mata el proceso
kill -9 <PID>
# O cambia el puerto en .env
PORT=3001
```

---

## 📊 Endpoints de la API

### POST /api/v1/formularios
Crea un nuevo formulario de soporte.

**Request:**
```json
{
  "nombre_completo": "Juan Pérez",
  "email": "juan@ejemplo.com",
  "telefono": "+56912345678",
  "tipo_producto": "Ratón",
  "modelo_producto": "MX Master 3S",
  "numero_serie": "SN123456",
  "categoria_problema": "Problema de conexión",
  "descripcion_problema": "El ratón no se conecta por Bluetooth...",
  "sistema_operativo": "Windows 11",
  "fecha_compra": "2024-01-15",
  "lugar_compra": "Amazon",
  "metodo_contacto_preferido": "email",
  "horario_contacto": "Mañana (9:00 - 12:00)",
  "acepta_terminos": true,
  "acepta_marketing": false
}
```

**Response (200):**
```json
{
  "error": false,
  "mensaje": "Formulario enviado exitosamente...",
  "id": 123
}
```

### GET /api/v1/formularios
Obtiene todos los formularios (para uso interno).

**Response (200):**
```json
{
  "error": false,
  "total": 10,
  "formularios": [...]
}
```

### GET /api/v1/formularios/:id
Obtiene un formulario específico por ID.

**Response (200):**
```json
{
  "error": false,
  "formulario": {...}
}
```

---

## 📝 Notas para Desarrolladores

### Agregar Nuevos Campos al Formulario

1. **Actualizar la base de datos:**
   ```sql
   ALTER TABLE formularios_soporte ADD COLUMN nuevo_campo VARCHAR(255);
   ```

2. **Actualizar el modelo Go:**
   ```go
   // backend/models/formulario.go
   type FormularioSoporte struct {
       // ... campos existentes
       NuevoCampo string `gorm:"type:varchar(255)" json:"nuevo_campo"`
   }
   ```

3. **Actualizar el formulario Svelte:**
   ```svelte
   <!-- frontend/src/routes/+page.svelte -->
   <input bind:value={formData.nuevo_campo} />
   ```

### Cambiar el Puerto del Backend

Edita `backend/.env`:
```env
PORT=8080
```

### Habilitar Logs de Debug

Edita `backend/.env`:
```env
ENVIRONMENT=development
LOG_LEVEL=debug
```

---

## 🎉 ¡Listo!

Tu sistema de formularios Logitech está ahora desplegado y funcionando.

**Próximos pasos sugeridos:**
- [ ] Configurar notificaciones por email cuando llegue un formulario
- [ ] Crear panel de administración para ver formularios
- [ ] Agregar autenticación para endpoints GET
- [ ] Implementar rate limiting
- [ ] Agregar tests automatizados

---

## 📞 Soporte

Si tienes problemas, revisa:
1. Logs del backend: `journalctl -u logitech-api -f`
2. Logs de Nginx: `tail -f /var/log/nginx/error.log`
3. Consola del navegador (F12)

**Creado con ❤️ para testing de formularios**
