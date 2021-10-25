const express = require('express');
const redis = require('redis');
const client = redis.createClient();
const app = express();

app.use(express.json());

app.get('/kanye', (req, resp) => {
    resp.send('WEST');
});

app.listen(3000)

client.on('error', err => {
    console.log(err);
});

client.on('connect', () => {
    console.log('connected...');
});

client.set('kan', 'ye');