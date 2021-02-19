package v1beta1

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
		return errors.New("non-hub version is sent to ConvertTo of v1beta1")
	}
	fmt.Println("convert from v1beta1 to HUB")
	dstYaml, err := yaml.Marshal(e)
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
	dst.Spec.FooBeta2 = e.Spec.FooBeta1
	dst.ObjectMeta = e.ObjectMeta
	return nil
}

func (e *ExampleKind) ConvertFrom(srcRaw conversion.Hub) error {
	src, ok := srcRaw.(*v1beta2.ExampleKind)
	if !ok {
		return errors.New("non-hub version is sent to ConvertFrom of v1beta1")
	}
	fmt.Println("convert from HUB to v1beta1")
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
	e.Spec.FooBeta1 = src.Spec.FooBeta2
	e.ObjectMeta = src.ObjectMeta
	return nil
}
