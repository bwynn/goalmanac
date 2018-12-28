# GoAlmanac

## v0.2

GoAlmanac is a simple service that retrieves weather information from 
[Wunderground](https://api.wunderground.com), formats and writes the response
to a separate json file. 

```sh
v0.1 - Weather retrieval and file write service
v0.2 - Cron job to automagically fetch and store weather data
v0.2x - migrate service to docker container
v0.3x++ - postgres migration, data import/seed db service
v1 - gui - visualizations and ui
v2++ - Start training ml algorithm for short term forecasting
v3++ - training on longer term trends and anomaly training
v5 - it takes over the world, or tells you tomorrow's weather with impeccable
accuracy
```

# Weather Tracker Service ******
A db collection of weather data and several software 
services centered around the management of that data;
managed via a docker image on a remote server. 

## nginx web server

## go-cron service(s)
a collection of services that run on the web server

### fetches weather data every 6 hrs
6 hour interval is data retrieval is based on the 
updates from global weather services (`GFS`, `EMWCF`, etc)
intervals. 

When service is dispatched, the handler function gracefully
manages the response from the third-party weather service: 
[wunderground](https://api.wunderground.com/developer). 

The service is contingent on the `WUNDERGROUND_API_FORECAST` 
environment variable, which contains the third-party service
api key. This is/can be configured in the .bashrc or .bash_profile
files in your root directory, or by using the `export
WUNDERGROUND_API_FORECAST=<your_secret_key>` within the terminal
session where the service is being configured.

Responses are forwarded to their appropriate handlers:

Response meta is entered into log db - 200/ok responses are
given abbreviated detail entries, all other responses are given
more detail for debugging.

Errors are directed to a secondary handler/service and cron-job that changes
the interval for weather data to 5 minutes and logs all responses
meta information to a log db

200 Response data is entered into db, after passing through
validation services to ensure proper parsing/entry.

## postgres db
data table/collection of weather statistics

    * forecast from GFS for location
        - precip %
        - temp
        - etc...

    * actual values

table/collection log for cron service?
    - successful response receipts
    - capture detailed error/log info

## client web app
something using the latest, greatest javascript libraries,
svgs, and some stylin' sass

### future ml service 
ml program/service that is trained on some set of data and then 
smooshed into the dataset found within the db - applies ml 
algorithm to short-term weather forecasting (for one location) 
vs GFS ensemble.
