package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisIIHLocalGRLastAttemptStatus *****
type isisIIHLocalGRLastAttemptStatus struct {
	validation
	obj             *otg.IsisIIHLocalGRLastAttemptStatus
	marshaller      marshalIsisIIHLocalGRLastAttemptStatus
	unMarshaller    unMarshalIsisIIHLocalGRLastAttemptStatus
	succeededHolder IsisIIHLocalGRLastAttemptSucceeded
	failedHolder    IsisIIHLocalGRLastAttemptFailed
}

func NewIsisIIHLocalGRLastAttemptStatus() IsisIIHLocalGRLastAttemptStatus {
	obj := isisIIHLocalGRLastAttemptStatus{obj: &otg.IsisIIHLocalGRLastAttemptStatus{}}
	obj.setDefault()
	return &obj
}

func (obj *isisIIHLocalGRLastAttemptStatus) msg() *otg.IsisIIHLocalGRLastAttemptStatus {
	return obj.obj
}

func (obj *isisIIHLocalGRLastAttemptStatus) setMsg(msg *otg.IsisIIHLocalGRLastAttemptStatus) IsisIIHLocalGRLastAttemptStatus {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisIIHLocalGRLastAttemptStatus struct {
	obj *isisIIHLocalGRLastAttemptStatus
}

type marshalIsisIIHLocalGRLastAttemptStatus interface {
	// ToProto marshals IsisIIHLocalGRLastAttemptStatus to protobuf object *otg.IsisIIHLocalGRLastAttemptStatus
	ToProto() (*otg.IsisIIHLocalGRLastAttemptStatus, error)
	// ToPbText marshals IsisIIHLocalGRLastAttemptStatus to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisIIHLocalGRLastAttemptStatus to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisIIHLocalGRLastAttemptStatus to JSON text
	ToJson() (string, error)
}

type unMarshalisisIIHLocalGRLastAttemptStatus struct {
	obj *isisIIHLocalGRLastAttemptStatus
}

type unMarshalIsisIIHLocalGRLastAttemptStatus interface {
	// FromProto unmarshals IsisIIHLocalGRLastAttemptStatus from protobuf object *otg.IsisIIHLocalGRLastAttemptStatus
	FromProto(msg *otg.IsisIIHLocalGRLastAttemptStatus) (IsisIIHLocalGRLastAttemptStatus, error)
	// FromPbText unmarshals IsisIIHLocalGRLastAttemptStatus from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisIIHLocalGRLastAttemptStatus from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisIIHLocalGRLastAttemptStatus from JSON text
	FromJson(value string) error
}

func (obj *isisIIHLocalGRLastAttemptStatus) Marshal() marshalIsisIIHLocalGRLastAttemptStatus {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisIIHLocalGRLastAttemptStatus{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisIIHLocalGRLastAttemptStatus) Unmarshal() unMarshalIsisIIHLocalGRLastAttemptStatus {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisIIHLocalGRLastAttemptStatus{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisIIHLocalGRLastAttemptStatus) ToProto() (*otg.IsisIIHLocalGRLastAttemptStatus, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisIIHLocalGRLastAttemptStatus) FromProto(msg *otg.IsisIIHLocalGRLastAttemptStatus) (IsisIIHLocalGRLastAttemptStatus, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisIIHLocalGRLastAttemptStatus) ToPbText() (string, error) {
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

func (m *unMarshalisisIIHLocalGRLastAttemptStatus) FromPbText(value string) error {
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

func (m *marshalisisIIHLocalGRLastAttemptStatus) ToYaml() (string, error) {
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

func (m *unMarshalisisIIHLocalGRLastAttemptStatus) FromYaml(value string) error {
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

func (m *marshalisisIIHLocalGRLastAttemptStatus) ToJson() (string, error) {
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

func (m *unMarshalisisIIHLocalGRLastAttemptStatus) FromJson(value string) error {
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

func (obj *isisIIHLocalGRLastAttemptStatus) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisIIHLocalGRLastAttemptStatus) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisIIHLocalGRLastAttemptStatus) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisIIHLocalGRLastAttemptStatus) Clone() (IsisIIHLocalGRLastAttemptStatus, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisIIHLocalGRLastAttemptStatus()
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

func (obj *isisIIHLocalGRLastAttemptStatus) setNil() {
	obj.succeededHolder = nil
	obj.failedHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// IsisIIHLocalGRLastAttemptStatus is this object contains the status of the last attempted Graceful Restart status of this router.
// - succeeded: Choice is set if the last Graceful Status is successful.
// - failed: The last Graceful Status is unsuccessful.
// - inprogress: The last Graceful Restart status is in progress.
// - unavailable: The last Graceful Restart status is not initiated.
type IsisIIHLocalGRLastAttemptStatus interface {
	Validation
	// msg marshals IsisIIHLocalGRLastAttemptStatus to protobuf object *otg.IsisIIHLocalGRLastAttemptStatus
	// and doesn't set defaults
	msg() *otg.IsisIIHLocalGRLastAttemptStatus
	// setMsg unmarshals IsisIIHLocalGRLastAttemptStatus from protobuf object *otg.IsisIIHLocalGRLastAttemptStatus
	// and doesn't set defaults
	setMsg(*otg.IsisIIHLocalGRLastAttemptStatus) IsisIIHLocalGRLastAttemptStatus
	// provides marshal interface
	Marshal() marshalIsisIIHLocalGRLastAttemptStatus
	// provides unmarshal interface
	Unmarshal() unMarshalIsisIIHLocalGRLastAttemptStatus
	// validate validates IsisIIHLocalGRLastAttemptStatus
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisIIHLocalGRLastAttemptStatus, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns IsisIIHLocalGRLastAttemptStatusChoiceEnum, set in IsisIIHLocalGRLastAttemptStatus
	Choice() IsisIIHLocalGRLastAttemptStatusChoiceEnum
	// setChoice assigns IsisIIHLocalGRLastAttemptStatusChoiceEnum provided by user to IsisIIHLocalGRLastAttemptStatus
	setChoice(value IsisIIHLocalGRLastAttemptStatusChoiceEnum) IsisIIHLocalGRLastAttemptStatus
	// HasChoice checks if Choice has been set in IsisIIHLocalGRLastAttemptStatus
	HasChoice() bool
	// Succeeded returns IsisIIHLocalGRLastAttemptSucceeded, set in IsisIIHLocalGRLastAttemptStatus.
	// IsisIIHLocalGRLastAttemptSucceeded is this container contains details about the successful status of the last Graceful Restart initiated by this router.
	Succeeded() IsisIIHLocalGRLastAttemptSucceeded
	// SetSucceeded assigns IsisIIHLocalGRLastAttemptSucceeded provided by user to IsisIIHLocalGRLastAttemptStatus.
	// IsisIIHLocalGRLastAttemptSucceeded is this container contains details about the successful status of the last Graceful Restart initiated by this router.
	SetSucceeded(value IsisIIHLocalGRLastAttemptSucceeded) IsisIIHLocalGRLastAttemptStatus
	// HasSucceeded checks if Succeeded has been set in IsisIIHLocalGRLastAttemptStatus
	HasSucceeded() bool
	// Failed returns IsisIIHLocalGRLastAttemptFailed, set in IsisIIHLocalGRLastAttemptStatus.
	// IsisIIHLocalGRLastAttemptFailed is this container contains the failure status of the last Graceful Restart initiated by this router.
	Failed() IsisIIHLocalGRLastAttemptFailed
	// SetFailed assigns IsisIIHLocalGRLastAttemptFailed provided by user to IsisIIHLocalGRLastAttemptStatus.
	// IsisIIHLocalGRLastAttemptFailed is this container contains the failure status of the last Graceful Restart initiated by this router.
	SetFailed(value IsisIIHLocalGRLastAttemptFailed) IsisIIHLocalGRLastAttemptStatus
	// HasFailed checks if Failed has been set in IsisIIHLocalGRLastAttemptStatus
	HasFailed() bool
	// Inprogress returns bool, set in IsisIIHLocalGRLastAttemptStatus.
	Inprogress() bool
	// SetInprogress assigns bool provided by user to IsisIIHLocalGRLastAttemptStatus
	SetInprogress(value bool) IsisIIHLocalGRLastAttemptStatus
	// HasInprogress checks if Inprogress has been set in IsisIIHLocalGRLastAttemptStatus
	HasInprogress() bool
	// Unavailable returns bool, set in IsisIIHLocalGRLastAttemptStatus.
	Unavailable() bool
	// SetUnavailable assigns bool provided by user to IsisIIHLocalGRLastAttemptStatus
	SetUnavailable(value bool) IsisIIHLocalGRLastAttemptStatus
	// HasUnavailable checks if Unavailable has been set in IsisIIHLocalGRLastAttemptStatus
	HasUnavailable() bool
	setNil()
}

type IsisIIHLocalGRLastAttemptStatusChoiceEnum string

// Enum of Choice on IsisIIHLocalGRLastAttemptStatus
var IsisIIHLocalGRLastAttemptStatusChoice = struct {
	SUCCEEDED   IsisIIHLocalGRLastAttemptStatusChoiceEnum
	FAILED      IsisIIHLocalGRLastAttemptStatusChoiceEnum
	INPROGRESS  IsisIIHLocalGRLastAttemptStatusChoiceEnum
	UNAVAILABLE IsisIIHLocalGRLastAttemptStatusChoiceEnum
}{
	SUCCEEDED:   IsisIIHLocalGRLastAttemptStatusChoiceEnum("succeeded"),
	FAILED:      IsisIIHLocalGRLastAttemptStatusChoiceEnum("failed"),
	INPROGRESS:  IsisIIHLocalGRLastAttemptStatusChoiceEnum("inprogress"),
	UNAVAILABLE: IsisIIHLocalGRLastAttemptStatusChoiceEnum("unavailable"),
}

func (obj *isisIIHLocalGRLastAttemptStatus) Choice() IsisIIHLocalGRLastAttemptStatusChoiceEnum {
	return IsisIIHLocalGRLastAttemptStatusChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *isisIIHLocalGRLastAttemptStatus) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *isisIIHLocalGRLastAttemptStatus) setChoice(value IsisIIHLocalGRLastAttemptStatusChoiceEnum) IsisIIHLocalGRLastAttemptStatus {
	intValue, ok := otg.IsisIIHLocalGRLastAttemptStatus_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on IsisIIHLocalGRLastAttemptStatusChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.IsisIIHLocalGRLastAttemptStatus_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Unavailable = nil
	obj.obj.Inprogress = nil
	obj.obj.Failed = nil
	obj.failedHolder = nil
	obj.obj.Succeeded = nil
	obj.succeededHolder = nil

	if value == IsisIIHLocalGRLastAttemptStatusChoice.SUCCEEDED {
		obj.obj.Succeeded = NewIsisIIHLocalGRLastAttemptSucceeded().msg()
	}

	if value == IsisIIHLocalGRLastAttemptStatusChoice.FAILED {
		obj.obj.Failed = NewIsisIIHLocalGRLastAttemptFailed().msg()
	}

	return obj
}

// description is TBD
// Succeeded returns a IsisIIHLocalGRLastAttemptSucceeded
func (obj *isisIIHLocalGRLastAttemptStatus) Succeeded() IsisIIHLocalGRLastAttemptSucceeded {
	if obj.obj.Succeeded == nil {
		obj.setChoice(IsisIIHLocalGRLastAttemptStatusChoice.SUCCEEDED)
	}
	if obj.succeededHolder == nil {
		obj.succeededHolder = &isisIIHLocalGRLastAttemptSucceeded{obj: obj.obj.Succeeded}
	}
	return obj.succeededHolder
}

// description is TBD
// Succeeded returns a IsisIIHLocalGRLastAttemptSucceeded
func (obj *isisIIHLocalGRLastAttemptStatus) HasSucceeded() bool {
	return obj.obj.Succeeded != nil
}

// description is TBD
// SetSucceeded sets the IsisIIHLocalGRLastAttemptSucceeded value in the IsisIIHLocalGRLastAttemptStatus object
func (obj *isisIIHLocalGRLastAttemptStatus) SetSucceeded(value IsisIIHLocalGRLastAttemptSucceeded) IsisIIHLocalGRLastAttemptStatus {
	obj.setChoice(IsisIIHLocalGRLastAttemptStatusChoice.SUCCEEDED)
	obj.succeededHolder = nil
	obj.obj.Succeeded = value.msg()

	return obj
}

// description is TBD
// Failed returns a IsisIIHLocalGRLastAttemptFailed
func (obj *isisIIHLocalGRLastAttemptStatus) Failed() IsisIIHLocalGRLastAttemptFailed {
	if obj.obj.Failed == nil {
		obj.setChoice(IsisIIHLocalGRLastAttemptStatusChoice.FAILED)
	}
	if obj.failedHolder == nil {
		obj.failedHolder = &isisIIHLocalGRLastAttemptFailed{obj: obj.obj.Failed}
	}
	return obj.failedHolder
}

// description is TBD
// Failed returns a IsisIIHLocalGRLastAttemptFailed
func (obj *isisIIHLocalGRLastAttemptStatus) HasFailed() bool {
	return obj.obj.Failed != nil
}

// description is TBD
// SetFailed sets the IsisIIHLocalGRLastAttemptFailed value in the IsisIIHLocalGRLastAttemptStatus object
func (obj *isisIIHLocalGRLastAttemptStatus) SetFailed(value IsisIIHLocalGRLastAttemptFailed) IsisIIHLocalGRLastAttemptStatus {
	obj.setChoice(IsisIIHLocalGRLastAttemptStatusChoice.FAILED)
	obj.failedHolder = nil
	obj.obj.Failed = value.msg()

	return obj
}

// description is TBD
// Inprogress returns a bool
func (obj *isisIIHLocalGRLastAttemptStatus) Inprogress() bool {

	if obj.obj.Inprogress == nil {
		obj.setChoice(IsisIIHLocalGRLastAttemptStatusChoice.INPROGRESS)
	}

	return *obj.obj.Inprogress

}

// description is TBD
// Inprogress returns a bool
func (obj *isisIIHLocalGRLastAttemptStatus) HasInprogress() bool {
	return obj.obj.Inprogress != nil
}

// description is TBD
// SetInprogress sets the bool value in the IsisIIHLocalGRLastAttemptStatus object
func (obj *isisIIHLocalGRLastAttemptStatus) SetInprogress(value bool) IsisIIHLocalGRLastAttemptStatus {
	obj.setChoice(IsisIIHLocalGRLastAttemptStatusChoice.INPROGRESS)
	obj.obj.Inprogress = &value
	return obj
}

// description is TBD
// Unavailable returns a bool
func (obj *isisIIHLocalGRLastAttemptStatus) Unavailable() bool {

	if obj.obj.Unavailable == nil {
		obj.setChoice(IsisIIHLocalGRLastAttemptStatusChoice.UNAVAILABLE)
	}

	return *obj.obj.Unavailable

}

// description is TBD
// Unavailable returns a bool
func (obj *isisIIHLocalGRLastAttemptStatus) HasUnavailable() bool {
	return obj.obj.Unavailable != nil
}

// description is TBD
// SetUnavailable sets the bool value in the IsisIIHLocalGRLastAttemptStatus object
func (obj *isisIIHLocalGRLastAttemptStatus) SetUnavailable(value bool) IsisIIHLocalGRLastAttemptStatus {
	obj.setChoice(IsisIIHLocalGRLastAttemptStatusChoice.UNAVAILABLE)
	obj.obj.Unavailable = &value
	return obj
}

func (obj *isisIIHLocalGRLastAttemptStatus) validateObj(vObj *validation, set_default bool) {
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

func (obj *isisIIHLocalGRLastAttemptStatus) setDefault() {
	var choices_set int = 0
	var choice IsisIIHLocalGRLastAttemptStatusChoiceEnum

	if obj.obj.Succeeded != nil {
		choices_set += 1
		choice = IsisIIHLocalGRLastAttemptStatusChoice.SUCCEEDED
	}

	if obj.obj.Failed != nil {
		choices_set += 1
		choice = IsisIIHLocalGRLastAttemptStatusChoice.FAILED
	}

	if obj.obj.Inprogress != nil {
		choices_set += 1
		choice = IsisIIHLocalGRLastAttemptStatusChoice.INPROGRESS
	}

	if obj.obj.Unavailable != nil {
		choices_set += 1
		choice = IsisIIHLocalGRLastAttemptStatusChoice.UNAVAILABLE
	}
	if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in IsisIIHLocalGRLastAttemptStatus")
			}
		} else {
			intVal := otg.IsisIIHLocalGRLastAttemptStatus_Choice_Enum_value[string(choice)]
			enumValue := otg.IsisIIHLocalGRLastAttemptStatus_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
