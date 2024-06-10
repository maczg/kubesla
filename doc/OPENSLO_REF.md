
## Datasource
Datasource crd is used to define the datasource for the application. The datasource is used to reference
which datasource has to be used by other resources like SLI, Alert, etc.

```yaml
apiVersion: openslo.io/v1alpha1
kind: Datasource
metadata:
  name: datasource-sample
spec:
  displayName: Example Prometheus
  description: This is an example Prometheus data source for monitoring.
  type: Prometheus
  connectionDetails:
    url: http://prometheus.example.com
``` 

### Status 
Controller will update the status of the datasource every minute. The status will contain the following fields:
- LastChecked (displayed) 
- StatusMessage (displayed)
- IsHealthy (displayed)

## SLO 
A service level objective (SLO) is a target value or a range of values for a service level that is described by a service level indicator (SLI).

It contains SLI spec by inline or reference to SLI crd.

One or multiple objectives can be defined in the SLO. In this case, each Objective
will share the same SLI but have different thresholds and conditions.

:warning: Folliwing field are explained: 

- timeWindow.Rolling time window: Duration should be provided in shorthand format e.g. 5m, 4w, 31d
- timeWindow.CalendarAligned:  Duration should be provided in shorthand format eg. 1d, 2M, 1Q, 366d
- budgetingMethod( Occurrencies | Timeslices| RatioTimeSlices):
    - Occurrences method uses a ratio of counts of good events to the total count of the events
    - Timeslices method uses a ratio of good time slices to total time slices in a budgeting period.
    - RatioTimeslices method uses an average of all time slices' success ratios in a budgeting period. 
- objectives Threshold  If thresholdMetric has been defined, only one Threshold can be defined. However, if using ratioMetric then any number of Thresholds can be defined.

## SLI 

SLI crd is used to define the Service Level Indicator for the application. The SLI is used to define the
metrics that are used to measure the service level of the application.

One of ratioMetric or thresholdMetric **MUST** be defined.

### ratioMetric
{good, total}, {bad, total} or raw data is used to calculate the ratio.

- counter (true/false): specifies whether the metric is monitonically increasing counter.
- good: represents the query used for gathering data from metric sources as numerator. 
Received data is used to compare objectives (threshold) values to find a good values. 
Exclusively declared with Bad
- bad: represents the query used for gathering data from metric sources as denominator.
Received data is used to compare objectives (threshold) values to find a bad values.
- total: represents the query used for gathering data from metric sources as denominator.
- rawType: enum(success | failure), required when using raw, specifies whether the ratios represented by the "raw" ratio metric are of successes or failures. Not to be used with good and bad as picking one of those determines the type of ratio.
- raw:  represents the query used for gathering already precomputed ratios. The type of ratio (success or failure) is specified using rawType.





### thresholdMetric
Metric, represents the query used for gathering data from metric sources. 
Raw data is used to compare objectives (threshold) values.