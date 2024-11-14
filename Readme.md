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

### Tabla `roles`
La Tabla de `roles` almacena que tipo de rol tiene cada usuario almacenado en la base de datos.

| Campo        | Tipo           | Descripcion                                    |
|--------------|----------------|------------------------------------------------|
| `ID`         | `SERIAL`       | Clave primaria, identificador unico de usuario.
|`name`        | `VARCHAR(255)` | Nombre del rol (cliente, manager, admin)                  
|`descripcion` | `TEXT`         | Descripcion de los permisos o privilegios del rol.                           

**Relación:** `usuarios.rol_id` es uan clave foranea que hace referencia a `roles.id`, definiendo el rol de cada usuario. 

### Tabla `promociones`
La tabla `promociones` almacena las promociones disponibles en el sistema. Cada promoción tiene un requisito de nivel minimo y fechas de inicio y fin para controlar su validez. 


### Tabla `consumed_promo`
La tabla `consumed_promo` registra el consumo de promociones por parte de los usuarios. Esto asegura que una pormocion especifica no se pueda volver a utilizar una vez consumida (a menos que la promocion permita varios usos).

**Relación:** `consumed_promo.usuario_id` se refiere a `usuarios.id` y `consumed_promo.promocion_id` se refiere a `promocion.id`

### Tabla `tickets` 
La tabla `tickets` registra las transacciones en las que un usuario ha ganado puntos a través de compras en la tienda. Esta tabla permite mantener un historial de puntos ganados que se pueden auditar en cualquier momento. 

**Relación:** `tickets.usuario_id` se refiere a `usuario.id`.

### Resumen de relaciones entre tablas
1. `usuarios` ↔ `roles`: relacion de muchos a uno (N:1)
   - Un usuario tiene un rol (cliente, manager, admin)
   - `usuarios.rol_id` es una clave foranea de `roles.id`

2. `usuarios` ↔ `promociones` a través de `consumed_promo`: relacion de muchos a muchos (N:N)
   - Un usuario puede consumir varias promociones, y una promocion puede ser consumida por varios usuarios. 
   - La tabla `consumed_promo` actua como tabla de union registrando cada consumo unico. 

3. `usuarios` ↔ `tickets`: relación de uno a muchos (1:N).
   - Un usuario puede tener multiples tickets que registran sus transacciones y los puntos obtenidos. 
   - `tickets.usuarios_id` es una clave foranea a `usuarios.id`. 
  