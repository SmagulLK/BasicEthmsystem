package service

import "TestProjectEthereum/internal/repository"

type OperationService struct {
	repo repository.Operation
}

func NewOperationService(repo repository.Operation) *OperationService {
	return &OperationService{repo: repo}
}
func (OperationServ *OperationService) UpdateBalance(balance int32) error {
	return OperationServ.repo.BalanceUpdate(balance)
}
