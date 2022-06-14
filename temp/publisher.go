// 필요해 보이는 부분만 솎기

package main

import (
	"errors"
	"fmt"
	"github.com/streadway/amqp"
)

//MessageBody is the struct for the body passed in the AMQP message. The type will be set on the Request header
type MessageBody struct {
	Data []byte
	Type string
}

//Message is the amqp request to publish
type Message struct {
	Queue           string
	DeliveryMode    uint8
	ReplyTo         string
	ContentType     string
	ContentEncoding string
	CorrelationID   string
	Priority        uint8
	Body            MessageBody
}

//Connection is the connection created
type Connection struct {
	name     string
	uri      string
	conn     *amqp.Connection
	channel  *amqp.Channel
	exchange string
	queues   []string
	err      chan error
}

//var (
//	connectionPool = make(map[string]*Connection)
//)


////NewConnection returns the new connection object
//func NewConnection(name, uri, exchange string, queues []string) *Connection {
//	if c, ok := connectionPool[name]; ok {
//		return c
//	}
//	c := &Connection{
//		exchange: exchange,
//		uri:      uri,
//		queues:   queues,
//		err:      make(chan error),
//	}
//	connectionPool[name] = c
//	return c
//}

////GetConnection returns the connection which was instantiated
//func GetConnection(name string) *Connection {
//	return connectionPool[name]
//}

func (c *Connection) Connect() error {
	var err error
	c.conn, err = amqp.Dial(c.uri)
	if err != nil {
		return fmt.Errorf("error in creating rabbitmq connection : %s", err.Error())
	}
	go func() {
		<-c.conn.NotifyClose(make(chan *amqp.Error)) //Listen to NotifyClose
		c.err <- errors.New("connection Closed")
	}()
	c.channel, err = c.conn.Channel()
	if err != nil {
		return fmt.Errorf("channel: %s", err)
	}
	if err := c.channel.ExchangeDeclare(
		c.exchange, // name
		"direct",   // type
		true,       // durable
		false,      // auto-deleted
		false,      // internal
		false,      // noWait
		nil,        // arguments
	); err != nil {
		return fmt.Errorf("error in Exchange Declare: %s", err)
	}
	return nil
}

//func (c *Connection) BindQueue() error {
//	for _, q := range c.queues {
//		if _, err := c.channel.QueueDeclare(q, true, false, false, false, nil); err != nil {
//			return fmt.Errorf("error in declaring the queue %s", err)
//		}
//		if err := c.channel.QueueBind(q, q, c.exchange, false, nil); err != nil {
//			return fmt.Errorf("Queue  Bind error: %s", err)
//		}
//	}
//	return nil
//}

////Reconnect reconnects the connection
//func (c *Connection) Reconnect() error {
//	if err := c.Connect(); err != nil {
//		return err
//	}
//	if err := c.BindQueue(); err != nil {
//		return err
//	}
//	return nil
//}

//Publish publishes a request to the amqp queue
func (c *Connection) Publish(m Message) error {
	//select { //non blocking channel - if there is no error will go to default where we do nothing
	//case err := <-c.err:
	//	if err != nil {
	//		c.Reconnect()
	//	}
	//default:
	//}

	p := amqp.Publishing{
		Headers:         amqp.Table{"type": m.Body.Type},
		DeliveryMode:    m.DeliveryMode,
		ContentType:     m.ContentType,
		ContentEncoding: m.ContentEncoding,
		CorrelationId:   m.CorrelationID,
		Body:            m.Body.Data,
		ReplyTo:         m.ReplyTo,
	}
	if err := c.channel.Publish(c.exchange, m.Queue, false, false, p); err != nil {
		return fmt.Errorf("error in Publishing: %s", err)
	}
	return nil
}

//
////Consume consumes the messages from the queues and passes it as map of chan of amqp.Delivery
//func (c *Connection) Consume() (map[string]<-chan amqp.Delivery, error) {
//	m := make(map[string]<-chan amqp.Delivery)
//	for _, q := range c.queues {
//		deliveries, err := c.channel.Consume(q, "", false, false, false, false, nil)
//		if err != nil {
//			return nil, err
//		}
//		m[q] = deliveries
//	}
//	return m, nil
//}
//
////HandleConsumedDeliveries handles the consumed deliveries from the queues. Should be called only for a consumer connection
//func (c *Connection) HandleConsumedDeliveries(q string, delivery <-chan amqp.Delivery, fn func(Connection, string, <-chan amqp.Delivery)) {
//	for {
//		go fn(*c, q, delivery)
//		if err := <-c.err; err != nil {
//			c.Reconnect()
//			deliveries, err := c.Consume()
//			if err != nil {
//				panic(err) //raising panic if consume fails even after reconnecting
//			}
//			delivery = deliveries[q]
//		}
//	}
//}
