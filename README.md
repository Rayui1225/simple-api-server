# simple-api-server
This is a simple-api-server imitate Dcard Backend Intern Assignment  
use golang(gin) + mongodb  
request get http://<host>/Advertise will get all advertisements  
request post http://<host>/Advertise and some data ex:
```
'{
"title" "AD 55",
"startAt" "2023-12-10T03:00:00.000Z",
"endAt" "2023-12-31T16:00:00.000Z",
"conditions": {
{
"ageStart": 20,
"ageEnd": 30,
"country: ["TW", "JP"],
"platform": ["android", "ios"]
}
}
}'
```
will insert this data to DB  

request get http://<host>/Ad?offset=10&limit=3&age=24&gender=F&country=TW&platform=ios will use request's paremeter to search data from DB each request can be none 
