package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisSRAlgorithm *****
type isisSRAlgorithm struct {
	validation
	obj          *otg.IsisSRAlgorithm
	marshaller   marshalIsisSRAlgorithm
	unMarshaller unMarshalIsisSRAlgorithm
}

func NewIsisSRAlgorithm() IsisSRAlgorithm {
	obj := isisSRAlgorithm{obj: &otg.IsisSRAlgorithm{}}
	obj.setDefault()
	return &obj
}

func (obj *isisSRAlgorithm) msg() *otg.IsisSRAlgorithm {
	return obj.obj
}

func (obj *isisSRAlgorithm) setMsg(msg *otg.IsisSRAlgorithm) IsisSRAlgorithm {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisSRAlgorithm struct {
	obj *isisSRAlgorithm
}

type marshalIsisSRAlgorithm interface {
	// ToProto marshals IsisSRAlgorithm to protobuf object *otg.IsisSRAlgorithm
	ToProto() (*otg.IsisSRAlgorithm, error)
	// ToPbText marshals IsisSRAlgorithm to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisSRAlgorithm to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisSRAlgorithm to JSON text
	ToJson() (string, error)
}

type unMarshalisisSRAlgorithm struct {
	obj *isisSRAlgorithm
}

type unMarshalIsisSRAlgorithm interface {
	// FromProto unmarshals IsisSRAlgorithm from protobuf object *otg.IsisSRAlgorithm
	FromProto(msg *otg.IsisSRAlgorithm) (IsisSRAlgorithm, error)
	// FromPbText unmarshals IsisSRAlgorithm from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisSRAlgorithm from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisSRAlgorithm from JSON text
	FromJson(value string) error
}

func (obj *isisSRAlgorithm) Marshal() marshalIsisSRAlgorithm {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisSRAlgorithm{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisSRAlgorithm) Unmarshal() unMarshalIsisSRAlgorithm {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisSRAlgorithm{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisSRAlgorithm) ToProto() (*otg.IsisSRAlgorithm, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisSRAlgorithm) FromProto(msg *otg.IsisSRAlgorithm) (IsisSRAlgorithm, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisSRAlgorithm) ToPbText() (string, error) {
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

func (m *unMarshalisisSRAlgorithm) FromPbText(value string) error {
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

func (m *marshalisisSRAlgorithm) ToYaml() (string, error) {
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

func (m *unMarshalisisSRAlgorithm) FromYaml(value string) error {
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

func (m *marshalisisSRAlgorithm) ToJson() (string, error) {
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

func (m *unMarshalisisSRAlgorithm) FromJson(value string) error {
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

func (obj *isisSRAlgorithm) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisSRAlgorithm) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisSRAlgorithm) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisSRAlgorithm) Clone() (IsisSRAlgorithm, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisSRAlgorithm()
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

// IsisSRAlgorithm is container for SR-Algorithms. Reference: https://datatracker.ietf.org/doc/html/rfc8667#name-sr-algorithm-sub-tlv.
type IsisSRAlgorithm interface {
	Validation
	// msg marshals IsisSRAlgorithm to protobuf object *otg.IsisSRAlgorithm
	// and doesn't set defaults
	msg() *otg.IsisSRAlgorithm
	// setMsg unmarshals IsisSRAlgorithm from protobuf object *otg.IsisSRAlgorithm
	// and doesn't set defaults
	setMsg(*otg.IsisSRAlgorithm) IsisSRAlgorithm
	// provides marshal interface
	Marshal() marshalIsisSRAlgorithm
	// provides unmarshal interface
	Unmarshal() unMarshalIsisSRAlgorithm
	// validate validates IsisSRAlgorithm
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisSRAlgorithm, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Algorithm returns uint32, set in IsisSRAlgorithm.
	Algorithm() uint32
	// SetAlgorithm assigns uint32 provided by user to IsisSRAlgorithm
	SetAlgorithm(value uint32) IsisSRAlgorithm
	// HasAlgorithm checks if Algorithm has been set in IsisSRAlgorithm
	HasAlgorithm() bool
}

// The Isis may use various algorithms when calculating reachability to other nodes or to prefixes attached to these nodes.
// Algorithm returns a uint32
func (obj *isisSRAlgorithm) Algorithm() uint32 {

	return *obj.obj.Algorithm

}

// The Isis may use various algorithms when calculating reachability to other nodes or to prefixes attached to these nodes.
// Algorithm returns a uint32
func (obj *isisSRAlgorithm) HasAlgorithm() bool {
	return obj.obj.Algorithm != nil
}

// The Isis may use various algorithms when calculating reachability to other nodes or to prefixes attached to these nodes.
// SetAlgorithm sets the uint32 value in the IsisSRAlgorithm object
func (obj *isisSRAlgorithm) SetAlgorithm(value uint32) IsisSRAlgorithm {

	obj.obj.Algorithm = &value
	return obj
}

func (obj *isisSRAlgorithm) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Algorithm != nil {

		if *obj.obj.Algorithm > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= IsisSRAlgorithm.Algorithm <= 255 but Got %d", *obj.obj.Algorithm))
		}

	}

}

func (obj *isisSRAlgorithm) setDefault() {
	if obj.obj.Algorithm == nil {
		obj.SetAlgorithm(0)
	}

}
