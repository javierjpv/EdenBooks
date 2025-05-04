package main

import (
	"log"
	"net/http"
	"os"

	// Importaciones locales con alias para evitar conflictos
	addressHandler "github.com/javierjpv/edenBooks/internal/modules/addresses/adapters/handlers"
	addressRepository "github.com/javierjpv/edenBooks/internal/modules/addresses/adapters/repositories"
	addressRoute "github.com/javierjpv/edenBooks/internal/modules/addresses/adapters/routes"
	addressService "github.com/javierjpv/edenBooks/internal/modules/addresses/adapters/services"
	addressUseCase "github.com/javierjpv/edenBooks/internal/modules/addresses/application/useCases"

	messageHandler "github.com/javierjpv/edenBooks/internal/modules/messages/adapters/handlers"
	messageRepository "github.com/javierjpv/edenBooks/internal/modules/messages/adapters/repositories"
	messageRoute "github.com/javierjpv/edenBooks/internal/modules/messages/adapters/routes"
	messageService "github.com/javierjpv/edenBooks/internal/modules/messages/adapters/services"
	messageUseCase "github.com/javierjpv/edenBooks/internal/modules/messages/application/useCases"

	chatHandler "github.com/javierjpv/edenBooks/internal/modules/chats/adapters/handlers"
	chatRepository "github.com/javierjpv/edenBooks/internal/modules/chats/adapters/repositories"
	chatRoute "github.com/javierjpv/edenBooks/internal/modules/chats/adapters/routes"
	chatService "github.com/javierjpv/edenBooks/internal/modules/chats/adapters/services"
	chatUseCase "github.com/javierjpv/edenBooks/internal/modules/chats/application/useCases"
	"github.com/javierjpv/edenBooks/internal/modules/chats/server"

	"github.com/joho/godotenv"
	"github.com/stripe/stripe-go/v72"

	userHandler "github.com/javierjpv/edenBooks/internal/modules/users/adapters/handlers"
	userRepository "github.com/javierjpv/edenBooks/internal/modules/users/adapters/repositories"
	userRoute "github.com/javierjpv/edenBooks/internal/modules/users/adapters/routes"
	userService "github.com/javierjpv/edenBooks/internal/modules/users/adapters/services"
	userUseCase "github.com/javierjpv/edenBooks/internal/modules/users/application/useCases"

	carrierHandler "github.com/javierjpv/edenBooks/internal/modules/carriers/adapters/handlers"
	carrierRepository "github.com/javierjpv/edenBooks/internal/modules/carriers/adapters/repositories"
	carrierRoute "github.com/javierjpv/edenBooks/internal/modules/carriers/adapters/routes"
	carrierService "github.com/javierjpv/edenBooks/internal/modules/carriers/adapters/services"
	carrierUseCase "github.com/javierjpv/edenBooks/internal/modules/carriers/application/useCases"

	categoryHandler "github.com/javierjpv/edenBooks/internal/modules/categories/adapters/handlers"
	categoryRepository "github.com/javierjpv/edenBooks/internal/modules/categories/adapters/repositories"
	categoryRoute "github.com/javierjpv/edenBooks/internal/modules/categories/adapters/routes"
	categoryService "github.com/javierjpv/edenBooks/internal/modules/categories/adapters/services"
	categoryUseCase "github.com/javierjpv/edenBooks/internal/modules/categories/application/useCases"

	productHandler "github.com/javierjpv/edenBooks/internal/modules/products/adapters/handlers"
	productRepository "github.com/javierjpv/edenBooks/internal/modules/products/adapters/repositories"
	productRoute "github.com/javierjpv/edenBooks/internal/modules/products/adapters/routes"
	productService "github.com/javierjpv/edenBooks/internal/modules/products/adapters/services"
	productUseCase "github.com/javierjpv/edenBooks/internal/modules/products/application/useCases"

	orderHandler "github.com/javierjpv/edenBooks/internal/modules/orders/adapters/handlers"
	orderRepository "github.com/javierjpv/edenBooks/internal/modules/orders/adapters/repositories"
	orderRoute "github.com/javierjpv/edenBooks/internal/modules/orders/adapters/routes"
	orderService "github.com/javierjpv/edenBooks/internal/modules/orders/adapters/services"
	orderUseCase "github.com/javierjpv/edenBooks/internal/modules/orders/application/useCases"

	notificationHandler "github.com/javierjpv/edenBooks/internal/modules/notifications/adapters/handlers"
	notificationRepository "github.com/javierjpv/edenBooks/internal/modules/notifications/adapters/repositories"
	notificationRoute "github.com/javierjpv/edenBooks/internal/modules/notifications/adapters/routes"
	notificationService "github.com/javierjpv/edenBooks/internal/modules/notifications/adapters/services"
	notificationUseCase "github.com/javierjpv/edenBooks/internal/modules/notifications/application/useCases"

	reviewHandler "github.com/javierjpv/edenBooks/internal/modules/reviews/adapters/handlers"
	reviewRepository "github.com/javierjpv/edenBooks/internal/modules/reviews/adapters/repositories"
	reviewRoute "github.com/javierjpv/edenBooks/internal/modules/reviews/adapters/routes"
	reviewService "github.com/javierjpv/edenBooks/internal/modules/reviews/adapters/services"
	reviewUseCase "github.com/javierjpv/edenBooks/internal/modules/reviews/application/useCases"

	transactionHandler "github.com/javierjpv/edenBooks/internal/modules/transactions/adapters/handlers"
	transactionRepository "github.com/javierjpv/edenBooks/internal/modules/transactions/adapters/repositories"
	transactionRoute "github.com/javierjpv/edenBooks/internal/modules/transactions/adapters/routes"
	transactionService "github.com/javierjpv/edenBooks/internal/modules/transactions/adapters/services"
	transactionUseCase "github.com/javierjpv/edenBooks/internal/modules/transactions/application/useCases"

	eventBusService "github.com/javierjpv/edenBooks/internal/shared/infrastructure/services"

	stripeHandler "github.com/javierjpv/edenBooks/internal/modules/stripe/adapters/handlers"
	stripeRoute "github.com/javierjpv/edenBooks/internal/modules/stripe/adapters/routes"

	addressEntities "github.com/javierjpv/edenBooks/internal/modules/addresses/domain/entities"
	carrierEntities "github.com/javierjpv/edenBooks/internal/modules/carriers/domain/entities"
	categoryEntities "github.com/javierjpv/edenBooks/internal/modules/categories/domain/entities"
	chatEntities "github.com/javierjpv/edenBooks/internal/modules/chats/domain/entities"
	messageEntities "github.com/javierjpv/edenBooks/internal/modules/messages/domain/entities"
	notificationEntities "github.com/javierjpv/edenBooks/internal/modules/notifications/domain/entities"
	orderEntities "github.com/javierjpv/edenBooks/internal/modules/orders/domain/entities"
	productEntities "github.com/javierjpv/edenBooks/internal/modules/products/domain/entities"
	reviewEntities "github.com/javierjpv/edenBooks/internal/modules/reviews/domain/entities"
	transactionEntities "github.com/javierjpv/edenBooks/internal/modules/transactions/domain/entities"
	userEntities "github.com/javierjpv/edenBooks/internal/modules/users/domain/entities"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println("No se pudo cargar el archivo .env, asegurarse de que STRIPE_SECRET_KEY esté configurado")
	}

	// Configurar clave de Stripe
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")
	if stripe.Key == "" {
		log.Fatal("STRIPE_SECRET_KEY no está definido en el entorno")
	}

	// Conectar a la base de datos SQLite
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error al conectar a la base de datos:", err)
	}

	// Migrar las estructuras de la base de datos
	err = db.AutoMigrate(
		&userEntities.User{},
		&carrierEntities.Carrier{},
		&chatEntities.Chat{},
		&messageEntities.Message{},
		&addressEntities.Address{},
		&notificationEntities.Notification{},
		&orderEntities.Order{},
		&categoryEntities.Category{},
		&productEntities.Product{},
		&reviewEntities.Review{},
		&transactionEntities.Transaction{},
	)
	if err != nil {
		log.Fatal("Error al migrar las tablas:", err)
	}
	eventBus := eventBusService.NewGoEventBus()

	addressRepository:=addressRepository.NewAddressRepository(db)
	addressService:=addressService.NewAddressService(addressRepository)
	addressUseCase:=addressUseCase.NewAddressUseCase(addressService)
	addressHandler:=addressHandler.NewAddressHandler(*addressUseCase)

	messageRepository:=messageRepository.NewMessageRepository(db)
	messageService:=messageService.NewMessageService(messageRepository)
	MessageUseCase:=messageUseCase.NewMessageUseCase(messageService)
	messageHandler:=messageHandler.NewMessageHandler(*MessageUseCase)
    
	chatRepository:=chatRepository.NewChatRepository(db)
	chatService:=chatService.NewChatService(chatRepository)
	chatUseCase:=chatUseCase.NewChatUseCase(chatService)
	chatHandler:=chatHandler.NewChatHandler(*chatUseCase)


	userRepository:=userRepository.NewUserRepository(db)
	userService:=userService.NewUserService(userRepository)
	userUseCase:=userUseCase.NewUserUseCase(userService)
	userHandler:=userHandler.NewUserHandler(*userUseCase)


	carrierRepository:=carrierRepository.NewCarrierRepository(db)
	carrierService:=carrierService.NewCarrierService(carrierRepository)
	carrierUseCase:=carrierUseCase.NewCarrierUseCase(carrierService)
	carrierHandler:=carrierHandler.NewCarrierHandler(*carrierUseCase)

	categoryRepository:=categoryRepository.NewCategoryRepository(db)
	categoryService:=categoryService.NewCategoryService(categoryRepository)
	categoryUseCase:=categoryUseCase.NewCategoryUseCase(categoryService)
	categoryHandler:=categoryHandler.NewCategoryHandler(*categoryUseCase)


	//Evenbus
	notificationRepository:=notificationRepository.NewNotificationRepository(db)
	notificationService:=notificationService.NewNotificationService(notificationRepository,eventBus)
	notificationService.ListenOrderCreated()
	notificationUseCase:=notificationUseCase.NewNotificationUseCase(notificationService)
	notificationHandler:=notificationHandler.NewNotificationHandler(*notificationUseCase)

	productRepository:=productRepository.NewProductRepository(db)
	productService:=productService.NewProductService(productRepository,eventBus)
	productUseCase:=productUseCase.NewProductUseCase(productService)
	productHandler:=productHandler.NewProductHandler(*productUseCase)

	orderRepository:=orderRepository.NewOrderRepository(db)
	orderService:=orderService.NewOrderService(orderRepository,productService,addressService,carrierService,userService,eventBus)
	orderService.ListenPaymentCreated()
	orderUseCase:=orderUseCase.NewOrderUseCase(orderService)
	orderHandler:=orderHandler.NewOrderHandler(*orderUseCase)

	transactionRepository:=transactionRepository.NewTransactionRepository(db)
	transactionService:=transactionService.NewTransactionService(transactionRepository)
	transactionUseCase:=transactionUseCase.NewTransactionUseCase(transactionService)
	transactionHandler:=transactionHandler.NewTransactionHandler(*transactionUseCase)
	stripeHandler:=stripeHandler.NewStripeHandler(eventBus,orderUseCase,transactionUseCase,addressUseCase)

    //Evenbus



	reviewRepository:=reviewRepository.NewReviewRepository(db)
	reviewService:=reviewService.NewReviewService(reviewRepository)
	reviewUseCase:=reviewUseCase.NewReviewUseCase(reviewService)
	reviewHandler:=reviewHandler.NewReviewHandler(*reviewUseCase)

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
	}))

	// e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
	// 	TokenLookup: "header:X-XSRF-TOKEN",
	//   }))

	//   e.Use(middleware.SecureWithConfig(middleware.SecureConfig{
	// 	XSSProtection:         "",
	// 	ContentTypeNosniff:    "",
	// 	XFrameOptions:         "",
	// 	HSTSMaxAge:            3600,
	// 	ContentSecurityPolicy: "default-src 'self'",
	// }))


	addressRoute.RegisterAddressRoutes(e,addressHandler)

    messageRoute.RegisterMessageRoutes(e,messageHandler)

	chatRoute.RegisterChatRoutes(e,chatHandler)
	userRoute.RegisterUserRoutes(e,userHandler)
	carrierRoute.RegisterCarrierRoutes(e,carrierHandler)
	categoryRoute.RegisterCategoryRoutes(e,categoryHandler)
	productRoute.RegisterProductRoutes(e,productHandler)
	reviewRoute.RegisterReviewRoutes(e,reviewHandler)
	orderRoute.RegisterOrderRoutes(e,orderHandler)
	notificationRoute.RegisterNotificationRoutes(e,notificationHandler)
	transactionRoute.RegisterTransactionRoutes(e,transactionHandler)
	stripeRoute.RegisterStripeRoutes(e,stripeHandler)


	//inicio websocket
    go func(){
		server.Run(MessageUseCase)
	}()
	
	//el servidor de Echo y el servidor de Gorilla websocket 
	// estaran disponibles y usaran la misma terminal para logs


	//Inicio Servidor Echo
	e.Logger.Fatal(e.Start(":6969"))
}
