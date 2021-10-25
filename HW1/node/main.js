const express = require('express');
const redis = require('redis');
const crypto = require('crypto');

const router = express.Router();
const app = express();
const client = redis.createClient();

app.use(express.json());

router.post('/node/sha256', function(req, res){
    console.log(req.body);
    const { data } = req.body;
    if (data.length < 8) return res.status(400).send(JSON.stringify({"message": "Your message length must be more than 8 characters!"}));

    hashed_data = crypto.createHash('sha256').update(data).digest('hex');
    client.set(hashed_data, data);
    return res.status(200).send(JSON.stringify({"message": "Data successfully saved."}));
});

router.get('/node/sha256/:data', function(req, res){
    const { data } = req.params;
    client.get(data, function(err, value){
        if (err) return res.status(500).send(JSON.stringify({"message": "Internal server error!", "error": err.message}));

        let message = value || 'Data does not exist in the database';
        return res.status(200).send(JSON.stringify({"message": message}));
    });
});

client.on('error', err => {
    console.log(err);
});

client.on('connect', () => {
    console.log('connected...');
});

app.listen(3000);
app.use('/', router);