# Client Service

Client service represents component which simulates end clients of the system, which want just to order and pick up food from
restaurant.
Client service have to generate end users orders and send orders to Food Ordering service. Client service consists of
clients which will actually generates orders and represents abstraction of a real system client.
Main work unit of client service is Client . Each client is a separate work unit which works completely separate and isolated
with respect to other clients from client service.

In Client service we have multiple independent work units which are clients .
Each client have to be implemented as a dedicated thread. Each client implements logic of generating new totally random order
with random number of foods and random foods assigned to this order. Actually client is just a combination of table and
waiter from dinning hall .
In order to know which restaurants are connected to the Food ordering system, each client first of all have to request from
Food ordering service, data about available restaurants which will include each restaurant menu.

Having all restaurants menu, Client generates random order and send generated order to food ordering service. Client can order from multiple
restaurants, that means that client final order send to Food ordering service can include multiple restaurants orders.
Client order is represented by at least 1 restaurant order, but usually client should order from multiple restaurants. For each
restaurant order maximum wait time that a client is willing to wait before order pick up, should be calculated by taking the item
with the highest preparation-time from the restaurant order and multiply it by 1.8 .

As a response sending generated order to Food ordering service, client receives data which includes estimated
preparation time for each order for a dedicated restaurant. Client waits for this time plus some additional realistic time
coefficient and after this, performs request to restaurant dinning hall to pick up the order.
If order is not ready at the time client performs request to the restaurant dinning hall in order to pick up his order.
Client waits for the order the time specified in response from restaurant dinning hall .
Each client should be created and destroyed for each new order request. This means, clients have to be destroyed after
they pick up their orders. New clients are created instead of destroyed cients .
