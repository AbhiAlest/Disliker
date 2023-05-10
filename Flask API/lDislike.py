#adds multilingual support for Spanish, Portuguese, Hindi, Korean, Japanese, French, and German

from flask import Flask, jsonify
import requests

app = Flask(__name__)

@app.route('/dislike/<video_id>/<language>', methods=['GET'])
def get_dislike_count(video_id, language):
    api_key = 'YOUTUBE_DATA_API_KEY' 
    url = f'https://www.googleapis.com/youtube/v3/videos?part=statistics&id={video_id}&key={api_key}&hl={language}'
    # replace language with appropriate abbreviation
    response = requests.get(url)
    data = response.json()

    if 'items' in data and len(data['items']) > 0:
        dislike_count = data['items'][0]['statistics']['dislikeCount']
        return jsonify({'dislike_count': dislike_count})
    else:
        return jsonify({'error': 'Invalid video ID'}), 404
