package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowIpv4 *****
type flowIpv4 struct {
	validation
	obj                  *otg.FlowIpv4
	marshaller           marshalFlowIpv4
	unMarshaller         unMarshalFlowIpv4
	versionHolder        PatternFlowIpv4Version
	headerLengthHolder   PatternFlowIpv4HeaderLength
	priorityHolder       FlowIpv4Priority
	totalLengthHolder    PatternFlowIpv4TotalLength
	identificationHolder PatternFlowIpv4Identification
	reservedHolder       PatternFlowIpv4Reserved
	dontFragmentHolder   PatternFlowIpv4DontFragment
	moreFragmentsHolder  PatternFlowIpv4MoreFragments
	fragmentOffsetHolder PatternFlowIpv4FragmentOffset
	timeToLiveHolder     PatternFlowIpv4TimeToLive
	protocolHolder       PatternFlowIpv4Protocol
	headerChecksumHolder PatternFlowIpv4HeaderChecksum
	srcHolder            PatternFlowIpv4Src
	dstHolder            PatternFlowIpv4Dst
	optionsHolder        FlowIpv4FlowIpv4OptionsIter
}

func NewFlowIpv4() FlowIpv4 {
	obj := flowIpv4{obj: &otg.FlowIpv4{}}
	obj.setDefault()
	return &obj
}

func (obj *flowIpv4) msg() *otg.FlowIpv4 {
	return obj.obj
}

func (obj *flowIpv4) setMsg(msg *otg.FlowIpv4) FlowIpv4 {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowIpv4 struct {
	obj *flowIpv4
}

type marshalFlowIpv4 interface {
	// ToProto marshals FlowIpv4 to protobuf object *otg.FlowIpv4
	ToProto() (*otg.FlowIpv4, error)
	// ToPbText marshals FlowIpv4 to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowIpv4 to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowIpv4 to JSON text
	ToJson() (string, error)
}

type unMarshalflowIpv4 struct {
	obj *flowIpv4
}

type unMarshalFlowIpv4 interface {
	// FromProto unmarshals FlowIpv4 from protobuf object *otg.FlowIpv4
	FromProto(msg *otg.FlowIpv4) (FlowIpv4, error)
	// FromPbText unmarshals FlowIpv4 from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowIpv4 from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowIpv4 from JSON text
	FromJson(value string) error
}

func (obj *flowIpv4) Marshal() marshalFlowIpv4 {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowIpv4{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowIpv4) Unmarshal() unMarshalFlowIpv4 {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowIpv4{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowIpv4) ToProto() (*otg.FlowIpv4, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowIpv4) FromProto(msg *otg.FlowIpv4) (FlowIpv4, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowIpv4) ToPbText() (string, error) {
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

func (m *unMarshalflowIpv4) FromPbText(value string) error {
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

func (m *marshalflowIpv4) ToYaml() (string, error) {
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

func (m *unMarshalflowIpv4) FromYaml(value string) error {
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

func (m *marshalflowIpv4) ToJson() (string, error) {
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

func (m *unMarshalflowIpv4) FromJson(value string) error {
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

func (obj *flowIpv4) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowIpv4) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowIpv4) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowIpv4) Clone() (FlowIpv4, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowIpv4()
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

func (obj *flowIpv4) setNil() {
	obj.versionHolder = nil
	obj.headerLengthHolder = nil
	obj.priorityHolder = nil
	obj.totalLengthHolder = nil
	obj.identificationHolder = nil
	obj.reservedHolder = nil
	obj.dontFragmentHolder = nil
	obj.moreFragmentsHolder = nil
	obj.fragmentOffsetHolder = nil
	obj.timeToLiveHolder = nil
	obj.protocolHolder = nil
	obj.headerChecksumHolder = nil
	obj.srcHolder = nil
	obj.dstHolder = nil
	obj.optionsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowIpv4 is iPv4 packet header
type FlowIpv4 interface {
	Validation
	// msg marshals FlowIpv4 to protobuf object *otg.FlowIpv4
	// and doesn't set defaults
	msg() *otg.FlowIpv4
	// setMsg unmarshals FlowIpv4 from protobuf object *otg.FlowIpv4
	// and doesn't set defaults
	setMsg(*otg.FlowIpv4) FlowIpv4
	// provides marshal interface
	Marshal() marshalFlowIpv4
	// provides unmarshal interface
	Unmarshal() unMarshalFlowIpv4
	// validate validates FlowIpv4
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowIpv4, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Version returns PatternFlowIpv4Version, set in FlowIpv4.
	// PatternFlowIpv4Version is version
	Version() PatternFlowIpv4Version
	// SetVersion assigns PatternFlowIpv4Version provided by user to FlowIpv4.
	// PatternFlowIpv4Version is version
	SetVersion(value PatternFlowIpv4Version) FlowIpv4
	// HasVersion checks if Version has been set in FlowIpv4
	HasVersion() bool
	// HeaderLength returns PatternFlowIpv4HeaderLength, set in FlowIpv4.
	// PatternFlowIpv4HeaderLength is header length
	HeaderLength() PatternFlowIpv4HeaderLength
	// SetHeaderLength assigns PatternFlowIpv4HeaderLength provided by user to FlowIpv4.
	// PatternFlowIpv4HeaderLength is header length
	SetHeaderLength(value PatternFlowIpv4HeaderLength) FlowIpv4
	// HasHeaderLength checks if HeaderLength has been set in FlowIpv4
	HasHeaderLength() bool
	// Priority returns FlowIpv4Priority, set in FlowIpv4.
	// FlowIpv4Priority is a container for ipv4 raw, tos, dscp ip priorities.
	Priority() FlowIpv4Priority
	// SetPriority assigns FlowIpv4Priority provided by user to FlowIpv4.
	// FlowIpv4Priority is a container for ipv4 raw, tos, dscp ip priorities.
	SetPriority(value FlowIpv4Priority) FlowIpv4
	// HasPriority checks if Priority has been set in FlowIpv4
	HasPriority() bool
	// TotalLength returns PatternFlowIpv4TotalLength, set in FlowIpv4.
	// PatternFlowIpv4TotalLength is total length
	TotalLength() PatternFlowIpv4TotalLength
	// SetTotalLength assigns PatternFlowIpv4TotalLength provided by user to FlowIpv4.
	// PatternFlowIpv4TotalLength is total length
	SetTotalLength(value PatternFlowIpv4TotalLength) FlowIpv4
	// HasTotalLength checks if TotalLength has been set in FlowIpv4
	HasTotalLength() bool
	// Identification returns PatternFlowIpv4Identification, set in FlowIpv4.
	// PatternFlowIpv4Identification is identification
	Identification() PatternFlowIpv4Identification
	// SetIdentification assigns PatternFlowIpv4Identification provided by user to FlowIpv4.
	// PatternFlowIpv4Identification is identification
	SetIdentification(value PatternFlowIpv4Identification) FlowIpv4
	// HasIdentification checks if Identification has been set in FlowIpv4
	HasIdentification() bool
	// Reserved returns PatternFlowIpv4Reserved, set in FlowIpv4.
	// PatternFlowIpv4Reserved is reserved flag.
	Reserved() PatternFlowIpv4Reserved
	// SetReserved assigns PatternFlowIpv4Reserved provided by user to FlowIpv4.
	// PatternFlowIpv4Reserved is reserved flag.
	SetReserved(value PatternFlowIpv4Reserved) FlowIpv4
	// HasReserved checks if Reserved has been set in FlowIpv4
	HasReserved() bool
	// DontFragment returns PatternFlowIpv4DontFragment, set in FlowIpv4.
	// PatternFlowIpv4DontFragment is dont fragment flag If the dont_fragment flag is set and fragmentation is required to route the packet then the packet is dropped.
	DontFragment() PatternFlowIpv4DontFragment
	// SetDontFragment assigns PatternFlowIpv4DontFragment provided by user to FlowIpv4.
	// PatternFlowIpv4DontFragment is dont fragment flag If the dont_fragment flag is set and fragmentation is required to route the packet then the packet is dropped.
	SetDontFragment(value PatternFlowIpv4DontFragment) FlowIpv4
	// HasDontFragment checks if DontFragment has been set in FlowIpv4
	HasDontFragment() bool
	// MoreFragments returns PatternFlowIpv4MoreFragments, set in FlowIpv4.
	// PatternFlowIpv4MoreFragments is more fragments flag
	MoreFragments() PatternFlowIpv4MoreFragments
	// SetMoreFragments assigns PatternFlowIpv4MoreFragments provided by user to FlowIpv4.
	// PatternFlowIpv4MoreFragments is more fragments flag
	SetMoreFragments(value PatternFlowIpv4MoreFragments) FlowIpv4
	// HasMoreFragments checks if MoreFragments has been set in FlowIpv4
	HasMoreFragments() bool
	// FragmentOffset returns PatternFlowIpv4FragmentOffset, set in FlowIpv4.
	// PatternFlowIpv4FragmentOffset is fragment offset
	FragmentOffset() PatternFlowIpv4FragmentOffset
	// SetFragmentOffset assigns PatternFlowIpv4FragmentOffset provided by user to FlowIpv4.
	// PatternFlowIpv4FragmentOffset is fragment offset
	SetFragmentOffset(value PatternFlowIpv4FragmentOffset) FlowIpv4
	// HasFragmentOffset checks if FragmentOffset has been set in FlowIpv4
	HasFragmentOffset() bool
	// TimeToLive returns PatternFlowIpv4TimeToLive, set in FlowIpv4.
	// PatternFlowIpv4TimeToLive is time to live
	TimeToLive() PatternFlowIpv4TimeToLive
	// SetTimeToLive assigns PatternFlowIpv4TimeToLive provided by user to FlowIpv4.
	// PatternFlowIpv4TimeToLive is time to live
	SetTimeToLive(value PatternFlowIpv4TimeToLive) FlowIpv4
	// HasTimeToLive checks if TimeToLive has been set in FlowIpv4
	HasTimeToLive() bool
	// Protocol returns PatternFlowIpv4Protocol, set in FlowIpv4.
	// PatternFlowIpv4Protocol is protocol, default is 61 any host internal protocol
	Protocol() PatternFlowIpv4Protocol
	// SetProtocol assigns PatternFlowIpv4Protocol provided by user to FlowIpv4.
	// PatternFlowIpv4Protocol is protocol, default is 61 any host internal protocol
	SetProtocol(value PatternFlowIpv4Protocol) FlowIpv4
	// HasProtocol checks if Protocol has been set in FlowIpv4
	HasProtocol() bool
	// HeaderChecksum returns PatternFlowIpv4HeaderChecksum, set in FlowIpv4.
	// PatternFlowIpv4HeaderChecksum is header checksum
	HeaderChecksum() PatternFlowIpv4HeaderChecksum
	// SetHeaderChecksum assigns PatternFlowIpv4HeaderChecksum provided by user to FlowIpv4.
	// PatternFlowIpv4HeaderChecksum is header checksum
	SetHeaderChecksum(value PatternFlowIpv4HeaderChecksum) FlowIpv4
	// HasHeaderChecksum checks if HeaderChecksum has been set in FlowIpv4
	HasHeaderChecksum() bool
	// Src returns PatternFlowIpv4Src, set in FlowIpv4.
	// PatternFlowIpv4Src is source address
	Src() PatternFlowIpv4Src
	// SetSrc assigns PatternFlowIpv4Src provided by user to FlowIpv4.
	// PatternFlowIpv4Src is source address
	SetSrc(value PatternFlowIpv4Src) FlowIpv4
	// HasSrc checks if Src has been set in FlowIpv4
	HasSrc() bool
	// Dst returns PatternFlowIpv4Dst, set in FlowIpv4.
	// PatternFlowIpv4Dst is destination address
	Dst() PatternFlowIpv4Dst
	// SetDst assigns PatternFlowIpv4Dst provided by user to FlowIpv4.
	// PatternFlowIpv4Dst is destination address
	SetDst(value PatternFlowIpv4Dst) FlowIpv4
	// HasDst checks if Dst has been set in FlowIpv4
	HasDst() bool
	// Options returns FlowIpv4FlowIpv4OptionsIterIter, set in FlowIpv4
	Options() FlowIpv4FlowIpv4OptionsIter
	setNil()
}

// description is TBD
// Version returns a PatternFlowIpv4Version
func (obj *flowIpv4) Version() PatternFlowIpv4Version {
	if obj.obj.Version == nil {
		obj.obj.Version = NewPatternFlowIpv4Version().msg()
	}
	if obj.versionHolder == nil {
		obj.versionHolder = &patternFlowIpv4Version{obj: obj.obj.Version}
	}
	return obj.versionHolder
}

// description is TBD
// Version returns a PatternFlowIpv4Version
func (obj *flowIpv4) HasVersion() bool {
	return obj.obj.Version != nil
}

// description is TBD
// SetVersion sets the PatternFlowIpv4Version value in the FlowIpv4 object
func (obj *flowIpv4) SetVersion(value PatternFlowIpv4Version) FlowIpv4 {

	obj.versionHolder = nil
	obj.obj.Version = value.msg()

	return obj
}

// description is TBD
// HeaderLength returns a PatternFlowIpv4HeaderLength
func (obj *flowIpv4) HeaderLength() PatternFlowIpv4HeaderLength {
	if obj.obj.HeaderLength == nil {
		obj.obj.HeaderLength = NewPatternFlowIpv4HeaderLength().msg()
	}
	if obj.headerLengthHolder == nil {
		obj.headerLengthHolder = &patternFlowIpv4HeaderLength{obj: obj.obj.HeaderLength}
	}
	return obj.headerLengthHolder
}

// description is TBD
// HeaderLength returns a PatternFlowIpv4HeaderLength
func (obj *flowIpv4) HasHeaderLength() bool {
	return obj.obj.HeaderLength != nil
}

// description is TBD
// SetHeaderLength sets the PatternFlowIpv4HeaderLength value in the FlowIpv4 object
func (obj *flowIpv4) SetHeaderLength(value PatternFlowIpv4HeaderLength) FlowIpv4 {

	obj.headerLengthHolder = nil
	obj.obj.HeaderLength = value.msg()

	return obj
}

// description is TBD
// Priority returns a FlowIpv4Priority
func (obj *flowIpv4) Priority() FlowIpv4Priority {
	if obj.obj.Priority == nil {
		obj.obj.Priority = NewFlowIpv4Priority().msg()
	}
	if obj.priorityHolder == nil {
		obj.priorityHolder = &flowIpv4Priority{obj: obj.obj.Priority}
	}
	return obj.priorityHolder
}

// description is TBD
// Priority returns a FlowIpv4Priority
func (obj *flowIpv4) HasPriority() bool {
	return obj.obj.Priority != nil
}

// description is TBD
// SetPriority sets the FlowIpv4Priority value in the FlowIpv4 object
func (obj *flowIpv4) SetPriority(value FlowIpv4Priority) FlowIpv4 {

	obj.priorityHolder = nil
	obj.obj.Priority = value.msg()

	return obj
}

// description is TBD
// TotalLength returns a PatternFlowIpv4TotalLength
func (obj *flowIpv4) TotalLength() PatternFlowIpv4TotalLength {
	if obj.obj.TotalLength == nil {
		obj.obj.TotalLength = NewPatternFlowIpv4TotalLength().msg()
	}
	if obj.totalLengthHolder == nil {
		obj.totalLengthHolder = &patternFlowIpv4TotalLength{obj: obj.obj.TotalLength}
	}
	return obj.totalLengthHolder
}

// description is TBD
// TotalLength returns a PatternFlowIpv4TotalLength
func (obj *flowIpv4) HasTotalLength() bool {
	return obj.obj.TotalLength != nil
}

// description is TBD
// SetTotalLength sets the PatternFlowIpv4TotalLength value in the FlowIpv4 object
func (obj *flowIpv4) SetTotalLength(value PatternFlowIpv4TotalLength) FlowIpv4 {

	obj.totalLengthHolder = nil
	obj.obj.TotalLength = value.msg()

	return obj
}

// description is TBD
// Identification returns a PatternFlowIpv4Identification
func (obj *flowIpv4) Identification() PatternFlowIpv4Identification {
	if obj.obj.Identification == nil {
		obj.obj.Identification = NewPatternFlowIpv4Identification().msg()
	}
	if obj.identificationHolder == nil {
		obj.identificationHolder = &patternFlowIpv4Identification{obj: obj.obj.Identification}
	}
	return obj.identificationHolder
}

// description is TBD
// Identification returns a PatternFlowIpv4Identification
func (obj *flowIpv4) HasIdentification() bool {
	return obj.obj.Identification != nil
}

// description is TBD
// SetIdentification sets the PatternFlowIpv4Identification value in the FlowIpv4 object
func (obj *flowIpv4) SetIdentification(value PatternFlowIpv4Identification) FlowIpv4 {

	obj.identificationHolder = nil
	obj.obj.Identification = value.msg()

	return obj
}

// description is TBD
// Reserved returns a PatternFlowIpv4Reserved
func (obj *flowIpv4) Reserved() PatternFlowIpv4Reserved {
	if obj.obj.Reserved == nil {
		obj.obj.Reserved = NewPatternFlowIpv4Reserved().msg()
	}
	if obj.reservedHolder == nil {
		obj.reservedHolder = &patternFlowIpv4Reserved{obj: obj.obj.Reserved}
	}
	return obj.reservedHolder
}

// description is TBD
// Reserved returns a PatternFlowIpv4Reserved
func (obj *flowIpv4) HasReserved() bool {
	return obj.obj.Reserved != nil
}

// description is TBD
// SetReserved sets the PatternFlowIpv4Reserved value in the FlowIpv4 object
func (obj *flowIpv4) SetReserved(value PatternFlowIpv4Reserved) FlowIpv4 {

	obj.reservedHolder = nil
	obj.obj.Reserved = value.msg()

	return obj
}

// description is TBD
// DontFragment returns a PatternFlowIpv4DontFragment
func (obj *flowIpv4) DontFragment() PatternFlowIpv4DontFragment {
	if obj.obj.DontFragment == nil {
		obj.obj.DontFragment = NewPatternFlowIpv4DontFragment().msg()
	}
	if obj.dontFragmentHolder == nil {
		obj.dontFragmentHolder = &patternFlowIpv4DontFragment{obj: obj.obj.DontFragment}
	}
	return obj.dontFragmentHolder
}

// description is TBD
// DontFragment returns a PatternFlowIpv4DontFragment
func (obj *flowIpv4) HasDontFragment() bool {
	return obj.obj.DontFragment != nil
}

// description is TBD
// SetDontFragment sets the PatternFlowIpv4DontFragment value in the FlowIpv4 object
func (obj *flowIpv4) SetDontFragment(value PatternFlowIpv4DontFragment) FlowIpv4 {

	obj.dontFragmentHolder = nil
	obj.obj.DontFragment = value.msg()

	return obj
}

// description is TBD
// MoreFragments returns a PatternFlowIpv4MoreFragments
func (obj *flowIpv4) MoreFragments() PatternFlowIpv4MoreFragments {
	if obj.obj.MoreFragments == nil {
		obj.obj.MoreFragments = NewPatternFlowIpv4MoreFragments().msg()
	}
	if obj.moreFragmentsHolder == nil {
		obj.moreFragmentsHolder = &patternFlowIpv4MoreFragments{obj: obj.obj.MoreFragments}
	}
	return obj.moreFragmentsHolder
}

// description is TBD
// MoreFragments returns a PatternFlowIpv4MoreFragments
func (obj *flowIpv4) HasMoreFragments() bool {
	return obj.obj.MoreFragments != nil
}

// description is TBD
// SetMoreFragments sets the PatternFlowIpv4MoreFragments value in the FlowIpv4 object
func (obj *flowIpv4) SetMoreFragments(value PatternFlowIpv4MoreFragments) FlowIpv4 {

	obj.moreFragmentsHolder = nil
	obj.obj.MoreFragments = value.msg()

	return obj
}

// description is TBD
// FragmentOffset returns a PatternFlowIpv4FragmentOffset
func (obj *flowIpv4) FragmentOffset() PatternFlowIpv4FragmentOffset {
	if obj.obj.FragmentOffset == nil {
		obj.obj.FragmentOffset = NewPatternFlowIpv4FragmentOffset().msg()
	}
	if obj.fragmentOffsetHolder == nil {
		obj.fragmentOffsetHolder = &patternFlowIpv4FragmentOffset{obj: obj.obj.FragmentOffset}
	}
	return obj.fragmentOffsetHolder
}

// description is TBD
// FragmentOffset returns a PatternFlowIpv4FragmentOffset
func (obj *flowIpv4) HasFragmentOffset() bool {
	return obj.obj.FragmentOffset != nil
}

// description is TBD
// SetFragmentOffset sets the PatternFlowIpv4FragmentOffset value in the FlowIpv4 object
func (obj *flowIpv4) SetFragmentOffset(value PatternFlowIpv4FragmentOffset) FlowIpv4 {

	obj.fragmentOffsetHolder = nil
	obj.obj.FragmentOffset = value.msg()

	return obj
}

// description is TBD
// TimeToLive returns a PatternFlowIpv4TimeToLive
func (obj *flowIpv4) TimeToLive() PatternFlowIpv4TimeToLive {
	if obj.obj.TimeToLive == nil {
		obj.obj.TimeToLive = NewPatternFlowIpv4TimeToLive().msg()
	}
	if obj.timeToLiveHolder == nil {
		obj.timeToLiveHolder = &patternFlowIpv4TimeToLive{obj: obj.obj.TimeToLive}
	}
	return obj.timeToLiveHolder
}

// description is TBD
// TimeToLive returns a PatternFlowIpv4TimeToLive
func (obj *flowIpv4) HasTimeToLive() bool {
	return obj.obj.TimeToLive != nil
}

// description is TBD
// SetTimeToLive sets the PatternFlowIpv4TimeToLive value in the FlowIpv4 object
func (obj *flowIpv4) SetTimeToLive(value PatternFlowIpv4TimeToLive) FlowIpv4 {

	obj.timeToLiveHolder = nil
	obj.obj.TimeToLive = value.msg()

	return obj
}

// description is TBD
// Protocol returns a PatternFlowIpv4Protocol
func (obj *flowIpv4) Protocol() PatternFlowIpv4Protocol {
	if obj.obj.Protocol == nil {
		obj.obj.Protocol = NewPatternFlowIpv4Protocol().msg()
	}
	if obj.protocolHolder == nil {
		obj.protocolHolder = &patternFlowIpv4Protocol{obj: obj.obj.Protocol}
	}
	return obj.protocolHolder
}

// description is TBD
// Protocol returns a PatternFlowIpv4Protocol
func (obj *flowIpv4) HasProtocol() bool {
	return obj.obj.Protocol != nil
}

// description is TBD
// SetProtocol sets the PatternFlowIpv4Protocol value in the FlowIpv4 object
func (obj *flowIpv4) SetProtocol(value PatternFlowIpv4Protocol) FlowIpv4 {

	obj.protocolHolder = nil
	obj.obj.Protocol = value.msg()

	return obj
}

// description is TBD
// HeaderChecksum returns a PatternFlowIpv4HeaderChecksum
func (obj *flowIpv4) HeaderChecksum() PatternFlowIpv4HeaderChecksum {
	if obj.obj.HeaderChecksum == nil {
		obj.obj.HeaderChecksum = NewPatternFlowIpv4HeaderChecksum().msg()
	}
	if obj.headerChecksumHolder == nil {
		obj.headerChecksumHolder = &patternFlowIpv4HeaderChecksum{obj: obj.obj.HeaderChecksum}
	}
	return obj.headerChecksumHolder
}

// description is TBD
// HeaderChecksum returns a PatternFlowIpv4HeaderChecksum
func (obj *flowIpv4) HasHeaderChecksum() bool {
	return obj.obj.HeaderChecksum != nil
}

// description is TBD
// SetHeaderChecksum sets the PatternFlowIpv4HeaderChecksum value in the FlowIpv4 object
func (obj *flowIpv4) SetHeaderChecksum(value PatternFlowIpv4HeaderChecksum) FlowIpv4 {

	obj.headerChecksumHolder = nil
	obj.obj.HeaderChecksum = value.msg()

	return obj
}

// description is TBD
// Src returns a PatternFlowIpv4Src
func (obj *flowIpv4) Src() PatternFlowIpv4Src {
	if obj.obj.Src == nil {
		obj.obj.Src = NewPatternFlowIpv4Src().msg()
	}
	if obj.srcHolder == nil {
		obj.srcHolder = &patternFlowIpv4Src{obj: obj.obj.Src}
	}
	return obj.srcHolder
}

// description is TBD
// Src returns a PatternFlowIpv4Src
func (obj *flowIpv4) HasSrc() bool {
	return obj.obj.Src != nil
}

// description is TBD
// SetSrc sets the PatternFlowIpv4Src value in the FlowIpv4 object
func (obj *flowIpv4) SetSrc(value PatternFlowIpv4Src) FlowIpv4 {

	obj.srcHolder = nil
	obj.obj.Src = value.msg()

	return obj
}

// description is TBD
// Dst returns a PatternFlowIpv4Dst
func (obj *flowIpv4) Dst() PatternFlowIpv4Dst {
	if obj.obj.Dst == nil {
		obj.obj.Dst = NewPatternFlowIpv4Dst().msg()
	}
	if obj.dstHolder == nil {
		obj.dstHolder = &patternFlowIpv4Dst{obj: obj.obj.Dst}
	}
	return obj.dstHolder
}

// description is TBD
// Dst returns a PatternFlowIpv4Dst
func (obj *flowIpv4) HasDst() bool {
	return obj.obj.Dst != nil
}

// description is TBD
// SetDst sets the PatternFlowIpv4Dst value in the FlowIpv4 object
func (obj *flowIpv4) SetDst(value PatternFlowIpv4Dst) FlowIpv4 {

	obj.dstHolder = nil
	obj.obj.Dst = value.msg()

	return obj
}

// description is TBD
// Options returns a []FlowIpv4Options
func (obj *flowIpv4) Options() FlowIpv4FlowIpv4OptionsIter {
	if len(obj.obj.Options) == 0 {
		obj.obj.Options = []*otg.FlowIpv4Options{}
	}
	if obj.optionsHolder == nil {
		obj.optionsHolder = newFlowIpv4FlowIpv4OptionsIter(&obj.obj.Options).setMsg(obj)
	}
	return obj.optionsHolder
}

type flowIpv4FlowIpv4OptionsIter struct {
	obj                  *flowIpv4
	flowIpv4OptionsSlice []FlowIpv4Options
	fieldPtr             *[]*otg.FlowIpv4Options
}

func newFlowIpv4FlowIpv4OptionsIter(ptr *[]*otg.FlowIpv4Options) FlowIpv4FlowIpv4OptionsIter {
	return &flowIpv4FlowIpv4OptionsIter{fieldPtr: ptr}
}

type FlowIpv4FlowIpv4OptionsIter interface {
	setMsg(*flowIpv4) FlowIpv4FlowIpv4OptionsIter
	Items() []FlowIpv4Options
	Add() FlowIpv4Options
	Append(items ...FlowIpv4Options) FlowIpv4FlowIpv4OptionsIter
	Set(index int, newObj FlowIpv4Options) FlowIpv4FlowIpv4OptionsIter
	Clear() FlowIpv4FlowIpv4OptionsIter
	clearHolderSlice() FlowIpv4FlowIpv4OptionsIter
	appendHolderSlice(item FlowIpv4Options) FlowIpv4FlowIpv4OptionsIter
}

func (obj *flowIpv4FlowIpv4OptionsIter) setMsg(msg *flowIpv4) FlowIpv4FlowIpv4OptionsIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&flowIpv4Options{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *flowIpv4FlowIpv4OptionsIter) Items() []FlowIpv4Options {
	return obj.flowIpv4OptionsSlice
}

func (obj *flowIpv4FlowIpv4OptionsIter) Add() FlowIpv4Options {
	newObj := &otg.FlowIpv4Options{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &flowIpv4Options{obj: newObj}
	newLibObj.setDefault()
	obj.flowIpv4OptionsSlice = append(obj.flowIpv4OptionsSlice, newLibObj)
	return newLibObj
}

func (obj *flowIpv4FlowIpv4OptionsIter) Append(items ...FlowIpv4Options) FlowIpv4FlowIpv4OptionsIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.flowIpv4OptionsSlice = append(obj.flowIpv4OptionsSlice, item)
	}
	return obj
}

func (obj *flowIpv4FlowIpv4OptionsIter) Set(index int, newObj FlowIpv4Options) FlowIpv4FlowIpv4OptionsIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.flowIpv4OptionsSlice[index] = newObj
	return obj
}
func (obj *flowIpv4FlowIpv4OptionsIter) Clear() FlowIpv4FlowIpv4OptionsIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.FlowIpv4Options{}
		obj.flowIpv4OptionsSlice = []FlowIpv4Options{}
	}
	return obj
}
func (obj *flowIpv4FlowIpv4OptionsIter) clearHolderSlice() FlowIpv4FlowIpv4OptionsIter {
	if len(obj.flowIpv4OptionsSlice) > 0 {
		obj.flowIpv4OptionsSlice = []FlowIpv4Options{}
	}
	return obj
}
func (obj *flowIpv4FlowIpv4OptionsIter) appendHolderSlice(item FlowIpv4Options) FlowIpv4FlowIpv4OptionsIter {
	obj.flowIpv4OptionsSlice = append(obj.flowIpv4OptionsSlice, item)
	return obj
}

func (obj *flowIpv4) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Version != nil {

		obj.Version().validateObj(vObj, set_default)
	}

	if obj.obj.HeaderLength != nil {

		obj.HeaderLength().validateObj(vObj, set_default)
	}

	if obj.obj.Priority != nil {

		obj.Priority().validateObj(vObj, set_default)
	}

	if obj.obj.TotalLength != nil {

		obj.TotalLength().validateObj(vObj, set_default)
	}

	if obj.obj.Identification != nil {

		obj.Identification().validateObj(vObj, set_default)
	}

	if obj.obj.Reserved != nil {

		obj.Reserved().validateObj(vObj, set_default)
	}

	if obj.obj.DontFragment != nil {

		obj.DontFragment().validateObj(vObj, set_default)
	}

	if obj.obj.MoreFragments != nil {

		obj.MoreFragments().validateObj(vObj, set_default)
	}

	if obj.obj.FragmentOffset != nil {

		obj.FragmentOffset().validateObj(vObj, set_default)
	}

	if obj.obj.TimeToLive != nil {

		obj.TimeToLive().validateObj(vObj, set_default)
	}

	if obj.obj.Protocol != nil {

		obj.Protocol().validateObj(vObj, set_default)
	}

	if obj.obj.HeaderChecksum != nil {

		obj.HeaderChecksum().validateObj(vObj, set_default)
	}

	if obj.obj.Src != nil {

		obj.Src().validateObj(vObj, set_default)
	}

	if obj.obj.Dst != nil {

		obj.Dst().validateObj(vObj, set_default)
	}

	if len(obj.obj.Options) != 0 {

		if set_default {
			obj.Options().clearHolderSlice()
			for _, item := range obj.obj.Options {
				obj.Options().appendHolderSlice(&flowIpv4Options{obj: item})
			}
		}
		for _, item := range obj.Options().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *flowIpv4) setDefault() {

}
