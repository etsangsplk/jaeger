package zapdb

import (
	"flag"
    "fmt"

	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/jaegertracing/jaeger/storage/dependencystore"
	"github.com/jaegertracing/jaeger/storage/spanstore"
	"github.com/uber/jaeger-lib/metrics"
)

// Factory implements storage.Factory for zapdb backend.
type Factory struct {
	Options *Options

	metricsFactory metrics.Factory
	logger         *zap.Logger
}

// NewFactory creates a new Factory.
func NewFactory() *Factory {
    fmt.Println("ZABDB STORAGE NewFactory!!!")
	return &Factory{}
}

// AddFlags implements plugin.Configurable
func (f *Factory) AddFlags(flagSet *flag.FlagSet) {
    fmt.Println("ZABDB STORAGE AddFlags!!!")
	return
}

// InitFromViper implements plugin.Configurable
func (f *Factory) InitFromViper(v *viper.Viper) {
    fmt.Println("ZABDB STORAGE InitFromViper!!!")
	return
}

// Initialize implements storage.Factory
func (f *Factory) Initialize(metricsFactory metrics.Factory, logger *zap.Logger) error {
    fmt.Println("ZABDB STORAGE Initialize!!!")
	return nil
}

// CreateSpanReader implements storage.Factory
func (f *Factory) CreateSpanReader() (spanstore.Reader, error) {
    fmt.Println("ZABDB STORAGE CreateSpanReader!!!")
	return nil, nil
}

// CreateSpanWriter implements storage.Factory
func (f *Factory) CreateSpanWriter() (spanstore.Writer, error) {
    fmt.Println("ZABDB STORAGE CreateSpanWriter!!!")
	return nil, nil
}

// CreateDependencyReader implements storage.Factory
func (f *Factory) CreateDependencyReader() (dependencystore.Reader, error) {
    fmt.Println("ZABDB STORAGE CreateDependencyReader!!!")
	return nil, nil
}

// CreateArchiveSpanReader implements storage.ArchiveFactory
func (f *Factory) CreateArchiveSpanReader() (spanstore.Reader, error) {
    fmt.Println("ZABDB STORAGE CreateArchiveSpanReader!!!")
	return nil, nil
}

// CreateArchiveSpanWriter implements storage.ArchiveFactory
func (f *Factory) CreateArchiveSpanWriter() (spanstore.Writer, error) {
    fmt.Println("ZABDB STORAGE CreateArchiveSpanWriter!!!")
	return nil, nil
}
