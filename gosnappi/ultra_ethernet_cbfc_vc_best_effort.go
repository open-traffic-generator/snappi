package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** UltraEthernetCbfcVcBestEffort *****
type ultraEthernetCbfcVcBestEffort struct {
	validation
	obj          *otg.UltraEthernetCbfcVcBestEffort
	marshaller   marshalUltraEthernetCbfcVcBestEffort
	unMarshaller unMarshalUltraEthernetCbfcVcBestEffort
}

func NewUltraEthernetCbfcVcBestEffort() UltraEthernetCbfcVcBestEffort {
	obj := ultraEthernetCbfcVcBestEffort{obj: &otg.UltraEthernetCbfcVcBestEffort{}}
	obj.setDefault()
	return &obj
}

func (obj *ultraEthernetCbfcVcBestEffort) msg() *otg.UltraEthernetCbfcVcBestEffort {
	return obj.obj
}

func (obj *ultraEthernetCbfcVcBestEffort) setMsg(msg *otg.UltraEthernetCbfcVcBestEffort) UltraEthernetCbfcVcBestEffort {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalultraEthernetCbfcVcBestEffort struct {
	obj *ultraEthernetCbfcVcBestEffort
}

type marshalUltraEthernetCbfcVcBestEffort interface {
	// ToProto marshals UltraEthernetCbfcVcBestEffort to protobuf object *otg.UltraEthernetCbfcVcBestEffort
	ToProto() (*otg.UltraEthernetCbfcVcBestEffort, error)
	// ToPbText marshals UltraEthernetCbfcVcBestEffort to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals UltraEthernetCbfcVcBestEffort to YAML text
	ToYaml() (string, error)
	// ToJson marshals UltraEthernetCbfcVcBestEffort to JSON text
	ToJson() (string, error)
}

type unMarshalultraEthernetCbfcVcBestEffort struct {
	obj *ultraEthernetCbfcVcBestEffort
}

type unMarshalUltraEthernetCbfcVcBestEffort interface {
	// FromProto unmarshals UltraEthernetCbfcVcBestEffort from protobuf object *otg.UltraEthernetCbfcVcBestEffort
	FromProto(msg *otg.UltraEthernetCbfcVcBestEffort) (UltraEthernetCbfcVcBestEffort, error)
	// FromPbText unmarshals UltraEthernetCbfcVcBestEffort from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals UltraEthernetCbfcVcBestEffort from YAML text
	FromYaml(value string) error
	// FromJson unmarshals UltraEthernetCbfcVcBestEffort from JSON text
	FromJson(value string) error
}

func (obj *ultraEthernetCbfcVcBestEffort) Marshal() marshalUltraEthernetCbfcVcBestEffort {
	if obj.marshaller == nil {
		obj.marshaller = &marshalultraEthernetCbfcVcBestEffort{obj: obj}
	}
	return obj.marshaller
}

func (obj *ultraEthernetCbfcVcBestEffort) Unmarshal() unMarshalUltraEthernetCbfcVcBestEffort {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalultraEthernetCbfcVcBestEffort{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalultraEthernetCbfcVcBestEffort) ToProto() (*otg.UltraEthernetCbfcVcBestEffort, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalultraEthernetCbfcVcBestEffort) FromProto(msg *otg.UltraEthernetCbfcVcBestEffort) (UltraEthernetCbfcVcBestEffort, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalultraEthernetCbfcVcBestEffort) ToPbText() (string, error) {
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

func (m *unMarshalultraEthernetCbfcVcBestEffort) FromPbText(value string) error {
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

func (m *marshalultraEthernetCbfcVcBestEffort) ToYaml() (string, error) {
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

func (m *unMarshalultraEthernetCbfcVcBestEffort) FromYaml(value string) error {
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

func (m *marshalultraEthernetCbfcVcBestEffort) ToJson() (string, error) {
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

func (m *unMarshalultraEthernetCbfcVcBestEffort) FromJson(value string) error {
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

func (obj *ultraEthernetCbfcVcBestEffort) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ultraEthernetCbfcVcBestEffort) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ultraEthernetCbfcVcBestEffort) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ultraEthernetCbfcVcBestEffort) Clone() (UltraEthernetCbfcVcBestEffort, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewUltraEthernetCbfcVcBestEffort()
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

// UltraEthernetCbfcVcBestEffort is best effort virtual channel settings. A best effort VC does not use CBFC
// credits.
type UltraEthernetCbfcVcBestEffort interface {
	Validation
	// msg marshals UltraEthernetCbfcVcBestEffort to protobuf object *otg.UltraEthernetCbfcVcBestEffort
	// and doesn't set defaults
	msg() *otg.UltraEthernetCbfcVcBestEffort
	// setMsg unmarshals UltraEthernetCbfcVcBestEffort from protobuf object *otg.UltraEthernetCbfcVcBestEffort
	// and doesn't set defaults
	setMsg(*otg.UltraEthernetCbfcVcBestEffort) UltraEthernetCbfcVcBestEffort
	// provides marshal interface
	Marshal() marshalUltraEthernetCbfcVcBestEffort
	// provides unmarshal interface
	Unmarshal() unMarshalUltraEthernetCbfcVcBestEffort
	// validate validates UltraEthernetCbfcVcBestEffort
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (UltraEthernetCbfcVcBestEffort, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
}

func (obj *ultraEthernetCbfcVcBestEffort) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *ultraEthernetCbfcVcBestEffort) setDefault() {

}
