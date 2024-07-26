package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpAttributesBsidMpls *****
type bgpAttributesBsidMpls struct {
	validation
	obj           *otg.BgpAttributesBsidMpls
	marshaller    marshalBgpAttributesBsidMpls
	unMarshaller  unMarshalBgpAttributesBsidMpls
	mplsSidHolder BgpAttributesSidMpls
}

func NewBgpAttributesBsidMpls() BgpAttributesBsidMpls {
	obj := bgpAttributesBsidMpls{obj: &otg.BgpAttributesBsidMpls{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpAttributesBsidMpls) msg() *otg.BgpAttributesBsidMpls {
	return obj.obj
}

func (obj *bgpAttributesBsidMpls) setMsg(msg *otg.BgpAttributesBsidMpls) BgpAttributesBsidMpls {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpAttributesBsidMpls struct {
	obj *bgpAttributesBsidMpls
}

type marshalBgpAttributesBsidMpls interface {
	// ToProto marshals BgpAttributesBsidMpls to protobuf object *otg.BgpAttributesBsidMpls
	ToProto() (*otg.BgpAttributesBsidMpls, error)
	// ToPbText marshals BgpAttributesBsidMpls to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpAttributesBsidMpls to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpAttributesBsidMpls to JSON text
	ToJson() (string, error)
}

type unMarshalbgpAttributesBsidMpls struct {
	obj *bgpAttributesBsidMpls
}

type unMarshalBgpAttributesBsidMpls interface {
	// FromProto unmarshals BgpAttributesBsidMpls from protobuf object *otg.BgpAttributesBsidMpls
	FromProto(msg *otg.BgpAttributesBsidMpls) (BgpAttributesBsidMpls, error)
	// FromPbText unmarshals BgpAttributesBsidMpls from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpAttributesBsidMpls from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpAttributesBsidMpls from JSON text
	FromJson(value string) error
}

func (obj *bgpAttributesBsidMpls) Marshal() marshalBgpAttributesBsidMpls {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpAttributesBsidMpls{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpAttributesBsidMpls) Unmarshal() unMarshalBgpAttributesBsidMpls {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpAttributesBsidMpls{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpAttributesBsidMpls) ToProto() (*otg.BgpAttributesBsidMpls, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpAttributesBsidMpls) FromProto(msg *otg.BgpAttributesBsidMpls) (BgpAttributesBsidMpls, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpAttributesBsidMpls) ToPbText() (string, error) {
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

func (m *unMarshalbgpAttributesBsidMpls) FromPbText(value string) error {
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

func (m *marshalbgpAttributesBsidMpls) ToYaml() (string, error) {
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

func (m *unMarshalbgpAttributesBsidMpls) FromYaml(value string) error {
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

func (m *marshalbgpAttributesBsidMpls) ToJson() (string, error) {
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

func (m *unMarshalbgpAttributesBsidMpls) FromJson(value string) error {
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

func (obj *bgpAttributesBsidMpls) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpAttributesBsidMpls) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpAttributesBsidMpls) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpAttributesBsidMpls) Clone() (BgpAttributesBsidMpls, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpAttributesBsidMpls()
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

func (obj *bgpAttributesBsidMpls) setNil() {
	obj.mplsSidHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpAttributesBsidMpls is when the active candidate path has a specified Binding Segment Identifier, the SR Policy uses that BSID defined
// as a MPLS label.The format of the sub-TLV is defined in draft-ietf-idr-sr-policy-safi-02  Section 2.4.2 .
type BgpAttributesBsidMpls interface {
	Validation
	// msg marshals BgpAttributesBsidMpls to protobuf object *otg.BgpAttributesBsidMpls
	// and doesn't set defaults
	msg() *otg.BgpAttributesBsidMpls
	// setMsg unmarshals BgpAttributesBsidMpls from protobuf object *otg.BgpAttributesBsidMpls
	// and doesn't set defaults
	setMsg(*otg.BgpAttributesBsidMpls) BgpAttributesBsidMpls
	// provides marshal interface
	Marshal() marshalBgpAttributesBsidMpls
	// provides unmarshal interface
	Unmarshal() unMarshalBgpAttributesBsidMpls
	// validate validates BgpAttributesBsidMpls
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpAttributesBsidMpls, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// FlagSpecifiedBsidOnly returns bool, set in BgpAttributesBsidMpls.
	FlagSpecifiedBsidOnly() bool
	// SetFlagSpecifiedBsidOnly assigns bool provided by user to BgpAttributesBsidMpls
	SetFlagSpecifiedBsidOnly(value bool) BgpAttributesBsidMpls
	// HasFlagSpecifiedBsidOnly checks if FlagSpecifiedBsidOnly has been set in BgpAttributesBsidMpls
	HasFlagSpecifiedBsidOnly() bool
	// FlagDropUponInvalid returns bool, set in BgpAttributesBsidMpls.
	FlagDropUponInvalid() bool
	// SetFlagDropUponInvalid assigns bool provided by user to BgpAttributesBsidMpls
	SetFlagDropUponInvalid(value bool) BgpAttributesBsidMpls
	// HasFlagDropUponInvalid checks if FlagDropUponInvalid has been set in BgpAttributesBsidMpls
	HasFlagDropUponInvalid() bool
	// MplsSid returns BgpAttributesSidMpls, set in BgpAttributesBsidMpls.
	// BgpAttributesSidMpls is this carries a 20 bit Multi Protocol Label Switching alongwith 3 bits traffic class, 1 bit indicating presence
	// or absence of Bottom-Of-Stack and 8 bits carrying the Time to Live value.
	MplsSid() BgpAttributesSidMpls
	// SetMplsSid assigns BgpAttributesSidMpls provided by user to BgpAttributesBsidMpls.
	// BgpAttributesSidMpls is this carries a 20 bit Multi Protocol Label Switching alongwith 3 bits traffic class, 1 bit indicating presence
	// or absence of Bottom-Of-Stack and 8 bits carrying the Time to Live value.
	SetMplsSid(value BgpAttributesSidMpls) BgpAttributesBsidMpls
	// HasMplsSid checks if MplsSid has been set in BgpAttributesBsidMpls
	HasMplsSid() bool
	setNil()
}

// S-Flag: This flag encodes the "Specified-BSID-only" behavior. It's usage is
// described in section 6.2.3 in [RFC9256].
// FlagSpecifiedBsidOnly returns a bool
func (obj *bgpAttributesBsidMpls) FlagSpecifiedBsidOnly() bool {

	return *obj.obj.FlagSpecifiedBsidOnly

}

// S-Flag: This flag encodes the "Specified-BSID-only" behavior. It's usage is
// described in section 6.2.3 in [RFC9256].
// FlagSpecifiedBsidOnly returns a bool
func (obj *bgpAttributesBsidMpls) HasFlagSpecifiedBsidOnly() bool {
	return obj.obj.FlagSpecifiedBsidOnly != nil
}

// S-Flag: This flag encodes the "Specified-BSID-only" behavior. It's usage is
// described in section 6.2.3 in [RFC9256].
// SetFlagSpecifiedBsidOnly sets the bool value in the BgpAttributesBsidMpls object
func (obj *bgpAttributesBsidMpls) SetFlagSpecifiedBsidOnly(value bool) BgpAttributesBsidMpls {

	obj.obj.FlagSpecifiedBsidOnly = &value
	return obj
}

// I-Flag: This flag encodes the "Drop Upon Invalid" behavior.
// It's usage is described in section 8.2 in [RFC9256].
// FlagDropUponInvalid returns a bool
func (obj *bgpAttributesBsidMpls) FlagDropUponInvalid() bool {

	return *obj.obj.FlagDropUponInvalid

}

// I-Flag: This flag encodes the "Drop Upon Invalid" behavior.
// It's usage is described in section 8.2 in [RFC9256].
// FlagDropUponInvalid returns a bool
func (obj *bgpAttributesBsidMpls) HasFlagDropUponInvalid() bool {
	return obj.obj.FlagDropUponInvalid != nil
}

// I-Flag: This flag encodes the "Drop Upon Invalid" behavior.
// It's usage is described in section 8.2 in [RFC9256].
// SetFlagDropUponInvalid sets the bool value in the BgpAttributesBsidMpls object
func (obj *bgpAttributesBsidMpls) SetFlagDropUponInvalid(value bool) BgpAttributesBsidMpls {

	obj.obj.FlagDropUponInvalid = &value
	return obj
}

// description is TBD
// MplsSid returns a BgpAttributesSidMpls
func (obj *bgpAttributesBsidMpls) MplsSid() BgpAttributesSidMpls {
	if obj.obj.MplsSid == nil {
		obj.obj.MplsSid = NewBgpAttributesSidMpls().msg()
	}
	if obj.mplsSidHolder == nil {
		obj.mplsSidHolder = &bgpAttributesSidMpls{obj: obj.obj.MplsSid}
	}
	return obj.mplsSidHolder
}

// description is TBD
// MplsSid returns a BgpAttributesSidMpls
func (obj *bgpAttributesBsidMpls) HasMplsSid() bool {
	return obj.obj.MplsSid != nil
}

// description is TBD
// SetMplsSid sets the BgpAttributesSidMpls value in the BgpAttributesBsidMpls object
func (obj *bgpAttributesBsidMpls) SetMplsSid(value BgpAttributesSidMpls) BgpAttributesBsidMpls {

	obj.mplsSidHolder = nil
	obj.obj.MplsSid = value.msg()

	return obj
}

func (obj *bgpAttributesBsidMpls) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.MplsSid != nil {

		obj.MplsSid().validateObj(vObj, set_default)
	}

}

func (obj *bgpAttributesBsidMpls) setDefault() {
	if obj.obj.FlagSpecifiedBsidOnly == nil {
		obj.SetFlagSpecifiedBsidOnly(false)
	}
	if obj.obj.FlagDropUponInvalid == nil {
		obj.SetFlagDropUponInvalid(false)
	}

}
