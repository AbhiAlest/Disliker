function getDislikeCount(videoId) {
  return fetch(`https://your-app-url/dislike/${videoId}`)
    .then(response => response.json())
    .then(data => data.dislike_count || 0);
}

function injectDislikeCount(dislikeCount) {
  browser.tabs.query({ active: true, currentWindow: true }).then(tabs => {
    const tabId = tabs[0].id;
    browser.tabs.sendMessage(tabId, { type: 'dislikeCount', count: dislikeCount });
  });
}

browser.runtime.onMessage.addListener(request => {
  if (request.type === 'getDislikeCount') {
    const videoId = new URLSearchParams(request.queryString).get('v');
    getDislikeCount(videoId).then(injectDislikeCount);
  }
});
