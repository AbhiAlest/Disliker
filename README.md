# Youtube-Dislike-Button-API

Youtube got rid of dislike button, which is kind of annoying. Run this API to see the public dislike counter.

Steps:

1. Input API key for Youtube Data API (via the Google Developers website)
2. <video_id> is a video's URL ID. Use route "/dislike/{video_id}"
3. Run API via the command "python youtube_api.py" for a flask app output
4. Make a GET request to the route's URL via the video ID. This will return a JSON object - - > {"dislike_count": "{insert dislike counter}"



This development is in Beta. Please be mindful of bugs. Currently working on turning this into an app/extension. 


Updates Listed:

04/14/2023 - 1.0 version debugging. Defined chrome extension. Working on formatting/installation, but dislike counter should be working. 
04/15/2023 - Defined Microsoft Edge extension.
