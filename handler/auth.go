package handler

import (
    "net/http"

    "github.com/labstack/echo"
    "github.com/WillyWilsen/Deall-Jobs_Technical-Test.git/model"
    "github.com/WillyWilsen/Deall-Jobs_Technical-Test.git/repository"
)

type AuthHandler struct {
    authRepository repository.AuthRepository
}

func NewAuthHandler(authRepository repository.AuthRepository) *AuthHandler {
    return &AuthHandler{
        authRepository: authRepository,
    }
}

func (h *AuthHandler) Register(c echo.Context) error {
    var user model.User

    // Bind payload
    if err := c.Bind(&user); err != nil || user.Name == "" {
        return c.JSON(http.StatusBadRequest, echo.Map{
            "code": http.StatusBadRequest,
            "status": "error",
            "message": "Invalid request payload",
        })
    }

    // Validate email
    if !model.ValidateEmail(user.Email) {
        return c.JSON(http.StatusBadRequest, echo.Map{
            "code": http.StatusBadRequest,
            "status": "error",
            "message": "Invalid email address",
        })
    }

    // Exist email
    exist_user, _ := h.authRepository.GetByEmail(user.Email)
    if exist_user != nil {
        return c.JSON(http.StatusConflict, echo.Map{
            "code": http.StatusConflict,
            "status": "error",
            "message": "Email already exists",
        })
    }

    // Hash password
    if err := user.HashPassword(user.Password); err != nil {
        return c.JSON(http.StatusInternalServerError, echo.Map{
            "code": http.StatusInternalServerError,
            "status": "error",
            "message": "Failed to hash password",
        })
    }

    // Register user
    if err := h.authRepository.Register(&user); err != nil {
        return c.JSON(http.StatusInternalServerError, echo.Map{
            "code": http.StatusInternalServerError,
            "status": "error",
            "message": "Failed to register user",
        })
    }

    // Success
    return c.JSON(http.StatusOK, echo.Map{
        "code": http.StatusOK,
        "status": "success",
        "message": "User registered successfully",
        "data": echo.Map{
            "name": user.Name,
            "email": user.Email,
        },
    })
}

func (h *AuthHandler) Login(c echo.Context) error {
    var user model.User

    // Bind payload
    if err := c.Bind(&user); err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{
            "code": http.StatusBadRequest,
            "status": "error",
            "message": "Invalid request payload",
        })
    }

    // Exist email
    exist_user, _ := h.authRepository.GetByEmail(user.Email)
    if exist_user == nil {
        return c.JSON(http.StatusUnauthorized, echo.Map{
            "code": http.StatusUnauthorized,
            "status": "error",
            "message": "Invalid credentials",
        })
    }

    // Verify password
    if err := exist_user.VerifyPassword(user.Password); err != nil {
        return c.JSON(http.StatusUnauthorized, echo.Map{
            "code": http.StatusUnauthorized,
            "status": "error",
            "message": "Invalid credentials",
        })
    }

    // Success
    return c.JSON(http.StatusOK, echo.Map{
        "code": http.StatusOK,
        "status": "success",
        "message": "Login successful",
        "data": echo.Map{
            "name": exist_user.Name,
            "email": exist_user.Email,
        },
    })
}