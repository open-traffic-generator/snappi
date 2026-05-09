package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisSRv6Msd *****
type isisSRv6Msd struct {
	validation
	obj          *otg.IsisSRv6Msd
	marshaller   marshalIsisSRv6Msd
	unMarshaller unMarshalIsisSRv6Msd
}

func NewIsisSRv6Msd() IsisSRv6Msd {
	obj := isisSRv6Msd{obj: &otg.IsisSRv6Msd{}}
	obj.setDefault()
	return &obj
}

func (obj *isisSRv6Msd) msg() *otg.IsisSRv6Msd {
	return obj.obj
}

func (obj *isisSRv6Msd) setMsg(msg *otg.IsisSRv6Msd) IsisSRv6Msd {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisSRv6Msd struct {
	obj *isisSRv6Msd
}

type marshalIsisSRv6Msd interface {
	// ToProto marshals IsisSRv6Msd to protobuf object *otg.IsisSRv6Msd
	ToProto() (*otg.IsisSRv6Msd, error)
	// ToPbText marshals IsisSRv6Msd to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisSRv6Msd to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisSRv6Msd to JSON text
	ToJson() (string, error)
}

type unMarshalisisSRv6Msd struct {
	obj *isisSRv6Msd
}

type unMarshalIsisSRv6Msd interface {
	// FromProto unmarshals IsisSRv6Msd from protobuf object *otg.IsisSRv6Msd
	FromProto(msg *otg.IsisSRv6Msd) (IsisSRv6Msd, error)
	// FromPbText unmarshals IsisSRv6Msd from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisSRv6Msd from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisSRv6Msd from JSON text
	FromJson(value string) error
}

func (obj *isisSRv6Msd) Marshal() marshalIsisSRv6Msd {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisSRv6Msd{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisSRv6Msd) Unmarshal() unMarshalIsisSRv6Msd {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisSRv6Msd{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisSRv6Msd) ToProto() (*otg.IsisSRv6Msd, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisSRv6Msd) FromProto(msg *otg.IsisSRv6Msd) (IsisSRv6Msd, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisSRv6Msd) ToPbText() (string, error) {
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

func (m *unMarshalisisSRv6Msd) FromPbText(value string) error {
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

func (m *marshalisisSRv6Msd) ToYaml() (string, error) {
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

func (m *unMarshalisisSRv6Msd) FromYaml(value string) error {
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

func (m *marshalisisSRv6Msd) ToJson() (string, error) {
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

func (m *unMarshalisisSRv6Msd) FromJson(value string) error {
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

func (obj *isisSRv6Msd) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisSRv6Msd) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisSRv6Msd) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisSRv6Msd) Clone() (IsisSRv6Msd, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisSRv6Msd()
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

// IsisSRv6Msd is sRv6 Maximum SID Depth (MSD) sub-TLV values, used for both node-level advertisement (carried in IS-IS Router CAPABILITY TLV 242 via Node MSD sub-TLV type 23) and link-level advertisement (carried in Extended IS Reachability TLV via Link MSD sub-TLV type 15). A field value of 0 suppresses advertisement of that sub-TLV. Reference: RFC 9352 Section 6.
type IsisSRv6Msd interface {
	Validation
	// msg marshals IsisSRv6Msd to protobuf object *otg.IsisSRv6Msd
	// and doesn't set defaults
	msg() *otg.IsisSRv6Msd
	// setMsg unmarshals IsisSRv6Msd from protobuf object *otg.IsisSRv6Msd
	// and doesn't set defaults
	setMsg(*otg.IsisSRv6Msd) IsisSRv6Msd
	// provides marshal interface
	Marshal() marshalIsisSRv6Msd
	// provides unmarshal interface
	Unmarshal() unMarshalIsisSRv6Msd
	// validate validates IsisSRv6Msd
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisSRv6Msd, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// MaxSl returns uint32, set in IsisSRv6Msd.
	MaxSl() uint32
	// SetMaxSl assigns uint32 provided by user to IsisSRv6Msd
	SetMaxSl(value uint32) IsisSRv6Msd
	// HasMaxSl checks if MaxSl has been set in IsisSRv6Msd
	HasMaxSl() bool
	// MaxEndPopSrh returns uint32, set in IsisSRv6Msd.
	MaxEndPopSrh() uint32
	// SetMaxEndPopSrh assigns uint32 provided by user to IsisSRv6Msd
	SetMaxEndPopSrh(value uint32) IsisSRv6Msd
	// HasMaxEndPopSrh checks if MaxEndPopSrh has been set in IsisSRv6Msd
	HasMaxEndPopSrh() bool
	// MaxHEncaps returns uint32, set in IsisSRv6Msd.
	MaxHEncaps() uint32
	// SetMaxHEncaps assigns uint32 provided by user to IsisSRv6Msd
	SetMaxHEncaps(value uint32) IsisSRv6Msd
	// HasMaxHEncaps checks if MaxHEncaps has been set in IsisSRv6Msd
	HasMaxHEncaps() bool
	// MaxEndDSrh returns uint32, set in IsisSRv6Msd.
	MaxEndDSrh() uint32
	// SetMaxEndDSrh assigns uint32 provided by user to IsisSRv6Msd
	SetMaxEndDSrh(value uint32) IsisSRv6Msd
	// HasMaxEndDSrh checks if MaxEndDSrh has been set in IsisSRv6Msd
	HasMaxEndDSrh() bool
	// MaxTInsert returns uint32, set in IsisSRv6Msd.
	MaxTInsert() uint32
	// SetMaxTInsert assigns uint32 provided by user to IsisSRv6Msd
	SetMaxTInsert(value uint32) IsisSRv6Msd
	// HasMaxTInsert checks if MaxTInsert has been set in IsisSRv6Msd
	HasMaxTInsert() bool
	// MaxTEncaps returns uint32, set in IsisSRv6Msd.
	MaxTEncaps() uint32
	// SetMaxTEncaps assigns uint32 provided by user to IsisSRv6Msd
	SetMaxTEncaps(value uint32) IsisSRv6Msd
	// HasMaxTEncaps checks if MaxTEncaps has been set in IsisSRv6Msd
	HasMaxTEncaps() bool
}

// Maximum Segments Left (Max-SL). Maximum value of the Segments Left field in an SRH that this node can correctly process. A value of 0 suppresses advertisement of this sub-TLV. Reference: RFC 9352 Section 6.
// MaxSl returns a uint32
func (obj *isisSRv6Msd) MaxSl() uint32 {

	return *obj.obj.MaxSl

}

// Maximum Segments Left (Max-SL). Maximum value of the Segments Left field in an SRH that this node can correctly process. A value of 0 suppresses advertisement of this sub-TLV. Reference: RFC 9352 Section 6.
// MaxSl returns a uint32
func (obj *isisSRv6Msd) HasMaxSl() bool {
	return obj.obj.MaxSl != nil
}

// Maximum Segments Left (Max-SL). Maximum value of the Segments Left field in an SRH that this node can correctly process. A value of 0 suppresses advertisement of this sub-TLV. Reference: RFC 9352 Section 6.
// SetMaxSl sets the uint32 value in the IsisSRv6Msd object
func (obj *isisSRv6Msd) SetMaxSl(value uint32) IsisSRv6Msd {

	obj.obj.MaxSl = &value
	return obj
}

// Maximum End Pop SRH (Max-End-Pop). Maximum size of the SRH this node can remove when processing an End.Pop operation. A value of 0 suppresses advertisement of this sub-TLV. Reference: RFC 9352 Section 6.
// MaxEndPopSrh returns a uint32
func (obj *isisSRv6Msd) MaxEndPopSrh() uint32 {

	return *obj.obj.MaxEndPopSrh

}

// Maximum End Pop SRH (Max-End-Pop). Maximum size of the SRH this node can remove when processing an End.Pop operation. A value of 0 suppresses advertisement of this sub-TLV. Reference: RFC 9352 Section 6.
// MaxEndPopSrh returns a uint32
func (obj *isisSRv6Msd) HasMaxEndPopSrh() bool {
	return obj.obj.MaxEndPopSrh != nil
}

// Maximum End Pop SRH (Max-End-Pop). Maximum size of the SRH this node can remove when processing an End.Pop operation. A value of 0 suppresses advertisement of this sub-TLV. Reference: RFC 9352 Section 6.
// SetMaxEndPopSrh sets the uint32 value in the IsisSRv6Msd object
func (obj *isisSRv6Msd) SetMaxEndPopSrh(value uint32) IsisSRv6Msd {

	obj.obj.MaxEndPopSrh = &value
	return obj
}

// Maximum H.Encaps (Max-H.Encaps). Maximum number of SRv6 SID encapsulations this node can perform as a headend. A value of 0 suppresses advertisement of this sub-TLV. Reference: RFC 9352 Section 6.
// MaxHEncaps returns a uint32
func (obj *isisSRv6Msd) MaxHEncaps() uint32 {

	return *obj.obj.MaxHEncaps

}

// Maximum H.Encaps (Max-H.Encaps). Maximum number of SRv6 SID encapsulations this node can perform as a headend. A value of 0 suppresses advertisement of this sub-TLV. Reference: RFC 9352 Section 6.
// MaxHEncaps returns a uint32
func (obj *isisSRv6Msd) HasMaxHEncaps() bool {
	return obj.obj.MaxHEncaps != nil
}

// Maximum H.Encaps (Max-H.Encaps). Maximum number of SRv6 SID encapsulations this node can perform as a headend. A value of 0 suppresses advertisement of this sub-TLV. Reference: RFC 9352 Section 6.
// SetMaxHEncaps sets the uint32 value in the IsisSRv6Msd object
func (obj *isisSRv6Msd) SetMaxHEncaps(value uint32) IsisSRv6Msd {

	obj.obj.MaxHEncaps = &value
	return obj
}

// Maximum End.D SRH (Max-End-D). Maximum size of the SRH this node can process when performing an End.D (decapsulation) function. A value of 0 suppresses advertisement of this sub-TLV. Reference: RFC 9352 Section 6.
// MaxEndDSrh returns a uint32
func (obj *isisSRv6Msd) MaxEndDSrh() uint32 {

	return *obj.obj.MaxEndDSrh

}

// Maximum End.D SRH (Max-End-D). Maximum size of the SRH this node can process when performing an End.D (decapsulation) function. A value of 0 suppresses advertisement of this sub-TLV. Reference: RFC 9352 Section 6.
// MaxEndDSrh returns a uint32
func (obj *isisSRv6Msd) HasMaxEndDSrh() bool {
	return obj.obj.MaxEndDSrh != nil
}

// Maximum End.D SRH (Max-End-D). Maximum size of the SRH this node can process when performing an End.D (decapsulation) function. A value of 0 suppresses advertisement of this sub-TLV. Reference: RFC 9352 Section 6.
// SetMaxEndDSrh sets the uint32 value in the IsisSRv6Msd object
func (obj *isisSRv6Msd) SetMaxEndDSrh(value uint32) IsisSRv6Msd {

	obj.obj.MaxEndDSrh = &value
	return obj
}

// Maximum T.Insert (Max-T.Insert). Maximum size of the SRH this node can insert using a T.Insert operation. A value of 0 suppresses advertisement of this sub-TLV. Reference: RFC 9352 Section 6.
// MaxTInsert returns a uint32
func (obj *isisSRv6Msd) MaxTInsert() uint32 {

	return *obj.obj.MaxTInsert

}

// Maximum T.Insert (Max-T.Insert). Maximum size of the SRH this node can insert using a T.Insert operation. A value of 0 suppresses advertisement of this sub-TLV. Reference: RFC 9352 Section 6.
// MaxTInsert returns a uint32
func (obj *isisSRv6Msd) HasMaxTInsert() bool {
	return obj.obj.MaxTInsert != nil
}

// Maximum T.Insert (Max-T.Insert). Maximum size of the SRH this node can insert using a T.Insert operation. A value of 0 suppresses advertisement of this sub-TLV. Reference: RFC 9352 Section 6.
// SetMaxTInsert sets the uint32 value in the IsisSRv6Msd object
func (obj *isisSRv6Msd) SetMaxTInsert(value uint32) IsisSRv6Msd {

	obj.obj.MaxTInsert = &value
	return obj
}

// Maximum T.Encaps (Max-T.Encaps). Maximum size of the SRH this node can insert using a T.Encaps operation. A value of 0 suppresses advertisement of this sub-TLV. Reference: RFC 9352 Section 6.
// MaxTEncaps returns a uint32
func (obj *isisSRv6Msd) MaxTEncaps() uint32 {

	return *obj.obj.MaxTEncaps

}

// Maximum T.Encaps (Max-T.Encaps). Maximum size of the SRH this node can insert using a T.Encaps operation. A value of 0 suppresses advertisement of this sub-TLV. Reference: RFC 9352 Section 6.
// MaxTEncaps returns a uint32
func (obj *isisSRv6Msd) HasMaxTEncaps() bool {
	return obj.obj.MaxTEncaps != nil
}

// Maximum T.Encaps (Max-T.Encaps). Maximum size of the SRH this node can insert using a T.Encaps operation. A value of 0 suppresses advertisement of this sub-TLV. Reference: RFC 9352 Section 6.
// SetMaxTEncaps sets the uint32 value in the IsisSRv6Msd object
func (obj *isisSRv6Msd) SetMaxTEncaps(value uint32) IsisSRv6Msd {

	obj.obj.MaxTEncaps = &value
	return obj
}

func (obj *isisSRv6Msd) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.MaxSl != nil {

		if *obj.obj.MaxSl > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= IsisSRv6Msd.MaxSl <= 255 but Got %d", *obj.obj.MaxSl))
		}

	}

	if obj.obj.MaxEndPopSrh != nil {

		if *obj.obj.MaxEndPopSrh > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= IsisSRv6Msd.MaxEndPopSrh <= 255 but Got %d", *obj.obj.MaxEndPopSrh))
		}

	}

	if obj.obj.MaxHEncaps != nil {

		if *obj.obj.MaxHEncaps > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= IsisSRv6Msd.MaxHEncaps <= 255 but Got %d", *obj.obj.MaxHEncaps))
		}

	}

	if obj.obj.MaxEndDSrh != nil {

		if *obj.obj.MaxEndDSrh > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= IsisSRv6Msd.MaxEndDSrh <= 255 but Got %d", *obj.obj.MaxEndDSrh))
		}

	}

	if obj.obj.MaxTInsert != nil {

		if *obj.obj.MaxTInsert > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= IsisSRv6Msd.MaxTInsert <= 255 but Got %d", *obj.obj.MaxTInsert))
		}

	}

	if obj.obj.MaxTEncaps != nil {

		if *obj.obj.MaxTEncaps > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= IsisSRv6Msd.MaxTEncaps <= 255 but Got %d", *obj.obj.MaxTEncaps))
		}

	}

}

func (obj *isisSRv6Msd) setDefault() {
	if obj.obj.MaxSl == nil {
		obj.SetMaxSl(0)
	}
	if obj.obj.MaxEndPopSrh == nil {
		obj.SetMaxEndPopSrh(0)
	}
	if obj.obj.MaxHEncaps == nil {
		obj.SetMaxHEncaps(0)
	}
	if obj.obj.MaxEndDSrh == nil {
		obj.SetMaxEndDSrh(0)
	}
	if obj.obj.MaxTInsert == nil {
		obj.SetMaxTInsert(0)
	}
	if obj.obj.MaxTEncaps == nil {
		obj.SetMaxTEncaps(0)
	}

}
