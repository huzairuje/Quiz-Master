# Essay
Given 2 or more different services, please design a system for all
services to maintain up-to-date data without sharing their
databases—for example, concrete use case: user & billing services.
a. Then you have to explain to us how the data flows and how the data sync between databases.
b. And what stacks that you recommend, and why?
Hints and Tips!
1. You choose which methodology you use, such as BFF, Hexagonal
Architecture, Event-driven Architecture, etc.
2. You also can draw a simple architecture diagram by using external
websites like plantuml, lucid.app, etc. And upload the image here with
your explanation.

## Answer
1. User Service is to use store the data for users data 
2. Billing Service is to use store the data for billings
3. tech stacks golang, and kafka as the actor on event driven architecture

i will choose the **Event Driven Architecture** , for the leverage the power of event driven architecture microservices we can shift the thinking from *invoke the service* to *initiate and capture the events*, A systems publishing events that can be consumed by zero or more downstream services and interpreted within the context of each service. The amount of flexibility and decoupling this provides is really great. Let's take another look at the architecture behind an E-commerce with that in mind. 

![flow data users and billing](https://github.com/huzairuje/Quiz-Master/blob/master/essay/assets/images/img1.png)

the services are fully decoupled and no longer know about each other. 
Our ecommerce site is now able to handle large demand spikes because the only services that need to be scaled are the WebApp, The other services can be scaled to average load as they no longer impact the customer's order experience. Let's look at the messaging patterns that make up our new event-driven microservices architecture.

Here you see how the generated event can be produced and consumed by the Inventory Manager, and a response event can be sent back to the customer once it's confirmed. There are a few key concepts here:

First that events are sent out with topics. These topic names are strings which can contain a hierarchy; thus supporting interest-based subscriptions.
Secondly, services will bind to a queue to consume events. These events will be attracted to the queue based on the queue's subscription.
Lastly, the queue will be non-exclusive, meaning that multiple consumers can receive messages from the queue in a round-robin fashion to support service scaling. Essentially the messaging component replaces the need for a load balancer!
Since the ecommerce store produced an event, this event can be consumed by zero or more recipients. In our case, it is consumed by the Inventory Manager, but what if we wanted to have those events consumed by an analytics engine as well? This is where the big benefit lies as you can add these kinds of capabilities without touching the Web App or Inventory Manager. After all, microservices are all about the ability to create new capabilities faster!

One-way Notification
The combination of messaging and event-driven architecture really begins to shine when events can be sent in a "fire and forget" way. The one-way notification pattern is a great example of this mode. It was performed in the previous step when the Inventory Manager sent a message indicating that an order's set of products were reserved. From a process perspective, one service that would be interested in those events would be billing. The below diagram shows the billing service consuming those events and acquiring funds from the customer.

Messaging Patterns for Event-Driven Microservices

A common challenge with microservices architecture is the need to avoid XA transactions which aren't supported. From the business perspective, if the credit card bounces, we need to cancel the order and remove the reservation on the product so others can buy it. By making use of the concept of "eventual consistency" we can avoid XA transactions entirely. The key to eventual consistency is that the events must never be lost, or the system will be in a perpetual state of inconsistency. Guaranteed messaging ensures that the information will be delivered to another application, even in the event of component or network problems. REST cannot provide this guarantee.

In our case, if there were no funds, the Inventory Manager needs to subscribe to these events so it can undo the product reservation and let the customer know their order has been cancelled when a payment problem occurs. The key is that the Billing service does not know what steps need to be taken to roll back the order; it simply produces events indicating "paid" or "not paid." Clean and elegant!

Single Request-Multi Response
If you have ever used Uber or Lyft, you will have seen that the drivers react immediately if their phone dings with another customer in the area needing a ride. Those notifications go to all drivers in the area, and it's the first to respond that wins the ride/business. This is exactly how the commerce store delivery mechanism is intended to work. The ability to send an event to multiple clients and receive multiple responses is key to this use case and messaging is a perfect enabler. This diagram illustrates how this can be implemented:

Messaging Patterns for Event-Driven Microservices

Once the Inventory Manager has fulfilled the order, it sends an event to the topic "order/ship//". If you're wondering why we would include delivery zip in a topic name, that's because if each car subscribes to events for a certain geography, we have a clean routing mechanism so drivers won't be overwhelmed by events outside of their area. In this case, the Logistics Manager receives ship notifications and sends out an event indicating that it's looking to locate a delivery vehicle.

A delivery vehicle (or the driver's mobile device) is presented with a notification that an order is ready for delivery. They choose to accept or deny the opportunity. As stated before, the Logistics service only needs a single accept, and for fairness, selects the driver who responded first.

Publish/Subscribe
The use of messaging is most commonly associated with the publish/subscribe messaging pattern. Publish/subscribe is an element of some of the messaging patterns we have explored, but it can also be used directly to enhance the user's experience. In our case, wouldn't it be great if the customer could track the delivery vehicle, in real-time, that contains our order? As far-fetched as it sounds, many believe that these types of capabilities will be the future of e-commerce. For example, what if you opted into a FedEx or UPS service where they could deliver a package to your parked car while you are at work? In our case, the user will navigate to our web app where a mapping service (such as Google Maps) will be tracking his shoes in real-time!

Messaging Patterns for Event-Driven Microservices

Because we used publish/subscribe, these events could be consumed by multiple components based on their interest. There is no reason for the Web App to consume these events unless a user navigates to that screen. Logistics, on the other hand, may wish to always receive these events to ensure the driver follows the predefined route and that they do not speed excessively.

Conclusion - Unlocking the Value of Messaging Patterns for Event-Driven Microservices
Hopefully the ecommerce will be implemented one day and fulfill my dreams! In the meantime, I hope that you found it an interesting way to look at the various ways messaging patterns can unlock the full benefits and value of event-driven microservices. Consider adding these patterns to your toolbox and experience the real benefits provided by this architecture.
