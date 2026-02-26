package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

// Services

type Log struct {
	logMessage string
	logSource  string
	logDate    time.Time
}

func CheckInventory(ctx context.Context, wg *sync.WaitGroup, errCh chan<- error, logCh chan<- Log) {
	defer wg.Done()

	fmt.Println("Envanter Servisi : Envanter Kontrol Ediliyor")
	select {
	case <-time.After(1 * time.Second):
		{
			logCh <- Log{
				logMessage: "yetersiz stock",
				logSource:  "Check Inventory",
				logDate:    time.Now(),
			}
			errCh <- errors.New("envanter Servisi : Stok Yetersiz")
		}

	case <-ctx.Done(): // Eğer üst katmandan (orkestratör) iptal sinyali gelirse kapanacak
		{
			logCh <- Log{
				logMessage: "İptal Sinyali Geldi",
				logSource:  "Check Inventory",
				logDate:    time.Now(),
			}
			fmt.Println("Envanter Servisi : İptal Sinyali Geldi BABA TAHLİYE")
		}

	}
}
func VerifyFraud(ctx context.Context, wg *sync.WaitGroup, logCh chan<- Log) {
	defer wg.Done()

	fmt.Println("Doğrulama Servisi : Kullanıcı Bilgileri Kontrol Ediliyor")
	select {
	case <-ctx.Done(): // Eğer üst katmandan (orkestratör) iptal sinyali gelirse kapanacak
		fmt.Println("Doğrulama Servisi : İptal Sinyali Geldi BABA TAHLİYE")
		logCh <- Log{
			logMessage: "İptal Sinyali Geldi",
			logSource:  "Dogrulama Servisi",
			logDate:    time.Now(),
		}

	case <-time.After(2 * time.Second):
		{
			fmt.Println("Doğrulama Servisi : Kullanici Bilgileri Dogrulandı")
			logCh <- Log{
				logMessage: "Dogrulama Basarili",
				logSource:  "Dogrulama Servisi",
				logDate:    time.Now(),
			}
		}
	}
}
func ProcessPayment(ctx context.Context, wg *sync.WaitGroup, logCh chan<- Log) {
	defer wg.Done()
	fmt.Println("Ödeme Servisi : Ödeme Bilgileri Kontrol Ediliyor")
	select {
	case <-ctx.Done(): // Eğer üst katmandan (orkestratör) iptal sinyali gelirse kapanacak
		logCh <- Log{
			logMessage: "İptal Sinyali Geldi",
			logSource:  "Ödeme Servisi",
			logDate:    time.Now(),
		}
		fmt.Println("Ödeme Servisi : İptal Sinyali Geldi BABA TAHLİYE")

	case <-time.After(3 * time.Second):
		fmt.Println("Ödeme Servisi : Ödeme Basarili")
		logCh <- Log{
			logMessage: "Ödeme Basarili",
			logSource:  "Ödeme Servisi",
			logDate:    time.Now(),
		}
	}
}

// Orchestrator
func PlaceOrder() error {
	logChannel := make(chan Log)
	errChannel := make(chan error, 3)
	timeoutCTX, cancel := context.WithTimeout(context.Background(), 88*time.Second)
	defer cancel()

	withCancelCTX, canc := context.WithCancel(timeoutCTX)
	defer canc()

	// Yaziyorum ama kafama yakılan bir sorun su neden timeoutContext olusturduk sonrada onnun uzerinden
	//withCancel olusturduk ne kadar gerekli anlamadım
	var wg sync.WaitGroup
	wg.Add(3)

	// LOGGER (consumer)
	go func() {
		for log := range logChannel {
			fmt.Printf("[LOG] %s | %s | %s\n",
				log.logDate.Format(time.RFC3339),
				log.logSource,
				log.logMessage,
			)
		}
	}()

	go CheckInventory(withCancelCTX, &wg, errChannel, logChannel)
	go VerifyFraud(withCancelCTX, &wg, logChannel)
	go ProcessPayment(withCancelCTX, &wg, logChannel)

	doneCh := make(chan struct{})
	go func() {
		wg.Wait()
		close(doneCh)
		close(logChannel)
	}()
	// Gözlemci Goroutine (Mutlu senaryo için)
	select {

	case <-doneCh:
		{
			return nil // Her şey kusursuz bitti.
		}

	case err := <-errChannel:
		{
			canc()
			return fmt.Errorf("Sipariş iptal edildi, sebep: %v", err)
		}
	case <-timeoutCTX.Done():
		return fmt.Errorf("Sistem Zaman Asimina Ugradi")

	}

}

func main() {
	siparisErr := PlaceOrder()

	if siparisErr != nil {
		fmt.Printf("BAŞARISIZ: %v\n", siparisErr)
	} else {
		fmt.Println("BASARİLİ")
	}

	time.Sleep(100 * time.Millisecond)

}
