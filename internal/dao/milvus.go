package dao

import (
	"context"
	"fmt"

	"gomind/internal/config"

	"github.com/milvus-io/milvus/client/v2/milvusclient"
)

func InitMilvus(ctx context.Context, cfg config.MilvusConfig, addr string) (*milvusclient.Client, error) {
	client, err := milvusclient.New(ctx, &milvusclient.ClientConfig{
		Address:  addr,
		Username: cfg.Username,
		Password: cfg.Password,
		DBName:   cfg.Database,
	})
	if err != nil {
		return nil, fmt.Errorf("connect milvus: %w", err)
	}

	if _, err := client.GetServerVersion(ctx, milvusclient.NewGetServerVersionOption()); err != nil {
		if closeErr := client.Close(ctx); closeErr != nil {
			return nil, fmt.Errorf("check milvus server version: %w; close milvus client: %v", err, closeErr)
		}
		return nil, fmt.Errorf("check milvus server version: %w", err)
	}

	return client, nil
}
