# 🔐 Configuración de Seguridad

## Variables de Entorno

Este proyecto utiliza variables de entorno para manejar información sensible. **NUNCA** subas el archivo `.env` a GitHub.

### Configuración para Producción

1. Copia el archivo de ejemplo:
   ```bash
   cp backend/.env.example backend/.env
   ```

2. Edita `backend/.env` con tus credenciales reales:
   ```env
   DB_HOST=tu_servidor_mysql
   DB_USER=tu_usuario
   DB_PASSWORD=tu_password_seguro
   DB_NAME=tu_base_de_datos
   ```

3. **IMPORTANTE**: El archivo `.env` está en `.gitignore` y nunca debe ser versionado.

## Credenciales de Base de Datos

Las credenciales de la base de datos deben ser proporcionadas por tu proveedor de hosting. Asegúrate de:

- Usar contraseñas fuertes
- Limitar los permisos del usuario de DB solo a lo necesario (SELECT, INSERT, UPDATE, DELETE)
- Nunca compartir credenciales en código fuente

## CORS

En producción, configura `ALLOWED_ORIGINS` con tu dominio específico:

```env
ALLOWED_ORIGINS=https://tudominio.com,https://www.tudominio.com
```

No uses `*` en producción.

## HTTPS

Siempre usa HTTPS en producción. Configura certificados SSL con Let's Encrypt:

```bash
sudo certbot --nginx -d tudominio.com
```
