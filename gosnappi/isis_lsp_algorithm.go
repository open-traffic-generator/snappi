package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisLspAlgorithm *****
type isisLspAlgorithm struct {
	validation
	obj          *otg.IsisLspAlgorithm
	marshaller   marshalIsisLspAlgorithm
	unMarshaller unMarshalIsisLspAlgorithm
}

func NewIsisLspAlgorithm() IsisLspAlgorithm {
	obj := isisLspAlgorithm{obj: &otg.IsisLspAlgorithm{}}
	obj.setDefault()
	return &obj
}

func (obj *isisLspAlgorithm) msg() *otg.IsisLspAlgorithm {
	return obj.obj
}

func (obj *isisLspAlgorithm) setMsg(msg *otg.IsisLspAlgorithm) IsisLspAlgorithm {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisLspAlgorithm struct {
	obj *isisLspAlgorithm
}

type marshalIsisLspAlgorithm interface {
	// ToProto marshals IsisLspAlgorithm to protobuf object *otg.IsisLspAlgorithm
	ToProto() (*otg.IsisLspAlgorithm, error)
	// ToPbText marshals IsisLspAlgorithm to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisLspAlgorithm to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisLspAlgorithm to JSON text
	ToJson() (string, error)
}

type unMarshalisisLspAlgorithm struct {
	obj *isisLspAlgorithm
}

type unMarshalIsisLspAlgorithm interface {
	// FromProto unmarshals IsisLspAlgorithm from protobuf object *otg.IsisLspAlgorithm
	FromProto(msg *otg.IsisLspAlgorithm) (IsisLspAlgorithm, error)
	// FromPbText unmarshals IsisLspAlgorithm from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisLspAlgorithm from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisLspAlgorithm from JSON text
	FromJson(value string) error
}

func (obj *isisLspAlgorithm) Marshal() marshalIsisLspAlgorithm {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisLspAlgorithm{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisLspAlgorithm) Unmarshal() unMarshalIsisLspAlgorithm {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisLspAlgorithm{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisLspAlgorithm) ToProto() (*otg.IsisLspAlgorithm, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisLspAlgorithm) FromProto(msg *otg.IsisLspAlgorithm) (IsisLspAlgorithm, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisLspAlgorithm) ToPbText() (string, error) {
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

func (m *unMarshalisisLspAlgorithm) FromPbText(value string) error {
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

func (m *marshalisisLspAlgorithm) ToYaml() (string, error) {
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

func (m *unMarshalisisLspAlgorithm) FromYaml(value string) error {
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

func (m *marshalisisLspAlgorithm) ToJson() (string, error) {
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

func (m *unMarshalisisLspAlgorithm) FromJson(value string) error {
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

func (obj *isisLspAlgorithm) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisLspAlgorithm) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisLspAlgorithm) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisLspAlgorithm) Clone() (IsisLspAlgorithm, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisLspAlgorithm()
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

// IsisLspAlgorithm is container for SR-Algorithms. Reference: https://datatracker.ietf.org/doc/html/rfc8667#name-sr-algorithm-sub-tlv.
type IsisLspAlgorithm interface {
	Validation
	// msg marshals IsisLspAlgorithm to protobuf object *otg.IsisLspAlgorithm
	// and doesn't set defaults
	msg() *otg.IsisLspAlgorithm
	// setMsg unmarshals IsisLspAlgorithm from protobuf object *otg.IsisLspAlgorithm
	// and doesn't set defaults
	setMsg(*otg.IsisLspAlgorithm) IsisLspAlgorithm
	// provides marshal interface
	Marshal() marshalIsisLspAlgorithm
	// provides unmarshal interface
	Unmarshal() unMarshalIsisLspAlgorithm
	// validate validates IsisLspAlgorithm
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisLspAlgorithm, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Algorithm returns uint32, set in IsisLspAlgorithm.
	Algorithm() uint32
	// SetAlgorithm assigns uint32 provided by user to IsisLspAlgorithm
	SetAlgorithm(value uint32) IsisLspAlgorithm
	// HasAlgorithm checks if Algorithm has been set in IsisLspAlgorithm
	HasAlgorithm() bool
}

// The Isis may use various algorithms when calculating reachability to other nodes or to prefixes attached to these nodes.
// Algorithm returns a uint32
func (obj *isisLspAlgorithm) Algorithm() uint32 {

	return *obj.obj.Algorithm

}

// The Isis may use various algorithms when calculating reachability to other nodes or to prefixes attached to these nodes.
// Algorithm returns a uint32
func (obj *isisLspAlgorithm) HasAlgorithm() bool {
	return obj.obj.Algorithm != nil
}

// The Isis may use various algorithms when calculating reachability to other nodes or to prefixes attached to these nodes.
// SetAlgorithm sets the uint32 value in the IsisLspAlgorithm object
func (obj *isisLspAlgorithm) SetAlgorithm(value uint32) IsisLspAlgorithm {

	obj.obj.Algorithm = &value
	return obj
}

func (obj *isisLspAlgorithm) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *isisLspAlgorithm) setDefault() {

}
