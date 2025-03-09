# Backend Architecture

El backend de EdenBooks está construido utilizando las siguientes arquitecturas y principios:

## Arquitectura Hexagonal

La arquitectura hexagonal, también conocida como Ports and Adapters, separa la lógica del negocio de las implementaciones técnicas, lo que facilita la prueba y el mantenimiento del código.

## Domain-Driven Design (DDD)

DDD se centra en el diseño del software basado en el dominio del problema, organizando el código en torno a conceptos del negocio y no en torno a detalles técnicos.

## Clean Architecture

Clean Architecture promueve la separación de preocupaciones, permitiendo que el código sea más flexible, escalable y fácil de mantener.

## Event-Driven

El sistema está diseñado para reaccionar a eventos, lo que permite una comunicación más eficiente y desacoplada entre los diferentes componentes del sistema.

## Estructura del Proyecto

- **Domain**: Contiene los modelos y las interfaces de repositorios y servicios.
  - **entities/**: Define las entidades del dominio.
  - **repositories/**: Define las interfaces de los repositorios.
  - **services/**: Define las interfaces de los servicios del dominio.
- **Application**: Contiene casos de uso y dtos.
  - **usecases/**: Contiene los casos de uso que implementan la lógica de la aplicación.
  - **dtos/**: Contiene los objetos de transferencia de datos.
- **Adapters**: Contiene las implementaciones de las interfaces de la carpeta Domain.
  - **repositories/**: Implementaciones de las interfaces de repositorios.
  - **services/**: Implementaciones de las interfaces de servicios.
  - **handlers/**: Controladores y manejadores de rutas.
  - **routes/**: Rutas.

Para más detalles sobre la implementación específica, revisa los archivos en el directorio correspondiente.

## Ejecución del Backend

Para ejecutar el backend, sigue las instrucciones en el [README principal](../README.md).
