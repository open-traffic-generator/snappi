package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisMetric *****
type isisMetric struct {
	validation
	obj          *otg.IsisMetric
	marshaller   marshalIsisMetric
	unMarshaller unMarshalIsisMetric
}

func NewIsisMetric() IsisMetric {
	obj := isisMetric{obj: &otg.IsisMetric{}}
	obj.setDefault()
	return &obj
}

func (obj *isisMetric) msg() *otg.IsisMetric {
	return obj.obj
}

func (obj *isisMetric) setMsg(msg *otg.IsisMetric) IsisMetric {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisMetric struct {
	obj *isisMetric
}

type marshalIsisMetric interface {
	// ToProto marshals IsisMetric to protobuf object *otg.IsisMetric
	ToProto() (*otg.IsisMetric, error)
	// ToPbText marshals IsisMetric to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisMetric to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisMetric to JSON text
	ToJson() (string, error)
}

type unMarshalisisMetric struct {
	obj *isisMetric
}

type unMarshalIsisMetric interface {
	// FromProto unmarshals IsisMetric from protobuf object *otg.IsisMetric
	FromProto(msg *otg.IsisMetric) (IsisMetric, error)
	// FromPbText unmarshals IsisMetric from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisMetric from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisMetric from JSON text
	FromJson(value string) error
}

func (obj *isisMetric) Marshal() marshalIsisMetric {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisMetric{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisMetric) Unmarshal() unMarshalIsisMetric {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisMetric{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisMetric) ToProto() (*otg.IsisMetric, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisMetric) FromProto(msg *otg.IsisMetric) (IsisMetric, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisMetric) ToPbText() (string, error) {
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

func (m *unMarshalisisMetric) FromPbText(value string) error {
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

func (m *marshalisisMetric) ToYaml() (string, error) {
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

func (m *unMarshalisisMetric) FromYaml(value string) error {
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

func (m *marshalisisMetric) ToJson() (string, error) {
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

func (m *unMarshalisisMetric) FromJson(value string) error {
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

func (obj *isisMetric) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisMetric) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisMetric) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisMetric) Clone() (IsisMetric, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisMetric()
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

// IsisMetric is iSIS per router statistics information.
type IsisMetric interface {
	Validation
	// msg marshals IsisMetric to protobuf object *otg.IsisMetric
	// and doesn't set defaults
	msg() *otg.IsisMetric
	// setMsg unmarshals IsisMetric from protobuf object *otg.IsisMetric
	// and doesn't set defaults
	setMsg(*otg.IsisMetric) IsisMetric
	// provides marshal interface
	Marshal() marshalIsisMetric
	// provides unmarshal interface
	Unmarshal() unMarshalIsisMetric
	// validate validates IsisMetric
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisMetric, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in IsisMetric.
	Name() string
	// SetName assigns string provided by user to IsisMetric
	SetName(value string) IsisMetric
	// HasName checks if Name has been set in IsisMetric
	HasName() bool
	// L1SessionsUp returns uint32, set in IsisMetric.
	L1SessionsUp() uint32
	// SetL1SessionsUp assigns uint32 provided by user to IsisMetric
	SetL1SessionsUp(value uint32) IsisMetric
	// HasL1SessionsUp checks if L1SessionsUp has been set in IsisMetric
	HasL1SessionsUp() bool
	// L1SessionFlap returns uint64, set in IsisMetric.
	L1SessionFlap() uint64
	// SetL1SessionFlap assigns uint64 provided by user to IsisMetric
	SetL1SessionFlap(value uint64) IsisMetric
	// HasL1SessionFlap checks if L1SessionFlap has been set in IsisMetric
	HasL1SessionFlap() bool
	// L1BroadcastHellosSent returns uint64, set in IsisMetric.
	L1BroadcastHellosSent() uint64
	// SetL1BroadcastHellosSent assigns uint64 provided by user to IsisMetric
	SetL1BroadcastHellosSent(value uint64) IsisMetric
	// HasL1BroadcastHellosSent checks if L1BroadcastHellosSent has been set in IsisMetric
	HasL1BroadcastHellosSent() bool
	// L1BroadcastHellosReceived returns uint64, set in IsisMetric.
	L1BroadcastHellosReceived() uint64
	// SetL1BroadcastHellosReceived assigns uint64 provided by user to IsisMetric
	SetL1BroadcastHellosReceived(value uint64) IsisMetric
	// HasL1BroadcastHellosReceived checks if L1BroadcastHellosReceived has been set in IsisMetric
	HasL1BroadcastHellosReceived() bool
	// L1PointToPointHellosSent returns uint64, set in IsisMetric.
	L1PointToPointHellosSent() uint64
	// SetL1PointToPointHellosSent assigns uint64 provided by user to IsisMetric
	SetL1PointToPointHellosSent(value uint64) IsisMetric
	// HasL1PointToPointHellosSent checks if L1PointToPointHellosSent has been set in IsisMetric
	HasL1PointToPointHellosSent() bool
	// L1PointToPointHellosReceived returns uint64, set in IsisMetric.
	L1PointToPointHellosReceived() uint64
	// SetL1PointToPointHellosReceived assigns uint64 provided by user to IsisMetric
	SetL1PointToPointHellosReceived(value uint64) IsisMetric
	// HasL1PointToPointHellosReceived checks if L1PointToPointHellosReceived has been set in IsisMetric
	HasL1PointToPointHellosReceived() bool
	// L1DatabaseSize returns uint64, set in IsisMetric.
	L1DatabaseSize() uint64
	// SetL1DatabaseSize assigns uint64 provided by user to IsisMetric
	SetL1DatabaseSize(value uint64) IsisMetric
	// HasL1DatabaseSize checks if L1DatabaseSize has been set in IsisMetric
	HasL1DatabaseSize() bool
	// L1PsnpSent returns uint64, set in IsisMetric.
	L1PsnpSent() uint64
	// SetL1PsnpSent assigns uint64 provided by user to IsisMetric
	SetL1PsnpSent(value uint64) IsisMetric
	// HasL1PsnpSent checks if L1PsnpSent has been set in IsisMetric
	HasL1PsnpSent() bool
	// L1PsnpReceived returns uint64, set in IsisMetric.
	L1PsnpReceived() uint64
	// SetL1PsnpReceived assigns uint64 provided by user to IsisMetric
	SetL1PsnpReceived(value uint64) IsisMetric
	// HasL1PsnpReceived checks if L1PsnpReceived has been set in IsisMetric
	HasL1PsnpReceived() bool
	// L1CsnpSent returns uint64, set in IsisMetric.
	L1CsnpSent() uint64
	// SetL1CsnpSent assigns uint64 provided by user to IsisMetric
	SetL1CsnpSent(value uint64) IsisMetric
	// HasL1CsnpSent checks if L1CsnpSent has been set in IsisMetric
	HasL1CsnpSent() bool
	// L1CsnpReceived returns uint64, set in IsisMetric.
	L1CsnpReceived() uint64
	// SetL1CsnpReceived assigns uint64 provided by user to IsisMetric
	SetL1CsnpReceived(value uint64) IsisMetric
	// HasL1CsnpReceived checks if L1CsnpReceived has been set in IsisMetric
	HasL1CsnpReceived() bool
	// L1LspSent returns uint64, set in IsisMetric.
	L1LspSent() uint64
	// SetL1LspSent assigns uint64 provided by user to IsisMetric
	SetL1LspSent(value uint64) IsisMetric
	// HasL1LspSent checks if L1LspSent has been set in IsisMetric
	HasL1LspSent() bool
	// L1LspReceived returns uint64, set in IsisMetric.
	L1LspReceived() uint64
	// SetL1LspReceived assigns uint64 provided by user to IsisMetric
	SetL1LspReceived(value uint64) IsisMetric
	// HasL1LspReceived checks if L1LspReceived has been set in IsisMetric
	HasL1LspReceived() bool
	// L2SessionsUp returns uint32, set in IsisMetric.
	L2SessionsUp() uint32
	// SetL2SessionsUp assigns uint32 provided by user to IsisMetric
	SetL2SessionsUp(value uint32) IsisMetric
	// HasL2SessionsUp checks if L2SessionsUp has been set in IsisMetric
	HasL2SessionsUp() bool
	// L2SessionFlap returns uint64, set in IsisMetric.
	L2SessionFlap() uint64
	// SetL2SessionFlap assigns uint64 provided by user to IsisMetric
	SetL2SessionFlap(value uint64) IsisMetric
	// HasL2SessionFlap checks if L2SessionFlap has been set in IsisMetric
	HasL2SessionFlap() bool
	// L2BroadcastHellosSent returns uint64, set in IsisMetric.
	L2BroadcastHellosSent() uint64
	// SetL2BroadcastHellosSent assigns uint64 provided by user to IsisMetric
	SetL2BroadcastHellosSent(value uint64) IsisMetric
	// HasL2BroadcastHellosSent checks if L2BroadcastHellosSent has been set in IsisMetric
	HasL2BroadcastHellosSent() bool
	// L2BroadcastHellosReceived returns uint64, set in IsisMetric.
	L2BroadcastHellosReceived() uint64
	// SetL2BroadcastHellosReceived assigns uint64 provided by user to IsisMetric
	SetL2BroadcastHellosReceived(value uint64) IsisMetric
	// HasL2BroadcastHellosReceived checks if L2BroadcastHellosReceived has been set in IsisMetric
	HasL2BroadcastHellosReceived() bool
	// L2PointToPointHellosSent returns uint64, set in IsisMetric.
	L2PointToPointHellosSent() uint64
	// SetL2PointToPointHellosSent assigns uint64 provided by user to IsisMetric
	SetL2PointToPointHellosSent(value uint64) IsisMetric
	// HasL2PointToPointHellosSent checks if L2PointToPointHellosSent has been set in IsisMetric
	HasL2PointToPointHellosSent() bool
	// L2PointToPointHellosReceived returns uint64, set in IsisMetric.
	L2PointToPointHellosReceived() uint64
	// SetL2PointToPointHellosReceived assigns uint64 provided by user to IsisMetric
	SetL2PointToPointHellosReceived(value uint64) IsisMetric
	// HasL2PointToPointHellosReceived checks if L2PointToPointHellosReceived has been set in IsisMetric
	HasL2PointToPointHellosReceived() bool
	// L2DatabaseSize returns uint64, set in IsisMetric.
	L2DatabaseSize() uint64
	// SetL2DatabaseSize assigns uint64 provided by user to IsisMetric
	SetL2DatabaseSize(value uint64) IsisMetric
	// HasL2DatabaseSize checks if L2DatabaseSize has been set in IsisMetric
	HasL2DatabaseSize() bool
	// L2PsnpSent returns uint64, set in IsisMetric.
	L2PsnpSent() uint64
	// SetL2PsnpSent assigns uint64 provided by user to IsisMetric
	SetL2PsnpSent(value uint64) IsisMetric
	// HasL2PsnpSent checks if L2PsnpSent has been set in IsisMetric
	HasL2PsnpSent() bool
	// L2PsnpReceived returns uint64, set in IsisMetric.
	L2PsnpReceived() uint64
	// SetL2PsnpReceived assigns uint64 provided by user to IsisMetric
	SetL2PsnpReceived(value uint64) IsisMetric
	// HasL2PsnpReceived checks if L2PsnpReceived has been set in IsisMetric
	HasL2PsnpReceived() bool
	// L2CsnpSent returns uint64, set in IsisMetric.
	L2CsnpSent() uint64
	// SetL2CsnpSent assigns uint64 provided by user to IsisMetric
	SetL2CsnpSent(value uint64) IsisMetric
	// HasL2CsnpSent checks if L2CsnpSent has been set in IsisMetric
	HasL2CsnpSent() bool
	// L2CsnpReceived returns uint64, set in IsisMetric.
	L2CsnpReceived() uint64
	// SetL2CsnpReceived assigns uint64 provided by user to IsisMetric
	SetL2CsnpReceived(value uint64) IsisMetric
	// HasL2CsnpReceived checks if L2CsnpReceived has been set in IsisMetric
	HasL2CsnpReceived() bool
	// L2LspSent returns uint64, set in IsisMetric.
	L2LspSent() uint64
	// SetL2LspSent assigns uint64 provided by user to IsisMetric
	SetL2LspSent(value uint64) IsisMetric
	// HasL2LspSent checks if L2LspSent has been set in IsisMetric
	HasL2LspSent() bool
	// L2LspReceived returns uint64, set in IsisMetric.
	L2LspReceived() uint64
	// SetL2LspReceived assigns uint64 provided by user to IsisMetric
	SetL2LspReceived(value uint64) IsisMetric
	// HasL2LspReceived checks if L2LspReceived has been set in IsisMetric
	HasL2LspReceived() bool
	// GrInitiated returns uint64, set in IsisMetric.
	GrInitiated() uint64
	// SetGrInitiated assigns uint64 provided by user to IsisMetric
	SetGrInitiated(value uint64) IsisMetric
	// HasGrInitiated checks if GrInitiated has been set in IsisMetric
	HasGrInitiated() bool
	// GrSucceeded returns uint64, set in IsisMetric.
	GrSucceeded() uint64
	// SetGrSucceeded assigns uint64 provided by user to IsisMetric
	SetGrSucceeded(value uint64) IsisMetric
	// HasGrSucceeded checks if GrSucceeded has been set in IsisMetric
	HasGrSucceeded() bool
	// NeighborGrInitiated returns uint64, set in IsisMetric.
	NeighborGrInitiated() uint64
	// SetNeighborGrInitiated assigns uint64 provided by user to IsisMetric
	SetNeighborGrInitiated(value uint64) IsisMetric
	// HasNeighborGrInitiated checks if NeighborGrInitiated has been set in IsisMetric
	HasNeighborGrInitiated() bool
	// NeighborGrSucceeded returns uint64, set in IsisMetric.
	NeighborGrSucceeded() uint64
	// SetNeighborGrSucceeded assigns uint64 provided by user to IsisMetric
	SetNeighborGrSucceeded(value uint64) IsisMetric
	// HasNeighborGrSucceeded checks if NeighborGrSucceeded has been set in IsisMetric
	HasNeighborGrSucceeded() bool
}

// The name of a configured ISIS router.
// Name returns a string
func (obj *isisMetric) Name() string {

	return *obj.obj.Name

}

// The name of a configured ISIS router.
// Name returns a string
func (obj *isisMetric) HasName() bool {
	return obj.obj.Name != nil
}

// The name of a configured ISIS router.
// SetName sets the string value in the IsisMetric object
func (obj *isisMetric) SetName(value string) IsisMetric {

	obj.obj.Name = &value
	return obj
}

// The number of Level 1 (L1) sessions that are fully up.
// L1SessionsUp returns a uint32
func (obj *isisMetric) L1SessionsUp() uint32 {

	return *obj.obj.L1SessionsUp

}

// The number of Level 1 (L1) sessions that are fully up.
// L1SessionsUp returns a uint32
func (obj *isisMetric) HasL1SessionsUp() bool {
	return obj.obj.L1SessionsUp != nil
}

// The number of Level 1 (L1) sessions that are fully up.
// SetL1SessionsUp sets the uint32 value in the IsisMetric object
func (obj *isisMetric) SetL1SessionsUp(value uint32) IsisMetric {

	obj.obj.L1SessionsUp = &value
	return obj
}

// The number of Level 1 Sessions Flap.
// L1SessionFlap returns a uint64
func (obj *isisMetric) L1SessionFlap() uint64 {

	return *obj.obj.L1SessionFlap

}

// The number of Level 1 Sessions Flap.
// L1SessionFlap returns a uint64
func (obj *isisMetric) HasL1SessionFlap() bool {
	return obj.obj.L1SessionFlap != nil
}

// The number of Level 1 Sessions Flap.
// SetL1SessionFlap sets the uint64 value in the IsisMetric object
func (obj *isisMetric) SetL1SessionFlap(value uint64) IsisMetric {

	obj.obj.L1SessionFlap = &value
	return obj
}

// Number of Level 1 Hello messages sent.
// L1BroadcastHellosSent returns a uint64
func (obj *isisMetric) L1BroadcastHellosSent() uint64 {

	return *obj.obj.L1BroadcastHellosSent

}

// Number of Level 1 Hello messages sent.
// L1BroadcastHellosSent returns a uint64
func (obj *isisMetric) HasL1BroadcastHellosSent() bool {
	return obj.obj.L1BroadcastHellosSent != nil
}

// Number of Level 1 Hello messages sent.
// SetL1BroadcastHellosSent sets the uint64 value in the IsisMetric object
func (obj *isisMetric) SetL1BroadcastHellosSent(value uint64) IsisMetric {

	obj.obj.L1BroadcastHellosSent = &value
	return obj
}

// Number of Level 1 Hello messages received.
// L1BroadcastHellosReceived returns a uint64
func (obj *isisMetric) L1BroadcastHellosReceived() uint64 {

	return *obj.obj.L1BroadcastHellosReceived

}

// Number of Level 1 Hello messages received.
// L1BroadcastHellosReceived returns a uint64
func (obj *isisMetric) HasL1BroadcastHellosReceived() bool {
	return obj.obj.L1BroadcastHellosReceived != nil
}

// Number of Level 1 Hello messages received.
// SetL1BroadcastHellosReceived sets the uint64 value in the IsisMetric object
func (obj *isisMetric) SetL1BroadcastHellosReceived(value uint64) IsisMetric {

	obj.obj.L1BroadcastHellosReceived = &value
	return obj
}

// Number of Level 1 Point-to-Point(P2P) Hello messages sent.
// L1PointToPointHellosSent returns a uint64
func (obj *isisMetric) L1PointToPointHellosSent() uint64 {

	return *obj.obj.L1PointToPointHellosSent

}

// Number of Level 1 Point-to-Point(P2P) Hello messages sent.
// L1PointToPointHellosSent returns a uint64
func (obj *isisMetric) HasL1PointToPointHellosSent() bool {
	return obj.obj.L1PointToPointHellosSent != nil
}

// Number of Level 1 Point-to-Point(P2P) Hello messages sent.
// SetL1PointToPointHellosSent sets the uint64 value in the IsisMetric object
func (obj *isisMetric) SetL1PointToPointHellosSent(value uint64) IsisMetric {

	obj.obj.L1PointToPointHellosSent = &value
	return obj
}

// Number of Level 1 Point-to-Point(P2P) Hello messages received.
// L1PointToPointHellosReceived returns a uint64
func (obj *isisMetric) L1PointToPointHellosReceived() uint64 {

	return *obj.obj.L1PointToPointHellosReceived

}

// Number of Level 1 Point-to-Point(P2P) Hello messages received.
// L1PointToPointHellosReceived returns a uint64
func (obj *isisMetric) HasL1PointToPointHellosReceived() bool {
	return obj.obj.L1PointToPointHellosReceived != nil
}

// Number of Level 1 Point-to-Point(P2P) Hello messages received.
// SetL1PointToPointHellosReceived sets the uint64 value in the IsisMetric object
func (obj *isisMetric) SetL1PointToPointHellosReceived(value uint64) IsisMetric {

	obj.obj.L1PointToPointHellosReceived = &value
	return obj
}

// Number of Link State Updates (LSPs) in the Level 1 LSP Databases.
// L1DatabaseSize returns a uint64
func (obj *isisMetric) L1DatabaseSize() uint64 {

	return *obj.obj.L1DatabaseSize

}

// Number of Link State Updates (LSPs) in the Level 1 LSP Databases.
// L1DatabaseSize returns a uint64
func (obj *isisMetric) HasL1DatabaseSize() bool {
	return obj.obj.L1DatabaseSize != nil
}

// Number of Link State Updates (LSPs) in the Level 1 LSP Databases.
// SetL1DatabaseSize sets the uint64 value in the IsisMetric object
func (obj *isisMetric) SetL1DatabaseSize(value uint64) IsisMetric {

	obj.obj.L1DatabaseSize = &value
	return obj
}

// Number of Level 1 (L1) Partial Sequence Number Packet (PSNPs) sent.
// L1PsnpSent returns a uint64
func (obj *isisMetric) L1PsnpSent() uint64 {

	return *obj.obj.L1PsnpSent

}

// Number of Level 1 (L1) Partial Sequence Number Packet (PSNPs) sent.
// L1PsnpSent returns a uint64
func (obj *isisMetric) HasL1PsnpSent() bool {
	return obj.obj.L1PsnpSent != nil
}

// Number of Level 1 (L1) Partial Sequence Number Packet (PSNPs) sent.
// SetL1PsnpSent sets the uint64 value in the IsisMetric object
func (obj *isisMetric) SetL1PsnpSent(value uint64) IsisMetric {

	obj.obj.L1PsnpSent = &value
	return obj
}

// Number of Level 1 (L1) Complete Sequence Number Packet (PSNPs) received.
// L1PsnpReceived returns a uint64
func (obj *isisMetric) L1PsnpReceived() uint64 {

	return *obj.obj.L1PsnpReceived

}

// Number of Level 1 (L1) Complete Sequence Number Packet (PSNPs) received.
// L1PsnpReceived returns a uint64
func (obj *isisMetric) HasL1PsnpReceived() bool {
	return obj.obj.L1PsnpReceived != nil
}

// Number of Level 1 (L1) Complete Sequence Number Packet (PSNPs) received.
// SetL1PsnpReceived sets the uint64 value in the IsisMetric object
func (obj *isisMetric) SetL1PsnpReceived(value uint64) IsisMetric {

	obj.obj.L1PsnpReceived = &value
	return obj
}

// Number of Level 1 (L1) Complete Sequence Number Packet (CSNPs) sent.
// L1CsnpSent returns a uint64
func (obj *isisMetric) L1CsnpSent() uint64 {

	return *obj.obj.L1CsnpSent

}

// Number of Level 1 (L1) Complete Sequence Number Packet (CSNPs) sent.
// L1CsnpSent returns a uint64
func (obj *isisMetric) HasL1CsnpSent() bool {
	return obj.obj.L1CsnpSent != nil
}

// Number of Level 1 (L1) Complete Sequence Number Packet (CSNPs) sent.
// SetL1CsnpSent sets the uint64 value in the IsisMetric object
func (obj *isisMetric) SetL1CsnpSent(value uint64) IsisMetric {

	obj.obj.L1CsnpSent = &value
	return obj
}

// Number of Level 1 (L1) Complete Sequence Number Packet (CSNPs) received.
// L1CsnpReceived returns a uint64
func (obj *isisMetric) L1CsnpReceived() uint64 {

	return *obj.obj.L1CsnpReceived

}

// Number of Level 1 (L1) Complete Sequence Number Packet (CSNPs) received.
// L1CsnpReceived returns a uint64
func (obj *isisMetric) HasL1CsnpReceived() bool {
	return obj.obj.L1CsnpReceived != nil
}

// Number of Level 1 (L1) Complete Sequence Number Packet (CSNPs) received.
// SetL1CsnpReceived sets the uint64 value in the IsisMetric object
func (obj *isisMetric) SetL1CsnpReceived(value uint64) IsisMetric {

	obj.obj.L1CsnpReceived = &value
	return obj
}

// Number of Level 1 (L1) Link State Protocol Data Units (LSPs) sent.
// L1LspSent returns a uint64
func (obj *isisMetric) L1LspSent() uint64 {

	return *obj.obj.L1LspSent

}

// Number of Level 1 (L1) Link State Protocol Data Units (LSPs) sent.
// L1LspSent returns a uint64
func (obj *isisMetric) HasL1LspSent() bool {
	return obj.obj.L1LspSent != nil
}

// Number of Level 1 (L1) Link State Protocol Data Units (LSPs) sent.
// SetL1LspSent sets the uint64 value in the IsisMetric object
func (obj *isisMetric) SetL1LspSent(value uint64) IsisMetric {

	obj.obj.L1LspSent = &value
	return obj
}

// Number of Level 1 (L1) Link State Protocol Data Units (LSPs) received.
// L1LspReceived returns a uint64
func (obj *isisMetric) L1LspReceived() uint64 {

	return *obj.obj.L1LspReceived

}

// Number of Level 1 (L1) Link State Protocol Data Units (LSPs) received.
// L1LspReceived returns a uint64
func (obj *isisMetric) HasL1LspReceived() bool {
	return obj.obj.L1LspReceived != nil
}

// Number of Level 1 (L1) Link State Protocol Data Units (LSPs) received.
// SetL1LspReceived sets the uint64 value in the IsisMetric object
func (obj *isisMetric) SetL1LspReceived(value uint64) IsisMetric {

	obj.obj.L1LspReceived = &value
	return obj
}

// The number of Level 2 (L2) sessions that are fully up.
// L2SessionsUp returns a uint32
func (obj *isisMetric) L2SessionsUp() uint32 {

	return *obj.obj.L2SessionsUp

}

// The number of Level 2 (L2) sessions that are fully up.
// L2SessionsUp returns a uint32
func (obj *isisMetric) HasL2SessionsUp() bool {
	return obj.obj.L2SessionsUp != nil
}

// The number of Level 2 (L2) sessions that are fully up.
// SetL2SessionsUp sets the uint32 value in the IsisMetric object
func (obj *isisMetric) SetL2SessionsUp(value uint32) IsisMetric {

	obj.obj.L2SessionsUp = &value
	return obj
}

// The number of Level 2 Sessions Flap.
// L2SessionFlap returns a uint64
func (obj *isisMetric) L2SessionFlap() uint64 {

	return *obj.obj.L2SessionFlap

}

// The number of Level 2 Sessions Flap.
// L2SessionFlap returns a uint64
func (obj *isisMetric) HasL2SessionFlap() bool {
	return obj.obj.L2SessionFlap != nil
}

// The number of Level 2 Sessions Flap.
// SetL2SessionFlap sets the uint64 value in the IsisMetric object
func (obj *isisMetric) SetL2SessionFlap(value uint64) IsisMetric {

	obj.obj.L2SessionFlap = &value
	return obj
}

// Number of Level 2 Hello messages sent.
// L2BroadcastHellosSent returns a uint64
func (obj *isisMetric) L2BroadcastHellosSent() uint64 {

	return *obj.obj.L2BroadcastHellosSent

}

// Number of Level 2 Hello messages sent.
// L2BroadcastHellosSent returns a uint64
func (obj *isisMetric) HasL2BroadcastHellosSent() bool {
	return obj.obj.L2BroadcastHellosSent != nil
}

// Number of Level 2 Hello messages sent.
// SetL2BroadcastHellosSent sets the uint64 value in the IsisMetric object
func (obj *isisMetric) SetL2BroadcastHellosSent(value uint64) IsisMetric {

	obj.obj.L2BroadcastHellosSent = &value
	return obj
}

// Number of Level 2 Hello messages received.
// L2BroadcastHellosReceived returns a uint64
func (obj *isisMetric) L2BroadcastHellosReceived() uint64 {

	return *obj.obj.L2BroadcastHellosReceived

}

// Number of Level 2 Hello messages received.
// L2BroadcastHellosReceived returns a uint64
func (obj *isisMetric) HasL2BroadcastHellosReceived() bool {
	return obj.obj.L2BroadcastHellosReceived != nil
}

// Number of Level 2 Hello messages received.
// SetL2BroadcastHellosReceived sets the uint64 value in the IsisMetric object
func (obj *isisMetric) SetL2BroadcastHellosReceived(value uint64) IsisMetric {

	obj.obj.L2BroadcastHellosReceived = &value
	return obj
}

// Number of Level 2 Point-to-Point(P2P) Hello messages sent.
// L2PointToPointHellosSent returns a uint64
func (obj *isisMetric) L2PointToPointHellosSent() uint64 {

	return *obj.obj.L2PointToPointHellosSent

}

// Number of Level 2 Point-to-Point(P2P) Hello messages sent.
// L2PointToPointHellosSent returns a uint64
func (obj *isisMetric) HasL2PointToPointHellosSent() bool {
	return obj.obj.L2PointToPointHellosSent != nil
}

// Number of Level 2 Point-to-Point(P2P) Hello messages sent.
// SetL2PointToPointHellosSent sets the uint64 value in the IsisMetric object
func (obj *isisMetric) SetL2PointToPointHellosSent(value uint64) IsisMetric {

	obj.obj.L2PointToPointHellosSent = &value
	return obj
}

// Number of Level 2 Point-to-Point(P2P) Hello messages received.
// L2PointToPointHellosReceived returns a uint64
func (obj *isisMetric) L2PointToPointHellosReceived() uint64 {

	return *obj.obj.L2PointToPointHellosReceived

}

// Number of Level 2 Point-to-Point(P2P) Hello messages received.
// L2PointToPointHellosReceived returns a uint64
func (obj *isisMetric) HasL2PointToPointHellosReceived() bool {
	return obj.obj.L2PointToPointHellosReceived != nil
}

// Number of Level 2 Point-to-Point(P2P) Hello messages received.
// SetL2PointToPointHellosReceived sets the uint64 value in the IsisMetric object
func (obj *isisMetric) SetL2PointToPointHellosReceived(value uint64) IsisMetric {

	obj.obj.L2PointToPointHellosReceived = &value
	return obj
}

// Number of Link State Updates (LSPs) in the Level 2 LSP Databases.
// L2DatabaseSize returns a uint64
func (obj *isisMetric) L2DatabaseSize() uint64 {

	return *obj.obj.L2DatabaseSize

}

// Number of Link State Updates (LSPs) in the Level 2 LSP Databases.
// L2DatabaseSize returns a uint64
func (obj *isisMetric) HasL2DatabaseSize() bool {
	return obj.obj.L2DatabaseSize != nil
}

// Number of Link State Updates (LSPs) in the Level 2 LSP Databases.
// SetL2DatabaseSize sets the uint64 value in the IsisMetric object
func (obj *isisMetric) SetL2DatabaseSize(value uint64) IsisMetric {

	obj.obj.L2DatabaseSize = &value
	return obj
}

// Number of Level 2 (L2) Partial Sequence Number Packet (PSNPs) sent.
// L2PsnpSent returns a uint64
func (obj *isisMetric) L2PsnpSent() uint64 {

	return *obj.obj.L2PsnpSent

}

// Number of Level 2 (L2) Partial Sequence Number Packet (PSNPs) sent.
// L2PsnpSent returns a uint64
func (obj *isisMetric) HasL2PsnpSent() bool {
	return obj.obj.L2PsnpSent != nil
}

// Number of Level 2 (L2) Partial Sequence Number Packet (PSNPs) sent.
// SetL2PsnpSent sets the uint64 value in the IsisMetric object
func (obj *isisMetric) SetL2PsnpSent(value uint64) IsisMetric {

	obj.obj.L2PsnpSent = &value
	return obj
}

// Number of Level 2 (L2) Complete Sequence Number Packet (PSNPs) received.
// L2PsnpReceived returns a uint64
func (obj *isisMetric) L2PsnpReceived() uint64 {

	return *obj.obj.L2PsnpReceived

}

// Number of Level 2 (L2) Complete Sequence Number Packet (PSNPs) received.
// L2PsnpReceived returns a uint64
func (obj *isisMetric) HasL2PsnpReceived() bool {
	return obj.obj.L2PsnpReceived != nil
}

// Number of Level 2 (L2) Complete Sequence Number Packet (PSNPs) received.
// SetL2PsnpReceived sets the uint64 value in the IsisMetric object
func (obj *isisMetric) SetL2PsnpReceived(value uint64) IsisMetric {

	obj.obj.L2PsnpReceived = &value
	return obj
}

// Number of Level 2 (L2) Complete Sequence Number Packet (CSNPs) sent.
// L2CsnpSent returns a uint64
func (obj *isisMetric) L2CsnpSent() uint64 {

	return *obj.obj.L2CsnpSent

}

// Number of Level 2 (L2) Complete Sequence Number Packet (CSNPs) sent.
// L2CsnpSent returns a uint64
func (obj *isisMetric) HasL2CsnpSent() bool {
	return obj.obj.L2CsnpSent != nil
}

// Number of Level 2 (L2) Complete Sequence Number Packet (CSNPs) sent.
// SetL2CsnpSent sets the uint64 value in the IsisMetric object
func (obj *isisMetric) SetL2CsnpSent(value uint64) IsisMetric {

	obj.obj.L2CsnpSent = &value
	return obj
}

// Number of Level 2 (L2) Complete Sequence Number Packet (CSNPs) received.
// L2CsnpReceived returns a uint64
func (obj *isisMetric) L2CsnpReceived() uint64 {

	return *obj.obj.L2CsnpReceived

}

// Number of Level 2 (L2) Complete Sequence Number Packet (CSNPs) received.
// L2CsnpReceived returns a uint64
func (obj *isisMetric) HasL2CsnpReceived() bool {
	return obj.obj.L2CsnpReceived != nil
}

// Number of Level 2 (L2) Complete Sequence Number Packet (CSNPs) received.
// SetL2CsnpReceived sets the uint64 value in the IsisMetric object
func (obj *isisMetric) SetL2CsnpReceived(value uint64) IsisMetric {

	obj.obj.L2CsnpReceived = &value
	return obj
}

// Number of Level 2 (L2) Link State Protocol Data Units (LSPs) sent.
// L2LspSent returns a uint64
func (obj *isisMetric) L2LspSent() uint64 {

	return *obj.obj.L2LspSent

}

// Number of Level 2 (L2) Link State Protocol Data Units (LSPs) sent.
// L2LspSent returns a uint64
func (obj *isisMetric) HasL2LspSent() bool {
	return obj.obj.L2LspSent != nil
}

// Number of Level 2 (L2) Link State Protocol Data Units (LSPs) sent.
// SetL2LspSent sets the uint64 value in the IsisMetric object
func (obj *isisMetric) SetL2LspSent(value uint64) IsisMetric {

	obj.obj.L2LspSent = &value
	return obj
}

// Number of Level 2 (L2) Link State Protocol Data Units (LSPs) received.
// L2LspReceived returns a uint64
func (obj *isisMetric) L2LspReceived() uint64 {

	return *obj.obj.L2LspReceived

}

// Number of Level 2 (L2) Link State Protocol Data Units (LSPs) received.
// L2LspReceived returns a uint64
func (obj *isisMetric) HasL2LspReceived() bool {
	return obj.obj.L2LspReceived != nil
}

// Number of Level 2 (L2) Link State Protocol Data Units (LSPs) received.
// SetL2LspReceived sets the uint64 value in the IsisMetric object
func (obj *isisMetric) SetL2LspReceived(value uint64) IsisMetric {

	obj.obj.L2LspReceived = &value
	return obj
}

// Number of Graceful Restarts that were initiated by this router.
// GrInitiated returns a uint64
func (obj *isisMetric) GrInitiated() uint64 {

	return *obj.obj.GrInitiated

}

// Number of Graceful Restarts that were initiated by this router.
// GrInitiated returns a uint64
func (obj *isisMetric) HasGrInitiated() bool {
	return obj.obj.GrInitiated != nil
}

// Number of Graceful Restarts that were initiated by this router.
// SetGrInitiated sets the uint64 value in the IsisMetric object
func (obj *isisMetric) SetGrInitiated(value uint64) IsisMetric {

	obj.obj.GrInitiated = &value
	return obj
}

// Number of Graceful Restarts succeeded that were initiated by a this router. This counter is incremented if the Graceful Restart completes succesfully before the T3 timer expires. Timer T3 is maintained for the entire system after which  the router will declare that it has failed to achieve database synchronization.
// GrSucceeded returns a uint64
func (obj *isisMetric) GrSucceeded() uint64 {

	return *obj.obj.GrSucceeded

}

// Number of Graceful Restarts succeeded that were initiated by a this router. This counter is incremented if the Graceful Restart completes succesfully before the T3 timer expires. Timer T3 is maintained for the entire system after which  the router will declare that it has failed to achieve database synchronization.
// GrSucceeded returns a uint64
func (obj *isisMetric) HasGrSucceeded() bool {
	return obj.obj.GrSucceeded != nil
}

// Number of Graceful Restarts succeeded that were initiated by a this router. This counter is incremented if the Graceful Restart completes succesfully before the T3 timer expires. Timer T3 is maintained for the entire system after which  the router will declare that it has failed to achieve database synchronization.
// SetGrSucceeded sets the uint64 value in the IsisMetric object
func (obj *isisMetric) SetGrSucceeded(value uint64) IsisMetric {

	obj.obj.GrSucceeded = &value
	return obj
}

// Number of Graceful Restarts that were initiated by a Neighbor. This counter is incremented for Restart TLV having RR bit set in the received IIH PDU.
// NeighborGrInitiated returns a uint64
func (obj *isisMetric) NeighborGrInitiated() uint64 {

	return *obj.obj.NeighborGrInitiated

}

// Number of Graceful Restarts that were initiated by a Neighbor. This counter is incremented for Restart TLV having RR bit set in the received IIH PDU.
// NeighborGrInitiated returns a uint64
func (obj *isisMetric) HasNeighborGrInitiated() bool {
	return obj.obj.NeighborGrInitiated != nil
}

// Number of Graceful Restarts that were initiated by a Neighbor. This counter is incremented for Restart TLV having RR bit set in the received IIH PDU.
// SetNeighborGrInitiated sets the uint64 value in the IsisMetric object
func (obj *isisMetric) SetNeighborGrInitiated(value uint64) IsisMetric {

	obj.obj.NeighborGrInitiated = &value
	return obj
}

// Number of Graceful Restarts succeeded that were initiated by a Neighbor. This counter is incremented when Restart TLV having RR bit unset in the received IIH PDU after the Graceful Restart was initiated by a Neighbor.
// NeighborGrSucceeded returns a uint64
func (obj *isisMetric) NeighborGrSucceeded() uint64 {

	return *obj.obj.NeighborGrSucceeded

}

// Number of Graceful Restarts succeeded that were initiated by a Neighbor. This counter is incremented when Restart TLV having RR bit unset in the received IIH PDU after the Graceful Restart was initiated by a Neighbor.
// NeighborGrSucceeded returns a uint64
func (obj *isisMetric) HasNeighborGrSucceeded() bool {
	return obj.obj.NeighborGrSucceeded != nil
}

// Number of Graceful Restarts succeeded that were initiated by a Neighbor. This counter is incremented when Restart TLV having RR bit unset in the received IIH PDU after the Graceful Restart was initiated by a Neighbor.
// SetNeighborGrSucceeded sets the uint64 value in the IsisMetric object
func (obj *isisMetric) SetNeighborGrSucceeded(value uint64) IsisMetric {

	obj.obj.NeighborGrSucceeded = &value
	return obj
}

func (obj *isisMetric) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *isisMetric) setDefault() {

}
