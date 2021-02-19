package v1alpha1

import (
	"fmt"
	"github.com/muvaf/kubebuilder-init/api/v1beta2"
	"github.com/pkg/errors"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
	"sigs.k8s.io/yaml"
)

func (e *ExampleKind) ConvertTo(dstRaw conversion.Hub) error {
	dst, ok := dstRaw.(*v1beta2.ExampleKind)
	if !ok {
		return errors.New("non-hub version is sent to ConvertTo of v1alpha1")
	}
	dstYaml, err := yaml.Marshal(dst)
	if err != nil {
		return err
	}
	srcYaml, err := yaml.Marshal(e)
	if err != nil {
		return err
	}

	fmt.Println("dst yaml:")
	fmt.Println(string(dstYaml))
	fmt.Println("src yaml:")
	fmt.Println(string(srcYaml))
	fmt.Println("convert from v1alpha1 to HUB")
	dst.Spec.FooBeta2 = e.Spec.Foo
	dst.ObjectMeta = e.ObjectMeta

	return nil
}

func (e *ExampleKind) ConvertFrom(srcRaw conversion.Hub) error {
	src, ok := srcRaw.(*v1beta2.ExampleKind)
	if !ok {
		return errors.New("non-hub version is sent to ConvertFrom of v1alpha1")
	}
	fmt.Println("convert from HUB to v1alpha1")
	dstYaml, err := yaml.Marshal(e)
	if err != nil {
		return err
	}
	srcYaml, err := yaml.Marshal(src)
	if err != nil {
		return err
	}

	fmt.Println("dst yaml:")
	fmt.Println(string(dstYaml))
	fmt.Println("src yaml:")
	fmt.Println(string(srcYaml))
	e.Spec.Foo = src.Spec.FooBeta2
	e.ObjectMeta = src.ObjectMeta

	return nil
}
