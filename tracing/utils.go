package tracing

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/segmentio/kafka-go"
	"google.golang.org/grpc/metadata"
)

func StartKafkaConsumerTracerSpan(ctx context.Context, headers []kafka.Header, operationName string) (context.Context, opentracing.Span) {
	carrierFromKafkaHeaders := TextMapCarrierFromKafkaMessageHeaders(headers)

	spanCtx, err := opentracing.GlobalTracer().Extract(opentracing.TextMap, carrierFromKafkaHeaders)
	if err != nil {
		serverSpan := opentracing.GlobalTracer().StartSpan(operationName)
		ctx = opentracing.ContextWithSpan(ctx, serverSpan)
		return ctx, serverSpan
	}

	serverSpan := opentracing.GlobalTracer().StartSpan(operationName, ext.RPCServerOption(spanCtx))
	ctx = opentracing.ContextWithSpan(ctx, serverSpan)

	return ctx, serverSpan
}
func TextMapCarrierFromKafkaMessageHeaders(headers []kafka.Header) opentracing.TextMapCarrier {
	textMap := make(map[string]string, len(headers))
	for _, header := range headers {
		textMap[header.Key] = string(header.Value)
	}
	return opentracing.TextMapCarrier(textMap)
}
func TextMapCarrierToKafkaMessageHeaders(textMap opentracing.TextMapCarrier) []kafka.Header {
	headers := make([]kafka.Header, 0, len(textMap))

	if err := textMap.ForeachKey(func(key, val string) error {
		headers = append(headers, kafka.Header{
			Key:   key,
			Value: []byte(val),
		})
		return nil
	}); err != nil {
		return headers
	}

	return headers
}
func StartHttpServerTracerSpan(c *gin.Context, operationName string) (context.Context, opentracing.Span) {
	spanCtx, err := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
	if err != nil {
		serverSpan := opentracing.GlobalTracer().StartSpan(operationName)
		ctx := opentracing.ContextWithSpan(c.Request.Context(), serverSpan)
		return ctx, serverSpan
	}

	serverSpan := opentracing.GlobalTracer().StartSpan(operationName, ext.RPCServerOption(spanCtx))
	ctx := opentracing.ContextWithSpan(c.Request.Context(), serverSpan)

	return ctx, serverSpan
}

func GetTextMapCarrierFromMetaData(ctx context.Context) opentracing.TextMapCarrier {
	metadataMap := make(opentracing.TextMapCarrier)
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		for key := range md.Copy() {
			metadataMap.Set(key, md.Get(key)[0])
		}
	}
	return metadataMap
}

func StartGrpcServerTracerSpan(ctx context.Context, operationName string) (context.Context, opentracing.Span) {
	textMapCarrierFromMetaData := GetTextMapCarrierFromMetaData(ctx)

	span, err := opentracing.GlobalTracer().Extract(opentracing.TextMap, textMapCarrierFromMetaData)
	if err != nil {
		serverSpan := opentracing.GlobalTracer().StartSpan(operationName)
		ctx = opentracing.ContextWithSpan(ctx, serverSpan)
		return ctx, serverSpan
	}

	serverSpan := opentracing.GlobalTracer().StartSpan(operationName, ext.RPCServerOption(span))
	ctx = opentracing.ContextWithSpan(ctx, serverSpan)

	return ctx, serverSpan
}

func InjectTextMapCarrier(spanCtx opentracing.SpanContext) (opentracing.TextMapCarrier, error) {
	m := make(opentracing.TextMapCarrier)
	if err := opentracing.GlobalTracer().Inject(spanCtx, opentracing.TextMap, m); err != nil {
		return nil, err
	}
	return m, nil
}

func ExtractTextMapCarrier(spanCtx opentracing.SpanContext) opentracing.TextMapCarrier {
	textMapCarrier, err := InjectTextMapCarrier(spanCtx)
	if err != nil {
		return make(opentracing.TextMapCarrier)
	}
	return textMapCarrier
}

func ExtractTextMapCarrierBytes(spanCtx opentracing.SpanContext) []byte {
	textMapCarrier, err := InjectTextMapCarrier(spanCtx)
	if err != nil {
		return []byte("")
	}

	dataBytes, err := json.Marshal(&textMapCarrier)
	if err != nil {
		return []byte("")
	}
	return dataBytes
}

func InjectTextMapCarrierToGrpcMetaData(ctx context.Context, spanCtx opentracing.SpanContext) context.Context {
	if textMapCarrier, err := InjectTextMapCarrier(spanCtx); err == nil {
		md := metadata.New(textMapCarrier)
		ctx = metadata.NewOutgoingContext(ctx, md)
	}
	return ctx
}

func TraceErr(span opentracing.Span, err error) {
	span.SetTag("error", true)
	span.LogKV("error_code", err.Error())
}
