package main

import (
	"fmt"

	"github.com/piquette/finance-go/chart"
	"github.com/piquette/finance-go/datetime"
	"github.com/piquette/finance-go/equity"
	"github.com/piquette/finance-go/quote"
)

// Prueba funcional para obtener información del ticker NVDA (NVIDIA)
func main() {
	ticker := "NVDA"

	fmt.Printf("=== Obtener información de %s (NVIDIA) ===\n\n", ticker)

	// 1. Obtener quote básico
	fmt.Println("1. Quote Básico:")
	fmt.Println("====================")
	q, err := quote.Get(ticker)
	if err != nil {
		fmt.Printf("Error obteniendo quote: %v\n", err)
	} else {
		fmt.Printf("Símbolo: %s\n", q.Symbol)
		fmt.Printf("Precio actual: %v\n", q.RegularMarketPrice)
		fmt.Printf("Cambio: %v\n", q.RegularMarketChange)
		fmt.Printf("Porcentaje de cambio: %v\n", q.RegularMarketChangePercent)
		fmt.Printf("Volumen: %v\n", q.RegularMarketVolume)
		fmt.Printf("Estado del mercado: %s\n", q.MarketState)
		fmt.Printf("Hora del mercado: %v\n", q.RegularMarketTime)
		fmt.Println()
	}

	// 2. Obtener información de equity (más detallada)
	fmt.Println("2. Equity Detallado:")
	fmt.Println("=====================")
	e, err := equity.Get(ticker)
	if err != nil {
		fmt.Printf("Error obteniendo equity: %v\n", err)
	} else {
		fmt.Printf("Símbolo: %s\n", e.Symbol)
		fmt.Printf("Nombre: %s\n", e.ShortName)
		fmt.Printf("Precio actual: %v\n", e.RegularMarketPrice)
		fmt.Printf("Precio anterior de cierre: %v\n", e.RegularMarketPreviousClose)
		fmt.Printf("Apertura: %v\n", e.RegularMarketOpen)
		fmt.Printf("Máximo del día: %v\n", e.RegularMarketDayHigh)
		fmt.Printf("Mínimo del día: %v\n", e.RegularMarketDayLow)
		fmt.Printf("Bid: %v\n", e.Bid)
		fmt.Printf("Ask: %v\n", e.Ask)
		fmt.Printf("Capitalización de mercado: %v\n", e.MarketCap)
		fmt.Printf("P/E Ratio: %v\n", e.TrailingPE)
		fmt.Printf("EPS: %v\n", e.EpsTrailingTwelveMonths)
		fmt.Println()
	}

	// 3. Obtener datos históricos (chart)
	fmt.Println("3. Datos Históricos (últimos 5 días):")
	fmt.Println("=========================================")
	params := &chart.Params{
		Symbol:   ticker,
		Interval: datetime.OneDay,
	}
	iter := chart.Get(params)

	count := 0
	for iter.Next() && count < 5 {
		b := iter.Bar()
		fmt.Printf("Fecha: %s\n", datetime.FromUnix(b.Timestamp))
		fmt.Printf("  Apertura: %s\n", b.Open)
		fmt.Printf("  Máximo: %s\n", b.High)
		fmt.Printf("  Mínimo: %s\n", b.Low)
		fmt.Printf("  Cierre: %s\n", b.Close)
		fmt.Printf("  Volumen: %v\n", b.Volume)
		fmt.Printf("  Ajustado: %s\n", b.AdjClose)
		fmt.Println()
		count++
	}

	if iter.Err() != nil {
		fmt.Printf("Error iterando datos históricos: %v\n", iter.Err())
	}

	fmt.Println("\n=== Prueba completada para NVDA ===")
}
