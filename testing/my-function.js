// my-functions.js

module.exports = { createTimestampedObject };

function createTimestampedObject(userContext, events, done) {
  const color = Math.floor(Math.random() * (16 + 1));
  const random = Math.floor(Math.random() * (1600 + 1));

  const data = new Uint8Array([color, Math.floor(random / 40), random % 40])
    .buffer;

  userContext.vars.color = 1;
  userContext.vars.y = 1;
  userContext.vars.x = 1;

  return done();
}
