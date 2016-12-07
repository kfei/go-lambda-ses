const failRes = (err) => {
    return JSON.stringify({
        statusCode: '400',
        errorCode: 5566,
        detailedMessage: err,
    })
}

exports.handler = (event, context) => {

    process.on('uncaughtException', err => context.fail(failRes(err)))

    event.sender= process.env.SENDER

    var child = require('child_process').spawn('./go-lambda-ses', [JSON.stringify(event)], { stdio: 'inherit' })

    child.on('close', (code) => {
        if (code !== 0 ) {
            context.fail(failRes("Process exited with non-zero status code: " + code))
        } else {
            context.succeed(null, null)
        }
    });
};
