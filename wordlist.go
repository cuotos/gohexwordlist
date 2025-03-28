package main

var wordlist = map[byte][2]string{
	0x00: {"aardvark", "adroitness"},
	0x01: {"absurd", "adviser"},
	0x02: {"accrue", "aftermath"},
	0x03: {"acme", "aggregate"},
	0x04: {"adrift", "alkali"},
	0x05: {"adult", "almighty"},
	0x06: {"afflict", "amulet"},
	0x07: {"ahead", "amusement"},
	0x08: {"aimless", "antenna"},
	0x09: {"Algol", "applicant"},
	0x0A: {"allow", "Apollo"},
	0x0B: {"alone", "armistice"},
	0x0C: {"ammo", "article"},
	0x0D: {"ancient", "asteroid"},
	0x0E: {"apple", "Atlantic"},
	0x0F: {"artist", "atmosphere"},
	0x10: {"assume", "autopsy"},
	0x11: {"Athens", "Babylon"},
	0x12: {"atlas", "backwater"},
	0x13: {"Aztec", "barbecue"},
	0x14: {"baboon", "belowground"},
	0x15: {"backfield", "bifocals"},
	0x16: {"backward", "bodyguard"},
	0x17: {"banjo", "bookseller"},
	0x18: {"beaming", "borderline"},
	0x19: {"bedlamp", "bottomless"},
	0x1A: {"beehive", "Bradbury"},
	0x1B: {"beeswax", "bravado"},
	0x1C: {"befriend", "Brazilian"},
	0x1D: {"Belfast", "breakaway"},
	0x1E: {"berserk", "Burlington"},
	0x1F: {"billiard", "businessman"},
	0x20: {"bison", "butterfat"},
	0x21: {"blackjack", "Camelot"},
	0x22: {"blockade", "candidate"},
	0x23: {"blowtorch", "cannonball"},
	0x24: {"bluebird", "Capricorn"},
	0x25: {"bombast", "caravan"},
	0x26: {"bookshelf", "caretaker"},
	0x27: {"brackish", "celebrate"},
	0x28: {"breadline", "cellulose"},
	0x29: {"breakup", "certify"},
	0x2A: {"brickyard", "chambermaid"},
	0x2B: {"briefcase", "Cherokee"},
	0x2C: {"Burbank", "Chicago"},
	0x2D: {"button", "clergyman"},
	0x2E: {"buzzard", "coherence"},
	0x2F: {"cement", "combustion"},
	0x30: {"chairlift", "commando"},
	0x31: {"chatter", "company"},
	0x32: {"checkup", "component"},
	0x33: {"chisel", "concurrent"},
	0x34: {"choking", "confidence"},
	0x35: {"chopper", "conformist"},
	0x36: {"Christmas", "congregate"},
	0x37: {"clamshell", "consensus"},
	0x38: {"classic", "consulting"},
	0x39: {"classroom", "corporate"},
	0x3A: {"cleanup", "corrosion"},
	0x3B: {"clockwork", "councilman"},
	0x3C: {"cobra", "crossover"},
	0x3D: {"commence", "crucifix"},
	0x3E: {"concert", "cumbersome"},
	0x3F: {"cowbell", "customer"},
	0x40: {"crackdown", "Dakota"},
	0x41: {"cranky", "decadence"},
	0x42: {"crowfoot", "December"},
	0x43: {"crucial", "decimal"},
	0x44: {"crumpled", "designing"},
	0x45: {"crusade", "detector"},
	0x46: {"cubic", "detergent"},
	0x47: {"dashboard", "determine"},
	0x48: {"deadbolt", "dictator"},
	0x49: {"deckhand", "dinosaur"},
	0x4A: {"dogsled", "direction"},
	0x4B: {"dragnet", "disable"},
	0x4C: {"drainage", "disbelief"},
	0x4D: {"dreadful", "disruptive"},
	0x4E: {"drifter", "distortion"},
	0x4F: {"dropper", "document"},
	0x50: {"drumbeat", "embezzle"},
	0x51: {"drunken", "enchanting"},
	0x52: {"Dupont", "enrollment"},
	0x53: {"dwelling", "enterprise"},
	0x54: {"eating", "equation"},
	0x55: {"edict", "equipment"},
	0x56: {"egghead", "escapade"},
	0x57: {"eightball", "Eskimo"},
	0x58: {"endorse", "everyday"},
	0x59: {"endow", "examine"},
	0x5A: {"enlist", "existence"},
	0x5B: {"erase", "exodus"},
	0x5C: {"escape", "fascinate"},
	0x5D: {"exceed", "filament"},
	0x5E: {"eyeglass", "finicky"},
	0x5F: {"eyetooth", "forever"},
	0x60: {"facial", "fortitude"},
	0x61: {"fallout", "frequency"},
	0x62: {"flagpole", "gadgetry"},
	0x63: {"flatfoot", "Galveston"},
	0x64: {"flytrap", "getaway"},
	0x65: {"fracture", "glossary"},
	0x66: {"framework", "gossamer"},
	0x67: {"freedom", "graduate"},
	0x68: {"frighten", "gravity"},
	0x69: {"gazelle", "guitarist"},
	0x6A: {"Geiger", "hamburger"},
	0x6B: {"glitter", "Hamilton"},
	0x6C: {"glucose", "handiwork"},
	0x6D: {"goggles", "hazardous"},
	0x6E: {"goldfish", "headwaters"},
	0x6F: {"gremlin", "hemisphere"},
	0x70: {"guidance", "hesitate"},
	0x71: {"hamlet", "hideaway"},
	0x72: {"highchair", "holiness"},
	0x73: {"hockey", "hurricane"},
	0x74: {"indoors", "hydraulic"},
	0x75: {"indulge", "impartial"},
	0x76: {"inverse", "impetus"},
	0x77: {"involve", "inception"},
	0x78: {"island", "indigo"},
	0x79: {"jawbone", "inertia"},
	0x7A: {"keyboard", "infancy"},
	0x7B: {"kickoff", "inferno"},
	0x7C: {"kiwi", "informant"},
	0x7D: {"klaxon", "insincere"},
	0x7E: {"locale", "insurgent"},
	0x7F: {"lockup", "integrate"},
	0x80: {"merit", "intention"},
	0x81: {"minnow", "inventive"},
	0x82: {"miser", "Istanbul"},
	0x83: {"Mohawk", "Jamaica"},
	0x84: {"mural", "Jupiter"},
	0x85: {"music", "leprosy"},
	0x86: {"necklace", "letterhead"},
	0x87: {"Neptune", "liberty"},
	0x88: {"newborn", "maritime"},
	0x89: {"nightbird", "matchmaker"},
	0x8A: {"Oakland", "maverick"},
	0x8B: {"obtuse", "Medusa"},
	0x8C: {"offload", "megaton"},
	0x8D: {"optic", "microscope"},
	0x8E: {"orca", "microwave"},
	0x8F: {"payday", "midsummer"},
	0x90: {"peachy", "millionaire"},
	0x91: {"pheasant", "miracle"},
	0x92: {"physique", "misnomer"},
	0x93: {"playhouse", "molasses"},
	0x94: {"Pluto", "molecule"},
	0x95: {"preclude", "Montana"},
	0x96: {"prefer", "monument"},
	0x97: {"preshrunk", "mosquito"},
	0x98: {"printer", "narrative"},
	0x99: {"prowler", "nebula"},
	0x9A: {"pupil", "newsletter"},
	0x9B: {"puppy", "Norwegian"},
	0x9C: {"python", "October"},
	0x9D: {"quadrant", "Ohio"},
	0x9E: {"quiver", "onlooker"},
	0x9F: {"quota", "opulent"},
	0xA0: {"ragtime", "Orlando"},
	0xA1: {"ratchet", "outfielder"},
	0xA2: {"rebirth", "Pacific"},
	0xA3: {"reform", "pandemic"},
	0xA4: {"regain", "Pandora"},
	0xA5: {"reindeer", "paperweight"},
	0xA6: {"rematch", "paragon"},
	0xA7: {"repay", "paragraph"},
	0xA8: {"retouch", "paramount"},
	0xA9: {"revenge", "passenger"},
	0xAA: {"reward", "pedigree"},
	0xAB: {"rhythm", "Pegasus"},
	0xAC: {"ribcage", "penetrate"},
	0xAD: {"ringbolt", "perceptive"},
	0xAE: {"robust", "performance"},
	0xAF: {"rocker", "pharmacy"},
	0xB0: {"ruffled", "phonetic"},
	0xB1: {"sailboat", "photograph"},
	0xB2: {"sawdust", "pioneer"},
	0xB3: {"scallion", "pocketful"},
	0xB4: {"scenic", "politeness"},
	0xB5: {"scorecard", "positive"},
	0xB6: {"Scotland", "potato"},
	0xB7: {"seabird", "processor"},
	0xB8: {"select", "provincial"},
	0xB9: {"sentence", "proximate"},
	0xBA: {"shadow", "puberty"},
	0xBB: {"shamrock", "publisher"},
	0xBC: {"showgirl", "pyramid"},
	0xBD: {"skullcap", "quantity"},
	0xBE: {"skydive", "racketeer"},
	0xBF: {"slingshot", "rebellion"},
	0xC0: {"slowdown", "recipe"},
	0xC1: {"snapline", "recover"},
	0xC2: {"snapshot", "repellent"},
	0xC3: {"snowcap", "replica"},
	0xC4: {"snowslide", "reproduce"},
	0xC5: {"solo", "resistor"},
	0xC6: {"southward", "responsive"},
	0xC7: {"soybean", "retraction"},
	0xC8: {"spaniel", "retrieval"},
	0xC9: {"spearhead", "retrospect"},
	0xCA: {"spellbind", "revenue"},
	0xCB: {"spheroid", "revival"},
	0xCC: {"spigot", "revolver"},
	0xCD: {"spindle", "sandalwood"},
	0xCE: {"spyglass", "sardonic"},
	0xCF: {"stagehand", "Saturday"},
	0xD0: {"stagnate", "savagery"},
	0xD1: {"stairway", "scavenger"},
	0xD2: {"standard", "sensation"},
	0xD3: {"stapler", "sociable"},
	0xD4: {"steamship", "souvenir"},
	0xD5: {"sterling", "specialist"},
	0xD6: {"stockman", "speculate"},
	0xD7: {"stopwatch", "stethoscope"},
	0xD8: {"stormy", "stupendous"},
	0xD9: {"sugar", "supportive"},
	0xDA: {"surmount", "surrender"},
	0xDB: {"suspense", "suspicious"},
	0xDC: {"sweatband", "sympathy"},
	0xDD: {"swelter", "tambourine"},
	0xDE: {"tactics", "telephone"},
	0xDF: {"talon", "therapist"},
	0xE0: {"tapeworm", "tobacco"},
	0xE1: {"tempest", "tolerance"},
	0xE2: {"tiger", "tomorrow"},
	0xE3: {"tissue", "torpedo"},
	0xE4: {"tonic", "tradition"},
	0xE5: {"topmost", "travesty"},
	0xE6: {"tracker", "trombonist"},
	0xE7: {"transit", "truncated"},
	0xE8: {"trauma", "typewriter"},
	0xE9: {"treadmill", "ultimate"},
	0xEA: {"Trojan", "undaunted"},
	0xEB: {"trouble", "underfoot"},
	0xEC: {"tumor", "unicorn"},
	0xED: {"tunnel", "unify"},
	0xEE: {"tycoon", "universe"},
	0xEF: {"uncut", "unravel"},
	0xF0: {"unearth", "upcoming"},
	0xF1: {"unwind", "vacancy"},
	0xF2: {"uproot", "vagabond"},
	0xF3: {"upset", "vertigo"},
	0xF4: {"upshot", "Virginia"},
	0xF5: {"vapor", "visitor"},
	0xF6: {"village", "vocalist"},
	0xF7: {"virus", "voyager"},
	0xF8: {"Vulcan", "warranty"},
	0xF9: {"waffle", "Waterloo"},
	0xFA: {"wallet", "whimsical"},
	0xFB: {"watchword", "Wichita"},
	0xFC: {"wayside", "Wilmington"},
	0xFD: {"willow", "Wyoming"},
	0xFE: {"woodlark", "yesteryear"},
	0xFF: {"Zulu", "Yucatan"},
}
