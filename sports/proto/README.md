## Sports Proto

### API

#### SportsEvent
* **ID**: The unique identifier for the sports event (int64)
* **Name**: The name given to the event (string)
* **SportType**: The type of sport being played at the event (string)
* **State**: The State where the sports event is located in (string)
* **Visible**: Whether the sports event is visible (bool)
* **AdvertisedStartTime**: The time the sports event is advertised to start (Timestamp)
* **Status**: Uses the AdvertisedStartTime to determine if the sports event is open or closed (string)

#### RPCs
* `ListSportsEvents(ListSportsEventsRequest) returns ListSportsEventsResponse`
* `GetSportsEventById(GetSportsEventRequest) returns GetSportsEventResponse`

#### Requests
* **ListSportsEventsRequest**: Request which takes `ListSportsEventsRequestFilter` and `order_by` as inputs 
* **GetSportsEventRequest**: Request which accepts an `int64` field `id`

#### Responses
* **ListSportsEventsResponse**: The response containing all sports events that match the filtering criteria
* **GetSportsEventResponse**: The response containing a single sports event corresponding to the id specified in `GetSportsEventRequest`

#### Filters
* **ListSportsEventsRequestFilter**: Input into ListSportsEventsRequest that filters the sports events result. Can provide three optional inputs into the filter:
    1. A `bool` variable to filter on the sports event visibility
    2. A list of `string`s specifying the sport type to filter and view
    3. A list of `string`s specifying which states the events are in to filter and view