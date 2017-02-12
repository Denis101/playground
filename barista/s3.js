const barista = require('barista-forwarder');
const tail = require('barista-input-tail');
const s3 = require('barista-output-s3');

barista.input('tail', tail('/var/log/nginx/error.log', (stream, resolve) => {
  return resolve(stream);
}));

barista.output('default', s3('aws_key', 'aws_secret', 'bucket', (record, resolve) => {
  return resolve(record);
}));

barista.run();
