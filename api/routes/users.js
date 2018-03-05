const config = require('../config');
const db = require('monk')(config.mongoUrl);
const userCollection = db.get('users');

const users = require('express').Router();

users.get('/', (req, res) => {
  //Pagination. Use if set, otherwise set to 1
  req.query.page = req.query.page || 1;

  //Return all users
  userCollection.find({}, {
    "limit": 10,
    "skip": (10 * req.query.page) - 10
  }).then((data) => {
    res.json(data);
  });

});

users.get('/:id', (req, res) => {
  let id = req.params.id;
  //Find user by Mongo ID
  userCollection.find({"_id": id}, { castIds: false }).then((data) => {
    res.json(data);
    //TODO return a message for an invalid/unknown id
  }).catch((err) => {
    res.json({"error": "Mongo connection failed. Please try again"});
    //TODO implement better error handling
    console.error(err);
  });

});

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

users.delete('/:id', (req, res) => {
  let id = req.params.id;
  userCollection.remove({"_id": id }).then(data => {
    res.json(data);
  }).catch((err) => {
    //TODO implement better error handling
    console.error(err);
  }).then(() => db.close());
});

module.exports = users;
