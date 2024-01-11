package ping

import "fmt"

const (
	Anime rune = 'A'
	STD        = "STD"
	GNR        = "GNR"
	LNG        = "LNG"
	ACT        = "ACT"
)

type SegmentKey int64

func (x SegmentKey) LNG() KeyGenrator {
	return PingKey(fmt.Sprintf("LNG:%d", x))
}

func (x SegmentKey) STD() KeyGenrator {
	return PingKey(fmt.Sprintf("STD:%d", x))
}

func (x SegmentKey) GNR() KeyGenrator {
	return PingKey(fmt.Sprintf("GNR:%d", x))
}

func (x SegmentKey) ACT() KeyGenrator {
	return PingKey(fmt.Sprintf("ACT:%d", x))
}

func (x SegmentKey) CHR(v rune) KeyGenrator {
	return PingKey(fmt.Sprintf("CHR:X-%c:%d", v, x))
}

func (x SegmentKey) TAG(v rune) KeyGenrator {
	return PingKey(fmt.Sprintf("TAG:X-%c:%d", v, x))
}

func (x SegmentKey) IMG(v rune) KeyGenrator {
	return PingKey(fmt.Sprintf("IMG:X-%c:%d", v, x))
}

func (x SegmentKey) TRL(v rune) KeyGenrator {
	return PingKey(fmt.Sprintf("TRL:X-%c:%d", v, x))
}

func CTM(v rune, w, y string) KeyGenrator {
	return PingKey(fmt.Sprintf("G-%c:%s:%s", v, w, y))
}
