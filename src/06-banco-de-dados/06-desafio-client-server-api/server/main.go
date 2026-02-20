package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type CurrencyRate struct {
	Code       string  `json:"code"`
	CodeIn     string  `json:"codein"`
	Name       string  `json:"name"`
	High       float64 `json:"high,string"`
	Low        float64 `json:"low,string"`
	VarBid     float64 `json:"varBid,string"`
	PctChange  float64 `json:"pctChange,string"`
	Bid        float64 `json:"bid,string"`
	Ask        float64 `json:"ask,string"`
	Timestamp  string  `json:"timestamp"`
	CreateDate string  `json:"create_date"`
}

type CurrencyRateModel struct {
	ID        uuid.UUID `gorm:"type:text;primaryKey"`
	Code      string    `gorm:"size:3;not null"`
	CodeIn    string    `gorm:"size:3;not null"`
	Name      string    `gorm:"size:255;not null"`
	High      float64   `gorm:"type:real;not null"`
	Low       float64   `gorm:"type:real;not null"`
	VarBid    float64   `gorm:"type:real;not null"`
	PctChange float64   `gorm:"type:real;not null"`
	Bid       float64   `gorm:"type:real;not null"`
	Ask       float64   `gorm:"type:real;not null"`
	gorm.Model
}

func (currencyRate *CurrencyRateModel) BeforeCreate(tx *gorm.DB) error {
	currencyRate.ID = uuid.New()
	return nil
}

func main() {
	db, error := gorm.Open(sqlite.Open("currency.db"), &gorm.Config{})
	if error != nil {
		panic(error)
	}
	db.AutoMigrate(&CurrencyRateModel{})

	http.HandleFunc("/cotacao", func(w http.ResponseWriter, r *http.Request) {
		currencyRate, error := getCurrencyRate()
		if error != nil {
			http.Error(w, "Erro ao buscar cotação na API", http.StatusInternalServerError)
			return
		}

		saveCurrencyRate(currencyRate, db)

		lastCurrencyRate, err := getLastCurrencyRate(db)
		if err != nil {
			http.Error(w, "Erro ao buscar cotação no banco", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]float64{"bid": lastCurrencyRate.Bid})
	})
	http.ListenAndServe(":8080", nil)
}

func getCurrencyRate() (CurrencyRate, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	req, error := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if error != nil {
		return CurrencyRate{}, error
	}

	res, error := http.DefaultClient.Do(req)
	if error != nil {
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("Timeout ao buscar cotação na API: contexto excedeu o limite de 200ms")
		}
		return CurrencyRate{}, error
	}
	defer res.Body.Close()

	var rates map[string]CurrencyRate
	error = json.NewDecoder(res.Body).Decode(&rates)
	if error != nil {
		return CurrencyRate{}, error
	}

	return rates["USDBRL"], nil
}

func saveCurrencyRate(currencyRate CurrencyRate, db *gorm.DB) error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel()

	result := db.WithContext(ctx).Create(&CurrencyRateModel{
		Code:      currencyRate.Code,
		CodeIn:    currencyRate.CodeIn,
		Name:      currencyRate.Name,
		High:      currencyRate.High,
		Low:       currencyRate.Low,
		VarBid:    currencyRate.VarBid,
		PctChange: currencyRate.PctChange,
		Bid:       currencyRate.Bid,
		Ask:       currencyRate.Ask,
	})

	if result.Error != nil {
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("Timeout ao salvar cotação no banco: contexto excedeu o limite de 10ms")
		}
		return result.Error
	}

	return nil
}

func getLastCurrencyRate(db *gorm.DB) (CurrencyRateModel, error) {
	var currencyRate CurrencyRateModel
	result := db.Order("created_at desc").First(&currencyRate)
	return currencyRate, result.Error
}
