package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisIIHNeighborGRLastAttemptStatus *****
type isisIIHNeighborGRLastAttemptStatus struct {
	validation
	obj             *otg.IsisIIHNeighborGRLastAttemptStatus
	marshaller      marshalIsisIIHNeighborGRLastAttemptStatus
	unMarshaller    unMarshalIsisIIHNeighborGRLastAttemptStatus
	succeededHolder IsisIIHNeighborGRLastAttemptSucceeded
	failedHolder    IsisIIHNeighborGRLastAttemptFailed
}

func NewIsisIIHNeighborGRLastAttemptStatus() IsisIIHNeighborGRLastAttemptStatus {
	obj := isisIIHNeighborGRLastAttemptStatus{obj: &otg.IsisIIHNeighborGRLastAttemptStatus{}}
	obj.setDefault()
	return &obj
}

func (obj *isisIIHNeighborGRLastAttemptStatus) msg() *otg.IsisIIHNeighborGRLastAttemptStatus {
	return obj.obj
}

func (obj *isisIIHNeighborGRLastAttemptStatus) setMsg(msg *otg.IsisIIHNeighborGRLastAttemptStatus) IsisIIHNeighborGRLastAttemptStatus {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisIIHNeighborGRLastAttemptStatus struct {
	obj *isisIIHNeighborGRLastAttemptStatus
}

type marshalIsisIIHNeighborGRLastAttemptStatus interface {
	// ToProto marshals IsisIIHNeighborGRLastAttemptStatus to protobuf object *otg.IsisIIHNeighborGRLastAttemptStatus
	ToProto() (*otg.IsisIIHNeighborGRLastAttemptStatus, error)
	// ToPbText marshals IsisIIHNeighborGRLastAttemptStatus to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisIIHNeighborGRLastAttemptStatus to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisIIHNeighborGRLastAttemptStatus to JSON text
	ToJson() (string, error)
}

type unMarshalisisIIHNeighborGRLastAttemptStatus struct {
	obj *isisIIHNeighborGRLastAttemptStatus
}

type unMarshalIsisIIHNeighborGRLastAttemptStatus interface {
	// FromProto unmarshals IsisIIHNeighborGRLastAttemptStatus from protobuf object *otg.IsisIIHNeighborGRLastAttemptStatus
	FromProto(msg *otg.IsisIIHNeighborGRLastAttemptStatus) (IsisIIHNeighborGRLastAttemptStatus, error)
	// FromPbText unmarshals IsisIIHNeighborGRLastAttemptStatus from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisIIHNeighborGRLastAttemptStatus from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisIIHNeighborGRLastAttemptStatus from JSON text
	FromJson(value string) error
}

func (obj *isisIIHNeighborGRLastAttemptStatus) Marshal() marshalIsisIIHNeighborGRLastAttemptStatus {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisIIHNeighborGRLastAttemptStatus{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisIIHNeighborGRLastAttemptStatus) Unmarshal() unMarshalIsisIIHNeighborGRLastAttemptStatus {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisIIHNeighborGRLastAttemptStatus{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisIIHNeighborGRLastAttemptStatus) ToProto() (*otg.IsisIIHNeighborGRLastAttemptStatus, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisIIHNeighborGRLastAttemptStatus) FromProto(msg *otg.IsisIIHNeighborGRLastAttemptStatus) (IsisIIHNeighborGRLastAttemptStatus, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisIIHNeighborGRLastAttemptStatus) ToPbText() (string, error) {
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

func (m *unMarshalisisIIHNeighborGRLastAttemptStatus) FromPbText(value string) error {
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

func (m *marshalisisIIHNeighborGRLastAttemptStatus) ToYaml() (string, error) {
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

func (m *unMarshalisisIIHNeighborGRLastAttemptStatus) FromYaml(value string) error {
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

func (m *marshalisisIIHNeighborGRLastAttemptStatus) ToJson() (string, error) {
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

func (m *unMarshalisisIIHNeighborGRLastAttemptStatus) FromJson(value string) error {
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

func (obj *isisIIHNeighborGRLastAttemptStatus) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisIIHNeighborGRLastAttemptStatus) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisIIHNeighborGRLastAttemptStatus) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisIIHNeighborGRLastAttemptStatus) Clone() (IsisIIHNeighborGRLastAttemptStatus, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisIIHNeighborGRLastAttemptStatus()
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

func (obj *isisIIHNeighborGRLastAttemptStatus) setNil() {
	obj.succeededHolder = nil
	obj.failedHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// IsisIIHNeighborGRLastAttemptStatus is this object contains the status of the last attempted Graceful Restart status of a ISIS neighbor.
// - succeeded: Choice is set if the last Graceful Status is successful.
// - failed: The last Graceful Status is unsuccessful.
// - inprogress: The last Graceful Restart status is in progress.
// - unavailable: The last Graceful Restart status is not initiated.
type IsisIIHNeighborGRLastAttemptStatus interface {
	Validation
	// msg marshals IsisIIHNeighborGRLastAttemptStatus to protobuf object *otg.IsisIIHNeighborGRLastAttemptStatus
	// and doesn't set defaults
	msg() *otg.IsisIIHNeighborGRLastAttemptStatus
	// setMsg unmarshals IsisIIHNeighborGRLastAttemptStatus from protobuf object *otg.IsisIIHNeighborGRLastAttemptStatus
	// and doesn't set defaults
	setMsg(*otg.IsisIIHNeighborGRLastAttemptStatus) IsisIIHNeighborGRLastAttemptStatus
	// provides marshal interface
	Marshal() marshalIsisIIHNeighborGRLastAttemptStatus
	// provides unmarshal interface
	Unmarshal() unMarshalIsisIIHNeighborGRLastAttemptStatus
	// validate validates IsisIIHNeighborGRLastAttemptStatus
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisIIHNeighborGRLastAttemptStatus, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns IsisIIHNeighborGRLastAttemptStatusChoiceEnum, set in IsisIIHNeighborGRLastAttemptStatus
	Choice() IsisIIHNeighborGRLastAttemptStatusChoiceEnum
	// setChoice assigns IsisIIHNeighborGRLastAttemptStatusChoiceEnum provided by user to IsisIIHNeighborGRLastAttemptStatus
	setChoice(value IsisIIHNeighborGRLastAttemptStatusChoiceEnum) IsisIIHNeighborGRLastAttemptStatus
	// HasChoice checks if Choice has been set in IsisIIHNeighborGRLastAttemptStatus
	HasChoice() bool
	// Succeeded returns IsisIIHNeighborGRLastAttemptSucceeded, set in IsisIIHNeighborGRLastAttemptStatus.
	// IsisIIHNeighborGRLastAttemptSucceeded is this object contains the result of a successful graceful restart status in the last attempted by a Neighbor.
	Succeeded() IsisIIHNeighborGRLastAttemptSucceeded
	// SetSucceeded assigns IsisIIHNeighborGRLastAttemptSucceeded provided by user to IsisIIHNeighborGRLastAttemptStatus.
	// IsisIIHNeighborGRLastAttemptSucceeded is this object contains the result of a successful graceful restart status in the last attempted by a Neighbor.
	SetSucceeded(value IsisIIHNeighborGRLastAttemptSucceeded) IsisIIHNeighborGRLastAttemptStatus
	// HasSucceeded checks if Succeeded has been set in IsisIIHNeighborGRLastAttemptStatus
	HasSucceeded() bool
	// Failed returns IsisIIHNeighborGRLastAttemptFailed, set in IsisIIHNeighborGRLastAttemptStatus.
	// IsisIIHNeighborGRLastAttemptFailed is this container contains the failure status of the last Graceful Restart initiated by this neighbor.
	Failed() IsisIIHNeighborGRLastAttemptFailed
	// SetFailed assigns IsisIIHNeighborGRLastAttemptFailed provided by user to IsisIIHNeighborGRLastAttemptStatus.
	// IsisIIHNeighborGRLastAttemptFailed is this container contains the failure status of the last Graceful Restart initiated by this neighbor.
	SetFailed(value IsisIIHNeighborGRLastAttemptFailed) IsisIIHNeighborGRLastAttemptStatus
	// HasFailed checks if Failed has been set in IsisIIHNeighborGRLastAttemptStatus
	HasFailed() bool
	// Inprogress returns bool, set in IsisIIHNeighborGRLastAttemptStatus.
	Inprogress() bool
	// SetInprogress assigns bool provided by user to IsisIIHNeighborGRLastAttemptStatus
	SetInprogress(value bool) IsisIIHNeighborGRLastAttemptStatus
	// HasInprogress checks if Inprogress has been set in IsisIIHNeighborGRLastAttemptStatus
	HasInprogress() bool
	// Unavailable returns bool, set in IsisIIHNeighborGRLastAttemptStatus.
	Unavailable() bool
	// SetUnavailable assigns bool provided by user to IsisIIHNeighborGRLastAttemptStatus
	SetUnavailable(value bool) IsisIIHNeighborGRLastAttemptStatus
	// HasUnavailable checks if Unavailable has been set in IsisIIHNeighborGRLastAttemptStatus
	HasUnavailable() bool
	setNil()
}

type IsisIIHNeighborGRLastAttemptStatusChoiceEnum string

// Enum of Choice on IsisIIHNeighborGRLastAttemptStatus
var IsisIIHNeighborGRLastAttemptStatusChoice = struct {
	SUCCEEDED   IsisIIHNeighborGRLastAttemptStatusChoiceEnum
	FAILED      IsisIIHNeighborGRLastAttemptStatusChoiceEnum
	INPROGRESS  IsisIIHNeighborGRLastAttemptStatusChoiceEnum
	UNAVAILABLE IsisIIHNeighborGRLastAttemptStatusChoiceEnum
}{
	SUCCEEDED:   IsisIIHNeighborGRLastAttemptStatusChoiceEnum("succeeded"),
	FAILED:      IsisIIHNeighborGRLastAttemptStatusChoiceEnum("failed"),
	INPROGRESS:  IsisIIHNeighborGRLastAttemptStatusChoiceEnum("inprogress"),
	UNAVAILABLE: IsisIIHNeighborGRLastAttemptStatusChoiceEnum("unavailable"),
}

func (obj *isisIIHNeighborGRLastAttemptStatus) Choice() IsisIIHNeighborGRLastAttemptStatusChoiceEnum {
	return IsisIIHNeighborGRLastAttemptStatusChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *isisIIHNeighborGRLastAttemptStatus) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *isisIIHNeighborGRLastAttemptStatus) setChoice(value IsisIIHNeighborGRLastAttemptStatusChoiceEnum) IsisIIHNeighborGRLastAttemptStatus {
	intValue, ok := otg.IsisIIHNeighborGRLastAttemptStatus_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on IsisIIHNeighborGRLastAttemptStatusChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.IsisIIHNeighborGRLastAttemptStatus_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Unavailable = nil
	obj.obj.Inprogress = nil
	obj.obj.Failed = nil
	obj.failedHolder = nil
	obj.obj.Succeeded = nil
	obj.succeededHolder = nil

	if value == IsisIIHNeighborGRLastAttemptStatusChoice.SUCCEEDED {
		obj.obj.Succeeded = NewIsisIIHNeighborGRLastAttemptSucceeded().msg()
	}

	if value == IsisIIHNeighborGRLastAttemptStatusChoice.FAILED {
		obj.obj.Failed = NewIsisIIHNeighborGRLastAttemptFailed().msg()
	}

	return obj
}

// description is TBD
// Succeeded returns a IsisIIHNeighborGRLastAttemptSucceeded
func (obj *isisIIHNeighborGRLastAttemptStatus) Succeeded() IsisIIHNeighborGRLastAttemptSucceeded {
	if obj.obj.Succeeded == nil {
		obj.setChoice(IsisIIHNeighborGRLastAttemptStatusChoice.SUCCEEDED)
	}
	if obj.succeededHolder == nil {
		obj.succeededHolder = &isisIIHNeighborGRLastAttemptSucceeded{obj: obj.obj.Succeeded}
	}
	return obj.succeededHolder
}

// description is TBD
// Succeeded returns a IsisIIHNeighborGRLastAttemptSucceeded
func (obj *isisIIHNeighborGRLastAttemptStatus) HasSucceeded() bool {
	return obj.obj.Succeeded != nil
}

// description is TBD
// SetSucceeded sets the IsisIIHNeighborGRLastAttemptSucceeded value in the IsisIIHNeighborGRLastAttemptStatus object
func (obj *isisIIHNeighborGRLastAttemptStatus) SetSucceeded(value IsisIIHNeighborGRLastAttemptSucceeded) IsisIIHNeighborGRLastAttemptStatus {
	obj.setChoice(IsisIIHNeighborGRLastAttemptStatusChoice.SUCCEEDED)
	obj.succeededHolder = nil
	obj.obj.Succeeded = value.msg()

	return obj
}

// description is TBD
// Failed returns a IsisIIHNeighborGRLastAttemptFailed
func (obj *isisIIHNeighborGRLastAttemptStatus) Failed() IsisIIHNeighborGRLastAttemptFailed {
	if obj.obj.Failed == nil {
		obj.setChoice(IsisIIHNeighborGRLastAttemptStatusChoice.FAILED)
	}
	if obj.failedHolder == nil {
		obj.failedHolder = &isisIIHNeighborGRLastAttemptFailed{obj: obj.obj.Failed}
	}
	return obj.failedHolder
}

// description is TBD
// Failed returns a IsisIIHNeighborGRLastAttemptFailed
func (obj *isisIIHNeighborGRLastAttemptStatus) HasFailed() bool {
	return obj.obj.Failed != nil
}

// description is TBD
// SetFailed sets the IsisIIHNeighborGRLastAttemptFailed value in the IsisIIHNeighborGRLastAttemptStatus object
func (obj *isisIIHNeighborGRLastAttemptStatus) SetFailed(value IsisIIHNeighborGRLastAttemptFailed) IsisIIHNeighborGRLastAttemptStatus {
	obj.setChoice(IsisIIHNeighborGRLastAttemptStatusChoice.FAILED)
	obj.failedHolder = nil
	obj.obj.Failed = value.msg()

	return obj
}

// description is TBD
// Inprogress returns a bool
func (obj *isisIIHNeighborGRLastAttemptStatus) Inprogress() bool {

	if obj.obj.Inprogress == nil {
		obj.setChoice(IsisIIHNeighborGRLastAttemptStatusChoice.INPROGRESS)
	}

	return *obj.obj.Inprogress

}

// description is TBD
// Inprogress returns a bool
func (obj *isisIIHNeighborGRLastAttemptStatus) HasInprogress() bool {
	return obj.obj.Inprogress != nil
}

// description is TBD
// SetInprogress sets the bool value in the IsisIIHNeighborGRLastAttemptStatus object
func (obj *isisIIHNeighborGRLastAttemptStatus) SetInprogress(value bool) IsisIIHNeighborGRLastAttemptStatus {
	obj.setChoice(IsisIIHNeighborGRLastAttemptStatusChoice.INPROGRESS)
	obj.obj.Inprogress = &value
	return obj
}

// description is TBD
// Unavailable returns a bool
func (obj *isisIIHNeighborGRLastAttemptStatus) Unavailable() bool {

	if obj.obj.Unavailable == nil {
		obj.setChoice(IsisIIHNeighborGRLastAttemptStatusChoice.UNAVAILABLE)
	}

	return *obj.obj.Unavailable

}

// description is TBD
// Unavailable returns a bool
func (obj *isisIIHNeighborGRLastAttemptStatus) HasUnavailable() bool {
	return obj.obj.Unavailable != nil
}

// description is TBD
// SetUnavailable sets the bool value in the IsisIIHNeighborGRLastAttemptStatus object
func (obj *isisIIHNeighborGRLastAttemptStatus) SetUnavailable(value bool) IsisIIHNeighborGRLastAttemptStatus {
	obj.setChoice(IsisIIHNeighborGRLastAttemptStatusChoice.UNAVAILABLE)
	obj.obj.Unavailable = &value
	return obj
}

func (obj *isisIIHNeighborGRLastAttemptStatus) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Succeeded != nil {

		obj.Succeeded().validateObj(vObj, set_default)
	}

	if obj.obj.Failed != nil {

		obj.Failed().validateObj(vObj, set_default)
	}

}

func (obj *isisIIHNeighborGRLastAttemptStatus) setDefault() {
	var choices_set int = 0
	var choice IsisIIHNeighborGRLastAttemptStatusChoiceEnum

	if obj.obj.Succeeded != nil {
		choices_set += 1
		choice = IsisIIHNeighborGRLastAttemptStatusChoice.SUCCEEDED
	}

	if obj.obj.Failed != nil {
		choices_set += 1
		choice = IsisIIHNeighborGRLastAttemptStatusChoice.FAILED
	}

	if obj.obj.Inprogress != nil {
		choices_set += 1
		choice = IsisIIHNeighborGRLastAttemptStatusChoice.INPROGRESS
	}

	if obj.obj.Unavailable != nil {
		choices_set += 1
		choice = IsisIIHNeighborGRLastAttemptStatusChoice.UNAVAILABLE
	}
	if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in IsisIIHNeighborGRLastAttemptStatus")
			}
		} else {
			intVal := otg.IsisIIHNeighborGRLastAttemptStatus_Choice_Enum_value[string(choice)]
			enumValue := otg.IsisIIHNeighborGRLastAttemptStatus_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
