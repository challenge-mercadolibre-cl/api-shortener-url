# Tecnologias.

- [Golang 1.19](https://go.dev/)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/) (sólo para desarrollo local)
- [Redis](https://redis.io/)

---

# Razonamiento del Lenguaje.

## ¿Por qué Golang?.

Al tener muchas tecnologias en el mercado presente y siguiendo el stack técnologico el cual se encuentra hoy la empresa (Java, JavaScript / TypeScript, Python y Golang), se escoge la necesidad de impartir por Golang por el hecho de:

1. Manejar concurrencias de procesos a través de go routines que permitiran tener un mejor procesamiento en menor tiempo.
2. Tener tiempos de respuestas bajo ante la necesidad de altos RPM's.

## ¿Por qué Air en Golang?

Para tener el watch mode al momento de desarrollar, dado que Golang es un lenguaje compilado muchas veces se pierde tiempo en _"bajar"_ la aplicación y volver a subir para que está compile. Al utilizar Air podemos despreocuparnos de este problema.

---

# Razonamiento Arquitectura Software.

## ¿Por qué Arquitectura Hexagonal?

Al poder utilizar arquitectura hexagonal y el desacople que se visualiza en este proyecto, nos permita grandes ventajas tales como:

1. **Testing**: Al tener todo desacoplado, es más fácil el hecho de testear y poder tener la lógica separada entre tantas capas, además.
2. **Escalabilidad**: Al cambiar la fuente de datos (por ej. de Redis a Postgress o Memcached u otra fuente de datos) es mucho más fácil, dado que solo es necesario escribir el repositorio que realizará la conexión hacía dicho servicio.

---

# Razonamiento Infraestructura.

## ¿Por qué Kubernetes?

Para realizar un escalamiento horizontal y vertical de forma más efectiva, permitiendo que en este caso el proveedor de la nube (Google escogido para la solución) sea responsable del manejo y gestión de los recursos del cluster, permitiendo evitar estrés de forma innecesaria por la mantención de los servicios.

## ¿Por qué Google?, ¿la solución es agnostica a la nube?.

La solución fue implementada en Google basicamente por el background técnico del desarrollador, la solución es agnostica al servicio cloud que nos encontremos (AWS, Azure, etc.).

**ACLARACIÓN**: Las herramientas utilizadas para el flujo de analitica / estadisticas si bien se pueden observar que sean de una "herramienta" especifica, estás se pueden cambiar sin problema alguno, dado que en estos procesos son muy "de la mano" con tools donde está el servicio.

---

# Razonamiento Bases de Datos.

## ¿Por qué Redis y no bases de datos "tradicionales" (NoSQL o SQL)

Es necesario tener un servicio de fuente de datos que permita:

1. **Rápida lectura y escritura sobre los datos:** Permitir tener alto tráfico y que responda en forma adecuada en los tiempos correctos.
2. **Bajo costo:** Tener un bajo costo en la solución para que pueda escalar durante el tiempo.
3. **Implementación algoritmo de desalojo:** Si eventualmente se decide ir por un algoritmo de desalojo, será fácil la transferencía de datos y el "cómo" lo podremos manejar entre bases de datos.

Está solución a nivel de código se encuentra en una fase Beta, la solución "real" con escenarios catastróficos se encuentra documentada en el README del proyecto.

### ¿Por qué se utilizo Redis y no Memcached?

Cuando se tiene _multiples nodos / instancias configuradas_ y se utiliza **Memcached** es necesario implementar un algoritmo distribuido para que se realice la sincronización correcta entre los nodos. En cambio, **Redis** lo tiene implementado por **defecto**.

---

# Razonamiento Testing.

## ¿Por qué test container?.

**Testcontainers** nos permita realizar testing de integración directamente, de está forma, levantando contenedores de docker durante mi suite de caso de pruebas puedo revisar que la integración este correcta _(en este caso, fue la integración realizada hacía Redis donde NO debo realizar ningun mock sobre los servicios)_

---

# ¿Como se calculo el Id de las URL's?.

Si bien se verifico distintas opciones para generas las **ids de urls a partir de string**, tales como:

- Base58
- Base62
- UUID (indistinto de la versión)

Se opto por la idea de ir con un **UUID aleatorio** _(a diferencia de que se genere uno a partir de un string, porque este se podría eventualmente "repetir". El string pensado era la combinación de userId:longUrl)_ que solo utilice 8 caracteres, para optar por el buen uso de:

- Bytes persistidos en Redis.
- Acortar url y dar preferencia al contenido del Tweet para compartir mediante twitter.

# ¿Cómo se podrá obtener estadisticas sobre las urls?.

Para obtener estadisticas sobre las URLS utilizaremos Kafka para realizar streaming de eventos, los eventos que vamos a tener en consideración serán:

- **URL_LINK_WAS_SEE**: Cuando una URL fue visitada.
- **URL_LINK_WAS_CREATED**: Cuando una URL fue creada.
- **URL_LINK_WAS_UPDATED**: Cuando una URL fue modificada.
- **URL_LINK_WAS_DISABLED**: Cuando una URL fue deshabilitada.
- **URL_LINK_WAS_ENABLED**: Cuando una URL fue habilitada.

Al tener **Kafka** como nuestro streaming de eventos y message broker, nos permite fácilmente utilizar **Kafka Connect** para enviar los eventos a otra fuente de datos y de esa forma realizar dashboard / procesos de analitica avanzada.

# ¿Que ocurre ante una catastrofe que se pierda el Redis?.

Como tenemos el streaming de eventos ocurriendo en Kafka y tenemos un periodo de una semana (que es configurable para que los mensajes de los topicos no se puedan leer), se pueden realizar consumidores que estos permiten persistir la información a un postgress (está información será más "historica" para tener en caso que ocurra una catastrofe)

# ¿Podemos implementar un algoritmo de desalojo de datos?.

Dado que el requerimiento es que los datos **NUNCA SE ELIMINAN** se opta porque no se puedan desalojar dichos datos y no requiere que se implemente un algoritmo de desalojo en la fuente de datos principal.

Si se requiere realizar un desalojo de datos en la fuente de datos "caché", se sugiere:

- Realizar a partir del timestamp del último acceso (LAST_ACCESS_TIME)

- ¿Por qué?, por que eso informará que ya no tiene tráfico dicha url pasado cierto tiempo. De todas formas, se puede implementar otras formas donde se opte por el timestamp de creación o actualización, eso se debe definir al momento de realizar el flujo. Considerando además si existirá una división entre cuentas "free" y cuentas de pago.
