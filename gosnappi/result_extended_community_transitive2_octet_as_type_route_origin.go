package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ResultExtendedCommunityTransitive2OctetAsTypeRouteOrigin *****
type resultExtendedCommunityTransitive2OctetAsTypeRouteOrigin struct {
	validation
	obj          *otg.ResultExtendedCommunityTransitive2OctetAsTypeRouteOrigin
	marshaller   marshalResultExtendedCommunityTransitive2OctetAsTypeRouteOrigin
	unMarshaller unMarshalResultExtendedCommunityTransitive2OctetAsTypeRouteOrigin
}

func NewResultExtendedCommunityTransitive2OctetAsTypeRouteOrigin() ResultExtendedCommunityTransitive2OctetAsTypeRouteOrigin {
	obj := resultExtendedCommunityTransitive2OctetAsTypeRouteOrigin{obj: &otg.ResultExtendedCommunityTransitive2OctetAsTypeRouteOrigin{}}
	obj.setDefault()
	return &obj
}

func (obj *resultExtendedCommunityTransitive2OctetAsTypeRouteOrigin) msg() *otg.ResultExtendedCommunityTransitive2OctetAsTypeRouteOrigin {
	return obj.obj
}

func (obj *resultExtendedCommunityTransitive2OctetAsTypeRouteOrigin) setMsg(msg *otg.ResultExtendedCommunityTransitive2OctetAsTypeRouteOrigin) ResultExtendedCommunityTransitive2OctetAsTypeRouteOrigin {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalresultExtendedCommunityTransitive2OctetAsTypeRouteOrigin struct {
	obj *resultExtendedCommunityTransitive2OctetAsTypeRouteOrigin
}

type marshalResultExtendedCommunityTransitive2OctetAsTypeRouteOrigin interface {
	// ToProto marshals ResultExtendedCommunityTransitive2OctetAsTypeRouteOrigin to protobuf object *otg.ResultExtendedCommunityTransitive2OctetAsTypeRouteOrigin
	ToProto() (*otg.ResultExtendedCommunityTransitive2OctetAsTypeRouteOrigin, error)
	// ToPbText marshals ResultExtendedCommunityTransitive2OctetAsTypeRouteOrigin to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ResultExtendedCommunityTransitive2OctetAsTypeRouteOrigin to YAML text
	ToYaml() (string, error)
	// ToJson marshals ResultExtendedCommunityTransitive2OctetAsTypeRouteOrigin to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals ResultExtendedCommunityTransitive2OctetAsTypeRouteOrigin to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalresultExtendedCommunityTransitive2OctetAsTypeRouteOrigin struct {
	obj *resultExtendedCommunityTransitive2OctetAsTypeRouteOrigin
}

type unMarshalResultExtendedCommunityTransitive2OctetAsTypeRouteOrigin interface {
	// FromProto unmarshals ResultExtendedCommunityTransitive2OctetAsTypeRouteOrigin from protobuf object *otg.ResultExtendedCommunityTransitive2OctetAsTypeRouteOrigin
	FromProto(msg *otg.ResultExtendedCommunityTransitive2OctetAsTypeRouteOrigin) (ResultExtendedCommunityTransitive2OctetAsTypeRouteOrigin, error)
	// FromPbText unmarshals ResultExtendedCommunityTransitive2OctetAsTypeRouteOrigin from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ResultExtendedCommunityTransitive2OctetAsTypeRouteOrigin from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ResultExtendedCommunityTransitive2OctetAsTypeRouteOrigin from JSON text
	FromJson(value string) error
}

func (obj *resultExtendedCommunityTransitive2OctetAsTypeRouteOrigin) Marshal() marshalResultExtendedCommunityTransitive2OctetAsTypeRouteOrigin {
	if obj.marshaller == nil {
		obj.marshaller = &marshalresultExtendedCommunityTransitive2OctetAsTypeRouteOrigin{obj: obj}
	}
	return obj.marshaller
}

func (obj *resultExtendedCommunityTransitive2OctetAsTypeRouteOrigin) Unmarshal() unMarshalResultExtendedCommunityTransitive2OctetAsTypeRouteOrigin {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalresultExtendedCommunityTransitive2OctetAsTypeRouteOrigin{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalresultExtendedCommunityTransitive2OctetAsTypeRouteOrigin) ToProto() (*otg.ResultExtendedCommunityTransitive2OctetAsTypeRouteOrigin, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalresultExtendedCommunityTransitive2OctetAsTypeRouteOrigin) FromProto(msg *otg.ResultExtendedCommunityTransitive2OctetAsTypeRouteOrigin) (ResultExtendedCommunityTransitive2OctetAsTypeRouteOrigin, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalresultExtendedCommunityTransitive2OctetAsTypeRouteOrigin) ToPbText() (string, error) {
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

func (m *unMarshalresultExtendedCommunityTransitive2OctetAsTypeRouteOrigin) FromPbText(value string) error {
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

func (m *marshalresultExtendedCommunityTransitive2OctetAsTypeRouteOrigin) ToYaml() (string, error) {
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

func (m *unMarshalresultExtendedCommunityTransitive2OctetAsTypeRouteOrigin) FromYaml(value string) error {
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

func (m *marshalresultExtendedCommunityTransitive2OctetAsTypeRouteOrigin) ToJsonRaw() (string, error) {
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

func (m *marshalresultExtendedCommunityTransitive2OctetAsTypeRouteOrigin) ToJson() (string, error) {
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

func (m *unMarshalresultExtendedCommunityTransitive2OctetAsTypeRouteOrigin) FromJson(value string) error {
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

func (obj *resultExtendedCommunityTransitive2OctetAsTypeRouteOrigin) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *resultExtendedCommunityTransitive2OctetAsTypeRouteOrigin) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *resultExtendedCommunityTransitive2OctetAsTypeRouteOrigin) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *resultExtendedCommunityTransitive2OctetAsTypeRouteOrigin) Clone() (ResultExtendedCommunityTransitive2OctetAsTypeRouteOrigin, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewResultExtendedCommunityTransitive2OctetAsTypeRouteOrigin()
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

// ResultExtendedCommunityTransitive2OctetAsTypeRouteOrigin is the Route Origin Community identifies one or more routers that inject a set of routes (that carry this Community) into BGP. It is sent with sub-type as 0x03 .
type ResultExtendedCommunityTransitive2OctetAsTypeRouteOrigin interface {
	Validation
	// msg marshals ResultExtendedCommunityTransitive2OctetAsTypeRouteOrigin to protobuf object *otg.ResultExtendedCommunityTransitive2OctetAsTypeRouteOrigin
	// and doesn't set defaults
	msg() *otg.ResultExtendedCommunityTransitive2OctetAsTypeRouteOrigin
	// setMsg unmarshals ResultExtendedCommunityTransitive2OctetAsTypeRouteOrigin from protobuf object *otg.ResultExtendedCommunityTransitive2OctetAsTypeRouteOrigin
	// and doesn't set defaults
	setMsg(*otg.ResultExtendedCommunityTransitive2OctetAsTypeRouteOrigin) ResultExtendedCommunityTransitive2OctetAsTypeRouteOrigin
	// provides marshal interface
	Marshal() marshalResultExtendedCommunityTransitive2OctetAsTypeRouteOrigin
	// provides unmarshal interface
	Unmarshal() unMarshalResultExtendedCommunityTransitive2OctetAsTypeRouteOrigin
	// validate validates ResultExtendedCommunityTransitive2OctetAsTypeRouteOrigin
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ResultExtendedCommunityTransitive2OctetAsTypeRouteOrigin, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Global2ByteAs returns uint32, set in ResultExtendedCommunityTransitive2OctetAsTypeRouteOrigin.
	Global2ByteAs() uint32
	// SetGlobal2ByteAs assigns uint32 provided by user to ResultExtendedCommunityTransitive2OctetAsTypeRouteOrigin
	SetGlobal2ByteAs(value uint32) ResultExtendedCommunityTransitive2OctetAsTypeRouteOrigin
	// HasGlobal2ByteAs checks if Global2ByteAs has been set in ResultExtendedCommunityTransitive2OctetAsTypeRouteOrigin
	HasGlobal2ByteAs() bool
	// Local4ByteAdmin returns uint32, set in ResultExtendedCommunityTransitive2OctetAsTypeRouteOrigin.
	Local4ByteAdmin() uint32
	// SetLocal4ByteAdmin assigns uint32 provided by user to ResultExtendedCommunityTransitive2OctetAsTypeRouteOrigin
	SetLocal4ByteAdmin(value uint32) ResultExtendedCommunityTransitive2OctetAsTypeRouteOrigin
	// HasLocal4ByteAdmin checks if Local4ByteAdmin has been set in ResultExtendedCommunityTransitive2OctetAsTypeRouteOrigin
	HasLocal4ByteAdmin() bool
}

// The two octet IANA assigned AS value assigned to the Autonomous System.
// Global2ByteAs returns a uint32
func (obj *resultExtendedCommunityTransitive2OctetAsTypeRouteOrigin) Global2ByteAs() uint32 {

	return *obj.obj.Global_2ByteAs

}

// The two octet IANA assigned AS value assigned to the Autonomous System.
// Global2ByteAs returns a uint32
func (obj *resultExtendedCommunityTransitive2OctetAsTypeRouteOrigin) HasGlobal2ByteAs() bool {
	return obj.obj.Global_2ByteAs != nil
}

// The two octet IANA assigned AS value assigned to the Autonomous System.
// SetGlobal2ByteAs sets the uint32 value in the ResultExtendedCommunityTransitive2OctetAsTypeRouteOrigin object
func (obj *resultExtendedCommunityTransitive2OctetAsTypeRouteOrigin) SetGlobal2ByteAs(value uint32) ResultExtendedCommunityTransitive2OctetAsTypeRouteOrigin {

	obj.obj.Global_2ByteAs = &value
	return obj
}

// The Local Administrator sub-field contains a number from a numbering space that is administered by the organization to which the Autonomous System number carried in the Global Administrator sub-field has been assigned by an appropriate authority.
// Local4ByteAdmin returns a uint32
func (obj *resultExtendedCommunityTransitive2OctetAsTypeRouteOrigin) Local4ByteAdmin() uint32 {

	return *obj.obj.Local_4ByteAdmin

}

// The Local Administrator sub-field contains a number from a numbering space that is administered by the organization to which the Autonomous System number carried in the Global Administrator sub-field has been assigned by an appropriate authority.
// Local4ByteAdmin returns a uint32
func (obj *resultExtendedCommunityTransitive2OctetAsTypeRouteOrigin) HasLocal4ByteAdmin() bool {
	return obj.obj.Local_4ByteAdmin != nil
}

// The Local Administrator sub-field contains a number from a numbering space that is administered by the organization to which the Autonomous System number carried in the Global Administrator sub-field has been assigned by an appropriate authority.
// SetLocal4ByteAdmin sets the uint32 value in the ResultExtendedCommunityTransitive2OctetAsTypeRouteOrigin object
func (obj *resultExtendedCommunityTransitive2OctetAsTypeRouteOrigin) SetLocal4ByteAdmin(value uint32) ResultExtendedCommunityTransitive2OctetAsTypeRouteOrigin {

	obj.obj.Local_4ByteAdmin = &value
	return obj
}

func (obj *resultExtendedCommunityTransitive2OctetAsTypeRouteOrigin) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Global_2ByteAs != nil {

		if *obj.obj.Global_2ByteAs > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= ResultExtendedCommunityTransitive2OctetAsTypeRouteOrigin.Global_2ByteAs <= 65535 but Got %d", *obj.obj.Global_2ByteAs))
		}

	}

}

func (obj *resultExtendedCommunityTransitive2OctetAsTypeRouteOrigin) setDefault() {

}
