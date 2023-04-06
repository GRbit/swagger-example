package codegen_test

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/grbit/swagger-example/internal/codegen"
	"github.com/grbit/swagger-example/pkg/codegen/client"
	"github.com/grbit/swagger-example/pkg/codegen/models"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/samber/lo"
)

const (
	addr = "localhost:8080"
	host = "http://" + addr + "/api/v1"
)

var (
	startServerOnce sync.Once
	codegenClient   *oapi_client.Client
)

func setup() {
	lg := log.Logger.Level(zerolog.Disabled)

	go func() {
		if err := codegen.StartServer(addr, lg); err != nil {
			panic(err)
		}
	}()

	// wait for server to start
	time.Sleep(time.Second)

	var err error

	codegenClient, err = oapi_client.NewClient(host)
	if err != nil {
		panic(err)
	}
}

func BenchmarkCodegenGetPropertiesInfo(b *testing.B) {
	startServerOnce.Do(setup)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := codegenClient.GetPropertiesInfo(context.Background())
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkCodegenCreateUpdatePropertyInfoById(b *testing.B) {
	startServerOnce.Do(setup)
	body := models.CreateUpdatePropertyInfoByIdJSONRequestBody{
		PropertyRating: lo.ToPtr(float32(5)),
		PropertyStatus: lo.ToPtr(models.PropertyInfoDataPropertyStatus("active")),
		PropertyUrl:    "https://example.com",
		RatingScale:    lo.ToPtr(10),
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := codegenClient.CreateUpdatePropertyInfoById(
			context.Background(),
			"123",
			body,
		)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkCodegenGetPropertyInfoById(b *testing.B) {
	startServerOnce.Do(setup)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := codegenClient.GetPropertyInfoById(context.Background(), "123")
		if err != nil {
			panic(err)
		}
	}
}
