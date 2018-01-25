const AWS = require('aws-sdk');

const DEFAULT_HEADERS = {'Access-Control-Allow-Origin': '*'};

module.exports.status = function (event, context, callback) {
    console.log(`Request to see status`);

    const response = {
        statusCode: 200,
        headers: DEFAULT_HEADERS,
        body: JSON.stringify({
            message: 'Status: Good',
            input: event,
        }),
    };

    callback(null, response);
}
