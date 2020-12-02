package gnb

const MODULE_NAME = "gnb"
const MODULE_GNB = 0x01

var PROTOMAP map[string]uint8 = map[string]uint8{
	"udp": 0x11,
	"tcp": 0x06,
}

var QOSMAP map[string]uint8 = map[string]uint8{
	"cs1": 0b001000,
	"cs2": 0b010000,
	"cs3": 0b011000,
	"cs4": 0b100000,
	"cs5": 0b101000,
	"cs6": 0b110000,
	"cs7": 0b111000,

	"af11": 0b001010,
	"af12": 0b001100,
	"af13": 0b001110,
	"af21": 0b010010,
	"af22": 0b010100,
	"af32": 0b010110,
	"af33": 0b011110,
	"af41": 0b100010,
	"af42": 0b100100,
	"af43": 0b100110,

	"be": 0b000000,
	"ef": 0b101110,
}
