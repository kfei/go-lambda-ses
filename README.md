# go-lambda-ses

> An example of writing AWS Lambda function in Golang (through the Node shell),
> and sending emails via SES.

## Usage

  1. Clone this repo.
  2. `glide install && make build`
  3. Upload *lambda.zip* to create a new Lambda function.
  4. Deploy the API Gateway.
  5. Make sure all privilege/CORS/policy settings are good.
  6. Call the API with request in the form of:
```json
{
    "from": "<from@mail.address>",
    "to": "<to@mail.address>",
    "subject": "The interesting subject",
    "body": "..."
}
```

## Credits

The idea of shelling out Golang program comes from
[rjocoleman/lambda-ses](https://github.com/rjocoleman/lambda-ses).
