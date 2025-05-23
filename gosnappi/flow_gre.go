package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowGre *****
type flowGre struct {
	validation
	obj                   *otg.FlowGre
	marshaller            marshalFlowGre
	unMarshaller          unMarshalFlowGre
	checksumPresentHolder PatternFlowGreChecksumPresent
	reserved0Holder       PatternFlowGreReserved0
	versionHolder         PatternFlowGreVersion
	protocolHolder        PatternFlowGreProtocol
	checksumHolder        PatternFlowGreChecksum
	reserved1Holder       PatternFlowGreReserved1
}

func NewFlowGre() FlowGre {
	obj := flowGre{obj: &otg.FlowGre{}}
	obj.setDefault()
	return &obj
}

func (obj *flowGre) msg() *otg.FlowGre {
	return obj.obj
}

func (obj *flowGre) setMsg(msg *otg.FlowGre) FlowGre {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowGre struct {
	obj *flowGre
}

type marshalFlowGre interface {
	// ToProto marshals FlowGre to protobuf object *otg.FlowGre
	ToProto() (*otg.FlowGre, error)
	// ToPbText marshals FlowGre to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowGre to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowGre to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals FlowGre to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalflowGre struct {
	obj *flowGre
}

type unMarshalFlowGre interface {
	// FromProto unmarshals FlowGre from protobuf object *otg.FlowGre
	FromProto(msg *otg.FlowGre) (FlowGre, error)
	// FromPbText unmarshals FlowGre from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowGre from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowGre from JSON text
	FromJson(value string) error
}

func (obj *flowGre) Marshal() marshalFlowGre {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowGre{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowGre) Unmarshal() unMarshalFlowGre {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowGre{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowGre) ToProto() (*otg.FlowGre, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowGre) FromProto(msg *otg.FlowGre) (FlowGre, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowGre) ToPbText() (string, error) {
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

func (m *unMarshalflowGre) FromPbText(value string) error {
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

func (m *marshalflowGre) ToYaml() (string, error) {
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

func (m *unMarshalflowGre) FromYaml(value string) error {
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

func (m *marshalflowGre) ToJsonRaw() (string, error) {
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

func (m *marshalflowGre) ToJson() (string, error) {
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

func (m *unMarshalflowGre) FromJson(value string) error {
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

func (obj *flowGre) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowGre) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowGre) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowGre) Clone() (FlowGre, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowGre()
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

func (obj *flowGre) setNil() {
	obj.checksumPresentHolder = nil
	obj.reserved0Holder = nil
	obj.versionHolder = nil
	obj.protocolHolder = nil
	obj.checksumHolder = nil
	obj.reserved1Holder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowGre is standard GRE packet header (RFC2784)
type FlowGre interface {
	Validation
	// msg marshals FlowGre to protobuf object *otg.FlowGre
	// and doesn't set defaults
	msg() *otg.FlowGre
	// setMsg unmarshals FlowGre from protobuf object *otg.FlowGre
	// and doesn't set defaults
	setMsg(*otg.FlowGre) FlowGre
	// provides marshal interface
	Marshal() marshalFlowGre
	// provides unmarshal interface
	Unmarshal() unMarshalFlowGre
	// validate validates FlowGre
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowGre, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// ChecksumPresent returns PatternFlowGreChecksumPresent, set in FlowGre.
	// PatternFlowGreChecksumPresent is checksum present bit
	ChecksumPresent() PatternFlowGreChecksumPresent
	// SetChecksumPresent assigns PatternFlowGreChecksumPresent provided by user to FlowGre.
	// PatternFlowGreChecksumPresent is checksum present bit
	SetChecksumPresent(value PatternFlowGreChecksumPresent) FlowGre
	// HasChecksumPresent checks if ChecksumPresent has been set in FlowGre
	HasChecksumPresent() bool
	// Reserved0 returns PatternFlowGreReserved0, set in FlowGre.
	// PatternFlowGreReserved0 is reserved bits
	Reserved0() PatternFlowGreReserved0
	// SetReserved0 assigns PatternFlowGreReserved0 provided by user to FlowGre.
	// PatternFlowGreReserved0 is reserved bits
	SetReserved0(value PatternFlowGreReserved0) FlowGre
	// HasReserved0 checks if Reserved0 has been set in FlowGre
	HasReserved0() bool
	// Version returns PatternFlowGreVersion, set in FlowGre.
	// PatternFlowGreVersion is gRE version number
	Version() PatternFlowGreVersion
	// SetVersion assigns PatternFlowGreVersion provided by user to FlowGre.
	// PatternFlowGreVersion is gRE version number
	SetVersion(value PatternFlowGreVersion) FlowGre
	// HasVersion checks if Version has been set in FlowGre
	HasVersion() bool
	// Protocol returns PatternFlowGreProtocol, set in FlowGre.
	// PatternFlowGreProtocol is protocol type of encapsulated payload
	Protocol() PatternFlowGreProtocol
	// SetProtocol assigns PatternFlowGreProtocol provided by user to FlowGre.
	// PatternFlowGreProtocol is protocol type of encapsulated payload
	SetProtocol(value PatternFlowGreProtocol) FlowGre
	// HasProtocol checks if Protocol has been set in FlowGre
	HasProtocol() bool
	// Checksum returns PatternFlowGreChecksum, set in FlowGre.
	// PatternFlowGreChecksum is optional checksum of GRE header and payload. Only present if the checksum_present bit is set.
	Checksum() PatternFlowGreChecksum
	// SetChecksum assigns PatternFlowGreChecksum provided by user to FlowGre.
	// PatternFlowGreChecksum is optional checksum of GRE header and payload. Only present if the checksum_present bit is set.
	SetChecksum(value PatternFlowGreChecksum) FlowGre
	// HasChecksum checks if Checksum has been set in FlowGre
	HasChecksum() bool
	// Reserved1 returns PatternFlowGreReserved1, set in FlowGre.
	// PatternFlowGreReserved1 is optional reserved field. Only present if the checksum_present bit is set.
	Reserved1() PatternFlowGreReserved1
	// SetReserved1 assigns PatternFlowGreReserved1 provided by user to FlowGre.
	// PatternFlowGreReserved1 is optional reserved field. Only present if the checksum_present bit is set.
	SetReserved1(value PatternFlowGreReserved1) FlowGre
	// HasReserved1 checks if Reserved1 has been set in FlowGre
	HasReserved1() bool
	setNil()
}

// description is TBD
// ChecksumPresent returns a PatternFlowGreChecksumPresent
func (obj *flowGre) ChecksumPresent() PatternFlowGreChecksumPresent {
	if obj.obj.ChecksumPresent == nil {
		obj.obj.ChecksumPresent = NewPatternFlowGreChecksumPresent().msg()
	}
	if obj.checksumPresentHolder == nil {
		obj.checksumPresentHolder = &patternFlowGreChecksumPresent{obj: obj.obj.ChecksumPresent}
	}
	return obj.checksumPresentHolder
}

// description is TBD
// ChecksumPresent returns a PatternFlowGreChecksumPresent
func (obj *flowGre) HasChecksumPresent() bool {
	return obj.obj.ChecksumPresent != nil
}

// description is TBD
// SetChecksumPresent sets the PatternFlowGreChecksumPresent value in the FlowGre object
func (obj *flowGre) SetChecksumPresent(value PatternFlowGreChecksumPresent) FlowGre {

	obj.checksumPresentHolder = nil
	obj.obj.ChecksumPresent = value.msg()

	return obj
}

// description is TBD
// Reserved0 returns a PatternFlowGreReserved0
func (obj *flowGre) Reserved0() PatternFlowGreReserved0 {
	if obj.obj.Reserved0 == nil {
		obj.obj.Reserved0 = NewPatternFlowGreReserved0().msg()
	}
	if obj.reserved0Holder == nil {
		obj.reserved0Holder = &patternFlowGreReserved0{obj: obj.obj.Reserved0}
	}
	return obj.reserved0Holder
}

// description is TBD
// Reserved0 returns a PatternFlowGreReserved0
func (obj *flowGre) HasReserved0() bool {
	return obj.obj.Reserved0 != nil
}

// description is TBD
// SetReserved0 sets the PatternFlowGreReserved0 value in the FlowGre object
func (obj *flowGre) SetReserved0(value PatternFlowGreReserved0) FlowGre {

	obj.reserved0Holder = nil
	obj.obj.Reserved0 = value.msg()

	return obj
}

// description is TBD
// Version returns a PatternFlowGreVersion
func (obj *flowGre) Version() PatternFlowGreVersion {
	if obj.obj.Version == nil {
		obj.obj.Version = NewPatternFlowGreVersion().msg()
	}
	if obj.versionHolder == nil {
		obj.versionHolder = &patternFlowGreVersion{obj: obj.obj.Version}
	}
	return obj.versionHolder
}

// description is TBD
// Version returns a PatternFlowGreVersion
func (obj *flowGre) HasVersion() bool {
	return obj.obj.Version != nil
}

// description is TBD
// SetVersion sets the PatternFlowGreVersion value in the FlowGre object
func (obj *flowGre) SetVersion(value PatternFlowGreVersion) FlowGre {

	obj.versionHolder = nil
	obj.obj.Version = value.msg()

	return obj
}

// description is TBD
// Protocol returns a PatternFlowGreProtocol
func (obj *flowGre) Protocol() PatternFlowGreProtocol {
	if obj.obj.Protocol == nil {
		obj.obj.Protocol = NewPatternFlowGreProtocol().msg()
	}
	if obj.protocolHolder == nil {
		obj.protocolHolder = &patternFlowGreProtocol{obj: obj.obj.Protocol}
	}
	return obj.protocolHolder
}

// description is TBD
// Protocol returns a PatternFlowGreProtocol
func (obj *flowGre) HasProtocol() bool {
	return obj.obj.Protocol != nil
}

// description is TBD
// SetProtocol sets the PatternFlowGreProtocol value in the FlowGre object
func (obj *flowGre) SetProtocol(value PatternFlowGreProtocol) FlowGre {

	obj.protocolHolder = nil
	obj.obj.Protocol = value.msg()

	return obj
}

// description is TBD
// Checksum returns a PatternFlowGreChecksum
func (obj *flowGre) Checksum() PatternFlowGreChecksum {
	if obj.obj.Checksum == nil {
		obj.obj.Checksum = NewPatternFlowGreChecksum().msg()
	}
	if obj.checksumHolder == nil {
		obj.checksumHolder = &patternFlowGreChecksum{obj: obj.obj.Checksum}
	}
	return obj.checksumHolder
}

// description is TBD
// Checksum returns a PatternFlowGreChecksum
func (obj *flowGre) HasChecksum() bool {
	return obj.obj.Checksum != nil
}

// description is TBD
// SetChecksum sets the PatternFlowGreChecksum value in the FlowGre object
func (obj *flowGre) SetChecksum(value PatternFlowGreChecksum) FlowGre {

	obj.checksumHolder = nil
	obj.obj.Checksum = value.msg()

	return obj
}

// description is TBD
// Reserved1 returns a PatternFlowGreReserved1
func (obj *flowGre) Reserved1() PatternFlowGreReserved1 {
	if obj.obj.Reserved1 == nil {
		obj.obj.Reserved1 = NewPatternFlowGreReserved1().msg()
	}
	if obj.reserved1Holder == nil {
		obj.reserved1Holder = &patternFlowGreReserved1{obj: obj.obj.Reserved1}
	}
	return obj.reserved1Holder
}

// description is TBD
// Reserved1 returns a PatternFlowGreReserved1
func (obj *flowGre) HasReserved1() bool {
	return obj.obj.Reserved1 != nil
}

// description is TBD
// SetReserved1 sets the PatternFlowGreReserved1 value in the FlowGre object
func (obj *flowGre) SetReserved1(value PatternFlowGreReserved1) FlowGre {

	obj.reserved1Holder = nil
	obj.obj.Reserved1 = value.msg()

	return obj
}

func (obj *flowGre) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.ChecksumPresent != nil {

		obj.ChecksumPresent().validateObj(vObj, set_default)
	}

	if obj.obj.Reserved0 != nil {

		obj.Reserved0().validateObj(vObj, set_default)
	}

	if obj.obj.Version != nil {

		obj.Version().validateObj(vObj, set_default)
	}

	if obj.obj.Protocol != nil {

		obj.Protocol().validateObj(vObj, set_default)
	}

	if obj.obj.Checksum != nil {

		obj.Checksum().validateObj(vObj, set_default)
	}

	if obj.obj.Reserved1 != nil {

		obj.Reserved1().validateObj(vObj, set_default)
	}

}

func (obj *flowGre) setDefault() {

}
