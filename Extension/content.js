function getDislikeCount(videoId) {
  return fetch(`https://your-app-url/dislike/${videoId}`)
    .then(response => response.json())
    .then(data => data.dislike_count || 0);
}

function injectDislikeCount(dislikeCount) {
  chrome.tabs.query({ active: true, currentWindow: true }, function(tabs) {
    const tabId = tabs[0].id;
    chrome.tabs.sendMessage(tabId, { type: 'dislikeCount', count: dislikeCount });
  });
}

chrome.webNavigation.onCompleted.addListener(function(details) {
  if (details.frameId === 0 && details.url.startsWith('https://www.youtube.com/watch')) {
    const videoId = new URLSearchParams(details.url.split('?')[1]).get('v');
    getDislikeCount(videoId).then(injectDislikeCount);
  }
});
