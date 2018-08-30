# GoAlmanac

GoAlmanac is a simple service that retrieves weather information from 
[Wunderground](https://api.wunderground.com), formats and writes the response
to a separate json file. 

## Services
As currently written, a static location is requested when `main.go` is invoked.
The `WundergroundApi` returns a large set of data - of which the minimal data
points and service information is stored and written to file. This relies on a
valid `WUNDERGROUND_API_FORECAST` key. 

### Current Roadmap
The early inceptions of `GoAlmanac` are intended to gather and store sets of
weather data.

```sh
v0.1 - Weather retrieval and file write service
## going forward =>
v0.2 - Cron job to automagically fetch and store weather data
v0.3x++ - postgres migration, data import/seed db service
v1 - gui - visualizations and ui
v2++ - Start training ml algorithm for short term forecasting
v3++ - training on longer term trends and anomaly training
v5 - it takes over the world, or tells you tomorrow's weather with impeccable
accuracy
```
