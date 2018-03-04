const config = require('../config');
const db = require('monk')(config.mongoUrl);
const userCollection = db.get('users');

const users = require('express').Router();

users.post('/', (req, res) => {
  let post = req.body;

  userCollection.insert(
      {firstName: post.firstName, lastName: post.lastName, email: post.email }
  ).then((data) => {
      res.json(data);
  }).catch((err) => {
    //TODO implement better error handling
    console.error(err);
  }).then(() => db.close());

});

module.exports = users;
