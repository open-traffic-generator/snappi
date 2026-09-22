package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** UltraEthernetCbfcPortCreditPerVc *****
type ultraEthernetCbfcPortCreditPerVc struct {
	validation
	obj          *otg.UltraEthernetCbfcPortCreditPerVc
	marshaller   marshalUltraEthernetCbfcPortCreditPerVc
	unMarshaller unMarshalUltraEthernetCbfcPortCreditPerVc
}

func NewUltraEthernetCbfcPortCreditPerVc() UltraEthernetCbfcPortCreditPerVc {
	obj := ultraEthernetCbfcPortCreditPerVc{obj: &otg.UltraEthernetCbfcPortCreditPerVc{}}
	obj.setDefault()
	return &obj
}

func (obj *ultraEthernetCbfcPortCreditPerVc) msg() *otg.UltraEthernetCbfcPortCreditPerVc {
	return obj.obj
}

func (obj *ultraEthernetCbfcPortCreditPerVc) setMsg(msg *otg.UltraEthernetCbfcPortCreditPerVc) UltraEthernetCbfcPortCreditPerVc {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalultraEthernetCbfcPortCreditPerVc struct {
	obj *ultraEthernetCbfcPortCreditPerVc
}

type marshalUltraEthernetCbfcPortCreditPerVc interface {
	// ToProto marshals UltraEthernetCbfcPortCreditPerVc to protobuf object *otg.UltraEthernetCbfcPortCreditPerVc
	ToProto() (*otg.UltraEthernetCbfcPortCreditPerVc, error)
	// ToPbText marshals UltraEthernetCbfcPortCreditPerVc to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals UltraEthernetCbfcPortCreditPerVc to YAML text
	ToYaml() (string, error)
	// ToJson marshals UltraEthernetCbfcPortCreditPerVc to JSON text
	ToJson() (string, error)
}

type unMarshalultraEthernetCbfcPortCreditPerVc struct {
	obj *ultraEthernetCbfcPortCreditPerVc
}

type unMarshalUltraEthernetCbfcPortCreditPerVc interface {
	// FromProto unmarshals UltraEthernetCbfcPortCreditPerVc from protobuf object *otg.UltraEthernetCbfcPortCreditPerVc
	FromProto(msg *otg.UltraEthernetCbfcPortCreditPerVc) (UltraEthernetCbfcPortCreditPerVc, error)
	// FromPbText unmarshals UltraEthernetCbfcPortCreditPerVc from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals UltraEthernetCbfcPortCreditPerVc from YAML text
	FromYaml(value string) error
	// FromJson unmarshals UltraEthernetCbfcPortCreditPerVc from JSON text
	FromJson(value string) error
}

func (obj *ultraEthernetCbfcPortCreditPerVc) Marshal() marshalUltraEthernetCbfcPortCreditPerVc {
	if obj.marshaller == nil {
		obj.marshaller = &marshalultraEthernetCbfcPortCreditPerVc{obj: obj}
	}
	return obj.marshaller
}

func (obj *ultraEthernetCbfcPortCreditPerVc) Unmarshal() unMarshalUltraEthernetCbfcPortCreditPerVc {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalultraEthernetCbfcPortCreditPerVc{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalultraEthernetCbfcPortCreditPerVc) ToProto() (*otg.UltraEthernetCbfcPortCreditPerVc, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalultraEthernetCbfcPortCreditPerVc) FromProto(msg *otg.UltraEthernetCbfcPortCreditPerVc) (UltraEthernetCbfcPortCreditPerVc, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalultraEthernetCbfcPortCreditPerVc) ToPbText() (string, error) {
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	protoMarshal, err := proto.Marshal(m.obj.msg())
	if err != nil {
		return "", err
	}
	return string(protoMarshal), nil
}

func (m *unMarshalultraEthernetCbfcPortCreditPerVc) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), m.obj.msg())
	if retObj != nil {
		return retObj
	}

	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (m *marshalultraEthernetCbfcPortCreditPerVc) ToYaml() (string, error) {
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(m.obj.msg())
	if err != nil {
		return "", err
	}
	data, err = yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (m *unMarshalultraEthernetCbfcPortCreditPerVc) FromYaml(value string) error {
	if value == "" {
		value = "{}"
	}
	data, err := yaml.YAMLToJSON([]byte(value))
	if err != nil {
		return err
	}
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	uError := opts.Unmarshal([]byte(data), m.obj.msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (m *marshalultraEthernetCbfcPortCreditPerVc) ToJson() (string, error) {
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, err := opts.Marshal(m.obj.msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (m *unMarshalultraEthernetCbfcPortCreditPerVc) FromJson(value string) error {
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	if value == "" {
		value = "{}"
	}
	uError := opts.Unmarshal([]byte(value), m.obj.msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	err := m.obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *ultraEthernetCbfcPortCreditPerVc) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ultraEthernetCbfcPortCreditPerVc) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ultraEthernetCbfcPortCreditPerVc) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ultraEthernetCbfcPortCreditPerVc) Clone() (UltraEthernetCbfcPortCreditPerVc, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewUltraEthernetCbfcPortCreditPerVc()
	data, err := proto.Marshal(obj.msg())
	if err != nil {
		return nil, err
	}
	pbErr := proto.Unmarshal(data, newObj.msg())
	if pbErr != nil {
		return nil, pbErr
	}
	return newObj, nil
}

// UltraEthernetCbfcPortCreditPerVc is per virtual channel credit limit configuration. The credit limit for each
// lossless VC is set individually using the credit_limit field of the VC.
type UltraEthernetCbfcPortCreditPerVc interface {
	Validation
	// msg marshals UltraEthernetCbfcPortCreditPerVc to protobuf object *otg.UltraEthernetCbfcPortCreditPerVc
	// and doesn't set defaults
	msg() *otg.UltraEthernetCbfcPortCreditPerVc
	// setMsg unmarshals UltraEthernetCbfcPortCreditPerVc from protobuf object *otg.UltraEthernetCbfcPortCreditPerVc
	// and doesn't set defaults
	setMsg(*otg.UltraEthernetCbfcPortCreditPerVc) UltraEthernetCbfcPortCreditPerVc
	// provides marshal interface
	Marshal() marshalUltraEthernetCbfcPortCreditPerVc
	// provides unmarshal interface
	Unmarshal() unMarshalUltraEthernetCbfcPortCreditPerVc
	// validate validates UltraEthernetCbfcPortCreditPerVc
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (UltraEthernetCbfcPortCreditPerVc, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
}

func (obj *ultraEthernetCbfcPortCreditPerVc) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *ultraEthernetCbfcPortCreditPerVc) setDefault() {

}
