package handlers

import "github.com/Yaduvamsikrishna123/Sweet-Shop-Management-System/internal/repository"

type Handler struct {
	SweetRepo       repository.SweetRepository
	UserRepo        repository.UserRepository
	TransactionRepo repository.TransactionRepository
}

func NewHandler(sweetRepo repository.SweetRepository, userRepo repository.UserRepository, transactionRepo repository.TransactionRepository) *Handler {
	return &Handler{
		SweetRepo:       sweetRepo,
		UserRepo:        userRepo,
		TransactionRepo: transactionRepo,
	}
}
