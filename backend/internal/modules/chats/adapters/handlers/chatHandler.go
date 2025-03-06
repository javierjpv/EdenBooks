package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/javierjpv/edenBooks/internal/modules/chats/application/dto"
	usecases "github.com/javierjpv/edenBooks/internal/modules/chats/application/useCases"
	"github.com/javierjpv/edenBooks/internal/shared/auth"
	"github.com/labstack/echo/v4"
)

type ChatHandler struct{
	useCase usecases.ChatUseCase
}

func NewChatHandler(useCase usecases.ChatUseCase)*ChatHandler{
	return &ChatHandler{useCase: useCase}
}

func (h *ChatHandler)CreateChat(c echo.Context)error{
	var chat dto.ChatDTO
	
	if err:=c.Bind(&chat);err!=nil{
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "error al recoger los datos"})
	}

	
	if _,err:=h.useCase.CreateChat(chat);err!=nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al crear la chat"})
	}
	return c.JSON(http.StatusCreated,map[string]string{"chat": "chat creada correctamente"})
}

func (h *ChatHandler)CreateChatWithUsers(c echo.Context)error{
	var req struct {
		// Chat dto.ChatDTO `json:"chat"`
		UserIDs []uint       `json:"userIDs"`
	}
	// var chat dto.ChatDTO
	
	if err:=c.Bind(&req);err!=nil{
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "error al recoger los datos"})
	}

	chat,err:=h.useCase.CreateChatWithUsers(req.UserIDs);
	if err!=nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al crear la chat"})
	}
	return c.JSON(http.StatusCreated, chat)

}
func (h *ChatHandler)UpdateChat(c echo.Context)error{
	var chat dto.ChatDTO
	chatID:=c.Param("id")
    id, err := strconv.Atoi(chatID)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
    }
	if err:=c.Bind(&chat);err!=nil{
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "error al recoger los datos"})
	}

	
	if err:=h.useCase.UpdateChat(uint(id),chat);err!=nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al actualizar la chat"})
	}
	return c.JSON(http.StatusOK,map[string]string{"chat": "chat actualizada correctamente"})
	// return c.NoContent(http.StatusNoContent) convencion rest, aunque si quieres enviar un mensaje esta bien

}


func (h *ChatHandler)DeleteChat(c echo.Context)error{
	chatID:=c.Param("id")
    id, err := strconv.Atoi(chatID)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
    }
	
	if err:=h.useCase.DeleteChat(uint(id));err!=nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al borrar la chat"})
	}
	return c.JSON(http.StatusOK,map[string]string{"chat": "chat borrada correctamente"})
	// return c.NoContent(http.StatusNoContent) convencion rest, aunque si quieres enviar un mensaje esta bien
}


func (h *ChatHandler)GetChatByID(c echo.Context)error{
	chatID:=c.Param("id")
    id, err := strconv.Atoi(chatID)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
    }
	chat,err:=h.useCase.GetChatByID(uint(id));
	if err!=nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Error al obtener la chat"})
	}
	return c.JSON(http.StatusOK, chat)
}

func (h *ChatHandler) GetFilteredChats(c echo.Context) error {
	// Extraer todos los filtros dinámicos de la URL
	// filters := make(map[string]string)
	// for key, values := range c.QueryParams() {
	// 	if len(values) > 0 {
	// 		filters[key] = values[0]
	// 	}
	// }
	fmt.Println("Procesando solicitud...");
	// Obtener el token del encabezado Authorization
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Token no proporcionado"})
	}


	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	
	userID, err :=auth.ExtractUserIDFromToken(tokenString)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Token inválido: " + err.Error()})
	}

	fmt.Println("user_id",userID);
	// Extraer filtros dinámicos de la URL
	filters := make(map[string]string)
	for key, values := range c.QueryParams() {
		if len(values) > 0 {
			filters[key] = values[0]
		}
	}
	// Agregar el userID como un filtro obligatorio
	filters["user_id"] = fmt.Sprintf("%d", userID)
	chats, err := h.useCase.GetFilteredChats(filters)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, chats)
}

