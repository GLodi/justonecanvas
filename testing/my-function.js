// my-functions.js

module.exports = { createTimestampedObject };

function createTimestampedObject(userContext, _, done) {
  userContext.vars.color = 1;
  userContext.vars.y = 1;
  userContext.vars.x = 1;

  return done();
}
