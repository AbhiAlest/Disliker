# Youtube-Dislike-Button-API

On November 10, 2021 YouTube removed the dislike button, greatly impacting the ability of individuals to determine the best possible content. Run this API/extension to see the public dislike counter.

Steps:

1. Input API key for Youtube Data API (via the Google Developers website)
2. <video_id> is a video's URL ID. Use route "/dislike/{video_id}"
3. Run API via the command "python youtube_api.py" for a flask app output
4. Make a GET request to the route's URL via the video ID. This will return a JSON object - - > {"dislike_count": "{insert dislike counter}"



This development is in Beta. Please be mindful of bugs. Currently working on turning this into an app/extension. 
