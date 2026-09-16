package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/javierjpv/edenBooks/internal/modules/users/application/dto"
	usecases "github.com/javierjpv/edenBooks/internal/modules/users/application/useCases"
	"github.com/labstack/echo/v4"
)

// Clave secreta (debería ir en variables de entorno)
var secretKey = []byte("supersecret")

type UserHandler struct {
	useCase usecases.UserUseCase
}

func NewUserHandler(useCase usecases.UserUseCase) *UserHandler {
	return &UserHandler{useCase: useCase}
}

func (h *UserHandler) Login(c echo.Context) error {
	var user dto.UserRequest
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al recoger los datos"})
	}
	registredUser, err := h.useCase.Login(user.Email, user.Password)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Credenciales incorrectas"})
	}
	// Crear token JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"ID":  registredUser.ID,
		"exp": time.Now().Add(time.Hour * 2).Unix(),
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al generar el token"})
	}
    registredUser.Token=tokenString
	return c.JSON(http.StatusOK, registredUser)
}

func (h *UserHandler) Register(c echo.Context) error {
	var user dto.UserRequest
	if err := c.Bind(&user); err != nil {
		// c.Logger().Error("Error al recoger los datos:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al recoger los datos"})
	}

	if err := h.useCase.Register(user); err != nil {
		// c.Logger().Error("Error al registrar usuario:", err)
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Error al crear el usuario"})
	}
	registredUser, err := h.useCase.Login(user.Email, user.Password)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Credenciales incorrectas"})
	}
	// Crear token JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"ID":  registredUser.ID,
		"exp": time.Now().Add(time.Hour * 2).Unix(),
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al generar el token"})
	}
	registredUser.Token=tokenString
	return c.JSON(http.StatusOK, registredUser)
}

func (h *UserHandler) GetUserByID(c echo.Context) error {
	userID := c.Param("id")
	id, err := strconv.Atoi(userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}
	userResponse, err := h.useCase.GetUserByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Error al obtener el user"})
	}
	return c.JSON(http.StatusOK, userResponse)
}

func (h *UserHandler) UpdateUser(c echo.Context) error {
	var user dto.UserRequest
	userID := c.Param("id")
	id, err := strconv.Atoi(userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "error al recoger los datos"})
	}

	if err := h.useCase.UpdateUser(uint(id), user); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al actualizar el user"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "user actualizado correctamente"})
	// return c.NoContent(http.StatusNoContent) convencion rest, aunque si quieres enviar un mensaje esta bien

}
