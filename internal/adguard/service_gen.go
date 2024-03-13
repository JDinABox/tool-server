package adguard

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// MarshalMsg implements msgp.Marshaler
func (z *AdguardService) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "ID"
	o = append(o, 0x84, 0xa2, 0x49, 0x44)
	o = msgp.AppendString(o, z.ID)
	// string "Name"
	o = append(o, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "Rules"
	o = append(o, 0xa5, 0x52, 0x75, 0x6c, 0x65, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Rules)))
	for za0001 := range z.Rules {
		o = msgp.AppendString(o, z.Rules[za0001])
	}
	// string "IconSvg"
	o = append(o, 0xa7, 0x49, 0x63, 0x6f, 0x6e, 0x53, 0x76, 0x67)
	o = msgp.AppendString(o, z.IconSvg)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *AdguardService) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "ID":
			z.ID, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "ID")
				return
			}
		case "Name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Name")
				return
			}
		case "Rules":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Rules")
				return
			}
			if cap(z.Rules) >= int(zb0002) {
				z.Rules = (z.Rules)[:zb0002]
			} else {
				z.Rules = make([]string, zb0002)
			}
			for za0001 := range z.Rules {
				z.Rules[za0001], bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Rules", za0001)
					return
				}
			}
		case "IconSvg":
			z.IconSvg, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "IconSvg")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *AdguardService) Msgsize() (s int) {
	s = 1 + 3 + msgp.StringPrefixSize + len(z.ID) + 5 + msgp.StringPrefixSize + len(z.Name) + 6 + msgp.ArrayHeaderSize
	for za0001 := range z.Rules {
		s += msgp.StringPrefixSize + len(z.Rules[za0001])
	}
	s += 8 + msgp.StringPrefixSize + len(z.IconSvg)
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *AdguardServices) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "BlockedServices"
	o = append(o, 0x82, 0xaf, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.BlockedServices)))
	for za0001 := range z.BlockedServices {
		o, err = z.BlockedServices[za0001].MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "BlockedServices", za0001)
			return
		}
	}
	// string "MappedServices"
	o = append(o, 0xae, 0x4d, 0x61, 0x70, 0x70, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73)
	o = msgp.AppendMapHeader(o, uint32(len(z.MappedServices)))
	for za0002, za0003 := range z.MappedServices {
		o = msgp.AppendString(o, za0002)
		o, err = za0003.MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "MappedServices", za0002)
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *AdguardServices) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "BlockedServices":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "BlockedServices")
				return
			}
			if cap(z.BlockedServices) >= int(zb0002) {
				z.BlockedServices = (z.BlockedServices)[:zb0002]
			} else {
				z.BlockedServices = make([]AdguardService, zb0002)
			}
			for za0001 := range z.BlockedServices {
				bts, err = z.BlockedServices[za0001].UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "BlockedServices", za0001)
					return
				}
			}
		case "MappedServices":
			var zb0003 uint32
			zb0003, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "MappedServices")
				return
			}
			if z.MappedServices == nil {
				z.MappedServices = make(map[string]AdguardService, zb0003)
			} else if len(z.MappedServices) > 0 {
				for key := range z.MappedServices {
					delete(z.MappedServices, key)
				}
			}
			for zb0003 > 0 {
				var za0002 string
				var za0003 AdguardService
				zb0003--
				za0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "MappedServices")
					return
				}
				bts, err = za0003.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "MappedServices", za0002)
					return
				}
				z.MappedServices[za0002] = za0003
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *AdguardServices) Msgsize() (s int) {
	s = 1 + 16 + msgp.ArrayHeaderSize
	for za0001 := range z.BlockedServices {
		s += z.BlockedServices[za0001].Msgsize()
	}
	s += 15 + msgp.MapHeaderSize
	if z.MappedServices != nil {
		for za0002, za0003 := range z.MappedServices {
			_ = za0003
			s += msgp.StringPrefixSize + len(za0002) + za0003.Msgsize()
		}
	}
	return
}
