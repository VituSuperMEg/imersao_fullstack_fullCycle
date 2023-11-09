package usecase

import (
	"errors"

	"github.com/codeedu/imersao/codepix-go/domain/model"
)

type PixUseCast struct {
	PixKeyRepository model.PixKeyRepositoryInterface
}

func (p *PixUseCast) RegisterKey(key string, kind string, accountId string) (*model.PixKey, error) {
	account, err := p.PixKeyRepository.FindAccount(accountId)

	if err != nil {
		return nil, err
	}

	pixKey, err := model.NewPixKey(kind, account, key)
	if err != nil {
		return nil, err
	}

	p.PixKeyRepository.RegisterKey(pixKey)
	if pixKey.ID == "" {
		return nil, errors.New("unable to create new key is moment")
	}

	return pixKey, nil
}

func (p *PixUseCast) FindKey(key string, kind string) (*model.PixKey, error) {
	pixKey, err := p.PixKeyRepository.FindKeyByKind(key, kind)
	if err != nil {
		return nil, err
	}
	return pixKey, nil
}
