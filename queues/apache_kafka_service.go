package queues

import (
	"github.com/sivaosorg/govm/apix"
)

type KafkaService interface {
}

type kafkaServiceImpl struct {
	conf []apix.ApiRequestConfig
}

func NewKafkaService(conf []apix.ApiRequestConfig) KafkaService {
	return &kafkaServiceImpl{
		conf: conf,
	}
}

func NewKafkaServiceSlices(conf ...apix.ApiRequestConfig) KafkaService {
	return &kafkaServiceImpl{
		conf: conf,
	}
}

// func (s *kafkaServiceImpl) ProduceCallback(request KafkaPublisherRequest) (*restify.Response, error) {
// err := KafkaPublisherRequestValidator(request)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if len(s.conf) == 0 {
// 		return nil, fmt.Errorf("API Conf is required")
// 	}
// 	configs, ok := apix.Get(s.conf, request.TenantKey)
// 	if !ok {
// 		return nil, fmt.Errorf("Tenant Key undefined: %v", request.TenantKey)
// 	}
// 	endpoint, err := configs.GetEndpoint(request.TopicKey)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if len(request.Payload) > 0 {
// 		for k, v := range request.Payload {
// 			endpoint.AppendBodyWith(k, v)
// 		}
// 	}
// 	svc := apix.NewApiService(configs)
// 	return svc.Do(nil, endpoint)
// }

// func (s *kafkaServiceImpl) ProduceCallbackNoneWait(request KafkaPublisherRequest) {
// 	go s.ProduceCallback(request)
// }
