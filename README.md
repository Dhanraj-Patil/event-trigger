# Event-Trigger
Sends scheduled text messages by making api calls to Twilio API when scheduled time is me. Achieved by using asynq to schedule tasks in future. Docker compose image consits of event-trigger server, worker server, mongodb, redis, asynqmon. Asynqmon helps in monitoring the data in queue. Event trigger data logs are retained for 2hr and are written to 'event-logs' dir which are later deleted in 46hrs. Total docker image size < 1Gib. Cost of running is upto $4 on digital ocean.

Could not deploy to cloud due to lack of free tier. Already used up free credits on GCP, Azure, digiralOcean, etc.

## How to run
- Run the following commands
```bash
https://github.com/Dhanraj-Patil/event-trigger.git
cd event-trigger
docker compose up -d
```

You can now view the API's via swaggerui at `http://localhost:8080/swagger/index.html` and view the tasks in queue or completed at `http://localhost:8081`.

## Example payload for create and test.
- Run the following commands
```json
{
  "message": "Hey",
  "phoneNo": "0000000000", //Enter your number here
  "schedule": "2025-02-16T22:24:00+05:30", //Edit time
  "userId": "dhanraj" //Enter a name as userId
}
```

## Example payload for update.
- Run the following commands
```params
  userId params: dhanraj //Enter a name as userId
```
```json
{
  "message": "Hey",
  "phoneNo": "0000000000", //Enter your number here
  "schedule": "2025-02-16T22:24:00+05:30", //Edit time
}
```
