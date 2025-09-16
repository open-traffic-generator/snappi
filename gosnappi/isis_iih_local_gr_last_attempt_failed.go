package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisIIHLocalGRLastAttemptFailed *****
type isisIIHLocalGRLastAttemptFailed struct {
	validation
	obj          *otg.IsisIIHLocalGRLastAttemptFailed
	marshaller   marshalIsisIIHLocalGRLastAttemptFailed
	unMarshaller unMarshalIsisIIHLocalGRLastAttemptFailed
}

func NewIsisIIHLocalGRLastAttemptFailed() IsisIIHLocalGRLastAttemptFailed {
	obj := isisIIHLocalGRLastAttemptFailed{obj: &otg.IsisIIHLocalGRLastAttemptFailed{}}
	obj.setDefault()
	return &obj
}

func (obj *isisIIHLocalGRLastAttemptFailed) msg() *otg.IsisIIHLocalGRLastAttemptFailed {
	return obj.obj
}

func (obj *isisIIHLocalGRLastAttemptFailed) setMsg(msg *otg.IsisIIHLocalGRLastAttemptFailed) IsisIIHLocalGRLastAttemptFailed {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisIIHLocalGRLastAttemptFailed struct {
	obj *isisIIHLocalGRLastAttemptFailed
}

type marshalIsisIIHLocalGRLastAttemptFailed interface {
	// ToProto marshals IsisIIHLocalGRLastAttemptFailed to protobuf object *otg.IsisIIHLocalGRLastAttemptFailed
	ToProto() (*otg.IsisIIHLocalGRLastAttemptFailed, error)
	// ToPbText marshals IsisIIHLocalGRLastAttemptFailed to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisIIHLocalGRLastAttemptFailed to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisIIHLocalGRLastAttemptFailed to JSON text
	ToJson() (string, error)
}

type unMarshalisisIIHLocalGRLastAttemptFailed struct {
	obj *isisIIHLocalGRLastAttemptFailed
}

type unMarshalIsisIIHLocalGRLastAttemptFailed interface {
	// FromProto unmarshals IsisIIHLocalGRLastAttemptFailed from protobuf object *otg.IsisIIHLocalGRLastAttemptFailed
	FromProto(msg *otg.IsisIIHLocalGRLastAttemptFailed) (IsisIIHLocalGRLastAttemptFailed, error)
	// FromPbText unmarshals IsisIIHLocalGRLastAttemptFailed from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisIIHLocalGRLastAttemptFailed from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisIIHLocalGRLastAttemptFailed from JSON text
	FromJson(value string) error
}

func (obj *isisIIHLocalGRLastAttemptFailed) Marshal() marshalIsisIIHLocalGRLastAttemptFailed {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisIIHLocalGRLastAttemptFailed{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisIIHLocalGRLastAttemptFailed) Unmarshal() unMarshalIsisIIHLocalGRLastAttemptFailed {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisIIHLocalGRLastAttemptFailed{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisIIHLocalGRLastAttemptFailed) ToProto() (*otg.IsisIIHLocalGRLastAttemptFailed, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisIIHLocalGRLastAttemptFailed) FromProto(msg *otg.IsisIIHLocalGRLastAttemptFailed) (IsisIIHLocalGRLastAttemptFailed, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisIIHLocalGRLastAttemptFailed) ToPbText() (string, error) {
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

func (m *unMarshalisisIIHLocalGRLastAttemptFailed) FromPbText(value string) error {
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

func (m *marshalisisIIHLocalGRLastAttemptFailed) ToYaml() (string, error) {
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

func (m *unMarshalisisIIHLocalGRLastAttemptFailed) FromYaml(value string) error {
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

func (m *marshalisisIIHLocalGRLastAttemptFailed) ToJson() (string, error) {
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

func (m *unMarshalisisIIHLocalGRLastAttemptFailed) FromJson(value string) error {
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

func (obj *isisIIHLocalGRLastAttemptFailed) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisIIHLocalGRLastAttemptFailed) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisIIHLocalGRLastAttemptFailed) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisIIHLocalGRLastAttemptFailed) Clone() (IsisIIHLocalGRLastAttemptFailed, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisIIHLocalGRLastAttemptFailed()
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

// IsisIIHLocalGRLastAttemptFailed is this container contains the failure status of the last Graceful Restart initiated by this router.
type IsisIIHLocalGRLastAttemptFailed interface {
	Validation
	// msg marshals IsisIIHLocalGRLastAttemptFailed to protobuf object *otg.IsisIIHLocalGRLastAttemptFailed
	// and doesn't set defaults
	msg() *otg.IsisIIHLocalGRLastAttemptFailed
	// setMsg unmarshals IsisIIHLocalGRLastAttemptFailed from protobuf object *otg.IsisIIHLocalGRLastAttemptFailed
	// and doesn't set defaults
	setMsg(*otg.IsisIIHLocalGRLastAttemptFailed) IsisIIHLocalGRLastAttemptFailed
	// provides marshal interface
	Marshal() marshalIsisIIHLocalGRLastAttemptFailed
	// provides unmarshal interface
	Unmarshal() unMarshalIsisIIHLocalGRLastAttemptFailed
	// validate validates IsisIIHLocalGRLastAttemptFailed
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisIIHLocalGRLastAttemptFailed, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Reason returns string, set in IsisIIHLocalGRLastAttemptFailed.
	Reason() string
	// SetReason assigns string provided by user to IsisIIHLocalGRLastAttemptFailed
	SetReason(value string) IsisIIHLocalGRLastAttemptFailed
	// HasReason checks if Reason has been set in IsisIIHLocalGRLastAttemptFailed
	HasReason() bool
}

// Failure reason of last Graceful Restart.
// Reason returns a string
func (obj *isisIIHLocalGRLastAttemptFailed) Reason() string {

	return *obj.obj.Reason

}

// Failure reason of last Graceful Restart.
// Reason returns a string
func (obj *isisIIHLocalGRLastAttemptFailed) HasReason() bool {
	return obj.obj.Reason != nil
}

// Failure reason of last Graceful Restart.
// SetReason sets the string value in the IsisIIHLocalGRLastAttemptFailed object
func (obj *isisIIHLocalGRLastAttemptFailed) SetReason(value string) IsisIIHLocalGRLastAttemptFailed {

	obj.obj.Reason = &value
	return obj
}

func (obj *isisIIHLocalGRLastAttemptFailed) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *isisIIHLocalGRLastAttemptFailed) setDefault() {

}
