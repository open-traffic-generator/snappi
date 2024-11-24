package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpEthernetSegmentDfElection *****
type bgpEthernetSegmentDfElection struct {
	validation
	obj          *otg.BgpEthernetSegmentDfElection
	marshaller   marshalBgpEthernetSegmentDfElection
	unMarshaller unMarshalBgpEthernetSegmentDfElection
}

func NewBgpEthernetSegmentDfElection() BgpEthernetSegmentDfElection {
	obj := bgpEthernetSegmentDfElection{obj: &otg.BgpEthernetSegmentDfElection{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpEthernetSegmentDfElection) msg() *otg.BgpEthernetSegmentDfElection {
	return obj.obj
}

func (obj *bgpEthernetSegmentDfElection) setMsg(msg *otg.BgpEthernetSegmentDfElection) BgpEthernetSegmentDfElection {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpEthernetSegmentDfElection struct {
	obj *bgpEthernetSegmentDfElection
}

type marshalBgpEthernetSegmentDfElection interface {
	// ToProto marshals BgpEthernetSegmentDfElection to protobuf object *otg.BgpEthernetSegmentDfElection
	ToProto() (*otg.BgpEthernetSegmentDfElection, error)
	// ToPbText marshals BgpEthernetSegmentDfElection to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpEthernetSegmentDfElection to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpEthernetSegmentDfElection to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals BgpEthernetSegmentDfElection to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalbgpEthernetSegmentDfElection struct {
	obj *bgpEthernetSegmentDfElection
}

type unMarshalBgpEthernetSegmentDfElection interface {
	// FromProto unmarshals BgpEthernetSegmentDfElection from protobuf object *otg.BgpEthernetSegmentDfElection
	FromProto(msg *otg.BgpEthernetSegmentDfElection) (BgpEthernetSegmentDfElection, error)
	// FromPbText unmarshals BgpEthernetSegmentDfElection from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpEthernetSegmentDfElection from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpEthernetSegmentDfElection from JSON text
	FromJson(value string) error
}

func (obj *bgpEthernetSegmentDfElection) Marshal() marshalBgpEthernetSegmentDfElection {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpEthernetSegmentDfElection{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpEthernetSegmentDfElection) Unmarshal() unMarshalBgpEthernetSegmentDfElection {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpEthernetSegmentDfElection{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpEthernetSegmentDfElection) ToProto() (*otg.BgpEthernetSegmentDfElection, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpEthernetSegmentDfElection) FromProto(msg *otg.BgpEthernetSegmentDfElection) (BgpEthernetSegmentDfElection, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpEthernetSegmentDfElection) ToPbText() (string, error) {
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

func (m *unMarshalbgpEthernetSegmentDfElection) FromPbText(value string) error {
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

func (m *marshalbgpEthernetSegmentDfElection) ToYaml() (string, error) {
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

func (m *unMarshalbgpEthernetSegmentDfElection) FromYaml(value string) error {
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

func (m *marshalbgpEthernetSegmentDfElection) ToJsonRaw() (string, error) {
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

func (m *marshalbgpEthernetSegmentDfElection) ToJson() (string, error) {
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

func (m *unMarshalbgpEthernetSegmentDfElection) FromJson(value string) error {
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

func (obj *bgpEthernetSegmentDfElection) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpEthernetSegmentDfElection) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpEthernetSegmentDfElection) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpEthernetSegmentDfElection) Clone() (BgpEthernetSegmentDfElection, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpEthernetSegmentDfElection()
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

// BgpEthernetSegmentDfElection is configuration for Designated Forwarder (DF) election among the Provider Edge (PE) routers on the same Ethernet Segment.
type BgpEthernetSegmentDfElection interface {
	Validation
	// msg marshals BgpEthernetSegmentDfElection to protobuf object *otg.BgpEthernetSegmentDfElection
	// and doesn't set defaults
	msg() *otg.BgpEthernetSegmentDfElection
	// setMsg unmarshals BgpEthernetSegmentDfElection from protobuf object *otg.BgpEthernetSegmentDfElection
	// and doesn't set defaults
	setMsg(*otg.BgpEthernetSegmentDfElection) BgpEthernetSegmentDfElection
	// provides marshal interface
	Marshal() marshalBgpEthernetSegmentDfElection
	// provides unmarshal interface
	Unmarshal() unMarshalBgpEthernetSegmentDfElection
	// validate validates BgpEthernetSegmentDfElection
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpEthernetSegmentDfElection, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// ElectionTimer returns uint32, set in BgpEthernetSegmentDfElection.
	ElectionTimer() uint32
	// SetElectionTimer assigns uint32 provided by user to BgpEthernetSegmentDfElection
	SetElectionTimer(value uint32) BgpEthernetSegmentDfElection
	// HasElectionTimer checks if ElectionTimer has been set in BgpEthernetSegmentDfElection
	HasElectionTimer() bool
}

// The DF election timer in seconds.
// ElectionTimer returns a uint32
func (obj *bgpEthernetSegmentDfElection) ElectionTimer() uint32 {

	return *obj.obj.ElectionTimer

}

// The DF election timer in seconds.
// ElectionTimer returns a uint32
func (obj *bgpEthernetSegmentDfElection) HasElectionTimer() bool {
	return obj.obj.ElectionTimer != nil
}

// The DF election timer in seconds.
// SetElectionTimer sets the uint32 value in the BgpEthernetSegmentDfElection object
func (obj *bgpEthernetSegmentDfElection) SetElectionTimer(value uint32) BgpEthernetSegmentDfElection {

	obj.obj.ElectionTimer = &value
	return obj
}

func (obj *bgpEthernetSegmentDfElection) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.ElectionTimer != nil {

		if *obj.obj.ElectionTimer > 300 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= BgpEthernetSegmentDfElection.ElectionTimer <= 300 but Got %d", *obj.obj.ElectionTimer))
		}

	}

}

func (obj *bgpEthernetSegmentDfElection) setDefault() {
	if obj.obj.ElectionTimer == nil {
		obj.SetElectionTimer(3)
	}

}
