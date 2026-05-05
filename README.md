# Usuarios Agro CRUD

API REST desarrollada en Go para gestionar usuarios y datos relacionados dentro del esquema `"Usuarios"` de la base de datos PostgreSQL `Agrocampo`.

El proyecto expone operaciones CRUD para roles, usuarios, contrasenas, tokens de recuperacion, verificacion en dos pasos, auditoria de usuarios y perfiles extendidos.

## Tecnologias

- Go 1.22
- PostgreSQL
- Gorilla Mux
- Driver `github.com/lib/pq`

## Estructura

```text
config/       Conexion a PostgreSQL
controllers/  Logica HTTP y consultas SQL
models/       Estructuras JSON de cada tabla
routes/       Registro de endpoints
main.go       Punto de entrada del servidor
```

## Conexion a base de datos

La conexion esta configurada en `config/db.go`:

```text
host: localhost
port: 5432
user: postgres
password: postgres
database: Agrocampo
schema: "Usuarios"
```

Antes de ejecutar la API, verifica que PostgreSQL este activo, que exista la base de datos `Agrocampo` y que las tablas necesarias esten dentro del esquema `"Usuarios"`.

## Tablas gestionadas

- `Rol`: administra roles de usuario.
- `Usuario`: administra usuarios principales del sistema.
- `Contrasena`: guarda hashes de contrasenas asociadas a usuarios.
- `TokenRecuperacion`: gestiona tokens para recuperacion de cuenta.
- `VerificacionDosPasos`: gestiona metodos y codigos OTP.
- `AuditoriaUsuario`: registra eventos o acciones de usuarios.
- `PerfilExtendido`: guarda informacion adicional del usuario.

## Ejecutar el proyecto

Instala dependencias:

```bash
go mod tidy
```

Ejecuta el servidor:

```bash
go run .
```

El servidor queda disponible en:

```text
http://localhost:8094
```

## Endpoints

Cada recurso tiene las mismas operaciones CRUD:

| Metodo | Ruta | Funcion |
| --- | --- | --- |
| GET | `/rol` | Listar roles |
| GET | `/rol/{id}` | Obtener rol por ID |
| POST | `/rol` | Crear rol |
| PUT | `/rol/{id}` | Actualizar rol |
| DELETE | `/rol/{id}` | Eliminar rol |
| GET | `/usuario` | Listar usuarios |
| GET | `/usuario/{id}` | Obtener usuario por ID |
| POST | `/usuario` | Crear usuario |
| PUT | `/usuario/{id}` | Actualizar usuario |
| DELETE | `/usuario/{id}` | Eliminar usuario |
| GET | `/contrasena` | Listar contrasenas |
| GET | `/contrasena/{id}` | Obtener contrasena por ID |
| POST | `/contrasena` | Crear contrasena |
| PUT | `/contrasena/{id}` | Actualizar contrasena |
| DELETE | `/contrasena/{id}` | Eliminar contrasena |
| GET | `/token_recuperacion` | Listar tokens |
| GET | `/token_recuperacion/{id}` | Obtener token por ID |
| POST | `/token_recuperacion` | Crear token |
| PUT | `/token_recuperacion/{id}` | Actualizar token |
| DELETE | `/token_recuperacion/{id}` | Eliminar token |
| GET | `/verificacion_dos_pasos` | Listar verificaciones |
| GET | `/verificacion_dos_pasos/{id}` | Obtener verificacion por ID |
| POST | `/verificacion_dos_pasos` | Crear verificacion |
| PUT | `/verificacion_dos_pasos/{id}` | Actualizar verificacion |
| DELETE | `/verificacion_dos_pasos/{id}` | Eliminar verificacion |
| GET | `/auditoria_usuario` | Listar auditorias |
| GET | `/auditoria_usuario/{id}` | Obtener auditoria por ID |
| POST | `/auditoria_usuario` | Crear auditoria |
| PUT | `/auditoria_usuario/{id}` | Actualizar auditoria |
| DELETE | `/auditoria_usuario/{id}` | Eliminar auditoria |
| GET | `/perfil_extendido` | Listar perfiles extendidos |
| GET | `/perfil_extendido/{id}` | Obtener perfil por ID |
| POST | `/perfil_extendido` | Crear perfil |
| PUT | `/perfil_extendido/{id}` | Actualizar perfil |
| DELETE | `/perfil_extendido/{id}` | Eliminar perfil |

## Ejemplos de uso

Crear un rol:

```bash
curl -X POST http://localhost:8094/rol \
  -H "Content-Type: application/json" \
  -d "{\"nombre_rol\":\"Administrador\",\"descripcion\":\"Acceso completo\",\"activo\":true}"
```

Listar roles:

```bash
curl http://localhost:8094/rol
```

Crear un usuario:

```bash
curl -X POST http://localhost:8094/usuario \
  -H "Content-Type: application/json" \
  -d "{\"nombre_completo\":\"Ana Perez\",\"correo\":\"ana@example.com\",\"telefono\":\"3001234567\",\"id_rol\":1,\"verificacion_dos_pasos\":false,\"activo\":true}"
```

## Respuestas

Las respuestas se devuelven en formato JSON.

Respuesta correcta:

```json
{
  "id_rol": 1,
  "nombre_rol": "Administrador",
  "descripcion": "Acceso completo",
  "activo": true,
  "fecha_creacion": "2026-05-04T20:00:00Z",
  "fecha_modificacion": null
}
```

Respuesta de error:

```json
{
  "error": "rol no encontrado"
}
```

## Validaciones principales

- Los IDs recibidos por ruta deben ser numericos.
- Al crear o actualizar registros relacionados con usuarios, se valida que `id_usuario` exista.
- Al crear o actualizar usuarios, se valida que `id_rol` exista.
- Si un registro no existe, la API responde con `404`.
- Si el JSON enviado no es valido, la API responde con `400`.
- Si hay relaciones en base de datos que impiden eliminar un registro, la API puede responder con `409`.

## Pruebas

Ejecuta las pruebas con:

```bash
go test ./...
```

Actualmente hay pruebas para verificar el registro de rutas principales.

## Notas

- La API escucha por defecto en el puerto `8094`.
- El proyecto usa nombres de tablas con mayusculas, por eso algunas consultas califican el esquema como `"Usuarios"`.
- Las contrasenas se almacenan como hash en el campo `contrasena_hash`; no se deben guardar contrasenas en texto plano.
