A basic comment posting using Go, event emitters, and json. The json
can be replaced with the chosen database connection.
The purpose of this was to attempt to recreate the listeners of Flux 
on the server side; I am still not clear how the differential between
Go server code and client Javascript works out, and how much 
is actually considered "flux". This code uses no javascript

Flux works like this:

Client action or server action on webpage creates an action, which is sent to the Dispatcher
The Dispatcher maintains a list of stores, and broadcasts a signal with the action type and the new data entered
  If there are dependencies the Dispatcher can wait for a view to finish updating before broadcasting
There are stores listening in to the dispatcher with callback functions registered, they execute
these functions based on the action type they recieve using a switch statement. When done, the store
broadcasts an update signal.
There are views listening in on the stores. When they recieve the store update signal, the view
refreshes the entire part it is responsible for, or rerenders that part of the DOM.
From what I understand, the store does not send the new data with its signal; the view, 
when recieving an update signal, simply rereads the database again and refreshes itself entirely. 

The listening can be done in several ways such as EventEmitters or websockets.

See these links for more information:
https://medium.com/brigade-engineering/what-is-the-flux-application-architecture-b57ebca85b9e#.r3mnvbkve
https://scotch.io/tutorials/getting-to-know-flux-the-react-js-architecture

