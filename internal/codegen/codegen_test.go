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
	chiAddr  = "localhost:18080"
	echoAddr = "localhost:18081"
)

var (
	startServerOnce   sync.Once
	chiCodegenClient  *oapi_client.Client
	echoCodegenClient *oapi_client.Client
)

func setup() {
	lg := log.Logger.Level(zerolog.Disabled)

	go func() {
		if err := codegen.StartServer(chiAddr, lg, codegen.NewChiHandler(lg)); err != nil {
			panic(err)
		}
	}()

	go func() {
		if err := codegen.StartServer(echoAddr, lg, codegen.NewEchoHandler(lg)); err != nil {
			panic(err)
		}
	}()

	// wait for server to start
	time.Sleep(time.Second)

	var err error

	chiCodegenClient, err = oapi_client.NewClient("http://" + chiAddr + "/api/v1")
	if err != nil {
		panic(err)
	}

	echoCodegenClient, err = oapi_client.NewClient("http://" + echoAddr + "/api/v1")
	if err != nil {
		panic(err)
	}
}

func BenchmarkCodegen_Chi_GetPropertiesInfo(b *testing.B) {
	startServerOnce.Do(setup)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := chiCodegenClient.GetPropertiesInfo(context.Background())
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkCodegen_Chi_CreateUpdatePropertyInfoById(b *testing.B) {
	startServerOnce.Do(setup)
	body := models.CreateUpdatePropertyInfoByIdJSONRequestBody{
		PropertyRating: lo.ToPtr(float32(5)),
		PropertyStatus: lo.ToPtr(models.PropertyInfoDataPropertyStatus("active")),
		PropertyUrl:    "https://example.com",
		RatingScale:    lo.ToPtr(10),
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := chiCodegenClient.CreateUpdatePropertyInfoById(
			context.Background(),
			"123",
			body,
		)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkCodegen_Chi_GetPropertyInfoById(b *testing.B) {
	startServerOnce.Do(setup)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := chiCodegenClient.GetPropertyInfoById(context.Background(), "123")
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkCodegen_Echo_GetPropertiesInfo(b *testing.B) {
	startServerOnce.Do(setup)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := echoCodegenClient.GetPropertiesInfo(context.Background())
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkCodegen_Echo_CreateUpdatePropertyInfoById(b *testing.B) {
	startServerOnce.Do(setup)
	body := models.CreateUpdatePropertyInfoByIdJSONRequestBody{
		PropertyRating: lo.ToPtr(float32(5)),
		PropertyStatus: lo.ToPtr(models.PropertyInfoDataPropertyStatus("active")),
		PropertyUrl:    "https://example.com",
		RatingScale:    lo.ToPtr(10),
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := echoCodegenClient.CreateUpdatePropertyInfoById(
			context.Background(),
			"123",
			body,
		)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkCodegen_Echo_GetPropertyInfoById(b *testing.B) {
	startServerOnce.Do(setup)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := echoCodegenClient.GetPropertyInfoById(context.Background(), "123")
		if err != nil {
			panic(err)
		}
	}
}
