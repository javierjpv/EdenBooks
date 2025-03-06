package auth

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func ExtractUserIDFromToken(tokenString string) (uint, error) {
	// Usar la misma clave secreta que se usa para firmar
	// fmt.Println("un intento:", 1)
	var secretKey = []byte("supersecret")

	// Parsear el token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verificar que el método de firma sea el correcto
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("método de firma inesperado: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return 0, fmt.Errorf("error al validar el token: %v", err)
	}

	// Verificar que el token sea válido
	if !token.Valid {
		return 0, errors.New("token inválido")
	}

	// Extraer los claims (información) del token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("no se pudieron extraer los claims del token")
	}

	// Extraer el ID del usuario
	userIDFloat, ok := claims["ID"].(float64)
	if !ok {
		return 0, errors.New("ID de usuario no encontrado en el token")
	}

	// Convertir a uint
	userID := uint(userIDFloat)

	return userID, nil
}