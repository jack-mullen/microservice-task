const config = require('../config');
const db = require('monk')(config.mongoUrl);
const userCollection = db.get('users');

const users = require('express').Router();

users.post('/', (req, res) => {
  let post = req.body; //request data

  req.checkBody("email", "Enter a valid email address.").isEmail();
  req.checkBody("firstName", "Enter a first name").isAlphanumeric();
  req.checkBody("lastName", "Enter a last name").isAlphanumeric();

  // If there are errors, send them out and stop processing
  let errors = req.validationErrors();
  if (errors) {
    res.send(errors);
    return;
  }

  // insert data into the 'users' collection
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
