package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics *****
type updateProtocolConfigIsisInterfaceTrafficEngineeringMetrics struct {
	validation
	obj          *otg.UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics
	marshaller   marshalUpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics
	unMarshaller unMarshalUpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics
}

func NewUpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics() UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics {
	obj := updateProtocolConfigIsisInterfaceTrafficEngineeringMetrics{obj: &otg.UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics{}}
	obj.setDefault()
	return &obj
}

func (obj *updateProtocolConfigIsisInterfaceTrafficEngineeringMetrics) msg() *otg.UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics {
	return obj.obj
}

func (obj *updateProtocolConfigIsisInterfaceTrafficEngineeringMetrics) setMsg(msg *otg.UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics) UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalupdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics struct {
	obj *updateProtocolConfigIsisInterfaceTrafficEngineeringMetrics
}

type marshalUpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics interface {
	// ToProto marshals UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics to protobuf object *otg.UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics
	ToProto() (*otg.UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics, error)
	// ToPbText marshals UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics to YAML text
	ToYaml() (string, error)
	// ToJson marshals UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics to JSON text
	ToJson() (string, error)
}

type unMarshalupdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics struct {
	obj *updateProtocolConfigIsisInterfaceTrafficEngineeringMetrics
}

type unMarshalUpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics interface {
	// FromProto unmarshals UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics from protobuf object *otg.UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics
	FromProto(msg *otg.UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics) (UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics, error)
	// FromPbText unmarshals UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics from YAML text
	FromYaml(value string) error
	// FromJson unmarshals UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics from JSON text
	FromJson(value string) error
}

func (obj *updateProtocolConfigIsisInterfaceTrafficEngineeringMetrics) Marshal() marshalUpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics {
	if obj.marshaller == nil {
		obj.marshaller = &marshalupdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics{obj: obj}
	}
	return obj.marshaller
}

func (obj *updateProtocolConfigIsisInterfaceTrafficEngineeringMetrics) Unmarshal() unMarshalUpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalupdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalupdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics) ToProto() (*otg.UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalupdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics) FromProto(msg *otg.UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics) (UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalupdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics) ToPbText() (string, error) {
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

func (m *unMarshalupdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics) FromPbText(value string) error {
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

func (m *marshalupdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics) ToYaml() (string, error) {
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

func (m *unMarshalupdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics) FromYaml(value string) error {
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

func (m *marshalupdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics) ToJson() (string, error) {
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

func (m *unMarshalupdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics) FromJson(value string) error {
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

func (obj *updateProtocolConfigIsisInterfaceTrafficEngineeringMetrics) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *updateProtocolConfigIsisInterfaceTrafficEngineeringMetrics) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *updateProtocolConfigIsisInterfaceTrafficEngineeringMetrics) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *updateProtocolConfigIsisInterfaceTrafficEngineeringMetrics) Clone() (UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewUpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics()
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

// UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics is the traffic engineering metric values to be updated on the IS-IS interface.
type UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics interface {
	Validation
	// msg marshals UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics to protobuf object *otg.UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics
	// and doesn't set defaults
	msg() *otg.UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics
	// setMsg unmarshals UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics from protobuf object *otg.UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics
	// and doesn't set defaults
	setMsg(*otg.UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics) UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics
	// provides marshal interface
	Marshal() marshalUpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics
	// provides unmarshal interface
	Unmarshal() unMarshalUpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics
	// validate validates UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetricsChoiceEnum, set in UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics
	Choice() UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetricsChoiceEnum
	// setChoice assigns UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetricsChoiceEnum provided by user to UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics
	setChoice(value UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetricsChoiceEnum) UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics
	// HasChoice checks if Choice has been set in UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics
	HasChoice() bool
	// AdministrativeGroup returns string, set in UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics.
	AdministrativeGroup() string
	// SetAdministrativeGroup assigns string provided by user to UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics
	SetAdministrativeGroup(value string) UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics
	// HasAdministrativeGroup checks if AdministrativeGroup has been set in UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics
	HasAdministrativeGroup() bool
	// MetricLevel returns uint32, set in UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics.
	MetricLevel() uint32
	// SetMetricLevel assigns uint32 provided by user to UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics
	SetMetricLevel(value uint32) UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics
	// HasMetricLevel checks if MetricLevel has been set in UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics
	HasMetricLevel() bool
	// MaxBandwidth returns uint32, set in UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics.
	MaxBandwidth() uint32
	// SetMaxBandwidth assigns uint32 provided by user to UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics
	SetMaxBandwidth(value uint32) UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics
	// HasMaxBandwidth checks if MaxBandwidth has been set in UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics
	HasMaxBandwidth() bool
}

type UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetricsChoiceEnum string

// Enum of Choice on UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics
var UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetricsChoice = struct {
	ADMINISTRATIVE_GROUP UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetricsChoiceEnum
	METRIC_LEVEL         UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetricsChoiceEnum
	MAX_BANDWIDTH        UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetricsChoiceEnum
}{
	ADMINISTRATIVE_GROUP: UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetricsChoiceEnum("administrative_group"),
	METRIC_LEVEL:         UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetricsChoiceEnum("metric_level"),
	MAX_BANDWIDTH:        UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetricsChoiceEnum("max_bandwidth"),
}

func (obj *updateProtocolConfigIsisInterfaceTrafficEngineeringMetrics) Choice() UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetricsChoiceEnum {
	return UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetricsChoiceEnum(obj.obj.Choice.Enum().String())
}

// The traffic engineering metric to be updated on the IS-IS interface.
// Choice returns a string
func (obj *updateProtocolConfigIsisInterfaceTrafficEngineeringMetrics) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *updateProtocolConfigIsisInterfaceTrafficEngineeringMetrics) setChoice(value UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetricsChoiceEnum) UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics {
	intValue, ok := otg.UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetricsChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.MaxBandwidth = nil
	obj.obj.MetricLevel = nil
	obj.obj.AdministrativeGroup = nil

	if value == UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetricsChoice.ADMINISTRATIVE_GROUP {
		defaultValue := "0000000"
		obj.obj.AdministrativeGroup = &defaultValue
	}

	if value == UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetricsChoice.METRIC_LEVEL {
		defaultValue := uint32(0)
		obj.obj.MetricLevel = &defaultValue
	}

	if value == UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetricsChoice.MAX_BANDWIDTH {
		defaultValue := uint32(125000000)
		obj.obj.MaxBandwidth = &defaultValue
	}

	return obj
}

// TThe Administrative group sub-TLV (sub-TLV 3). It is a 4-octet  user-defined bit mask used to update administrative group numbers  to the interface.
// AdministrativeGroup returns a string
func (obj *updateProtocolConfigIsisInterfaceTrafficEngineeringMetrics) AdministrativeGroup() string {

	if obj.obj.AdministrativeGroup == nil {
		obj.setChoice(UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetricsChoice.ADMINISTRATIVE_GROUP)
	}

	return *obj.obj.AdministrativeGroup

}

// TThe Administrative group sub-TLV (sub-TLV 3). It is a 4-octet  user-defined bit mask used to update administrative group numbers  to the interface.
// AdministrativeGroup returns a string
func (obj *updateProtocolConfigIsisInterfaceTrafficEngineeringMetrics) HasAdministrativeGroup() bool {
	return obj.obj.AdministrativeGroup != nil
}

// TThe Administrative group sub-TLV (sub-TLV 3). It is a 4-octet  user-defined bit mask used to update administrative group numbers  to the interface.
// SetAdministrativeGroup sets the string value in the UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics object
func (obj *updateProtocolConfigIsisInterfaceTrafficEngineeringMetrics) SetAdministrativeGroup(value string) UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics {
	obj.setChoice(UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetricsChoice.ADMINISTRATIVE_GROUP)
	obj.obj.AdministrativeGroup = &value
	return obj
}

// The user-assigned link metric to update for Traffic Engineering.
// MetricLevel returns a uint32
func (obj *updateProtocolConfigIsisInterfaceTrafficEngineeringMetrics) MetricLevel() uint32 {

	if obj.obj.MetricLevel == nil {
		obj.setChoice(UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetricsChoice.METRIC_LEVEL)
	}

	return *obj.obj.MetricLevel

}

// The user-assigned link metric to update for Traffic Engineering.
// MetricLevel returns a uint32
func (obj *updateProtocolConfigIsisInterfaceTrafficEngineeringMetrics) HasMetricLevel() bool {
	return obj.obj.MetricLevel != nil
}

// The user-assigned link metric to update for Traffic Engineering.
// SetMetricLevel sets the uint32 value in the UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics object
func (obj *updateProtocolConfigIsisInterfaceTrafficEngineeringMetrics) SetMetricLevel(value uint32) UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics {
	obj.setChoice(UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetricsChoice.METRIC_LEVEL)
	obj.obj.MetricLevel = &value
	return obj
}

// The maximum link bandwidth (sub-TLV 9) in bytes/sec to update for this link for a direction.
// MaxBandwidth returns a uint32
func (obj *updateProtocolConfigIsisInterfaceTrafficEngineeringMetrics) MaxBandwidth() uint32 {

	if obj.obj.MaxBandwidth == nil {
		obj.setChoice(UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetricsChoice.MAX_BANDWIDTH)
	}

	return *obj.obj.MaxBandwidth

}

// The maximum link bandwidth (sub-TLV 9) in bytes/sec to update for this link for a direction.
// MaxBandwidth returns a uint32
func (obj *updateProtocolConfigIsisInterfaceTrafficEngineeringMetrics) HasMaxBandwidth() bool {
	return obj.obj.MaxBandwidth != nil
}

// The maximum link bandwidth (sub-TLV 9) in bytes/sec to update for this link for a direction.
// SetMaxBandwidth sets the uint32 value in the UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics object
func (obj *updateProtocolConfigIsisInterfaceTrafficEngineeringMetrics) SetMaxBandwidth(value uint32) UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics {
	obj.setChoice(UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetricsChoice.MAX_BANDWIDTH)
	obj.obj.MaxBandwidth = &value
	return obj
}

func (obj *updateProtocolConfigIsisInterfaceTrafficEngineeringMetrics) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.AdministrativeGroup != nil {

		err := obj.validateHex(obj.AdministrativeGroup())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics.AdministrativeGroup"))
		}

	}

}

func (obj *updateProtocolConfigIsisInterfaceTrafficEngineeringMetrics) setDefault() {
	var choices_set int = 0
	var choice UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetricsChoiceEnum

	if obj.obj.AdministrativeGroup != nil {
		choices_set += 1
		choice = UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetricsChoice.ADMINISTRATIVE_GROUP
	}

	if obj.obj.MetricLevel != nil {
		choices_set += 1
		choice = UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetricsChoice.METRIC_LEVEL
	}

	if obj.obj.MaxBandwidth != nil {
		choices_set += 1
		choice = UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetricsChoice.MAX_BANDWIDTH
	}
	if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics")
			}
		} else {
			intVal := otg.UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics_Choice_Enum_value[string(choice)]
			enumValue := otg.UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

	if obj.obj.AdministrativeGroup == nil && choice == UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetricsChoice.ADMINISTRATIVE_GROUP {
		obj.SetAdministrativeGroup("0000000")
	}
	if obj.obj.MetricLevel == nil && choice == UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetricsChoice.METRIC_LEVEL {
		obj.SetMetricLevel(0)
	}
	if obj.obj.MaxBandwidth == nil && choice == UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetricsChoice.MAX_BANDWIDTH {
		obj.SetMaxBandwidth(125000000)
	}

}
