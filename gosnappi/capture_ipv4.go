package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** CaptureIpv4 *****
type captureIpv4 struct {
	validation
	obj                  *otg.CaptureIpv4
	marshaller           marshalCaptureIpv4
	unMarshaller         unMarshalCaptureIpv4
	versionHolder        CaptureField
	headerLengthHolder   CaptureField
	priorityHolder       CaptureField
	totalLengthHolder    CaptureField
	identificationHolder CaptureField
	reservedHolder       CaptureField
	dontFragmentHolder   CaptureField
	moreFragmentsHolder  CaptureField
	fragmentOffsetHolder CaptureField
	timeToLiveHolder     CaptureField
	protocolHolder       CaptureField
	headerChecksumHolder CaptureField
	srcHolder            CaptureField
	dstHolder            CaptureField
}

func NewCaptureIpv4() CaptureIpv4 {
	obj := captureIpv4{obj: &otg.CaptureIpv4{}}
	obj.setDefault()
	return &obj
}

func (obj *captureIpv4) msg() *otg.CaptureIpv4 {
	return obj.obj
}

func (obj *captureIpv4) setMsg(msg *otg.CaptureIpv4) CaptureIpv4 {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalcaptureIpv4 struct {
	obj *captureIpv4
}

type marshalCaptureIpv4 interface {
	// ToProto marshals CaptureIpv4 to protobuf object *otg.CaptureIpv4
	ToProto() (*otg.CaptureIpv4, error)
	// ToPbText marshals CaptureIpv4 to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals CaptureIpv4 to YAML text
	ToYaml() (string, error)
	// ToJson marshals CaptureIpv4 to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals CaptureIpv4 to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalcaptureIpv4 struct {
	obj *captureIpv4
}

type unMarshalCaptureIpv4 interface {
	// FromProto unmarshals CaptureIpv4 from protobuf object *otg.CaptureIpv4
	FromProto(msg *otg.CaptureIpv4) (CaptureIpv4, error)
	// FromPbText unmarshals CaptureIpv4 from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals CaptureIpv4 from YAML text
	FromYaml(value string) error
	// FromJson unmarshals CaptureIpv4 from JSON text
	FromJson(value string) error
}

func (obj *captureIpv4) Marshal() marshalCaptureIpv4 {
	if obj.marshaller == nil {
		obj.marshaller = &marshalcaptureIpv4{obj: obj}
	}
	return obj.marshaller
}

func (obj *captureIpv4) Unmarshal() unMarshalCaptureIpv4 {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalcaptureIpv4{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalcaptureIpv4) ToProto() (*otg.CaptureIpv4, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalcaptureIpv4) FromProto(msg *otg.CaptureIpv4) (CaptureIpv4, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalcaptureIpv4) ToPbText() (string, error) {
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

func (m *unMarshalcaptureIpv4) FromPbText(value string) error {
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

func (m *marshalcaptureIpv4) ToYaml() (string, error) {
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

func (m *unMarshalcaptureIpv4) FromYaml(value string) error {
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

func (m *marshalcaptureIpv4) ToJsonRaw() (string, error) {
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

func (m *marshalcaptureIpv4) ToJson() (string, error) {
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

func (m *unMarshalcaptureIpv4) FromJson(value string) error {
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

func (obj *captureIpv4) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *captureIpv4) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *captureIpv4) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *captureIpv4) Clone() (CaptureIpv4, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewCaptureIpv4()
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

func (obj *captureIpv4) setNil() {
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
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// CaptureIpv4 is description is TBD
type CaptureIpv4 interface {
	Validation
	// msg marshals CaptureIpv4 to protobuf object *otg.CaptureIpv4
	// and doesn't set defaults
	msg() *otg.CaptureIpv4
	// setMsg unmarshals CaptureIpv4 from protobuf object *otg.CaptureIpv4
	// and doesn't set defaults
	setMsg(*otg.CaptureIpv4) CaptureIpv4
	// provides marshal interface
	Marshal() marshalCaptureIpv4
	// provides unmarshal interface
	Unmarshal() unMarshalCaptureIpv4
	// validate validates CaptureIpv4
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (CaptureIpv4, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Version returns CaptureField, set in CaptureIpv4.
	// CaptureField is description is TBD
	Version() CaptureField
	// SetVersion assigns CaptureField provided by user to CaptureIpv4.
	// CaptureField is description is TBD
	SetVersion(value CaptureField) CaptureIpv4
	// HasVersion checks if Version has been set in CaptureIpv4
	HasVersion() bool
	// HeaderLength returns CaptureField, set in CaptureIpv4.
	// CaptureField is description is TBD
	HeaderLength() CaptureField
	// SetHeaderLength assigns CaptureField provided by user to CaptureIpv4.
	// CaptureField is description is TBD
	SetHeaderLength(value CaptureField) CaptureIpv4
	// HasHeaderLength checks if HeaderLength has been set in CaptureIpv4
	HasHeaderLength() bool
	// Priority returns CaptureField, set in CaptureIpv4.
	// CaptureField is description is TBD
	Priority() CaptureField
	// SetPriority assigns CaptureField provided by user to CaptureIpv4.
	// CaptureField is description is TBD
	SetPriority(value CaptureField) CaptureIpv4
	// HasPriority checks if Priority has been set in CaptureIpv4
	HasPriority() bool
	// TotalLength returns CaptureField, set in CaptureIpv4.
	// CaptureField is description is TBD
	TotalLength() CaptureField
	// SetTotalLength assigns CaptureField provided by user to CaptureIpv4.
	// CaptureField is description is TBD
	SetTotalLength(value CaptureField) CaptureIpv4
	// HasTotalLength checks if TotalLength has been set in CaptureIpv4
	HasTotalLength() bool
	// Identification returns CaptureField, set in CaptureIpv4.
	// CaptureField is description is TBD
	Identification() CaptureField
	// SetIdentification assigns CaptureField provided by user to CaptureIpv4.
	// CaptureField is description is TBD
	SetIdentification(value CaptureField) CaptureIpv4
	// HasIdentification checks if Identification has been set in CaptureIpv4
	HasIdentification() bool
	// Reserved returns CaptureField, set in CaptureIpv4.
	// CaptureField is description is TBD
	Reserved() CaptureField
	// SetReserved assigns CaptureField provided by user to CaptureIpv4.
	// CaptureField is description is TBD
	SetReserved(value CaptureField) CaptureIpv4
	// HasReserved checks if Reserved has been set in CaptureIpv4
	HasReserved() bool
	// DontFragment returns CaptureField, set in CaptureIpv4.
	// CaptureField is description is TBD
	DontFragment() CaptureField
	// SetDontFragment assigns CaptureField provided by user to CaptureIpv4.
	// CaptureField is description is TBD
	SetDontFragment(value CaptureField) CaptureIpv4
	// HasDontFragment checks if DontFragment has been set in CaptureIpv4
	HasDontFragment() bool
	// MoreFragments returns CaptureField, set in CaptureIpv4.
	// CaptureField is description is TBD
	MoreFragments() CaptureField
	// SetMoreFragments assigns CaptureField provided by user to CaptureIpv4.
	// CaptureField is description is TBD
	SetMoreFragments(value CaptureField) CaptureIpv4
	// HasMoreFragments checks if MoreFragments has been set in CaptureIpv4
	HasMoreFragments() bool
	// FragmentOffset returns CaptureField, set in CaptureIpv4.
	// CaptureField is description is TBD
	FragmentOffset() CaptureField
	// SetFragmentOffset assigns CaptureField provided by user to CaptureIpv4.
	// CaptureField is description is TBD
	SetFragmentOffset(value CaptureField) CaptureIpv4
	// HasFragmentOffset checks if FragmentOffset has been set in CaptureIpv4
	HasFragmentOffset() bool
	// TimeToLive returns CaptureField, set in CaptureIpv4.
	// CaptureField is description is TBD
	TimeToLive() CaptureField
	// SetTimeToLive assigns CaptureField provided by user to CaptureIpv4.
	// CaptureField is description is TBD
	SetTimeToLive(value CaptureField) CaptureIpv4
	// HasTimeToLive checks if TimeToLive has been set in CaptureIpv4
	HasTimeToLive() bool
	// Protocol returns CaptureField, set in CaptureIpv4.
	// CaptureField is description is TBD
	Protocol() CaptureField
	// SetProtocol assigns CaptureField provided by user to CaptureIpv4.
	// CaptureField is description is TBD
	SetProtocol(value CaptureField) CaptureIpv4
	// HasProtocol checks if Protocol has been set in CaptureIpv4
	HasProtocol() bool
	// HeaderChecksum returns CaptureField, set in CaptureIpv4.
	// CaptureField is description is TBD
	HeaderChecksum() CaptureField
	// SetHeaderChecksum assigns CaptureField provided by user to CaptureIpv4.
	// CaptureField is description is TBD
	SetHeaderChecksum(value CaptureField) CaptureIpv4
	// HasHeaderChecksum checks if HeaderChecksum has been set in CaptureIpv4
	HasHeaderChecksum() bool
	// Src returns CaptureField, set in CaptureIpv4.
	// CaptureField is description is TBD
	Src() CaptureField
	// SetSrc assigns CaptureField provided by user to CaptureIpv4.
	// CaptureField is description is TBD
	SetSrc(value CaptureField) CaptureIpv4
	// HasSrc checks if Src has been set in CaptureIpv4
	HasSrc() bool
	// Dst returns CaptureField, set in CaptureIpv4.
	// CaptureField is description is TBD
	Dst() CaptureField
	// SetDst assigns CaptureField provided by user to CaptureIpv4.
	// CaptureField is description is TBD
	SetDst(value CaptureField) CaptureIpv4
	// HasDst checks if Dst has been set in CaptureIpv4
	HasDst() bool
	setNil()
}

// description is TBD
// Version returns a CaptureField
func (obj *captureIpv4) Version() CaptureField {
	if obj.obj.Version == nil {
		obj.obj.Version = NewCaptureField().msg()
	}
	if obj.versionHolder == nil {
		obj.versionHolder = &captureField{obj: obj.obj.Version}
	}
	return obj.versionHolder
}

// description is TBD
// Version returns a CaptureField
func (obj *captureIpv4) HasVersion() bool {
	return obj.obj.Version != nil
}

// description is TBD
// SetVersion sets the CaptureField value in the CaptureIpv4 object
func (obj *captureIpv4) SetVersion(value CaptureField) CaptureIpv4 {

	obj.versionHolder = nil
	obj.obj.Version = value.msg()

	return obj
}

// description is TBD
// HeaderLength returns a CaptureField
func (obj *captureIpv4) HeaderLength() CaptureField {
	if obj.obj.HeaderLength == nil {
		obj.obj.HeaderLength = NewCaptureField().msg()
	}
	if obj.headerLengthHolder == nil {
		obj.headerLengthHolder = &captureField{obj: obj.obj.HeaderLength}
	}
	return obj.headerLengthHolder
}

// description is TBD
// HeaderLength returns a CaptureField
func (obj *captureIpv4) HasHeaderLength() bool {
	return obj.obj.HeaderLength != nil
}

// description is TBD
// SetHeaderLength sets the CaptureField value in the CaptureIpv4 object
func (obj *captureIpv4) SetHeaderLength(value CaptureField) CaptureIpv4 {

	obj.headerLengthHolder = nil
	obj.obj.HeaderLength = value.msg()

	return obj
}

// description is TBD
// Priority returns a CaptureField
func (obj *captureIpv4) Priority() CaptureField {
	if obj.obj.Priority == nil {
		obj.obj.Priority = NewCaptureField().msg()
	}
	if obj.priorityHolder == nil {
		obj.priorityHolder = &captureField{obj: obj.obj.Priority}
	}
	return obj.priorityHolder
}

// description is TBD
// Priority returns a CaptureField
func (obj *captureIpv4) HasPriority() bool {
	return obj.obj.Priority != nil
}

// description is TBD
// SetPriority sets the CaptureField value in the CaptureIpv4 object
func (obj *captureIpv4) SetPriority(value CaptureField) CaptureIpv4 {

	obj.priorityHolder = nil
	obj.obj.Priority = value.msg()

	return obj
}

// description is TBD
// TotalLength returns a CaptureField
func (obj *captureIpv4) TotalLength() CaptureField {
	if obj.obj.TotalLength == nil {
		obj.obj.TotalLength = NewCaptureField().msg()
	}
	if obj.totalLengthHolder == nil {
		obj.totalLengthHolder = &captureField{obj: obj.obj.TotalLength}
	}
	return obj.totalLengthHolder
}

// description is TBD
// TotalLength returns a CaptureField
func (obj *captureIpv4) HasTotalLength() bool {
	return obj.obj.TotalLength != nil
}

// description is TBD
// SetTotalLength sets the CaptureField value in the CaptureIpv4 object
func (obj *captureIpv4) SetTotalLength(value CaptureField) CaptureIpv4 {

	obj.totalLengthHolder = nil
	obj.obj.TotalLength = value.msg()

	return obj
}

// description is TBD
// Identification returns a CaptureField
func (obj *captureIpv4) Identification() CaptureField {
	if obj.obj.Identification == nil {
		obj.obj.Identification = NewCaptureField().msg()
	}
	if obj.identificationHolder == nil {
		obj.identificationHolder = &captureField{obj: obj.obj.Identification}
	}
	return obj.identificationHolder
}

// description is TBD
// Identification returns a CaptureField
func (obj *captureIpv4) HasIdentification() bool {
	return obj.obj.Identification != nil
}

// description is TBD
// SetIdentification sets the CaptureField value in the CaptureIpv4 object
func (obj *captureIpv4) SetIdentification(value CaptureField) CaptureIpv4 {

	obj.identificationHolder = nil
	obj.obj.Identification = value.msg()

	return obj
}

// description is TBD
// Reserved returns a CaptureField
func (obj *captureIpv4) Reserved() CaptureField {
	if obj.obj.Reserved == nil {
		obj.obj.Reserved = NewCaptureField().msg()
	}
	if obj.reservedHolder == nil {
		obj.reservedHolder = &captureField{obj: obj.obj.Reserved}
	}
	return obj.reservedHolder
}

// description is TBD
// Reserved returns a CaptureField
func (obj *captureIpv4) HasReserved() bool {
	return obj.obj.Reserved != nil
}

// description is TBD
// SetReserved sets the CaptureField value in the CaptureIpv4 object
func (obj *captureIpv4) SetReserved(value CaptureField) CaptureIpv4 {

	obj.reservedHolder = nil
	obj.obj.Reserved = value.msg()

	return obj
}

// description is TBD
// DontFragment returns a CaptureField
func (obj *captureIpv4) DontFragment() CaptureField {
	if obj.obj.DontFragment == nil {
		obj.obj.DontFragment = NewCaptureField().msg()
	}
	if obj.dontFragmentHolder == nil {
		obj.dontFragmentHolder = &captureField{obj: obj.obj.DontFragment}
	}
	return obj.dontFragmentHolder
}

// description is TBD
// DontFragment returns a CaptureField
func (obj *captureIpv4) HasDontFragment() bool {
	return obj.obj.DontFragment != nil
}

// description is TBD
// SetDontFragment sets the CaptureField value in the CaptureIpv4 object
func (obj *captureIpv4) SetDontFragment(value CaptureField) CaptureIpv4 {

	obj.dontFragmentHolder = nil
	obj.obj.DontFragment = value.msg()

	return obj
}

// description is TBD
// MoreFragments returns a CaptureField
func (obj *captureIpv4) MoreFragments() CaptureField {
	if obj.obj.MoreFragments == nil {
		obj.obj.MoreFragments = NewCaptureField().msg()
	}
	if obj.moreFragmentsHolder == nil {
		obj.moreFragmentsHolder = &captureField{obj: obj.obj.MoreFragments}
	}
	return obj.moreFragmentsHolder
}

// description is TBD
// MoreFragments returns a CaptureField
func (obj *captureIpv4) HasMoreFragments() bool {
	return obj.obj.MoreFragments != nil
}

// description is TBD
// SetMoreFragments sets the CaptureField value in the CaptureIpv4 object
func (obj *captureIpv4) SetMoreFragments(value CaptureField) CaptureIpv4 {

	obj.moreFragmentsHolder = nil
	obj.obj.MoreFragments = value.msg()

	return obj
}

// description is TBD
// FragmentOffset returns a CaptureField
func (obj *captureIpv4) FragmentOffset() CaptureField {
	if obj.obj.FragmentOffset == nil {
		obj.obj.FragmentOffset = NewCaptureField().msg()
	}
	if obj.fragmentOffsetHolder == nil {
		obj.fragmentOffsetHolder = &captureField{obj: obj.obj.FragmentOffset}
	}
	return obj.fragmentOffsetHolder
}

// description is TBD
// FragmentOffset returns a CaptureField
func (obj *captureIpv4) HasFragmentOffset() bool {
	return obj.obj.FragmentOffset != nil
}

// description is TBD
// SetFragmentOffset sets the CaptureField value in the CaptureIpv4 object
func (obj *captureIpv4) SetFragmentOffset(value CaptureField) CaptureIpv4 {

	obj.fragmentOffsetHolder = nil
	obj.obj.FragmentOffset = value.msg()

	return obj
}

// description is TBD
// TimeToLive returns a CaptureField
func (obj *captureIpv4) TimeToLive() CaptureField {
	if obj.obj.TimeToLive == nil {
		obj.obj.TimeToLive = NewCaptureField().msg()
	}
	if obj.timeToLiveHolder == nil {
		obj.timeToLiveHolder = &captureField{obj: obj.obj.TimeToLive}
	}
	return obj.timeToLiveHolder
}

// description is TBD
// TimeToLive returns a CaptureField
func (obj *captureIpv4) HasTimeToLive() bool {
	return obj.obj.TimeToLive != nil
}

// description is TBD
// SetTimeToLive sets the CaptureField value in the CaptureIpv4 object
func (obj *captureIpv4) SetTimeToLive(value CaptureField) CaptureIpv4 {

	obj.timeToLiveHolder = nil
	obj.obj.TimeToLive = value.msg()

	return obj
}

// description is TBD
// Protocol returns a CaptureField
func (obj *captureIpv4) Protocol() CaptureField {
	if obj.obj.Protocol == nil {
		obj.obj.Protocol = NewCaptureField().msg()
	}
	if obj.protocolHolder == nil {
		obj.protocolHolder = &captureField{obj: obj.obj.Protocol}
	}
	return obj.protocolHolder
}

// description is TBD
// Protocol returns a CaptureField
func (obj *captureIpv4) HasProtocol() bool {
	return obj.obj.Protocol != nil
}

// description is TBD
// SetProtocol sets the CaptureField value in the CaptureIpv4 object
func (obj *captureIpv4) SetProtocol(value CaptureField) CaptureIpv4 {

	obj.protocolHolder = nil
	obj.obj.Protocol = value.msg()

	return obj
}

// description is TBD
// HeaderChecksum returns a CaptureField
func (obj *captureIpv4) HeaderChecksum() CaptureField {
	if obj.obj.HeaderChecksum == nil {
		obj.obj.HeaderChecksum = NewCaptureField().msg()
	}
	if obj.headerChecksumHolder == nil {
		obj.headerChecksumHolder = &captureField{obj: obj.obj.HeaderChecksum}
	}
	return obj.headerChecksumHolder
}

// description is TBD
// HeaderChecksum returns a CaptureField
func (obj *captureIpv4) HasHeaderChecksum() bool {
	return obj.obj.HeaderChecksum != nil
}

// description is TBD
// SetHeaderChecksum sets the CaptureField value in the CaptureIpv4 object
func (obj *captureIpv4) SetHeaderChecksum(value CaptureField) CaptureIpv4 {

	obj.headerChecksumHolder = nil
	obj.obj.HeaderChecksum = value.msg()

	return obj
}

// description is TBD
// Src returns a CaptureField
func (obj *captureIpv4) Src() CaptureField {
	if obj.obj.Src == nil {
		obj.obj.Src = NewCaptureField().msg()
	}
	if obj.srcHolder == nil {
		obj.srcHolder = &captureField{obj: obj.obj.Src}
	}
	return obj.srcHolder
}

// description is TBD
// Src returns a CaptureField
func (obj *captureIpv4) HasSrc() bool {
	return obj.obj.Src != nil
}

// description is TBD
// SetSrc sets the CaptureField value in the CaptureIpv4 object
func (obj *captureIpv4) SetSrc(value CaptureField) CaptureIpv4 {

	obj.srcHolder = nil
	obj.obj.Src = value.msg()

	return obj
}

// description is TBD
// Dst returns a CaptureField
func (obj *captureIpv4) Dst() CaptureField {
	if obj.obj.Dst == nil {
		obj.obj.Dst = NewCaptureField().msg()
	}
	if obj.dstHolder == nil {
		obj.dstHolder = &captureField{obj: obj.obj.Dst}
	}
	return obj.dstHolder
}

// description is TBD
// Dst returns a CaptureField
func (obj *captureIpv4) HasDst() bool {
	return obj.obj.Dst != nil
}

// description is TBD
// SetDst sets the CaptureField value in the CaptureIpv4 object
func (obj *captureIpv4) SetDst(value CaptureField) CaptureIpv4 {

	obj.dstHolder = nil
	obj.obj.Dst = value.msg()

	return obj
}

func (obj *captureIpv4) validateObj(vObj *validation, set_default bool) {
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

}

func (obj *captureIpv4) setDefault() {

}
