## Racing Proto

### API

#### Race
* **ID**: The unique identifier for the race (int64)
* **MeetingID**: The unique identifier for the races meeting (int64)
* **Name**: The name given to the race (string)
* **Number**: The number of the race (int64)
* **Visible**: Whether the race is visible (bool)
* **AdvertisedStartTime**: The time the race is advertised to run (Timestamp)
* **Status**: Uses the AdvertisedStartTime to determine if the race is open or closed (string)

#### RPCs
* `ListRaces(ListRacesRequest) returns ListRacesResponse`
* `GetRaceById(GetRaceRequest) returns GetRaceResponse`

#### Requests
* **ListRacesRequest**: Request which takes `ListRacesRequestFilter` and `order_by` as inputs 
* **GetRaceRequest**: Request which accepts an `int64` field `id`

#### Responses
* **ListRacesResponse**: The response containing all races that match the filtering criteria
* **GetRaceResponse**: The response containing a single Race corresponding to the id specified in `GetRaceRequest`

#### Filters
* **ListRacesRequestFilter**: Input into ListRaceRequest that filters the races result. Can provide two optional inputs into the filter:
    1. List of `int64` meeting ids to filter a subset of the results
    2. A `bool` variable to filter on race visibility