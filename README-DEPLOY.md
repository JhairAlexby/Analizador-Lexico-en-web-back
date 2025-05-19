# Instrucciones para Desplegar en Railway

Este documento contiene las instrucciones para desplegar correctamente el Analizador Léxico en Railway.

## Requisitos Previos

- Cuenta en [Railway](https://railway.app/)
- Git instalado en tu sistema
- Repositorio subido a GitHub

## Pasos para Desplegar

1. **Inicia sesión en Railway**
   - Ve a [Railway](https://railway.app/) y accede con tu cuenta

2. **Crea un nuevo proyecto**
   - Haz clic en "New Project"
   - Selecciona "Deploy from GitHub repo"
   - Conecta tu repositorio de GitHub y selecciona el repositorio del Analizador Léxico

3. **Configuración del Despliegue**
   - Railway detectará automáticamente el Dockerfile en tu repositorio
   - No es necesario configurar variables de entorno adicionales, ya que la aplicación usa el puerto proporcionado por Railway

4. **Despliegue**
   - Haz clic en "Deploy" y espera a que Railway construya y despliegue tu aplicación
   - Una vez completado, Railway te proporcionará una URL para acceder a tu aplicación

## Solución de Problemas

Si encuentras el error `/bin/bash: line 1: go: command not found`, significa que Railway está intentando ejecutar el comando Go directamente en lugar de usar Docker. Asegúrate de que:

1. El Dockerfile está presente en la raíz del proyecto
2. El Procfile está configurado correctamente con `web: ./main`
3. Railway está configurado para usar Docker

## Estructura de Archivos para Despliegue

```
├── Dockerfile          # Configuración para construir la imagen Docker
├── .dockerignore       # Archivos a ignorar en la construcción de Docker
├── Procfile            # Instrucciones para Railway sobre cómo ejecutar la aplicación
├── go.mod              # Dependencias de Go
└── main.go             # Código fuente principal
```

## Verificación del Despliegue

Para verificar que tu aplicación se ha desplegado correctamente, puedes hacer una solicitud POST a la URL proporcionada por Railway:

```bash
curl -X POST -H "Content-Type: application/json" -d '{"text":"juan 890 juan"}' https://tu-url-de-railway.app/analyze
```

Deberías recibir una respuesta JSON con el análisis léxico del texto proporcionado.