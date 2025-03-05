package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** RoCEv2PerPortSettings *****
type roCEv2PerPortSettings struct {
	validation
	obj              *otg.RoCEv2PerPortSettings
	marshaller       marshalRoCEv2PerPortSettings
	unMarshaller     unMarshalRoCEv2PerPortSettings
	interBatchHolder RoCEv2InterBatch
}

func NewRoCEv2PerPortSettings() RoCEv2PerPortSettings {
	obj := roCEv2PerPortSettings{obj: &otg.RoCEv2PerPortSettings{}}
	obj.setDefault()
	return &obj
}

func (obj *roCEv2PerPortSettings) msg() *otg.RoCEv2PerPortSettings {
	return obj.obj
}

func (obj *roCEv2PerPortSettings) setMsg(msg *otg.RoCEv2PerPortSettings) RoCEv2PerPortSettings {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalroCEv2PerPortSettings struct {
	obj *roCEv2PerPortSettings
}

type marshalRoCEv2PerPortSettings interface {
	// ToProto marshals RoCEv2PerPortSettings to protobuf object *otg.RoCEv2PerPortSettings
	ToProto() (*otg.RoCEv2PerPortSettings, error)
	// ToPbText marshals RoCEv2PerPortSettings to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals RoCEv2PerPortSettings to YAML text
	ToYaml() (string, error)
	// ToJson marshals RoCEv2PerPortSettings to JSON text
	ToJson() (string, error)
}

type unMarshalroCEv2PerPortSettings struct {
	obj *roCEv2PerPortSettings
}

type unMarshalRoCEv2PerPortSettings interface {
	// FromProto unmarshals RoCEv2PerPortSettings from protobuf object *otg.RoCEv2PerPortSettings
	FromProto(msg *otg.RoCEv2PerPortSettings) (RoCEv2PerPortSettings, error)
	// FromPbText unmarshals RoCEv2PerPortSettings from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals RoCEv2PerPortSettings from YAML text
	FromYaml(value string) error
	// FromJson unmarshals RoCEv2PerPortSettings from JSON text
	FromJson(value string) error
}

func (obj *roCEv2PerPortSettings) Marshal() marshalRoCEv2PerPortSettings {
	if obj.marshaller == nil {
		obj.marshaller = &marshalroCEv2PerPortSettings{obj: obj}
	}
	return obj.marshaller
}

func (obj *roCEv2PerPortSettings) Unmarshal() unMarshalRoCEv2PerPortSettings {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalroCEv2PerPortSettings{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalroCEv2PerPortSettings) ToProto() (*otg.RoCEv2PerPortSettings, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalroCEv2PerPortSettings) FromProto(msg *otg.RoCEv2PerPortSettings) (RoCEv2PerPortSettings, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalroCEv2PerPortSettings) ToPbText() (string, error) {
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

func (m *unMarshalroCEv2PerPortSettings) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), m.obj.msg())
	if retObj != nil {
		return retObj
	}
	m.obj.setNil()
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (m *marshalroCEv2PerPortSettings) ToYaml() (string, error) {
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

func (m *unMarshalroCEv2PerPortSettings) FromYaml(value string) error {
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
	m.obj.setNil()
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (m *marshalroCEv2PerPortSettings) ToJson() (string, error) {
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

func (m *unMarshalroCEv2PerPortSettings) FromJson(value string) error {
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
	m.obj.setNil()
	err := m.obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *roCEv2PerPortSettings) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *roCEv2PerPortSettings) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *roCEv2PerPortSettings) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *roCEv2PerPortSettings) Clone() (RoCEv2PerPortSettings, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRoCEv2PerPortSettings()
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

func (obj *roCEv2PerPortSettings) setNil() {
	obj.interBatchHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// RoCEv2PerPortSettings is roCEv2 traffic flow configuration.
type RoCEv2PerPortSettings interface {
	Validation
	// msg marshals RoCEv2PerPortSettings to protobuf object *otg.RoCEv2PerPortSettings
	// and doesn't set defaults
	msg() *otg.RoCEv2PerPortSettings
	// setMsg unmarshals RoCEv2PerPortSettings from protobuf object *otg.RoCEv2PerPortSettings
	// and doesn't set defaults
	setMsg(*otg.RoCEv2PerPortSettings) RoCEv2PerPortSettings
	// provides marshal interface
	Marshal() marshalRoCEv2PerPortSettings
	// provides unmarshal interface
	Unmarshal() unMarshalRoCEv2PerPortSettings
	// validate validates RoCEv2PerPortSettings
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (RoCEv2PerPortSettings, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// TargetLineRate returns uint64, set in RoCEv2PerPortSettings.
	TargetLineRate() uint64
	// SetTargetLineRate assigns uint64 provided by user to RoCEv2PerPortSettings
	SetTargetLineRate(value uint64) RoCEv2PerPortSettings
	// HasTargetLineRate checks if TargetLineRate has been set in RoCEv2PerPortSettings
	HasTargetLineRate() bool
	// InterBatch returns RoCEv2InterBatch, set in RoCEv2PerPortSettings.
	// RoCEv2InterBatch is roCEv2 inter batch settigns applicable per port
	InterBatch() RoCEv2InterBatch
	// SetInterBatch assigns RoCEv2InterBatch provided by user to RoCEv2PerPortSettings.
	// RoCEv2InterBatch is roCEv2 inter batch settigns applicable per port
	SetInterBatch(value RoCEv2InterBatch) RoCEv2PerPortSettings
	// HasInterBatch checks if InterBatch has been set in RoCEv2PerPortSettings
	HasInterBatch() bool
	setNil()
}

// Target Line Rate in percentage.
// TargetLineRate returns a uint64
func (obj *roCEv2PerPortSettings) TargetLineRate() uint64 {

	return *obj.obj.TargetLineRate

}

// Target Line Rate in percentage.
// TargetLineRate returns a uint64
func (obj *roCEv2PerPortSettings) HasTargetLineRate() bool {
	return obj.obj.TargetLineRate != nil
}

// Target Line Rate in percentage.
// SetTargetLineRate sets the uint64 value in the RoCEv2PerPortSettings object
func (obj *roCEv2PerPortSettings) SetTargetLineRate(value uint64) RoCEv2PerPortSettings {

	obj.obj.TargetLineRate = &value
	return obj
}

// description is TBD
// InterBatch returns a RoCEv2InterBatch
func (obj *roCEv2PerPortSettings) InterBatch() RoCEv2InterBatch {
	if obj.obj.InterBatch == nil {
		obj.obj.InterBatch = NewRoCEv2InterBatch().msg()
	}
	if obj.interBatchHolder == nil {
		obj.interBatchHolder = &roCEv2InterBatch{obj: obj.obj.InterBatch}
	}
	return obj.interBatchHolder
}

// description is TBD
// InterBatch returns a RoCEv2InterBatch
func (obj *roCEv2PerPortSettings) HasInterBatch() bool {
	return obj.obj.InterBatch != nil
}

// description is TBD
// SetInterBatch sets the RoCEv2InterBatch value in the RoCEv2PerPortSettings object
func (obj *roCEv2PerPortSettings) SetInterBatch(value RoCEv2InterBatch) RoCEv2PerPortSettings {

	obj.interBatchHolder = nil
	obj.obj.InterBatch = value.msg()

	return obj
}

func (obj *roCEv2PerPortSettings) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.InterBatch != nil {

		obj.InterBatch().validateObj(vObj, set_default)
	}

}

func (obj *roCEv2PerPortSettings) setDefault() {

}
