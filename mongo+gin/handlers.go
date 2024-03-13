package main

import (
    "context"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "net/http"
    "strconv"
    "log"
)

type AdvertiseHandler struct {
    client *mongo.Client
}

func NewAdvertiseHandler(client *mongo.Client) *AdvertiseHandler {
    return &AdvertiseHandler{
        client: client,
    }
}

func (h *AdvertiseHandler) GetAdvertises(c *gin.Context) {
    var Advertises []Advertise
    db := h.client.Database("Advertise")
    collection := db.Collection("Advertise")
     cur, err := collection.Find(context.TODO(), bson.D{{}})
     if err != nil {
         c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
         return
     }
     defer cur.Close(context.TODO())
 
     if err = cur.All(context.TODO(), &Advertises); err != nil {
         c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
         return
     }
 
     c.IndentedJSON(http.StatusOK, Advertises)
}

func (h *AdvertiseHandler) getAdvertiseByCondition(c *gin.Context) {
    offsetStr := c.Query("offset")   // 获取offset参数的值
    limitStr := c.Query("limit")     // 获取limit参数的值
    ageStr := c.Query("age")         // 获取age参数的值
    gender := c.Query("gender")   // 获取gender参数的值
    country := c.Query("country") // 获取country参数的值
    platform := c.Query("platform") // 获取platform参数的值
    var Advertises []Advertise
    db := h.client.Database("Advertise")
    collection := db.Collection("Advertise")
    filter := bson.D{}
    age ,err := strconv.Atoi(ageStr)
    if age != 0 {
        filter = append(filter, bson.E{Key: "conditions.ageStart", Value: bson.D{{Key: "$lte", Value: age}}})
        filter = append(filter, bson.E{Key: "conditions.ageEnd", Value: bson.D{{Key: "$gte", Value: age}}})
    }
    if gender != "" {
        filter = append(filter, bson.E{Key: "$or", Value: bson.A{
            bson.D{{"conditions.Gender", bson.D{{"$in", bson.A{gender}}}}},
            bson.D{{"conditions.Gender", bson.A{}}},                       
        }})
    }
    if country != "" {
        filter = append(filter, bson.E{Key: "conditions.country", Value: bson.D{{Key: "$in", Value: bson.A{country}}}})
    }
    if platform != "" {
        filter = append(filter, bson.E{Key: "conditions.platform", Value: bson.D{{Key: "$in", Value: bson.A{platform}}}})
    }
   
    findOptions := options.Find()
    findOptions.SetSort(bson.D{{"endAt", 1}})
    offset, err := strconv.ParseInt(offsetStr, 10, 64)
    limit, err := strconv.ParseInt(limitStr, 10, 64)
    findOptions.SetSkip(offset)
    findOptions.SetLimit(limit)
    cur, err := collection.Find(context.TODO(), filter,findOptions)
    if err != nil {
        c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer cur.Close(context.TODO())

    if err = cur.All(context.TODO(), &Advertises); err != nil {
        c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.IndentedJSON(http.StatusOK, Advertises)

}

func (h *AdvertiseHandler) postAdvertises(c *gin.Context) {
    var newAdvertise Advertise
    db := h.client.Database("Advertise")
    collection := db.Collection("Advertise")
    if err := c.BindJSON(&newAdvertise); err != nil {
        return
    }
    _, err := collection.InsertOne(context.TODO(), newAdvertise)
    if err != nil {
        log.Fatal(err)
    }
    c.IndentedJSON(http.StatusCreated, newAdvertise)
}