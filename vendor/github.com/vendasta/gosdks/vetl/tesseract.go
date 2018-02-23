package vetl

import (
	"github.com/vendasta/gosdks/vstore"
	"fmt"
	"context"
	"github.com/vendasta/gosdks/logging"
)

// VStoreModelToTesseract provides a helper method for easily registering a vStore model with Tesseract
func VStoreModelToTesseract(ctx context.Context, client Interface, id string, model vstore.Model, PubsubID string, fields []string, primaryKey []string) error {
	schema := model.Schema()
	logging.Debugf(ctx, "Attempting to Register %s Kind with vETL.", schema.Kind)

	// vStore Schema -> vETL Schema
	vetlSchema, err := DataSourceFromVStoreModel(model, PubsubID)
	if err != nil {
		return fmt.Errorf("failed to initialize the vETL schema for %s because: %s", id, err.Error())
	}

	// Register vStore with vETL for the given Model
	err = client.CreateDataSource(ctx, id, vetlSchema)
	if err != nil {
		return fmt.Errorf("failed to create vETL Datasource: %s", err.Error())
	}

	// Register a transformation that only keeps the provided fields
	transformID := fmt.Sprintf("%s-dimensional-fields", id)
	err = client.UpsertTransform(ctx, []string{id}, transformID, KeepPropertiesTransform(fields), true)
	if err != nil {
		return fmt.Errorf("failed to upsert vETL Transform: %s", err.Error())
	}

	// Registers the (dataset -> transform) to be pushed into Tesseract
	err = client.CreateSubscription(ctx, fmt.Sprintf("%s-tesseract", id), transformID, TesseractDataSink(schema.Namespace, schema.Kind, primaryKey))
	if err != nil {
		return fmt.Errorf("failed to create vETL Subscription: %s", err.Error())
	}

	return nil
}
