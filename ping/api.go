package ping

import "fmt"

type CacheKey struct {
	ID     int64
	Target string
}

func (x *CacheKey) Main() KeyGenrator {
	return PingKey(fmt.Sprintf("%s:ID:%d", x.Target, x.ID))
}

func (x *CacheKey) Meta() KeyGenrator {
	return PingKey(fmt.Sprintf("%s:MTD:%d", x.Target, x.ID))
}

func (x *CacheKey) Resources() KeyGenrator {
	return PingKey(fmt.Sprintf("%s:RSC:%d", x.Target, x.ID))
}

func (x *CacheKey) Links() KeyGenrator {
	return PingKey(fmt.Sprintf("%s:LNK:%d", x.Target, x.ID))
}

func (x *CacheKey) Server() KeyGenrator {
	return PingKey(fmt.Sprintf("%s:SRV:%d", x.Target, x.ID))
}

func (x *CacheKey) SubV() KeyGenrator {
	return PingKey(fmt.Sprintf("%s:SUB:VDO:%d", x.Target, x.ID))
}

func (x *CacheKey) SubT() KeyGenrator {
	return PingKey(fmt.Sprintf("%s:SUB:TRN:%d", x.Target, x.ID))
}

func (x *CacheKey) DubV() KeyGenrator {
	return PingKey(fmt.Sprintf("%s:DUB:VDO:%d", x.Target, x.ID))
}

func (x *CacheKey) DubT() KeyGenrator {
	return PingKey(fmt.Sprintf("%s:DUB:TRN:%d", x.Target, x.ID))
}

func (x *CacheKey) Posters() KeyGenrator {
	return PingKey(fmt.Sprintf("%s:PST:%d", x.Target, x.ID))
}

func (x *CacheKey) Backdrops() KeyGenrator {
	return PingKey(fmt.Sprintf("%s:BKD:%d", x.Target, x.ID))
}

func (x *CacheKey) Logos() KeyGenrator {
	return PingKey(fmt.Sprintf("%s:LOG:%d", x.Target, x.ID))
}

func (x *CacheKey) Studio() KeyGenrator {
	return PingKey(fmt.Sprintf("%s:STD:%d", x.Target, x.ID))
}

func (x *CacheKey) Genre() KeyGenrator {
	return PingKey(fmt.Sprintf("%s:GNR:%d", x.Target, x.ID))
}

func (x *CacheKey) Tags() KeyGenrator {
	return PingKey(fmt.Sprintf("%s:TAG:%d", x.Target, x.ID))
}

func (x *CacheKey) Trailers() KeyGenrator {
	return PingKey(fmt.Sprintf("%s:TRL:%d", x.Target, x.ID))
}

func (x *CacheKey) Characters() KeyGenrator {
	return PingKey(fmt.Sprintf("%s:CHR:%d", x.Target, x.ID))
}

func (x *CacheKey) Seasons(Limit, Offset int32) KeyGenrator {
	return PingKey(fmt.Sprintf("%s:SSN:%d:%d-%d", x.Target, x.ID, Limit, Offset))
}

func (x *CacheKey) Episodes(Limit, Offset int32) KeyGenrator {
	return PingKey(fmt.Sprintf("%s:EPS:%d:%d-%d", x.Target, x.ID, Limit, Offset))
}
