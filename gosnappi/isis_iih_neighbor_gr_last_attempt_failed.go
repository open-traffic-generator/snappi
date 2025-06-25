package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisIIHNeighborGRLastAttemptFailed *****
type isisIIHNeighborGRLastAttemptFailed struct {
	validation
	obj          *otg.IsisIIHNeighborGRLastAttemptFailed
	marshaller   marshalIsisIIHNeighborGRLastAttemptFailed
	unMarshaller unMarshalIsisIIHNeighborGRLastAttemptFailed
}

func NewIsisIIHNeighborGRLastAttemptFailed() IsisIIHNeighborGRLastAttemptFailed {
	obj := isisIIHNeighborGRLastAttemptFailed{obj: &otg.IsisIIHNeighborGRLastAttemptFailed{}}
	obj.setDefault()
	return &obj
}

func (obj *isisIIHNeighborGRLastAttemptFailed) msg() *otg.IsisIIHNeighborGRLastAttemptFailed {
	return obj.obj
}

func (obj *isisIIHNeighborGRLastAttemptFailed) setMsg(msg *otg.IsisIIHNeighborGRLastAttemptFailed) IsisIIHNeighborGRLastAttemptFailed {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisIIHNeighborGRLastAttemptFailed struct {
	obj *isisIIHNeighborGRLastAttemptFailed
}

type marshalIsisIIHNeighborGRLastAttemptFailed interface {
	// ToProto marshals IsisIIHNeighborGRLastAttemptFailed to protobuf object *otg.IsisIIHNeighborGRLastAttemptFailed
	ToProto() (*otg.IsisIIHNeighborGRLastAttemptFailed, error)
	// ToPbText marshals IsisIIHNeighborGRLastAttemptFailed to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisIIHNeighborGRLastAttemptFailed to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisIIHNeighborGRLastAttemptFailed to JSON text
	ToJson() (string, error)
}

type unMarshalisisIIHNeighborGRLastAttemptFailed struct {
	obj *isisIIHNeighborGRLastAttemptFailed
}

type unMarshalIsisIIHNeighborGRLastAttemptFailed interface {
	// FromProto unmarshals IsisIIHNeighborGRLastAttemptFailed from protobuf object *otg.IsisIIHNeighborGRLastAttemptFailed
	FromProto(msg *otg.IsisIIHNeighborGRLastAttemptFailed) (IsisIIHNeighborGRLastAttemptFailed, error)
	// FromPbText unmarshals IsisIIHNeighborGRLastAttemptFailed from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisIIHNeighborGRLastAttemptFailed from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisIIHNeighborGRLastAttemptFailed from JSON text
	FromJson(value string) error
}

func (obj *isisIIHNeighborGRLastAttemptFailed) Marshal() marshalIsisIIHNeighborGRLastAttemptFailed {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisIIHNeighborGRLastAttemptFailed{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisIIHNeighborGRLastAttemptFailed) Unmarshal() unMarshalIsisIIHNeighborGRLastAttemptFailed {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisIIHNeighborGRLastAttemptFailed{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisIIHNeighborGRLastAttemptFailed) ToProto() (*otg.IsisIIHNeighborGRLastAttemptFailed, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisIIHNeighborGRLastAttemptFailed) FromProto(msg *otg.IsisIIHNeighborGRLastAttemptFailed) (IsisIIHNeighborGRLastAttemptFailed, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisIIHNeighborGRLastAttemptFailed) ToPbText() (string, error) {
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

func (m *unMarshalisisIIHNeighborGRLastAttemptFailed) FromPbText(value string) error {
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

func (m *marshalisisIIHNeighborGRLastAttemptFailed) ToYaml() (string, error) {
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

func (m *unMarshalisisIIHNeighborGRLastAttemptFailed) FromYaml(value string) error {
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

func (m *marshalisisIIHNeighborGRLastAttemptFailed) ToJson() (string, error) {
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

func (m *unMarshalisisIIHNeighborGRLastAttemptFailed) FromJson(value string) error {
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

func (obj *isisIIHNeighborGRLastAttemptFailed) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisIIHNeighborGRLastAttemptFailed) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisIIHNeighborGRLastAttemptFailed) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisIIHNeighborGRLastAttemptFailed) Clone() (IsisIIHNeighborGRLastAttemptFailed, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisIIHNeighborGRLastAttemptFailed()
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

// IsisIIHNeighborGRLastAttemptFailed is this container contains the failure status of the last Graceful Restart initiated by this neighbor.
type IsisIIHNeighborGRLastAttemptFailed interface {
	Validation
	// msg marshals IsisIIHNeighborGRLastAttemptFailed to protobuf object *otg.IsisIIHNeighborGRLastAttemptFailed
	// and doesn't set defaults
	msg() *otg.IsisIIHNeighborGRLastAttemptFailed
	// setMsg unmarshals IsisIIHNeighborGRLastAttemptFailed from protobuf object *otg.IsisIIHNeighborGRLastAttemptFailed
	// and doesn't set defaults
	setMsg(*otg.IsisIIHNeighborGRLastAttemptFailed) IsisIIHNeighborGRLastAttemptFailed
	// provides marshal interface
	Marshal() marshalIsisIIHNeighborGRLastAttemptFailed
	// provides unmarshal interface
	Unmarshal() unMarshalIsisIIHNeighborGRLastAttemptFailed
	// validate validates IsisIIHNeighborGRLastAttemptFailed
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisIIHNeighborGRLastAttemptFailed, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Reason returns string, set in IsisIIHNeighborGRLastAttemptFailed.
	Reason() string
	// SetReason assigns string provided by user to IsisIIHNeighborGRLastAttemptFailed
	SetReason(value string) IsisIIHNeighborGRLastAttemptFailed
	// HasReason checks if Reason has been set in IsisIIHNeighborGRLastAttemptFailed
	HasReason() bool
}

// Failure reason of last Graceful Restart in readable string.
// Reason returns a string
func (obj *isisIIHNeighborGRLastAttemptFailed) Reason() string {

	return *obj.obj.Reason

}

// Failure reason of last Graceful Restart in readable string.
// Reason returns a string
func (obj *isisIIHNeighborGRLastAttemptFailed) HasReason() bool {
	return obj.obj.Reason != nil
}

// Failure reason of last Graceful Restart in readable string.
// SetReason sets the string value in the IsisIIHNeighborGRLastAttemptFailed object
func (obj *isisIIHNeighborGRLastAttemptFailed) SetReason(value string) IsisIIHNeighborGRLastAttemptFailed {

	obj.obj.Reason = &value
	return obj
}

func (obj *isisIIHNeighborGRLastAttemptFailed) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *isisIIHNeighborGRLastAttemptFailed) setDefault() {

}
