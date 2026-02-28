package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand/v2"
	"time"
)

type TransferPayload struct {
	ToAccount string
	Amount    float64
}

type contextKey string

const (
	TraceId contextKey = "traceID"
	UserId  contextKey = "userID"
)

func GetTraceID(ctx context.Context) string {

	if x, ok := ctx.Value(TraceId).(string); ok {
		return x
	}

	return ""
}

func GetUserID(ctx context.Context) (int, error) {
	id, ok := ctx.Value(UserId).(int)
	if ok != true {
		return 0, errors.New("Kullanici Bulunamadi")
	}
	return id, nil
}

type MockDb struct {
}

func (d MockDb) SaveTransaction(userID int, amount float64) {
	fmt.Printf("DB: %d ID'li kullanıcı %.2f TL transfer yaptı\n", userID, amount)
}

// Middlewares

// m1
func RequestLoggerMiddleware(ctx context.Context) context.Context {

	newTrace := fmt.Sprintf("%v-%d", time.Now().Second(), rand.IntN(999))
	fmt.Printf("[LOGGER] Middleware Working, new Trace :  %v \n", newTrace)
	newCtx := context.WithValue(ctx, TraceId, newTrace)
	return newCtx
}

// m2
func AuthGuardMiddleware(ctx context.Context) context.Context {
	trace := GetTraceID(ctx)
	newId := rand.IntN(88)
	fmt.Printf("[Auth] Middleware Working, on Trace : %v , User : %v Checking  \n", trace, newId)
	authCtx := context.WithValue(ctx, UserId, newId)
	return authCtx
}

type TransferService struct {
	Db *MockDb
}

// burda bişeyi merak ettim neden constructor yaparken mockDb için pointre kullandık ?
func NewTransferService(bp *MockDb) *TransferService {
	return &TransferService{
		Db: bp,
	}
}

func (s TransferService) ProcessTransfer(ctx context.Context, payload TransferPayload) error {

	traceID := GetTraceID(ctx)
	userID, ok := GetUserID(ctx)
	if ok != nil {
		return fmt.Errorf("Hata : %e\n", ok)
	}

	fmt.Printf("[Transfer Service] User : %v , Basariyla Giris Yapti Trace : %v\n", userID, traceID)
	s.Db.SaveTransaction(userID, payload.Amount)
	return nil
}

func main() {

	payload := TransferPayload{
		ToAccount: "Faruk",
		Amount:    rand.Float64() + 80,
	}

	rootCtx := context.Background()
	traceCtx := RequestLoggerMiddleware(rootCtx)
	authCtx := AuthGuardMiddleware(traceCtx)
	db := MockDb{}
	service := NewTransferService(&db)

	err := service.ProcessTransfer(authCtx, payload)
	if err != nil {
		fmt.Printf("İşlem Başarısız: %v\n", err)
	}
	fmt.Printf("HTTP_STATUS_CODE : 200\n")

}
