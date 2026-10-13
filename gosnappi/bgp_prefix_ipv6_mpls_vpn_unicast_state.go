package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpPrefixIpv6MplsVpnUnicastState *****
type bgpPrefixIpv6MplsVpnUnicastState struct {
	validation
	obj                    *otg.BgpPrefixIpv6MplsVpnUnicastState
	marshaller             marshalBgpPrefixIpv6MplsVpnUnicastState
	unMarshaller           unMarshalBgpPrefixIpv6MplsVpnUnicastState
	matchedVrfRoutesHolder BgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes
}

func NewBgpPrefixIpv6MplsVpnUnicastState() BgpPrefixIpv6MplsVpnUnicastState {
	obj := bgpPrefixIpv6MplsVpnUnicastState{obj: &otg.BgpPrefixIpv6MplsVpnUnicastState{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpPrefixIpv6MplsVpnUnicastState) msg() *otg.BgpPrefixIpv6MplsVpnUnicastState {
	return obj.obj
}

func (obj *bgpPrefixIpv6MplsVpnUnicastState) setMsg(msg *otg.BgpPrefixIpv6MplsVpnUnicastState) BgpPrefixIpv6MplsVpnUnicastState {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpPrefixIpv6MplsVpnUnicastState struct {
	obj *bgpPrefixIpv6MplsVpnUnicastState
}

type marshalBgpPrefixIpv6MplsVpnUnicastState interface {
	// ToProto marshals BgpPrefixIpv6MplsVpnUnicastState to protobuf object *otg.BgpPrefixIpv6MplsVpnUnicastState
	ToProto() (*otg.BgpPrefixIpv6MplsVpnUnicastState, error)
	// ToPbText marshals BgpPrefixIpv6MplsVpnUnicastState to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpPrefixIpv6MplsVpnUnicastState to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpPrefixIpv6MplsVpnUnicastState to JSON text
	ToJson() (string, error)
}

type unMarshalbgpPrefixIpv6MplsVpnUnicastState struct {
	obj *bgpPrefixIpv6MplsVpnUnicastState
}

type unMarshalBgpPrefixIpv6MplsVpnUnicastState interface {
	// FromProto unmarshals BgpPrefixIpv6MplsVpnUnicastState from protobuf object *otg.BgpPrefixIpv6MplsVpnUnicastState
	FromProto(msg *otg.BgpPrefixIpv6MplsVpnUnicastState) (BgpPrefixIpv6MplsVpnUnicastState, error)
	// FromPbText unmarshals BgpPrefixIpv6MplsVpnUnicastState from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpPrefixIpv6MplsVpnUnicastState from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpPrefixIpv6MplsVpnUnicastState from JSON text
	FromJson(value string) error
}

func (obj *bgpPrefixIpv6MplsVpnUnicastState) Marshal() marshalBgpPrefixIpv6MplsVpnUnicastState {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpPrefixIpv6MplsVpnUnicastState{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpPrefixIpv6MplsVpnUnicastState) Unmarshal() unMarshalBgpPrefixIpv6MplsVpnUnicastState {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpPrefixIpv6MplsVpnUnicastState{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpPrefixIpv6MplsVpnUnicastState) ToProto() (*otg.BgpPrefixIpv6MplsVpnUnicastState, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpPrefixIpv6MplsVpnUnicastState) FromProto(msg *otg.BgpPrefixIpv6MplsVpnUnicastState) (BgpPrefixIpv6MplsVpnUnicastState, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpPrefixIpv6MplsVpnUnicastState) ToPbText() (string, error) {
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

func (m *unMarshalbgpPrefixIpv6MplsVpnUnicastState) FromPbText(value string) error {
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

func (m *marshalbgpPrefixIpv6MplsVpnUnicastState) ToYaml() (string, error) {
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

func (m *unMarshalbgpPrefixIpv6MplsVpnUnicastState) FromYaml(value string) error {
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

func (m *marshalbgpPrefixIpv6MplsVpnUnicastState) ToJson() (string, error) {
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

func (m *unMarshalbgpPrefixIpv6MplsVpnUnicastState) FromJson(value string) error {
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

func (obj *bgpPrefixIpv6MplsVpnUnicastState) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpPrefixIpv6MplsVpnUnicastState) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpPrefixIpv6MplsVpnUnicastState) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpPrefixIpv6MplsVpnUnicastState) Clone() (BgpPrefixIpv6MplsVpnUnicastState, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpPrefixIpv6MplsVpnUnicastState()
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

func (obj *bgpPrefixIpv6MplsVpnUnicastState) setNil() {
	obj.matchedVrfRoutesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpPrefixIpv6MplsVpnUnicastState is bGP/MPLS L3VPN (RFC 4659) VPN-IPv6 (6VPE) learned prefix, received under the VPN-IPv6 AFI/SAFI (AFI 2, SAFI 128).
type BgpPrefixIpv6MplsVpnUnicastState interface {
	Validation
	// msg marshals BgpPrefixIpv6MplsVpnUnicastState to protobuf object *otg.BgpPrefixIpv6MplsVpnUnicastState
	// and doesn't set defaults
	msg() *otg.BgpPrefixIpv6MplsVpnUnicastState
	// setMsg unmarshals BgpPrefixIpv6MplsVpnUnicastState from protobuf object *otg.BgpPrefixIpv6MplsVpnUnicastState
	// and doesn't set defaults
	setMsg(*otg.BgpPrefixIpv6MplsVpnUnicastState) BgpPrefixIpv6MplsVpnUnicastState
	// provides marshal interface
	Marshal() marshalBgpPrefixIpv6MplsVpnUnicastState
	// provides unmarshal interface
	Unmarshal() unMarshalBgpPrefixIpv6MplsVpnUnicastState
	// validate validates BgpPrefixIpv6MplsVpnUnicastState
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpPrefixIpv6MplsVpnUnicastState, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns BgpPrefixIpv6MplsVpnUnicastStateChoiceEnum, set in BgpPrefixIpv6MplsVpnUnicastState
	Choice() BgpPrefixIpv6MplsVpnUnicastStateChoiceEnum
	// setChoice assigns BgpPrefixIpv6MplsVpnUnicastStateChoiceEnum provided by user to BgpPrefixIpv6MplsVpnUnicastState
	setChoice(value BgpPrefixIpv6MplsVpnUnicastStateChoiceEnum) BgpPrefixIpv6MplsVpnUnicastState
	// HasChoice checks if Choice has been set in BgpPrefixIpv6MplsVpnUnicastState
	HasChoice() bool
	// MatchedVrfRoutes returns BgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes, set in BgpPrefixIpv6MplsVpnUnicastState.
	// BgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes is the local VRFs that this received VPN-IPv6 NLRI was imported into via Route Target matching (RFC 4364 Section 4.3.1), each with its own copy of the received path attributes.
	MatchedVrfRoutes() BgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes
	// SetMatchedVrfRoutes assigns BgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes provided by user to BgpPrefixIpv6MplsVpnUnicastState.
	// BgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes is the local VRFs that this received VPN-IPv6 NLRI was imported into via Route Target matching (RFC 4364 Section 4.3.1), each with its own copy of the received path attributes.
	SetMatchedVrfRoutes(value BgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes) BgpPrefixIpv6MplsVpnUnicastState
	// HasMatchedVrfRoutes checks if MatchedVrfRoutes has been set in BgpPrefixIpv6MplsVpnUnicastState
	HasMatchedVrfRoutes() bool
	setNil()
}

type BgpPrefixIpv6MplsVpnUnicastStateChoiceEnum string

// Enum of Choice on BgpPrefixIpv6MplsVpnUnicastState
var BgpPrefixIpv6MplsVpnUnicastStateChoice = struct {
	MATCHED_VRF_ROUTES BgpPrefixIpv6MplsVpnUnicastStateChoiceEnum
}{
	MATCHED_VRF_ROUTES: BgpPrefixIpv6MplsVpnUnicastStateChoiceEnum("matched_vrf_routes"),
}

func (obj *bgpPrefixIpv6MplsVpnUnicastState) Choice() BgpPrefixIpv6MplsVpnUnicastStateChoiceEnum {
	return BgpPrefixIpv6MplsVpnUnicastStateChoiceEnum(obj.obj.Choice.Enum().String())
}

// Selects how this received VPN-IPv6 NLRI is reported. Currently the only supported view groups the prefix by the local VRF(s) (RFC 4364 Section 4.3.1 Route Target import) it was matched into.
// Choice returns a string
func (obj *bgpPrefixIpv6MplsVpnUnicastState) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *bgpPrefixIpv6MplsVpnUnicastState) setChoice(value BgpPrefixIpv6MplsVpnUnicastStateChoiceEnum) BgpPrefixIpv6MplsVpnUnicastState {
	intValue, ok := otg.BgpPrefixIpv6MplsVpnUnicastState_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpPrefixIpv6MplsVpnUnicastStateChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpPrefixIpv6MplsVpnUnicastState_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.MatchedVrfRoutes = nil
	obj.matchedVrfRoutesHolder = nil

	if value == BgpPrefixIpv6MplsVpnUnicastStateChoice.MATCHED_VRF_ROUTES {
		obj.obj.MatchedVrfRoutes = NewBgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes().msg()
	}

	return obj
}

// description is TBD
// MatchedVrfRoutes returns a BgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes
func (obj *bgpPrefixIpv6MplsVpnUnicastState) MatchedVrfRoutes() BgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes {
	if obj.obj.MatchedVrfRoutes == nil {
		obj.setChoice(BgpPrefixIpv6MplsVpnUnicastStateChoice.MATCHED_VRF_ROUTES)
	}
	if obj.matchedVrfRoutesHolder == nil {
		obj.matchedVrfRoutesHolder = &bgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes{obj: obj.obj.MatchedVrfRoutes}
	}
	return obj.matchedVrfRoutesHolder
}

// description is TBD
// MatchedVrfRoutes returns a BgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes
func (obj *bgpPrefixIpv6MplsVpnUnicastState) HasMatchedVrfRoutes() bool {
	return obj.obj.MatchedVrfRoutes != nil
}

// description is TBD
// SetMatchedVrfRoutes sets the BgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes value in the BgpPrefixIpv6MplsVpnUnicastState object
func (obj *bgpPrefixIpv6MplsVpnUnicastState) SetMatchedVrfRoutes(value BgpPrefixIpv6MplsVpnUnicastMatchedVrfRoutes) BgpPrefixIpv6MplsVpnUnicastState {
	obj.setChoice(BgpPrefixIpv6MplsVpnUnicastStateChoice.MATCHED_VRF_ROUTES)
	obj.matchedVrfRoutesHolder = nil
	obj.obj.MatchedVrfRoutes = value.msg()

	return obj
}

func (obj *bgpPrefixIpv6MplsVpnUnicastState) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.MatchedVrfRoutes != nil {

		obj.MatchedVrfRoutes().validateObj(vObj, set_default)
	}

}

func (obj *bgpPrefixIpv6MplsVpnUnicastState) setDefault() {
	var choices_set int = 0
	var choice BgpPrefixIpv6MplsVpnUnicastStateChoiceEnum

	if obj.obj.MatchedVrfRoutes != nil {
		choices_set += 1
		choice = BgpPrefixIpv6MplsVpnUnicastStateChoice.MATCHED_VRF_ROUTES
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(BgpPrefixIpv6MplsVpnUnicastStateChoice.MATCHED_VRF_ROUTES)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in BgpPrefixIpv6MplsVpnUnicastState")
			}
		} else {
			intVal := otg.BgpPrefixIpv6MplsVpnUnicastState_Choice_Enum_value[string(choice)]
			enumValue := otg.BgpPrefixIpv6MplsVpnUnicastState_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
