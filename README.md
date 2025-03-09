# EdenBooks

App where you can buy and sell second-hand books.

## Overview

EdenBooks es una plataforma diseñada para facilitar la compra y venta de libros de segunda mano. Nuestro objetivo es proporcionar un mercado amigable donde los usuarios puedan intercambiar libros usados, haciendo más fácil encontrar libros asequibles y promoviendo la sostenibilidad.

## Features

- **Comprar y Vender Libros**: Los usuarios pueden listar sus libros usados para la venta y buscar libros disponibles para comprar.
- **Cuentas de Usuario**: Crea y gestiona cuentas de usuario para una experiencia personalizada.
- **Buscar y Filtrar**: Busca fácilmente libros por título, autor o categoría y filtra los resultados para encontrar exactamente lo que necesitas.
- **Transacciones Seguras**: Asegura transacciones seguras entre compradores y vendedores mediante Stripe.
- **Chat**: Los usuarios pueden comunicarse entre sí antes de comprar un libro mediante chats.
- **Whislist**: Cada usuario puede guardar sus libros favoritos en su lista de deseos

## Tecnologías Utilizadas

- **Backend**: Go (55.4%)
- **Frontend**: TypeScript y React (44.3%)
- **Herramienta de Construcción**: Vite
- **Otros**: (0.3%)

## Getting Started

### Prerequisitos

- [Go](https://golang.org/doc/install)
- [Node.js](https://nodejs.org/en/download/)
- [Yarn](https://classic.yarnpkg.com/en/docs/install)

### Instalación

1. Clona el repositorio:
   git clone https://github.com/javierjpv/EdenBooks.git
   cd EdenBooks

2. Instala las dependencias del backend:
go mod download

3. Instala las dependencias del frontend:
cd frontend/edenBooks
yarn

### Ejecución de la Aplicación

1. Inicia el servidor backend:
go run main.go

2. Inicia el servidor de desarrollo del frontend:
yarn dev

3. Abre tu navegador y navega a http://localhost:5173 para acceder a la aplicación.

## Contribuyendo
¡Acepto contribuciones! Por favor, lee las Guías de Contribución para más detalles.

Licencia
Este proyecto está licenciado bajo la Licencia MIT.

## Agradecimientos
Un agradecimiento especial a la comunidad de código abierto.