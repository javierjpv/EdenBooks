
### README del Backend (backend/README.md)

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

- **Domain**: Contiene los modelos y lógica del negocio.
- **Application**: Contiene la lógica de la aplicación, como casos de uso y servicios.
- **Infrastructure**: Contiene las implementaciones técnicas, como acceso a la base de datos y servicios externos.
- **Interfaces**: Contiene los adaptadores para la interacción con el mundo exterior, como controladores HTTP y suscripción a eventos.

Para más detalles sobre la implementación específica, revisa los archivos en el directorio correspondiente.

## Ejecución del Backend

Para ejecutar el backend, sigue las instrucciones en el [README principal](../README.md).
