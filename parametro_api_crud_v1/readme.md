# agenda_api_crud_v1

API Restful para el manejo de los datos de la agenda personal.

## Especificaciones técnicas

- Go 1.18
- Bee 1.12.0

## Variables de entorno
En un archivo `.env` agregar
```bash
export PARAMETROS_API_V1_PGUSER=postgres
export PARAMETROS_API_V1_PGPASS=avion123
export PARAMETROS_API_V1_PGHOST=localhost
export PARAMETROS_API_V1_PGPORT=5432
export PARAMETROS_API_V1_PGDB=core
export PARAMETROS_API_V1_PGSCHEMA=parametros
```
Y antes de ejecutar el programa ejecutar `source .env` para cargar las variables de entorno

```bash
bee run -downdoc=true -gendoc=true
```
