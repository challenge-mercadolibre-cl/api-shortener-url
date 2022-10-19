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
2. Tener tiempos de respuestas bajo ante la necesidad de altos TPM y TPS eventualmente.

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

La solución fue implementada en Google basicamente por el background técnico del desarrollador, la solución es agnostica al servicio cloud que nos encontremos (AWS, Azure, etc.)

---

# Razonamiento Bases de Datos.

## ¿Por qué Redis y no bases de datos "tradicionales" (NoSQL o SQL)

Es necesario tener un servicio de fuente de datos que permita:

1. Rápida lectura y escritura sobre los datos: Permitir tener alto tráfico y que responda en forma adecuada en los tiempos correctos.
2. Bajo costo: Tener un bajo costo en la solución para que pueda escalar durante el tiempo.

### ¿Por qué se utilizo Redis y no Memcached?

Cuando se tiene _multiples nodos / instancias configuradas_ y se utiliza **Memcached** es necesario implementar un algoritmo distribuido para que se realice la sincronización correcta entre los nodos. En cambio, **Redis** lo tiene implementado por **defecto**.

---

# Razonamiento Testing.

## ¿Por qué test container?.

**Testcontainers** nos permita realizar testing de integración directamente, de está forma, levantando contenedores de docker durante mi suite de caso de pruebas puedo revisar que la integración este correcta _(en este caso, fue la integración realizada hacía Redis donde NO debo realizar ningun mock sobre los servicios)_
