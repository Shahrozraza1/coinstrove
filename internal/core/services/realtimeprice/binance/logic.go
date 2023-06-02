package binance

import (
	"coinscience/consts"
	"coinscience/internal/core/ports"
	"fmt"
)

type newBinanceService struct {
	priceRepo ports.PriceRepository
	data      []map[string]interface{}
}

func NewBinanceService(priceRepo ports.PriceRepository) ports.PriceService {
	return &newBinanceService{
		priceRepo: priceRepo,
	}
}

func (repo *newBinanceService) GetThePrice() {
	repo.data = repo.priceRepo.Get(consts.BINANCE)
	fmt.Println(repo.data)
}

func (repo *newBinanceService) Publish() {

}
