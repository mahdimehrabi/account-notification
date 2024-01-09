package service

import (
	"errors"
	http "github.com/mahdimehrabi/account-notification/account/application/rest/dpi"
	"github.com/mahdimehrabi/account-notification/account/internal/entity"
	accountRepo "github.com/mahdimehrabi/account-notification/account/internal/repository/account"
	"github.com/rs/zerolog/log"
	"time"
)

type AccountService struct {
	accountRepository accountRepo.Account
}

func NewAccountService() *AccountService {
	return &AccountService{
		accountRepository: http.AppDPI.AccountRepo,
	}
}

func (s AccountService) Create(account *entity.Account) error {
	if err := s.accountRepository.Create(account); err != nil {
		log.Error().Msg(err.Error())
		return err
	}

	return nil
}

func (s AccountService) List() ([]entity.Account, error) {
	accounts, err := s.accountRepository.List()
	if err != nil {
		log.Error().Msg(err.Error())
		return accounts, nil
	}
	return accounts, nil
}

func (s AccountService) Send(transaction *entity.Transaction) error {
	if err := s.accountRepository.Send(transaction); err != nil {
		if !errors.Is(err, accountRepo.ErrNotFound) && !errors.Is(err, accountRepo.ErrInsufficientAmount) {
			log.Error().Msg(err.Error())
		}
		return err
	}
	if err := http.AppDPI.NotificationSender.Send(*transaction); err != nil {
		log.Error().Msg(err.Error())
		for {
			go func() {
				for {
					if err := s.accountRepository.SendContemplate(transaction); err != nil {
						log.Error().Msg(err.Error())
						//oh its horrible , trying to give user money back :D
						time.Sleep(60 * time.Second)
						continue
					}
					break
				}
			}()
		}
	}

	return nil
}
