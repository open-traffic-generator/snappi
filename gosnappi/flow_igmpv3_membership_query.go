package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowIgmpv3MembershipQuery *****
type flowIgmpv3MembershipQuery struct {
	validation
	obj                   *otg.FlowIgmpv3MembershipQuery
	marshaller            marshalFlowIgmpv3MembershipQuery
	unMarshaller          unMarshalFlowIgmpv3MembershipQuery
	typeHolder            PatternFlowIgmpv3MembershipQueryType
	maxRespCodeHolder     PatternFlowIgmpv3MembershipQueryMaxRespCode
	checksumHolder        PatternFlowIgmpv3MembershipQueryChecksum
	groupAddressHolder    PatternFlowIgmpv3MembershipQueryGroupAddress
	flagsHolder           PatternFlowIgmpv3MembershipQueryFlags
	sFlagHolder           PatternFlowIgmpv3MembershipQuerySFlag
	qrvHolder             PatternFlowIgmpv3MembershipQueryQrv
	qqicHolder            PatternFlowIgmpv3MembershipQueryQqic
	numberOfSourcesHolder PatternFlowIgmpv3MembershipQueryNumberOfSources
	sourcesHolder         FlowIgmpv3MembershipQueryFlowIgmpv3SourceIter
}

func NewFlowIgmpv3MembershipQuery() FlowIgmpv3MembershipQuery {
	obj := flowIgmpv3MembershipQuery{obj: &otg.FlowIgmpv3MembershipQuery{}}
	obj.setDefault()
	return &obj
}

func (obj *flowIgmpv3MembershipQuery) msg() *otg.FlowIgmpv3MembershipQuery {
	return obj.obj
}

func (obj *flowIgmpv3MembershipQuery) setMsg(msg *otg.FlowIgmpv3MembershipQuery) FlowIgmpv3MembershipQuery {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowIgmpv3MembershipQuery struct {
	obj *flowIgmpv3MembershipQuery
}

type marshalFlowIgmpv3MembershipQuery interface {
	// ToProto marshals FlowIgmpv3MembershipQuery to protobuf object *otg.FlowIgmpv3MembershipQuery
	ToProto() (*otg.FlowIgmpv3MembershipQuery, error)
	// ToPbText marshals FlowIgmpv3MembershipQuery to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowIgmpv3MembershipQuery to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowIgmpv3MembershipQuery to JSON text
	ToJson() (string, error)
}

type unMarshalflowIgmpv3MembershipQuery struct {
	obj *flowIgmpv3MembershipQuery
}

type unMarshalFlowIgmpv3MembershipQuery interface {
	// FromProto unmarshals FlowIgmpv3MembershipQuery from protobuf object *otg.FlowIgmpv3MembershipQuery
	FromProto(msg *otg.FlowIgmpv3MembershipQuery) (FlowIgmpv3MembershipQuery, error)
	// FromPbText unmarshals FlowIgmpv3MembershipQuery from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowIgmpv3MembershipQuery from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowIgmpv3MembershipQuery from JSON text
	FromJson(value string) error
}

func (obj *flowIgmpv3MembershipQuery) Marshal() marshalFlowIgmpv3MembershipQuery {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowIgmpv3MembershipQuery{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowIgmpv3MembershipQuery) Unmarshal() unMarshalFlowIgmpv3MembershipQuery {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowIgmpv3MembershipQuery{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowIgmpv3MembershipQuery) ToProto() (*otg.FlowIgmpv3MembershipQuery, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowIgmpv3MembershipQuery) FromProto(msg *otg.FlowIgmpv3MembershipQuery) (FlowIgmpv3MembershipQuery, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowIgmpv3MembershipQuery) ToPbText() (string, error) {
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

func (m *unMarshalflowIgmpv3MembershipQuery) FromPbText(value string) error {
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

func (m *marshalflowIgmpv3MembershipQuery) ToYaml() (string, error) {
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

func (m *unMarshalflowIgmpv3MembershipQuery) FromYaml(value string) error {
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

func (m *marshalflowIgmpv3MembershipQuery) ToJson() (string, error) {
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

func (m *unMarshalflowIgmpv3MembershipQuery) FromJson(value string) error {
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

func (obj *flowIgmpv3MembershipQuery) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowIgmpv3MembershipQuery) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowIgmpv3MembershipQuery) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowIgmpv3MembershipQuery) Clone() (FlowIgmpv3MembershipQuery, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowIgmpv3MembershipQuery()
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

func (obj *flowIgmpv3MembershipQuery) setNil() {
	obj.typeHolder = nil
	obj.maxRespCodeHolder = nil
	obj.checksumHolder = nil
	obj.groupAddressHolder = nil
	obj.flagsHolder = nil
	obj.sFlagHolder = nil
	obj.qrvHolder = nil
	obj.qqicHolder = nil
	obj.numberOfSourcesHolder = nil
	obj.sourcesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowIgmpv3MembershipQuery is iGMPv3 Membership Query message (RFC 9776 Section 4.1).
type FlowIgmpv3MembershipQuery interface {
	Validation
	// msg marshals FlowIgmpv3MembershipQuery to protobuf object *otg.FlowIgmpv3MembershipQuery
	// and doesn't set defaults
	msg() *otg.FlowIgmpv3MembershipQuery
	// setMsg unmarshals FlowIgmpv3MembershipQuery from protobuf object *otg.FlowIgmpv3MembershipQuery
	// and doesn't set defaults
	setMsg(*otg.FlowIgmpv3MembershipQuery) FlowIgmpv3MembershipQuery
	// provides marshal interface
	Marshal() marshalFlowIgmpv3MembershipQuery
	// provides unmarshal interface
	Unmarshal() unMarshalFlowIgmpv3MembershipQuery
	// validate validates FlowIgmpv3MembershipQuery
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowIgmpv3MembershipQuery, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Type returns PatternFlowIgmpv3MembershipQueryType, set in FlowIgmpv3MembershipQuery.
	// PatternFlowIgmpv3MembershipQueryType is iGMP message type, 17 (0x11) for a Membership Query.
	Type() PatternFlowIgmpv3MembershipQueryType
	// SetType assigns PatternFlowIgmpv3MembershipQueryType provided by user to FlowIgmpv3MembershipQuery.
	// PatternFlowIgmpv3MembershipQueryType is iGMP message type, 17 (0x11) for a Membership Query.
	SetType(value PatternFlowIgmpv3MembershipQueryType) FlowIgmpv3MembershipQuery
	// HasType checks if Type has been set in FlowIgmpv3MembershipQuery
	HasType() bool
	// MaxRespCode returns PatternFlowIgmpv3MembershipQueryMaxRespCode, set in FlowIgmpv3MembershipQuery.
	// PatternFlowIgmpv3MembershipQueryMaxRespCode is maximum time in units of 1/10 second within which hosts must respond to the query, so the default 100 is 10 seconds. Values of 128 or more use the exponent and mantissa encoding of RFC 9776 Section 4.1.1.
	MaxRespCode() PatternFlowIgmpv3MembershipQueryMaxRespCode
	// SetMaxRespCode assigns PatternFlowIgmpv3MembershipQueryMaxRespCode provided by user to FlowIgmpv3MembershipQuery.
	// PatternFlowIgmpv3MembershipQueryMaxRespCode is maximum time in units of 1/10 second within which hosts must respond to the query, so the default 100 is 10 seconds. Values of 128 or more use the exponent and mantissa encoding of RFC 9776 Section 4.1.1.
	SetMaxRespCode(value PatternFlowIgmpv3MembershipQueryMaxRespCode) FlowIgmpv3MembershipQuery
	// HasMaxRespCode checks if MaxRespCode has been set in FlowIgmpv3MembershipQuery
	HasMaxRespCode() bool
	// Checksum returns PatternFlowIgmpv3MembershipQueryChecksum, set in FlowIgmpv3MembershipQuery.
	// PatternFlowIgmpv3MembershipQueryChecksum is checksum of the whole IGMP message, a 16-bit one's complement sum with no pseudo-header.
	Checksum() PatternFlowIgmpv3MembershipQueryChecksum
	// SetChecksum assigns PatternFlowIgmpv3MembershipQueryChecksum provided by user to FlowIgmpv3MembershipQuery.
	// PatternFlowIgmpv3MembershipQueryChecksum is checksum of the whole IGMP message, a 16-bit one's complement sum with no pseudo-header.
	SetChecksum(value PatternFlowIgmpv3MembershipQueryChecksum) FlowIgmpv3MembershipQuery
	// HasChecksum checks if Checksum has been set in FlowIgmpv3MembershipQuery
	HasChecksum() bool
	// GroupAddress returns PatternFlowIgmpv3MembershipQueryGroupAddress, set in FlowIgmpv3MembershipQuery.
	// PatternFlowIgmpv3MembershipQueryGroupAddress is group address, zero for a general query and the group being queried for a group-specific or group-and-source-specific query.
	GroupAddress() PatternFlowIgmpv3MembershipQueryGroupAddress
	// SetGroupAddress assigns PatternFlowIgmpv3MembershipQueryGroupAddress provided by user to FlowIgmpv3MembershipQuery.
	// PatternFlowIgmpv3MembershipQueryGroupAddress is group address, zero for a general query and the group being queried for a group-specific or group-and-source-specific query.
	SetGroupAddress(value PatternFlowIgmpv3MembershipQueryGroupAddress) FlowIgmpv3MembershipQuery
	// HasGroupAddress checks if GroupAddress has been set in FlowIgmpv3MembershipQuery
	HasGroupAddress() bool
	// Flags returns PatternFlowIgmpv3MembershipQueryFlags, set in FlowIgmpv3MembershipQuery.
	// PatternFlowIgmpv3MembershipQueryFlags is flags field of the query, 4 bits that are managed by IANA and are set to zero by a conforming sender.
	Flags() PatternFlowIgmpv3MembershipQueryFlags
	// SetFlags assigns PatternFlowIgmpv3MembershipQueryFlags provided by user to FlowIgmpv3MembershipQuery.
	// PatternFlowIgmpv3MembershipQueryFlags is flags field of the query, 4 bits that are managed by IANA and are set to zero by a conforming sender.
	SetFlags(value PatternFlowIgmpv3MembershipQueryFlags) FlowIgmpv3MembershipQuery
	// HasFlags checks if Flags has been set in FlowIgmpv3MembershipQuery
	HasFlags() bool
	// SFlag returns PatternFlowIgmpv3MembershipQuerySFlag, set in FlowIgmpv3MembershipQuery.
	// PatternFlowIgmpv3MembershipQuerySFlag is suppress router-side processing flag, when set receiving routers do not update their timers.
	SFlag() PatternFlowIgmpv3MembershipQuerySFlag
	// SetSFlag assigns PatternFlowIgmpv3MembershipQuerySFlag provided by user to FlowIgmpv3MembershipQuery.
	// PatternFlowIgmpv3MembershipQuerySFlag is suppress router-side processing flag, when set receiving routers do not update their timers.
	SetSFlag(value PatternFlowIgmpv3MembershipQuerySFlag) FlowIgmpv3MembershipQuery
	// HasSFlag checks if SFlag has been set in FlowIgmpv3MembershipQuery
	HasSFlag() bool
	// Qrv returns PatternFlowIgmpv3MembershipQueryQrv, set in FlowIgmpv3MembershipQuery.
	// PatternFlowIgmpv3MembershipQueryQrv is querier's robustness variable. Zero means the receiver uses its own value, otherwise it is the value used by the querier.
	Qrv() PatternFlowIgmpv3MembershipQueryQrv
	// SetQrv assigns PatternFlowIgmpv3MembershipQueryQrv provided by user to FlowIgmpv3MembershipQuery.
	// PatternFlowIgmpv3MembershipQueryQrv is querier's robustness variable. Zero means the receiver uses its own value, otherwise it is the value used by the querier.
	SetQrv(value PatternFlowIgmpv3MembershipQueryQrv) FlowIgmpv3MembershipQuery
	// HasQrv checks if Qrv has been set in FlowIgmpv3MembershipQuery
	HasQrv() bool
	// Qqic returns PatternFlowIgmpv3MembershipQueryQqic, set in FlowIgmpv3MembershipQuery.
	// PatternFlowIgmpv3MembershipQueryQqic is querier's query interval code in seconds, so the default 125 is 125 seconds. Values of 128 or more use the exponent and mantissa encoding of RFC 9776 Section 4.1.7.
	Qqic() PatternFlowIgmpv3MembershipQueryQqic
	// SetQqic assigns PatternFlowIgmpv3MembershipQueryQqic provided by user to FlowIgmpv3MembershipQuery.
	// PatternFlowIgmpv3MembershipQueryQqic is querier's query interval code in seconds, so the default 125 is 125 seconds. Values of 128 or more use the exponent and mantissa encoding of RFC 9776 Section 4.1.7.
	SetQqic(value PatternFlowIgmpv3MembershipQueryQqic) FlowIgmpv3MembershipQuery
	// HasQqic checks if Qqic has been set in FlowIgmpv3MembershipQuery
	HasQqic() bool
	// NumberOfSources returns PatternFlowIgmpv3MembershipQueryNumberOfSources, set in FlowIgmpv3MembershipQuery.
	// PatternFlowIgmpv3MembershipQueryNumberOfSources is number of source addresses in the query, zero for a general and a group-specific query. When auto, it is set to the number of entries in sources.
	NumberOfSources() PatternFlowIgmpv3MembershipQueryNumberOfSources
	// SetNumberOfSources assigns PatternFlowIgmpv3MembershipQueryNumberOfSources provided by user to FlowIgmpv3MembershipQuery.
	// PatternFlowIgmpv3MembershipQueryNumberOfSources is number of source addresses in the query, zero for a general and a group-specific query. When auto, it is set to the number of entries in sources.
	SetNumberOfSources(value PatternFlowIgmpv3MembershipQueryNumberOfSources) FlowIgmpv3MembershipQuery
	// HasNumberOfSources checks if NumberOfSources has been set in FlowIgmpv3MembershipQuery
	HasNumberOfSources() bool
	// Sources returns FlowIgmpv3MembershipQueryFlowIgmpv3SourceIterIter, set in FlowIgmpv3MembershipQuery
	Sources() FlowIgmpv3MembershipQueryFlowIgmpv3SourceIter
	setNil()
}

// description is TBD
// Type returns a PatternFlowIgmpv3MembershipQueryType
func (obj *flowIgmpv3MembershipQuery) Type() PatternFlowIgmpv3MembershipQueryType {
	if obj.obj.Type == nil {
		obj.obj.Type = NewPatternFlowIgmpv3MembershipQueryType().msg()
	}
	if obj.typeHolder == nil {
		obj.typeHolder = &patternFlowIgmpv3MembershipQueryType{obj: obj.obj.Type}
	}
	return obj.typeHolder
}

// description is TBD
// Type returns a PatternFlowIgmpv3MembershipQueryType
func (obj *flowIgmpv3MembershipQuery) HasType() bool {
	return obj.obj.Type != nil
}

// description is TBD
// SetType sets the PatternFlowIgmpv3MembershipQueryType value in the FlowIgmpv3MembershipQuery object
func (obj *flowIgmpv3MembershipQuery) SetType(value PatternFlowIgmpv3MembershipQueryType) FlowIgmpv3MembershipQuery {

	obj.typeHolder = nil
	obj.obj.Type = value.msg()

	return obj
}

// description is TBD
// MaxRespCode returns a PatternFlowIgmpv3MembershipQueryMaxRespCode
func (obj *flowIgmpv3MembershipQuery) MaxRespCode() PatternFlowIgmpv3MembershipQueryMaxRespCode {
	if obj.obj.MaxRespCode == nil {
		obj.obj.MaxRespCode = NewPatternFlowIgmpv3MembershipQueryMaxRespCode().msg()
	}
	if obj.maxRespCodeHolder == nil {
		obj.maxRespCodeHolder = &patternFlowIgmpv3MembershipQueryMaxRespCode{obj: obj.obj.MaxRespCode}
	}
	return obj.maxRespCodeHolder
}

// description is TBD
// MaxRespCode returns a PatternFlowIgmpv3MembershipQueryMaxRespCode
func (obj *flowIgmpv3MembershipQuery) HasMaxRespCode() bool {
	return obj.obj.MaxRespCode != nil
}

// description is TBD
// SetMaxRespCode sets the PatternFlowIgmpv3MembershipQueryMaxRespCode value in the FlowIgmpv3MembershipQuery object
func (obj *flowIgmpv3MembershipQuery) SetMaxRespCode(value PatternFlowIgmpv3MembershipQueryMaxRespCode) FlowIgmpv3MembershipQuery {

	obj.maxRespCodeHolder = nil
	obj.obj.MaxRespCode = value.msg()

	return obj
}

// description is TBD
// Checksum returns a PatternFlowIgmpv3MembershipQueryChecksum
func (obj *flowIgmpv3MembershipQuery) Checksum() PatternFlowIgmpv3MembershipQueryChecksum {
	if obj.obj.Checksum == nil {
		obj.obj.Checksum = NewPatternFlowIgmpv3MembershipQueryChecksum().msg()
	}
	if obj.checksumHolder == nil {
		obj.checksumHolder = &patternFlowIgmpv3MembershipQueryChecksum{obj: obj.obj.Checksum}
	}
	return obj.checksumHolder
}

// description is TBD
// Checksum returns a PatternFlowIgmpv3MembershipQueryChecksum
func (obj *flowIgmpv3MembershipQuery) HasChecksum() bool {
	return obj.obj.Checksum != nil
}

// description is TBD
// SetChecksum sets the PatternFlowIgmpv3MembershipQueryChecksum value in the FlowIgmpv3MembershipQuery object
func (obj *flowIgmpv3MembershipQuery) SetChecksum(value PatternFlowIgmpv3MembershipQueryChecksum) FlowIgmpv3MembershipQuery {

	obj.checksumHolder = nil
	obj.obj.Checksum = value.msg()

	return obj
}

// description is TBD
// GroupAddress returns a PatternFlowIgmpv3MembershipQueryGroupAddress
func (obj *flowIgmpv3MembershipQuery) GroupAddress() PatternFlowIgmpv3MembershipQueryGroupAddress {
	if obj.obj.GroupAddress == nil {
		obj.obj.GroupAddress = NewPatternFlowIgmpv3MembershipQueryGroupAddress().msg()
	}
	if obj.groupAddressHolder == nil {
		obj.groupAddressHolder = &patternFlowIgmpv3MembershipQueryGroupAddress{obj: obj.obj.GroupAddress}
	}
	return obj.groupAddressHolder
}

// description is TBD
// GroupAddress returns a PatternFlowIgmpv3MembershipQueryGroupAddress
func (obj *flowIgmpv3MembershipQuery) HasGroupAddress() bool {
	return obj.obj.GroupAddress != nil
}

// description is TBD
// SetGroupAddress sets the PatternFlowIgmpv3MembershipQueryGroupAddress value in the FlowIgmpv3MembershipQuery object
func (obj *flowIgmpv3MembershipQuery) SetGroupAddress(value PatternFlowIgmpv3MembershipQueryGroupAddress) FlowIgmpv3MembershipQuery {

	obj.groupAddressHolder = nil
	obj.obj.GroupAddress = value.msg()

	return obj
}

// description is TBD
// Flags returns a PatternFlowIgmpv3MembershipQueryFlags
func (obj *flowIgmpv3MembershipQuery) Flags() PatternFlowIgmpv3MembershipQueryFlags {
	if obj.obj.Flags == nil {
		obj.obj.Flags = NewPatternFlowIgmpv3MembershipQueryFlags().msg()
	}
	if obj.flagsHolder == nil {
		obj.flagsHolder = &patternFlowIgmpv3MembershipQueryFlags{obj: obj.obj.Flags}
	}
	return obj.flagsHolder
}

// description is TBD
// Flags returns a PatternFlowIgmpv3MembershipQueryFlags
func (obj *flowIgmpv3MembershipQuery) HasFlags() bool {
	return obj.obj.Flags != nil
}

// description is TBD
// SetFlags sets the PatternFlowIgmpv3MembershipQueryFlags value in the FlowIgmpv3MembershipQuery object
func (obj *flowIgmpv3MembershipQuery) SetFlags(value PatternFlowIgmpv3MembershipQueryFlags) FlowIgmpv3MembershipQuery {

	obj.flagsHolder = nil
	obj.obj.Flags = value.msg()

	return obj
}

// description is TBD
// SFlag returns a PatternFlowIgmpv3MembershipQuerySFlag
func (obj *flowIgmpv3MembershipQuery) SFlag() PatternFlowIgmpv3MembershipQuerySFlag {
	if obj.obj.SFlag == nil {
		obj.obj.SFlag = NewPatternFlowIgmpv3MembershipQuerySFlag().msg()
	}
	if obj.sFlagHolder == nil {
		obj.sFlagHolder = &patternFlowIgmpv3MembershipQuerySFlag{obj: obj.obj.SFlag}
	}
	return obj.sFlagHolder
}

// description is TBD
// SFlag returns a PatternFlowIgmpv3MembershipQuerySFlag
func (obj *flowIgmpv3MembershipQuery) HasSFlag() bool {
	return obj.obj.SFlag != nil
}

// description is TBD
// SetSFlag sets the PatternFlowIgmpv3MembershipQuerySFlag value in the FlowIgmpv3MembershipQuery object
func (obj *flowIgmpv3MembershipQuery) SetSFlag(value PatternFlowIgmpv3MembershipQuerySFlag) FlowIgmpv3MembershipQuery {

	obj.sFlagHolder = nil
	obj.obj.SFlag = value.msg()

	return obj
}

// description is TBD
// Qrv returns a PatternFlowIgmpv3MembershipQueryQrv
func (obj *flowIgmpv3MembershipQuery) Qrv() PatternFlowIgmpv3MembershipQueryQrv {
	if obj.obj.Qrv == nil {
		obj.obj.Qrv = NewPatternFlowIgmpv3MembershipQueryQrv().msg()
	}
	if obj.qrvHolder == nil {
		obj.qrvHolder = &patternFlowIgmpv3MembershipQueryQrv{obj: obj.obj.Qrv}
	}
	return obj.qrvHolder
}

// description is TBD
// Qrv returns a PatternFlowIgmpv3MembershipQueryQrv
func (obj *flowIgmpv3MembershipQuery) HasQrv() bool {
	return obj.obj.Qrv != nil
}

// description is TBD
// SetQrv sets the PatternFlowIgmpv3MembershipQueryQrv value in the FlowIgmpv3MembershipQuery object
func (obj *flowIgmpv3MembershipQuery) SetQrv(value PatternFlowIgmpv3MembershipQueryQrv) FlowIgmpv3MembershipQuery {

	obj.qrvHolder = nil
	obj.obj.Qrv = value.msg()

	return obj
}

// description is TBD
// Qqic returns a PatternFlowIgmpv3MembershipQueryQqic
func (obj *flowIgmpv3MembershipQuery) Qqic() PatternFlowIgmpv3MembershipQueryQqic {
	if obj.obj.Qqic == nil {
		obj.obj.Qqic = NewPatternFlowIgmpv3MembershipQueryQqic().msg()
	}
	if obj.qqicHolder == nil {
		obj.qqicHolder = &patternFlowIgmpv3MembershipQueryQqic{obj: obj.obj.Qqic}
	}
	return obj.qqicHolder
}

// description is TBD
// Qqic returns a PatternFlowIgmpv3MembershipQueryQqic
func (obj *flowIgmpv3MembershipQuery) HasQqic() bool {
	return obj.obj.Qqic != nil
}

// description is TBD
// SetQqic sets the PatternFlowIgmpv3MembershipQueryQqic value in the FlowIgmpv3MembershipQuery object
func (obj *flowIgmpv3MembershipQuery) SetQqic(value PatternFlowIgmpv3MembershipQueryQqic) FlowIgmpv3MembershipQuery {

	obj.qqicHolder = nil
	obj.obj.Qqic = value.msg()

	return obj
}

// description is TBD
// NumberOfSources returns a PatternFlowIgmpv3MembershipQueryNumberOfSources
func (obj *flowIgmpv3MembershipQuery) NumberOfSources() PatternFlowIgmpv3MembershipQueryNumberOfSources {
	if obj.obj.NumberOfSources == nil {
		obj.obj.NumberOfSources = NewPatternFlowIgmpv3MembershipQueryNumberOfSources().msg()
	}
	if obj.numberOfSourcesHolder == nil {
		obj.numberOfSourcesHolder = &patternFlowIgmpv3MembershipQueryNumberOfSources{obj: obj.obj.NumberOfSources}
	}
	return obj.numberOfSourcesHolder
}

// description is TBD
// NumberOfSources returns a PatternFlowIgmpv3MembershipQueryNumberOfSources
func (obj *flowIgmpv3MembershipQuery) HasNumberOfSources() bool {
	return obj.obj.NumberOfSources != nil
}

// description is TBD
// SetNumberOfSources sets the PatternFlowIgmpv3MembershipQueryNumberOfSources value in the FlowIgmpv3MembershipQuery object
func (obj *flowIgmpv3MembershipQuery) SetNumberOfSources(value PatternFlowIgmpv3MembershipQueryNumberOfSources) FlowIgmpv3MembershipQuery {

	obj.numberOfSourcesHolder = nil
	obj.obj.NumberOfSources = value.msg()

	return obj
}

// List of unicast source addresses of a group-and-source-specific query.
// Sources returns a []FlowIgmpv3Source
func (obj *flowIgmpv3MembershipQuery) Sources() FlowIgmpv3MembershipQueryFlowIgmpv3SourceIter {
	if len(obj.obj.Sources) == 0 {
		obj.obj.Sources = []*otg.FlowIgmpv3Source{}
	}
	if obj.sourcesHolder == nil {
		obj.sourcesHolder = newFlowIgmpv3MembershipQueryFlowIgmpv3SourceIter(&obj.obj.Sources).setMsg(obj)
	}
	return obj.sourcesHolder
}

type flowIgmpv3MembershipQueryFlowIgmpv3SourceIter struct {
	obj                   *flowIgmpv3MembershipQuery
	flowIgmpv3SourceSlice []FlowIgmpv3Source
	fieldPtr              *[]*otg.FlowIgmpv3Source
}

func newFlowIgmpv3MembershipQueryFlowIgmpv3SourceIter(ptr *[]*otg.FlowIgmpv3Source) FlowIgmpv3MembershipQueryFlowIgmpv3SourceIter {
	return &flowIgmpv3MembershipQueryFlowIgmpv3SourceIter{fieldPtr: ptr}
}

type FlowIgmpv3MembershipQueryFlowIgmpv3SourceIter interface {
	setMsg(*flowIgmpv3MembershipQuery) FlowIgmpv3MembershipQueryFlowIgmpv3SourceIter
	Items() []FlowIgmpv3Source
	Add() FlowIgmpv3Source
	Append(items ...FlowIgmpv3Source) FlowIgmpv3MembershipQueryFlowIgmpv3SourceIter
	Set(index int, newObj FlowIgmpv3Source) FlowIgmpv3MembershipQueryFlowIgmpv3SourceIter
	Clear() FlowIgmpv3MembershipQueryFlowIgmpv3SourceIter
	clearHolderSlice() FlowIgmpv3MembershipQueryFlowIgmpv3SourceIter
	appendHolderSlice(item FlowIgmpv3Source) FlowIgmpv3MembershipQueryFlowIgmpv3SourceIter
}

func (obj *flowIgmpv3MembershipQueryFlowIgmpv3SourceIter) setMsg(msg *flowIgmpv3MembershipQuery) FlowIgmpv3MembershipQueryFlowIgmpv3SourceIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&flowIgmpv3Source{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *flowIgmpv3MembershipQueryFlowIgmpv3SourceIter) Items() []FlowIgmpv3Source {
	return obj.flowIgmpv3SourceSlice
}

func (obj *flowIgmpv3MembershipQueryFlowIgmpv3SourceIter) Add() FlowIgmpv3Source {
	newObj := &otg.FlowIgmpv3Source{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &flowIgmpv3Source{obj: newObj}
	newLibObj.setDefault()
	obj.flowIgmpv3SourceSlice = append(obj.flowIgmpv3SourceSlice, newLibObj)
	return newLibObj
}

func (obj *flowIgmpv3MembershipQueryFlowIgmpv3SourceIter) Append(items ...FlowIgmpv3Source) FlowIgmpv3MembershipQueryFlowIgmpv3SourceIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.flowIgmpv3SourceSlice = append(obj.flowIgmpv3SourceSlice, item)
	}
	return obj
}

func (obj *flowIgmpv3MembershipQueryFlowIgmpv3SourceIter) Set(index int, newObj FlowIgmpv3Source) FlowIgmpv3MembershipQueryFlowIgmpv3SourceIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.flowIgmpv3SourceSlice[index] = newObj
	return obj
}
func (obj *flowIgmpv3MembershipQueryFlowIgmpv3SourceIter) Clear() FlowIgmpv3MembershipQueryFlowIgmpv3SourceIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.FlowIgmpv3Source{}
		obj.flowIgmpv3SourceSlice = []FlowIgmpv3Source{}
	}
	return obj
}
func (obj *flowIgmpv3MembershipQueryFlowIgmpv3SourceIter) clearHolderSlice() FlowIgmpv3MembershipQueryFlowIgmpv3SourceIter {
	if len(obj.flowIgmpv3SourceSlice) > 0 {
		obj.flowIgmpv3SourceSlice = []FlowIgmpv3Source{}
	}
	return obj
}
func (obj *flowIgmpv3MembershipQueryFlowIgmpv3SourceIter) appendHolderSlice(item FlowIgmpv3Source) FlowIgmpv3MembershipQueryFlowIgmpv3SourceIter {
	obj.flowIgmpv3SourceSlice = append(obj.flowIgmpv3SourceSlice, item)
	return obj
}

func (obj *flowIgmpv3MembershipQuery) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Type != nil {

		obj.Type().validateObj(vObj, set_default)
	}

	if obj.obj.MaxRespCode != nil {

		obj.MaxRespCode().validateObj(vObj, set_default)
	}

	if obj.obj.Checksum != nil {

		obj.Checksum().validateObj(vObj, set_default)
	}

	if obj.obj.GroupAddress != nil {

		obj.GroupAddress().validateObj(vObj, set_default)
	}

	if obj.obj.Flags != nil {

		obj.Flags().validateObj(vObj, set_default)
	}

	if obj.obj.SFlag != nil {

		obj.SFlag().validateObj(vObj, set_default)
	}

	if obj.obj.Qrv != nil {

		obj.Qrv().validateObj(vObj, set_default)
	}

	if obj.obj.Qqic != nil {

		obj.Qqic().validateObj(vObj, set_default)
	}

	if obj.obj.NumberOfSources != nil {

		obj.NumberOfSources().validateObj(vObj, set_default)
	}

	if len(obj.obj.Sources) != 0 {

		if set_default {
			obj.Sources().clearHolderSlice()
			for _, item := range obj.obj.Sources {
				obj.Sources().appendHolderSlice(&flowIgmpv3Source{obj: item})
			}
		}
		for _, item := range obj.Sources().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *flowIgmpv3MembershipQuery) setDefault() {

}
