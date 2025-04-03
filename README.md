# Universidad del Valle de Guatemala
# Sistemas y Tecnologías Web - sección 20
# Roberto Barreda #23354

# Series Tracker  
Programa para gestionar una lista de series o películas 

## **Instalación**  
**Requisitos**:  
- Docker instalado ([Guía de instalación](https://docs.docker.com/get-docker/)).  
- Git o Github Desktop (opcional, para clonar el repositorio).  

**Pasos**:  
1. **Clonar el repositorio** (o descargar los archivos):  
   ```bash  
   git clone [URL_DEL_REPOSITORIO]  
   cd [NOMBRE_DEL_REPOSITORIO]  
2. **Instalar dependencias del backend (Go):**
    cd backend  
    go mod tidy
3. **Generar documentación Swagger:**
    go install github.com/swaggo/swag/cmd/swag@latest  
    swag init -g main.go -o docs  
4. **Construir y ejecutar con Docker Compose:**
    docker-compose up --build  