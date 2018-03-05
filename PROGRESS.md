# Structure #
I created a decoupled structure of an API service (/api) and a frontend service (/web).

The frontend is built in Vue, which has been compiled and is served through an Nginx server. The backend is built in express/node, as it is lightweight and quick for prototyping API structures.

# Challenges #
A lot of my time was spent working out how to correctly use Docker. I have not used Docker a lot before outside of personal testing and projects, and therefore spent considerable time learning about Docker, the composer feature, and the network links.

Time was also a factor, so I tried to get a barebones API structure and UI created, with the hope that I could add features and iron out issues in increments.

# To Do and Critique #
There is still quite a bit which needs doing. There is currently no way to update a user (I would do this via a PUT request to the API). There is also no demonstration of a single GET for a particular user using the UI, though this functionality does work through the api directly if you want to try it out.

There are some minor issues with the UI that I would also resolve. The "Delete" button works but does not reactively update the users list (you need to refresh). The data flow through Vue, while not terrible, could be improved by using a state library (such as VueX).
