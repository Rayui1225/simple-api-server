# simple-api-server
This is a simple-api-server imitate Dcard Backend Intern Assignment  
Use golang(gin) + mongodb  

# system design
This system use client-server architecture  
![ApiServerArch drawio](https://github.com/Rayui1225/simple-api-server/assets/49279418/0ffb3e33-7277-4850-82d9-65be264c6394)
# API endpoint  
## Get all advertisements  
Method : GET  
Path : /Advertise  
Parameter : none  
Success response :
```
[
    {
        "title": "AD 55",
        "startAt": "2023-12-10T03:00:00Z",
        "endAt": "2023-12-30T03:00:00Z",
        "conditions": {
            "ageStart": 18,
            "ageEnd": 27,
            "gender": ["F", "M"],
            "country": ["JP", "TW"],
            "platform": ["android"]
        }
    }
]
```
Fail response :
```
{
    "error": "Internal Server Error"
}
```

## Add new advertisement
Method : POST  
Path : /Advertise  
Request data example : 
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
Success response :
```
[
 "message": "Advertise created successfully"
]
```
Fail response :
```
{
    "error": "Internal Server Error"
}
```

## Get advertisements by Parameter
Method : GET  
Path : /Ad  
Parameter (Optional) :  
```
offset : The number of records you want to pass  
limit : The maximum number of records to return  
age : A specific age within the target age range for the advertisement  
gender : The target gender for the advertisement, can be 'F' or 'M'  
country : The target country for the advertisement, such as 'TW' or 'JP'  
platform : The target platform for the advertisement, such as 'ios' or 'android'
```  
Success response :
```
[
    {
        "title": "AD 55",
        "startAt": "2023-12-10T03:00:00Z",
        "endAt": "2023-12-30T03:00:00Z",
        "conditions": {
            "ageStart": 18,
            "ageEnd": 27,
            "gender": ["F", "M"],
            "country": ["JP", "TW"],
            "platform": ["android"]
        }
    }
]
```
Fail response :
```
{
    "error": "Internal Server Error"
}
```
# run this project 
```
git clone https://github.com/Rayui1225/simple-api-server.git
```  
```
cd simple-api-server
```  
```
go run main.go
```
## You can also use the provided Dockerfile to package this project into a Docker image.  
# Usage example
```curl -X GET "http://localhost:8080/Advertise"```
```
curl -X POST "http://localhost:8080/Advertise" -H "Content-Type: application/json" -d '{
    "title": "New Ad 66",
    "startAt": "2024-01-01T00:00:00Z",
    "endAt": "2024-01-31T00:00:00Z",
    "conditions": {
        "ageStart": 20,
        "ageEnd": 35,
        "gender": ["M"],
        "country": ["US"],
        "platform": ["ios", "android"]
    }
}'
```
