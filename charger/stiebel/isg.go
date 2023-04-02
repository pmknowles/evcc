package stiebel

import (
	"math"

	"github.com/volkszaehler/mbmd/encoding"
)

type Type int

const (
	_ Type = iota
	Int16
	Uint16
)

type Register struct {
	Addr                uint16
	Name, Comment, Unit string
	Typ                 Type
	Divider             float64
}

func (reg Register) Float(b []byte) float64 {
	var i int64

	switch reg.Typ {
	case Int16:
		i = int64(encoding.Int16(b))
		if i == math.MinInt16 {
			return math.NaN()
		}
	case Uint16:
		i = int64(encoding.Uint16(b))
	default:
		panic("invalid register type")
	}

	f := float64(i)
	if reg.Divider != 0 {
		f = f / reg.Divider
	}

	return f
}

var IsgInput = []Register{
	{501, "ISTTEMPERATUR FE7", "", "°C", Int16, 10},
	{502, "SOLLTEMPERATUR FE7", "", "°C", Int16, 10},
	{503, "ISTTEMPERATUR FEK", "", "°C", Int16, 10},
	{504, "SOLLTEMPERATUR FEK", "", "°C", Int16, 10},
	{505, "RAUMFEUCHTE", "", "%", Int16, 10},
	{506, "TAUPUNKTTEMPERATUR", "", "°C", Int16, 10},
	{507, "AUSSENTEMPERATUR", "", "°C", Int16, 100},
	{508, "ISTTEMPERATUR HK 1", "", "°C", Int16, 10},
	{509, "SOLLTEMPERATUR HK 1", "", "°C", Int16, 10},
	{510, "SOLLTEMPERATUR HK 1", "", "°C", Int16, 10},
	{511, "ISTTEMPERATUR HK 2", "", "°C", Int16, 10},
	{512, "SOLLTEMPERATUR HK 2", "", "°C", Int16, 10},
	{513, "VORLAUFISTTEMPERATUR WP", "MFG", "°C", Int16, 10},
	{514, "VORLAUFISTTEMPERATUR NHZ", "MFG", "°C", Int16, 10},
	{515, "VORLAUFISTTEMPERATUR", "", "°C", Int16, 10},
	{516, "RUECKLAUISTTEMPERATUR", "", "°C", Int16, -1e3},
	{517, "FESTWERTSOLLTEMPERATUR", "", "°C", Int16, 10},
	{518, "PUFFERISTTEMPERATUR", "", "°C", Int16, 10},
	{519, "PUFFERSOLLTEMPERATUR", "", "°C", Int16, 10},
	{521, "VOLUMENSTROM", "MFG", "l/min", Int16, 10},
	{522, "ISTTEMPERATUR", "Warmwasser", "°C", Int16, 10},
	{523, "SOLLTEMPERATUR", "Warmwasser", "°C", Int16, 10},
	{524, "ISTTEMPERATUR GEBLAESE", "Kühlen", "K", Int16, 10},
	{525, "SOLLTEMPERATUR GEBLAESE", "Kühlen", "K", Int16, 10},
	{526, "ISTTEMPERATUR FLAECHE", "Kühlen", "K", Int16, 10},
	{527, "SOLLTEMPERATUR FLAECHE", "Kühlen", "K", Int16, 10},
	{528, "KOLLEKTORTEMPERATUR", "Solar", "°C", Int16, 10},
	{529, "SPEICHERTEMPERATUR", "Solar", "°C", Int16, 10},
	{531, "ISTTEMPERATUR", "Wärmeerzeuger extern", "°C", Int16, 10},
	{532, "SOLLTEMPERATUR", "Wärmeerzeuger extern", "K", Int16, -1e3},
	{533, "EINSATZGRENZE HZG", "", "°C", Int16, -1e3},
	{534, "EINSATZGRENZE WW", "", "°C", Int16, 10},
	{536, "QUELLENTEMPERATUR", "", "°C", Int16, -10},
	{537, "QUELLENTEMPERATUR MIN", "", "°C", Int16, 10},
	{539, "HEISSGASTEMPERATUR", "", "°C", Int16, 10},
	{540, "DRUCK HOCHDRUCK", "", "bar", Int16, 10},
	{541, "DRUCK NIEDERDRUCK", "", "bar", Int16, 10},
	{542, "RUECKLAUFTEMPERATUR", "Wärmepumpe 1", "°C", Int16, 10},
	{543, "VORLAUFTEMPERATUR", "Wärmepumpe 1", "°C", Int16, 10},
	{544, "HEISSGASTEMPERATUR", "Wärmepumpe 1", "°C", Int16, 10},
	{548, "WP WASSERVOLUMENSTROM", "Wärmepumpe 1", "l/min", Int16, 10},
	{549, "RUECKLAUFTEMPERATUR", "Wärmepumpe 2", "°C", Int16, 10},
	{550, "VORLAUFTEMPERATUR", "Wärmepumpe 2", "°C", Int16, 10},
	{551, "HEISSGASTEMPERATUR", "Wärmepumpe 2", "°C", Int16, 10},
	{555, "WP WASSERVOLUMENSTROM", "Wärmepumpe 2", "l/min", Int16, 10},
	{556, "RUECKLAUFTEMPERATUR", "Wärmepumpe 3", "°C", Int16, 10},
	{557, "VORLAUFTEMPERATUR", "Wärmepumpe 3", "°C", Int16, 10},
	{558, "HEISSGASTEMPERATUR", "Wärmepumpe 3", "°C", Int16, 10},
	{562, "WP WASSERVOLUMENSTROM", "Wärmepumpe 3", "l/min", Int16, 10},
	{563, "RUECKLAUFTEMPERATUR", "Wärmepumpe 4", "°C", Int16, 10},
	{564, "VORLAUFTEMPERATUR", "Wärmepumpe 4", "°C", Int16, 10},
	{565, "HEISSGASTEMPERATUR", "Wärmepumpe 4", "°C", Int16, 10},
	{569, "WP WASSERVOLUMENSTROM", "Wärmepumpe 4", "l/min", Int16, 10},
	{570, "RUECKLAUFTEMPERATUR", "Wärmepumpe 5", "°C", Int16, 10},
	{571, "VORLAUFTEMPERATUR", "Wärmepumpe 5", "°C", Int16, 10},
	{572, "HEISSGASTEMPERATUR", "Wärmepumpe 5", "°C", Int16, 10},
	{576, "WP WASSERVOLUMENSTROM", "Wärmepumpe 5", "l/min", Int16, 10},
	{577, "RUECKLAUFTEMPERATUR", "Wärmepumpe 6", "°C", Int16, 10},
	{578, "VORLAUFTEMPERATUR", "Wärmepumpe 6", "°C", Int16, 10},
	{579, "HEISSGASTEMPERATUR", "Wärmepumpe 6", "°C", Int16, 10},
	{583, "WP WASSERVOLUMENSTROM", "Wärmepumpe 6", "l/min", Int16, 10},
	{584, "ISTTEMPERATUR", "Raumtemperatur Heizkreis 1", "°C", Int16, 10},
	{585, "SOLLTEMPERATUR", "Raumtemperatur Heizkreis 1", "°C", Int16, 10},
	{586, "RAUMFEUCHTE", "Heizkreis 1", "%", Int16, 10},
	{587, "TAUPUNKTTEMPERATUR", "Heizkreis 1", "°C", Int16, 10},
	{588, "ISTTEMPERATUR", "Raumtemperatur Heizkreis 2", "°C", Int16, 10},
	{589, "SOLLTEMPERATUR", "Raumtemperatur Heizkreis 2", "°C", Int16, 10},
	{590, "RAUMFEUCHTE", "Heizkreis 2", "%", Int16, 10},
	{591, "TAUPUNKTTEMPERATUR", "Heizkreis 2", "°C", Int16, 10},
	{592, "ISTTEMPERATUR", "Raumtemperatur Heizkreis 3", "°C", Int16, 10},
	{593, "SOLLTEMPERATUR", "Raumtemperatur Heizkreis 3", "°C", Int16, 10},
	{594, "RAUMFEUCHTE", "Heizkreis 3", "%", Int16, 10},
	{595, "TAUPUNKTTEMPERATUR", "Heizkreis 3", "°C", Int16, 10},
	{596, "ISTTEMPERATUR", "Raumtemperatur Heizkreis 4", "°C", Int16, 10},
	{597, "SOLLTEMPERATUR", "Raumtemperatur Heizkreis 4", "°C", Int16, 10},
	{598, "RAUMFEUCHTE", "Heizkreis 4", "%", Int16, 10},
	{599, "TAUPUNKTTEMPERATUR", "Heizkreis 4", "°C", Int16, 10},
	{600, "ISTTEMPERATUR", "Raumtemperatur Heizkreis 5", "°C", Int16, 10},
	{601, "SOLLTEMPERATUR", "Raumtemperatur Heizkreis 5", "°C", Int16, 10},
	{602, "RAUMFEUCHTE", "Heizkreis 5", "%", Int16, 10},
	{603, "TAUPUNKTTEMPERATUR", "Heizkreis 5", "°C", Int16, 10},
	{604, "SOLLTEMPERATUR", "Raumtemperatur Kühlkreis 1", "°C", Int16, 10},
	{605, "SOLLTEMPERATUR", "Raumtemperatur Kühlkreis 2", "°C", Int16, 10},
	{606, "SOLLTEMPERATUR", "Raumtemperatur Kühlkreis 3", "°C", Int16, 10},
	{607, "SOLLTEMPERATUR", "Raumtemperatur Kühlkreis 4", "°C", Int16, 10},
	{608, "SOLLTEMPERATUR", "Raumtemperatur Kühlkreis 5", "°C", Int16, 10},
}

var IsgHolding = []Register{
	{1501, "BETRIEBSART", "", "", Uint16, 0},
	{1502, "KOMFORT TEMPERATUR", "Heizkreis 1", "°C", Int16, 10},
	{1503, "ECO TEMPERATUR", "Heizkreis 1", "°C", Int16, 10},
	{1505, "KOMFORT TEMPERATUR", "Heizkreis 2", "°C", Int16, 10},
	{1506, "ECO TEMPERATUR", "Heizkreis 2", "°C", Int16, 10},
	{1508, "FESTWERTBETRIEB", "", "°C", Int16, -10},
	{1509, "BIVALENZTEMPERATUR HZG", "", "°C", Int16, 10},
	{1510, "KOMFORT TEMPERATUR", "Warmwasser", "°C", Int16, 10},
	{1511, "ECO TEMPERATUR", "Warmwasser", "°C", Int16, 10},
	{1513, "BIVALENZTEMPERATUR WW", "", "°C", Int16, 10},
	{1514, "VORLAUFSOLLTEMPERATUR", "", "°C", Int16, 10},
	{1515, "HYSTERESE VORLAUFTEMP", "", "K", Int16, 10},
	{1516, "RAUMSOLLTEMPERATUR", "", "°C", Int16, 10},
	{1517, "VORLAUFSOLLTEMPERATUR", "", "°C", Int16, 10},
	{1518, "HYSTERESE VORLAUFTEMP", "", "K", Int16, 10},
	{1519, "RAUMSOLLTEMPERATUR", "", "°C", Int16, 10},
	{1520, "RESET", "", "", Uint16, 0},
}