# Aplicación de fidelización de clientes

Este proyecto es un backend en GO diseñado para gesionar una aplicación de fidelizacion de clientes mediante acumulación de puntos. Los clientes pueden acumular puntos a través de transaccione, consumir pormociones a las que tienen acceso según su nivel y los managers pueden leer codigos QR para registrar las transacciones.

## Tecnologias y dependencias
**Golang:** Lenguaje principal para el desarrollo del backend.
**Fiber:** Framework web para manejar las rutas y peticiones HTTP.
**GORM:** ORM para interactuar con la base de datos.
**PostgreSQL:** Base de datos para almacenar los datos de usuarios, pormociones...
**Bcrypt:** Libreria para encriptar contraseñas.
**JWT:** Metodo de autenticacion y autorizacion de usuarios y permisos.


## Esquema de la Base de Datos 

### Tabla `usuarios`
La tabla `usuarios` almacena la informacion de todos los usuarios, incluidos clientes, managers y administradores. Cada usuario tiene un nivel de puntos que define el acceso a las promociones, así como un rol que determina sus permisos en la aplicación.

| Campo      | Tipo          | Descripcion                                    |
|------------|---------------|------------------------------------------------|
|`ID`        | `UUID`        | Clave primaria, identificador unico de usuario.
|`name`      | `VARCHAR(255` | Nombre del usuario.                            
|`gender`    | `VARCHAR(255` | Genero del usuario.                            
|`birth_date`| `VARCHAR(255` | Fecha de nacimiento del usuario.               
|`email`     | `VARCHAR(255` | Correo electronico del usuario, unico.         
|`password`  | `VARCHAR(255` | Contraseña encriptada del usuario.                        
|`nivel`     | `INTEGER`     | Nivel del usuario para determinar accesos.                
|`puntos`    | `INTEGER`     | Total de puntos acumulados del usuario.                   
|`rol_id`    | `INTEGER`     | CLave foranea a `roles(id)`, determina el rol del usuario. 
|`created_at`| `TIMESTAMP`   | Fecha de creacion del usuario.
|`updated_at`| `TIMESTAMP`   | Fecha de la ultima actualizacion del usuario.

### Tabla `promociones`
La tabla `promociones` almacena las promociones disponibles en el sistema. Cada promoción tiene un requisito de nivel minimo y fechas de inicio y fin para controlar su validez. 

| Campo          | Tipo           | Descripcion                                    |
|----------------|----------------|------------------------------------------------|
| `id`           | `UUID`         | Clave primaria, identificador unico de promocion.
| `nombre`       | `VARCHAR(255)` | Nombre de la promocion.
| `descripcion`  | `TEXT`         | Descripcion de la promocion.
| `nivel_min`    | `INTEGER`      | Nivel minimo necesario para acceder a la promocion.
| `fecha_inicio` | `DATE`         | Fecha de inicio de la promocion.
| `fecha_fin`    | `DATE`         | Fecha opcional de fin de la promocion (nulo si no tiene fin).



### Tabla `consumed_promo`
La tabla `consumed_promo` registra el consumo de promociones por parte de los usuarios. Esto asegura que una pormocion especifica no se pueda volver a utilizar una vez consumida (a menos que la promocion permita varios usos).

| Campo          | Tipo           | Descripcion                                    |
|----------------|----------------|------------------------------------------------|
| `id`           | `SERIAL`       | Clave primaria de promociones consumidas
| `usuario_id`   | `UUID`         | Clave foranea a `usuarios(id)`, identifica al usuario
| `promocion_id` | `UUID`         | Clave foranea a `promocion(id)`, identifica la promocion
| `fecha_consumo`| `TIMESTAMP`    | Fecha en la que se consumio la promocion.

**Relación:** `consumed_promo.usuario_id` se refiere a `usuarios.id` y `consumed_promo.promocion_id` se refiere a `promocion.id`

### Tabla `tickets` 
La tabla `tickets` registra las transacciones en las que un usuario ha ganado puntos a través de compras en la tienda. Esta tabla permite mantener un historial de puntos ganados que se pueden auditar en cualquier momento. 

| Campo          | Tipo           | Descripcion                                    |
|----------------|----------------|------------------------------------------------|
| `id`           | `SERIAL`       | Clave primaria de puntos obtenidos.
| `usuario_id`   | `UUID`         | Clave foranea a `usuarios(id)`, identifica al usuario.
| `tienda`       | `UUID`         | Clave foranea a `tiendas(id)`, identifica a la tienda.
| `puntos`       | `INTEGER`      | Numero de puntos acumulados en esa compra.
| `fecha`        | `TIMESTAP`     | Fecha en la que se registro la compra.

**Relación:** `tickets.usuario_id` se refiere a `usuario.id`.

  

## Licencia

Este proyecto está licenciado bajo la licencia **Creative Commons Attribution-NonCommercial 4.0 International (CC BY-NC 4.0)**. Puedes ver los términos completos de la licencia en el archivo `LICENSE.md` o visitar [https://creativecommons.org/licenses/by-nc/4.0/](https://creativecommons.org/licenses/by-nc/4.0/).

Esto significa que eres libre de:

- Compartir: Copiar y redistribuir el material en cualquier medio o formato.
- Adaptar: Remezclar, transformar y construir a partir del material.

**Bajo las siguientes condiciones**:

- **Atribución**: Debes dar el crédito adecuado, proporcionar un enlace a la licencia, e indicar si se realizaron cambios. Puedes hacerlo de cualquier manera razonable, pero no de manera que sugiera que el licenciante te respalda o apoya de forma especial.
- **No Comercial**: No puedes utilizar el material para fines comerciales.
