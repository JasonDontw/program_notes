import axios from 'axios';

const KEY = 'AIzaSyDJ9ytNHSjtjGm87vd8e6iQ52dE6KhZdqM';

export const baseParams = {
    part: "snippet",
    maxResults: 5,
    key: KEY
  };

export default axios.create({
    baseURL: "https://www.googleapis.com/youtube/v3",
    params: baseParams
});
