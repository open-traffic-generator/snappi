package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisLspFlags *****
type isisLspFlags struct {
	validation
	obj          *otg.IsisLspFlags
	marshaller   marshalIsisLspFlags
	unMarshaller unMarshalIsisLspFlags
}

func NewIsisLspFlags() IsisLspFlags {
	obj := isisLspFlags{obj: &otg.IsisLspFlags{}}
	obj.setDefault()
	return &obj
}

func (obj *isisLspFlags) msg() *otg.IsisLspFlags {
	return obj.obj
}

func (obj *isisLspFlags) setMsg(msg *otg.IsisLspFlags) IsisLspFlags {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisLspFlags struct {
	obj *isisLspFlags
}

type marshalIsisLspFlags interface {
	// ToProto marshals IsisLspFlags to protobuf object *otg.IsisLspFlags
	ToProto() (*otg.IsisLspFlags, error)
	// ToPbText marshals IsisLspFlags to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisLspFlags to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisLspFlags to JSON text
	ToJson() (string, error)
}

type unMarshalisisLspFlags struct {
	obj *isisLspFlags
}

type unMarshalIsisLspFlags interface {
	// FromProto unmarshals IsisLspFlags from protobuf object *otg.IsisLspFlags
	FromProto(msg *otg.IsisLspFlags) (IsisLspFlags, error)
	// FromPbText unmarshals IsisLspFlags from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisLspFlags from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisLspFlags from JSON text
	FromJson(value string) error
}

func (obj *isisLspFlags) Marshal() marshalIsisLspFlags {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisLspFlags{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisLspFlags) Unmarshal() unMarshalIsisLspFlags {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisLspFlags{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisLspFlags) ToProto() (*otg.IsisLspFlags, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisLspFlags) FromProto(msg *otg.IsisLspFlags) (IsisLspFlags, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisLspFlags) ToPbText() (string, error) {
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

func (m *unMarshalisisLspFlags) FromPbText(value string) error {
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

func (m *marshalisisLspFlags) ToYaml() (string, error) {
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

func (m *unMarshalisisLspFlags) FromYaml(value string) error {
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

func (m *marshalisisLspFlags) ToJson() (string, error) {
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

func (m *unMarshalisisLspFlags) FromJson(value string) error {
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

func (obj *isisLspFlags) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisLspFlags) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisLspFlags) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisLspFlags) Clone() (IsisLspFlags, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisLspFlags()
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

// IsisLspFlags is lSP Type flags.
type IsisLspFlags interface {
	Validation
	// msg marshals IsisLspFlags to protobuf object *otg.IsisLspFlags
	// and doesn't set defaults
	msg() *otg.IsisLspFlags
	// setMsg unmarshals IsisLspFlags from protobuf object *otg.IsisLspFlags
	// and doesn't set defaults
	setMsg(*otg.IsisLspFlags) IsisLspFlags
	// provides marshal interface
	Marshal() marshalIsisLspFlags
	// provides unmarshal interface
	Unmarshal() unMarshalIsisLspFlags
	// validate validates IsisLspFlags
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisLspFlags, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// PartitionRepair returns bool, set in IsisLspFlags.
	PartitionRepair() bool
	// SetPartitionRepair assigns bool provided by user to IsisLspFlags
	SetPartitionRepair(value bool) IsisLspFlags
	// HasPartitionRepair checks if PartitionRepair has been set in IsisLspFlags
	HasPartitionRepair() bool
	// AttachedError returns bool, set in IsisLspFlags.
	AttachedError() bool
	// SetAttachedError assigns bool provided by user to IsisLspFlags
	SetAttachedError(value bool) IsisLspFlags
	// HasAttachedError checks if AttachedError has been set in IsisLspFlags
	HasAttachedError() bool
	// AttachedExpense returns bool, set in IsisLspFlags.
	AttachedExpense() bool
	// SetAttachedExpense assigns bool provided by user to IsisLspFlags
	SetAttachedExpense(value bool) IsisLspFlags
	// HasAttachedExpense checks if AttachedExpense has been set in IsisLspFlags
	HasAttachedExpense() bool
	// AttachedDelay returns bool, set in IsisLspFlags.
	AttachedDelay() bool
	// SetAttachedDelay assigns bool provided by user to IsisLspFlags
	SetAttachedDelay(value bool) IsisLspFlags
	// HasAttachedDelay checks if AttachedDelay has been set in IsisLspFlags
	HasAttachedDelay() bool
	// AttachedDefault returns bool, set in IsisLspFlags.
	AttachedDefault() bool
	// SetAttachedDefault assigns bool provided by user to IsisLspFlags
	SetAttachedDefault(value bool) IsisLspFlags
	// HasAttachedDefault checks if AttachedDefault has been set in IsisLspFlags
	HasAttachedDefault() bool
	// Overload returns bool, set in IsisLspFlags.
	Overload() bool
	// SetOverload assigns bool provided by user to IsisLspFlags
	SetOverload(value bool) IsisLspFlags
	// HasOverload checks if Overload has been set in IsisLspFlags
	HasOverload() bool
}

// When set, the originator supports partition repair.
// PartitionRepair returns a bool
func (obj *isisLspFlags) PartitionRepair() bool {

	return *obj.obj.PartitionRepair

}

// When set, the originator supports partition repair.
// PartitionRepair returns a bool
func (obj *isisLspFlags) HasPartitionRepair() bool {
	return obj.obj.PartitionRepair != nil
}

// When set, the originator supports partition repair.
// SetPartitionRepair sets the bool value in the IsisLspFlags object
func (obj *isisLspFlags) SetPartitionRepair(value bool) IsisLspFlags {

	obj.obj.PartitionRepair = &value
	return obj
}

// When set, the originator is attached to another area using the referred metric.
// AttachedError returns a bool
func (obj *isisLspFlags) AttachedError() bool {

	return *obj.obj.AttachedError

}

// When set, the originator is attached to another area using the referred metric.
// AttachedError returns a bool
func (obj *isisLspFlags) HasAttachedError() bool {
	return obj.obj.AttachedError != nil
}

// When set, the originator is attached to another area using the referred metric.
// SetAttachedError sets the bool value in the IsisLspFlags object
func (obj *isisLspFlags) SetAttachedError(value bool) IsisLspFlags {

	obj.obj.AttachedError = &value
	return obj
}

// When set, the originator is attached to another
// area using the referred metric.
// AttachedExpense returns a bool
func (obj *isisLspFlags) AttachedExpense() bool {

	return *obj.obj.AttachedExpense

}

// When set, the originator is attached to another
// area using the referred metric.
// AttachedExpense returns a bool
func (obj *isisLspFlags) HasAttachedExpense() bool {
	return obj.obj.AttachedExpense != nil
}

// When set, the originator is attached to another
// area using the referred metric.
// SetAttachedExpense sets the bool value in the IsisLspFlags object
func (obj *isisLspFlags) SetAttachedExpense(value bool) IsisLspFlags {

	obj.obj.AttachedExpense = &value
	return obj
}

// Delay Metric - when set, the originator is attached to another
// area using the referred metric.
// AttachedDelay returns a bool
func (obj *isisLspFlags) AttachedDelay() bool {

	return *obj.obj.AttachedDelay

}

// Delay Metric - when set, the originator is attached to another
// area using the referred metric.
// AttachedDelay returns a bool
func (obj *isisLspFlags) HasAttachedDelay() bool {
	return obj.obj.AttachedDelay != nil
}

// Delay Metric - when set, the originator is attached to another
// area using the referred metric.
// SetAttachedDelay sets the bool value in the IsisLspFlags object
func (obj *isisLspFlags) SetAttachedDelay(value bool) IsisLspFlags {

	obj.obj.AttachedDelay = &value
	return obj
}

// Default Metric - when set, the originator is attached to another
// area using the referred metric.
// AttachedDefault returns a bool
func (obj *isisLspFlags) AttachedDefault() bool {

	return *obj.obj.AttachedDefault

}

// Default Metric - when set, the originator is attached to another
// area using the referred metric.
// AttachedDefault returns a bool
func (obj *isisLspFlags) HasAttachedDefault() bool {
	return obj.obj.AttachedDefault != nil
}

// Default Metric - when set, the originator is attached to another
// area using the referred metric.
// SetAttachedDefault sets the bool value in the IsisLspFlags object
func (obj *isisLspFlags) SetAttachedDefault(value bool) IsisLspFlags {

	obj.obj.AttachedDefault = &value
	return obj
}

// Overload bit - when set, the originator is overloaded, and must
// be avoided in path calculation.
// Overload returns a bool
func (obj *isisLspFlags) Overload() bool {

	return *obj.obj.Overload

}

// Overload bit - when set, the originator is overloaded, and must
// be avoided in path calculation.
// Overload returns a bool
func (obj *isisLspFlags) HasOverload() bool {
	return obj.obj.Overload != nil
}

// Overload bit - when set, the originator is overloaded, and must
// be avoided in path calculation.
// SetOverload sets the bool value in the IsisLspFlags object
func (obj *isisLspFlags) SetOverload(value bool) IsisLspFlags {

	obj.obj.Overload = &value
	return obj
}

func (obj *isisLspFlags) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *isisLspFlags) setDefault() {

}
