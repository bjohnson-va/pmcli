## 1.0.1
- Add customer id to PredictionWithMetrics to indicate who that prediction/data was for

## 1.0.0
- List allows you to list both churn predictions and historical data for a specific customer, bundled with the associated metrics for each time period
- Prioritize allows you to list predictions across all customers or a defined subset of customers for the latest time period, and returns those predictions ordered by priority
- IngestPredictions allows you to command the microservice to ingest the requests/results of a complete prediction job for a model id 
- IngestTraining allows you to command the microservice to ingest the requests of a complete train job for a model id 
