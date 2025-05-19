# Usar una imagen oficial de Go como base
FROM golang:1.20-alpine

# Establecer el directorio de trabajo en el contenedor
WORKDIR /app

# Copiar los archivos go.mod y go.sum primero para aprovechar la caché de Docker
COPY go.mod ./

# Copiar el código fuente
COPY . .

# Compilar la aplicación
RUN go build -o main .

# Exponer el puerto que usa la aplicación
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./main"]