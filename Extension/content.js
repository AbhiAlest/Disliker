function getDislikeCount(videoId) {
  return fetch(`https://your-app-url/dislike/${videoId}`)
    .then(response => response.json())
    .then(data => data.dislike_count || 0);
}

function injectDislikeCount(dislikeCount) {
  const container = document.querySelector('.ytd-video-primary-info-renderer');
  if (container) {
    const element = container.querySelector('#dislike-button');
    if (element) {
      const existingCount = element.querySelector('#dislikes');
      if (existingCount) {
        existingCount.textContent = dislikeCount.toLocaleString();
      } else {
        const count = document.createElement('span');
        count.id = 'dislikes';
        count.className = 'ytd-toggle-button-renderer-count-text';
        count.textContent = dislikeCount.toLocaleString();
        element.appendChild(count);
      }
    }
  }
}

function init() {
  const videoId = new URLSearchParams(location.search).get('v');
  if (videoId) {
    getDislikeCount(videoId).then(injectDislikeCount);
  }
}

init();
