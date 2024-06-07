package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowRSVPPathSenderTspecIntServ *****
type flowRSVPPathSenderTspecIntServ struct {
	validation
	obj                               *otg.FlowRSVPPathSenderTspecIntServ
	marshaller                        marshalFlowRSVPPathSenderTspecIntServ
	unMarshaller                      unMarshalFlowRSVPPathSenderTspecIntServ
	versionHolder                     PatternFlowRSVPPathSenderTspecIntServVersion
	reserved1Holder                   PatternFlowRSVPPathSenderTspecIntServReserved1
	overallLengthHolder               PatternFlowRSVPPathSenderTspecIntServOverallLength
	serviceHeaderHolder               PatternFlowRSVPPathSenderTspecIntServServiceHeader
	zeroBitHolder                     PatternFlowRSVPPathSenderTspecIntServZeroBit
	reserved2Holder                   PatternFlowRSVPPathSenderTspecIntServReserved2
	lengthOfServiceDataHolder         PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData
	parameterIdTokenBucketTspecHolder PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec
	parameter_127FlagHolder           PatternFlowRSVPPathSenderTspecIntServParameter127Flag
	parameter_127LengthHolder         PatternFlowRSVPPathSenderTspecIntServParameter127Length
	minimumPolicedUnitHolder          PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit
	maximumPacketSizeHolder           PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize
}

func NewFlowRSVPPathSenderTspecIntServ() FlowRSVPPathSenderTspecIntServ {
	obj := flowRSVPPathSenderTspecIntServ{obj: &otg.FlowRSVPPathSenderTspecIntServ{}}
	obj.setDefault()
	return &obj
}

func (obj *flowRSVPPathSenderTspecIntServ) msg() *otg.FlowRSVPPathSenderTspecIntServ {
	return obj.obj
}

func (obj *flowRSVPPathSenderTspecIntServ) setMsg(msg *otg.FlowRSVPPathSenderTspecIntServ) FlowRSVPPathSenderTspecIntServ {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowRSVPPathSenderTspecIntServ struct {
	obj *flowRSVPPathSenderTspecIntServ
}

type marshalFlowRSVPPathSenderTspecIntServ interface {
	// ToProto marshals FlowRSVPPathSenderTspecIntServ to protobuf object *otg.FlowRSVPPathSenderTspecIntServ
	ToProto() (*otg.FlowRSVPPathSenderTspecIntServ, error)
	// ToPbText marshals FlowRSVPPathSenderTspecIntServ to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowRSVPPathSenderTspecIntServ to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowRSVPPathSenderTspecIntServ to JSON text
	ToJson() (string, error)
}

type unMarshalflowRSVPPathSenderTspecIntServ struct {
	obj *flowRSVPPathSenderTspecIntServ
}

type unMarshalFlowRSVPPathSenderTspecIntServ interface {
	// FromProto unmarshals FlowRSVPPathSenderTspecIntServ from protobuf object *otg.FlowRSVPPathSenderTspecIntServ
	FromProto(msg *otg.FlowRSVPPathSenderTspecIntServ) (FlowRSVPPathSenderTspecIntServ, error)
	// FromPbText unmarshals FlowRSVPPathSenderTspecIntServ from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowRSVPPathSenderTspecIntServ from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowRSVPPathSenderTspecIntServ from JSON text
	FromJson(value string) error
}

func (obj *flowRSVPPathSenderTspecIntServ) Marshal() marshalFlowRSVPPathSenderTspecIntServ {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowRSVPPathSenderTspecIntServ{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowRSVPPathSenderTspecIntServ) Unmarshal() unMarshalFlowRSVPPathSenderTspecIntServ {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowRSVPPathSenderTspecIntServ{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowRSVPPathSenderTspecIntServ) ToProto() (*otg.FlowRSVPPathSenderTspecIntServ, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowRSVPPathSenderTspecIntServ) FromProto(msg *otg.FlowRSVPPathSenderTspecIntServ) (FlowRSVPPathSenderTspecIntServ, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowRSVPPathSenderTspecIntServ) ToPbText() (string, error) {
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

func (m *unMarshalflowRSVPPathSenderTspecIntServ) FromPbText(value string) error {
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

func (m *marshalflowRSVPPathSenderTspecIntServ) ToYaml() (string, error) {
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

func (m *unMarshalflowRSVPPathSenderTspecIntServ) FromYaml(value string) error {
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

func (m *marshalflowRSVPPathSenderTspecIntServ) ToJson() (string, error) {
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

func (m *unMarshalflowRSVPPathSenderTspecIntServ) FromJson(value string) error {
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

func (obj *flowRSVPPathSenderTspecIntServ) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowRSVPPathSenderTspecIntServ) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowRSVPPathSenderTspecIntServ) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowRSVPPathSenderTspecIntServ) Clone() (FlowRSVPPathSenderTspecIntServ, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowRSVPPathSenderTspecIntServ()
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

func (obj *flowRSVPPathSenderTspecIntServ) setNil() {
	obj.versionHolder = nil
	obj.reserved1Holder = nil
	obj.overallLengthHolder = nil
	obj.serviceHeaderHolder = nil
	obj.zeroBitHolder = nil
	obj.reserved2Holder = nil
	obj.lengthOfServiceDataHolder = nil
	obj.parameterIdTokenBucketTspecHolder = nil
	obj.parameter_127FlagHolder = nil
	obj.parameter_127LengthHolder = nil
	obj.minimumPolicedUnitHolder = nil
	obj.maximumPacketSizeHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowRSVPPathSenderTspecIntServ is int-serv SENDER_TSPEC object: Class = 12, C-Type = 2
type FlowRSVPPathSenderTspecIntServ interface {
	Validation
	// msg marshals FlowRSVPPathSenderTspecIntServ to protobuf object *otg.FlowRSVPPathSenderTspecIntServ
	// and doesn't set defaults
	msg() *otg.FlowRSVPPathSenderTspecIntServ
	// setMsg unmarshals FlowRSVPPathSenderTspecIntServ from protobuf object *otg.FlowRSVPPathSenderTspecIntServ
	// and doesn't set defaults
	setMsg(*otg.FlowRSVPPathSenderTspecIntServ) FlowRSVPPathSenderTspecIntServ
	// provides marshal interface
	Marshal() marshalFlowRSVPPathSenderTspecIntServ
	// provides unmarshal interface
	Unmarshal() unMarshalFlowRSVPPathSenderTspecIntServ
	// validate validates FlowRSVPPathSenderTspecIntServ
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowRSVPPathSenderTspecIntServ, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Version returns PatternFlowRSVPPathSenderTspecIntServVersion, set in FlowRSVPPathSenderTspecIntServ.
	// PatternFlowRSVPPathSenderTspecIntServVersion is message format version number.
	Version() PatternFlowRSVPPathSenderTspecIntServVersion
	// SetVersion assigns PatternFlowRSVPPathSenderTspecIntServVersion provided by user to FlowRSVPPathSenderTspecIntServ.
	// PatternFlowRSVPPathSenderTspecIntServVersion is message format version number.
	SetVersion(value PatternFlowRSVPPathSenderTspecIntServVersion) FlowRSVPPathSenderTspecIntServ
	// HasVersion checks if Version has been set in FlowRSVPPathSenderTspecIntServ
	HasVersion() bool
	// Reserved1 returns PatternFlowRSVPPathSenderTspecIntServReserved1, set in FlowRSVPPathSenderTspecIntServ.
	// PatternFlowRSVPPathSenderTspecIntServReserved1 is reserved.
	Reserved1() PatternFlowRSVPPathSenderTspecIntServReserved1
	// SetReserved1 assigns PatternFlowRSVPPathSenderTspecIntServReserved1 provided by user to FlowRSVPPathSenderTspecIntServ.
	// PatternFlowRSVPPathSenderTspecIntServReserved1 is reserved.
	SetReserved1(value PatternFlowRSVPPathSenderTspecIntServReserved1) FlowRSVPPathSenderTspecIntServ
	// HasReserved1 checks if Reserved1 has been set in FlowRSVPPathSenderTspecIntServ
	HasReserved1() bool
	// OverallLength returns PatternFlowRSVPPathSenderTspecIntServOverallLength, set in FlowRSVPPathSenderTspecIntServ.
	// PatternFlowRSVPPathSenderTspecIntServOverallLength is overall length (7 words not including header).
	OverallLength() PatternFlowRSVPPathSenderTspecIntServOverallLength
	// SetOverallLength assigns PatternFlowRSVPPathSenderTspecIntServOverallLength provided by user to FlowRSVPPathSenderTspecIntServ.
	// PatternFlowRSVPPathSenderTspecIntServOverallLength is overall length (7 words not including header).
	SetOverallLength(value PatternFlowRSVPPathSenderTspecIntServOverallLength) FlowRSVPPathSenderTspecIntServ
	// HasOverallLength checks if OverallLength has been set in FlowRSVPPathSenderTspecIntServ
	HasOverallLength() bool
	// ServiceHeader returns PatternFlowRSVPPathSenderTspecIntServServiceHeader, set in FlowRSVPPathSenderTspecIntServ.
	// PatternFlowRSVPPathSenderTspecIntServServiceHeader is service header, service number - '1' (Generic information) if in a PATH message.
	ServiceHeader() PatternFlowRSVPPathSenderTspecIntServServiceHeader
	// SetServiceHeader assigns PatternFlowRSVPPathSenderTspecIntServServiceHeader provided by user to FlowRSVPPathSenderTspecIntServ.
	// PatternFlowRSVPPathSenderTspecIntServServiceHeader is service header, service number - '1' (Generic information) if in a PATH message.
	SetServiceHeader(value PatternFlowRSVPPathSenderTspecIntServServiceHeader) FlowRSVPPathSenderTspecIntServ
	// HasServiceHeader checks if ServiceHeader has been set in FlowRSVPPathSenderTspecIntServ
	HasServiceHeader() bool
	// ZeroBit returns PatternFlowRSVPPathSenderTspecIntServZeroBit, set in FlowRSVPPathSenderTspecIntServ.
	// PatternFlowRSVPPathSenderTspecIntServZeroBit is mUST be 0.
	ZeroBit() PatternFlowRSVPPathSenderTspecIntServZeroBit
	// SetZeroBit assigns PatternFlowRSVPPathSenderTspecIntServZeroBit provided by user to FlowRSVPPathSenderTspecIntServ.
	// PatternFlowRSVPPathSenderTspecIntServZeroBit is mUST be 0.
	SetZeroBit(value PatternFlowRSVPPathSenderTspecIntServZeroBit) FlowRSVPPathSenderTspecIntServ
	// HasZeroBit checks if ZeroBit has been set in FlowRSVPPathSenderTspecIntServ
	HasZeroBit() bool
	// Reserved2 returns PatternFlowRSVPPathSenderTspecIntServReserved2, set in FlowRSVPPathSenderTspecIntServ.
	// PatternFlowRSVPPathSenderTspecIntServReserved2 is reserved.
	Reserved2() PatternFlowRSVPPathSenderTspecIntServReserved2
	// SetReserved2 assigns PatternFlowRSVPPathSenderTspecIntServReserved2 provided by user to FlowRSVPPathSenderTspecIntServ.
	// PatternFlowRSVPPathSenderTspecIntServReserved2 is reserved.
	SetReserved2(value PatternFlowRSVPPathSenderTspecIntServReserved2) FlowRSVPPathSenderTspecIntServ
	// HasReserved2 checks if Reserved2 has been set in FlowRSVPPathSenderTspecIntServ
	HasReserved2() bool
	// LengthOfServiceData returns PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData, set in FlowRSVPPathSenderTspecIntServ.
	// PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData is length of service data, 6 words not including per-service header.
	LengthOfServiceData() PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData
	// SetLengthOfServiceData assigns PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData provided by user to FlowRSVPPathSenderTspecIntServ.
	// PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData is length of service data, 6 words not including per-service header.
	SetLengthOfServiceData(value PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData) FlowRSVPPathSenderTspecIntServ
	// HasLengthOfServiceData checks if LengthOfServiceData has been set in FlowRSVPPathSenderTspecIntServ
	HasLengthOfServiceData() bool
	// ParameterIdTokenBucketTspec returns PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec, set in FlowRSVPPathSenderTspecIntServ.
	// PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec is parameter ID, parameter 127 (Token Bucket TSpec)
	ParameterIdTokenBucketTspec() PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec
	// SetParameterIdTokenBucketTspec assigns PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec provided by user to FlowRSVPPathSenderTspecIntServ.
	// PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec is parameter ID, parameter 127 (Token Bucket TSpec)
	SetParameterIdTokenBucketTspec(value PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec) FlowRSVPPathSenderTspecIntServ
	// HasParameterIdTokenBucketTspec checks if ParameterIdTokenBucketTspec has been set in FlowRSVPPathSenderTspecIntServ
	HasParameterIdTokenBucketTspec() bool
	// Parameter127Flag returns PatternFlowRSVPPathSenderTspecIntServParameter127Flag, set in FlowRSVPPathSenderTspecIntServ.
	// PatternFlowRSVPPathSenderTspecIntServParameter127Flag is parameter 127 flags (none set)
	Parameter127Flag() PatternFlowRSVPPathSenderTspecIntServParameter127Flag
	// SetParameter127Flag assigns PatternFlowRSVPPathSenderTspecIntServParameter127Flag provided by user to FlowRSVPPathSenderTspecIntServ.
	// PatternFlowRSVPPathSenderTspecIntServParameter127Flag is parameter 127 flags (none set)
	SetParameter127Flag(value PatternFlowRSVPPathSenderTspecIntServParameter127Flag) FlowRSVPPathSenderTspecIntServ
	// HasParameter127Flag checks if Parameter127Flag has been set in FlowRSVPPathSenderTspecIntServ
	HasParameter127Flag() bool
	// Parameter127Length returns PatternFlowRSVPPathSenderTspecIntServParameter127Length, set in FlowRSVPPathSenderTspecIntServ.
	// PatternFlowRSVPPathSenderTspecIntServParameter127Length is parameter 127 length, 5 words not including per-service header
	Parameter127Length() PatternFlowRSVPPathSenderTspecIntServParameter127Length
	// SetParameter127Length assigns PatternFlowRSVPPathSenderTspecIntServParameter127Length provided by user to FlowRSVPPathSenderTspecIntServ.
	// PatternFlowRSVPPathSenderTspecIntServParameter127Length is parameter 127 length, 5 words not including per-service header
	SetParameter127Length(value PatternFlowRSVPPathSenderTspecIntServParameter127Length) FlowRSVPPathSenderTspecIntServ
	// HasParameter127Length checks if Parameter127Length has been set in FlowRSVPPathSenderTspecIntServ
	HasParameter127Length() bool
	// TokenBucketRate returns float32, set in FlowRSVPPathSenderTspecIntServ.
	TokenBucketRate() float32
	// SetTokenBucketRate assigns float32 provided by user to FlowRSVPPathSenderTspecIntServ
	SetTokenBucketRate(value float32) FlowRSVPPathSenderTspecIntServ
	// HasTokenBucketRate checks if TokenBucketRate has been set in FlowRSVPPathSenderTspecIntServ
	HasTokenBucketRate() bool
	// TokenBucketSize returns float32, set in FlowRSVPPathSenderTspecIntServ.
	TokenBucketSize() float32
	// SetTokenBucketSize assigns float32 provided by user to FlowRSVPPathSenderTspecIntServ
	SetTokenBucketSize(value float32) FlowRSVPPathSenderTspecIntServ
	// HasTokenBucketSize checks if TokenBucketSize has been set in FlowRSVPPathSenderTspecIntServ
	HasTokenBucketSize() bool
	// PeakDataRate returns float32, set in FlowRSVPPathSenderTspecIntServ.
	PeakDataRate() float32
	// SetPeakDataRate assigns float32 provided by user to FlowRSVPPathSenderTspecIntServ
	SetPeakDataRate(value float32) FlowRSVPPathSenderTspecIntServ
	// HasPeakDataRate checks if PeakDataRate has been set in FlowRSVPPathSenderTspecIntServ
	HasPeakDataRate() bool
	// MinimumPolicedUnit returns PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit, set in FlowRSVPPathSenderTspecIntServ.
	// PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit is the minimum policed unit parameter should generally be set equal to the size of the smallest packet generated by the application.
	MinimumPolicedUnit() PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit
	// SetMinimumPolicedUnit assigns PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit provided by user to FlowRSVPPathSenderTspecIntServ.
	// PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit is the minimum policed unit parameter should generally be set equal to the size of the smallest packet generated by the application.
	SetMinimumPolicedUnit(value PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit) FlowRSVPPathSenderTspecIntServ
	// HasMinimumPolicedUnit checks if MinimumPolicedUnit has been set in FlowRSVPPathSenderTspecIntServ
	HasMinimumPolicedUnit() bool
	// MaximumPacketSize returns PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize, set in FlowRSVPPathSenderTspecIntServ.
	// PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize is the maximum packet size parameter should be set to the size of the largest packet the application might wish to generate. This value must, by definition, be equal to or larger than the value of The minimum policed unit.
	MaximumPacketSize() PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize
	// SetMaximumPacketSize assigns PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize provided by user to FlowRSVPPathSenderTspecIntServ.
	// PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize is the maximum packet size parameter should be set to the size of the largest packet the application might wish to generate. This value must, by definition, be equal to or larger than the value of The minimum policed unit.
	SetMaximumPacketSize(value PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize) FlowRSVPPathSenderTspecIntServ
	// HasMaximumPacketSize checks if MaximumPacketSize has been set in FlowRSVPPathSenderTspecIntServ
	HasMaximumPacketSize() bool
	setNil()
}

// description is TBD
// Version returns a PatternFlowRSVPPathSenderTspecIntServVersion
func (obj *flowRSVPPathSenderTspecIntServ) Version() PatternFlowRSVPPathSenderTspecIntServVersion {
	if obj.obj.Version == nil {
		obj.obj.Version = NewPatternFlowRSVPPathSenderTspecIntServVersion().msg()
	}
	if obj.versionHolder == nil {
		obj.versionHolder = &patternFlowRSVPPathSenderTspecIntServVersion{obj: obj.obj.Version}
	}
	return obj.versionHolder
}

// description is TBD
// Version returns a PatternFlowRSVPPathSenderTspecIntServVersion
func (obj *flowRSVPPathSenderTspecIntServ) HasVersion() bool {
	return obj.obj.Version != nil
}

// description is TBD
// SetVersion sets the PatternFlowRSVPPathSenderTspecIntServVersion value in the FlowRSVPPathSenderTspecIntServ object
func (obj *flowRSVPPathSenderTspecIntServ) SetVersion(value PatternFlowRSVPPathSenderTspecIntServVersion) FlowRSVPPathSenderTspecIntServ {

	obj.versionHolder = nil
	obj.obj.Version = value.msg()

	return obj
}

// description is TBD
// Reserved1 returns a PatternFlowRSVPPathSenderTspecIntServReserved1
func (obj *flowRSVPPathSenderTspecIntServ) Reserved1() PatternFlowRSVPPathSenderTspecIntServReserved1 {
	if obj.obj.Reserved1 == nil {
		obj.obj.Reserved1 = NewPatternFlowRSVPPathSenderTspecIntServReserved1().msg()
	}
	if obj.reserved1Holder == nil {
		obj.reserved1Holder = &patternFlowRSVPPathSenderTspecIntServReserved1{obj: obj.obj.Reserved1}
	}
	return obj.reserved1Holder
}

// description is TBD
// Reserved1 returns a PatternFlowRSVPPathSenderTspecIntServReserved1
func (obj *flowRSVPPathSenderTspecIntServ) HasReserved1() bool {
	return obj.obj.Reserved1 != nil
}

// description is TBD
// SetReserved1 sets the PatternFlowRSVPPathSenderTspecIntServReserved1 value in the FlowRSVPPathSenderTspecIntServ object
func (obj *flowRSVPPathSenderTspecIntServ) SetReserved1(value PatternFlowRSVPPathSenderTspecIntServReserved1) FlowRSVPPathSenderTspecIntServ {

	obj.reserved1Holder = nil
	obj.obj.Reserved1 = value.msg()

	return obj
}

// description is TBD
// OverallLength returns a PatternFlowRSVPPathSenderTspecIntServOverallLength
func (obj *flowRSVPPathSenderTspecIntServ) OverallLength() PatternFlowRSVPPathSenderTspecIntServOverallLength {
	if obj.obj.OverallLength == nil {
		obj.obj.OverallLength = NewPatternFlowRSVPPathSenderTspecIntServOverallLength().msg()
	}
	if obj.overallLengthHolder == nil {
		obj.overallLengthHolder = &patternFlowRSVPPathSenderTspecIntServOverallLength{obj: obj.obj.OverallLength}
	}
	return obj.overallLengthHolder
}

// description is TBD
// OverallLength returns a PatternFlowRSVPPathSenderTspecIntServOverallLength
func (obj *flowRSVPPathSenderTspecIntServ) HasOverallLength() bool {
	return obj.obj.OverallLength != nil
}

// description is TBD
// SetOverallLength sets the PatternFlowRSVPPathSenderTspecIntServOverallLength value in the FlowRSVPPathSenderTspecIntServ object
func (obj *flowRSVPPathSenderTspecIntServ) SetOverallLength(value PatternFlowRSVPPathSenderTspecIntServOverallLength) FlowRSVPPathSenderTspecIntServ {

	obj.overallLengthHolder = nil
	obj.obj.OverallLength = value.msg()

	return obj
}

// description is TBD
// ServiceHeader returns a PatternFlowRSVPPathSenderTspecIntServServiceHeader
func (obj *flowRSVPPathSenderTspecIntServ) ServiceHeader() PatternFlowRSVPPathSenderTspecIntServServiceHeader {
	if obj.obj.ServiceHeader == nil {
		obj.obj.ServiceHeader = NewPatternFlowRSVPPathSenderTspecIntServServiceHeader().msg()
	}
	if obj.serviceHeaderHolder == nil {
		obj.serviceHeaderHolder = &patternFlowRSVPPathSenderTspecIntServServiceHeader{obj: obj.obj.ServiceHeader}
	}
	return obj.serviceHeaderHolder
}

// description is TBD
// ServiceHeader returns a PatternFlowRSVPPathSenderTspecIntServServiceHeader
func (obj *flowRSVPPathSenderTspecIntServ) HasServiceHeader() bool {
	return obj.obj.ServiceHeader != nil
}

// description is TBD
// SetServiceHeader sets the PatternFlowRSVPPathSenderTspecIntServServiceHeader value in the FlowRSVPPathSenderTspecIntServ object
func (obj *flowRSVPPathSenderTspecIntServ) SetServiceHeader(value PatternFlowRSVPPathSenderTspecIntServServiceHeader) FlowRSVPPathSenderTspecIntServ {

	obj.serviceHeaderHolder = nil
	obj.obj.ServiceHeader = value.msg()

	return obj
}

// description is TBD
// ZeroBit returns a PatternFlowRSVPPathSenderTspecIntServZeroBit
func (obj *flowRSVPPathSenderTspecIntServ) ZeroBit() PatternFlowRSVPPathSenderTspecIntServZeroBit {
	if obj.obj.ZeroBit == nil {
		obj.obj.ZeroBit = NewPatternFlowRSVPPathSenderTspecIntServZeroBit().msg()
	}
	if obj.zeroBitHolder == nil {
		obj.zeroBitHolder = &patternFlowRSVPPathSenderTspecIntServZeroBit{obj: obj.obj.ZeroBit}
	}
	return obj.zeroBitHolder
}

// description is TBD
// ZeroBit returns a PatternFlowRSVPPathSenderTspecIntServZeroBit
func (obj *flowRSVPPathSenderTspecIntServ) HasZeroBit() bool {
	return obj.obj.ZeroBit != nil
}

// description is TBD
// SetZeroBit sets the PatternFlowRSVPPathSenderTspecIntServZeroBit value in the FlowRSVPPathSenderTspecIntServ object
func (obj *flowRSVPPathSenderTspecIntServ) SetZeroBit(value PatternFlowRSVPPathSenderTspecIntServZeroBit) FlowRSVPPathSenderTspecIntServ {

	obj.zeroBitHolder = nil
	obj.obj.ZeroBit = value.msg()

	return obj
}

// description is TBD
// Reserved2 returns a PatternFlowRSVPPathSenderTspecIntServReserved2
func (obj *flowRSVPPathSenderTspecIntServ) Reserved2() PatternFlowRSVPPathSenderTspecIntServReserved2 {
	if obj.obj.Reserved2 == nil {
		obj.obj.Reserved2 = NewPatternFlowRSVPPathSenderTspecIntServReserved2().msg()
	}
	if obj.reserved2Holder == nil {
		obj.reserved2Holder = &patternFlowRSVPPathSenderTspecIntServReserved2{obj: obj.obj.Reserved2}
	}
	return obj.reserved2Holder
}

// description is TBD
// Reserved2 returns a PatternFlowRSVPPathSenderTspecIntServReserved2
func (obj *flowRSVPPathSenderTspecIntServ) HasReserved2() bool {
	return obj.obj.Reserved2 != nil
}

// description is TBD
// SetReserved2 sets the PatternFlowRSVPPathSenderTspecIntServReserved2 value in the FlowRSVPPathSenderTspecIntServ object
func (obj *flowRSVPPathSenderTspecIntServ) SetReserved2(value PatternFlowRSVPPathSenderTspecIntServReserved2) FlowRSVPPathSenderTspecIntServ {

	obj.reserved2Holder = nil
	obj.obj.Reserved2 = value.msg()

	return obj
}

// description is TBD
// LengthOfServiceData returns a PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData
func (obj *flowRSVPPathSenderTspecIntServ) LengthOfServiceData() PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData {
	if obj.obj.LengthOfServiceData == nil {
		obj.obj.LengthOfServiceData = NewPatternFlowRSVPPathSenderTspecIntServLengthOfServiceData().msg()
	}
	if obj.lengthOfServiceDataHolder == nil {
		obj.lengthOfServiceDataHolder = &patternFlowRSVPPathSenderTspecIntServLengthOfServiceData{obj: obj.obj.LengthOfServiceData}
	}
	return obj.lengthOfServiceDataHolder
}

// description is TBD
// LengthOfServiceData returns a PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData
func (obj *flowRSVPPathSenderTspecIntServ) HasLengthOfServiceData() bool {
	return obj.obj.LengthOfServiceData != nil
}

// description is TBD
// SetLengthOfServiceData sets the PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData value in the FlowRSVPPathSenderTspecIntServ object
func (obj *flowRSVPPathSenderTspecIntServ) SetLengthOfServiceData(value PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData) FlowRSVPPathSenderTspecIntServ {

	obj.lengthOfServiceDataHolder = nil
	obj.obj.LengthOfServiceData = value.msg()

	return obj
}

// description is TBD
// ParameterIdTokenBucketTspec returns a PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec
func (obj *flowRSVPPathSenderTspecIntServ) ParameterIdTokenBucketTspec() PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec {
	if obj.obj.ParameterIdTokenBucketTspec == nil {
		obj.obj.ParameterIdTokenBucketTspec = NewPatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec().msg()
	}
	if obj.parameterIdTokenBucketTspecHolder == nil {
		obj.parameterIdTokenBucketTspecHolder = &patternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec{obj: obj.obj.ParameterIdTokenBucketTspec}
	}
	return obj.parameterIdTokenBucketTspecHolder
}

// description is TBD
// ParameterIdTokenBucketTspec returns a PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec
func (obj *flowRSVPPathSenderTspecIntServ) HasParameterIdTokenBucketTspec() bool {
	return obj.obj.ParameterIdTokenBucketTspec != nil
}

// description is TBD
// SetParameterIdTokenBucketTspec sets the PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec value in the FlowRSVPPathSenderTspecIntServ object
func (obj *flowRSVPPathSenderTspecIntServ) SetParameterIdTokenBucketTspec(value PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspec) FlowRSVPPathSenderTspecIntServ {

	obj.parameterIdTokenBucketTspecHolder = nil
	obj.obj.ParameterIdTokenBucketTspec = value.msg()

	return obj
}

// description is TBD
// Parameter127Flag returns a PatternFlowRSVPPathSenderTspecIntServParameter127Flag
func (obj *flowRSVPPathSenderTspecIntServ) Parameter127Flag() PatternFlowRSVPPathSenderTspecIntServParameter127Flag {
	if obj.obj.Parameter_127Flag == nil {
		obj.obj.Parameter_127Flag = NewPatternFlowRSVPPathSenderTspecIntServParameter127Flag().msg()
	}
	if obj.parameter_127FlagHolder == nil {
		obj.parameter_127FlagHolder = &patternFlowRSVPPathSenderTspecIntServParameter127Flag{obj: obj.obj.Parameter_127Flag}
	}
	return obj.parameter_127FlagHolder
}

// description is TBD
// Parameter127Flag returns a PatternFlowRSVPPathSenderTspecIntServParameter127Flag
func (obj *flowRSVPPathSenderTspecIntServ) HasParameter127Flag() bool {
	return obj.obj.Parameter_127Flag != nil
}

// description is TBD
// SetParameter127Flag sets the PatternFlowRSVPPathSenderTspecIntServParameter127Flag value in the FlowRSVPPathSenderTspecIntServ object
func (obj *flowRSVPPathSenderTspecIntServ) SetParameter127Flag(value PatternFlowRSVPPathSenderTspecIntServParameter127Flag) FlowRSVPPathSenderTspecIntServ {

	obj.parameter_127FlagHolder = nil
	obj.obj.Parameter_127Flag = value.msg()

	return obj
}

// description is TBD
// Parameter127Length returns a PatternFlowRSVPPathSenderTspecIntServParameter127Length
func (obj *flowRSVPPathSenderTspecIntServ) Parameter127Length() PatternFlowRSVPPathSenderTspecIntServParameter127Length {
	if obj.obj.Parameter_127Length == nil {
		obj.obj.Parameter_127Length = NewPatternFlowRSVPPathSenderTspecIntServParameter127Length().msg()
	}
	if obj.parameter_127LengthHolder == nil {
		obj.parameter_127LengthHolder = &patternFlowRSVPPathSenderTspecIntServParameter127Length{obj: obj.obj.Parameter_127Length}
	}
	return obj.parameter_127LengthHolder
}

// description is TBD
// Parameter127Length returns a PatternFlowRSVPPathSenderTspecIntServParameter127Length
func (obj *flowRSVPPathSenderTspecIntServ) HasParameter127Length() bool {
	return obj.obj.Parameter_127Length != nil
}

// description is TBD
// SetParameter127Length sets the PatternFlowRSVPPathSenderTspecIntServParameter127Length value in the FlowRSVPPathSenderTspecIntServ object
func (obj *flowRSVPPathSenderTspecIntServ) SetParameter127Length(value PatternFlowRSVPPathSenderTspecIntServParameter127Length) FlowRSVPPathSenderTspecIntServ {

	obj.parameter_127LengthHolder = nil
	obj.obj.Parameter_127Length = value.msg()

	return obj
}

// Token bucket rate is set to sender's view of its generated traffic.
// TokenBucketRate returns a float32
func (obj *flowRSVPPathSenderTspecIntServ) TokenBucketRate() float32 {

	return *obj.obj.TokenBucketRate

}

// Token bucket rate is set to sender's view of its generated traffic.
// TokenBucketRate returns a float32
func (obj *flowRSVPPathSenderTspecIntServ) HasTokenBucketRate() bool {
	return obj.obj.TokenBucketRate != nil
}

// Token bucket rate is set to sender's view of its generated traffic.
// SetTokenBucketRate sets the float32 value in the FlowRSVPPathSenderTspecIntServ object
func (obj *flowRSVPPathSenderTspecIntServ) SetTokenBucketRate(value float32) FlowRSVPPathSenderTspecIntServ {

	obj.obj.TokenBucketRate = &value
	return obj
}

// Token bucket size is set to sender's view of its generated traffic.
// TokenBucketSize returns a float32
func (obj *flowRSVPPathSenderTspecIntServ) TokenBucketSize() float32 {

	return *obj.obj.TokenBucketSize

}

// Token bucket size is set to sender's view of its generated traffic.
// TokenBucketSize returns a float32
func (obj *flowRSVPPathSenderTspecIntServ) HasTokenBucketSize() bool {
	return obj.obj.TokenBucketSize != nil
}

// Token bucket size is set to sender's view of its generated traffic.
// SetTokenBucketSize sets the float32 value in the FlowRSVPPathSenderTspecIntServ object
func (obj *flowRSVPPathSenderTspecIntServ) SetTokenBucketSize(value float32) FlowRSVPPathSenderTspecIntServ {

	obj.obj.TokenBucketSize = &value
	return obj
}

// The peak rate may be set to the sender's peak traffic generation rate (if known and controlled), the physical interface line rate (if known), or positive infinity (if no better value is available).
// PeakDataRate returns a float32
func (obj *flowRSVPPathSenderTspecIntServ) PeakDataRate() float32 {

	return *obj.obj.PeakDataRate

}

// The peak rate may be set to the sender's peak traffic generation rate (if known and controlled), the physical interface line rate (if known), or positive infinity (if no better value is available).
// PeakDataRate returns a float32
func (obj *flowRSVPPathSenderTspecIntServ) HasPeakDataRate() bool {
	return obj.obj.PeakDataRate != nil
}

// The peak rate may be set to the sender's peak traffic generation rate (if known and controlled), the physical interface line rate (if known), or positive infinity (if no better value is available).
// SetPeakDataRate sets the float32 value in the FlowRSVPPathSenderTspecIntServ object
func (obj *flowRSVPPathSenderTspecIntServ) SetPeakDataRate(value float32) FlowRSVPPathSenderTspecIntServ {

	obj.obj.PeakDataRate = &value
	return obj
}

// description is TBD
// MinimumPolicedUnit returns a PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit
func (obj *flowRSVPPathSenderTspecIntServ) MinimumPolicedUnit() PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit {
	if obj.obj.MinimumPolicedUnit == nil {
		obj.obj.MinimumPolicedUnit = NewPatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit().msg()
	}
	if obj.minimumPolicedUnitHolder == nil {
		obj.minimumPolicedUnitHolder = &patternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit{obj: obj.obj.MinimumPolicedUnit}
	}
	return obj.minimumPolicedUnitHolder
}

// description is TBD
// MinimumPolicedUnit returns a PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit
func (obj *flowRSVPPathSenderTspecIntServ) HasMinimumPolicedUnit() bool {
	return obj.obj.MinimumPolicedUnit != nil
}

// description is TBD
// SetMinimumPolicedUnit sets the PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit value in the FlowRSVPPathSenderTspecIntServ object
func (obj *flowRSVPPathSenderTspecIntServ) SetMinimumPolicedUnit(value PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit) FlowRSVPPathSenderTspecIntServ {

	obj.minimumPolicedUnitHolder = nil
	obj.obj.MinimumPolicedUnit = value.msg()

	return obj
}

// description is TBD
// MaximumPacketSize returns a PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize
func (obj *flowRSVPPathSenderTspecIntServ) MaximumPacketSize() PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize {
	if obj.obj.MaximumPacketSize == nil {
		obj.obj.MaximumPacketSize = NewPatternFlowRSVPPathSenderTspecIntServMaximumPacketSize().msg()
	}
	if obj.maximumPacketSizeHolder == nil {
		obj.maximumPacketSizeHolder = &patternFlowRSVPPathSenderTspecIntServMaximumPacketSize{obj: obj.obj.MaximumPacketSize}
	}
	return obj.maximumPacketSizeHolder
}

// description is TBD
// MaximumPacketSize returns a PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize
func (obj *flowRSVPPathSenderTspecIntServ) HasMaximumPacketSize() bool {
	return obj.obj.MaximumPacketSize != nil
}

// description is TBD
// SetMaximumPacketSize sets the PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize value in the FlowRSVPPathSenderTspecIntServ object
func (obj *flowRSVPPathSenderTspecIntServ) SetMaximumPacketSize(value PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize) FlowRSVPPathSenderTspecIntServ {

	obj.maximumPacketSizeHolder = nil
	obj.obj.MaximumPacketSize = value.msg()

	return obj
}

func (obj *flowRSVPPathSenderTspecIntServ) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Version != nil {

		obj.Version().validateObj(vObj, set_default)
	}

	if obj.obj.Reserved1 != nil {

		obj.Reserved1().validateObj(vObj, set_default)
	}

	if obj.obj.OverallLength != nil {

		obj.OverallLength().validateObj(vObj, set_default)
	}

	if obj.obj.ServiceHeader != nil {

		obj.ServiceHeader().validateObj(vObj, set_default)
	}

	if obj.obj.ZeroBit != nil {

		obj.ZeroBit().validateObj(vObj, set_default)
	}

	if obj.obj.Reserved2 != nil {

		obj.Reserved2().validateObj(vObj, set_default)
	}

	if obj.obj.LengthOfServiceData != nil {

		obj.LengthOfServiceData().validateObj(vObj, set_default)
	}

	if obj.obj.ParameterIdTokenBucketTspec != nil {

		obj.ParameterIdTokenBucketTspec().validateObj(vObj, set_default)
	}

	if obj.obj.Parameter_127Flag != nil {

		obj.Parameter127Flag().validateObj(vObj, set_default)
	}

	if obj.obj.Parameter_127Length != nil {

		obj.Parameter127Length().validateObj(vObj, set_default)
	}

	if obj.obj.MinimumPolicedUnit != nil {

		obj.MinimumPolicedUnit().validateObj(vObj, set_default)
	}

	if obj.obj.MaximumPacketSize != nil {

		obj.MaximumPacketSize().validateObj(vObj, set_default)
	}

}

func (obj *flowRSVPPathSenderTspecIntServ) setDefault() {
	if obj.obj.TokenBucketRate == nil {
		obj.SetTokenBucketRate(0)
	}
	if obj.obj.TokenBucketSize == nil {
		obj.SetTokenBucketSize(0)
	}
	if obj.obj.PeakDataRate == nil {
		obj.SetPeakDataRate(0)
	}

}
