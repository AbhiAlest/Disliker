<h1 align="center">Disliker</h1>
<br />

<p>
**Disliker is an open-source project that provides a simple and convenient way to view the number of dislikes for any YouTube video. **
</p>
<p>
The project consists of two main components: an API that extracts the dislike count from the YouTube video page, and a browser extension that displays the dislike count to the user.
</p>
<p>
This development is in Beta. Please be mindful of bugs. Currently working on turning this into an app/extension. 
</p>

---

<h2>API Usage</h2>

1. Input API key for Youtube Data API (via the Google Developers website).
2. <video_id> is a video's URL ID. Use route 
```
/dislike/{video_id}
```
4. Run API via the command for a flask app output
```python
python youtube_api.py
```
5. Make a GET request to the route's URL via the video ID. This will return a JSON object 
```json
{
"dislike_count": {insert dislike counter} 
}
```
Multilingual support: Currently, Disliker is being adapted to display the public dislike count for a variety of non-English languages including Spanish, Portuguese, Hindi, Korean, Japanese, French, and German. 
Given this example local server: http://localhost:5000/dislike/VIDEO_ID/{language}, replace {language} with appropriate lingual abbreviations: 'es', 'pt', 'hi', 'ko', 'ja', 'fr', or 'de'




