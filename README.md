# Analizador Léxico - Backend

Este es el backend para un analizador léxico que identifica y clasifica tokens según reglas específicas. El proyecto está configurado para ser desplegado en Railway.

## Reglas de clasificación

- **Identificador válido**: Cadenas que contienen solo letras (ej. "juan", "Upchiapas")
- **Número válido**: Cadenas que contienen solo dígitos (ej. "9890")
- **Error/No válido**: Cadenas que mezclan letras y números (ej. "9890LUpchiapas")

## Funcionalidades

- Tokeniza una cadena de texto separándola por espacios
- Clasifica cada token según las reglas anteriores
- Cuenta el total de cada tipo de token
- Identifica tokens únicos (no cuenta repeticiones)

## Cómo ejecutar

### Localmente

1. Asegúrate de tener Go instalado en tu sistema
2. Navega al directorio del proyecto
3. Ejecuta el servidor:

```bash
go run main.go
```

El servidor se iniciará en http://localhost:8080

### Despliegue en Railway

1. Asegúrate de tener una cuenta en [Railway](https://railway.app/)
2. Conecta tu repositorio de GitHub a Railway
3. Railway detectará automáticamente que es un proyecto Go y utilizará el Procfile para iniciar la aplicación
4. La variable de entorno `PORT` será configurada automáticamente por Railway

## API

### Endpoint: POST /analyze

**Solicitud**:

```json
{
  "text": "juan 890 juan 9890LUpchiapas"
}
```

**Respuesta**:

```json
{
  "tokens": [
    {"value": "juan", "type": "identificador"},
    {"value": "890", "type": "numero"},
    {"value": "juan", "type": "identificador"},
    {"value": "9890LUpchiapas", "type": "error"}
  ],
  "totalsByType": {
    "identificador": 2,
    "identificador_unicos": 1,
    "numero": 1,
    "numero_unicos": 1,
    "error": 1
  },
  "uniqueTokens": {
    "890": {"value": "890", "type": "numero"},
    "juan": {"value": "juan", "type": "identificador"}
  }
}
```

## Ejemplo

Para el texto "juan 890 juan", el resultado mostrará:
- Total identificador: 1 (único)
- Total número: 1 (único)
- Total tokens de error: 0