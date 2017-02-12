const barista = require('barista-forwarder');
const tail = require('barista-input-tail');
const elasticsearch = require('barista-output-elasticsearch');
const s3 = require('barista-output-s3');

barista.input('tail', tail('/var/log/nginx/error.log', (stream, resolve) => {
  return resolve(stream);
}));

barista.filter((type, record, resolve) => {
  return resolve(type === 'tail' ? 'elasticsearch' : 's3');
});

barista.output('elasticsearch', elasticsearch('host', 'port', (record, resolve) => {
  return resolve(record);
}));

barista.output('s3', s3('aws_key', 'aws_secret', 'bucket', (record, resolve) => {
  return resolve(record);
}));

barista.run();
