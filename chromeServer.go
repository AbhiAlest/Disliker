package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/analytics/v3"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

type VideoStats struct {
	DislikeCount int64 `json:"dislike_count"`
}

func getDislikeCount(w http.ResponseWriter, r *http.Request) {
	videoID := mux.Vars(r)["video_id"]

	if !isAuthorized(r) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	dislikeCount, found := getFromCache(videoID)
	if found {
		videoStats := VideoStats{
			DislikeCount: dislikeCount,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(videoStats)
		return
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer client.Disconnect(ctx)

	videoCollection := client.Database("videos").Collection("stats")
	filter := bson.M{"video_id": videoID}
	var videoStat VideoStat
	err = videoCollection.FindOne(ctx, filter).Decode(&videoStat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if time.Since(videoStat.Timestamp) < 1*time.Hour {
		videoStats := VideoStats{
			DislikeCount: videoStat.DislikeCount,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(videoStats)

		addToCache(videoID, videoStat.DislikeCount)

		return
	}

	youtubeService, err := youtube.NewService(context.Background(), option.WithAPIKey(os.Getenv("YOUTUBE_API_KEY")))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	videoResponse, err := youtubeService.Videos.List("statistics").
		Id(videoID).
		Do()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	videoStat.DislikeCount = videoResponse.Items[0].Statistics.DislikeCount
	videoStat.Timestamp = time.Now()
	update := bson.M{"$set": videoStat}
	_, err = videoCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Query the YouTube API for the video's metadata.
	videoMetadata, err := youtubeService.Videos.List("snippet").
		Id(videoID).
		Do()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get user's authentication token from headers and perform authorization.
	authToken := r.Header.Get("Authorization")
	if authToken != "myAuthToken" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Connect to MongoDB.
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	videosCollection := client.Database("myDB").Collection("videos")
	_, err = videosCollection.InsertOne(ctx, videoMetadata.Items[0])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	usersCollection := client.Database("myDB").Collection("users")
	_, err = usersCollection.InsertOne(ctx, bson.M{
		"auth_token": authToken,
		"timestamp":  time.Now().Unix(),
		"video_id":   videoID,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	gaService, err := analytics.NewService(ctx, option.WithCredentialsFile("path/to/credentials.json"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = gaService.Events.Insert("myGAID", &analytics.Event{
		Category: "Video",
		Action:   "Watched",
		Label:    videoMetadata.Items[0].Snippet.Title,
	}).Do()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}


  mongoURI := os.Getenv("MONGO_URI")
  client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  defer client.Disconnect(ctx)

  database := client.Database("youtube_stats")
  collection := database.Collection("video_stats")

_, err = collection.InsertOne(ctx, bson.M{"video_id": videoID, "dislike_count": videoStats.DislikeCount})
    if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
    }


    cursor, err := collection.Aggregate(ctx, mongo.Pipeline{{"$group", bson.D{{"_id", ""}, {"average_dislikes", bson.D{{"$avg", "$dislike_count"}}}}}})
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer cursor.Close(ctx)

    var results []bson.M
    if err = cursor.All(ctx, &results); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(results)

  
	key := "video_stats_" + videoID
	if !cache.Contains(key) {
		cache.Set(key, videoStats, 5*time.Minute)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(videoStats)
}

