package main

import (
	"fmt"
	"log"
	"strings"

	normalize "github.com/pajlada/lidl-normalize"
)

var tests = map[string]string{
	"a":           "a",
	"Æ":           "AE",
	"æ":           "ae",
	"Å":           "A",
	"å":           "a",
	"Ǻ":           "A",
	"ǻ":           "a",
	"Ḁ":           "A",
	"ḁ":           "a",
	"ẚ":           "a",
	"Ă":           "A",
	"ă":           "a",
	"Ặ":           "A",
	"ặ":           "a",
	"Ắ":           "A",
	"ắ":           "a",
	"Ằ":           "A",
	"ằ":           "a",
	"Ẳ":           "A",
	"ẳ":           "a",
	"Ẵ":           "A",
	"ẵ":           "a",
	"Ȃ":           "A",
	"ȃ":           "a",
	"Â":           "A",
	"â":           "a",
	"Ậ":           "A",
	"ậ":           "a",
	"Ấ":           "A",
	"ấ":           "a",
	"Ầ":           "A",
	"ầ":           "a",
	"Ẫ":           "A",
	"ẫ":           "a",
	"Ẩ":           "A",
	"ẩ":           "a",
	"Ả":           "A",
	"ả":           "a",
	"Ǎ":           "A",
	"ǎ":           "a",
	"Ⱥ":           "A",
	"ⱥ":           "a",
	"Ȧ":           "A",
	"ȧ":           "a",
	"Ǡ":           "A",
	"ǡ":           "a",
	"Ạ":           "A",
	"ạ":           "a",
	"Ä":           "A",
	"ä":           "a",
	"Ǟ":           "A",
	"ǟ":           "a",
	"À":           "A",
	"à":           "a",
	"Ȁ":           "A",
	"ȁ":           "a",
	"Á":           "A",
	"á":           "a",
	"Ā":           "A",
	"ā":           "a",
	"Ā̀":          "A",
	"ā̀":          "a",
	"Ã":           "A",
	"ã":           "a",
	"Ą":           "A",
	"ą":           "a",
	"Ą́":          "A",
	"ą́":          "a",
	"Ą̃":          "A",
	"ą̃":          "a",
	"ᶏ":           "a",
	"Ɑ":           "a",
	"ɑ":           "a",
	"ᶐ":           "a",
	"Ɐ":           "A",
	"ɐ":           "a",
	"Λ":           "A",
	"ʌ":           "A",
	"Ɒ":           "a",
	"ɒ":           "a",
	"ᶛ":           "a",
	"ᴀ":           "A",
	"ᴬ":           "A",
	"ᵃ":           "a",
	"ᵄ":           "a",
	"ₐ":           "a",
	"ꬱ":           "a",
	"⍺":           "a",
	"abc å def":   "abc a def",
	"ˢᵐᵒˡ ⁿᵃᵗᶦᵒⁿ": "smol nation", // smol nation
	// "Ниг":           "Nig",
	"🇦":     "A",
	"🇺🇦XD":  "UAXD",
	"🆓 ICE": "FREE ICE",
	// "ер":            "er",
	"chocolate 🇳🇮b": "chocolate NIb",
	"🅱lueberry":     "Blueberry",
	"⒝":             "b",
	"ü Ü ö Ö ä Ä":   "u U o O a A",

	"ᴭ": "AE",
	"ᴮ": "B",
	"ᴯ": "B",
	"ᴰ": "D",
	"ᴱ": "E",
	"ᴲ": "E",
	"ᴳ": "G",
	"ᴴ": "H",
	"ᴵ": "I",
	"ᴶ": "J",
	"ᴷ": "K",
	"ᴸ": "L",
	"ᴹ": "M",
	"ᴺ": "N",
	"ᴻ": "N",
	"ᴼ": "O",
	"ᴾ": "P",
	"ᴿ": "R",
	"ᵀ": "T",
	"ᵁ": "U",
	"ᵂ": "W",

	"ʝ": "j",
	"ʞ": "k",
	"ʟ": "L",
	"ʠ": "q",
	"ʦ": "ts",
	"ʧ": "tf",
	"ʨ": "tc",
	"ʩ": "fn",
	"ʪ": "ls",

	"ʒ": "3",
	"ʓ": "3",
	"ʗ": "C",
	"ʘ": "0",
	"ʙ": "B",
	"ʚ": "a",
	"ʛ": "G",
	"ʜ": "H",

	"ʆ": "l",
	"ʇ": "t",
	"ʈ": "t",
	"ʋ": "v",
	"ʍ": "m",
	"ʎ": "h",
	"ʏ": "Y",
	"ʐ": "z",
	"ʑ": "z",
	"ʀ": "R",
	"ʁ": "R",
	"ʂ": "s",
	"ɲ": "n",
	"ɳ": "n",
	"ɴ": "N",

	"ɓ": "b",
	"ɔ": "c",

	"ɠ": "g",
	"ɡ": "g",
	"ɢ": "G",

	"ɵ": "o",

	"ɖ": "d",
	"ɗ": "d",

	"ɘ": "e",

	"ɕ": "c",

	"ɦ": "h",
	"ɧ": "h",

	"ɭ": "l",

	"ɩ": "i",
	"ɪ": "I",
	"ɫ": "l",
	"ɬ": "l",
	"ɨ": "i",

	"ɼ": "r",
	"ɽ": "r",

	"і": "i",
	"Ї": "I",
	"є": "e",

	"⒜ ⒝ ⒞ ⒟ ⒠ ⒡ ⒢ ⒣ ⒤ ⒥ ⒦ ⒧ ⒨ ⒩ ⒪ ⒫ ⒬ ⒭ ⒮ ⒯ ⒰ ⒱ ⒲ ⒳ ⒴": "a b c d e f g h i j k l m n o p q r s t u v w x y",

	"Ⓩⓐⓑⓒⓓⓔⓕⓖⓗⓘⓙⓚⓛⓜⓝⓞⓟⓠⓡⓢⓣⓤⓥⓦⓧⓨⓩ⓪": "Zabcdefghijklmnopqrstuvwxyz0",

	"𝕒𝕓𝕔𝕕𝕖𝕗𝕘𝕙𝕚𝕛𝕜𝕝𝕞𝕟𝕠𝕡𝕢𝕣𝕤𝕥𝕦𝕧𝕨𝕩𝕪𝕫": "abcdefghijklmnopqrstuvwxyz",

	"🄰🄱🄲🄳🄴🄵🄶🄷🄸🄹🄺🄻🄼🄽🄾🄿🅀🅁🅂🅃🅄🅅🅆🅇🅈🅉": "ABCDEFGHIJKLMNOPQRSTUVWXYZ",

	"₳฿₵ĐɆ₣₲ⱧłJ₭Ⱡ₥₦Ø₱QⱤ₴₮ɄV₩ӾɎⱫ": "ABCDEFGHIJKLMNOPQRSTUVWXYZ",

	"𝖆𝖇𝖈𝖉𝖊𝖋𝖌𝖍𝖎𝖏𝖐𝖑𝖒𝖓𝖔𝖕𝖖𝖗𝖘𝖙𝖚𝖛𝖜𝖝𝖞𝖟": "abcdefghijklmnopqrstuvwxyz",

	"①": "1",
	"②": "2",
	"③": "3",
	"④": "4",
	"⑤": "5",
	"⑥": "6",
	"⑦": "7",
	"⑧": "8",
	"⑨": "9",

	"⑴": "1",
	"⑵": "2",
	"⑶": "3",
	"⑷": "4",
	"⑸": "5",
	"⑹": "6",
	"⑺": "7",
	"⑻": "8",
	"⑼": "9",

	"⑩": "10",
	"⑪": "11",
	"⑫": "12",
	"⑬": "13",
	"⑭": "14",
	"⑮": "15",
	"⑯": "16",
	"⑰": "17",
	"⑱": "18",
	"⑲": "19",
	"⑳": "20",

	"⑽": "10",
	"⑾": "11",
	"⑿": "12",
	"⒀": "13",
	"⒁": "14",
	"⒂": "15",
	"⒃": "16",
	"⒄": "17",
	"⒅": "18",
	"⒆": "19",
	"⒇": "20",

	"ⓐⓑⓒⓓⓔⓕⓖⓗⓘⓙⓚⓛⓜⓝⓞⓟⓠⓡⓢⓣⓤⓥⓦⓧⓨⓩ⓪①②③④⑤⑥⑦⑧⑨⓪": "abcdefghijklmnopqrstuvwxyz01234567890",

	"ᴀʙᴄᴅᴇғɢʜɪᴊᴋʟᴍɴᴏᴘǫʀ":                    "ABCDEFGHIJKLMNOPQR",
	"ᴛᴜᴠᴡxʏᴢ𝟶𝟷𝟸𝟹𝟺𝟻𝟼𝟽𝟾𝟿𝟶":                    "TUVWXYZ01234567890",
	"𝕒𝕓𝕔𝕕𝕖𝕗𝕘𝕙𝕚𝕛𝕜𝕝𝕞𝕟𝕠𝕡𝕢𝕣𝕤𝕥𝕦𝕧𝕨𝕩𝕪𝕫𝟘𝟙𝟚𝟛𝟜𝟝𝟞𝟟𝟠𝟡𝟘": "abcdefghijklmnopqrstuvwxyz01234567890",

	"𝔞𝔟𝔠𝔡𝔢𝔣𝔤𝔥𝔦𝔧𝔨𝔩𝔪𝔫𝔬𝔭𝔮𝔯𝔰𝔱𝔲𝔳𝔴𝔵𝔶𝔷": "abcdefghijklmnopqrstuvwxyz",

	"𝐚𝐛𝐜𝐝𝐞𝐟𝐠𝐡𝐢𝐣𝐤𝐥𝐦𝐧𝐨𝐩𝐪𝐫𝐬𝐭𝐮𝐯𝐰𝐱𝐲𝐳𝟎𝟏𝟐𝟑𝟒𝟓𝟔𝟕𝟖𝟗𝟎": "abcdefghijklmnopqrstuvwxyz01234567890",

	// ₳฿₵ĐɆ₣₲ⱧłJ₭Ⱡ₥₦Ø₱QⱤ₴₮ɄV₩ӾɎⱫ

	// αв¢∂єfgнιנкℓмиσρqяѕтυνωχуz

	// αвcdєfghíjklmnσpqrstuvwхчz

	"𝓪𝓫𝓬𝓭𝓮𝓯𝓰𝓱𝓲𝓳𝓴𝓵𝓶𝓷𝓸𝓹𝓺𝓻𝓼𝓽𝓾𝓿𝔀𝔁𝔂𝔃": "abcdefghijklmnopqrstuvwxyz",

	"𝒶𝒷𝒸𝒹𝑒𝒻𝑔𝒽𝒾𝒿𝓀𝓁𝓂𝓃𝑜𝓅𝓆𝓇𝓈𝓉𝓊𝓋𝓌𝓍𝓎𝓏": "abcdefghijklmnopqrstuvwxyz",

	// αႦƈԃҽϝɠԋιʝƙʅɱɳσρϙɾʂƚυʋɯxყȥ

	"𝘢𝘣𝘤𝘥𝘦𝘧𝘨𝘩𝘪𝘫𝘬𝘭𝘮𝘯𝘰𝘱𝘲𝘳𝘴𝘵𝘶𝘷𝘸𝘹𝘺𝘻": "abcdefghijklmnopqrstuvwxyz",

	"𝙖𝙗𝙘𝙙𝙚𝙛𝙜𝙝𝙞𝙟𝙠𝙡𝙢𝙣𝙤𝙥𝙦𝙧𝙨𝙩𝙪𝙫𝙬𝙭𝙮𝙯": "abcdefghijklmnopqrstuvwxyz",

	"ᗩᗷᑕᗪEᖴGᕼIᒍKᒪᗰᑎOᑭᑫᖇᔕTᑌᐯᗯ᙭Yᘔ": "ABCDEFGHIJKLMNOPQRSTUVWXYZ",

	"🅰🅱🅲🅳🅴🅵🅶🅷🅸🅹🅺🅻🅼🅽🅾🅿🆀🆁🆂🆃🆄🆅🆆🆇🆈🆉": "ABCDEFGHIJKLMNOPQRSTUVWXYZ",

	"𝚊𝚋𝚌𝚍𝚎𝚏𝚐𝚑𝚒𝚓𝚔𝚕𝚖𝚗𝚘𝚙𝚚𝚛𝚜𝚝𝚞𝚟𝚠𝚡𝚢𝚣𝟶𝟷𝟸𝟹𝟺𝟻𝟼𝟽𝟾𝟿𝟶": "abcdefghijklmnopqrstuvwxyz01234567890",

	// a̶b̶c̶d̶e̶f̶g̶h̶i̶j̶k̶l̶m̶n̶o̶p̶q̶r̶s̶t̶u̶v̶w̶x̶y̶z̶0̶1̶2̶3̶4̶5̶6̶7̶8̶9̶0̶

	// a̴b̴c̴d̴e̴f̴g̴h̴i̴j̴k̴l̴m̴n̴o̴p̴q̴r̴s̴t̴u̴v̴w̴x̴y̴z̴0̴1̴2̴3̴4̴5̴6̴7̴8̴9̴0̴

	// a̷b̷c̷d̷e̷f̷g̷h̷i̷j̷k̷l̷m̷n̷o̷p̷q̷r̷s̷t̷u̷v̷w̷x̷y̷z̷0̷1̷2̷3̷4̷5̷6̷7̷8̷9̷0̷

	// a̲b̲c̲d̲e̲f̲g̲h̲i̲j̲k̲l̲m̲n̲o̲p̲q̲r̲s̲t̲u̲v̲w̲x̲y̲z̲0̲1̲2̲3̲4̲5̲6̲7̲8̲9̲0̲

	// a̳b̳c̳d̳e̳f̳g̳h̳i̳j̳k̳l̳m̳n̳o̳p̳q̳r̳s̳t̳u̳v̳w̳x̳y̳z̳0̳1̳2̳3̳4̳5̳6̳7̳8̳9̳0̳

	// [̲̅a̲̅][̲̅b̲̅][̲̅c̲̅][̲̅d̲̅][̲̅e̲̅][̲̅f̲̅][̲̅g̲̅][̲̅h̲̅][̲̅i̲̅][̲̅j̲̅][̲̅k̲̅][̲̅l̲̅][̲̅m̲̅][̲̅n̲̅][̲̅o̲̅][̲̅p̲̅][̲̅q̲̅][̲̅r̲̅][̲̅s̲̅][̲̅t̲̅][̲̅u̲̅][̲̅v̲̅][̲̅w̲̅][̲̅x̲̅][̲̅y̲̅][̲̅z̲̅][̲̅0̲̅][̲̅1̲̅][̲̅2̲̅][̲̅3̲̅][̲̅4̲̅][̲̅5̲̅][̲̅6̲̅][̲̅7̲̅][̲̅8̲̅][̲̅9̲̅][̲̅0̲̅]

	"ａｂｃｄｅｆｇｈｉｊｋｌｍｎｏｐｑｒｓｔｕｖｗｘｙｚ": "abcdefghijklmnopqrstuvwxyz",
	"０１２３４５６７８９０":                "01234567890",

	// ᴀʙᴄᴅᴇғɢʜɪᴊᴋʟᴍɴᴏᴘϙʀꜱᴛᴜᴠᴡxʏᴢ01234567890

	// ɐbɔdǝɟƃɥıɾʞlɯnodbɹsʇnʌʍxʎz01234567890
}

var itests = map[string]string{
	"🇧🇱 🇺 🇪 🇧🇪 🇷 🇷 🇾":     "BL U E BE R R Y",
	"⁰ ¹ ² ³ ⁴ ⁵ ⁶ ⁷ ⁸ ⁹": "0 1 2 3 4 5 6 7 8 9",
	"⁺ ⁻ ⁼ ⁽ ⁾":           "+ - = ( )",
	"₀ ₁ ₂ ₃ ₄ ₅ ₆ ₇ ₈ ₉": "0 1 2 3 4 5 6 7 8 9",
	"₊ ₋ ₌ ₍ ₎":           "+ - = ( )",

	"ᵃ ᵇ ᶜ ᵈ ᵉ ᶠ ᵍ ʰ ⁱ ʲ ᵏ ˡ ᵐ ⁿ ᵒ ᵖ ʳ ˢ ᵗ ᵘ ᵛ ʷ ˣ ʸ ᶻ": "a b c d e f g h i j k l m n o p r s t u v w x y z",
	"ₐ ₑ ₕ ᵢ ⱼ ₖ ₗ ₘ ₙ ₒ ₚ ᵣ ₛ ₜ ᵤ ᵥ ₓ":                 "a e h i j k l m n o p r s t u v x",

	"ᴬ ᴮ ᴰ ᴱ ᴳ ᴴ ᴵ ᴶ ᴷ ᴸ ᴹ ᴺ ᴼ ᴾ ᴿ ᵀ ᵁ ⱽ ᵂ": "A B D E G H I J K L M N O P R T U V W",

	"a̶b̶c̶d̶e̶f̶g̶h̶i̶j̶k̶l̶m̶n̶o̶p̶q̶r̶s̶t̶u̶v̶w̶x̶y̶z": "abcdefghijklmnopqrstuvwxyz",
}

func main() {
	i := 0
	for in, result := range tests {
		out, err := normalize.Normalize(in)
		if err != nil {
			panic(err)
		}
		if out != result {
			fmt.Printf("(%s) failed - '%s' != '%s' \t%s:%s\n", in, out, result, in, result)
			inRunes := []rune(in)
			resultRunes := []rune(result)
			if len(inRunes) == len(resultRunes) {
				for i := 0; i < len(inRunes); i++ {
					fmt.Printf("%c:%c\n", inRunes[i], resultRunes[i])
				}
			}
		}
		i++
	}

	for in, result := range itests {
		out, err := normalize.Normalize(in)
		if err != nil {
			panic(err)
		}

		out = strings.ToLower(out)
		result = strings.ToLower(result)
		if out != result {
			fmt.Printf("(%s) failed - '%s' != '%s' %s:%s\n", in, out, result, in, result)
		}
		i++
	}

	in := "🅱🅡🅘🅓🅖🅔"
	out, _ := normalize.Normalize(in)
	log.Println(out)
}
