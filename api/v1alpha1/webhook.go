package v1alpha1

import (
	"github.com/muvaf/kubebuilder-init/api/v1beta2"
	"github.com/pkg/errors"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

func (e *ExampleKind) ConvertTo(dstRaw conversion.Hub) error {
	dst, ok := dstRaw.(*v1beta2.ExampleKind)
	if !ok {
		return errors.New("non-hub version is sent to ConvertTo of v1alpha1")
	}
	dst.Spec.FooBeta2 = e.Spec.Foo
	return nil
}

func (e *ExampleKind) ConvertFrom(srcRaw conversion.Hub) error {
	src, ok := srcRaw.(*v1beta2.ExampleKind)
	if !ok {
		return errors.New("non-hub version is sent to ConvertFrom of v1alpha1")
	}
	e.Spec.Foo = src.Spec.FooBeta2
	return nil
}
