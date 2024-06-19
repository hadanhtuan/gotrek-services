package api

import (
	"encoding/json"
	"fmt"
	"search-service/internal/model"
	"search-service/internal/util"

	"github.com/elastic/go-elasticsearch/v8/typedapi/core/update"
	"github.com/hadanhtuan/go-sdk/amqp"
	es "github.com/hadanhtuan/go-sdk/db/elasticsearch"
)

func (pc *SearchController) bindingMap() map[string]amqp.CallbackFunc {
	return map[string]amqp.CallbackFunc{
		util.PropertyCreated: pc.EventPropertyCreated,
		util.PropertyUpdated: pc.EventPropertyUpdated,
	}
}

func (pc *SearchController) InitRoutingAMQP() {
	instance := amqp.GetConnection()
	bindingMap := pc.bindingMap()

	for routingKey, execFunc := range bindingMap {
		instance.BindingQueue(util.SEARCH_EXCHANGE, util.SEARCH_QUEUE, routingKey, execFunc)
	}

	instance.StartConsume(util.SEARCH_QUEUE)
}

func (pc *SearchController) EventPropertyCreated(payload []byte) {
	var data model.Property
	json.Unmarshal(payload, &data)

	es.IndexDocument(util.PropertyIndex, data.ID, data)
}

func (ps *SearchController) EventPropertyUpdated(payload []byte) {
	var property model.Property

	json.Unmarshal(payload, &property)

	fmt.Println(property.ID)
	fmt.Println(property.HostName)

	// byteMerged, _ := json.Marshal(property)

	es.UpdateDocument(util.PropertyIndex, property.ID, &update.Request{
		Doc: json.RawMessage(payload),
	})
}
