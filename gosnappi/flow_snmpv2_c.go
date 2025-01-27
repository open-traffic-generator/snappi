package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowSnmpv2C *****
type flowSnmpv2C struct {
	validation
	obj           *otg.FlowSnmpv2C
	marshaller    marshalFlowSnmpv2C
	unMarshaller  unMarshalFlowSnmpv2C
	versionHolder PatternFlowSnmpv2CVersion
	dataHolder    FlowSnmpv2CData
}

func NewFlowSnmpv2C() FlowSnmpv2C {
	obj := flowSnmpv2C{obj: &otg.FlowSnmpv2C{}}
	obj.setDefault()
	return &obj
}

func (obj *flowSnmpv2C) msg() *otg.FlowSnmpv2C {
	return obj.obj
}

func (obj *flowSnmpv2C) setMsg(msg *otg.FlowSnmpv2C) FlowSnmpv2C {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowSnmpv2C struct {
	obj *flowSnmpv2C
}

type marshalFlowSnmpv2C interface {
	// ToProto marshals FlowSnmpv2C to protobuf object *otg.FlowSnmpv2C
	ToProto() (*otg.FlowSnmpv2C, error)
	// ToPbText marshals FlowSnmpv2C to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowSnmpv2C to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowSnmpv2C to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals FlowSnmpv2C to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalflowSnmpv2C struct {
	obj *flowSnmpv2C
}

type unMarshalFlowSnmpv2C interface {
	// FromProto unmarshals FlowSnmpv2C from protobuf object *otg.FlowSnmpv2C
	FromProto(msg *otg.FlowSnmpv2C) (FlowSnmpv2C, error)
	// FromPbText unmarshals FlowSnmpv2C from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowSnmpv2C from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowSnmpv2C from JSON text
	FromJson(value string) error
}

func (obj *flowSnmpv2C) Marshal() marshalFlowSnmpv2C {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowSnmpv2C{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowSnmpv2C) Unmarshal() unMarshalFlowSnmpv2C {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowSnmpv2C{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowSnmpv2C) ToProto() (*otg.FlowSnmpv2C, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowSnmpv2C) FromProto(msg *otg.FlowSnmpv2C) (FlowSnmpv2C, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowSnmpv2C) ToPbText() (string, error) {
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

func (m *unMarshalflowSnmpv2C) FromPbText(value string) error {
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

func (m *marshalflowSnmpv2C) ToYaml() (string, error) {
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

func (m *unMarshalflowSnmpv2C) FromYaml(value string) error {
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

func (m *marshalflowSnmpv2C) ToJsonRaw() (string, error) {
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
	return string(data), nil
}

func (m *marshalflowSnmpv2C) ToJson() (string, error) {
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

func (m *unMarshalflowSnmpv2C) FromJson(value string) error {
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

func (obj *flowSnmpv2C) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowSnmpv2C) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowSnmpv2C) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowSnmpv2C) Clone() (FlowSnmpv2C, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowSnmpv2C()
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

func (obj *flowSnmpv2C) setNil() {
	obj.versionHolder = nil
	obj.dataHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowSnmpv2C is sNMPv2C packet header as defined in RFC1901 and RFC3416.
type FlowSnmpv2C interface {
	Validation
	// msg marshals FlowSnmpv2C to protobuf object *otg.FlowSnmpv2C
	// and doesn't set defaults
	msg() *otg.FlowSnmpv2C
	// setMsg unmarshals FlowSnmpv2C from protobuf object *otg.FlowSnmpv2C
	// and doesn't set defaults
	setMsg(*otg.FlowSnmpv2C) FlowSnmpv2C
	// provides marshal interface
	Marshal() marshalFlowSnmpv2C
	// provides unmarshal interface
	Unmarshal() unMarshalFlowSnmpv2C
	// validate validates FlowSnmpv2C
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowSnmpv2C, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Version returns PatternFlowSnmpv2CVersion, set in FlowSnmpv2C.
	Version() PatternFlowSnmpv2CVersion
	// SetVersion assigns PatternFlowSnmpv2CVersion provided by user to FlowSnmpv2C.
	SetVersion(value PatternFlowSnmpv2CVersion) FlowSnmpv2C
	// HasVersion checks if Version has been set in FlowSnmpv2C
	HasVersion() bool
	// Community returns string, set in FlowSnmpv2C.
	Community() string
	// SetCommunity assigns string provided by user to FlowSnmpv2C
	SetCommunity(value string) FlowSnmpv2C
	// HasCommunity checks if Community has been set in FlowSnmpv2C
	HasCommunity() bool
	// Data returns FlowSnmpv2CData, set in FlowSnmpv2C.
	Data() FlowSnmpv2CData
	// SetData assigns FlowSnmpv2CData provided by user to FlowSnmpv2C.
	SetData(value FlowSnmpv2CData) FlowSnmpv2C
	setNil()
}

// description is TBD
// Version returns a PatternFlowSnmpv2CVersion
func (obj *flowSnmpv2C) Version() PatternFlowSnmpv2CVersion {
	if obj.obj.Version == nil {
		obj.obj.Version = NewPatternFlowSnmpv2CVersion().msg()
	}
	if obj.versionHolder == nil {
		obj.versionHolder = &patternFlowSnmpv2CVersion{obj: obj.obj.Version}
	}
	return obj.versionHolder
}

// description is TBD
// Version returns a PatternFlowSnmpv2CVersion
func (obj *flowSnmpv2C) HasVersion() bool {
	return obj.obj.Version != nil
}

// description is TBD
// SetVersion sets the PatternFlowSnmpv2CVersion value in the FlowSnmpv2C object
func (obj *flowSnmpv2C) SetVersion(value PatternFlowSnmpv2CVersion) FlowSnmpv2C {

	obj.versionHolder = nil
	obj.obj.Version = value.msg()

	return obj
}

// It is an ASCII based octet string which identifies the SNMP community in which the sender and recipient of this message are located. It should match the read-only or read-write community string configured on the recipient for the PDU to be accepted.
// Community returns a string
func (obj *flowSnmpv2C) Community() string {

	return *obj.obj.Community

}

// It is an ASCII based octet string which identifies the SNMP community in which the sender and recipient of this message are located. It should match the read-only or read-write community string configured on the recipient for the PDU to be accepted.
// Community returns a string
func (obj *flowSnmpv2C) HasCommunity() bool {
	return obj.obj.Community != nil
}

// It is an ASCII based octet string which identifies the SNMP community in which the sender and recipient of this message are located. It should match the read-only or read-write community string configured on the recipient for the PDU to be accepted.
// SetCommunity sets the string value in the FlowSnmpv2C object
func (obj *flowSnmpv2C) SetCommunity(value string) FlowSnmpv2C {

	obj.obj.Community = &value
	return obj
}

// description is TBD
// Data returns a FlowSnmpv2CData
func (obj *flowSnmpv2C) Data() FlowSnmpv2CData {
	if obj.obj.Data == nil {
		obj.obj.Data = NewFlowSnmpv2CData().msg()
	}
	if obj.dataHolder == nil {
		obj.dataHolder = &flowSnmpv2CData{obj: obj.obj.Data}
	}
	return obj.dataHolder
}

// description is TBD
// SetData sets the FlowSnmpv2CData value in the FlowSnmpv2C object
func (obj *flowSnmpv2C) SetData(value FlowSnmpv2CData) FlowSnmpv2C {

	obj.dataHolder = nil
	obj.obj.Data = value.msg()

	return obj
}

func (obj *flowSnmpv2C) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Version != nil {

		obj.Version().validateObj(vObj, set_default)
	}

	if obj.obj.Community != nil {

		if len(*obj.obj.Community) > 10000 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"None <= length of FlowSnmpv2C.Community <= 10000 but Got %d",
					len(*obj.obj.Community)))
		}

	}

	// Data is required
	if obj.obj.Data == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Data is required field on interface FlowSnmpv2C")
	}

	if obj.obj.Data != nil {

		obj.Data().validateObj(vObj, set_default)
	}

}

func (obj *flowSnmpv2C) setDefault() {
	if obj.obj.Community == nil {
		obj.SetCommunity("community")
	}

}
