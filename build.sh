#!/bin/bash

echo "🚀 Iniciando compilación del proyecto Logitech Form..."

# Colores para output
GREEN='\033[0;32m'
BLUE='\033[0;34m'
RED='\033[0;31m'
NC='\033[0m' # No Color

# Crear directorio de deployment
echo -e "${BLUE}📁 Creando directorio de deployment...${NC}"
mkdir -p deployment
rm -rf deployment/*

# Compilar Backend
echo -e "${BLUE}🔨 Compilando backend para Linux...${NC}"
cd backend
GOOS=linux GOARCH=amd64 go build -o logitech-api main.go
if [ $? -eq 0 ]; then
    echo -e "${GREEN}✅ Backend compilado exitosamente${NC}"
    cp logitech-api ../deployment/
    cp .env.example ../deployment/.env
    cp -r models ../deployment/
    cp -r handlers ../deployment/
    cp -r database ../deployment/
else
    echo -e "${RED}❌ Error al compilar el backend${NC}"
    exit 1
fi
cd ..

# Compilar Frontend
echo -e "${BLUE}🎨 Compilando frontend...${NC}"
cd frontend
bun install
bun run build
if [ $? -eq 0 ]; then
    echo -e "${GREEN}✅ Frontend compilado exitosamente${NC}"
    cp -r build ../deployment/frontend-build
else
    echo -e "${RED}❌ Error al compilar el frontend${NC}"
    exit 1
fi
cd ..

# Copiar archivos necesarios
echo -e "${BLUE}📋 Copiando archivos de configuración...${NC}"
cp -r database deployment/
cp -r docs deployment/
cp README.md deployment/

# Crear archivo comprimido
echo -e "${BLUE}📦 Creando archivo comprimido...${NC}"
cd deployment
tar -czf ../logitech-form-deploy.tar.gz .
cd ..

echo -e "${GREEN}✅ ¡Deployment preparado exitosamente!${NC}"
echo -e "${GREEN}📦 Archivo: logitech-form-deploy.tar.gz${NC}"
echo -e "${BLUE}📖 Lee docs/DEPLOYMENT.md para instrucciones de instalación${NC}"
