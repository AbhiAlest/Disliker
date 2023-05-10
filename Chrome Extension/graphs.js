const urlParams = new URLSearchParams(window.location.search);
const videoId = urlParams.get('v');

fetch(`https://www.googleapis.com/youtube/v3/videos?part=statistics&id=${videoId}&key=${API_KEY}`)
  .then(response => response.json())
  .then(data => {
    const likeCount = parseInt(data.items[0].statistics.likeCount);
    const dislikeCount = parseInt(data.items[0].statistics.dislikeCount);

    const ratio = likeCount / dislikeCount;

    const newDiv = document.createElement('div');
    newDiv.style.position = 'absolute';
    newDiv.style.top = '0';
    newDiv.style.left = '0';
    newDiv.style.width = '300px';
    newDiv.style.height = '300px';
    document.body.appendChild(newDiv);

    const ctx = newDiv.getContext('2d');
    const chart = new Chart(ctx, {
      type: 'line',
      data: {
        labels: ['1', '2', '3', '4', '5'],
        datasets: [{
          label: 'Like to Dislike Ratio',
          data: [ratio, ratio, ratio, ratio, ratio],
          fill: false,
          borderColor: 'rgb(75, 192, 192)',
          tension: 0.1
        }]
      }
    });

    newDiv.appendChild(chart.canvas);

    const ratioBtn = document.createElement('button');
    ratioBtn.innerText = 'Like to Dislike Ratio Graph';
    ratioBtn.addEventListener('click', () => {
      chart.data.datasets[0].label = 'Like to Dislike Ratio';
      chart.data.datasets[0].data = [ratio, ratio, ratio, ratio, ratio];
      chart.update();
    });

    const dislikesBtn = document.createElement('button');
    dislikesBtn.innerText = 'Dislikes over Time Graph';
    dislikesBtn.addEventListener('click', () => {
      fetch(`https://your-api-url/dislikes/${videoId}`)
        .then(response => response.json())
        .then(data => {
          chart.data.datasets[0].label = 'Dislikes over Time';
          chart.data.datasets[0].data = data.dislikes;
          chart.update();
        });
    });

    const buttonContainer = document.createElement('div');
    buttonContainer.style.position = 'absolute';
    buttonContainer.style.top = '0';
    buttonContainer.style.left = '310px';
    buttonContainer.style.display = 'flex';
    buttonContainer.appendChild(ratioBtn);
    buttonContainer.appendChild(dislikesBtn);
    newDiv.appendChild(buttonContainer);
  });
