import axios from 'axios';

export default axios.create({
    baseURL:'https://api.unsplash.com',
    headers: {
        Authorization:
        'Client-ID 6c1075170bebd4b6eb7f9f939482f24c90c87704c38a8c85604fd1a53d6e1c63',
    }
});