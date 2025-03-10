package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** RsvpLspIpv4Rro *****
type rsvpLspIpv4Rro struct {
	validation
	obj          *otg.RsvpLspIpv4Rro
	marshaller   marshalRsvpLspIpv4Rro
	unMarshaller unMarshalRsvpLspIpv4Rro
}

func NewRsvpLspIpv4Rro() RsvpLspIpv4Rro {
	obj := rsvpLspIpv4Rro{obj: &otg.RsvpLspIpv4Rro{}}
	obj.setDefault()
	return &obj
}

func (obj *rsvpLspIpv4Rro) msg() *otg.RsvpLspIpv4Rro {
	return obj.obj
}

func (obj *rsvpLspIpv4Rro) setMsg(msg *otg.RsvpLspIpv4Rro) RsvpLspIpv4Rro {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalrsvpLspIpv4Rro struct {
	obj *rsvpLspIpv4Rro
}

type marshalRsvpLspIpv4Rro interface {
	// ToProto marshals RsvpLspIpv4Rro to protobuf object *otg.RsvpLspIpv4Rro
	ToProto() (*otg.RsvpLspIpv4Rro, error)
	// ToPbText marshals RsvpLspIpv4Rro to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals RsvpLspIpv4Rro to YAML text
	ToYaml() (string, error)
	// ToJson marshals RsvpLspIpv4Rro to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals RsvpLspIpv4Rro to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalrsvpLspIpv4Rro struct {
	obj *rsvpLspIpv4Rro
}

type unMarshalRsvpLspIpv4Rro interface {
	// FromProto unmarshals RsvpLspIpv4Rro from protobuf object *otg.RsvpLspIpv4Rro
	FromProto(msg *otg.RsvpLspIpv4Rro) (RsvpLspIpv4Rro, error)
	// FromPbText unmarshals RsvpLspIpv4Rro from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals RsvpLspIpv4Rro from YAML text
	FromYaml(value string) error
	// FromJson unmarshals RsvpLspIpv4Rro from JSON text
	FromJson(value string) error
}

func (obj *rsvpLspIpv4Rro) Marshal() marshalRsvpLspIpv4Rro {
	if obj.marshaller == nil {
		obj.marshaller = &marshalrsvpLspIpv4Rro{obj: obj}
	}
	return obj.marshaller
}

func (obj *rsvpLspIpv4Rro) Unmarshal() unMarshalRsvpLspIpv4Rro {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalrsvpLspIpv4Rro{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalrsvpLspIpv4Rro) ToProto() (*otg.RsvpLspIpv4Rro, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalrsvpLspIpv4Rro) FromProto(msg *otg.RsvpLspIpv4Rro) (RsvpLspIpv4Rro, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalrsvpLspIpv4Rro) ToPbText() (string, error) {
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

func (m *unMarshalrsvpLspIpv4Rro) FromPbText(value string) error {
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

func (m *marshalrsvpLspIpv4Rro) ToYaml() (string, error) {
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

func (m *unMarshalrsvpLspIpv4Rro) FromYaml(value string) error {
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

func (m *marshalrsvpLspIpv4Rro) ToJsonRaw() (string, error) {
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

func (m *marshalrsvpLspIpv4Rro) ToJson() (string, error) {
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

func (m *unMarshalrsvpLspIpv4Rro) FromJson(value string) error {
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

func (obj *rsvpLspIpv4Rro) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *rsvpLspIpv4Rro) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *rsvpLspIpv4Rro) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *rsvpLspIpv4Rro) Clone() (RsvpLspIpv4Rro, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRsvpLspIpv4Rro()
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

// RsvpLspIpv4Rro is this contains the list of Record Route Object(RRO) objects associated with the traffic engineering tunnel. The  Record Route Object(RRO) is used in RSVP-TE to record the route traversed by the LSP. The RRO might be present in both   Path message and Resv message, the RRO stores the IP addresses of the routers that the traffic engineering tunnel  traversed and also the label generated and distributed by the routers. The RROs in the Resv message mirrors that of  the Path message, the only difference is that the RRO in a Resv message records the path information in the reverse  direction.
type RsvpLspIpv4Rro interface {
	Validation
	// msg marshals RsvpLspIpv4Rro to protobuf object *otg.RsvpLspIpv4Rro
	// and doesn't set defaults
	msg() *otg.RsvpLspIpv4Rro
	// setMsg unmarshals RsvpLspIpv4Rro from protobuf object *otg.RsvpLspIpv4Rro
	// and doesn't set defaults
	setMsg(*otg.RsvpLspIpv4Rro) RsvpLspIpv4Rro
	// provides marshal interface
	Marshal() marshalRsvpLspIpv4Rro
	// provides unmarshal interface
	Unmarshal() unMarshalRsvpLspIpv4Rro
	// validate validates RsvpLspIpv4Rro
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (RsvpLspIpv4Rro, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Address returns string, set in RsvpLspIpv4Rro.
	Address() string
	// SetAddress assigns string provided by user to RsvpLspIpv4Rro
	SetAddress(value string) RsvpLspIpv4Rro
	// HasAddress checks if Address has been set in RsvpLspIpv4Rro
	HasAddress() bool
	// ReportedLabel returns uint32, set in RsvpLspIpv4Rro.
	ReportedLabel() uint32
	// SetReportedLabel assigns uint32 provided by user to RsvpLspIpv4Rro
	SetReportedLabel(value uint32) RsvpLspIpv4Rro
	// HasReportedLabel checks if ReportedLabel has been set in RsvpLspIpv4Rro
	HasReportedLabel() bool
}

// The IPv4 addresses of the routers that the traffic engineering tunnel traversed.
// Address returns a string
func (obj *rsvpLspIpv4Rro) Address() string {

	return *obj.obj.Address

}

// The IPv4 addresses of the routers that the traffic engineering tunnel traversed.
// Address returns a string
func (obj *rsvpLspIpv4Rro) HasAddress() bool {
	return obj.obj.Address != nil
}

// The IPv4 addresses of the routers that the traffic engineering tunnel traversed.
// SetAddress sets the string value in the RsvpLspIpv4Rro object
func (obj *rsvpLspIpv4Rro) SetAddress(value string) RsvpLspIpv4Rro {

	obj.obj.Address = &value
	return obj
}

// Label reported for RRO hop. When the Label_Recording flag is set in the Session Attribute object, nodes doing route recording should include the Label Record subobject containing the reported label.
// ReportedLabel returns a uint32
func (obj *rsvpLspIpv4Rro) ReportedLabel() uint32 {

	return *obj.obj.ReportedLabel

}

// Label reported for RRO hop. When the Label_Recording flag is set in the Session Attribute object, nodes doing route recording should include the Label Record subobject containing the reported label.
// ReportedLabel returns a uint32
func (obj *rsvpLspIpv4Rro) HasReportedLabel() bool {
	return obj.obj.ReportedLabel != nil
}

// Label reported for RRO hop. When the Label_Recording flag is set in the Session Attribute object, nodes doing route recording should include the Label Record subobject containing the reported label.
// SetReportedLabel sets the uint32 value in the RsvpLspIpv4Rro object
func (obj *rsvpLspIpv4Rro) SetReportedLabel(value uint32) RsvpLspIpv4Rro {

	obj.obj.ReportedLabel = &value
	return obj
}

func (obj *rsvpLspIpv4Rro) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Address != nil {

		err := obj.validateIpv4(obj.Address())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on RsvpLspIpv4Rro.Address"))
		}

	}

}

func (obj *rsvpLspIpv4Rro) setDefault() {

}
