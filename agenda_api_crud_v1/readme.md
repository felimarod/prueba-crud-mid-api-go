# agenda_api_crud_v1

API Restful para el manejo de los datos de la agenda personal.

## Especificaciones técnicas

- Go 1.18
- Bee 1.12.0

## Variables de entorno
En un archivo `.env` agregar
```bash
export AGENDA_API_V1_PGUSER=
export AGENDA_API_V1_PGPASS=
export AGENDA_API_V1_PGHOST=
export AGENDA_API_V1_PGPORT=
export AGENDA_API_V1_PGDB=
export AGENDA_API_V1_PGSCHEMA=
```
Y antes de ejecutar el programa ejecutar `source .env` para cargar las variables de entorno

```bash
bee run -downdoc=true -gendoc=true
```
