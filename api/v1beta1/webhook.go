package v1beta1

import (
	"github.com/muvaf/kubebuilder-init/api/v1beta2"
	"github.com/pkg/errors"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

func (e *ExampleKind) ConvertTo(dstRaw conversion.Hub) error {
	dst, ok := dstRaw.(*v1beta2.ExampleKind)
	if !ok {
		return errors.New("non-hub version is sent to ConvertTo of v1beta1")
	}
	dst.Spec.FooBeta2 = e.Spec.FooBeta1
	return nil
}

func (e *ExampleKind) ConvertFrom(srcRaw conversion.Hub) error {
	src, ok := srcRaw.(*v1beta2.ExampleKind)
	if !ok {
		return errors.New("non-hub version is sent to ConvertFrom of v1beta1")
	}
	e.Spec.FooBeta1 = src.Spec.FooBeta2
	return nil
}