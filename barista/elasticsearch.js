const barista = require('barista-forwarder');
const tail = require('barista-input-tail');
const elasticsearch = require('barista-output-elasticsearch');

barista.input('tail', tail('/var/log/nginx/error.log', (stream, resolve) => {
  return resolve(stream);
}));

barista.output('default', elasticsearch('host', 'port', (record, resolve) => {
  return resolve(record);
}));

barista.run();
